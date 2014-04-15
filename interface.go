package libvirt

/*
#cgo LDFLAGS: -lvirt -ldl
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

import (
	"errors"
	"unsafe"
)

type VirInterface struct {
	ptr C.virInterfacePtr
}

func (n *VirInterface) Create(flags uint32) error {
	result := C.virInterfaceCreate(n.ptr, C.uint(flags))
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (n *VirInterface) Destroy(flags uint32) error {
	result := C.virInterfaceDestroy(n.ptr, C.uint(flags))
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (n *VirInterface) IsActive() (bool, error) {
	result := C.virInterfaceIsActive(n.ptr)
	if result == -1 {
		return false, errors.New(GetLastError())
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (n *VirInterface) GetConnect() (VirConnection, error) {
	virConPtr := C.virInterfaceGetConnect(n.ptr)
	if virConPtr == nil {
		return VirConnection{}, errors.New(GetLastError())
	}
	return VirConnection{virConPtr}, nil
}

func (n *VirInterface) GetMACString() (string, error) {
	result := C.virInterfaceGetMACString(n.ptr)
	if result == nil {
		return "", errors.New(GetLastError())
	}
	mac := C.GoString(result)
	return mac, nil
}
