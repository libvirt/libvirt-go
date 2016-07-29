package libvirt

/*
#cgo CFLAGS: -Wno-implicit-function-declaration
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>

void domainEventLifecycleCallback_cgo(virConnectPtr c, virDomainPtr d,
                                     int event, int detail, void *data)
{
    domainEventLifecycleCallback(c, d, event, detail, data);
}

void domainEventGenericCallback_cgo(virConnectPtr c, virDomainPtr d, void *data)
{
    domainEventGenericCallback(c, d, data);
}

void domainEventRTCChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                     long long utcoffset, void *data)
{
    domainEventRTCChangeCallback(c, d, utcoffset, data);
}

void domainEventWatchdogCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    int action, void *data)
{
    domainEventWatchdogCallback(c, d, action, data);
}

void domainEventIOErrorCallback_cgo(virConnectPtr c, virDomainPtr d,
                                   const char *srcPath, const char *devAlias,
                                   int action, void *data)
{
    domainEventIOErrorCallback(c, d, srcPath, devAlias, action, data);
}

void domainEventGraphicsCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    int phase, const virDomainEventGraphicsAddress *local,
                                    const virDomainEventGraphicsAddress *remote,
                                    const char *authScheme,
                                    const virDomainEventGraphicsSubject *subject, void *data)
{
    domainEventGraphicsCallback(c, d, phase, local, remote, authScheme, subject, data);
}

void domainEventIOErrorReasonCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         const char *srcPath, const char *devAlias,
                                         int action, const char *reason, void *data)
{
    domainEventIOErrorReasonCallback(c, d, srcPath, devAlias, action, reason, data);
}

void domainEventBlockJobCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    const char *disk, int type, int status, void *data)
{
    domainEventIOErrorReasonCallback(c, d, disk, type, status, data);
}

void domainEventDiskChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                      const char *oldSrcPath, const char *newSrcPath,
                                      const char *devAlias, int reason, void *data)
{
    domainEventDiskChangeCallback(c, d, oldSrcPath, newSrcPath, devAlias, reason, data);
}

void domainEventTrayChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                      const char *devAlias, int reason, void *data)
{
    domainEventTrayChangeCallback(c, d, devAlias, reason, data);
}

void domainEventReasonCallback_cgo(virConnectPtr c, virDomainPtr d,
                                  int reason, void *data)
{
    domainEventReasonCallback(c, d, reason, data);
}

void domainEventBalloonChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         unsigned long long actual, void *data)
{
    domainEventBalloonChangeCallback(c, d, actual, data);
}

void domainEventDeviceRemovedCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         const char *devAlias, void *data)
{
    domainEventDeviceRemovedCallback(c, d, devAlias, data);
}

void freeGoCallback_cgo(void* goCallbackId) {
   freeCallbackId((long)goCallbackId);
}

int virConnectDomainEventRegisterAny_cgo(virConnectPtr c,  virDomainPtr d,
                                         int eventID, virConnectDomainEventGenericCallback cb,
                                         long goCallbackId) {
    void* id = (void*)goCallbackId;
    return virConnectDomainEventRegisterAny(c, d, eventID, cb, id, freeGoCallback_cgo);
}

void errorGlobalCallback_cgo(void *userData, virErrorPtr error)
{
    globalErrorCallback(error);
}

void errorConnCallback_cgo(void *userData, virErrorPtr error)
{
    connErrorCallback((long)userData, error);
}

void virConnSetErrorFunc_cgo(virConnectPtr c, long goCallbackId, virErrorFunc cb)
{
    void* id = (void*)goCallbackId;
    virConnSetErrorFunc(c, id, cb);
}


*/
import "C"
