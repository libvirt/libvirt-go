package libvirt

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

import (
	"reflect"
	"time"
	"unsafe"
)

type VirIPAddrType int

const (
	VIR_IP_ADDR_TYPE_IPV4 = VirIPAddrType(C.VIR_IP_ADDR_TYPE_IPV4)
	VIR_IP_ADDR_TYPE_IPV6 = VirIPAddrType(C.VIR_IP_ADDR_TYPE_IPV6)
)

type VirNetworkXMLFlags int

const (
	VIR_NETWORK_XML_INACTIVE = VirNetworkXMLFlags(C.VIR_NETWORK_XML_INACTIVE)
)

type VirNetworkUpdateCommand int

const (
	VIR_NETWORK_UPDATE_COMMAND_NONE      = VirNetworkUpdateCommand(C.VIR_NETWORK_UPDATE_COMMAND_NONE)
	VIR_NETWORK_UPDATE_COMMAND_MODIFY    = VirNetworkUpdateCommand(C.VIR_NETWORK_UPDATE_COMMAND_MODIFY)
	VIR_NETWORK_UPDATE_COMMAND_DELETE    = VirNetworkUpdateCommand(C.VIR_NETWORK_UPDATE_COMMAND_DELETE)
	VIR_NETWORK_UPDATE_COMMAND_ADD_LAST  = VirNetworkUpdateCommand(C.VIR_NETWORK_UPDATE_COMMAND_ADD_LAST)
	VIR_NETWORK_UPDATE_COMMAND_ADD_FIRST = VirNetworkUpdateCommand(C.VIR_NETWORK_UPDATE_COMMAND_ADD_FIRST)
)

type VirNetworkUpdateSection int

const (
	VIR_NETWORK_SECTION_NONE              = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_NONE)
	VIR_NETWORK_SECTION_BRIDGE            = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_BRIDGE)
	VIR_NETWORK_SECTION_DOMAIN            = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_DOMAIN)
	VIR_NETWORK_SECTION_IP                = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_IP)
	VIR_NETWORK_SECTION_IP_DHCP_HOST      = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_IP_DHCP_HOST)
	VIR_NETWORK_SECTION_IP_DHCP_RANGE     = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_IP_DHCP_RANGE)
	VIR_NETWORK_SECTION_FORWARD           = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_FORWARD)
	VIR_NETWORK_SECTION_FORWARD_INTERFACE = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_FORWARD_INTERFACE)
	VIR_NETWORK_SECTION_FORWARD_PF        = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_FORWARD_PF)
	VIR_NETWORK_SECTION_PORTGROUP         = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_PORTGROUP)
	VIR_NETWORK_SECTION_DNS_HOST          = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_DNS_HOST)
	VIR_NETWORK_SECTION_DNS_TXT           = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_DNS_TXT)
	VIR_NETWORK_SECTION_DNS_SRV           = VirNetworkUpdateSection(C.VIR_NETWORK_SECTION_DNS_SRV)
)

type VirNetworkUpdateFlags int

const (
	VIR_NETWORK_UPDATE_AFFECT_CURRENT = VirNetworkUpdateFlags(C.VIR_NETWORK_UPDATE_AFFECT_CURRENT)
	VIR_NETWORK_UPDATE_AFFECT_LIVE    = VirNetworkUpdateFlags(C.VIR_NETWORK_UPDATE_AFFECT_LIVE)
	VIR_NETWORK_UPDATE_AFFECT_CONFIG  = VirNetworkUpdateFlags(C.VIR_NETWORK_UPDATE_AFFECT_CONFIG)
)

type VirNetworkEventLifecycleType int

const (
	VIR_NETWORK_EVENT_DEFINED   = VirNetworkEventLifecycleType(C.VIR_NETWORK_EVENT_DEFINED)
	VIR_NETWORK_EVENT_UNDEFINED = VirNetworkEventLifecycleType(C.VIR_NETWORK_EVENT_UNDEFINED)
	VIR_NETWORK_EVENT_STARTED   = VirNetworkEventLifecycleType(C.VIR_NETWORK_EVENT_STARTED)
	VIR_NETWORK_EVENT_STOPPED   = VirNetworkEventLifecycleType(C.VIR_NETWORK_EVENT_STOPPED)
)

type VirNetworkEventID int

const (
	VIR_NETWORK_EVENT_ID_LIFECYCLE = VirNetworkEventID(C.VIR_NETWORK_EVENT_ID_LIFECYCLE)
)

type VirNetwork struct {
	ptr C.virNetworkPtr
}

func (n *VirNetwork) Free() error {
	if result := C.virNetworkFree(n.ptr); result != 0 {
		return GetLastError()
	}
	return nil
}

func (n *VirNetwork) Create() error {
	result := C.virNetworkCreate(n.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *VirNetwork) Destroy() error {
	result := C.virNetworkDestroy(n.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *VirNetwork) IsActive() (bool, error) {
	result := C.virNetworkIsActive(n.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (n *VirNetwork) IsPersistent() (bool, error) {
	result := C.virNetworkIsPersistent(n.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (n *VirNetwork) GetAutostart() (bool, error) {
	var out C.int
	result := C.virNetworkGetAutostart(n.ptr, (*C.int)(unsafe.Pointer(&out)))
	if result == -1 {
		return false, GetLastError()
	}
	switch out {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (n *VirNetwork) SetAutostart(autostart bool) error {
	var cAutostart C.int
	switch autostart {
	case true:
		cAutostart = 1
	default:
		cAutostart = 0
	}
	result := C.virNetworkSetAutostart(n.ptr, cAutostart)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *VirNetwork) GetName() (string, error) {
	name := C.virNetworkGetName(n.ptr)
	if name == nil {
		return "", GetLastError()
	}
	return C.GoString(name), nil
}

func (n *VirNetwork) GetUUID() ([]byte, error) {
	var cUuid [C.VIR_UUID_BUFLEN](byte)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virNetworkGetUUID(n.ptr, (*C.uchar)(cuidPtr))
	if result != 0 {
		return []byte{}, GetLastError()
	}
	return C.GoBytes(cuidPtr, C.VIR_UUID_BUFLEN), nil
}

func (n *VirNetwork) GetUUIDString() (string, error) {
	var cUuid [C.VIR_UUID_STRING_BUFLEN](C.char)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virNetworkGetUUIDString(n.ptr, (*C.char)(cuidPtr))
	if result != 0 {
		return "", GetLastError()
	}
	return C.GoString((*C.char)(cuidPtr)), nil
}

func (n *VirNetwork) GetBridgeName() (string, error) {
	result := C.virNetworkGetBridgeName(n.ptr)
	if result == nil {
		return "", GetLastError()
	}
	bridge := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return bridge, nil
}

func (n *VirNetwork) GetXMLDesc(flags uint32) (string, error) {
	result := C.virNetworkGetXMLDesc(n.ptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

func (n *VirNetwork) Undefine() error {
	result := C.virNetworkUndefine(n.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *VirNetwork) GetDHCPLeases() ([]VirNetworkDHCPLease, error) {
	var cLeases *C.virNetworkDHCPLeasePtr
	numLeases := C.virNetworkGetDHCPLeases(n.ptr, nil, (**C.virNetworkDHCPLeasePtr)(&cLeases), C.uint(0))
	if numLeases == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cLeases)),
		Len:  int(numLeases),
		Cap:  int(numLeases),
	}
	var leases []VirNetworkDHCPLease
	slice := *(*[]C.virNetworkDHCPLeasePtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		leases = append(leases, VirNetworkDHCPLease{ptr})
	}
	C.free(unsafe.Pointer(cLeases))
	return leases, nil
}

type VirNetworkDHCPLease struct {
	ptr C.virNetworkDHCPLeasePtr
}

func (l *VirNetworkDHCPLease) GetIface() string {
	return C.GoString(l.ptr.iface)
}

func (l *VirNetworkDHCPLease) GetExpiryTime() time.Time {
	return time.Unix(int64(l.ptr.expirytime), 0)
}

func (l *VirNetworkDHCPLease) GetMACAddress() string {
	return C.GoString(l.ptr.mac)
}

func (l *VirNetworkDHCPLease) GetIPAddress() string {
	return C.GoString(l.ptr.ipaddr)
}

func (l *VirNetworkDHCPLease) GetIPPrefix() uint {
	return uint(l.ptr.prefix)
}

func (l *VirNetworkDHCPLease) GetHostname() string {
	return C.GoString(l.ptr.hostname)
}

func (l *VirNetworkDHCPLease) GetClientID() string {
	return C.GoString(l.ptr.clientid)
}
