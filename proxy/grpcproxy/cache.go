// Copyright 2024 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package grpcproxy

import (
	lru "github.com/hashicorp/golang-lru/v2"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/metadata"
)

// cacheKey is the key for caching a GRPC client connection.
type cacheKey struct {
	host               string
	port               string
	connectionSecurity string
}

// CacheKeyFn returns a key for caching a GRPC client connection.
//
// The returned key is a tuple of (connectionSecurity, host, port).
//
// The connectionSecurity should be a unique identifier for the type of security used for the GRPC
// connection. For example, "insecure", "tls-skip-verify", "mtls-zatar", "mtls-custom".
//
// The host is the host name of the GRPC server.
//
// The port is the port number of the GRPC server.
type CacheKeyFn func(md metadata.MD, dest string) (connectionSecurity string, host string, port string)

// CacheEvictFn returns true if the cache entry for that GRPC client connection should be evicted.
type CacheEvictFn func(md metadata.MD) bool

func (p *proxy) initCache() error {
	if p.cacheKeyFn == nil {
		p.cacheKeyFn = func(metadata.MD, string) (string, string, string) { return "", "", "" }
	}
	if p.cacheEvictFn == nil {
		p.cacheEvictFn = func(metadata.MD) bool { return false }
	}
	if p.cacheSize > 0 {
		cache, err := lru.NewWithEvict[cacheKey, *grpc.ClientConn](p.cacheSize, func(key cacheKey, conn *grpc.ClientConn) {
			conn.Close()
		})
		if err != nil {
			return err
		}
		p.cache = cache
	}
	return nil
}

func (p *proxy) connFromCache(md metadata.MD, dest string) (*grpc.ClientConn, *cacheKey, error) {
	if p.cache == nil {
		return nil, nil, nil // no cache
	}
	cs, host, port := p.cacheKeyFn(md, dest)
	if cs == "" || host == "" || port == "" {
		return nil, nil, nil // no cache
	}

	// if cache evict, remove all cache entries for that host.
	if p.cacheEvictFn(md) {
		for _, k := range p.cache.Keys() {
			if k.host == host {
				p.cache.Remove(k)
			}
		}
	}
	key := &cacheKey{host: host, port: port, connectionSecurity: cs}
	if conn, ok := p.cache.Get(*key); ok {
		if conn.GetState() == connectivity.Shutdown {
			// connectivity state is shutdown, remove from cache and retry.
			p.cache.Remove(*key)
			return nil, key, nil
		}
		return conn, key, nil
	}
	return nil, key, nil
}

func (p *proxy) addToCache(conn *grpc.ClientConn, ck *cacheKey) (added bool) {
	if p.cache == nil || ck == nil {
		return false // no cache
	}
	p.cache.Add(*ck, conn)
	return true
}
