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
