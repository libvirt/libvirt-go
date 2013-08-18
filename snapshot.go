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
	// "unsafe"
)

type VirDomainSnapshot struct {
	ptr C.virDomainSnapshotPtr
}

func (d *VirDomain) CreateSnapshotXML(xml string, flags uint32) (VirDomainSnapshot, error) {
	cXml := C.CString(xml)
	// defer C.free(unsafe.Pointer(cXml))
	result := C.virDomainSnapshotCreateXML(d.ptr, cXml, C.uint(flags))
	if result == nil {
		return VirDomainSnapshot{}, errors.New(GetLastError())
	}
	return VirDomainSnapshot{ptr: result}, nil
}
