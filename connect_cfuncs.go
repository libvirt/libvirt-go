package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include "connect_cfuncs.h"
#include "callbacks_cfuncs.h"

extern void closeCallback(virConnectPtr, int, long);
void closeCallback_cgo(virConnectPtr conn, int reason, void *opaque)
{
    closeCallback(conn, reason, (long)opaque);
}

int virConnectRegisterCloseCallback_cgo(virConnectPtr c, virConnectCloseFunc cb, long goCallbackId)
{
    void *id = (void*)goCallbackId;
    return virConnectRegisterCloseCallback(c, cb, id, freeGoCallback_cgo);
}

#include <stdio.h>

extern int connectAuthCallback(virConnectCredentialPtr, unsigned int, int);
int connectAuthCallback_cgo(virConnectCredentialPtr cred, unsigned int ncred, void *cbdata)
{
    int *callbackID = cbdata;

    return connectAuthCallback(cred, ncred, *callbackID);
}

virConnectPtr virConnectOpenAuthWrap(const char *name, int *credtype, uint ncredtype, int callbackID, unsigned int flags)
{
    virConnectAuth auth = {
       .credtype = credtype,
       .ncredtype = ncredtype,
       .cb = connectAuthCallback_cgo,
       .cbdata = &callbackID,
    };

    return virConnectOpenAuth(name, &auth, flags);
}

*/
import "C"
