#ifndef LIBVIRT_GO_CONNECT_COMPAT_H_
#define LIBVIRT_GO_CONNECT_COMPAT_H_

/* 1.2.6 */

#ifndef VIR_CONNECT_COMPARE_CPU_FAIL_INCOMPATIBLE
#define VIR_CONNECT_COMPARE_CPU_FAIL_INCOMPATIBLE 1 << 0
#endif

int virNodeGetFreePagesCompat(virConnectPtr conn,
			      unsigned int npages,
			      unsigned int *pages,
			      int startcell,
			      unsigned int cellcount,
			      unsigned long long *counts,
			      unsigned int flags);


/* 1.2.7 */

char * virConnectGetDomainCapabilitiesCompat(virConnectPtr conn,
					     const char *emulatorbin,
					     const char *arch,
					     const char *machine,
					     const char *virttype,
					     unsigned int flags);


/* 1.2.8 */

#ifndef VIR_CONNECT_GET_ALL_DOMAINS_STATS_ACTIVE
#define VIR_CONNECT_GET_ALL_DOMAINS_STATS_ACTIVE 1 << 0
#endif

#ifndef VIR_CONNECT_GET_ALL_DOMAINS_STATS_INACTIVE
#define VIR_CONNECT_GET_ALL_DOMAINS_STATS_INACTIVE 1 << 1
#endif

#ifndef VIR_CONNECT_GET_ALL_DOMAINS_STATS_PERSISTENT
#define VIR_CONNECT_GET_ALL_DOMAINS_STATS_PERSISTENT 1 << 2
#endif

#ifndef VIR_CONNECT_GET_ALL_DOMAINS_STATS_TRANSIENT
#define VIR_CONNECT_GET_ALL_DOMAINS_STATS_TRANSIENT 1 << 3
#endif

#ifndef VIR_CONNECT_GET_ALL_DOMAINS_STATS_RUNNING
#define VIR_CONNECT_GET_ALL_DOMAINS_STATS_RUNNING 1 << 4
#endif

#ifndef VIR_CONNECT_GET_ALL_DOMAINS_STATS_PAUSED
#define VIR_CONNECT_GET_ALL_DOMAINS_STATS_PAUSED 1 << 5
#endif

#ifndef VIR_CONNECT_GET_ALL_DOMAINS_STATS_SHUTOFF
#define VIR_CONNECT_GET_ALL_DOMAINS_STATS_SHUTOFF 1 << 6
#endif

#ifndef VIR_CONNECT_GET_ALL_DOMAINS_STATS_OTHER
#define VIR_CONNECT_GET_ALL_DOMAINS_STATS_OTHER 1 << 7
#endif

#ifndef VIR_CONNECT_GET_ALL_DOMAINS_STATS_ENFORCE_STATS
#define VIR_CONNECT_GET_ALL_DOMAINS_STATS_ENFORCE_STATS 1U << 31
#endif

#ifndef VIR_CONNECT_LIST_STORAGE_POOLS_ZFS
#define VIR_CONNECT_LIST_STORAGE_POOLS_ZFS 1 << 17
#endif

#if LIBVIR_VERSION_NUMBER < 1002008
typedef struct _virDomainStatsRecord virDomainStatsRecord;
typedef virDomainStatsRecord *virDomainStatsRecordPtr;
struct _virDomainStatsRecord {
    virDomainPtr dom;
    virTypedParameterPtr params;
    int nparams;
};
#endif

int virConnectGetAllDomainStatsCompat(virConnectPtr conn,
				      unsigned int stats,
				      virDomainStatsRecordPtr **retStats,
				      unsigned int flags);

int virDomainListGetStatsCompat(virDomainPtr *doms,
				unsigned int stats,
				virDomainStatsRecordPtr **retStats,
				unsigned int flags);

void virDomainStatsRecordListFreeCompat(virDomainStatsRecordPtr *stats);


/* 1.2.9 */
#ifndef VIR_NODE_ALLOC_PAGES_ADD
#define VIR_NODE_ALLOC_PAGES_ADD 0
#endif

#ifndef VIR_NODE_ALLOC_PAGES_SET
#define VIR_NODE_ALLOC_PAGES_SET 1 << 0
#endif

int virNodeAllocPagesCompat(virConnectPtr conn,
			    unsigned int npages,
			    unsigned int *pageSizes,
			    unsigned long long *pageCounts,
			    int startCell,
			    unsigned int cellCount,
			    unsigned int flags);


/* 1.2.11 */

#ifndef VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_CONNECTED
#define VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_CONNECTED 1
#endif

#ifndef VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_DISCONNECTED
#define VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_DISCONNECTED 2
#endif

#ifndef VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_UNKNOWN
#define VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_UNKNOWN 0
#endif

#ifndef VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_DOMAIN_STARTED
#define VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_DOMAIN_STARTED 1
#endif

#ifndef VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_CHANNEL
#define VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_CHANNEL 2
#endif


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
