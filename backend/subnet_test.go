package backend

import (
	"net"
	"testing"

	"github.com/digitalrebar/digitalrebar/go/common/store"
)

func TestSubnetCrud(t *testing.T) {
	bs := store.NewSimpleMemoryStore()
	dt := mkDT(bs)
	createTests := []crudTest{
		{"Create empty Subnet", dt.create, &Subnet{p: dt}, false},
		{"Create valid Subnet", dt.create, &Subnet{p: dt, Name: "test", Subnet: "192.168.124.0/24", ActiveStart: net.ParseIP("192.168.124.80"), ActiveEnd: net.ParseIP("192.168.124.254"), ActiveLeaseTime: 60, ReservedLeaseTime: 7200, Strategy: "mac"}, true},
		{"Create duplicate Subnet", dt.create, &Subnet{p: dt, Name: "test", Subnet: "192.168.124.0/24", ActiveStart: net.ParseIP("192.168.124.80"), ActiveEnd: net.ParseIP("192.168.124.254"), ActiveLeaseTime: 60, ReservedLeaseTime: 7200, Strategy: "mac"}, false},
		{"Create invalid Subnet(bad Subnet)", dt.create, &Subnet{p: dt, Name: "test2", Subnet: "127.0.0.0", ActiveStart: net.ParseIP("192.168.124.80"), ActiveEnd: net.ParseIP("192.168.124.254"), ActiveLeaseTime: 60, ReservedLeaseTime: 7200, Strategy: "mac"}, false},
		{"Create invalid Subnet(overlapping Subnet)", dt.create, &Subnet{p: dt, Name: "test2", Subnet: "192.168.124.0/24", ActiveStart: net.ParseIP("192.168.124.80"), ActiveEnd: net.ParseIP("192.168.124.254"), ActiveLeaseTime: 60, ReservedLeaseTime: 7200, Strategy: "mac"}, false},
		{"Create invalid Subnet(swapped Active range endpoints)", dt.create, &Subnet{p: dt, Name: "test2", Subnet: "192.168.125.0/24", ActiveStart: net.ParseIP("192.168.125.254"), ActiveEnd: net.ParseIP("192.168.125.80"), ActiveLeaseTime: 60, ReservedLeaseTime: 7200, Strategy: "mac"}, false},
		{"Create invalid Subnet(ActiveStart out of range)", dt.create, &Subnet{p: dt, Name: "test2", Subnet: "192.168.125.0/24", ActiveStart: net.ParseIP("192.168.124.80"), ActiveEnd: net.ParseIP("192.168.125.254"), ActiveLeaseTime: 60, ReservedLeaseTime: 7200, Strategy: "mac"}, false},
		{"Create invalid Subnet(ActiveEnd out of range)", dt.create, &Subnet{p: dt, Name: "test2", Subnet: "192.168.125.0/24", ActiveStart: net.ParseIP("192.168.125.80"), ActiveEnd: net.ParseIP("192.168.126.254"), ActiveLeaseTime: 60, ReservedLeaseTime: 7200, Strategy: "mac"}, false},
		{"Create invalid Subnet(ActiveLeaseTime too short)", dt.create, &Subnet{p: dt, Name: "test2", Subnet: "192.168.125.0/24", ActiveStart: net.ParseIP("192.168.125.80"), ActiveEnd: net.ParseIP("192.168.125.254"), ActiveLeaseTime: 59, ReservedLeaseTime: 7200, Strategy: "mac"}, false},
		{"Create invalid Subnet(ReservedLeaseTime too short)", dt.create, &Subnet{p: dt, Name: "test2", Subnet: "192.168.125.0/24", ActiveStart: net.ParseIP("192.168.125.80"), ActiveEnd: net.ParseIP("192.168.125.254"), ActiveLeaseTime: 60, ReservedLeaseTime: 7199, Strategy: "mac"}, false},
		{"Create invalid Subnet(no Strategy)", dt.create, &Subnet{p: dt, Name: "test2", Subnet: "192.168.125.0/24", ActiveStart: net.ParseIP("192.168.125.80"), ActiveEnd: net.ParseIP("192.168.125.254"), ActiveLeaseTime: 60, ReservedLeaseTime: 7200, Strategy: ""}, false},
	}
	for _, test := range createTests {
		test.Test(t)
	}
}
