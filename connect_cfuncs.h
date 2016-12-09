#ifndef GO_LIBVIRT_H
#define GO_LIBVIRT_H
void closeCallback_cgo(virConnectPtr conn, int reason, void *opaque);
int virConnectRegisterCloseCallback_cgo(virConnectPtr c, virConnectCloseFunc cb, long goCallbackId);

virConnectPtr virConnectOpenAuthWrap(const char *name, int *credtype, uint ncredtype, int callbackID, unsigned int flags);

#endif /* GO_LIBVIRT_H */
