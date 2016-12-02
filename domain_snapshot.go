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
	"unsafe"
)

type VirDomainSnapshotCreateFlags int

const (
	VIR_DOMAIN_SNAPSHOT_CREATE_REDEFINE    = VirDomainSnapshotCreateFlags(C.VIR_DOMAIN_SNAPSHOT_CREATE_REDEFINE)
	VIR_DOMAIN_SNAPSHOT_CREATE_CURRENT     = VirDomainSnapshotCreateFlags(C.VIR_DOMAIN_SNAPSHOT_CREATE_CURRENT)
	VIR_DOMAIN_SNAPSHOT_CREATE_NO_METADATA = VirDomainSnapshotCreateFlags(C.VIR_DOMAIN_SNAPSHOT_CREATE_NO_METADATA)
	VIR_DOMAIN_SNAPSHOT_CREATE_HALT        = VirDomainSnapshotCreateFlags(C.VIR_DOMAIN_SNAPSHOT_CREATE_HALT)
	VIR_DOMAIN_SNAPSHOT_CREATE_DISK_ONLY   = VirDomainSnapshotCreateFlags(C.VIR_DOMAIN_SNAPSHOT_CREATE_DISK_ONLY)
	VIR_DOMAIN_SNAPSHOT_CREATE_REUSE_EXT   = VirDomainSnapshotCreateFlags(C.VIR_DOMAIN_SNAPSHOT_CREATE_REUSE_EXT)
	VIR_DOMAIN_SNAPSHOT_CREATE_QUIESCE     = VirDomainSnapshotCreateFlags(C.VIR_DOMAIN_SNAPSHOT_CREATE_QUIESCE)
	VIR_DOMAIN_SNAPSHOT_CREATE_ATOMIC      = VirDomainSnapshotCreateFlags(C.VIR_DOMAIN_SNAPSHOT_CREATE_ATOMIC)
	VIR_DOMAIN_SNAPSHOT_CREATE_LIVE        = VirDomainSnapshotCreateFlags(C.VIR_DOMAIN_SNAPSHOT_CREATE_LIVE)
)

type VirDomainSnapshotListFlags int

const (
	VIR_DOMAIN_SNAPSHOT_LIST_ROOTS       = VirDomainSnapshotListFlags(C.VIR_DOMAIN_SNAPSHOT_LIST_ROOTS)
	VIR_DOMAIN_SNAPSHOT_LIST_DESCENDANTS = VirDomainSnapshotListFlags(C.VIR_DOMAIN_SNAPSHOT_LIST_DESCENDANTS)
	VIR_DOMAIN_SNAPSHOT_LIST_LEAVES      = VirDomainSnapshotListFlags(C.VIR_DOMAIN_SNAPSHOT_LIST_LEAVES)
	VIR_DOMAIN_SNAPSHOT_LIST_NO_LEAVES   = VirDomainSnapshotListFlags(C.VIR_DOMAIN_SNAPSHOT_LIST_NO_LEAVES)
	VIR_DOMAIN_SNAPSHOT_LIST_METADATA    = VirDomainSnapshotListFlags(C.VIR_DOMAIN_SNAPSHOT_LIST_METADATA)
	VIR_DOMAIN_SNAPSHOT_LIST_NO_METADATA = VirDomainSnapshotListFlags(C.VIR_DOMAIN_SNAPSHOT_LIST_NO_METADATA)
	VIR_DOMAIN_SNAPSHOT_LIST_INACTIVE    = VirDomainSnapshotListFlags(C.VIR_DOMAIN_SNAPSHOT_LIST_INACTIVE)
	VIR_DOMAIN_SNAPSHOT_LIST_ACTIVE      = VirDomainSnapshotListFlags(C.VIR_DOMAIN_SNAPSHOT_LIST_ACTIVE)
	VIR_DOMAIN_SNAPSHOT_LIST_DISK_ONLY   = VirDomainSnapshotListFlags(C.VIR_DOMAIN_SNAPSHOT_LIST_DISK_ONLY)
	VIR_DOMAIN_SNAPSHOT_LIST_INTERNAL    = VirDomainSnapshotListFlags(C.VIR_DOMAIN_SNAPSHOT_LIST_INTERNAL)
	VIR_DOMAIN_SNAPSHOT_LIST_EXTERNAL    = VirDomainSnapshotListFlags(C.VIR_DOMAIN_SNAPSHOT_LIST_EXTERNAL)
)

type VirDomainSnapshotRevertFlags int

const (
	VIR_DOMAIN_SNAPSHOT_REVERT_RUNNING = VirDomainSnapshotRevertFlags(C.VIR_DOMAIN_SNAPSHOT_REVERT_RUNNING)
	VIR_DOMAIN_SNAPSHOT_REVERT_PAUSED  = VirDomainSnapshotRevertFlags(C.VIR_DOMAIN_SNAPSHOT_REVERT_PAUSED)
	VIR_DOMAIN_SNAPSHOT_REVERT_FORCE   = VirDomainSnapshotRevertFlags(C.VIR_DOMAIN_SNAPSHOT_REVERT_FORCE)
)

type VirDomainSnapshotDeleteFlags int

const (
	VIR_DOMAIN_SNAPSHOT_DELETE_CHILDREN      = VirDomainSnapshotDeleteFlags(C.VIR_DOMAIN_SNAPSHOT_DELETE_CHILDREN)
	VIR_DOMAIN_SNAPSHOT_DELETE_METADATA_ONLY = VirDomainSnapshotDeleteFlags(C.VIR_DOMAIN_SNAPSHOT_DELETE_METADATA_ONLY)
	VIR_DOMAIN_SNAPSHOT_DELETE_CHILDREN_ONLY = VirDomainSnapshotDeleteFlags(C.VIR_DOMAIN_SNAPSHOT_DELETE_CHILDREN_ONLY)
)

type VirDomainSnapshot struct {
	ptr C.virDomainSnapshotPtr
}

func (s *VirDomainSnapshot) Free() error {
	if result := C.virDomainSnapshotFree(s.ptr); result != 0 {
		return GetLastError()
	}
	return nil
}

func (s *VirDomainSnapshot) Delete(flags uint32) error {
	result := C.virDomainSnapshotDelete(s.ptr, C.uint(flags))
	if result != 0 {
		return GetLastError()
	}
	return nil
}

func (s *VirDomainSnapshot) RevertToSnapshot(flags uint32) error {
	result := C.virDomainRevertToSnapshot(s.ptr, C.uint(flags))
	if result != 0 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) CreateSnapshotXML(xml string, flags uint32) (*VirDomainSnapshot, error) {
	cXml := C.CString(xml)
	defer C.free(unsafe.Pointer(cXml))
	result := C.virDomainSnapshotCreateXML(d.ptr, cXml, C.uint(flags))
	if result == nil {
		return nil, GetLastError()
	}
	return &VirDomainSnapshot{ptr: result}, nil
}

func (d *VirDomain) Save(destFile string) error {
	cPath := C.CString(destFile)
	defer C.free(unsafe.Pointer(cPath))
	result := C.virDomainSave(d.ptr, cPath)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (d *VirDomain) SaveFlags(destFile string, destXml string, flags uint32) error {
	cDestFile := C.CString(destFile)
	cDestXml := C.CString(destXml)
	defer C.free(unsafe.Pointer(cDestXml))
	defer C.free(unsafe.Pointer(cDestFile))
	result := C.virDomainSaveFlags(d.ptr, cDestFile, cDestXml, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (conn VirConnection) Restore(srcFile string) error {
	cPath := C.CString(srcFile)
	defer C.free(unsafe.Pointer(cPath))
	if result := C.virDomainRestore(conn.ptr, cPath); result == -1 {
		return GetLastError()
	}
	return nil
}

func (conn VirConnection) RestoreFlags(srcFile, xmlConf string, flags uint32) error {
	cPath := C.CString(srcFile)
	defer C.free(unsafe.Pointer(cPath))
	var cXmlConf *C.char
	if xmlConf != "" {
		cXmlConf = C.CString(xmlConf)
		defer C.free(unsafe.Pointer(cXmlConf))
	}
	if result := C.virDomainRestoreFlags(conn.ptr, cPath, cXmlConf, C.uint(flags)); result == -1 {
		return GetLastError()
	}
	return nil
}

func (s *VirDomainSnapshot) IsCurrent(flags uint32) (bool, error) {
	result := C.virDomainSnapshotIsCurrent(s.ptr, C.uint(flags))
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (s *VirDomainSnapshot) HasMetadata(flags uint32) (bool, error) {
	result := C.virDomainSnapshotHasMetadata(s.ptr, C.uint(flags))
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (s *VirDomainSnapshot) GetXMLDesc(flags uint32) (string, error) {
	result := C.virDomainSnapshotGetXMLDesc(s.ptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

func (s *VirDomainSnapshot) GetName() (string, error) {
	name := C.virDomainSnapshotGetName(s.ptr)
	if name == nil {
		return "", GetLastError()
	}
	return C.GoString(name), nil
}

func (s *VirDomainSnapshot) GetParent(flags uint32) (*VirDomainSnapshot, error) {
	ptr := C.virDomainSnapshotGetParent(s.ptr, C.uint(flags))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirDomainSnapshot{ptr: ptr}, nil
}

func (s *VirDomainSnapshot) NumChildren(flags VirDomainSnapshotListFlags) (int, error) {
	result := int(C.virDomainSnapshotNumChildren(s.ptr, C.uint(flags)))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (s *VirDomainSnapshot) ListChildrenNames(flags VirDomainSnapshotListFlags) ([]string, error) {
	const maxNames = 1024
	var names [maxNames](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numNames := C.virDomainSnapshotListChildrenNames(
		s.ptr,
		(**C.char)(namesPtr),
		maxNames, C.uint(flags))
	if numNames == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numNames)
	for k := 0; k < int(numNames); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (d *VirDomainSnapshot) ListAllChildren(flags VirDomainSnapshotListFlags) ([]VirDomainSnapshot, error) {
	var cList *C.virDomainSnapshotPtr
	numVols := C.virDomainSnapshotListAllChildren(d.ptr, (**C.virDomainSnapshotPtr)(&cList), C.uint(flags))
	if numVols == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numVols),
		Cap:  int(numVols),
	}
	var pools []VirDomainSnapshot
	slice := *(*[]C.virDomainSnapshotPtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		pools = append(pools, VirDomainSnapshot{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return pools, nil
}
