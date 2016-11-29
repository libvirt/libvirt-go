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

type VirNodeDeviceEventID int

const (
	VIR_NODE_DEVICE_EVENT_ID_LIFECYCLE = VirNodeDeviceEventID(C.VIR_NODE_DEVICE_EVENT_ID_LIFECYCLE)
	VIR_NODE_DEVICE_EVENT_ID_UPDATE    = VirNodeDeviceEventID(C.VIR_NODE_DEVICE_EVENT_ID_UPDATE)
)

type VirNodeDeviceEventLifecycleType int

const (
	VIR_NODE_DEVICE_EVENT_CREATED = VirNodeDeviceEventLifecycleType(C.VIR_NODE_DEVICE_EVENT_CREATED)
	VIR_NODE_DEVICE_EVENT_DELETED = VirNodeDeviceEventLifecycleType(C.VIR_NODE_DEVICE_EVENT_DELETED)
)

type VirNodeDevice struct {
	ptr C.virNodeDevicePtr
}

func (s *VirNodeDevice) Free() error {
	if result := C.virNodeDeviceFree(s.ptr); result != 0 {
		return GetLastError()
	}
	return nil
}

func (n *VirNodeDevice) Destroy() error {
	result := C.virNodeDeviceDestroy(n.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *VirNodeDevice) Reset() error {
	result := C.virNodeDeviceReset(n.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *VirNodeDevice) Detach() error {
	result := C.virNodeDeviceDettach(n.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *VirNodeDevice) DetachFlags(driverName string, flags uint32) error {
	cDriverName := C.CString(driverName)
	defer C.free(cDriverName)
	result := C.virNodeDeviceDetachFlags(n.ptr, cDriverName, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *VirNodeDevice) ReAttach() error {
	result := C.virNodeDeviceReAttach(n.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *VirNodeDevice) GetName() (string, error) {
	name := C.virNodeDeviceGetName(n.ptr)
	if name == nil {
		return "", GetLastError()
	}
	return C.GoString(name), nil
}

func (n *VirNodeDevice) GetXMLDesc(flags uint32) (string, error) {
	result := C.virNodeDeviceGetXMLDesc(n.ptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

func (n *VirNodeDevice) GetParent() (string, error) {
	result := C.virNodeDeviceGetParent(n.ptr)
	if result == nil {
		return "", GetLastError()
	}
	defer C.free(result)
	return C.GoString(result), nil
}

func (p *VirNodeDevice) NumOfStorageCaps() (int, error) {
	result := int(C.virNodeDeviceNumOfCaps(p.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (p *VirNodeDevice) ListStorageCaps() ([]string, error) {
	const maxCaps = 1024
	var names [maxCaps](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numCaps := C.virNodeDeviceListCaps(
		p.ptr,
		(**C.char)(namesPtr),
		maxCaps)
	if numCaps == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numCaps)
	for k := 0; k < int(numCaps); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}
