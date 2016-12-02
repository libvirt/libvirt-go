package libvirt

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>

*/
import "C"

type EventHandleType int

const (
	VIR_EVENT_HANDLE_READABLE = EventHandleType(C.VIR_EVENT_HANDLE_READABLE)
	VIR_EVENT_HANDLE_WRITABLE = EventHandleType(C.VIR_EVENT_HANDLE_WRITABLE)
	VIR_EVENT_HANDLE_ERROR    = EventHandleType(C.VIR_EVENT_HANDLE_ERROR)
	VIR_EVENT_HANDLE_HANGUP   = EventHandleType(C.VIR_EVENT_HANDLE_HANGUP)
)

func EventRegisterDefaultImpl() error {
	if i := int(C.virEventRegisterDefaultImpl()); i != 0 {
		return GetLastError()
	}
	return nil
}

func EventRunDefaultImpl() error {
	if i := int(C.virEventRunDefaultImpl()); i != 0 {
		return GetLastError()
	}
	return nil
}
