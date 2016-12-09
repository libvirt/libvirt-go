#ifndef STORAGE_POOL_EVENTS_CFUNCS_H__
#define STORAGE_POOL_EVENTS_CFUNCS_H__

void storagePoolEventLifecycleCallback_cgo(virConnectPtr c, virStoragePoolPtr d,
					   int event, int detail, void* data);

int virConnectStoragePoolEventRegisterAny_cgo(virConnectPtr c,  virStoragePoolPtr d,
					      int eventID, virConnectStoragePoolEventGenericCallback cb,
					      long goCallbackId);


#endif /* STORAGE_POOL_EVENTS_CFUNCS_H__ */
