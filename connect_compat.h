#ifndef LIBVIRT_GO_CONNECT_COMPAT_H_
#define LIBVIRT_GO_CONNECT_COMPAT_H_

/* 1.2.12 */

#ifndef VIR_CONNECT_GET_ALL_DOMAINS_STATS_BACKING
#define VIR_CONNECT_GET_ALL_DOMAINS_STATS_BACKING 1 << 30
#endif

virDomainPtr virDomainDefineXMLFlagsCompat(virConnectPtr conn,
					   const char *xml,
					   unsigned int flags);

/* 1.2.14 */

#ifndef VIR_CONNECT_BASELINE_CPU_MIGRATABLE
#define VIR_CONNECT_BASELINE_CPU_MIGRATABLE 1 << 1
#endif

#endif /* LIBVIRT_GO_CONNECT_COMPAT_H_ */
