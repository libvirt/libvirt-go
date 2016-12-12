#ifndef LIBVIRT_GO_STORAGE_POOL_COMPAT_H__
#define LIBVIRT_GO_STORAGE_POOL_COMPAT_H__

/* 1.3.1 */

#ifndef VIR_STORAGE_POOL_CREATE_NORMAL
#define VIR_STORAGE_POOL_CREATE_NORMAL 0
#endif

#ifndef VIR_STORAGE_POOL_CREATE_WITH_BUILD
#define VIR_STORAGE_POOL_CREATE_WITH_BUILD 1 << 0
#endif

#ifndef VIR_STORAGE_POOL_CREATE_WITH_BUILD_OVERWRITE
#define VIR_STORAGE_POOL_CREATE_WITH_BUILD_OVERWRITE 1 << 1
#endif

#ifndef VIR_STORAGE_POOL_CREATE_WITH_BUILD_NO_OVERWRITE
#define VIR_STORAGE_POOL_CREATE_WITH_BUILD_NO_OVERWRITE 1 << 2
#endif


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
