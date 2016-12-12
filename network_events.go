package libvirt

import (
	"fmt"
	"unsafe"
)

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include "network_compat.h"
#include "network_events_cfuncs.h"
*/
import "C"

type NetworkEventLifecycle struct {
	Event NetworkEventLifecycleType
	// TODO: we can make Detail typesafe somehow ?
	Detail int
}

type NetworkEventLifecycleCallback func(c *Connect, n *Network, event *NetworkEventLifecycle)

//export networkEventLifecycleCallback
func networkEventLifecycleCallback(c C.virConnectPtr, n C.virNetworkPtr,
	event int, detail int,
	goCallbackId int) {

	network := &Network{ptr: n}
	connection := &Connect{ptr: c}

	eventDetails := &NetworkEventLifecycle{
		Event:  NetworkEventLifecycleType(event),
		Detail: detail,
	}

	callbackFunc := getCallbackId(goCallbackId)
	callback, ok := callbackFunc.(NetworkEventLifecycleCallback)
	if !ok {
		panic("Inappropriate callback type called")
	}
	callback(connection, network, eventDetails)
}

func (c *Connect) NetworkEventLifecycleRegister(net *Network, callback NetworkEventLifecycleCallback) (int, error) {
	goCallBackId := registerCallbackId(callback)
	if C.LIBVIR_VERSION_NUMBER < 1002001 {
		return 0, GetNotImplementedError()
	}

	callbackPtr := unsafe.Pointer(C.networkEventLifecycleCallback_cgo)
	var cnet C.virNetworkPtr
	if net != nil {
		cnet = net.ptr
	}
	ret := C.virConnectNetworkEventRegisterAny_cgo(c.ptr, cnet,
		C.VIR_NETWORK_EVENT_ID_LIFECYCLE,
		C.virConnectNetworkEventGenericCallback(callbackPtr),
		C.long(goCallBackId))
	if ret == -1 {
		freeCallbackId(goCallBackId)
		return 0, GetLastError()
	}
	return int(ret), nil
}

func (c *Connect) NetworkEventDeregister(callbackId int) error {
	if C.LIBVIR_VERSION_NUMBER < 1002001 {
		return GetNotImplementedError()
	}
	// Deregister the callback
	if i := int(C.virConnectNetworkEventDeregisterAnyCompat(c.ptr, C.int(callbackId))); i != 0 {
		return GetLastError()
	}
	return nil
}

func (e NetworkEventLifecycle) String() string {
	var event string
	switch e.Event {
	case NETWORK_EVENT_DEFINED:
		event = "defined"

	case NETWORK_EVENT_UNDEFINED:
		event = "undefined"

	case NETWORK_EVENT_STARTED:
		event = "started"

	case NETWORK_EVENT_STOPPED:
		event = "stopped"

	default:
		event = "unknown"
	}

	return fmt.Sprintf("Network event=%q", event)
}
