package ondatra

import (
	"context"
	"fmt"
	"testing"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	gpb "github.com/openconfig/gnmi/proto/gnmi"
	"github.com/openconfig/ondatra"
	"github.com/openconfig/ondatra/binding"
	"github.com/openconfig/ondatra/fakebind"
	"github.com/openconfig/ondatra/gnmi"
	"github.com/openconfig/ondatra/gnmi/oc"
	"github.com/openconfig/ondatra/internal/display"
	opb "github.com/openconfig/ondatra/proto"
	"github.com/openconfig/ygnmi/ygnmi"
	"github.com/openconfig/ygot/ygot"
)

type DnBinding struct {
	binding.Binding
}

func DnDialGNMIInsecure(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	display.Action(display.MainT, "DnDevice.DialGNMI")
	//addr := "localhost:50051"
	// addr := "10.1.1.103:50051"
	// 100.64.7.159/20
	addr := "100.64.7.159:50051"
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// 5 second timeout:
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, err
	}
	return gpb.NewGNMIClient(conn), nil
}

func DnDialGNMISecure(ctx context.Context, opts ...grpc.DialOption) (gpb.GNMIClient, error) {
	/* Connects to the gNMI server using the TLS credentials and the login credentials.
	 * The TLS certificate and user/pass are read from the .ini file
	 */
	display.Action(display.MainT, "DnDialGNMISecure2")
	cfg, err := GetDnConfig()
	if err != nil {
		fmt.Println("Failed to get DnConfig: ", err)
		panic(err)
	}

	certFile := cfg.TLS.Certificate
	addr := cfg.Cluster.Host + ":" + cfg.Cluster.Port
	var userpass loginCreds = loginCreds{
		Username: cfg.Cluster.Username,
		Password: cfg.Cluster.Password,
	}
	// func NewTLSConfig(ca, cert, key, clientAuth string, skipVerify, genSelfSigned bool) (*tls.Config, error) {
	tls_config, err := NewTLSConfig("", certFile, "", "", true, false)
	if err != nil {
		fmt.Println("Failed to create TLS credentials: ", err)
		panic(err)
	}
	opts = append(opts, grpc.WithTransportCredentials(credentials.NewTLS(tls_config)))
	opts = append(opts, grpc.WithPerRPCCredentials(userpass))

	// 5 second timeout:
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	conn, err := grpc.DialContext(ctx, addr, opts...)
	if err != nil {
		return nil, err
	}
	return gpb.NewGNMIClient(conn), nil

}

func (b *DnBinding) Reserve(ctx context.Context, tb *opb.Testbed, runTime, waitTime time.Duration, partial map[string]string) (*binding.Reservation, error) {
	display.Action(display.MainT, "DnBinding.Reserve")
	fmt.Printf("%v %v %v %v %v\n", ctx, tb, runTime, waitTime, partial)
	dev := &fakebind.DUT{
		AbstractDUT: &binding.AbstractDUT{
			Dims: &binding.Dims{
				Name:   "fakeDUT",
				Vendor: opb.Device_DriveNets,
			},
		},
		DialGNMIFn: DnDialGNMISecure,
		//DialGNMIFn: DnDialGNMIInsecure,
	}
	return &binding.Reservation{
		ID: "DrivenetsReserveId",
		DUTs: map[string]binding.DUT{
			"dn_dut": dev,
		},
	}, nil
}

func (b *DnBinding) Release(ctx context.Context) error {
	display.Action(display.MainT, "DnBinding.Release")
	return nil
}
func (b *DnBinding) FetchReservation(ctx context.Context, id string) (*binding.Reservation, error) {
	display.Action(display.MainT, "DnBinding.FetchReservation")
	return &binding.Reservation{}, nil
}

func DnGetBinding() (binding.Binding, error) {
	display.Action(display.MainT, "DnGetBinding")
	return &DnBinding{}, nil
}

func TestMain(m *testing.M) {
	display.Action(display.MainT, "TestMain")
	ondatra.RunTests(m, DnGetBinding)
}

func TestDnGnmiGet(t *testing.T) {
	display.Action(display.MainT, "TestDnGnmiGet")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dut := ondatra.DUT(t, "dn_dut")
	fmt.Printf("DUT: %v\n", dut)

	ds := gnmi.Get(t, dut, gnmi.OC().Interface("bundle-1").OperStatus().State())
	fmt.Println("Got status: ", ds)
}

func TestDnGnmiBasicSet(t *testing.T) {
	display.Action(display.MainT, "TestDnGnmiBasicSet")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dut := ondatra.DUT(t, "dn_dut")
	fmt.Printf("DUT: %v\n", dut)

	cofig_container := gnmi.OC().Interface("bundle-2").Config()
	// val3 := "3"
	// val2 := oc.Interface_Ethernet{
	// 	AggregateId: &val3,
	// }
	val := oc.Interface{
		AdminStatus: oc.Interface_AdminStatus_UP,
		//Ethernet:    &val2,
	}
	gnmi.Update(t, dut, cofig_container, &val)

	ds := gnmi.Get(t, dut, gnmi.OC().Interface("bundle-2").Config())
	fmt.Printf("Read status from server: \n%+v\n", ds)

	if ds.AdminStatus != oc.Interface_AdminStatus_UP {
		t.Errorf("Get(DUT port1 status): got %v, want %v", ds, oc.Interface_AdminStatus_UP)
	}
	if ds.Tpid != oc.VlanTypes_TPID_TYPES_TPID_0X8100 {
		t.Errorf("Get(DUT port1 status): got %v, want %v", ds, oc.VlanTypes_TPID_TYPES_TPID_0X8100)
	}
}

func TestDnGnmiIdentityRefSet(t *testing.T) {
	// /oc-if:interfaces/oc-if:interface/oc-if:state/oc-vlan:tpid
	display.Action(display.MainT, "TestDnGnmiIdentityRefSet")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dut := ondatra.DUT(t, "dn_dut")
	fmt.Printf("DUT: %v\n", dut)

	cofig_container := gnmi.OC().Interface("bundle-3").Config()
	val := oc.Interface{
		Tpid: oc.VlanTypes_TPID_TYPES_TPID_0X8100,
		Type: oc.IETFInterfaces_InterfaceType_ieee8023adLag,
	}
	gnmi.Update(t, dut, cofig_container, &val)

}

func TestDnGnmiMultiSet(t *testing.T) {
	// /oc-if:interfaces/oc-if:interface/oc-if:state/oc-vlan:tpid
	display.Action(display.MainT, "TestDnGnmiIdentityRefSet")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dut := ondatra.DUT(t, "dn_dut")
	fmt.Printf("DUT: %v\n", dut)

	cofig_container := gnmi.OC().Interface("bundle-4").Config()
	// p1 := "port-1"
	// mtu := uint16(70)
	access_vlan := uint16(71)
	native_vlan := uint16(72)
	// min_links := uint16(73)

	lagspeed := uint32(3)
	members := []string{"port-1", "port-2"}
	trunk_vlans := []oc.Interface_Aggregation_SwitchedVlan_TrunkVlans_Union{
		// oc.To_Interface_Aggregation_SwitchedVlan_TrunkVlans_Union(uint16(16)),
	}
	switched_vlan := oc.Interface_Aggregation_SwitchedVlan{
		AccessVlan:    &access_vlan,
		InterfaceMode: oc.Vlan_VlanModeType_TRUNK,
		NativeVlan:    &native_vlan,
		TrunkVlans:    trunk_vlans,
	}
	agg_container := oc.Interface_Aggregation{
		LagSpeed: &lagspeed,
		LagType:  oc.IfAggregate_AggregationType_LACP,
		Member:   members,
		// MinLinks:     &min_links,
		SwitchedVlan: &switched_vlan,
	}
	val := oc.Interface{
		AdminStatus: oc.Interface_AdminStatus_UP,
		Tpid:        oc.VlanTypes_TPID_TYPES_TPID_0X8100,
		// HardwarePort: &p1,
		// Mtu: &mtu,
		Type:        oc.IETFInterfaces_InterfaceType_ieee8023adLag,
		Aggregation: &agg_container,
	}
	fmt.Printf("val: %v\n", val)
	gnmi.Update(t, dut, cofig_container, &val)

}

func TestDnGnmiLeafListSet(t *testing.T) {
	// /oc-if:interfaces/oc-if:interface/oc-if:state/oc-vlan:tpid
	display.Action(display.MainT, "TestDnGnmiIdentityRefSet")
	_, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	dut := ondatra.DUT(t, "dn_dut")
	fmt.Printf("DUT: %v\n", dut)

	cofig_container := gnmi.OC().Interface("bundle-5").Config()
	// p1 := "port-1"
	// mtu := uint16(70)
	access_vlan := uint16(71)
	native_vlan := uint16(72)
	// min_links := uint16(73)

	lagspeed := uint32(3)
	members := []string{"port-1", "port-2"}
	trunk_vlans := []oc.Interface_Aggregation_SwitchedVlan_TrunkVlans_Union{
		// oc.To_Interface_Aggregation_SwitchedVlan_TrunkVlans_Union(uint16(16)),
	}
	switched_vlan := oc.Interface_Aggregation_SwitchedVlan{
		AccessVlan:    &access_vlan,
		InterfaceMode: oc.Vlan_VlanModeType_TRUNK,
		NativeVlan:    &native_vlan,
		TrunkVlans:    trunk_vlans,
	}
	agg_container := oc.Interface_Aggregation{
		LagSpeed: &lagspeed,
		LagType:  oc.IfAggregate_AggregationType_LACP,
		Member:   members,
		// MinLinks:     &min_links,
		SwitchedVlan: &switched_vlan,
	}
	val := oc.Interface{
		AdminStatus: oc.Interface_AdminStatus_UP,
		Tpid:        oc.VlanTypes_TPID_TYPES_TPID_0X8100,
		// HardwarePort: &p1,
		// Mtu: &mtu,
		// Type:        oc.IETFInterfaces_InterfaceType_aal5,
		Aggregation: &agg_container,
	}
	gnmi.Update(t, dut, cofig_container, &val)

}

/////////////////////////////////////////////

func TestDnGnmiMultipleUpdates(t *testing.T) {
	dut := ondatra.DUT(t, "dn_dut")
	sb := &gnmi.SetBatch{}
	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Description().Config(), "Multiple Updates Test")
	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Mtu().Config(), uint16(1500))
	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Enabled().Config(), true)
	sb.Set(t, dut)
	// if err := sb.Set(t, dut).RawResponse.Response(); err != nil {
	// 	t.Fatalf("Multiple updates failed: %v", err)
	// }
}

func TestDnGnmiUpdateAndDelete(t *testing.T) {
	dut := ondatra.DUT(t, "dn_dut")
	sb := &gnmi.SetBatch{}
	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Description().Config(), "Update and Delete Test")
	gnmi.BatchDelete(sb, gnmi.OC().Interface("eth1").Mtu().Config())
	sb.Set(t, dut)
	// if err := sb.Set(t, dut).Error(); err != nil {
	// 	t.Fatalf("Update and delete failed: %v", err)
	// }
}

func TestDnGnmiReplaceOperation(t *testing.T) {
	dut := ondatra.DUT(t, "dn_dut")
	config := gnmi.OC().Interface("eth1").Config()
	intf := &oc.Interface{
		Name:        ygot.String("eth1"),
		Description: ygot.String("Update Test"),
		Type:        oc.IETFInterfaces_InterfaceType_ethernetCsmacd,
		Enabled:     ygot.Bool(true),
		Mtu:         ygot.Uint16(9000),
	}
	gnmi.Update(t, dut, config, intf)
}

func TestDnGnmiUpdateThenReplace(t *testing.T) {
	dut := ondatra.DUT(t, "dn_dut")
	sb := &gnmi.SetBatch{}
	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Description().Config(), "Update Then Replace - Update")
	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Config(), &oc.Interface{
		Name:        ygot.String("eth1"),
		Description: ygot.String("Update Then Replace - Replace"),
		Enabled:     ygot.Bool(true),
	})
	sb.Set(t, dut)
	// if err := sb.Set(t, dut).Error(); err != nil {
	// 	t.Fatalf("Update then replace failed: %v", err)
	// }
}

func TestDnGnmiDeleteThenUpdate(t *testing.T) {
	dut := ondatra.DUT(t, "dn_dut")
	sb := &gnmi.SetBatch{}
	gnmi.BatchDelete(sb, gnmi.OC().Interface("eth1").Description().Config())
	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Description().Config(), "Delete Then Update")
	sb.Set(t, dut)
	// if err := sb.Set(t, dut).Error(); err != nil {
	// 	t.Fatalf("Delete then update failed: %v", err)
	// }
}

// func TestDnGnmiInvalidPath(t *testing.T) {
// 	dut := ondatra.DUT(t, "dn_dut")
// 	invalidPath := gnmi.OC().Interface("eth1").Config().AugmentedConfig("invalid")
// 	err := gnmi.Replace(t, dut, invalidPath, "Invalid Path Test")
// 	if err == nil {
// 		t.Fatalf("Expected error for invalid path, got nil")
// 	}
// }

// func TestDnGnmiInvalidValue(t *testing.T) {
// 	dut := ondatra.DUT(t, "dn_dut")
// 	err := gnmi.Replace(t, dut, gnmi.OC().Interface("eth1").Mtu().Config(), uint16(65536)) // Assuming 65536 is an invalid MTU
// 	if err == nil {
// 		t.Fatalf("Expected error for invalid value, got nil")
// 	}
// }

// func TestDnGnmiMixedValidInvalid(t *testing.T) {
// 	dut := ondatra.DUT(t, "dn_dut")
// 	sb := &gnmi.SetBatch{}
// 	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Description().Config(), "Valid Update")
// 	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Mtu().Config(), uint16(65536)) // Invalid MTU
// 	err := sb.Set(t, dut).Error()
// 	if err == nil {
// 		t.Fatalf("Expected error for mixed valid/invalid operations, got nil")
// 	}
// }

func TestDnGnmiSetWithPrefixPath(t *testing.T) {
	dut := ondatra.DUT(t, "dn_dut")
	prefix := gnmi.OC().Interface("eth1")
	sb := &gnmi.SetBatch{}
	gnmi.BatchUpdate(sb, prefix.Description().Config(), "Prefix Path Test")
	gnmi.BatchUpdate(sb, prefix.Enabled().Config(), true)
	sb.Set(t, dut)
	// if err := sb.Set(t, dut).Error(); err != nil {
	// 	t.Fatalf("Set with prefix path failed: %v", err)
	// }
}

// func TestDnGnmiSetReadOnlyField(t *testing.T) {
// 	dut := ondatra.DUT(t, "dn_dut")
// 	err := gnmi.Replace(t, dut, gnmi.OC().Interface("eth1").Counters().InOctets().State(), uint64(1000))
// 	if err == nil {
// 		t.Fatalf("Expected error when setting read-only field, got nil")
// 	}
// }

func TestDnGnmiDeleteWithWildcards(t *testing.T) {
	dut := ondatra.DUT(t, "dn_dut")
	gnmi.Delete(t, dut, gnmi.OC().Interface("...").Config())

	// if err != nil {
	// 	t.Fatalf("Delete with wildcards failed: %v", err)
	// }
}

// func TestDnGnmiPartialFailure(t *testing.T) {
// 	dut := ondatra.DUT(t, "dn_dut")
// 	sb := &gnmi.SetBatch{}
// 	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Description().Config(), "Partial Failure Test")
// 	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Mtu().Config(), uint64(65536)) // Invalid MTU
// 	err := sb.Set(t, dut)

// }

func TestDnGnmiSetAndSubscribe(t *testing.T) {
	dut := ondatra.DUT(t, "dn_dut")
	intfPath := gnmi.OC().Interface("eth1")

	// Start subscription
	watch := gnmi.Watch(t, dut, intfPath.Description().State(), 30*time.Second, func(val *ygnmi.Value[string]) bool {
		return val.IsPresent()
	})

	// Perform Set operation
	gnmi.Update(t, dut, intfPath.Description().Config(), "Set and Subscribe Test")

	// Wait for subscription update
	if _, ok := watch.Await(t); !ok {
		t.Fatalf("Did not receive subscription update")
	}
}

func TestDnGnmiSetWithTypedValue(t *testing.T) {
	dut := ondatra.DUT(t, "dn_dut")
	sb := &gnmi.SetBatch{}
	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Description().Config(), "String Value") // string_val
	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Mtu().Config(), uint16(1500))           // uint_val
	gnmi.BatchUpdate(sb, gnmi.OC().Interface("eth1").Enabled().Config(), true)               // bool_val
	sb.Set(t, dut)
}

func aTestDnGnmiSetEmptyUpdateValue(t *testing.T) {
	dut := ondatra.DUT(t, "dn_dut")
	gnmi.Update(t, dut, gnmi.OC().Interface("eth1").Description().Config(), "")
}

func aTestDnGnmiSetDuplicatePaths(t *testing.T) {
	dut := ondatra.DUT(t, "dn_dut")
	sb := &gnmi.SetBatch{}
	gnmi.BatchUpdate(sb, gnmi.OC().Interface("bundle-6").Description().Config(), "First Update")
	gnmi.BatchUpdate(sb, gnmi.OC().Interface("bundle-6").Description().Config(), "Second Update")
	sb.Set(t, dut)

	// Verify the last update was applied
	desc := gnmi.Get(t, dut, gnmi.OC().Interface("bundle-6").Description().State())
	if desc != "Second Update" {
		t.Fatalf("Expected description 'Second Update', got '%s'", desc)
	}
}
