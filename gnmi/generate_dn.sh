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

DN_VERSION="dev_v19_2"

#git clone git@github.com:drivenets/cheetah.git --branch $DN_VERSION
#wget https://raw.githubusercontent.com/openconfig/gnmi/master/metadata/yang/gnmi-collector-metadata.yang
#git clone https://github.com/open-traffic-generator/models-yang.git

EXCLUDE_MODULES=ietf-interfaces,openconfig-bfd,openconfig-messages,rpc

YANG_FILES=(
# /dn/yangs/dn-isis.yang
# /dn/yangs/dn-ldp-oper.yang
# /dn/yangs/dn-sys-telnet.yang
# /dn/yangs/dn-qos.yang
# /dn/yangs/dn-vpws.yang
# /dn/yangs/dn-debug-segment-routing.yang
# /dn/yangs/dn-isis-oper.yang
# /dn/yangs/dn-sys-moninode.yang
# /dn/yangs/dn-debug-isis.yang
# /dn/yangs/dn-forwarding-options.yang
# /dn/yangs/dn-srv-mpls-oam.yang
# /dn/yangs/dn-community-list.yang
# /dn/yangs/dn-sys-event-manager.yang
# /dn/yangs/dn-srv-flow-monitoring.yang
# /dn/yangs/dn-sys-alarms.yang
# /dn/yangs/dn-snmp-common.yang
# /dn/yangs/dn-debug-lacp.yang
# /dn/yangs/dn-extensions.yang
# /dn/yangs/dn-sys-aaa.yang
# /dn/yangs/dn-mfib.yang
# /dn/yangs/dn-debug-rsvp.yang
# /dn/yangs/dn-bgp-oper.yang
# /dn/yangs/dn-multicast.yang
# /dn/yangs/dn-bfd.yang
# /dn/yangs/dn-tracking-policy.yang
# /dn/yangs/dn-mrib.yang
# /dn/yangs/dn-interfaces.yang
# /dn/yangs/dn-sys-ntp.yang
# /dn/yangs/dn-types.yang
# /dn/yangs/dn-if-ethernet.yang
# /dn/yangs/dn-services.yang
# /dn/yangs/dn-if-flowspec.yang
# /dn/yangs/dn-acl-ip-protocols.yang
# /dn/yangs/dn-routing-policy.yang
# /dn/yangs/dn-sys-logging.yang
# /dn/yangs/dn-srv-service-activation-testing.yang
# /dn/yangs/dn-nat.yang
# /dn/yangs/dn-if-breakout.yang
# /dn/yangs/dn-ldp.yang
# /dn/yangs/dn-nacm.yang
# /dn/yangs/dn-protocol.yang
# /dn/yangs/dn-debug-ospf.yang
# /dn/yangs/dn-debug-ldp.yang
# /dn/yangs/dn-multihoming.yang
# /dn/yangs/dn-evpn-vpws.yang
# /dn/yangs/dn-prefix-list.yang
# /dn/yangs/dn-dhcp-relay-agent.yang
# /dn/yangs/dn-srv-stamp.yang
# /dn/yangs/dn-ospf.yang
# /dn/yangs/dn-debug-msdp.yang
# /dn/yangs/dn-rib.yang
# /dn/yangs/dn-sys-ncm.yang
# /dn/yangs/dn-access-control-list.yang
# /dn/yangs/dn-sr-flex-algo.yang
# /dn/yangs/dn-evpn-all-common.yang
# /dn/yangs/dn-sys-resources.yang
# /dn/yangs/dn-vrf.yang
# /dn/yangs/dn-network-services.yang
# /dn/yangs/dn-oob-mgmt-route.yang
# /dn/yangs/dn-bmp.yang
# /dn/yangs/dn-if-l2nat.yang
# /dn/yangs/dn-vlan.yang
# /dn/yangs/dn-pim.yang
# /dn/yangs/dn-clock.yang
# /dn/yangs/dn-platform-security.yang
# /dn/yangs/dn-sys-ncc.yang
# /dn/yangs/dn-interfaces-common.yang
# /dn/yangs/dn-sys-debug.yang
# /dn/yangs/dn-igmp.yang
# /dn/yangs/dn-bgp.yang
# /dn/yangs/dn-high-availability.yang
# /dn/yangs/dn-if-nfcommon.yang
# /dn/yangs/dn-sys-ldap.yang
# /dn/yangs/dn-optical-transport.yang
# /dn/yangs/dn-rsvp.yang
# /dn/yangs/dn-if-ncs-ethernet.yang
# /dn/yangs/dn-transceivers.yang
# /dn/yangs/dn-performance-monitoring.yang
# /dn/yangs/dn-debug-fib-manager.yang
# /dn/yangs/dn-srv-connectivity-fault-management.yang
# /dn/yangs/dn-flowspec-local-policies.yang
# /dn/yangs/dn-routing-options.yang
# /dn/yangs/dn-bridge-domain.yang
# /dn/yangs/dn-sys-cprl.yang
# /dn/yangs/dn-platform.yang
# /dn/yangs/dn-fib.yang
# /dn/yangs/dn-srv-ethernet-oam.yang
# /dn/yangs/dn-evpn.yang
# /dn/yangs/dn-mldp.yang
# /dn/yangs/dn-srv-connectivity-fault-management-types.yang
# /dn/yangs/dn-load-balancing.yang
# /dn/yangs/dn-if-ip.yang
# /dn/yangs/dn-sys-install.yang
# /dn/yangs/dn-srv-link-fault-management.yang
# /dn/yangs/dn-debug-pim.yang
# /dn/yangs/dn-hw-model.yang
# /dn/yangs/dn-static-route.yang
# /dn/yangs/dn-sys-ncp.yang
# /dn/yangs/dn-srv-l2xc.yang
# /dn/yangs/dn-sys-ssh.yang
# /dn/yangs/dn-snmp.yang
# /dn/yangs/dn-sys-https.yang
# /dn/yangs/dn-debug-vrrp.yang
# /dn/yangs/dn-sys-connectivity.yang
# /dn/yangs/dn-debug-bgp.yang
# /dn/yangs/dn-sys-login.yang
# /dn/yangs/dn-policy.yang
# /dn/yangs/dn-srv-twamp.yang
# /dn/yangs/dn-if-bundle.yang
# /dn/yangs/dn-if-urpf.yang
# /dn/yangs/dn-debug-arp.yang
# /dn/yangs/dn-sys-sztp.yang
# /dn/yangs/dn-debug-rib-manager.yang
# /dn/yangs/dn-sys-netconf.yang
# /dn/yangs/dn-as-path-access-list.yang
# /dn/yangs/dn-if-ipsec.yang
# /dn/yangs/dn-sys-backplane.yang
# /dn/yangs/dn-traffic-engineering.yang
# /dn/yangs/dn-sys-ncf.yang
# /dn/yangs/dn-ipsec.yang
# /dn/yangs/dn-segment-routing.yang
# /dn/yangs/dn-msdp.yang
# /dn/yangs/dn-system.yang
# /dn/yangs/dn-sys-processes.yang
# /dn/yangs/dn-rpki.yang
# /dn/yangs/dn-ptp.yang
# /dn/yangs/dn-sys-containers.yang
# /dn/yangs/dn-sys-network-resources.yang
# /dn/yangs/dn-timezones.yang
# /dn/yangs/dn-vrrp.yang
# /dn/yangs/dn-debug-mpls-oam.yang
# /dn/yangs/dn-lldp.yang
# /dn/yangs/dn-srv-port-mirroring.yang
# /dn/yangs/dn-if-gre-tunnel.yang
# /dn/yangs/dn-sys-ncs.yang
# /dn/yangs/dn-qppb-policy.yang
# /dn/yangs/dn-sys-dns.yang
# /dn/yangs/dn-sys-ftp.yang
# /dn/yangs/dn-sys-diagnostics.yang
# /dn/yangs/dn-lacp.yang
# /dn/yangs/dn-packet-fields.yang
# /dn/yangs/dn-mpls.yang
# /dn/yangs/dn-sys-pmtud.yang
# /dn/yangs/dn-evpn-vpws-fxc.yang
# /dn/yangs/dn-vlan-types.yang
/dn/yangs/dn-top.yang
# /dn/yangs/dn-debug.yang
# /dn/yangs/dn-sys-fragmentation.yang
# /dn/yangs/dn-transport-types.yang
)


go run github.com/openconfig/ygnmi/app/ygnmi generator \
  --typedef_enum_with_defmod=false \
  --exclude_modules="${EXCLUDE_MODULES}" \
  --base_package_path=github.com/openconfig/ondatra/gnmi/oc \
  --output_dir=gnmi/dn \
  --ignore_deviate_notsupported \
  "${YANG_FILES[@]}"

# go run github.com/openconfig/ygnmi/app/ygnmi generator \
#   --trim_module_prefix=open-traffic-generator \
#   --base_package_path=github.com/openconfig/ondatra/gnmi/otg \
#   --output_dir=gnmi/otg \
#   --paths=models-yang/models/... \
#   --generate_atomic_lists=false \
#   "${OTG_YANG_FILES[@]}"

find gnmi -name "*.go" -exec goimports -w {} +
find gnmi -name "*.go" -exec gofmt -w -s {} +

#rm -rf cheetah
