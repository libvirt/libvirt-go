package libvirt

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>

void ignoreErrorFunc(void *userData, virErrorPtr error) {
     // no-op
}
*/
import "C"

import (
	"fmt"
)

func init() {
	C.virSetErrorFunc(nil, (C.virErrorFunc)(C.ignoreErrorFunc))
	C.virInitialize()
}

type ErrorLevel int

const (
	VIR_ERR_NONE    = ErrorLevel(C.VIR_ERR_NONE)
	VIR_ERR_WARNING = ErrorLevel(C.VIR_ERR_WARNING)
	VIR_ERR_ERROR   = ErrorLevel(C.VIR_ERR_ERROR)
)

type ErrorNumber int

const (
	VIR_ERR_OK = ErrorNumber(C.VIR_ERR_OK)

	// internal error
	VIR_ERR_INTERNAL_ERROR = ErrorNumber(C.VIR_ERR_INTERNAL_ERROR)

	// memory allocation failure
	VIR_ERR_NO_MEMORY = ErrorNumber(C.VIR_ERR_NO_MEMORY)

	// no support for this function
	VIR_ERR_NO_SUPPORT = ErrorNumber(C.VIR_ERR_NO_SUPPORT)

	// could not resolve hostname
	VIR_ERR_UNKNOWN_HOST = ErrorNumber(C.VIR_ERR_UNKNOWN_HOST)

	// can't connect to hypervisor
	VIR_ERR_NO_CONNECT = ErrorNumber(C.VIR_ERR_NO_CONNECT)

	// invalid connection object
	VIR_ERR_INVALID_CONN = ErrorNumber(C.VIR_ERR_INVALID_CONN)

	// invalid domain object
	VIR_ERR_INVALID_DOMAIN = ErrorNumber(C.VIR_ERR_INVALID_DOMAIN)

	// invalid function argument
	VIR_ERR_INVALID_ARG = ErrorNumber(C.VIR_ERR_INVALID_ARG)

	// a command to hypervisor failed
	VIR_ERR_OPERATION_FAILED = ErrorNumber(C.VIR_ERR_OPERATION_FAILED)

	// a HTTP GET command to failed
	VIR_ERR_GET_FAILED = ErrorNumber(C.VIR_ERR_GET_FAILED)

	// a HTTP POST command to failed
	VIR_ERR_POST_FAILED = ErrorNumber(C.VIR_ERR_POST_FAILED)

	// unexpected HTTP error code
	VIR_ERR_HTTP_ERROR = ErrorNumber(C.VIR_ERR_HTTP_ERROR)

	// failure to serialize an S-Expr
	VIR_ERR_SEXPR_SERIAL = ErrorNumber(C.VIR_ERR_SEXPR_SERIAL)

	// could not open Xen hypervisor control
	VIR_ERR_NO_XEN = ErrorNumber(C.VIR_ERR_NO_XEN)

	// failure doing an hypervisor call
	VIR_ERR_XEN_CALL = ErrorNumber(C.VIR_ERR_XEN_CALL)

	// unknown OS type
	VIR_ERR_OS_TYPE = ErrorNumber(C.VIR_ERR_OS_TYPE)

	// missing kernel information
	VIR_ERR_NO_KERNEL = ErrorNumber(C.VIR_ERR_NO_KERNEL)

	// missing root device information
	VIR_ERR_NO_ROOT = ErrorNumber(C.VIR_ERR_NO_ROOT)

	// missing source device information
	VIR_ERR_NO_SOURCE = ErrorNumber(C.VIR_ERR_NO_SOURCE)

	// missing target device information
	VIR_ERR_NO_TARGET = ErrorNumber(C.VIR_ERR_NO_TARGET)

	// missing domain name information
	VIR_ERR_NO_NAME = ErrorNumber(C.VIR_ERR_NO_NAME)

	// missing domain OS information
	VIR_ERR_NO_OS = ErrorNumber(C.VIR_ERR_NO_OS)

	// missing domain devices information
	VIR_ERR_NO_DEVICE = ErrorNumber(C.VIR_ERR_NO_DEVICE)

	// could not open Xen Store control
	VIR_ERR_NO_XENSTORE = ErrorNumber(C.VIR_ERR_NO_XENSTORE)

	// too many drivers registered
	VIR_ERR_DRIVER_FULL = ErrorNumber(C.VIR_ERR_DRIVER_FULL)

	// not supported by the drivers (DEPRECATED)
	VIR_ERR_CALL_FAILED = ErrorNumber(C.VIR_ERR_CALL_FAILED)

	// an XML description is not well formed or broken
	VIR_ERR_XML_ERROR = ErrorNumber(C.VIR_ERR_XML_ERROR)

	// the domain already exist
	VIR_ERR_DOM_EXIST = ErrorNumber(C.VIR_ERR_DOM_EXIST)

	// operation forbidden on read-only connections
	VIR_ERR_OPERATION_DENIED = ErrorNumber(C.VIR_ERR_OPERATION_DENIED)

	// failed to open a conf file
	VIR_ERR_OPEN_FAILED = ErrorNumber(C.VIR_ERR_OPEN_FAILED)

	// failed to read a conf file
	VIR_ERR_READ_FAILED = ErrorNumber(C.VIR_ERR_READ_FAILED)

	// failed to parse a conf file
	VIR_ERR_PARSE_FAILED = ErrorNumber(C.VIR_ERR_PARSE_FAILED)

	// failed to parse the syntax of a conf file
	VIR_ERR_CONF_SYNTAX = ErrorNumber(C.VIR_ERR_CONF_SYNTAX)

	// failed to write a conf file
	VIR_ERR_WRITE_FAILED = ErrorNumber(C.VIR_ERR_WRITE_FAILED)

	// detail of an XML error
	VIR_ERR_XML_DETAIL = ErrorNumber(C.VIR_ERR_XML_DETAIL)

	// invalid network object
	VIR_ERR_INVALID_NETWORK = ErrorNumber(C.VIR_ERR_INVALID_NETWORK)

	// the network already exist
	VIR_ERR_NETWORK_EXIST = ErrorNumber(C.VIR_ERR_NETWORK_EXIST)

	// general system call failure
	VIR_ERR_SYSTEM_ERROR = ErrorNumber(C.VIR_ERR_SYSTEM_ERROR)

	// some sort of RPC error
	VIR_ERR_RPC = ErrorNumber(C.VIR_ERR_RPC)

	// error from a GNUTLS call
	VIR_ERR_GNUTLS_ERROR = ErrorNumber(C.VIR_ERR_GNUTLS_ERROR)

	// failed to start network
	VIR_WAR_NO_NETWORK = ErrorNumber(C.VIR_WAR_NO_NETWORK)

	// domain not found or unexpectedly disappeared
	VIR_ERR_NO_DOMAIN = ErrorNumber(C.VIR_ERR_NO_DOMAIN)

	// network not found
	VIR_ERR_NO_NETWORK = ErrorNumber(C.VIR_ERR_NO_NETWORK)

	// invalid MAC address
	VIR_ERR_INVALID_MAC = ErrorNumber(C.VIR_ERR_INVALID_MAC)

	// authentication failed
	VIR_ERR_AUTH_FAILED = ErrorNumber(C.VIR_ERR_AUTH_FAILED)

	// invalid storage pool object
	VIR_ERR_INVALID_STORAGE_POOL = ErrorNumber(C.VIR_ERR_INVALID_STORAGE_POOL)

	// invalid storage vol object
	VIR_ERR_INVALID_STORAGE_VOL = ErrorNumber(C.VIR_ERR_INVALID_STORAGE_VOL)

	// failed to start storage
	VIR_WAR_NO_STORAGE = ErrorNumber(C.VIR_WAR_NO_STORAGE)

	// storage pool not found
	VIR_ERR_NO_STORAGE_POOL = ErrorNumber(C.VIR_ERR_NO_STORAGE_POOL)

	// storage volume not found
	VIR_ERR_NO_STORAGE_VOL = ErrorNumber(C.VIR_ERR_NO_STORAGE_VOL)

	// failed to start node driver
	VIR_WAR_NO_NODE = ErrorNumber(C.VIR_WAR_NO_NODE)

	// invalid node device object
	VIR_ERR_INVALID_NODE_DEVICE = ErrorNumber(C.VIR_ERR_INVALID_NODE_DEVICE)

	// node device not found
	VIR_ERR_NO_NODE_DEVICE = ErrorNumber(C.VIR_ERR_NO_NODE_DEVICE)

	// security model not found
	VIR_ERR_NO_SECURITY_MODEL = ErrorNumber(C.VIR_ERR_NO_SECURITY_MODEL)

	// operation is not applicable at this time
	VIR_ERR_OPERATION_INVALID = ErrorNumber(C.VIR_ERR_OPERATION_INVALID)

	// failed to start interface driver
	VIR_WAR_NO_INTERFACE = ErrorNumber(C.VIR_WAR_NO_INTERFACE)

	// interface driver not running
	VIR_ERR_NO_INTERFACE = ErrorNumber(C.VIR_ERR_NO_INTERFACE)

	// invalid interface object
	VIR_ERR_INVALID_INTERFACE = ErrorNumber(C.VIR_ERR_INVALID_INTERFACE)

	// more than one matching interface found
	VIR_ERR_MULTIPLE_INTERFACES = ErrorNumber(C.VIR_ERR_MULTIPLE_INTERFACES)

	// failed to start nwfilter driver
	VIR_WAR_NO_NWFILTER = ErrorNumber(C.VIR_WAR_NO_NWFILTER)

	// invalid nwfilter object
	VIR_ERR_INVALID_NWFILTER = ErrorNumber(C.VIR_ERR_INVALID_NWFILTER)

	// nw filter pool not found
	VIR_ERR_NO_NWFILTER = ErrorNumber(C.VIR_ERR_NO_NWFILTER)

	// nw filter pool not found
	VIR_ERR_BUILD_FIREWALL = ErrorNumber(C.VIR_ERR_BUILD_FIREWALL)

	// failed to start secret storage
	VIR_WAR_NO_SECRET = ErrorNumber(C.VIR_WAR_NO_SECRET)

	// invalid secret
	VIR_ERR_INVALID_SECRET = ErrorNumber(C.VIR_ERR_INVALID_SECRET)

	// secret not found
	VIR_ERR_NO_SECRET = ErrorNumber(C.VIR_ERR_NO_SECRET)

	// unsupported configuration construct
	VIR_ERR_CONFIG_UNSUPPORTED = ErrorNumber(C.VIR_ERR_CONFIG_UNSUPPORTED)

	// timeout occurred during operation
	VIR_ERR_OPERATION_TIMEOUT = ErrorNumber(C.VIR_ERR_OPERATION_TIMEOUT)

	// a migration worked, but making the VM persist on the dest host failed
	VIR_ERR_MIGRATE_PERSIST_FAILED = ErrorNumber(C.VIR_ERR_MIGRATE_PERSIST_FAILED)

	// a synchronous hook script failed
	VIR_ERR_HOOK_SCRIPT_FAILED = ErrorNumber(C.VIR_ERR_HOOK_SCRIPT_FAILED)

	// invalid domain snapshot
	VIR_ERR_INVALID_DOMAIN_SNAPSHOT = ErrorNumber(C.VIR_ERR_INVALID_DOMAIN_SNAPSHOT)

	// domain snapshot not found
	VIR_ERR_NO_DOMAIN_SNAPSHOT = ErrorNumber(C.VIR_ERR_NO_DOMAIN_SNAPSHOT)

	// stream pointer not valid
	VIR_ERR_INVALID_STREAM = ErrorNumber(C.VIR_ERR_INVALID_STREAM)

	// valid API use but unsupported by the given driver
	VIR_ERR_ARGUMENT_UNSUPPORTED = ErrorNumber(C.VIR_ERR_ARGUMENT_UNSUPPORTED)

	// storage pool probe failed
	VIR_ERR_STORAGE_PROBE_FAILED = ErrorNumber(C.VIR_ERR_STORAGE_PROBE_FAILED)

	// storage pool already built
	VIR_ERR_STORAGE_POOL_BUILT = ErrorNumber(C.VIR_ERR_STORAGE_POOL_BUILT)

	// force was not requested for a risky domain snapshot revert
	VIR_ERR_SNAPSHOT_REVERT_RISKY = ErrorNumber(C.VIR_ERR_SNAPSHOT_REVERT_RISKY)

	// operation on a domain was canceled/aborted by user
	VIR_ERR_OPERATION_ABORTED = ErrorNumber(C.VIR_ERR_OPERATION_ABORTED)

	// authentication cancelled
	VIR_ERR_AUTH_CANCELLED = ErrorNumber(C.VIR_ERR_AUTH_CANCELLED)

	// The metadata is not present
	VIR_ERR_NO_DOMAIN_METADATA = ErrorNumber(C.VIR_ERR_NO_DOMAIN_METADATA)

	// Migration is not safe
	VIR_ERR_MIGRATE_UNSAFE = ErrorNumber(C.VIR_ERR_MIGRATE_UNSAFE)

	// integer overflow
	VIR_ERR_OVERFLOW = ErrorNumber(C.VIR_ERR_OVERFLOW)

	// action prevented by block copy job
	VIR_ERR_BLOCK_COPY_ACTIVE = ErrorNumber(C.VIR_ERR_BLOCK_COPY_ACTIVE)

	// The requested operation is not supported
	VIR_ERR_OPERATION_UNSUPPORTED = ErrorNumber(C.VIR_ERR_OPERATION_UNSUPPORTED)

	// error in ssh transport driver
	VIR_ERR_SSH = ErrorNumber(C.VIR_ERR_SSH)

	// guest agent is unresponsive, not running or not usable
	VIR_ERR_AGENT_UNRESPONSIVE = ErrorNumber(C.VIR_ERR_AGENT_UNRESPONSIVE)

	// resource is already in use
	VIR_ERR_RESOURCE_BUSY = ErrorNumber(C.VIR_ERR_RESOURCE_BUSY)

	// operation on the object/resource was denied
	VIR_ERR_ACCESS_DENIED = ErrorNumber(C.VIR_ERR_ACCESS_DENIED)

	// error from a dbus service
	VIR_ERR_DBUS_SERVICE = ErrorNumber(C.VIR_ERR_DBUS_SERVICE)

	// the storage vol already exists
	VIR_ERR_STORAGE_VOL_EXIST = ErrorNumber(C.VIR_ERR_STORAGE_VOL_EXIST)

	// given CPU is incompatible with host CPU
	VIR_ERR_CPU_INCOMPATIBLE = ErrorNumber(C.VIR_ERR_CPU_INCOMPATIBLE)

	// XML document doesn't validate against schema
	VIR_ERR_XML_INVALID_SCHEMA = ErrorNumber(C.VIR_ERR_XML_INVALID_SCHEMA)

	// Finish API succeeded but it is expected to return NULL */
	VIR_ERR_MIGRATE_FINISH_OK = ErrorNumber(C.VIR_ERR_MIGRATE_FINISH_OK)

	// authentication unavailable
	VIR_ERR_AUTH_UNAVAILABLE = ErrorNumber(C.VIR_ERR_AUTH_UNAVAILABLE)

	// Server was not found
	VIR_ERR_NO_SERVER = ErrorNumber(C.VIR_ERR_NO_SERVER)

	// Client was not found
	VIR_ERR_NO_CLIENT = ErrorNumber(C.VIR_ERR_NO_CLIENT)

	// guest agent replies with wrong id to guest sync command
	VIR_ERR_AGENT_UNSYNCED = ErrorNumber(C.VIR_ERR_AGENT_UNSYNCED)

	// error in libssh transport driver
	VIR_ERR_LIBSSH = ErrorNumber(C.VIR_ERR_LIBSSH)
)

type ErrorDomain int

const (
	VIR_FROM_NONE = ErrorDomain(C.VIR_FROM_NONE)

	// Error at Xen hypervisor layer
	VIR_FROM_XEN = ErrorDomain(C.VIR_FROM_XEN)

	// Error at connection with xend daemon
	VIR_FROM_XEND = ErrorDomain(C.VIR_FROM_XEND)

	// Error at connection with xen store
	VIR_FROM_XENSTORE = ErrorDomain(C.VIR_FROM_XENSTORE)

	// Error in the S-Expression code
	VIR_FROM_SEXPR = ErrorDomain(C.VIR_FROM_SEXPR)

	// Error in the XML code
	VIR_FROM_XML = ErrorDomain(C.VIR_FROM_XML)

	// Error when operating on a domain
	VIR_FROM_DOM = ErrorDomain(C.VIR_FROM_DOM)

	// Error in the XML-RPC code
	VIR_FROM_RPC = ErrorDomain(C.VIR_FROM_RPC)

	// Error in the proxy code; unused since 0.8.6
	VIR_FROM_PROXY = ErrorDomain(C.VIR_FROM_PROXY)

	// Error in the configuration file handling
	VIR_FROM_CONF = ErrorDomain(C.VIR_FROM_CONF)

	// Error at the QEMU daemon
	VIR_FROM_QEMU = ErrorDomain(C.VIR_FROM_QEMU)

	// Error when operating on a network
	VIR_FROM_NET = ErrorDomain(C.VIR_FROM_NET)

	// Error from test driver
	VIR_FROM_TEST = ErrorDomain(C.VIR_FROM_TEST)

	// Error from remote driver
	VIR_FROM_REMOTE = ErrorDomain(C.VIR_FROM_REMOTE)

	// Error from OpenVZ driver
	VIR_FROM_OPENVZ = ErrorDomain(C.VIR_FROM_OPENVZ)

	// Error at Xen XM layer
	VIR_FROM_XENXM = ErrorDomain(C.VIR_FROM_XENXM)

	// Error in the Linux Stats code
	VIR_FROM_STATS_LINUX = ErrorDomain(C.VIR_FROM_STATS_LINUX)

	// Error from Linux Container driver
	VIR_FROM_LXC = ErrorDomain(C.VIR_FROM_LXC)

	// Error from storage driver
	VIR_FROM_STORAGE = ErrorDomain(C.VIR_FROM_STORAGE)

	// Error from network config
	VIR_FROM_NETWORK = ErrorDomain(C.VIR_FROM_NETWORK)

	// Error from domain config
	VIR_FROM_DOMAIN = ErrorDomain(C.VIR_FROM_DOMAIN)

	// Error at the UML driver
	VIR_FROM_UML = ErrorDomain(C.VIR_FROM_UML)

	// Error from node device monitor
	VIR_FROM_NODEDEV = ErrorDomain(C.VIR_FROM_NODEDEV)

	// Error from xen inotify layer
	VIR_FROM_XEN_INOTIFY = ErrorDomain(C.VIR_FROM_XEN_INOTIFY)

	// Error from security framework
	VIR_FROM_SECURITY = ErrorDomain(C.VIR_FROM_SECURITY)

	// Error from VirtualBox driver
	VIR_FROM_VBOX = ErrorDomain(C.VIR_FROM_VBOX)

	// Error when operating on an interface
	VIR_FROM_INTERFACE = ErrorDomain(C.VIR_FROM_INTERFACE)

	// The OpenNebula driver no longer exists. Retained for ABI/API compat only
	VIR_FROM_ONE = ErrorDomain(C.VIR_FROM_ONE)

	// Error from ESX driver
	VIR_FROM_ESX = ErrorDomain(C.VIR_FROM_ESX)

	// Error from IBM power hypervisor
	VIR_FROM_PHYP = ErrorDomain(C.VIR_FROM_PHYP)

	// Error from secret storage
	VIR_FROM_SECRET = ErrorDomain(C.VIR_FROM_SECRET)

	// Error from CPU driver
	VIR_FROM_CPU = ErrorDomain(C.VIR_FROM_CPU)

	// Error from XenAPI
	VIR_FROM_XENAPI = ErrorDomain(C.VIR_FROM_XENAPI)

	// Error from network filter driver
	VIR_FROM_NWFILTER = ErrorDomain(C.VIR_FROM_NWFILTER)

	// Error from Synchronous hooks
	VIR_FROM_HOOK = ErrorDomain(C.VIR_FROM_HOOK)

	// Error from domain snapshot
	VIR_FROM_DOMAIN_SNAPSHOT = ErrorDomain(C.VIR_FROM_DOMAIN_SNAPSHOT)

	// Error from auditing subsystem
	VIR_FROM_AUDIT = ErrorDomain(C.VIR_FROM_AUDIT)

	// Error from sysinfo/SMBIOS
	VIR_FROM_SYSINFO = ErrorDomain(C.VIR_FROM_SYSINFO)

	// Error from I/O streams
	VIR_FROM_STREAMS = ErrorDomain(C.VIR_FROM_STREAMS)

	// Error from VMware driver
	VIR_FROM_VMWARE = ErrorDomain(C.VIR_FROM_VMWARE)

	// Error from event loop impl
	VIR_FROM_EVENT = ErrorDomain(C.VIR_FROM_EVENT)

	// Error from libxenlight driver
	VIR_FROM_LIBXL = ErrorDomain(C.VIR_FROM_LIBXL)

	// Error from lock manager
	VIR_FROM_LOCKING = ErrorDomain(C.VIR_FROM_LOCKING)

	// Error from Hyper-V driver
	VIR_FROM_HYPERV = ErrorDomain(C.VIR_FROM_HYPERV)

	// Error from capabilities
	VIR_FROM_CAPABILITIES = ErrorDomain(C.VIR_FROM_CAPABILITIES)

	// Error from URI handling
	VIR_FROM_URI = ErrorDomain(C.VIR_FROM_URI)

	// Error from auth handling
	VIR_FROM_AUTH = ErrorDomain(C.VIR_FROM_AUTH)

	// Error from DBus
	VIR_FROM_DBUS = ErrorDomain(C.VIR_FROM_DBUS)

	// Error from Parallels
	VIR_FROM_PARALLELS = ErrorDomain(C.VIR_FROM_PARALLELS)

	// Error from Device
	VIR_FROM_DEVICE = ErrorDomain(C.VIR_FROM_DEVICE)

	// Error from libssh2 connection transport
	VIR_FROM_SSH = ErrorDomain(C.VIR_FROM_SSH)

	// Error from lockspace
	VIR_FROM_LOCKSPACE = ErrorDomain(C.VIR_FROM_LOCKSPACE)

	// Error from initctl device communication
	VIR_FROM_INITCTL = ErrorDomain(C.VIR_FROM_INITCTL)

	// Error from identity code
	VIR_FROM_IDENTITY = ErrorDomain(C.VIR_FROM_IDENTITY)

	// Error from cgroups
	VIR_FROM_CGROUP = ErrorDomain(C.VIR_FROM_CGROUP)

	// Error from access control manager
	VIR_FROM_ACCESS = ErrorDomain(C.VIR_FROM_ACCESS)

	// Error from systemd code
	VIR_FROM_SYSTEMD = ErrorDomain(C.VIR_FROM_SYSTEMD)

	// Error from bhyve driver
	VIR_FROM_BHYVE = ErrorDomain(C.VIR_FROM_BHYVE)

	// Error from crypto code
	VIR_FROM_CRYPTO = ErrorDomain(C.VIR_FROM_CRYPTO)

	// Error from firewall
	VIR_FROM_FIREWALL = ErrorDomain(C.VIR_FROM_FIREWALL)

	// Erorr from polkit code
	VIR_FROM_POLKIT = ErrorDomain(C.VIR_FROM_POLKIT)

	// Error from thread utils
	VIR_FROM_THREAD = ErrorDomain(C.VIR_FROM_THREAD)

	// Error from admin backend
	VIR_FROM_ADMIN = ErrorDomain(C.VIR_FROM_ADMIN)

	// Error from log manager
	VIR_FROM_LOGGING = ErrorDomain(C.VIR_FROM_LOGGING)

	// Error from Xen xl config code
	VIR_FROM_XENXL = ErrorDomain(C.VIR_FROM_XENXL)

	// Error from perf
	VIR_FROM_PERF = ErrorDomain(C.VIR_FROM_PERF)

	// Error from libssh
	VIR_FROM_LIBSSH = ErrorDomain(C.VIR_FROM_LIBSSH)
)

type Error struct {
	Code    ErrorNumber
	Domain  ErrorDomain
	Message string
	Level   ErrorLevel
}

func (err Error) Error() string {
	return fmt.Sprintf("[Code-%d] [Domain-%d] %s",
		err.Code, err.Domain, err.Message)
}

var ErrNoError = Error{
	Code:    VIR_ERR_OK,
	Domain:  VIR_FROM_NONE,
	Message: "",
	Level:   VIR_ERR_NONE,
}

func GetLastError() Error {
	err := C.virGetLastError()
	if err == nil {
		return ErrNoError
	}
	virErr := Error{
		Code:    ErrorNumber(err.code),
		Domain:  ErrorDomain(err.domain),
		Message: C.GoString(err.message),
		Level:   ErrorLevel(err.level),
	}
	C.virResetError(err)
	return virErr
}
