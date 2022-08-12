#!/bin/bash
#
# Copyright 2021 Google LLC
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at:
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

git clone https://github.com/openconfig/gnsi.git
# TODO: Remove the following line once the new gNSI version is ready for use
pushd gnsi && git checkout c2ebf3e && popd
git clone https://github.com/openconfig/public.git
wget https://raw.githubusercontent.com/openconfig/gnmi/master/metadata/yang/gnmi-collector-metadata.yang
git clone https://github.com/open-traffic-generator/models-yang.git

EXCLUDE_MODULES=ietf-interfaces,openconfig-bfd,openconfig-messages

YGOT_COMMON_ARGS=(
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
  -generate_populate_defaults
  -list_builder_key_threshold=4
)

YANG_FILES=(
  gnmi-collector-metadata.yang
  gnsi/authz/gnsi-authz.yang
  gnsi/cert/gnsi-cert.yang
  gnsi/console/gnsi-console.yang
  gnsi/pathz/gnsi-pathz.yang
  gnsi/ssh/gnsi-ssh.yang
  public/release/models/acl/openconfig-acl.yang
  public/release/models/acl/openconfig-packet-match.yang
  public/release/models/aft/openconfig-aft.yang
  public/release/models/aft/openconfig-aft-network-instance.yang
  public/release/models/ate/openconfig-ate-flow.yang
  public/release/models/ate/openconfig-ate-intf.yang
  public/release/models/bfd/openconfig-bfd.yang
  public/release/models/bgp/openconfig-bgp-policy.yang
  public/release/models/bgp/openconfig-bgp-types.yang
  public/release/models/interfaces/openconfig-if-aggregate.yang
  public/release/models/interfaces/openconfig-if-ethernet.yang
  public/release/models/interfaces/openconfig-if-ethernet-ext.yang
  public/release/models/interfaces/openconfig-if-ip-ext.yang
  public/release/models/interfaces/openconfig-if-ip.yang
  public/release/models/interfaces/openconfig-interfaces.yang
  public/release/models/isis/openconfig-isis.yang
  public/release/models/lacp/openconfig-lacp.yang
  public/release/models/lldp/openconfig-lldp-types.yang
  public/release/models/lldp/openconfig-lldp.yang
  public/release/models/local-routing/openconfig-local-routing.yang
  public/release/models/mpls/openconfig-mpls-types.yang
  public/release/models/multicast/openconfig-pim.yang
  public/release/models/network-instance/openconfig-network-instance.yang
  public/release/models/openconfig-extensions.yang
  public/release/models/optical-transport/openconfig-transport-types.yang
  public/release/models/ospf/openconfig-ospfv2.yang
  public/release/models/p4rt/openconfig-p4rt.yang
  public/release/models/platform/openconfig-platform-cpu.yang
  public/release/models/platform/openconfig-platform-fan.yang
  public/release/models/platform/openconfig-platform-integrated-circuit.yang
  public/release/models/platform/openconfig-platform-software.yang
  public/release/models/platform/openconfig-platform-transceiver.yang
  public/release/models/platform/openconfig-platform.yang
  public/release/models/policy-forwarding/openconfig-policy-forwarding.yang
  public/release/models/policy/openconfig-policy-types.yang
  public/release/models/qos/openconfig-qos-elements.yang
  public/release/models/qos/openconfig-qos-interfaces.yang
  public/release/models/qos/openconfig-qos-types.yang
  public/release/models/qos/openconfig-qos.yang
  public/release/models/rib/openconfig-rib-bgp.yang
  public/release/models/sampling/openconfig-sampling-sflow.yang
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
)

# Generate Schema Structs
go run github.com/openconfig/ygot/generator \
  -path=public/release/models,public/third_party/ietf \
  -generate_structs \
  -generate_path_structs=false \
  -prefer_operational_state \
  -output_dir=telemetry \
  -package_name=telemetry \
  -structs_split_files_count=10 \
  "${YGOT_COMMON_ARGS[@]}" \
  "${YANG_FILES[@]}"

go run internal/gnmigen/main/main.go \
  -output_dir=telemetry \
  -package_name=telemetry \
  -path=public/release/models,public/third_party/ietf \
  -exclude_modules="${EXCLUDE_MODULES}" \
  -gen_path_struct_api=false \
  -split_pathstructs_by_module=true \
  -telemetry_types_file_split=10 \
  "${YANG_FILES[@]}"

# Generate Config API.
mkdir -p config/device
go run github.com/openconfig/ygot/generator \
  -path=public/release/models,public/third_party/ietf \
  -generate_path_structs \
  -generate_structs=false \
  -exclude_state \
  -schema_struct_path=github.com/openconfig/ondatra/telemetry \
  -output_dir=config \
  -package_name=device \
  -path_structs_split_files_count=1 \
  -split_pathstructs_by_module=true \
  -path_structs_output_file=config/device/device.go \
  -base_import_path=github.com/openconfig/ondatra/config \
  -trim_path_package_prefix="openconfig-" \
  -path_struct_package_suffix="" \
  "${YGOT_COMMON_ARGS[@]}" \
  "${YANG_FILES[@]}"

go run internal/gnmigen/main/main.go \
  -output_dir=config \
  -package_name=device \
  -path=public/release/models,public/third_party/ietf \
  -exclude_modules="${EXCLUDE_MODULES}" \
  -generate_config_func \
  -prefer_shadow_path \
  -schema_struct_path=github.com/openconfig/ondatra/telemetry \
  -split_pathstructs_by_module=true \
  -fake_root_helper_filename=device/root_helper.go \
  -fake_root_gnmi_filename=device/device_telem.go \
  -config_import_path=github.com/openconfig/ondatra/config \
  -telemetry_funcs_file_split=5 \
  -telemetry_types_file_split=0 \
  -trim_path_package_prefix="openconfig-" \
  "${YANG_FILES[@]}"

# Generate Telemetry API.
mkdir -p telemetry/device
go run github.com/openconfig/ygot/generator \
  -path=public/release/models,public/third_party/ietf \
  -generate_path_structs \
  -generate_structs=false \
  -prefer_operational_state \
  -output_dir=telemetry \
  -package_name=device \
  -path_structs_output_file=telemetry/device/device.go \
  -split_pathstructs_by_module=true \
  -schema_struct_path=github.com/openconfig/ondatra/telemetry \
  -trim_path_package_prefix="openconfig-" \
  -path_struct_package_suffix="" \
  -base_import_path=github.com/openconfig/ondatra/telemetry \
  -path_structs_split_files_count=2 \
  "${YGOT_COMMON_ARGS[@]}" \
  "${YANG_FILES[@]}"

go run internal/gnmigen/main/main.go \
  -output_dir=telemetry \
  -package_name=device \
  -path=public/release/models,public/third_party/ietf \
  -exclude_modules="${EXCLUDE_MODULES}" \
  -schema_struct_path=github.com/openconfig/ondatra/telemetry \
  -split_pathstructs_by_module=true \
  -fake_root_helper_filename=device/root_helper.go \
  -fake_root_gnmi_filename=device/device_telem.go \
  -telemetry_funcs_file_split=10 \
  -telemetry_types_file_split=10 \
  -trim_path_package_prefix="openconfig-" \
  "${YANG_FILES[@]}"

OTG_YANG_FILES=(
  models-yang/models/isis/open-traffic-generator-isis.yang
  models-yang/models/types/open-traffic-generator-types.yang
  models-yang/models/flow/open-traffic-generator-flow.yang
  models-yang/models/discovery/open-traffic-generator-discovery.yang
  models-yang/models/interface/open-traffic-generator-port.yang
  models-yang/models/bgp/open-traffic-generator-bgp.yang
)

# Generate OTG Schema Structs
mkdir -p telemetry/otg
go run github.com/openconfig/ygot/generator \
    -path=models-yang/models \
    -generate_structs \
    -generate_path_structs=false \
    -prefer_operational_state \
    -output_dir=telemetry/otg \
    -package_name=otg \
    -structs_split_files_count=3 \
    -trim_path_package_prefix="open-traffic-generator-" \
    "${YGOT_COMMON_ARGS[@]}" \
    "${OTG_YANG_FILES[@]}"

go run internal/gnmigen/main/main.go \
  -path=models-yang/models \
  -output_dir=telemetry/otg \
  -package_name=otg \
  -gen_path_struct_api=false \
  -split_pathstructs_by_module=true \
  -telemetry_types_file_split=3 \
  -trim_path_package_prefix="open-traffic-generator-" \
  "${OTG_YANG_FILES[@]}"

# Generate OTG Telemetry API
mkdir -p telemetry/otg/device
go run github.com/openconfig/ygot/generator \
    -path=models-yang/models \
    -generate_structs=false \
    -generate_path_structs \
    -prefer_operational_state \
    -output_dir=telemetry/otg \
    -package_name=device \
    -path_structs_output_file=telemetry/otg/device/device.go \
    -split_pathstructs_by_module=true \
    -schema_struct_path=github.com/openconfig/ondatra/telemetry/otg \
    -path_struct_package_suffix="" \
    -base_import_path=github.com/openconfig/ondatra/telemetry/otg \
    -path_structs_split_files_count=1 \
    -trim_path_package_prefix="open-traffic-generator-" \
    "${YGOT_COMMON_ARGS[@]}" \
    "${OTG_YANG_FILES[@]}"

go run internal/gnmigen/main/main.go \
  -path=models-yang/models \
  -output_dir=telemetry/otg \
  -package_name=device \
  -schema_struct_path=github.com/openconfig/ondatra/telemetry/otg \
  -split_pathstructs_by_module=true \
  -fake_root_helper_filename=device/root_helper.go \
  -fake_root_gnmi_filename=device/device_telem.go \
  -telemetry_funcs_file_split=3 \
  -trim_path_package_prefix="open-traffic-generator-" \
  "${OTG_YANG_FILES[@]}"

# New Telemetry Library
go run github.com/openconfig/ygnmi/app/ygnmi generator \
  --trim_module_prefix=openconfig \
  --exclude_modules="${EXCLUDE_MODULES}" \
  --base_package_path=github.com/openconfig/ondatra/gnmi/oc \
  --output_dir=gnmi/oc \
  --paths=public/release/models/...,public/third_party/ietf/... \
  "${YANG_FILES[@]}"

go run github.com/openconfig/ygnmi/app/ygnmi generator \
  --trim_module_prefix=open-traffic-generator \
  --base_package_path=github.com/openconfig/ondatra/gnmi/otg \
  --output_dir=gnmi/otg \
  --paths=models-yang/models/... \
  "${OTG_YANG_FILES[@]}"

find gnmi/oc gnmi/otg -name "*.go" -exec sed -i '1s/^/\/\/go:build go1.18\n/' {} +

find gnmi config telemetry -name "*.go" -exec goimports -w {} +
find gnmi config telemetry -name "*.go" -exec gofmt -w -s {} +

rm -rf public gnmi-collector-metadata.yang models-yang
