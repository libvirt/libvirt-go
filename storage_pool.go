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

type VirStoragePoolState int

const (
	VIR_STORAGE_POOL_INACTIVE     = VirStoragePoolState(C.VIR_STORAGE_POOL_INACTIVE)     // Not running
	VIR_STORAGE_POOL_BUILDING     = VirStoragePoolState(C.VIR_STORAGE_POOL_BUILDING)     // Initializing pool,not available
	VIR_STORAGE_POOL_RUNNING      = VirStoragePoolState(C.VIR_STORAGE_POOL_RUNNING)      // Running normally
	VIR_STORAGE_POOL_DEGRADED     = VirStoragePoolState(C.VIR_STORAGE_POOL_DEGRADED)     // Running degraded
	VIR_STORAGE_POOL_INACCESSIBLE = VirStoragePoolState(C.VIR_STORAGE_POOL_INACCESSIBLE) // Running,but not accessible
)

type VirStoragePoolBuildFlags int

const (
	VIR_STORAGE_POOL_BUILD_NEW          = VirStoragePoolBuildFlags(C.VIR_STORAGE_POOL_BUILD_NEW)          // Regular build from scratch
	VIR_STORAGE_POOL_BUILD_REPAIR       = VirStoragePoolBuildFlags(C.VIR_STORAGE_POOL_BUILD_REPAIR)       // Repair / reinitialize
	VIR_STORAGE_POOL_BUILD_RESIZE       = VirStoragePoolBuildFlags(C.VIR_STORAGE_POOL_BUILD_RESIZE)       // Extend existing pool
	VIR_STORAGE_POOL_BUILD_NO_OVERWRITE = VirStoragePoolBuildFlags(C.VIR_STORAGE_POOL_BUILD_NO_OVERWRITE) // Do not overwrite existing pool
	VIR_STORAGE_POOL_BUILD_OVERWRITE    = VirStoragePoolBuildFlags(C.VIR_STORAGE_POOL_BUILD_OVERWRITE)    // Overwrite data
)

type VirStoragePoolCreateFlags int

const (
	VIR_STORAGE_POOL_CREATE_NORMAL                  = VirStoragePoolCreateFlags(C.VIR_STORAGE_POOL_CREATE_NORMAL)
	VIR_STORAGE_POOL_CREATE_WITH_BUILD              = VirStoragePoolCreateFlags(C.VIR_STORAGE_POOL_CREATE_WITH_BUILD)
	VIR_STORAGE_POOL_CREATE_WITH_BUILD_OVERWRITE    = VirStoragePoolCreateFlags(C.VIR_STORAGE_POOL_CREATE_WITH_BUILD_OVERWRITE)
	VIR_STORAGE_POOL_CREATE_WITH_BUILD_NO_OVERWRITE = VirStoragePoolCreateFlags(C.VIR_STORAGE_POOL_CREATE_WITH_BUILD_NO_OVERWRITE)
)

type VirStoragePoolDeleteFlags int

const (
	VIR_STORAGE_POOL_DELETE_NORMAL = VirStoragePoolDeleteFlags(C.VIR_STORAGE_POOL_DELETE_NORMAL)
	VIR_STORAGE_POOL_DELETE_ZEROED = VirStoragePoolDeleteFlags(C.VIR_STORAGE_POOL_DELETE_ZEROED)
)

type VirStoragePoolEventID int

const (
	VIR_STORAGE_POOL_EVENT_ID_LIFECYCLE = VirStoragePoolEventID(C.VIR_STORAGE_POOL_EVENT_ID_LIFECYCLE)
	VIR_STORAGE_POOL_EVENT_ID_REFRESH   = VirStoragePoolEventID(C.VIR_STORAGE_POOL_EVENT_ID_REFRESH)
)

type VirStoragePoolEventLifecycleType int

const (
	VIR_STORAGE_POOL_EVENT_DEFINED   = VirStoragePoolEventLifecycleType(C.VIR_STORAGE_POOL_EVENT_DEFINED)
	VIR_STORAGE_POOL_EVENT_UNDEFINED = VirStoragePoolEventLifecycleType(C.VIR_STORAGE_POOL_EVENT_UNDEFINED)
	VIR_STORAGE_POOL_EVENT_STARTED   = VirStoragePoolEventLifecycleType(C.VIR_STORAGE_POOL_EVENT_STARTED)
	VIR_STORAGE_POOL_EVENT_STOPPED   = VirStoragePoolEventLifecycleType(C.VIR_STORAGE_POOL_EVENT_STOPPED)
)

type VirStoragePool struct {
	ptr C.virStoragePoolPtr
}

type VirStoragePoolInfo struct {
	State      VirStoragePoolState
	Capacity   uint64
	Allocation uint64
	Available  uint64
}

func (p *VirStoragePool) Build(flags VirStoragePoolBuildFlags) error {
	result := C.virStoragePoolBuild(p.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (p *VirStoragePool) Create(flags uint32) error {
	result := C.virStoragePoolCreate(p.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (p *VirStoragePool) Delete(flags uint32) error {
	result := C.virStoragePoolDelete(p.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (p *VirStoragePool) Destroy() error {
	result := C.virStoragePoolDestroy(p.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (p *VirStoragePool) Free() error {
	if result := C.virStoragePoolFree(p.ptr); result != 0 {
		return GetLastError()
	}
	p.ptr = nil
	return nil
}

func (p *VirStoragePool) GetAutostart() (bool, error) {
	var out C.int
	result := C.virStoragePoolGetAutostart(p.ptr, (*C.int)(unsafe.Pointer(&out)))
	if result == -1 {
		return false, GetLastError()
	}
	switch out {
	case 1:
		return true, nil
	default:
		return false, nil
	}
}

func (p *VirStoragePool) GetInfo() (*VirStoragePoolInfo, error) {
	var cinfo C.virStoragePoolInfo
	result := C.virStoragePoolGetInfo(p.ptr, &cinfo)
	if result == -1 {
		return nil, GetLastError()
	}
	return &VirStoragePoolInfo{
		State:      VirStoragePoolState(cinfo.state),
		Capacity:   uint64(cinfo.capacity),
		Allocation: uint64(cinfo.allocation),
		Available:  uint64(cinfo.available),
	}, nil
}

func (p *VirStoragePool) GetName() (string, error) {
	name := C.virStoragePoolGetName(p.ptr)
	if name == nil {
		return "", GetLastError()
	}
	return C.GoString(name), nil
}

func (p *VirStoragePool) GetUUID() ([]byte, error) {
	var cUuid [C.VIR_UUID_BUFLEN](byte)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virStoragePoolGetUUID(p.ptr, (*C.uchar)(cuidPtr))
	if result != 0 {
		return []byte{}, GetLastError()
	}
	return C.GoBytes(cuidPtr, C.VIR_UUID_BUFLEN), nil
}

func (p *VirStoragePool) GetUUIDString() (string, error) {
	var cUuid [C.VIR_UUID_STRING_BUFLEN](C.char)
	cuidPtr := unsafe.Pointer(&cUuid)
	result := C.virStoragePoolGetUUIDString(p.ptr, (*C.char)(cuidPtr))
	if result != 0 {
		return "", GetLastError()
	}
	return C.GoString((*C.char)(cuidPtr)), nil
}

func (p *VirStoragePool) GetXMLDesc(flags uint32) (string, error) {
	result := C.virStoragePoolGetXMLDesc(p.ptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

func (p *VirStoragePool) IsActive() (bool, error) {
	result := C.virStoragePoolIsActive(p.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (p *VirStoragePool) IsPersistent() (bool, error) {
	result := C.virStoragePoolIsPersistent(p.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (p *VirStoragePool) SetAutostart(autostart bool) error {
	var cAutostart C.int
	switch autostart {
	case true:
		cAutostart = 1
	default:
		cAutostart = 0
	}
	result := C.virStoragePoolSetAutostart(p.ptr, cAutostart)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (p *VirStoragePool) Refresh(flags uint32) error {
	result := C.virStoragePoolRefresh(p.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (p *VirStoragePool) Undefine() error {
	result := C.virStoragePoolUndefine(p.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (p *VirStoragePool) StorageVolCreateXML(xmlConfig string, flags uint32) (*VirStorageVol, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virStorageVolCreateXML(p.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirStorageVol{ptr: ptr}, nil
}

func (p *VirStoragePool) StorageVolCreateXMLFrom(xmlConfig string, clonevol *VirStorageVol, flags uint32) (*VirStorageVol, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virStorageVolCreateXMLFrom(p.ptr, cXml, clonevol.ptr, C.uint(flags))
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirStorageVol{ptr: ptr}, nil
}

func (p *VirStoragePool) LookupStorageVolByName(name string) (*VirStorageVol, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virStorageVolLookupByName(p.ptr, cName)
	if ptr == nil {
		return nil, GetLastError()
	}
	return &VirStorageVol{ptr: ptr}, nil
}

func (p *VirStoragePool) NumOfStorageVolumes() (int, error) {
	result := int(C.virStoragePoolNumOfVolumes(p.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (p *VirStoragePool) ListStorageVolumes() ([]string, error) {
	const maxVols = 1024
	var names [maxVols](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numStorageVols := C.virStoragePoolListVolumes(
		p.ptr,
		(**C.char)(namesPtr),
		maxVols)
	if numStorageVols == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numStorageVols)
	for k := 0; k < int(numStorageVols); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (p *VirStoragePool) ListAllStorageVolumes(flags uint32) ([]VirStorageVol, error) {
	var cList *C.virStorageVolPtr
	numVols := C.virStoragePoolListAllVolumes(p.ptr, (**C.virStorageVolPtr)(&cList), C.uint(flags))
	if numVols == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numVols),
		Cap:  int(numVols),
	}
	var pools []VirStorageVol
	slice := *(*[]C.virStorageVolPtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		pools = append(pools, VirStorageVol{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return pools, nil
}
