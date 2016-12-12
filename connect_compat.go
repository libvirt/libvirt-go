package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <assert.h>
#include "connect_compat.h"

char * virConnectGetDomainCapabilitiesCompat(virConnectPtr conn,
					     const char *emulatorbin,
					     const char *arch,
					     const char *machine,
					     const char *virttype,
					     unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002007
    assert(0); // Caller should have checked version
#else
    return virConnectGetDomainCapabilities(conn, emulatorbin, arch, machine, virttype, flags);
#endif
}

int virConnectGetAllDomainStatsCompat(virConnectPtr conn,
				      unsigned int stats,
				      virDomainStatsRecordPtr **retStats,
				      unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002008
    assert(0); // Caller should have checked version
#else
    return virConnectGetAllDomainStats(conn, stats, retStats, flags);
#endif
}

int virDomainListGetStatsCompat(virDomainPtr *doms,
				unsigned int stats,
				virDomainStatsRecordPtr **retStats,
				unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002008
    assert(0); // Caller should have checked version
#else
    return virDomainListGetStats(doms, stats, retStats, flags);
#endif
}

void virDomainStatsRecordListFreeCompat(virDomainStatsRecordPtr *stats)
{
}

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
