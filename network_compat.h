#ifndef LIBVIRT_GO_NETWORK_COMPAT_H__
#define LIBVIRT_GO_NETWORK_COMPAT_H__


/* 1.2.5 */

#ifndef VIR_IP_ADDR_TYPE_IPV4
#define VIR_IP_ADDR_TYPE_IPV4 0
#endif

#ifndef VIR_IP_ADDR_TYPE_IPV6
#define VIR_IP_ADDR_TYPE_IPV6 1
#endif

#if LIBVIR_VERSION_NUMBER < 1002006
typedef struct _virNetworkDHCPLease virNetworkDHCPLease;
typedef virNetworkDHCPLease *virNetworkDHCPLeasePtr;
struct _virNetworkDHCPLease {
    char *iface;                /* Network interface name */
    long long expirytime;       /* Seconds since epoch */
    int type;                   /* virIPAddrType */
    char *mac;                  /* MAC address */
    char *iaid;                 /* IAID */
    char *ipaddr;               /* IP address */
    unsigned int prefix;        /* IP address prefix */
    char *hostname;             /* Hostname */
    char *clientid;             /* Client ID or DUID */
};
#endif

void virNetworkDHCPLeaseFreeCompat(virNetworkDHCPLeasePtr lease);

int virNetworkGetDHCPLeasesCompat(virNetworkPtr network,
				  const char *mac,
				  virNetworkDHCPLeasePtr **leases,
				  unsigned int flags);

#endif /* LIBVIRT_GO_NETWORK_COMPAT_H__ */
