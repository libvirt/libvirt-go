#ifndef LIBVIRT_GO_NODE_DEVICE_COMPAT_H__
#define LIBVIRT_GO_NODE_DEVICE_COMPAT_H__

/* 2.2.0 */

#ifndef VIR_NODE_DEVICE_EVENT_ID_LIFECYCLE
#define VIR_NODE_DEVICE_EVENT_ID_LIFECYCLE 0
#endif

#ifndef VIR_NODE_DEVICE_EVENT_ID_UPDATE
#define VIR_NODE_DEVICE_EVENT_ID_UPDATE 1
#endif

#ifndef VIR_NODE_DEVICE_EVENT_CREATED
#define VIR_NODE_DEVICE_EVENT_CREATED 0
#endif

#ifndef VIR_NODE_DEVICE_EVENT_DELETED
#define VIR_NODE_DEVICE_EVENT_DELETED 1
#endif

#if LIBVIR_VERSION_NUMBER < 2002000
typedef void (*virConnectNodeDeviceEventGenericCallback)(virConnectPtr conn,
                                                         virNodeDevicePtr dev,
                                                         void *opaque);
#endif

int virConnectNodeDeviceEventDeregisterAnyCompat(virConnectPtr conn,
						 int callbackID);


#endif /* LIBVIRT_GO_NODE_DEVICE_COMPAT_H__ */
