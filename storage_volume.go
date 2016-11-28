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

type VirStorageVolCreateFlags int

const (
	VIR_STORAGE_VOL_CREATE_PREALLOC_METADATA = VirStorageVolCreateFlags(C.VIR_STORAGE_VOL_CREATE_PREALLOC_METADATA)
)

type VirStorageVolDeleteFlags int

const (
	VIR_STORAGE_VOL_DELETE_NORMAL = VirStorageVolDeleteFlags(C.VIR_STORAGE_VOL_DELETE_NORMAL) // Delete metadata only (fast)
	VIR_STORAGE_VOL_DELETE_ZEROED = VirStorageVolDeleteFlags(C.VIR_STORAGE_VOL_DELETE_ZEROED) // Clear all data to zeros (slow)
)

type VirStorageVolResizeFlags int

const (
	VIR_STORAGE_VOL_RESIZE_ALLOCATE = VirStorageVolResizeFlags(C.VIR_STORAGE_VOL_RESIZE_ALLOCATE) // force allocation of new size
	VIR_STORAGE_VOL_RESIZE_DELTA    = VirStorageVolResizeFlags(C.VIR_STORAGE_VOL_RESIZE_DELTA)    // size is relative to current
	VIR_STORAGE_VOL_RESIZE_SHRINK   = VirStorageVolResizeFlags(C.VIR_STORAGE_VOL_RESIZE_SHRINK)   // allow decrease in capacity
)

type VirStorageVolType int

const (
	VIR_STORAGE_VOL_FILE    = VirStorageVolType(C.VIR_STORAGE_VOL_FILE)    // Regular file based volumes
	VIR_STORAGE_VOL_BLOCK   = VirStorageVolType(C.VIR_STORAGE_VOL_BLOCK)   // Block based volumes
	VIR_STORAGE_VOL_DIR     = VirStorageVolType(C.VIR_STORAGE_VOL_DIR)     // Directory-passthrough based volume
	VIR_STORAGE_VOL_NETWORK = VirStorageVolType(C.VIR_STORAGE_VOL_NETWORK) //Network volumes like RBD (RADOS Block Device)
	VIR_STORAGE_VOL_NETDIR  = VirStorageVolType(C.VIR_STORAGE_VOL_NETDIR)  // Network accessible directory that can contain other network volumes
)

type VirStorageVolWipeAlgorithm int

const (
	VIR_STORAGE_VOL_WIPE_ALG_ZERO       = VirStorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_ZERO)       // 1-pass, all zeroes
	VIR_STORAGE_VOL_WIPE_ALG_NNSA       = VirStorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_NNSA)       // 4-pass NNSA Policy Letter NAP-14.1-C (XVI-8)
	VIR_STORAGE_VOL_WIPE_ALG_DOD        = VirStorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_DOD)        // 4-pass DoD 5220.22-M section 8-306 procedure
	VIR_STORAGE_VOL_WIPE_ALG_BSI        = VirStorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_BSI)        // 9-pass method recommended by the German Center of Security in Information Technologies
	VIR_STORAGE_VOL_WIPE_ALG_GUTMANN    = VirStorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_GUTMANN)    // The canonical 35-pass sequence
	VIR_STORAGE_VOL_WIPE_ALG_SCHNEIER   = VirStorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_SCHNEIER)   // 7-pass method described by Bruce Schneier in "Applied Cryptography" (1996)
	VIR_STORAGE_VOL_WIPE_ALG_PFITZNER7  = VirStorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_PFITZNER7)  // 7-pass random
	VIR_STORAGE_VOL_WIPE_ALG_PFITZNER33 = VirStorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_PFITZNER33) // 33-pass random
	VIR_STORAGE_VOL_WIPE_ALG_RANDOM     = VirStorageVolWipeAlgorithm(C.VIR_STORAGE_VOL_WIPE_ALG_RANDOM)     // 1-pass random
)

type VirStorageVol struct {
	ptr C.virStorageVolPtr
}

type VirStorageVolInfo struct {
	ptr C.virStorageVolInfo
}

func (v *VirStorageVol) Delete(flags VirStorageVolDeleteFlags) error {
	result := C.virStorageVolDelete(v.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (v *VirStorageVol) Free() error {
	if result := C.virStorageVolFree(v.ptr); result != 0 {
		return GetLastError()
	}
	return nil
}

func (v *VirStorageVol) GetInfo() (VirStorageVolInfo, error) {
	vi := VirStorageVolInfo{}
	var ptr C.virStorageVolInfo
	result := C.virStorageVolGetInfo(v.ptr, (*C.virStorageVolInfo)(unsafe.Pointer(&ptr)))
	if result == -1 {
		return vi, GetLastError()
	}
	vi.ptr = ptr
	return vi, nil
}

func (i *VirStorageVolInfo) GetType() VirStorageVolType {
	return VirStorageVolType(i.ptr._type)
}

func (i *VirStorageVolInfo) GetCapacityInBytes() uint64 {
	return uint64(i.ptr.capacity)
}

func (i *VirStorageVolInfo) GetAllocationInBytes() uint64 {
	return uint64(i.ptr.allocation)
}

func (v *VirStorageVol) GetKey() (string, error) {
	key := C.virStorageVolGetKey(v.ptr)
	if key == nil {
		return "", GetLastError()
	}
	return C.GoString(key), nil
}

func (v *VirStorageVol) GetName() (string, error) {
	name := C.virStorageVolGetName(v.ptr)
	if name == nil {
		return "", GetLastError()
	}
	return C.GoString(name), nil
}

func (v *VirStorageVol) GetPath() (string, error) {
	result := C.virStorageVolGetPath(v.ptr)
	if result == nil {
		return "", GetLastError()
	}
	path := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return path, nil
}

func (v *VirStorageVol) GetXMLDesc(flags uint32) (string, error) {
	result := C.virStorageVolGetXMLDesc(v.ptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

func (v *VirStorageVol) Resize(capacity uint64, flags VirStorageVolResizeFlags) error {
	result := C.virStorageVolResize(v.ptr, C.ulonglong(capacity), C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (v *VirStorageVol) Wipe(flags uint32) error {
	result := C.virStorageVolWipe(v.ptr, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}
func (v *VirStorageVol) WipePattern(algorithm VirStorageVolWipeAlgorithm, flags uint32) error {
	result := C.virStorageVolWipePattern(v.ptr, C.uint(algorithm), C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (v *VirStorageVol) Upload(stream *VirStream, offset, length uint64, flags uint32) error {
	if C.virStorageVolUpload(v.ptr, stream.ptr, C.ulonglong(offset),
		C.ulonglong(length), C.uint(flags)) == -1 {
		return GetLastError()
	}
	return nil
}

func (v *VirStorageVol) Download(stream *VirStream, offset, length uint64, flags uint32) error {
	if C.virStorageVolDownload(v.ptr, stream.ptr, C.ulonglong(offset),
		C.ulonglong(length), C.uint(flags)) == -1 {
		return GetLastError()
	}
	return nil
}

func (v *VirStorageVol) LookupPoolByVolume() (VirStoragePool, error) {
	poolPtr := C.virStoragePoolLookupByVolume(v.ptr)
	if poolPtr == nil {
		return VirStoragePool{}, GetLastError()
	}
	return VirStoragePool{ptr: poolPtr}, nil
}
