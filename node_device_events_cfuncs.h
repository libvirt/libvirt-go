#ifndef NODE_DEVICE_EVENTS_CFUNCS_H__
#define NODE_DEVICE_EVENTS_CFUNCS_H__

void nodeDeviceEventLifecycleCallback_cgo(virConnectPtr c, virNodeDevicePtr d,
					  int event, int detail, void* data);

void nodeDeviceEventGenericCallback_cgo(virConnectPtr c, virNodeDevicePtr d, void* data);

int virConnectNodeDeviceEventRegisterAny_cgo(virConnectPtr c,  virNodeDevicePtr d,
					     int eventID, virConnectNodeDeviceEventGenericCallback cb,
					     long goCallbackId);


#endif /* NODE_DEVICE_EVENTS_CFUNCS_H__ */
