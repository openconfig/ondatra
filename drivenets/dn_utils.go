package ondatra

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

// DnConfig is the configuration struct for the Drivenets gNMI client.
// It is used to read the configuration from the file.
type DnConfig struct {
	Cluster struct {
		Port     string `yaml:"port"`
		Host     string `yaml:"host"`
		Username string `yaml:"username"`
		Password string `yaml:"password"`
	} `yaml:"cluster"`
	TLS struct {
		Certificate string `yaml:"certificate"`
	} `yaml:"TLS"`
}

// loginCreds implements the credentials.PerRPCCredentials interface.
// Used to pass the username and password to the server.
type loginCreds struct {
	Username, Password string
}

// GetRequestMetadata implements credentials.PerRPCCredentials.
func (c loginCreds) GetRequestMetadata(ctx context.Context, uri ...string) (map[string]string, error) {
	return map[string]string{
		"username": c.Username,
		"password": c.Password,
	}, nil
}

func (c loginCreds) RequireTransportSecurity() bool {
	return true
}

func NewTLSConfig(ca, cert, key, clientAuth string, skipVerify, genSelfSigned bool) (*tls.Config, error) {
	if !(skipVerify || ca != "" || (cert != "" && key != "")) {
		return nil, nil
	}

	tlsConfig := &tls.Config{
		InsecureSkipVerify: skipVerify,
	}

	// set clientAuth
	switch clientAuth {
	case "":
		if ca != "" {
			tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
		}
	case "request":
		tlsConfig.ClientAuth = tls.RequestClientCert
	case "require":
		tlsConfig.ClientAuth = tls.RequireAnyClientCert
	case "verify-if-given":
		tlsConfig.ClientAuth = tls.VerifyClientCertIfGiven
	case "require-verify":
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
	default:
		return nil, fmt.Errorf("unknown client-auth mode: %s", clientAuth)
	}
	if cert != "" && key != "" {
		panic("NOT IMPLEMENTED cert != '' && key != ''")
	} else if genSelfSigned {
		panic("NOT IMPLEMENTED genSelfSigned")
	}
	if ca != "" {
		certPool, err := LoadCACertificates(ca)
		if err != nil {
			return nil, err
		}
		tlsConfig.RootCAs = certPool
		tlsConfig.ClientCAs = certPool
	}
	return tlsConfig, nil
}

func LoadCACertificates(filePath string) (*x509.CertPool, error) {
	certPEMBlock, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("failed to read the cert file: %s: %w", filePath, err)
	}

	certPool := x509.NewCertPool()

	for {
		block, rest := pem.Decode(certPEMBlock)
		if block == nil {
			break
		}
		certPEMBlock = rest

		cert, err := x509.ParseCertificate(block.Bytes)
		if err != nil {
			return nil, fmt.Errorf("failed to parse certificate: %w", err)
		}

		if !cert.IsCA {
			return nil, fmt.Errorf("file %s contains a certificate that is not a CA", filePath)
		}
		certPool.AddCert(cert)
	}

	return certPool, nil
}

func GetDnConfig() (DnConfig, error) {
	var cfg DnConfig
	f, err := os.Open("test_conf.ini")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Read config: %v\n", cfg)
	return cfg, nil

}
