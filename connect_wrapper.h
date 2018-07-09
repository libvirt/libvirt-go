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

#ifndef LIBVIRT_GO_CONNECT_WRAPPER_H__
#define LIBVIRT_GO_CONNECT_WRAPPER_H__

#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include "connect_compat.h"

void
closeCallbackHelper(virConnectPtr conn,
                    int reason,
                    void *opaque);

int
virConnectRegisterCloseCallbackHelper(virConnectPtr c,
                                      virConnectCloseFunc cb,
                                      long goCallbackId);

virConnectPtr
virConnectOpenAuthWrapper(const char *name,
                          int *credtype,
                          uint ncredtype,
                          int callbackID,
                          unsigned int flags);

int
virNodeGetFreePagesWrapper(virConnectPtr conn,
                           unsigned int npages,
                           unsigned int *pages,
                           int startcell,
                           unsigned int cellcount,
                           unsigned long long *counts,
                           unsigned int flags);

char *
virConnectGetDomainCapabilitiesWrapper(virConnectPtr conn,
                                       const char *emulatorbin,
                                       const char *arch,
                                       const char *machine,
                                       const char *virttype,
                                       unsigned int flags);

int
virConnectGetAllDomainStatsWrapper(virConnectPtr conn,
                                   unsigned int stats,
                                   virDomainStatsRecordPtr **retStats,
                                   unsigned int flags);

int

virDomainListGetStatsWrapper(virDomainPtr *doms,
                             unsigned int stats,
                             virDomainStatsRecordPtr **retStats,
                             unsigned int flags);

void
virDomainStatsRecordListFreeWrapper(virDomainStatsRecordPtr *stats);

int
virNodeAllocPagesWrapper(virConnectPtr conn,
                         unsigned int npages,
                         unsigned int *pageSizes,
                         unsigned long long *pageCounts,
                         int startCell,
                         unsigned int cellCount,
                         unsigned int flags);

virDomainPtr
virDomainDefineXMLFlagsWrapper(virConnectPtr conn,
                               const char *xml,
                               unsigned int flags);

virStoragePoolPtr
virStoragePoolLookupByTargetPathWrapper(virConnectPtr conn,
                                        const char *path);

char *
virConnectBaselineHypervisorCPUWrapper(virConnectPtr conn,
                                       const char *emulator,
                                       const char *arch,
                                       const char *machine,
                                       const char *virttype,
                                       const char **xmlCPUs,
                                       unsigned int ncpus,
                                       unsigned int flags);

int
virConnectCompareHypervisorCPUWrapper(virConnectPtr conn,
                                      const char *emulator,
                                      const char *arch,
                                      const char *machine,
                                      const char *virttype,
                                      const char *xmlCPU,
                                      unsigned int flags);

int
virNodeGetSEVInfoWrapper(virConnectPtr conn,
                         virTypedParameterPtr *params,
                         int *nparams,
                         unsigned int flags);

int
virConnectListAllNWFilterBindingsWrapper(virConnectPtr conn,
                                         virNWFilterBindingPtr **bindings,
                                         unsigned int flags);

virNWFilterBindingPtr
virNWFilterBindingCreateXMLWrapper(virConnectPtr conn,
                                   const char *xml,
                                   unsigned int flags);

virNWFilterBindingPtr
virNWFilterBindingLookupByPortDevWrapper(virConnectPtr conn,
                                         const char *portdev);

#endif /* LIBVIRT_GO_CONNECT_WRAPPER_H__ */
