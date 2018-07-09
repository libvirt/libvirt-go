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
#include "connect_wrapper.h"
#include "callbacks_wrapper.h"

extern void closeCallback(virConnectPtr, int, long);
void closeCallbackHelper(virConnectPtr conn, int reason, void *opaque)
{
    closeCallback(conn, reason, (long)opaque);
}

int virConnectRegisterCloseCallbackHelper(virConnectPtr c, virConnectCloseFunc cb, long goCallbackId)
{
    void *id = (void*)goCallbackId;
    return virConnectRegisterCloseCallback(c, cb, id, freeGoCallbackHelper);
}

#include <stdio.h>

extern int connectAuthCallback(virConnectCredentialPtr, unsigned int, int);
int connectAuthCallbackHelper(virConnectCredentialPtr cred, unsigned int ncred, void *cbdata)
{
    int *callbackID = cbdata;

    return connectAuthCallback(cred, ncred, *callbackID);
}

virConnectPtr virConnectOpenAuthWrapper(const char *name, int *credtype, uint ncredtype, int callbackID, unsigned int flags)
{
    virConnectAuth auth = {
       .credtype = credtype,
       .ncredtype = ncredtype,
       .cb = connectAuthCallbackHelper,
       .cbdata = &callbackID,
    };

    return virConnectOpenAuth(name, &auth, flags);
}


int virNodeGetFreePagesWrapper(virConnectPtr conn,
                               unsigned int npages,
                               unsigned int *pages,
                               int startcell,
                               unsigned int cellcount,
                               unsigned long long *counts,
                               unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002006
    assert(0); // Caller should have checked version
#else
    return virNodeGetFreePages(conn, npages, pages, startcell, cellcount, counts, flags);
#endif
}

char * virConnectGetDomainCapabilitiesWrapper(virConnectPtr conn,
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

int virConnectGetAllDomainStatsWrapper(virConnectPtr conn,
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

int virDomainListGetStatsWrapper(virDomainPtr *doms,
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

void virDomainStatsRecordListFreeWrapper(virDomainStatsRecordPtr *stats)
{
}

int virNodeAllocPagesWrapper(virConnectPtr conn,
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


virDomainPtr virDomainDefineXMLFlagsWrapper(virConnectPtr conn,
                                            const char *xml,
                                            unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 1002012
    assert(0); // Caller should have checked version
#else
    return virDomainDefineXMLFlags(conn, xml, flags);
#endif
}

virStoragePoolPtr virStoragePoolLookupByTargetPathWrapper(virConnectPtr conn,
                                                          const char *path)
{
#if LIBVIR_VERSION_NUMBER < 4001000
    assert(0); // Caller should have checked version
#else
    return virStoragePoolLookupByTargetPath(conn, path);
#endif
}

char *virConnectBaselineHypervisorCPUWrapper(virConnectPtr conn,
                                             const char *emulator,
                                             const char *arch,
                                             const char *machine,
                                             const char *virttype,
                                             const char **xmlCPUs,
                                             unsigned int ncpus,
                                             unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 4004000
    assert(0); // Caller should have checked version
#else
    return virConnectBaselineHypervisorCPU(conn, emulator, arch, machine, virttype, xmlCPUs, ncpus, flags);
#endif
}

int virConnectCompareHypervisorCPUWrapper(virConnectPtr conn,
                                          const char *emulator,
                                          const char *arch,
                                          const char *machine,
                                          const char *virttype,
                                          const char *xmlCPU,
                                          unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 4004000
    assert(0); // Caller should have checked version
#else
    return virConnectCompareHypervisorCPU(conn, emulator, arch, machine, virttype, xmlCPU, flags);
#endif
}

int virNodeGetSEVInfoWrapper(virConnectPtr conn,
                             virTypedParameterPtr *params,
                             int *nparams,
                             unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 4005000
    assert(0); // Caller should have checked version
#else
    return virNodeGetSEVInfo(conn, params, nparams, flags);
#endif
}

int virConnectListAllNWFilterBindingsWrapper(virConnectPtr conn,
                                             virNWFilterBindingPtr **bindings,
                                             unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 4005000
    assert(0); // Caller should have checked version
#else
    return virConnectListAllNWFilterBindings(conn, bindings, flags);
#endif
}

virNWFilterBindingPtr virNWFilterBindingCreateXMLWrapper(virConnectPtr conn,
                                                         const char *xml,
                                                         unsigned int flags)
{
#if LIBVIR_VERSION_NUMBER < 4005000
    assert(0); // Caller should have checked version
#else
    return virNWFilterBindingCreateXML(conn, xml, flags);
#endif
}

virNWFilterBindingPtr virNWFilterBindingLookupByPortDevWrapper(virConnectPtr conn,
                                                               const char *portdev)
{
#if LIBVIR_VERSION_NUMBER < 4005000
    assert(0); // Caller should have checked version
#else
    return virNWFilterBindingLookupByPortDev(conn, portdev);
#endif
}


*/
import "C"
