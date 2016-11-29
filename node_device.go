package libvirt

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
*/
import "C"

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
