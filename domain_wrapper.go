/*
 * This file is part of the libvirt-go project
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 *
 * Copyright (c) 2013 Alex Zorin
 * Copyright (C) 2016 Red Hat, Inc.
 *
 */

package libvirt

/*
#cgo pkg-config: libvirt
#include <assert.h>
#include "domain_wrapper.h"

int virDomainCoreDumpWithFormatWrapper(virDomainPtr domain,
				      const char *to,
				      unsigned int dumpformat,
				      unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002003
    assert(0); // Caller should have checked version
#else
    return virDomainCoreDumpWithFormat(domain, to, dumpformat, flags);
#endif
}


int virDomainGetTimeWrapper(virDomainPtr dom,
			   long long *seconds,
			   unsigned int *nseconds,
			   unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002005
    assert(0); // Caller should have checked version
#else
    return virDomainGetTime(dom, seconds, nseconds, flags);
#endif
}

int virDomainSetTimeWrapper(virDomainPtr dom,
			   long long seconds,
			   unsigned int nseconds,
			   unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002005
    assert(0); // Caller should have checked version
#else
    return virDomainSetTime(dom, seconds, nseconds, flags);
#endif
}

int virDomainFSFreezeWrapper(virDomainPtr dom,
			    const char **mountpoints,
			    unsigned int nmountpoints,
			    unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002005
    assert(0); // Caller should have checked version
#else
    return virDomainFSFreeze(dom, mountpoints, nmountpoints, flags);
#endif
}

int virDomainFSThawWrapper(virDomainPtr dom,
			  const char **mountpoints,
			  unsigned int nmountpoints,
			  unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002005
    assert(0); // Caller should have checked version
#else
    return virDomainFSThaw(dom, mountpoints, nmountpoints, flags);
#endif
}

int virDomainBlockCopyWrapper(virDomainPtr dom, const char *disk,
			     const char *destxml,
			     virTypedParameterPtr params,
			     int nparams,
			     unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002008
    assert(0); // Caller should have checked version
#else
    return virDomainBlockCopy(dom, disk, destxml, params, nparams, flags);
#endif
}

int virDomainOpenGraphicsFDWrapper(virDomainPtr dom,
				  unsigned int idx,
				  unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002008
    assert(0); // Caller should have checked version
#else
    return virDomainOpenGraphicsFD(dom, idx, flags);
#endif
}

void virDomainFSInfoFreeWrapper(virDomainFSInfoPtr info)
{
}

int virDomainGetFSInfoWrapper(virDomainPtr dom,
			     virDomainFSInfoPtr **info,
			     unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002011
    assert(0); // Caller should have checked version
#else
    return virDomainGetFSInfo(dom, info, flags);
#endif
}

int virDomainInterfaceAddressesWrapper(virDomainPtr dom,
				      virDomainInterfacePtr **ifaces,
				      unsigned int source,
				      unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002014
    assert(0); // Caller should have checked version
#else
    return virDomainInterfaceAddresses(dom, ifaces, source, flags);
#endif
}

void virDomainInterfaceFreeWrapper(virDomainInterfacePtr iface)
{
}

void virDomainIOThreadInfoFreeWrapper(virDomainIOThreadInfoPtr info)
{
}

int virDomainGetIOThreadInfoWrapper(virDomainPtr domain,
				   virDomainIOThreadInfoPtr **info,
				   unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002014
    assert(0); // Caller should have checked version
#else
    return virDomainGetIOThreadInfo(domain, info, flags);
#endif
}
int virDomainPinIOThreadWrapper(virDomainPtr domain,
			       unsigned int iothread_id,
			       unsigned char *cpumap,
			       int maplen,
			       unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002014
    assert(0); // Caller should have checked version
#else
    return virDomainPinIOThread(domain, iothread_id, cpumap, maplen, flags);
#endif
}

int virDomainAddIOThreadWrapper(virDomainPtr domain,
			       unsigned int iothread_id,
			       unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002015
    assert(0); // Caller should have checked version
#else
    return virDomainAddIOThread(domain, iothread_id, flags);
#endif
}


int virDomainDelIOThreadWrapper(virDomainPtr domain,
			       unsigned int iothread_id,
			       unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002015
    assert(0); // Caller should have checked version
#else
    return virDomainDelIOThread(domain, iothread_id, flags);
#endif
}


int virDomainSetUserPasswordWrapper(virDomainPtr dom,
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


int virDomainRenameWrapper(virDomainPtr dom,
			  const char *new_name,
			  unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002019
    assert(0); // Caller should have checked version
#else
    return virDomainRename(dom, new_name, flags);
#endif
}


int virDomainGetPerfEventsWrapper(virDomainPtr dom,
				 virTypedParameterPtr *params,
				 int *nparams,
				 unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1003003
    assert(0); // Caller should have checked version
#else
    return virDomainGetPerfEvents(dom, params, nparams, flags);
#endif
}


int virDomainSetPerfEventsWrapper(virDomainPtr dom,
				 virTypedParameterPtr params,
				 int nparams,
				 unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1003003
    assert(0); // Caller should have checked version
#else
    return virDomainSetPerfEvents(dom, params, nparams, flags);
#endif
}


int virDomainMigrateStartPostCopyWrapper(virDomainPtr domain,
					unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1003003
    assert(0); // Caller should have checked version
#else
    return virDomainMigrateStartPostCopy(domain, flags);
#endif
}


int virDomainGetGuestVcpusWrapper(virDomainPtr domain,
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


int virDomainSetGuestVcpusWrapper(virDomainPtr domain,
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

int virDomainSetVcpuWrapper(virDomainPtr domain,
			   const char *cpumap,
			   int state,
			   unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 3001000
    assert(0); // Caller should have checked version
#else
    return virDomainSetVcpu(domain, cpumap, state, flags);
#endif
}


int virDomainSetBlockThresholdWrapper(virDomainPtr domain,
                                     const char *dev,
                                     unsigned long long threshold,
                                     unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 3002000
    assert(0); // Caller should have checked version
#else
    return virDomainSetBlockThreshold(domain, dev, threshold, flags);
#endif
}

int virDomainMigrateGetMaxDowntimeWrapper(virDomainPtr domain,
					 unsigned long long *downtime,
					 unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 3007000
    assert(0); // Caller should have checked version
#else
    return virDomainMigrateGetMaxDowntime(domain, downtime, flags);
#endif
}


char *virDomainManagedSaveGetXMLDescWrapper(virDomainPtr domain,
					   unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 3007000
    assert(0); // Caller should have checked version
#else
    return virDomainManagedSaveGetXMLDesc(domain, flags);
#endif
}


int virDomainManagedSaveDefineXMLWrapper(virDomainPtr domain,
					const char *dxml,
					unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 3007000
    assert(0); // Caller should have checked version
#else
    return virDomainManagedSaveDefineXML(domain, dxml, flags);
#endif
}

int virDomainSetLifecycleActionWrapper(virDomainPtr domain,
                                      unsigned int type,
                                      unsigned int action,
                                      unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 3009000
    assert(0); // Caller should have checked version
#else
    return virDomainSetLifecycleAction(domain, type, action, flags);
#endif
}

int virDomainDetachDeviceAliasWrapper(virDomainPtr domain,
				     const char *alias,
				     unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 4004000
    assert(0); // Caller should have checked version
#else
    return virDomainDetachDeviceAlias(domain, alias, flags);
#endif
}

int virDomainGetLaunchSecurityInfoWrapper(virDomainPtr domain,
					 virTypedParameterPtr *params,
					 int *nparams,
					 unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 4005000
    assert(0); // Caller should have checked version
#else
    return virDomainGetLaunchSecurityInfo(domain, params, nparams, flags);
#endif
}

*/
import "C"
