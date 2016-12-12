package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <assert.h>
#include "network_compat.h"

int virConnectNetworkEventDeregisterAnyCompat(virConnectPtr conn,
					      int callbackID)
{
#if LIBVIR_VERSION_NUMBER < 1002001
    assert(0); // Caller should have checked version
#else
    return virConnectNetworkEventDeregisterAny(conn, callbackID);
#endif
}

void virNetworkDHCPLeaseFreeCompat(virNetworkDHCPLeasePtr lease)
{
}

int virNetworkGetDHCPLeasesCompat(virNetworkPtr network,
				  const char *mac,
				  virNetworkDHCPLeasePtr **leases,
				  unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002006
    assert(0); // Caller should have checked version
#else
    return virNetworkGetDHCPLeases(network, mac, leases, flags);
#endif
}

*/
import "C"
