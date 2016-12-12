package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <assert.h>
#include "node_device_compat.h"

int virConnectNodeDeviceEventDeregisterAnyCompat(virConnectPtr conn,
						 int callbackID)
{
#if LIBVIR_VERSION_NUMBER < 2002000
    assert(0); // Caller should have checked version
#else
    return virConnectNodeDeviceEventDeregisterAny(conn, callbackID);
#endif
}

*/
import "C"
