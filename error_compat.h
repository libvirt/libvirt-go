#ifndef LIBVIRT_GO_ERROR_COMPAT_H__
#define LIBVIRT_GO_ERROR_COMPAT_H__

/* 1.3.0 */

#ifndef VIR_FROM_LOGGING
#define VIR_FROM_LOGGING 63
#endif

/* 1.3.2 */

#ifndef VIR_FROM_XENXL
#define VIR_FROM_XENXL 64
#endif


/* 1.3.3 */

#ifndef VIR_FROM_PERF
#define VIR_FROM_PERF 65
#endif

#ifndef VIR_ERR_AUTH_UNAVAILABLE
#define VIR_ERR_AUTH_UNAVAILABLE 94
#endif

#ifndef VIR_ERR_NO_SERVER
#define VIR_ERR_NO_SERVER 95
#endif


/* 1.3.5 */

#ifndef VIR_ERR_NO_CLIENT
#define VIR_ERR_NO_CLIENT 96
#endif


/* 2.3.0 */

#ifndef VIR_ERR_AGENT_UNSYNCED
#define VIR_ERR_AGENT_UNSYNCED 97
#endif

/* 2.5.0 */

#ifndef VIR_ERR_LIBSSH
#define VIR_ERR_LIBSSH 98
#endif

#ifndef VIR_FROM_LIBSSH
#define VIR_FROM_LIBSSH 66
#endif

#endif /* LIBVIRT_GO_ERROR_COMPAT_H__ */
