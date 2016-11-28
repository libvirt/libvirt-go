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

// virStorageVolCreateFlags
const (
	VIR_STORAGE_VOL_CREATE_PREALLOC_METADATA = C.VIR_STORAGE_VOL_CREATE_PREALLOC_METADATA
)

// virStorageVolDeleteFlags
const (
	VIR_STORAGE_VOL_DELETE_NORMAL = C.VIR_STORAGE_VOL_DELETE_NORMAL // Delete metadata only (fast)
	VIR_STORAGE_VOL_DELETE_ZEROED = C.VIR_STORAGE_VOL_DELETE_ZEROED // Clear all data to zeros (slow)
)

// virStorageVolResizeFlags
const (
	VIR_STORAGE_VOL_RESIZE_ALLOCATE = C.VIR_STORAGE_VOL_RESIZE_ALLOCATE // force allocation of new size
	VIR_STORAGE_VOL_RESIZE_DELTA    = C.VIR_STORAGE_VOL_RESIZE_DELTA    // size is relative to current
	VIR_STORAGE_VOL_RESIZE_SHRINK   = C.VIR_STORAGE_VOL_RESIZE_SHRINK   // allow decrease in capacity
)

// virStorageVolType
const (
	VIR_STORAGE_VOL_FILE    = C.VIR_STORAGE_VOL_FILE    // Regular file based volumes
	VIR_STORAGE_VOL_BLOCK   = C.VIR_STORAGE_VOL_BLOCK   // Block based volumes
	VIR_STORAGE_VOL_DIR     = C.VIR_STORAGE_VOL_DIR     // Directory-passthrough based volume
	VIR_STORAGE_VOL_NETWORK = C.VIR_STORAGE_VOL_NETWORK //Network volumes like RBD (RADOS Block Device)
	VIR_STORAGE_VOL_NETDIR  = C.VIR_STORAGE_VOL_NETDIR  // Network accessible directory that can contain other network volumes
)

// virStorageVolWipeAlgorithm
const (
	VIR_STORAGE_VOL_WIPE_ALG_ZERO       = C.VIR_STORAGE_VOL_WIPE_ALG_ZERO       // 1-pass, all zeroes
	VIR_STORAGE_VOL_WIPE_ALG_NNSA       = C.VIR_STORAGE_VOL_WIPE_ALG_NNSA       // 4-pass NNSA Policy Letter NAP-14.1-C (XVI-8)
	VIR_STORAGE_VOL_WIPE_ALG_DOD        = C.VIR_STORAGE_VOL_WIPE_ALG_DOD        // 4-pass DoD 5220.22-M section 8-306 procedure
	VIR_STORAGE_VOL_WIPE_ALG_BSI        = C.VIR_STORAGE_VOL_WIPE_ALG_BSI        // 9-pass method recommended by the German Center of Security in Information Technologies
	VIR_STORAGE_VOL_WIPE_ALG_GUTMANN    = C.VIR_STORAGE_VOL_WIPE_ALG_GUTMANN    // The canonical 35-pass sequence
	VIR_STORAGE_VOL_WIPE_ALG_SCHNEIER   = C.VIR_STORAGE_VOL_WIPE_ALG_SCHNEIER   // 7-pass method described by Bruce Schneier in "Applied Cryptography" (1996)
	VIR_STORAGE_VOL_WIPE_ALG_PFITZNER7  = C.VIR_STORAGE_VOL_WIPE_ALG_PFITZNER7  // 7-pass random
	VIR_STORAGE_VOL_WIPE_ALG_PFITZNER33 = C.VIR_STORAGE_VOL_WIPE_ALG_PFITZNER33 // 33-pass random
	VIR_STORAGE_VOL_WIPE_ALG_RANDOM     = C.VIR_STORAGE_VOL_WIPE_ALG_RANDOM     // 1-pass random
)

type VirStorageVol struct {
	ptr C.virStorageVolPtr
}

type VirStorageVolInfo struct {
	ptr C.virStorageVolInfo
}

func (v *VirStorageVol) Delete(flags uint32) error {
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

func (i *VirStorageVolInfo) GetType() int {
	return int(i.ptr._type)
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

func (v *VirStorageVol) Resize(capacity uint64, flags uint32) error {
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
func (v *VirStorageVol) WipePattern(algorithm uint32, flags uint32) error {
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
