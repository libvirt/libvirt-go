package libvirt

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>

*/
import "C"

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
