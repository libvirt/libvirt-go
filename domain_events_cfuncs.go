package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include "domain_events_cfuncs.h"
#include "callbacks_cfuncs.h"
#include <stdint.h>

extern void domainEventLifecycleCallback(virConnectPtr, virDomainPtr, int, int, int);
void domainEventLifecycleCallback_cgo(virConnectPtr c, virDomainPtr d,
                                     int event, int detail, void *data)
{
    domainEventLifecycleCallback(c, d, event, detail, (int)(intptr_t)data);
}

extern void domainEventGenericCallback(virConnectPtr, virDomainPtr, int);
void domainEventGenericCallback_cgo(virConnectPtr c, virDomainPtr d, void *data)
{
    domainEventGenericCallback(c, d, (int)(intptr_t)data);
}

extern void domainEventRTCChangeCallback(virConnectPtr, virDomainPtr, long long, int);
void domainEventRTCChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                     long long utcoffset, void *data)
{
    domainEventRTCChangeCallback(c, d, utcoffset, (int)(intptr_t)data);
}

extern void domainEventWatchdogCallback(virConnectPtr, virDomainPtr, int, int);
void domainEventWatchdogCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    int action, void *data)
{
    domainEventWatchdogCallback(c, d, action, (int)(intptr_t)data);
}

extern void domainEventIOErrorCallback(virConnectPtr, virDomainPtr, const char *, const char *, int, int);
void domainEventIOErrorCallback_cgo(virConnectPtr c, virDomainPtr d,
                                   const char *srcPath, const char *devAlias,
                                   int action, void *data)
{
    domainEventIOErrorCallback(c, d, srcPath, devAlias, action, (int)(intptr_t)data);
}

extern void domainEventGraphicsCallback(virConnectPtr, virDomainPtr, int, const virDomainEventGraphicsAddress *,
                                        const virDomainEventGraphicsAddress *, const char *,
                                        const virDomainEventGraphicsSubject *, int);
void domainEventGraphicsCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    int phase, const virDomainEventGraphicsAddress *local,
                                    const virDomainEventGraphicsAddress *remote,
                                    const char *authScheme,
                                    const virDomainEventGraphicsSubject *subject, void *data)
{
    domainEventGraphicsCallback(c, d, phase, local, remote, authScheme, subject, (int)(intptr_t)data);
}

extern void domainEventIOErrorReasonCallback(virConnectPtr, virDomainPtr, const char *, const char *,
                                             int, const char *, int);
void domainEventIOErrorReasonCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         const char *srcPath, const char *devAlias,
                                         int action, const char *reason, void *data)
{
    domainEventIOErrorReasonCallback(c, d, srcPath, devAlias, action, reason, (int)(intptr_t)data);
}

extern void domainEventBlockJobCallback(virConnectPtr, virDomainPtr, const char *, int, int, int);
void domainEventBlockJobCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    const char *disk, int type, int status, void *data)
{
    domainEventBlockJobCallback(c, d, disk, type, status, (int)(intptr_t)data);
}

extern void domainEventDiskChangeCallback(virConnectPtr, virDomainPtr, const char *, const char *,
                                          const char *, int, int);
void domainEventDiskChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                      const char *oldSrcPath, const char *newSrcPath,
                                      const char *devAlias, int reason, void *data)
{
    domainEventDiskChangeCallback(c, d, oldSrcPath, newSrcPath, devAlias, reason, (int)(intptr_t)data);
}

extern void domainEventTrayChangeCallback(virConnectPtr, virDomainPtr, const char *, int, int);
void domainEventTrayChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                      const char *devAlias, int reason, void *data)
{
    domainEventTrayChangeCallback(c, d, devAlias, reason, (int)(intptr_t)data);
}

extern void domainEventReasonCallback(virConnectPtr, virDomainPtr, int, int);
void domainEventReasonCallback_cgo(virConnectPtr c, virDomainPtr d,
                                  int reason, void *data)
{
    domainEventReasonCallback(c, d, reason, (int)(intptr_t)data);
}

extern void domainEventBalloonChangeCallback(virConnectPtr, virDomainPtr, unsigned long long, int);
void domainEventBalloonChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         unsigned long long actual, void *data)
{
    domainEventBalloonChangeCallback(c, d, actual, (int)(intptr_t)data);
}

extern void domainEventDeviceRemovedCallback(virConnectPtr, virDomainPtr, const char *, int);
void domainEventDeviceRemovedCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         const char *devAlias, void *data)
{
    domainEventDeviceRemovedCallback(c, d, devAlias, (int)(intptr_t)data);
}

int virConnectDomainEventRegisterAny_cgo(virConnectPtr c,  virDomainPtr d,
                                         int eventID, virConnectDomainEventGenericCallback cb,
                                         long goCallbackId) {
    void* id = (void*)goCallbackId;
    return virConnectDomainEventRegisterAny(c, d, eventID, cb, id, freeGoCallback_cgo);
}

*/
import "C"
