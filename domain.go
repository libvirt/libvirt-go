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

type VirDomain struct {
	ptr _Ctype_virDomainPtr
}

func (d *VirDomain) GetName() (string, error) {
	name := C.virDomainGetName(d.ptr)
	if name == nil {
		return "", errors.New(GetLastError())
	}
	return C.GoString(name), nil
}

func (d *VirDomain) GetState() ([]int, error) {
	var cState C.int
	var cReason C.int
	result := C.virDomainGetState(d.ptr,
		(*C.int)(unsafe.Pointer(&cState)),
		(*C.int)(unsafe.Pointer(&cReason)),
		0)
	if int(result) == -1 {
		return []int{}, errors.New(GetLastError())
	}
	return []int{int(cState), int(cReason)}, nil
}

func (d *VirDomain) GetUUID() ([]byte, error) {
	var cUuid [C.VIR_UUID_BUFLEN](byte)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virDomainGetUUID(d.ptr, (*C.uchar)(cuidPtr))
	if result != 0 {
		return []byte{}, errors.New(GetLastError())
	}
	return C.GoBytes(cuidPtr, C.VIR_UUID_BUFLEN), nil
}

func (d *VirDomain) GetUUIDString() (string, error) {
	var cUuid [C.VIR_UUID_STRING_BUFLEN](C.char)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virDomainGetUUIDString(d.ptr, (*C.char)(cuidPtr))
	if result != 0 {
		return "", errors.New(GetLastError())
	}
	return C.GoString((*C.char)(cuidPtr)), nil
}
