package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <assert.h>
#include "connect_compat.h"

int virNodeAllocPagesCompat(virConnectPtr conn,
			    unsigned int npages,
			    unsigned int *pageSizes,
			    unsigned long long *pageCounts,
			    int startCell,
			    unsigned int cellCount,
			    unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002009
    assert(0); // Caller should have checked version
#else
    return virNodeAllocPages(conn, npages, pageSizes, pageCounts, startCell, cellCount, flags);
#endif
}


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
