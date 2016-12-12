#ifndef LIBVIRT_GO_ERROR_COMPAT_H__
#define LIBVIRT_GO_ERROR_COMPAT_H__

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
