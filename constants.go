package libvirt

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>

*/
import "C"

// virStreamFlags
const (
	VIR_STREAM_NONBLOCK = C.VIR_STREAM_NONBLOCK
)
