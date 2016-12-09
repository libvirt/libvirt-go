package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
#include <string.h>
#include "connect_cfuncs.h"
#include "callbacks_cfuncs.h"

extern void closeCallback(virConnectPtr, int, long);
void closeCallback_cgo(virConnectPtr conn, int reason, void *opaque)
{
    closeCallback(conn, reason, (long)opaque);
}

int authCb(virConnectCredentialPtr cred, unsigned int ncred, void *cbdata)
{
	int i;

    auth_cb_data *data = (auth_cb_data*)cbdata;
    for (i = 0; i < ncred; i++) {
        if (cred[i].type == VIR_CRED_AUTHNAME) {
            cred[i].result = strndup(data->username, data->username_len);
            if (cred[i].result == NULL)
                return -1;
            cred[i].resultlen = strlen(cred[i].result);
        }
        else if (cred[i].type == VIR_CRED_PASSPHRASE) {
            cred[i].result = strndup(data->passphrase, data->passphrase_len);
            if (cred[i].result == NULL)
                return -1;
            cred[i].resultlen = strlen(cred[i].result);
        }
    }
    return 0;
}

auth_cb_data* authData(char* username, uint username_len, char* passphrase, uint passphrase_len) {
    auth_cb_data * data = malloc(sizeof(auth_cb_data));
    data->username = username;
    data->username_len = username_len;
    data->passphrase = passphrase;
    data->passphrase_len = passphrase_len;
    return data;
}

int* authMechs() {
    int* authMechs = malloc(2*sizeof(VIR_CRED_AUTHNAME));
    authMechs[0] = VIR_CRED_AUTHNAME;
    authMechs[1] = VIR_CRED_PASSPHRASE;
    return authMechs;
}

int virConnectRegisterCloseCallback_cgo(virConnectPtr c, virConnectCloseFunc cb, long goCallbackId)
{
    void *id = (void*)goCallbackId;
    return virConnectRegisterCloseCallback(c, cb, id, freeGoCallback_cgo);
}

*/
import "C"
