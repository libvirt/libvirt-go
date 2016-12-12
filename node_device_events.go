package libvirt

import (
	"fmt"
	"unsafe"
)

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include "node_device_compat.h"
#include "node_device_events_cfuncs.h"
*/
import "C"

type NodeDeviceEventGenericCallback func(c *Connect, d *NodeDevice)

type NodeDeviceEventLifecycle struct {
	Event NodeDeviceEventLifecycleType
	// TODO: we can make Detail typesafe somehow ?
	Detail int
}

type NodeDeviceEventLifecycleCallback func(c *Connect, n *NodeDevice, event *NodeDeviceEventLifecycle)

//export nodeDeviceEventLifecycleCallback
func nodeDeviceEventLifecycleCallback(c C.virConnectPtr, s C.virNodeDevicePtr,
	event int, detail int,
	goCallbackId int) {

	node_device := &NodeDevice{ptr: s}
	connection := &Connect{ptr: c}

	eventDetails := &NodeDeviceEventLifecycle{
		Event:  NodeDeviceEventLifecycleType(event),
		Detail: detail,
	}

	callbackFunc := getCallbackId(goCallbackId)
	callback, ok := callbackFunc.(NodeDeviceEventLifecycleCallback)
	if !ok {
		panic("Inappropriate callback type called")
	}
	callback(connection, node_device, eventDetails)
}

//export nodeDeviceEventGenericCallback
func nodeDeviceEventGenericCallback(c C.virConnectPtr, d C.virNodeDevicePtr,
	goCallbackId int) {

	node_device := &NodeDevice{ptr: d}
	connection := &Connect{ptr: c}

	callbackFunc := getCallbackId(goCallbackId)
	callback, ok := callbackFunc.(NodeDeviceEventGenericCallback)
	if !ok {
		panic("Inappropriate callback type called")
	}
	callback(connection, node_device)
}

func (c *Connect) NodeDeviceEventLifecycleRegister(device *NodeDevice, callback NodeDeviceEventLifecycleCallback) (int, error) {
	if C.LIBVIR_VERSION_NUMBER < 2002000 {
		return 0, GetNotImplementedError()
	}
	goCallBackId := registerCallbackId(callback)

	callbackPtr := unsafe.Pointer(C.nodeDeviceEventLifecycleCallback_cgo)
	var cdevice C.virNodeDevicePtr
	if device != nil {
		cdevice = device.ptr
	}
	ret := C.virConnectNodeDeviceEventRegisterAny_cgo(c.ptr, cdevice,
		C.VIR_NODE_DEVICE_EVENT_ID_LIFECYCLE,
		C.virConnectNodeDeviceEventGenericCallback(callbackPtr),
		C.long(goCallBackId))
	if ret == -1 {
		freeCallbackId(goCallBackId)
		return 0, GetLastError()
	}
	return int(ret), nil
}

func (c *Connect) NodeDeviceEventUpdateRegister(device *NodeDevice, callback NodeDeviceEventGenericCallback) (int, error) {
	goCallBackId := registerCallbackId(callback)

	callbackPtr := unsafe.Pointer(C.nodeDeviceEventGenericCallback_cgo)
	var cdevice C.virNodeDevicePtr
	if device != nil {
		cdevice = device.ptr
	}
	ret := C.virConnectNodeDeviceEventRegisterAny_cgo(c.ptr, cdevice,
		C.VIR_NODE_DEVICE_EVENT_ID_UPDATE,
		C.virConnectNodeDeviceEventGenericCallback(callbackPtr),
		C.long(goCallBackId))
	if ret == -1 {
		freeCallbackId(goCallBackId)
		return 0, GetLastError()
	}
	return int(ret), nil
}

func (c *Connect) NodeDeviceEventDeregister(callbackId int) error {
	if C.LIBVIR_VERSION_NUMBER < 2002000 {
		return GetNotImplementedError()
	}
	// Deregister the callback
	if i := int(C.virConnectNodeDeviceEventDeregisterAnyCompat(c.ptr, C.int(callbackId))); i != 0 {
		return GetLastError()
	}
	return nil
}

func (e NodeDeviceEventLifecycle) String() string {
	var event string
	switch e.Event {
	case NODE_DEVICE_EVENT_CREATED:
		event = "created"

	case NODE_DEVICE_EVENT_DELETED:
		event = "deleted"

	default:
		event = "unknown"
	}

	return fmt.Sprintf("NodeDevice event=%q", event)
}
