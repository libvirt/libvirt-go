#ifndef LIBVIRT_GO_STORAGE_POOL_COMPAT_H__
#define LIBVIRT_GO_STORAGE_POOL_COMPAT_H__

/* 2.0.0 */

#ifndef VIVIR_STORAGE_POOL_EVENT_DEFINED
#define VIR_STORAGE_POOL_EVENT_DEFINED 0
#endif

#ifndef VIR_STORAGE_POOL_EVENT_UNDEFINED
#define VIR_STORAGE_POOL_EVENT_UNDEFINED 1
#endif

#ifndef VIR_STORAGE_POOL_EVENT_STARTED
#define VIR_STORAGE_POOL_EVENT_STARTED 2
#endif

#ifndef VIR_STORAGE_POOL_EVENT_STOPPED
#define VIR_STORAGE_POOL_EVENT_STOPPED 3
#endif

#ifndef VIR_STORAGE_POOL_EVENT_ID_LIFECYCLE
#define VIR_STORAGE_POOL_EVENT_ID_LIFECYCLE 0
#endif

#ifndef VIR_STORAGE_POOL_EVENT_ID_REFRESH
#define VIR_STORAGE_POOL_EVENT_ID_REFRESH 1
#endif

#if LIBVIR_VERSION_NUMBER < 2000000
typedef void (*virConnectStoragePoolEventGenericCallback)(virConnectPtr conn,
                                                          virStoragePoolPtr pool,
                                                          void *opaque);
#endif

int virConnectStoragePoolEventDeregisterAnyCompat(virConnectPtr conn,
						  int callbackID);


#endif /* LIBVIRT_GO_STORAGE_POOL_COMPAT_H__ */
