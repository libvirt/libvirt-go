#ifndef LIBVIRT_GO_DOMAIN_COMPAT_H__
#define LIBVIRT_GO_ERROR_COMPAT_H__

/* 1.2.12 */

#ifndef VIR_DOMAIN_DEFINE_VALIDATE
#define VIR_DOMAIN_DEFINE_VALIDATE 1 << 0
#endif

#ifndef VIR_DOMAIN_START_VALIDATE
#define VIR_DOMAIN_START_VALIDATE 1 << 4
#endif


/* 1.2.14 */

#ifndef VIR_DOMAIN_CONTROL_ERROR_REASON_NONE
#define VIR_DOMAIN_CONTROL_ERROR_REASON_NONE 0
#endif

#ifndef VIR_DOMAIN_CONTROL_ERROR_REASON_UNKNOWN
#define VIR_DOMAIN_CONTROL_ERROR_REASON_UNKNOWN 1
#endif

#ifndef VIR_DOMAIN_CONTROL_ERROR_REASON_MONITOR
#define VIR_DOMAIN_CONTROL_ERROR_REASON_MONITOR 2
#endif

#ifndef VIR_DOMAIN_CONTROL_ERROR_REASON_INTERNAL
#define VIR_DOMAIN_CONTROL_ERROR_REASON_INTERNAL 3
#endif

#ifndef VIR_DOMAIN_INTERFACE_ADDRESSES_SRC_LEASE
#define VIR_DOMAIN_INTERFACE_ADDRESSES_SRC_LEASE 0
#endif

#ifndef VIR_DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT
#define VIR_DOMAIN_INTERFACE_ADDRESSES_SRC_AGENT 1
#endif

#ifndef VIR_DOMAIN_PAUSED_STARTING_UP
#define VIR_DOMAIN_PAUSED_STARTING_UP 11
#endif

#if LIBVIR_VERSION_NUMBER < 1002014
typedef struct _virDomainIOThreadInfo virDomainIOThreadInfo;
typedef virDomainIOThreadInfo *virDomainIOThreadInfoPtr;
struct _virDomainIOThreadInfo {
    unsigned int iothread_id;          /* IOThread ID */
    unsigned char *cpumap;             /* CPU map for thread. A pointer to an */
                                       /* array of real CPUs (in 8-bit bytes) */
    int cpumaplen;                     /* cpumap size */
};

typedef struct _virDomainInterfaceIPAddress virDomainIPAddress;
typedef virDomainIPAddress *virDomainIPAddressPtr;
struct _virDomainInterfaceIPAddress {
    int type;                /* virIPAddrType */
    char *addr;              /* IP address */
    unsigned int prefix;     /* IP address prefix */
};

typedef struct _virDomainInterface virDomainInterface;
typedef virDomainInterface *virDomainInterfacePtr;
struct _virDomainInterface {
    char *name;                     /* interface name */
    char *hwaddr;                   /* hardware address, may be NULL */
    unsigned int naddrs;            /* number of items in @addrs */
    virDomainIPAddressPtr addrs;    /* array of IP addresses */
};
#endif

int virDomainInterfaceAddressesCompat(virDomainPtr dom,
				      virDomainInterfacePtr **ifaces,
				      unsigned int source,
				      unsigned int flags);

void virDomainInterfaceFreeCompat(virDomainInterfacePtr iface);

void virDomainIOThreadInfoFreeCompat(virDomainIOThreadInfoPtr info);

int virDomainGetIOThreadInfoCompat(virDomainPtr domain,
				   virDomainIOThreadInfoPtr **info,
				   unsigned int flags);
int virDomainPinIOThreadCompat(virDomainPtr domain,
			       unsigned int iothread_id,
			       unsigned char *cpumap,
			       int maplen,
			       unsigned int flags);


/* 1.2.15 */

#ifndef VIR_DOMAIN_JOB_DOWNTIME_NET
#define VIR_DOMAIN_JOB_DOWNTIME_NET "downtime_net"
#endif

#ifndef VIR_DOMAIN_JOB_TIME_ELAPSED_NET
#define VIR_DOMAIN_JOB_TIME_ELAPSED_NET "time_elapsed_net"
#endif

#ifndef VIR_DOMAIN_EVENT_ID_DEVICE_ADDED
#define VIR_DOMAIN_EVENT_ID_DEVICE_ADDED 19
#endif


int virDomainAddIOThreadCompat(virDomainPtr domain,
			       unsigned int iothread_id,
			       unsigned int flags);
int virDomainDelIOThreadCompat(virDomainPtr domain,
			       unsigned int iothread_id,
			       unsigned int flags);


/* 1.2.16 */

#ifndef VIR_DOMAIN_PASSWORD_ENCRYPTED
#define VIR_DOMAIN_PASSWORD_ENCRYPTED 1 << 0
#endif

int virDomainSetUserPasswordCompat(virDomainPtr dom,
				   const char *user,
				   const char *password,
				   unsigned int flags);


/* 1.2.17 */

#ifndef VIR_DOMAIN_EVENT_WATCHDOG_INJECTNMI
#define VIR_DOMAIN_EVENT_WATCHDOG_INJECTNMI 6
#endif

#ifndef VIR_MIGRATE_PARAM_MIGRATE_DISKS
#define VIR_MIGRATE_PARAM_MIGRATE_DISKS "migrate_disks"
#endif


/* 1.2.19 */

#ifndef VIR_DOMAIN_BANDWIDTH_IN_FLOOR
#define VIR_DOMAIN_BANDWIDTH_IN_FLOOR "inbound.floor"
#endif

#ifndef VIR_DOMAIN_EVENT_DEFINED_RENAMED
#define VIR_DOMAIN_EVENT_DEFINED_RENAMED 2
#endif

#ifndef VIR_DOMAIN_EVENT_UNDEFINED_RENAMED
#define VIR_DOMAIN_EVENT_UNDEFINED_RENAMED 1
#endif

int virDomainRenameCompat(virDomainPtr dom,
			  const char *new_name,
			  unsigned int flags);


/* 1.3.1 */

#ifndef VIR_DOMAIN_JOB_MEMORY_DIRTY_RATE
#define VIR_DOMAIN_JOB_MEMORY_DIRTY_RATE "memory_dirty_rate"
#endif

#ifndef VIR_DOMAIN_JOB_MEMORY_ITERATION
#define VIR_DOMAIN_JOB_MEMORY_ITERATION "memory_iteration"
#endif


/* 1.3.2 */

#ifndef VIR_DOMAIN_EVENT_ID_MIGRATION_ITERATION
#define VIR_DOMAIN_EVENT_ID_MIGRATION_ITERATION 20
#endif


/* 1.3.3 */

#ifndef VIR_DOMAIN_EVENT_DEFINED_FROM_SNAPSHOT
#define VIR_DOMAIN_EVENT_DEFINED_FROM_SNAPSHOT 3
#endif

#ifndef VIR_DOMAIN_EVENT_RESUMED_POSTCOPY
#define VIR_DOMAIN_EVENT_RESUMED_POSTCOPY 3
#endif

#ifndef VIR_DOMAIN_EVENT_SUSPENDED_POSTCOPY
#define VIR_DOMAIN_EVENT_SUSPENDED_POSTCOPY 7
#endif

#ifndef VIR_DOMAIN_EVENT_SUSPENDED_POSTCOPY_FAILED
#define VIR_DOMAIN_EVENT_SUSPENDED_POSTCOPY_FAILED 8
#endif

#ifndef VIR_DOMAIN_PAUSED_POSTCOPY
#define VIR_DOMAIN_PAUSED_POSTCOPY 12
#endif

#ifndef VIR_DOMAIN_PAUSED_POSTCOPY_FAILED
#define VIR_DOMAIN_PAUSED_POSTCOPY_FAILED 13
#endif

#ifndef VIR_DOMAIN_RUNNING_POSTCOPY
#define VIR_DOMAIN_RUNNING_POSTCOPY 10
#endif

#ifndef VIR_DOMAIN_SCHEDULER_GLOBAL_PERIOD
#define VIR_DOMAIN_SCHEDULER_GLOBAL_PERIOD "global_period"
#endif

#ifndef VIR_DOMAIN_SCHEDULER_GLOBAL_QUOTA
#define VIR_DOMAIN_SCHEDULER_GLOBAL_QUOTA "global_quota"
#endif

#ifndef VIR_DOMAIN_STATS_PERF
#define VIR_DOMAIN_STATS_PERF (1 << 6)
#endif

#ifndef VIR_MIGRATE_PARAM_DISKS_PORT
#define VIR_MIGRATE_PARAM_DISKS_PORT "disks_port"
#endif

#ifndef VIR_PERF_PARAM_CMT
#define VIR_PERF_PARAM_CMT "cmt"
#endif

#ifndef VIR_MIGRATE_POSTCOPY
#define VIR_MIGRATE_POSTCOPY (1 << 15)
#endif

#ifndef VIR_DOMAIN_EVENT_ID_JOB_COMPLETED
#define VIR_DOMAIN_EVENT_ID_JOB_COMPLETED 21
#endif

#ifndef VIR_DOMAIN_TUNABLE_CPU_GLOBAL_PERIOD
#define VIR_DOMAIN_TUNABLE_CPU_GLOBAL_PERIOD "cputune.global_period"
#endif

#ifndef VIR_DOMAIN_TUNABLE_CPU_GLOBAL_QUOTA
#define VIR_DOMAIN_TUNABLE_CPU_GLOBAL_QUOTA "cputune.global_quota"
#endif

int virDomainGetPerfEventsCompat(virDomainPtr dom,
				 virTypedParameterPtr *params,
				 int *nparams,
				 unsigned int flags);
int virDomainSetPerfEventsCompat(virDomainPtr dom,
				 virTypedParameterPtr params,
				 int nparams,
				 unsigned int flags);
int virDomainMigrateStartPostCopyCompat(virDomainPtr domain,
					unsigned int flags);


/* 1.3.4 */

#ifndef VIR_MIGRATE_PARAM_COMPRESSION
#define VIR_MIGRATE_PARAM_COMPRESSION  "compression"
#endif

#ifndef VIR_MIGRATE_PARAM_COMPRESSION_MT_THREADS
#define VIR_MIGRATE_PARAM_COMPRESSION_MT_THREADS "compression.mt.threads"
#endif

#ifndef VIR_MIGRATE_PARAM_COMPRESSION_MT_DTHREADS
#define VIR_MIGRATE_PARAM_COMPRESSION_MT_DTHREADS "compression.mt.dthreads"
#endif

#ifndef VIR_MIGRATE_PARAM_COMPRESSION_MT_LEVEL
#define VIR_MIGRATE_PARAM_COMPRESSION_MT_LEVEL "compression.mt.level"
#endif

#ifndef VIR_MIGRATE_PARAM_COMPRESSION_XBZRLE_CACHE
#define VIR_MIGRATE_PARAM_COMPRESSION_XBZRLE_CACHE "compression.xbzrle.cache"
#endif

#ifndef VIR_MIGRATE_PARAM_PERSIST_XML
#define VIR_MIGRATE_PARAM_PERSIST_XML "persistent_xml"
#endif

#ifndef VIR_DOMAIN_EVENT_ID_DEVICE_REMOVAL_FAILED
#define VIR_DOMAIN_EVENT_ID_DEVICE_REMOVAL_FAILED 22
#endif


/* 1.3.5 */

#ifndef VIR_PERF_PARAM_MBML
#define VIR_PERF_PARAM_MBML "mbml"
#endif

#ifndef VIR_PERF_PARAM_MBMT
#define VIR_PERF_PARAM_MBMT "mbmt"
#endif


/* 2.0.0 */

#ifndef VIR_DOMAIN_JOB_AUTO_CONVERGE_THROTTLE
#define VIR_DOMAIN_JOB_AUTO_CONVERGE_THROTTLE "auto_converge_throttle"
#endif

#ifndef VIR_MIGRATE_PARAM_AUTO_CONVERGE_INITIAL
#define VIR_MIGRATE_PARAM_AUTO_CONVERGE_INITIAL "auto_converge.initial"
#endif

#ifndef VIR_MIGRATE_PARAM_AUTO_CONVERGE_INCREMENT
#define VIR_MIGRATE_PARAM_AUTO_CONVERGE_INCREMENT "auto_converge.increment"
#endif

int virDomainGetGuestVcpusCompat(virDomainPtr domain,
				 virTypedParameterPtr *params,
				 unsigned int *nparams,
				 unsigned int flags);

int virDomainSetGuestVcpusCompat(virDomainPtr domain,
				 const char *cpumap,
				 int state,
				 unsigned int flags);


/* 2.1.0 */

#ifndef VIR_DOMAIN_MEMORY_STAT_USABLE
#define VIR_DOMAIN_MEMORY_STAT_USABLE 8
#endif

#ifndef VIR_DOMAIN_MEMORY_STAT_LAST_UPDATE
#define VIR_DOMAIN_MEMORY_STAT_LAST_UPDATE 9
#endif

/* 2.2.0 */

#ifndef VIR_DOMAIN_SCHEDULER_IOTHREAD_PERIOD
#define VIR_DOMAIN_SCHEDULER_IOTHREAD_PERIOD "iothread_period"
#endif

#ifndef VIR_DOMAIN_SCHEDULER_IOTHREAD_QUOTA
#define VIR_DOMAIN_SCHEDULER_IOTHREAD_QUOTA "iothread_quota"
#endif

#ifndef VIR_DOMAIN_TUNABLE_CPU_IOTHREAD_PERIOD
#define VIR_DOMAIN_TUNABLE_CPU_IOTHREAD_PERIOD "cputune.iothread_period"
#endif

#ifndef VIR_DOMAIN_TUNABLE_CPU_IOTHREAD_QUOTA
# define VIR_DOMAIN_TUNABLE_CPU_IOTHREAD_QUOTA "cputune.iothread_quota"
#endif


/* 2.3.0 */

#ifndef VIR_DOMAIN_UNDEFINE_KEEP_NVRAM
#define VIR_DOMAIN_UNDEFINE_KEEP_NVRAM (1 << 3)
#endif

#ifndef VIR_PERF_PARAM_CACHE_MISSES
#define VIR_PERF_PARAM_CACHE_MISSES "cache_misses"
#endif

#ifndef VIR_PERF_PARAM_CACHE_REFERENCES
#define VIR_PERF_PARAM_CACHE_REFERENCES "cache_references"
#endif

#ifndef VIR_PERF_PARAM_INSTRUCTIONS
#define VIR_PERF_PARAM_INSTRUCTIONS "instructions"
#endif

#ifndef VIR_PERF_PARAM_CPU_CYCLES
#define VIR_PERF_PARAM_CPU_CYCLES "cpu_cycles"
#endif


/* 2.4.0 */

#ifndef VIR_DOMAIN_BLOCK_IOTUNE_READ_BYTES_SEC_MAX_LENGTH
#define VIR_DOMAIN_BLOCK_IOTUNE_READ_BYTES_SEC_MAX_LENGTH "read_bytes_sec_max_length"
#endif

#ifndef VIR_DOMAIN_BLOCK_IOTUNE_READ_IOPS_SEC_MAX_LENGTH
#define VIR_DOMAIN_BLOCK_IOTUNE_READ_IOPS_SEC_MAX_LENGTH "read_iops_sec_max_length"
#endif

#ifndef VIR_DOMAIN_BLOCK_IOTUNE_TOTAL_BYTES_SEC_MAX_LENGTH
#define VIR_DOMAIN_BLOCK_IOTUNE_TOTAL_BYTES_SEC_MAX_LENGTH "total_bytes_sec_max_length"
#endif

#ifndef VIR_DOMAIN_BLOCK_IOTUNE_TOTAL_IOPS_SEC_MAX_LENGTH
#define VIR_DOMAIN_BLOCK_IOTUNE_TOTAL_IOPS_SEC_MAX_LENGTH "total_iops_sec_max_length"
#endif

#ifndef VIR_DOMAIN_BLOCK_IOTUNE_WRITE_BYTES_SEC_MAX_LENGTH
#define VIR_DOMAIN_BLOCK_IOTUNE_WRITE_BYTES_SEC_MAX_LENGTH "write_bytes_sec_max_length"
#endif

#ifndef VIR_DOMAIN_BLOCK_IOTUNE_WRITE_IOPS_SEC_MAX_LENGTH
#define VIR_DOMAIN_BLOCK_IOTUNE_WRITE_IOPS_SEC_MAX_LENGTH "write_iopcs_sec_max_length"
#endif

#ifndef VIR_DOMAIN_TUNABLE_BLKDEV_TOTAL_BYTES_SEC_MAX_LENGTH
#define VIR_DOMAIN_TUNABLE_BLKDEV_TOTAL_BYTES_SEC_MAX_LENGTH "blkdeviotune.total_bytes_sec_max_length"
#endif

#ifndef VIR_DOMAIN_TUNABLE_BLKDEV_READ_BYTES_SEC_MAX_LENGTH
#define VIR_DOMAIN_TUNABLE_BLKDEV_READ_BYTES_SEC_MAX_LENGTH "blkdeviotune.read_bytes_sec_max_length"
#endif

#ifndef VIR_DOMAIN_TUNABLE_BLKDEV_WRITE_BYTES_SEC_MAX_LENGTH
#define VIR_DOMAIN_TUNABLE_BLKDEV_WRITE_BYTES_SEC_MAX_LENGTH "blkdeviotune.write_bytes_sec_max_length"
#endif

#ifndef VIR_DOMAIN_TUNABLE_BLKDEV_TOTAL_IOPS_SEC_MAX_LENGTH
#define VIR_DOMAIN_TUNABLE_BLKDEV_TOTAL_IOPS_SEC_MAX_LENGTH "blkdeviotune.total_iops_sec_max_length"
#endif

#ifndef VIR_DOMAIN_TUNABLE_BLKDEV_READ_IOPS_SEC_MAX_LENGTH
#define VIR_DOMAIN_TUNABLE_BLKDEV_READ_IOPS_SEC_MAX_LENGTH "blkdeviotune.read_iops_sec_max_length"
#endif

#ifndef VIR_DOMAIN_TUNABLE_BLKDEV_WRITE_IOPS_SEC_MAX_LENGTH
#define VIR_DOMAIN_TUNABLE_BLKDEV_WRITE_IOPS_SEC_MAX_LENGTH "blkdeviotune.write_iops_sec_max_length"
#endif

#ifndef VIR_DOMAIN_VCPU_HOTPLUGGABLE
#define VIR_DOMAIN_VCPU_HOTPLUGGABLE (1 << 4)
#endif

#endif /* LIBVIRT_GO_ERROR_COMPAT_H__ */
