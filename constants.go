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

// virIPAddrType
const (
	VIR_IP_ADDR_TYPE_IPV4 = C.VIR_IP_ADDR_TYPE_IPV4
	VIR_IP_ADDR_TYPE_IPV6 = C.VIR_IP_ADDR_TYPE_IPV6
)
