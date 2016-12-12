package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <assert.h>
#include "storage_pool_compat.h"

int virConnectStoragePoolEventDeregisterAnyCompat(virConnectPtr conn,
						  int callbackID)
{
#if LIBVIR_VERSION_NUMBER < 2000000
    assert(0); // Caller shouuld have checked version
#else
    return virConnectStoragePoolEventDeregisterAny(conn, callbackID);
#endif
}

*/
import "C"
