package libvirt

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"unsafe"
)

/*
#cgo LDFLAGS: -lvirt -ldl
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

type VirConnection struct {
	connection _Ctype_virConnectPtr
}

type VirDomain struct {
	domain _Ctype_virDomainPtr
}

func NewVirConnection(uri string) (VirConnection, error) {
	cUri := C.CString(uri)
	defer C.free(unsafe.Pointer(cUri))
	ptr := C.virConnectOpen(cUri)
	if ptr == nil {
		return VirConnection{}, errors.New(GetLastError())
	}
	obj := VirConnection{connection: ptr}
	return obj, nil
}

func GetLastError() string {
	err := C.virGetLastError()
	errMsg := fmt.Sprintf("[Code-%d] [Domain-%d] %s",
		err.code, err.domain, C.GoString(err.message))
	C.virResetError(err)
	return errMsg
}

func (c *VirConnection) CloseConnection() error {
	result := int(C.virConnectClose(c.connection))
	if result == -1 {
		return errors.New(GetLastError())
	}
	return nil
}

func (c *VirConnection) GetCapabilities() (string, error) {
	str := C.virConnectGetCapabilities(c.connection)
	if str == nil {
		return "", errors.New(GetLastError())
	}
	capabilities := C.GoString(str)
	C.free(unsafe.Pointer(str))
	return capabilities, nil
}

func (c *VirConnection) GetHostname() (string, error) {
	str := C.virConnectGetHostname(c.connection)
	if str == nil {
		return "", errors.New(GetLastError())
	}
	hostname := C.GoString(str)
	C.free(unsafe.Pointer(str))
	return hostname, nil
}

func (c *VirConnection) ListDefinedDomains() ([]string, error) {
	var names [1024](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numDomains := C.virConnectListDefinedDomains(
		c.connection,
		(**C.char)(namesPtr),
		1024)
	if numDomains == -1 {
		return nil, errors.New(GetLastError())
	}
	goNames := make([]string, numDomains)
	for k := 0; k < int(numDomains); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListDomains() ([]uint32, error) {
	domainIds := make([]int, 1024)
	domainIdsPtr := unsafe.Pointer(&domainIds)
	numDomains := C.virConnectListDomains(c.connection, (*C.int)(domainIdsPtr), 1024)
	if numDomains == -1 {
		return nil, errors.New(GetLastError())
	}

	domains := make([]uint32, numDomains)

	gBytes := C.GoBytes(domainIdsPtr, C.int(numDomains*32))
	buf := bytes.NewBuffer(gBytes)
	for k := 0; k < int(numDomains); k++ {
		binary.Read(buf, binary.LittleEndian, &domains[k])
	}
	return domains, nil
}

func (c *VirConnection) LookupDomainById(id uint32) (VirDomain, error) {
	ptr := C.virDomainLookupByID(c.connection, C.int(id))
	if ptr == nil {
		return VirDomain{}, errors.New(GetLastError())
	}
	return VirDomain{domain: ptr}, nil
}

func (c *VirConnection) LookupDomainByName(id string) (VirDomain, error) {
	cName := C.CString(id)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virDomainLookupByName(c.connection, cName)
	if ptr == nil {
		return VirDomain{}, errors.New(GetLastError())
	}
	return VirDomain{domain: ptr}, nil
}

func (d *VirDomain) GetName() (string, error) {
	name := C.virDomainGetName(d.domain)
	if name == nil {
		return "", errors.New(GetLastError())
	}
	return C.GoString(name), nil
}

func (d *VirDomain) GetState() ([]int, error) {
	var cState C.int
	var cReason C.int
	result := C.virDomainGetState(d.domain,
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
	result := C.virDomainGetUUID(d.domain, (*C.uchar)(cuidPtr))
	if result != 0 {
		return []byte{}, errors.New(GetLastError())
	}
	return C.GoBytes(cuidPtr, C.VIR_UUID_BUFLEN), nil
}

func (d *VirDomain) GetUUIDString() (string, error) {
	var cUuid [C.VIR_UUID_STRING_BUFLEN](C.char)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virDomainGetUUIDString(d.domain, (*C.char)(cuidPtr))
	if result != 0 {
		return "", errors.New(GetLastError())
	}
	return C.GoString((*C.char)(cuidPtr)), nil
}
