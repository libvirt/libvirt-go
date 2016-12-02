package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

import (
	"unsafe"
)

type NodeDeviceEventID int

const (
	NODE_DEVICE_EVENT_ID_LIFECYCLE = NodeDeviceEventID(C.VIR_NODE_DEVICE_EVENT_ID_LIFECYCLE)
	NODE_DEVICE_EVENT_ID_UPDATE    = NodeDeviceEventID(C.VIR_NODE_DEVICE_EVENT_ID_UPDATE)
)

type NodeDeviceEventLifecycleType int

const (
	NODE_DEVICE_EVENT_CREATED = NodeDeviceEventLifecycleType(C.VIR_NODE_DEVICE_EVENT_CREATED)
	NODE_DEVICE_EVENT_DELETED = NodeDeviceEventLifecycleType(C.VIR_NODE_DEVICE_EVENT_DELETED)
)

type NodeDevice struct {
	ptr C.virNodeDevicePtr
}

func (n *NodeDevice) Free() error {
	if result := C.virNodeDeviceFree(n.ptr); result != 0 {
		return GetLastError()
	}
	n.ptr = nil
	return nil
}

func (n *NodeDevice) Destroy() error {
	result := C.virNodeDeviceDestroy(n.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *NodeDevice) Reset() error {
	result := C.virNodeDeviceReset(n.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *NodeDevice) Detach() error {
	result := C.virNodeDeviceDettach(n.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *NodeDevice) DetachFlags(driverName string, flags uint32) error {
	cDriverName := C.CString(driverName)
	defer C.free(cDriverName)
	result := C.virNodeDeviceDetachFlags(n.ptr, cDriverName, C.uint(flags))
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *NodeDevice) ReAttach() error {
	result := C.virNodeDeviceReAttach(n.ptr)
	if result == -1 {
		return GetLastError()
	}
	return nil
}

func (n *NodeDevice) GetName() (string, error) {
	name := C.virNodeDeviceGetName(n.ptr)
	if name == nil {
		return "", GetLastError()
	}
	return C.GoString(name), nil
}

func (n *NodeDevice) GetXMLDesc(flags uint32) (string, error) {
	result := C.virNodeDeviceGetXMLDesc(n.ptr, C.uint(flags))
	if result == nil {
		return "", GetLastError()
	}
	xml := C.GoString(result)
	C.free(unsafe.Pointer(result))
	return xml, nil
}

func (n *NodeDevice) GetParent() (string, error) {
	result := C.virNodeDeviceGetParent(n.ptr)
	if result == nil {
		return "", GetLastError()
	}
	defer C.free(result)
	return C.GoString(result), nil
}

func (p *NodeDevice) NumOfStorageCaps() (int, error) {
	result := int(C.virNodeDeviceNumOfCaps(p.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (p *NodeDevice) ListStorageCaps() ([]string, error) {
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
