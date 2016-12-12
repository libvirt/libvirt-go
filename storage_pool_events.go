package libvirt

import (
	"fmt"
	"unsafe"
)

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include "storage_pool_compat.h"
#include "storage_pool_events_cfuncs.h"
*/
import "C"

type StoragePoolEventLifecycle struct {
	Event StoragePoolEventLifecycleType
	// TODO: we can make Detail typesafe somehow ?
	Detail int
}

type StoragePoolEventLifecycleCallback func(c *Connect, n *StoragePool, event *StoragePoolEventLifecycle)

//export storagePoolEventLifecycleCallback
func storagePoolEventLifecycleCallback(c C.virConnectPtr, s C.virStoragePoolPtr,
	event int, detail int,
	goCallbackId int) {

	storage_pool := &StoragePool{ptr: s}
	connection := &Connect{ptr: c}

	eventDetails := &StoragePoolEventLifecycle{
		Event:  StoragePoolEventLifecycleType(event),
		Detail: detail,
	}

	callbackFunc := getCallbackId(goCallbackId)
	callback, ok := callbackFunc.(StoragePoolEventLifecycleCallback)
	if !ok {
		panic("Inappropriate callback type called")
	}
	callback(connection, storage_pool, eventDetails)
}

func (c *Connect) StoragePoolEventLifecycleRegister(pool *StoragePool, callback StoragePoolEventLifecycleCallback) (int, error) {
	if C.LIBVIR_VERSION_NUMBER < 2000000 {
		return 0, GetNotImplementedError()
	}

	goCallBackId := registerCallbackId(callback)

	callbackPtr := unsafe.Pointer(C.storagePoolEventLifecycleCallback_cgo)
	var cpool C.virStoragePoolPtr
	if pool != nil {
		cpool = pool.ptr
	}
	ret := C.virConnectStoragePoolEventRegisterAny_cgo(c.ptr, cpool,
		C.VIR_STORAGE_POOL_EVENT_ID_LIFECYCLE,
		C.virConnectStoragePoolEventGenericCallback(callbackPtr),
		C.long(goCallBackId))
	if ret == -1 {
		freeCallbackId(goCallBackId)
		return 0, GetLastError()
	}
	return int(ret), nil
}

func (c *Connect) StoragePoolEventDeregister(callbackId int) error {
	if C.LIBVIR_VERSION_NUMBER < 2000000 {
		return GetNotImplementedError()
	}

	// Deregister the callback
	if i := int(C.virConnectStoragePoolEventDeregisterAnyCompat(c.ptr, C.int(callbackId))); i != 0 {
		return GetLastError()
	}
	return nil
}

func (e StoragePoolEventLifecycle) String() string {
	var event string
	switch e.Event {
	case STORAGE_POOL_EVENT_DEFINED:
		event = "defined"

	case STORAGE_POOL_EVENT_UNDEFINED:
		event = "undefined"

	case STORAGE_POOL_EVENT_STARTED:
		event = "started"

	case STORAGE_POOL_EVENT_STOPPED:
		event = "stopped"

	default:
		event = "unknown"
	}

	return fmt.Sprintf("StoragePool event=%q", event)
}
