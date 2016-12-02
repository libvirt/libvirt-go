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

type SecretUsageType int

const (
	VIR_SECRET_USAGE_TYPE_NONE   = SecretUsageType(C.VIR_SECRET_USAGE_TYPE_NONE)
	VIR_SECRET_USAGE_TYPE_VOLUME = SecretUsageType(C.VIR_SECRET_USAGE_TYPE_VOLUME)
	VIR_SECRET_USAGE_TYPE_CEPH   = SecretUsageType(C.VIR_SECRET_USAGE_TYPE_CEPH)
	VIR_SECRET_USAGE_TYPE_ISCSI  = SecretUsageType(C.VIR_SECRET_USAGE_TYPE_ISCSI)
	VIR_SECRET_USAGE_TYPE_TLS    = SecretUsageType(C.VIR_SECRET_USAGE_TYPE_TLS)
)

type Secret struct {
	ptr C.virSecretPtr
}

func (s *Secret) Free() error {
	if result := C.virSecretFree(s.ptr); result != 0 {
		return GetLastError()
	}
	s.ptr = nil
	return nil
}

func (s *Secret) Undefine() error {
	result := C.virSecretUndefine(s.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (s *Secret) GetUUID() ([]byte, error) {
	var cUuid [C.VIR_UUID_BUFLEN](byte)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virSecretGetUUID(s.ptr, (*C.uchar)(cuidPtr))
	if result != 0 {
		return []byte{}, GetLastError()
	}
	return C.GoBytes(cuidPtr, C.VIR_UUID_BUFLEN), nil
}

func (s *Secret) GetUUIDString() (string, error) {
	var cUuid [C.VIR_UUID_STRING_BUFLEN](C.char)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virSecretGetUUIDString(s.ptr, (*C.char)(cuidPtr))
	if result != 0 {
		return "", GetLastError()
	}
	return C.GoString((*C.char)(cuidPtr)), nil
}

func (s *Secret) GetUsageID() (string, error) {
	result := C.virSecretGetUsageID(s.ptr)
	if result == nil {
		return "", GetLastError()
	}
	return C.GoString(result), nil
}

func (s *Secret) GetUsageType() (SecretUsageType, error) {
	result := SecretUsageType(C.virSecretGetUsageType(s.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (s *Secret) GetXMLDesc(flags uint32) (string, error) {
	result := C.virSecretGetXMLDesc(s.ptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

func (s *Secret) GetValue(flags uint32) ([]byte, error) {
	var cvalue_size C.size_t

	cvalue := C.virSecretGetValue(s.ptr, &cvalue_size, C.uint(flags))
	if cvalue == nil {
		return nil, GetLastError()
	}
	defer C.free(cvalue)
	ret := C.GoBytes(unsafe.Pointer(cvalue), C.int(cvalue_size))
	return ret, nil
}

func (s *Secret) SetValue(value []byte, flags uint32) error {
	var cvalue_size C.size_t = C.size_t(len(value))
	var cvalue *C.uchar = (*C.uchar)(C.CBytes(value))

	defer C.free(cvalue)

	result := C.virSecretSetValue(s.ptr, cvalue, cvalue_size, C.uint(flags))

	if result == -1 {
		return GetLastError()
	}

	return nil
}
