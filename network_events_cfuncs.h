#ifndef NETWORK_EVENTS_CFUNCS_H__
#define NETWORK_EVENTS_CFUNCS_H__

void networkEventLifecycleCallback_cgo(virConnectPtr c, virNetworkPtr d,
                                     int event, int detail, void* data);

int virConnectNetworkEventRegisterAny_cgo(virConnectPtr c,  virNetworkPtr d,
                                         int eventID, virConnectNetworkEventGenericCallback cb,
                                         long goCallbackId);


#endif /* NETWORK_EVENTS_CFUNCS_H__ */
