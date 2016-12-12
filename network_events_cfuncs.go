package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <assert.h>
#include "network_compat.h"
#include "network_events_cfuncs.h"
#include "callbacks_cfuncs.h"
#include <stdint.h>

extern void networkEventLifecycleCallback(virConnectPtr, virNetworkPtr, int, int, int);
void networkEventLifecycleCallback_cgo(virConnectPtr c, virNetworkPtr d,
                                     int event, int detail, void *data)
{
    networkEventLifecycleCallback(c, d, event, detail, (int)(intptr_t)data);
}

int virConnectNetworkEventRegisterAny_cgo(virConnectPtr c,  virNetworkPtr d,
                                         int eventID, virConnectNetworkEventGenericCallback cb,
                                         long goCallbackId) {
    void* id = (void*)goCallbackId;
#if LIBVIR_VERSION_NUMBER < 1002001
    assert(0); // Caller should have checked version
#else
    return virConnectNetworkEventRegisterAny(c, d, eventID, cb, id, freeGoCallback_cgo);
#endif
}

*/
import "C"
