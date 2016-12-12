package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <assert.h>
#include "domain_compat.h"

int virDomainAddIOThreadCompat(virDomainPtr domain,
			       unsigned int iothread_id,
			       unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002015
    assert(0); // Caller should have checked version
#else
    return virDomainAddIOThread(domain, iothread_id, flags);
#endif
}


int virDomainDelIOThreadCompat(virDomainPtr domain,
			       unsigned int iothread_id,
			       unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002015
    assert(0); // Caller should have checked version
#else
    return virDomainDelIOThread(domain, iothread_id, flags);
#endif
}


int virDomainSetUserPasswordCompat(virDomainPtr dom,
				   const char *user,
				   const char *password,
				   unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002016
    assert(0); // Caller should have checked version
#else
    return virDomainSetUserPassword(dom, user, password, flags);
#endif
}


int virDomainRenameCompat(virDomainPtr dom,
			  const char *new_name,
			  unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002019
    assert(0); // Caller should have checked version
#else
    return virDomainRename(dom, new_name, flags);
#endif
}


int virDomainGetPerfEventsCompat(virDomainPtr dom,
				 virTypedParameterPtr *params,
				 int *nparams,
				 unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1003003
    assert(0); // Caller should have checked version
#else
    return virDomainGetPerfEventsCompat(dom, params, nparams, flags);
#endif
}


int virDomainSetPerfEventsCompat(virDomainPtr dom,
				 virTypedParameterPtr params,
				 int nparams,
				 unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1003003
    assert(0); // Caller should have checked version
#else
    return virDomainSetPerfEventsCompat(dom, params, nparams, flags);
#endif
}


int virDomainMigrateStartPostCopyCompat(virDomainPtr domain,
					unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1003003
    assert(0); // Caller should have checked version
#else
    return virDomainMigrateStartPostCopy(domain, flags);
#endif
}


int virDomainGetGuestVcpusCompat(virDomainPtr domain,
				 virTypedParameterPtr *params,
				 unsigned int *nparams,
				 unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 2000000
    assert(0); // Caller should have checked version
#else
    return virDomainGetGuestVcpus(domain, params, nparams, flags);
#endif
}


int virDomainSetGuestVcpusCompat(virDomainPtr domain,
				 const char *cpumap,
				 int state,
				 unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 2000000
    assert(0); // Caller should have checked version
#else
    return virDomainSetGuestVcpus(domain, cpumap, state, flags);
#endif
}

*/
import "C"
