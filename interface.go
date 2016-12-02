package libvirt

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

type InterfaceXMLFlags int

const (
	INTERFACE_XML_INACTIVE = InterfaceXMLFlags(C.VIR_INTERFACE_XML_INACTIVE)
)

type Interface struct {
	ptr C.virInterfacePtr
}

func (n *Interface) Create(flags uint32) error {
	result := C.virInterfaceCreate(n.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *Interface) Destroy(flags uint32) error {
	result := C.virInterfaceDestroy(n.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *Interface) IsActive() (bool, error) {
	result := C.virInterfaceIsActive(n.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (n *Interface) GetMACString() (string, error) {
	result := C.virInterfaceGetMACString(n.ptr)
	if result == nil {
		return "", GetLastError()
	}
	mac := C.GoString(result)
	return mac, nil
}

func (n *Interface) GetName() (string, error) {
	result := C.virInterfaceGetName(n.ptr)
	if result == nil {
		return "", GetLastError()
	}
	name := C.GoString(result)
	return name, nil
}

func (n *Interface) GetXMLDesc(flags uint32) (string, error) {
	result := C.virInterfaceGetXMLDesc(n.ptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

func (n *Interface) Undefine() error {
	result := C.virInterfaceUndefine(n.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *Interface) Free() error {
	if result := C.virInterfaceFree(n.ptr); result != 0 {
		return GetLastError()
	}
	n.ptr = nil
	return nil
}
