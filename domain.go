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

const (
	VIR_DOMAIN_RUNNING = C.VIR_DOMAIN_RUNNING
	VIR_DOMAIN_SHUTOFF = C.VIR_DOMAIN_SHUTOFF
)

type VirDomain struct {
	ptr C.virDomainPtr
}

type VirDomainInfo struct {
	ptr C.virDomainInfo
}

func (d *VirDomain) Create() error {
	result := C.virDomainCreate(d.ptr)
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (d *VirDomain) Destroy() error {
	result := C.virDomainDestroy(d.ptr)
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (d *VirDomain) Shutdown() error {
	result := C.virDomainShutdown(d.ptr)
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (d *VirDomain) Reboot(flags uint) error {
	result := C.virDomainReboot(d.ptr, C.uint(flags))
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (d *VirDomain) IsActive() (bool, error) {
	result := C.virDomainIsActive(d.ptr)
	if result == -1 {
		return false, errors.New(GetLastError())
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (d *VirDomain) SetAutostart(autostart bool) error {
	var cAutostart C.int
	switch autostart {
	case true:
		cAutostart = 1
	default:
		cAutostart = 0
	}
	result := C.virDomainSetAutostart(d.ptr, cAutostart)
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (d *VirDomain) GetAutostart() (bool, error) {
	var out C.int
	result := C.virDomainGetAutostart(d.ptr, (*C.int)(unsafe.Pointer(&out)))
	if result == -1 {
		return false, errors.New(GetLastError())
	}
	switch out {
	case 1:
		return true, nil
	default:
		return false, nil
	}
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

func (d *VirDomain) GetInfo() (VirDomainInfo, error) {
	di := VirDomainInfo{}
	var ptr C.virDomainInfo
	result := C.virDomainGetInfo(d.ptr, (*C.virDomainInfo)(unsafe.Pointer(&ptr)))
	if result == -1 {
		return di, errors.New(GetLastError())
	}
	di.ptr = ptr
	return di, nil
}

func (d *VirDomain) GetXMLDesc(flags uint32) (string, error) {
	result := C.virDomainGetXMLDesc(d.ptr, C.uint(flags))
	if result == nil {
		return "", errors.New(GetLastError())
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

func (i *VirDomainInfo) GetState() uint8 {
	return uint8(i.ptr.state)
}

func (i *VirDomainInfo) GetMaxMem() uint64 {
	return uint64(i.ptr.maxMem)
}

func (i *VirDomainInfo) GetMemory() uint64 {
	return uint64(i.ptr.memory)
}

func (i *VirDomainInfo) GetNrVirtCpu() uint16 {
	return uint16(i.ptr.nrVirtCpu)
}

func (i *VirDomainInfo) GetCpuTime() uint64 {
	return uint64(i.ptr.cpuTime)
}
