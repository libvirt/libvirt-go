package libvirt

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include "callbacks_cfuncs.h"

extern void freeCallbackId(long);
void freeGoCallback_cgo(void* goCallbackId) {
   freeCallbackId((long)goCallbackId);
}

*/
import "C"
