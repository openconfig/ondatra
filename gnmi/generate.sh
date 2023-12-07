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

git clone https://github.com/openconfig/public.git
wget https://raw.githubusercontent.com/openconfig/gnmi/master/metadata/yang/gnmi-collector-metadata.yang
git clone https://github.com/open-traffic-generator/models-yang.git
git clone https://github.com/openconfig/gnsi.git
# TODO(team): Remove the following line once the new gNSI version is ready for use
pushd gnsi && git checkout c2ebf3e && popd

EXCLUDE_MODULES=ietf-interfaces,openconfig-bfd,openconfig-messages

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
  public/release/models/aft/openconfig-aft-summary.yang
  public/release/models/ate/openconfig-ate-flow.yang
  public/release/models/ate/openconfig-ate-intf.yang
  public/release/models/bfd/openconfig-bfd.yang
  public/release/models/bgp/openconfig-bgp-policy.yang
  public/release/models/bgp/openconfig-bgp-types.yang
  public/release/models/extensions/openconfig-metadata.yang
  public/release/models/interfaces/openconfig-if-aggregate.yang
  public/release/models/interfaces/openconfig-if-ethernet.yang
  public/release/models/interfaces/openconfig-if-ethernet-ext.yang
  public/release/models/interfaces/openconfig-if-ip-ext.yang
  public/release/models/interfaces/openconfig-if-ip.yang
  public/release/models/interfaces/openconfig-if-sdn-ext.yang
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
  public/release/models/optical-transport/openconfig-terminal-device.yang
  public/release/models/optical-transport/openconfig-transport-types.yang
  public/release/models/ospf/openconfig-ospfv2.yang
  public/release/models/ospf/openconfig-ospf-policy.yang
  public/release/models/p4rt/openconfig-p4rt.yang
  public/release/models/platform/openconfig-platform-controller-card.yang
  public/release/models/platform/openconfig-platform-cpu.yang
  public/release/models/platform/openconfig-platform-ext.yang
  public/release/models/platform/openconfig-platform-fabric.yang
  public/release/models/platform/openconfig-platform-fan.yang
  public/release/models/platform/openconfig-platform-integrated-circuit.yang
  public/release/models/platform/openconfig-platform-linecard.yang
  public/release/models/platform/openconfig-platform-pipeline-counters.yang
  public/release/models/platform/openconfig-platform-psu.yang
  public/release/models/platform/openconfig-platform-software.yang
  public/release/models/platform/openconfig-platform-transceiver.yang
  public/release/models/platform/openconfig-platform.yang
  public/release/models/platform/openconfig-platform-common.yang
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
  public/release/models/system/openconfig-system-bootz.yang
  public/release/models/system/openconfig-system-controlplane.yang
  public/release/models/system/openconfig-system-utilization.yang
  public/release/models/types/openconfig-inet-types.yang
  public/release/models/types/openconfig-types.yang
  public/release/models/types/openconfig-yang-types.yang
  public/release/models/vlan/openconfig-vlan.yang
  public/third_party/ietf/iana-if-type.yang
  public/third_party/ietf/ietf-inet-types.yang
  public/third_party/ietf/ietf-interfaces.yang
  public/third_party/ietf/ietf-yang-types.yang
)

OTG_YANG_FILES=(
  models-yang/models/bgp/open-traffic-generator-bgp.yang
  models-yang/models/discovery/open-traffic-generator-discovery.yang
  models-yang/models/flow/open-traffic-generator-flow.yang
  models-yang/models/interface/open-traffic-generator-port.yang
  models-yang/models/isis/open-traffic-generator-isis.yang
  models-yang/models/lacp/open-traffic-generator-lacp.yang
  models-yang/models/lag/open-traffic-generator-lag.yang
  models-yang/models/lldp/open-traffic-generator-lldp.yang
  models-yang/models/rsvp/open-traffic-generator-rsvp.yang
  models-yang/models/types/open-traffic-generator-types.yang
)

go run github.com/openconfig/ygnmi/app/ygnmi generator \
  --trim_module_prefix=openconfig \
  --typedef_enum_with_defmod=false \
  --exclude_modules="${EXCLUDE_MODULES}" \
  --base_package_path=github.com/openconfig/ondatra/gnmi/oc \
  --output_dir=gnmi/oc \
  --paths=public/release/models/...,public/third_party/ietf/... \
  --split_package_paths="/network-instances/network-instance/protocols/protocol/isis=netinstisis,/network-instances/network-instance/protocols/protocol/bgp=netinstbgp" \
  --ignore_deviate_notsupported \
  "${YANG_FILES[@]}"

go run github.com/openconfig/ygnmi/app/ygnmi generator \
  --trim_module_prefix=open-traffic-generator \
  --base_package_path=github.com/openconfig/ondatra/gnmi/otg \
  --output_dir=gnmi/otg \
  --paths=models-yang/models/... \
  --generate_atomic_lists=false \
  "${OTG_YANG_FILES[@]}"

find gnmi -name "*.go" -exec goimports -w {} +
find gnmi -name "*.go" -exec gofmt -w -s {} +

rm -rf public gnmi-collector-metadata.yang models-yang gnsi
