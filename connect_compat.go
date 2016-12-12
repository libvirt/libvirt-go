package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <assert.h>
#include "connect_compat.h"

virDomainPtr virDomainDefineXMLFlagsCompat(virConnectPtr conn,
					   const char *xml,
					   unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002012
    assert(0); // Caller should have checked version
#else
    return virDomainDefineXMLFlags(conn, xml, flags);
#endif
}

*/
import "C"
