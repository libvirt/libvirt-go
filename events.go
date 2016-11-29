package libvirt

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>

*/
import "C"

type VirEventHandleType int

const (
	VIR_EVENT_HANDLE_READABLE = VirEventHandleType(C.VIR_EVENT_HANDLE_READABLE)
	VIR_EVENT_HANDLE_WRITABLE = VirEventHandleType(C.VIR_EVENT_HANDLE_WRITABLE)
	VIR_EVENT_HANDLE_ERROR    = VirEventHandleType(C.VIR_EVENT_HANDLE_ERROR)
	VIR_EVENT_HANDLE_HANGUP   = VirEventHandleType(C.VIR_EVENT_HANDLE_HANGUP)
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
