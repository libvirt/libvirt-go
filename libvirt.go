package libvirt

import (
	"bytes"
	"encoding/binary"
	"errors"
	"unsafe"
)

/*
#cgo LDFLAGS: -lvirt -ldl
#include <libvirt/libvirt.h>
*/
import "C"

type VirConnection struct {
	connection _Ctype_virConnectPtr
}

func NewVirConnection(uri string) (VirConnection, error) {
	ptr := C.virConnectOpen(C.CString(uri))
	if ptr == nil {
		return VirConnection{}, errors.New("Failed to connect to hypervisor")
	}
	obj := VirConnection{connection: ptr}
	return obj, nil
}

func (c *VirConnection) ListDomains() ([]uint32, error) {
	domainIds := make([]int, 1024)
	domainIdsPtr := unsafe.Pointer(&domainIds)
	numDomains := C.virConnectListDomains(c.connection, (*C.int)(domainIdsPtr), 1024)
	if numDomains == -1 {
		return nil, errors.New("Failed to list domains")
	}

	domains := make([]uint32, numDomains)

	gBytes := C.GoBytes(domainIdsPtr, C.int(numDomains*32))
	buf := bytes.NewBuffer(gBytes)
	for k := 0; k < int(numDomains); k++ {
		binary.Read(buf, binary.LittleEndian, &domains[k])
	}
	return domains, nil
}
