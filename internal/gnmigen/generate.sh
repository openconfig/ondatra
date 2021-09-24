#!/bin/bash
#
# Copyright 2021 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     https://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

# This script is used to generate the Ondatra Telemetry and Config Go APIs.

set -e

go install github.com/openconfig/ygot/generator@latest
git clone https://github.com/openconfig/public.git
wget https://raw.githubusercontent.com/openconfig/gnmi/master/metadata/yang/gnmi-collector-metadata.yang

EXCLUDE_MODULES=ietf-interfaces,openconfig-acl,openconfig-bfd,openconfig-messages

COMMON_ARGS=(
  -path=public/release/models,public/third_party/ietf
  -generate_path_structs
  -compress_paths
  -exclude_modules="${EXCLUDE_MODULES}"
  -generate_fakeroot
  -fakeroot_name=device
  -ignore_shadow_schema_paths
  -generate_simple_unions
  -shorten_enum_leaf_names
  -typedef_enum_with_defmod
  -enum_suffix_for_simple_union_enums
  -trim_enum_openconfig_prefix
  -include_schema
  -generate_append
  -generate_getters
  -generate_rename
  -generate_delete
  -generate_leaf_getters
  -structs_split_files_count=10
)

YANG_FILES=(
  public/release/models/openconfig-extensions.yang
  public/release/models/acl/openconfig-packet-match.yang
  public/release/models/aft/openconfig-aft.yang
  public/release/models/bfd/openconfig-bfd.yang
  public/release/models/bgp/openconfig-bgp-types.yang
  public/release/models/interfaces/openconfig-if-aggregate.yang
  public/release/models/interfaces/openconfig-if-ethernet.yang
  public/release/models/interfaces/openconfig-if-ip.yang
  public/release/models/interfaces/openconfig-if-ip-ext.yang
  public/release/models/interfaces/openconfig-interfaces.yang
  public/release/models/isis/openconfig-isis.yang
  public/release/models/lacp/openconfig-lacp.yang
  public/release/models/lldp/openconfig-lldp.yang
  public/release/models/lldp/openconfig-lldp-types.yang
  public/release/models/local-routing/openconfig-local-routing.yang
  public/release/models/mpls/openconfig-mpls-types.yang
  public/release/models/multicast/openconfig-pim.yang
  public/release/models/network-instance/openconfig-network-instance.yang
  public/release/models/optical-transport/openconfig-transport-types.yang
  public/release/models/ospf/openconfig-ospfv2.yang
  public/release/models/platform/openconfig-platform-cpu.yang
  public/release/models/platform/openconfig-platform-software.yang
  public/release/models/platform/openconfig-platform-transceiver.yang
  public/release/models/platform/openconfig-platform.yang
  public/release/models/policy/openconfig-policy-types.yang
  public/release/models/policy-forwarding/openconfig-policy-forwarding.yang
  public/release/models/qos/openconfig-qos.yang
  public/release/models/qos/openconfig-qos-elements.yang
  public/release/models/qos/openconfig-qos-interfaces.yang
  public/release/models/qos/openconfig-qos-types.yang
  public/release/models/rib/openconfig-rib-bgp.yang
  public/release/models/segment-routing/openconfig-segment-routing-types.yang
  public/release/models/system/openconfig-system.yang
  public/release/models/types/openconfig-inet-types.yang
  public/release/models/types/openconfig-types.yang
  public/release/models/types/openconfig-yang-types.yang
  public/release/models/vlan/openconfig-vlan.yang
  public/third_party/ietf/iana-if-type.yang
  public/third_party/ietf/ietf-inet-types.yang
  public/third_party/ietf/ietf-interfaces.yang
  public/third_party/ietf/ietf-yang-types.yang
  gnmi-collector-metadata.yang
)

# Generate Config API.
mkdir -p config
generator \
  -generate_structs=false \
  -exclude_state \
  -schema_struct_path=github.com/openconfig/ondatra/telemetry \
  -output_dir=config \
  -package_name=config \
  -path_structs_split_files_count=5 \
  "${COMMON_ARGS[@]}" \
  "${YANG_FILES[@]}"
go run internal/gnmigen/main/main.go \
  -output_dir=config \
  -package_name=config \
  -path=public/release/models,public/third_party/ietf \
  -exclude_modules="${EXCLUDE_MODULES}" \
  -generate_config_func \
  -prefer_shadow_path \
  -schema_struct_path=github.com/openconfig/ondatra/telemetry \
  -telemetry_funcs_file_split=10 \
  -telemetry_types_file_split=0 \
  "${YANG_FILES[@]}"

# Generate Telemetry API.
mkdir -p telemetry
generator \
  -generate_structs \
  -prefer_operational_state \
  -list_builder_key_threshold=4 \
  -output_dir=telemetry \
  -package_name=telemetry \
  -path_structs_split_files_count=10 \
  "${COMMON_ARGS[@]}" \
  "${YANG_FILES[@]}"
go run internal/gnmigen/main/main.go \
  -output_dir=telemetry \
  -package_name=telemetry \
  -path=public/release/models,public/third_party/ietf \
  -exclude_modules="${EXCLUDE_MODULES}" \
  -telemetry_funcs_file_split=40 \
  -telemetry_types_file_split=10 \
  "${YANG_FILES[@]}"

goimports -w config/*.go telemetry/*.go
gofmt -w -s config/*.go telemetry/*.go
rm -rf public gnmi-collector-metadata.yang
