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

#ifndef LIBVIRT_GO_DOMAIN_WRAPPER_H__
#define LIBVIRT_GO_DOMAIN_WRAPPER_H__

#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include "domain_compat.h"

int
virDomainCoreDumpWithFormatWrapper(virDomainPtr domain,
                                   const char *to,
                                   unsigned int dumpformat,
                                   unsigned int flags);


int
virDomainGetTimeWrapper(virDomainPtr dom,
                        long long *seconds,
                        unsigned int *nseconds,
                        unsigned int flags);

int
virDomainSetTimeWrapper(virDomainPtr dom,
                        long long seconds,
                        unsigned int nseconds,
                        unsigned int flags);

int
virDomainFSFreezeWrapper(virDomainPtr dom,
                         const char **mountpoints,
                         unsigned int nmountpoints,
                         unsigned int flags);

int
virDomainFSThawWrapper(virDomainPtr dom,
                       const char **mountpoints,
                       unsigned int nmountpoints,
                       unsigned int flags);


int
virDomainBlockCopyWrapper(virDomainPtr dom,
                          const char *disk,
                          const char *destxml,
                          virTypedParameterPtr params,
                          int nparams,
                          unsigned int flags);

int
virDomainOpenGraphicsFDWrapper(virDomainPtr dom,
                               unsigned int idx,
                               unsigned int flags);


void
virDomainFSInfoFreeWrapper(virDomainFSInfoPtr info);

int
virDomainGetFSInfoWrapper(virDomainPtr dom,
                          virDomainFSInfoPtr **info,
                          unsigned int flags);


int
virDomainInterfaceAddressesWrapper(virDomainPtr dom,
                                   virDomainInterfacePtr **ifaces,
                                   unsigned int source,
                                   unsigned int flags);

void
virDomainInterfaceFreeWrapper(virDomainInterfacePtr iface);

void
virDomainIOThreadInfoFreeWrapper(virDomainIOThreadInfoPtr info);

int
virDomainGetIOThreadInfoWrapper(virDomainPtr domain,
                                virDomainIOThreadInfoPtr **info,
                                unsigned int flags);
int
virDomainPinIOThreadWrapper(virDomainPtr domain,
                            unsigned int iothread_id,
                            unsigned char *cpumap,
                            int maplen,
                            unsigned int flags);

int
virDomainAddIOThreadWrapper(virDomainPtr domain,
                            unsigned int iothread_id,
                            unsigned int flags);

int
virDomainDelIOThreadWrapper(virDomainPtr domain,
                            unsigned int iothread_id,
                            unsigned int flags);

int
virDomainSetUserPasswordWrapper(virDomainPtr dom,
                                const char *user,
                                const char *password,
                                unsigned int flags);

int
virDomainRenameWrapper(virDomainPtr dom,
                       const char *new_name,
                       unsigned int flags);

int
virDomainGetPerfEventsWrapper(virDomainPtr dom,
                              virTypedParameterPtr *params,
                              int *nparams,
                              unsigned int flags);

int
virDomainSetPerfEventsWrapper(virDomainPtr dom,
                              virTypedParameterPtr params,
                              int nparams,
                              unsigned int flags);

int
virDomainMigrateStartPostCopyWrapper(virDomainPtr domain,
                                     unsigned int flags);

int
virDomainGetGuestVcpusWrapper(virDomainPtr domain,
                              virTypedParameterPtr *params,
                              unsigned int *nparams,
                              unsigned int flags);

int
virDomainSetGuestVcpusWrapper(virDomainPtr domain,
                              const char *cpumap,
                              int state,
                              unsigned int flags);

int
virDomainSetVcpuWrapper(virDomainPtr domain,
                        const char *cpumap,
                        int state,
                        unsigned int flags);

int
virDomainSetBlockThresholdWrapper(virDomainPtr domain,
                                  const char *dev,
                                  unsigned long long threshold,
                                  unsigned int flags);

int
virDomainMigrateGetMaxDowntimeWrapper(virDomainPtr domain,
                                      unsigned long long *downtime,
                                      unsigned int flags);

char *
virDomainManagedSaveGetXMLDescWrapper(virDomainPtr domain,
                                      unsigned int flags);

int
virDomainManagedSaveDefineXMLWrapper(virDomainPtr domain,
                                     const char *dxml,
                                     unsigned int flags);

int
virDomainSetLifecycleActionWrapper(virDomainPtr domain,
                                   unsigned int type,
                                   unsigned int action,
                                   unsigned int flags);

int
virDomainDetachDeviceAliasWrapper(virDomainPtr domain,
                                  const char *alias,
                                  unsigned int flags);

int
virDomainGetLaunchSecurityInfoWrapper(virDomainPtr domain,
                                      virTypedParameterPtr *params,
                                      int *nparams,
                                      unsigned int flags);

#endif /* LIBVIRT_GO_DOMAIN_WRAPPER_H__ */
