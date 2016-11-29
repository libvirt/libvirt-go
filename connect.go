package libvirt

import (
	"fmt"
	"io/ioutil"
	"reflect"
	"sync"
	"unsafe"
)

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>
#include <libvirt/virterror.h>
#include <stdlib.h>
#include "go_libvirt.h"
*/
import "C"

type VirConnectCloseReason int

const (
	VIR_CONNECT_CLOSE_REASON_ERROR     = VirConnectCloseReason(C.VIR_CONNECT_CLOSE_REASON_ERROR)
	VIR_CONNECT_CLOSE_REASON_EOF       = VirConnectCloseReason(C.VIR_CONNECT_CLOSE_REASON_EOF)
	VIR_CONNECT_CLOSE_REASON_KEEPALIVE = VirConnectCloseReason(C.VIR_CONNECT_CLOSE_REASON_KEEPALIVE)
	VIR_CONNECT_CLOSE_REASON_CLIENT    = VirConnectCloseReason(C.VIR_CONNECT_CLOSE_REASON_CLIENT)
)

type VirConnectListAllDomainsFlags int

const (
	VIR_CONNECT_LIST_DOMAINS_ACTIVE         = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_ACTIVE)
	VIR_CONNECT_LIST_DOMAINS_INACTIVE       = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_INACTIVE)
	VIR_CONNECT_LIST_DOMAINS_PERSISTENT     = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_PERSISTENT)
	VIR_CONNECT_LIST_DOMAINS_TRANSIENT      = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_TRANSIENT)
	VIR_CONNECT_LIST_DOMAINS_RUNNING        = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_RUNNING)
	VIR_CONNECT_LIST_DOMAINS_PAUSED         = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_PAUSED)
	VIR_CONNECT_LIST_DOMAINS_SHUTOFF        = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_SHUTOFF)
	VIR_CONNECT_LIST_DOMAINS_OTHER          = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_OTHER)
	VIR_CONNECT_LIST_DOMAINS_MANAGEDSAVE    = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_MANAGEDSAVE)
	VIR_CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_MANAGEDSAVE)
	VIR_CONNECT_LIST_DOMAINS_AUTOSTART      = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_AUTOSTART)
	VIR_CONNECT_LIST_DOMAINS_NO_AUTOSTART   = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_AUTOSTART)
	VIR_CONNECT_LIST_DOMAINS_HAS_SNAPSHOT   = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_HAS_SNAPSHOT)
	VIR_CONNECT_LIST_DOMAINS_NO_SNAPSHOT    = VirConnectListAllDomainsFlags(C.VIR_CONNECT_LIST_DOMAINS_NO_SNAPSHOT)
)

type VirConnectListAllNetworksFlags int

const (
	VIR_CONNECT_LIST_NETWORKS_INACTIVE     = VirConnectListAllNetworksFlags(C.VIR_CONNECT_LIST_NETWORKS_INACTIVE)
	VIR_CONNECT_LIST_NETWORKS_ACTIVE       = VirConnectListAllNetworksFlags(C.VIR_CONNECT_LIST_NETWORKS_ACTIVE)
	VIR_CONNECT_LIST_NETWORKS_PERSISTENT   = VirConnectListAllNetworksFlags(C.VIR_CONNECT_LIST_NETWORKS_PERSISTENT)
	VIR_CONNECT_LIST_NETWORKS_TRANSIENT    = VirConnectListAllNetworksFlags(C.VIR_CONNECT_LIST_NETWORKS_TRANSIENT)
	VIR_CONNECT_LIST_NETWORKS_AUTOSTART    = VirConnectListAllNetworksFlags(C.VIR_CONNECT_LIST_NETWORKS_AUTOSTART)
	VIR_CONNECT_LIST_NETWORKS_NO_AUTOSTART = VirConnectListAllNetworksFlags(C.VIR_CONNECT_LIST_NETWORKS_NO_AUTOSTART)
)

type VirConnectListAllStoragePoolsFlags int

const (
	VIR_CONNECT_LIST_STORAGE_POOLS_INACTIVE     = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_INACTIVE)
	VIR_CONNECT_LIST_STORAGE_POOLS_ACTIVE       = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_ACTIVE)
	VIR_CONNECT_LIST_STORAGE_POOLS_PERSISTENT   = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_PERSISTENT)
	VIR_CONNECT_LIST_STORAGE_POOLS_TRANSIENT    = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_TRANSIENT)
	VIR_CONNECT_LIST_STORAGE_POOLS_AUTOSTART    = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_AUTOSTART)
	VIR_CONNECT_LIST_STORAGE_POOLS_NO_AUTOSTART = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_NO_AUTOSTART)
	VIR_CONNECT_LIST_STORAGE_POOLS_DIR          = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_DIR)
	VIR_CONNECT_LIST_STORAGE_POOLS_FS           = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_FS)
	VIR_CONNECT_LIST_STORAGE_POOLS_NETFS        = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_NETFS)
	VIR_CONNECT_LIST_STORAGE_POOLS_LOGICAL      = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_LOGICAL)
	VIR_CONNECT_LIST_STORAGE_POOLS_DISK         = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_DISK)
	VIR_CONNECT_LIST_STORAGE_POOLS_ISCSI        = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_ISCSI)
	VIR_CONNECT_LIST_STORAGE_POOLS_SCSI         = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_SCSI)
	VIR_CONNECT_LIST_STORAGE_POOLS_MPATH        = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_MPATH)
	VIR_CONNECT_LIST_STORAGE_POOLS_RBD          = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_RBD)
	VIR_CONNECT_LIST_STORAGE_POOLS_SHEEPDOG     = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_SHEEPDOG)
	VIR_CONNECT_LIST_STORAGE_POOLS_GLUSTER      = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_GLUSTER)
	VIR_CONNECT_LIST_STORAGE_POOLS_ZFS          = VirConnectListAllStoragePoolsFlags(C.VIR_CONNECT_LIST_STORAGE_POOLS_ZFS)
)

type VirConnectBaselineCPUFlags int

const (
	VIR_CONNECT_BASELINE_CPU_EXPAND_FEATURES = VirConnectBaselineCPUFlags(C.VIR_CONNECT_BASELINE_CPU_EXPAND_FEATURES)
	VIR_CONNECT_BASELINE_CPU_MIGRATABLE      = VirConnectBaselineCPUFlags(C.VIR_CONNECT_BASELINE_CPU_MIGRATABLE)
)

type VirConnectCompareCPUFlags int

const (
	VIR_CONNECT_COMPARE_CPU_FAIL_INCOMPATIBLE = VirConnectCompareCPUFlags(C.VIR_CONNECT_COMPARE_CPU_FAIL_INCOMPATIBLE)
)

type VirConnectListAllInterfacesFlags int

const (
	VIR_CONNECT_LIST_INTERFACES_INACTIVE = VirConnectListAllInterfacesFlags(C.VIR_CONNECT_LIST_INTERFACES_INACTIVE)
	VIR_CONNECT_LIST_INTERFACES_ACTIVE   = VirConnectListAllInterfacesFlags(C.VIR_CONNECT_LIST_INTERFACES_ACTIVE)
)

type VirConnectListAllNodeDeviceFlags int

const (
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SYSTEM        = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SYSTEM)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_PCI_DEV       = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_PCI_DEV)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_DEV       = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_DEV)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_INTERFACE = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_USB_INTERFACE)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_NET           = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_NET)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_HOST     = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_HOST)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_TARGET   = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_TARGET)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI          = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_STORAGE       = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_STORAGE)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_FC_HOST       = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_FC_HOST)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_VPORTS        = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_VPORTS)
	VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_GENERIC  = VirConnectListAllNodeDeviceFlags(C.VIR_CONNECT_LIST_NODE_DEVICES_CAP_SCSI_GENERIC)
)

type VirConnectListAllSecrets int

const (
	VIR_CONNECT_LIST_SECRETS_EPHEMERAL    = VirConnectListAllSecrets(C.VIR_CONNECT_LIST_SECRETS_EPHEMERAL)
	VIR_CONNECT_LIST_SECRETS_NO_EPHEMERAL = VirConnectListAllSecrets(C.VIR_CONNECT_LIST_SECRETS_NO_EPHEMERAL)
	VIR_CONNECT_LIST_SECRETS_PRIVATE      = VirConnectListAllSecrets(C.VIR_CONNECT_LIST_SECRETS_PRIVATE)
	VIR_CONNECT_LIST_SECRETS_NO_PRIVATE   = VirConnectListAllSecrets(C.VIR_CONNECT_LIST_SECRETS_NO_PRIVATE)
)

type VirConnectGetAllDomainStatsFlags int

const (
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_ACTIVE        = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_ACTIVE)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_INACTIVE      = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_INACTIVE)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_PERSISTENT    = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_PERSISTENT)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_TRANSIENT     = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_TRANSIENT)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_RUNNING       = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_RUNNING)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_PAUSED        = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_PAUSED)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_SHUTOFF       = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_SHUTOFF)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_OTHER         = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_OTHER)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_BACKING       = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_BACKING)
	VIR_CONNECT_GET_ALL_DOMAINS_STATS_ENFORCE_STATS = VirConnectGetAllDomainStatsFlags(C.VIR_CONNECT_GET_ALL_DOMAINS_STATS_ENFORCE_STATS)
)

type VirConnectFlags int

const (
	VIR_CONNECT_RO         = VirConnectFlags(C.VIR_CONNECT_RO)
	VIR_CONNECT_NO_ALIASES = VirConnectFlags(C.VIR_CONNECT_NO_ALIASES)
)

type VirConnectDomainEventAgentLifecycleState int

const (
	VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_CONNECTED    = VirConnectDomainEventAgentLifecycleState(C.VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_CONNECTED)
	VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_DISCONNECTED = VirConnectDomainEventAgentLifecycleState(C.VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_STATE_DISCONNECTED)
)

type VirConnectDomainEventAgentLifecycleReason int

const (
	VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_UNKNOWN        = VirConnectDomainEventAgentLifecycleReason(C.VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_UNKNOWN)
	VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_DOMAIN_STARTED = VirConnectDomainEventAgentLifecycleReason(C.VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_DOMAIN_STARTED)
	VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_CHANNEL        = VirConnectDomainEventAgentLifecycleReason(C.VIR_CONNECT_DOMAIN_EVENT_AGENT_LIFECYCLE_REASON_CHANNEL)
)

type VirConnectCompareResult int

const (
	VIR_CPU_COMPARE_ERROR        = VirConnectCompareResult(C.VIR_CPU_COMPARE_ERROR)
	VIR_CPU_COMPARE_INCOMPATIBLE = VirConnectCompareResult(C.VIR_CPU_COMPARE_INCOMPATIBLE)
	VIR_CPU_COMPARE_IDENTICAL    = VirConnectCompareResult(C.VIR_CPU_COMPARE_IDENTICAL)
	VIR_CPU_COMPARE_SUPERSET     = VirConnectCompareResult(C.VIR_CPU_COMPARE_SUPERSET)
)

type VirNodeAllocPagesFlags int

const (
	VIR_NODE_ALLOC_PAGES_ADD = VirNodeAllocPagesFlags(C.VIR_NODE_ALLOC_PAGES_ADD)
	VIR_NODE_ALLOC_PAGES_SET = VirNodeAllocPagesFlags(C.VIR_NODE_ALLOC_PAGES_SET)
)

type VirNodeSuspendTarget int

const (
	VIR_NODE_SUSPEND_TARGET_MEM    = VirNodeSuspendTarget(C.VIR_NODE_SUSPEND_TARGET_MEM)
	VIR_NODE_SUSPEND_TARGET_DISK   = VirNodeSuspendTarget(C.VIR_NODE_SUSPEND_TARGET_DISK)
	VIR_NODE_SUSPEND_TARGET_HYBRID = VirNodeSuspendTarget(C.VIR_NODE_SUSPEND_TARGET_HYBRID)
)

type VirNodeGetCPUStatsAllCPUs int

const (
	VIR_NODE_CPU_STATS_ALL_CPUS = VirNodeGetCPUStatsAllCPUs(C.VIR_NODE_CPU_STATS_ALL_CPUS)
)

type VirNodeGetMemoryStatsAllCells int

const (
	VIR_NODE_MEMORY_STATS_ALL_CELLS = VirNodeGetMemoryStatsAllCells(C.VIR_NODE_MEMORY_STATS_ALL_CELLS)
)

type VirConnectCredentialType int

const (
	VIR_CRED_USERNAME     = VirConnectCredentialType(C.VIR_CRED_USERNAME)
	VIR_CRED_AUTHNAME     = VirConnectCredentialType(C.VIR_CRED_AUTHNAME)
	VIR_CRED_LANGUAGE     = VirConnectCredentialType(C.VIR_CRED_LANGUAGE)
	VIR_CRED_CNONCE       = VirConnectCredentialType(C.VIR_CRED_CNONCE)
	VIR_CRED_PASSPHRASE   = VirConnectCredentialType(C.VIR_CRED_PASSPHRASE)
	VIR_CRED_ECHOPROMPT   = VirConnectCredentialType(C.VIR_CRED_ECHOPROMPT)
	VIR_CRED_NOECHOPROMPT = VirConnectCredentialType(C.VIR_CRED_NOECHOPROMPT)
	VIR_CRED_REALM        = VirConnectCredentialType(C.VIR_CRED_REALM)
	VIR_CRED_EXTERNAL     = VirConnectCredentialType(C.VIR_CRED_EXTERNAL)
)

type VirConnection struct {
	ptr C.virConnectPtr
}

// Additional data associated to the connection.
type virConnectionData struct {
	errCallbackId   *int
	closeCallbackId *int
}

var connections map[C.virConnectPtr]*virConnectionData
var connectionsLock sync.RWMutex

func init() {
	connections = make(map[C.virConnectPtr]*virConnectionData)
}

func saveConnectionData(c *VirConnection, d *virConnectionData) {
	if c.ptr == nil {
		return // Or panic?
	}
	connectionsLock.Lock()
	defer connectionsLock.Unlock()
	connections[c.ptr] = d
}

func getConnectionData(c *VirConnection) *virConnectionData {
	connectionsLock.RLock()
	d := connections[c.ptr]
	connectionsLock.RUnlock()
	if d != nil {
		return d
	}
	d = &virConnectionData{}
	saveConnectionData(c, d)
	return d
}

func releaseConnectionData(c *VirConnection) {
	if c.ptr == nil {
		return
	}
	connectionsLock.Lock()
	defer connectionsLock.Unlock()
	delete(connections, c.ptr)
}

func GetVersion() (uint32, error) {
	var version C.ulong
	if err := C.virGetVersion(&version, nil, nil); err < 0 {
		return 0, GetLastError()
	}
	return uint32(version), nil
}

func NewVirConnection(uri string) (VirConnection, error) {
	var cUri *C.char
	if uri != "" {
		cUri = C.CString(uri)
		defer C.free(unsafe.Pointer(cUri))
	}
	ptr := C.virConnectOpen(cUri)
	if ptr == nil {
		return VirConnection{}, GetLastError()
	}
	obj := VirConnection{ptr: ptr}
	return obj, nil
}

func NewVirConnectionWithAuth(uri string, username string, password string) (VirConnection, error) {
	var cUri *C.char

	authMechs := C.authMechs()
	defer C.free(unsafe.Pointer(authMechs))
	cUsername := C.CString(username)
	defer C.free(unsafe.Pointer(cUsername))
	cPassword := C.CString(password)
	defer C.free(unsafe.Pointer(cPassword))
	cbData := C.authData(cUsername, C.uint(len(username)), cPassword, C.uint(len(password)))
	defer C.free(unsafe.Pointer(cbData))

	auth := C.virConnectAuth{
		credtype:  authMechs,
		ncredtype: C.uint(2),
		cb:        C.virConnectAuthCallbackPtr(unsafe.Pointer(C.authCb)),
		cbdata:    unsafe.Pointer(cbData),
	}

	if uri != "" {
		cUri = C.CString(uri)
		defer C.free(unsafe.Pointer(cUri))
	}
	ptr := C.virConnectOpenAuth(cUri, (*C.struct__virConnectAuth)(unsafe.Pointer(&auth)), C.uint(0))
	if ptr == nil {
		return VirConnection{}, GetLastError()
	}
	obj := VirConnection{ptr: ptr}
	return obj, nil
}

func NewVirConnectionReadOnly(uri string) (VirConnection, error) {
	var cUri *C.char
	if uri != "" {
		cUri = C.CString(uri)
		defer C.free(unsafe.Pointer(cUri))
	}
	ptr := C.virConnectOpenReadOnly(cUri)
	if ptr == nil {
		return VirConnection{}, GetLastError()
	}
	obj := VirConnection{ptr: ptr}
	return obj, nil
}

func (c *VirConnection) CloseConnection() (int, error) {
	c.UnsetErrorFunc()
	result := int(C.virConnectClose(c.ptr))
	if result == -1 {
		return result, GetLastError()
	}
	if result == 0 {
		// No more reference to this connection, release data.
		releaseConnectionData(c)
	}
	return result, nil
}

type CloseCallback func(conn VirConnection, reason VirConnectCloseReason, opaque func())
type closeContext struct {
	cb CloseCallback
	f  func()
}

// Register a close callback for the given destination. Only one
// callback per connection is allowed. Setting a callback will remove
// the previous one.
func (c *VirConnection) RegisterCloseCallback(cb CloseCallback, opaque func()) error {
	c.UnregisterCloseCallback()
	context := &closeContext{
		cb: cb,
		f:  opaque,
	}
	goCallbackId := registerCallbackId(context)
	callbackPtr := unsafe.Pointer(C.closeCallback_cgo)
	res := C.virConnectRegisterCloseCallback_cgo(c.ptr, C.virConnectCloseFunc(callbackPtr), C.long(goCallbackId))
	if res != 0 {
		freeCallbackId(goCallbackId)
		return GetLastError()
	}
	connData := getConnectionData(c)
	connData.closeCallbackId = &goCallbackId
	return nil
}

func (c *VirConnection) UnregisterCloseCallback() error {
	connData := getConnectionData(c)
	if connData.closeCallbackId == nil {
		return nil
	}
	callbackPtr := unsafe.Pointer(C.closeCallback_cgo)
	res := C.virConnectUnregisterCloseCallback(c.ptr, C.virConnectCloseFunc(callbackPtr))
	if res != 0 {
		return GetLastError()
	}
	connData.closeCallbackId = nil
	return nil
}

//export closeCallback
func closeCallback(conn C.virConnectPtr, reason VirConnectCloseReason, goCallbackId int) {
	ctx := getCallbackId(goCallbackId)
	switch cctx := ctx.(type) {
	case *closeContext:
		cctx.cb(VirConnection{ptr: conn}, reason, cctx.f)
	default:
		panic("Inappropriate callback type called")
	}
}

func (c *VirConnection) GetCapabilities() (string, error) {
	str := C.virConnectGetCapabilities(c.ptr)
	if str == nil {
		return "", GetLastError()
	}
	capabilities := C.GoString(str)
	C.free(unsafe.Pointer(str))
	return capabilities, nil
}

func (c *VirConnection) GetNodeInfo() (VirNodeInfo, error) {
	ni := VirNodeInfo{}
	var ptr C.virNodeInfo
	result := C.virNodeGetInfo(c.ptr, (*C.virNodeInfo)(unsafe.Pointer(&ptr)))
	if result == -1 {
		return ni, GetLastError()
	}
	ni.ptr = ptr
	return ni, nil
}

func (c *VirConnection) GetHostname() (string, error) {
	str := C.virConnectGetHostname(c.ptr)
	if str == nil {
		return "", GetLastError()
	}
	hostname := C.GoString(str)
	C.free(unsafe.Pointer(str))
	return hostname, nil
}

func (c *VirConnection) GetLibVersion() (uint32, error) {
	var version C.ulong
	if err := C.virConnectGetLibVersion(c.ptr, &version); err < 0 {
		return 0, GetLastError()
	}
	return uint32(version), nil
}

func (c *VirConnection) GetType() (string, error) {
	str := C.virConnectGetType(c.ptr)
	if str == nil {
		return "", GetLastError()
	}
	hypDriver := C.GoString(str)
	return hypDriver, nil
}

func (c *VirConnection) IsAlive() (bool, error) {
	result := C.virConnectIsAlive(c.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (c *VirConnection) IsEncrypted() (bool, error) {
	result := C.virConnectIsEncrypted(c.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (c *VirConnection) IsSecure() (bool, error) {
	result := C.virConnectIsSecure(c.ptr)
	if result == -1 {
		return false, GetLastError()
	}
	if result == 1 {
		return true, nil
	}
	return false, nil
}

func (c *VirConnection) ListDefinedDomains() ([]string, error) {
	var names [1024](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numDomains := C.virConnectListDefinedDomains(
		c.ptr,
		(**C.char)(namesPtr),
		1024)
	if numDomains == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numDomains)
	for k := 0; k < int(numDomains); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListDomains() ([]uint32, error) {
	var cDomainsIds [512](uint32)
	cDomainsPointer := unsafe.Pointer(&cDomainsIds)
	numDomains := C.virConnectListDomains(c.ptr, (*C.int)(cDomainsPointer), 512)
	if numDomains == -1 {
		return nil, GetLastError()
	}

	return cDomainsIds[:numDomains], nil
}

func (c *VirConnection) ListInterfaces() ([]string, error) {
	const maxIfaces = 1024
	var names [maxIfaces](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numIfaces := C.virConnectListInterfaces(
		c.ptr,
		(**C.char)(namesPtr),
		maxIfaces)
	if numIfaces == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numIfaces)
	for k := 0; k < int(numIfaces); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListNetworks() ([]string, error) {
	const maxNets = 1024
	var names [maxNets](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numNetworks := C.virConnectListNetworks(
		c.ptr,
		(**C.char)(namesPtr),
		maxNets)
	if numNetworks == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numNetworks)
	for k := 0; k < int(numNetworks); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListStoragePools() ([]string, error) {
	const maxPools = 1024
	var names [maxPools](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numStoragePools := C.virConnectListStoragePools(
		c.ptr,
		(**C.char)(namesPtr),
		maxPools)
	if numStoragePools == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numStoragePools)
	for k := 0; k < int(numStoragePools); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) LookupDomainById(id uint32) (VirDomain, error) {
	ptr := C.virDomainLookupByID(c.ptr, C.int(id))
	if ptr == nil {
		return VirDomain{}, GetLastError()
	}
	return VirDomain{ptr: ptr}, nil
}

func (c *VirConnection) LookupDomainByName(id string) (VirDomain, error) {
	cName := C.CString(id)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virDomainLookupByName(c.ptr, cName)
	if ptr == nil {
		return VirDomain{}, GetLastError()
	}
	return VirDomain{ptr: ptr}, nil
}

func (c *VirConnection) LookupByUUIDString(uuid string) (VirDomain, error) {
	cUuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virDomainLookupByUUIDString(c.ptr, cUuid)
	if ptr == nil {
		return VirDomain{}, GetLastError()
	}
	return VirDomain{ptr: ptr}, nil
}

func (c *VirConnection) DomainCreateXMLFromFile(xmlFile string, flags VirDomainCreateFlags) (VirDomain, error) {
	xmlConfig, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		return VirDomain{}, err
	}
	return c.DomainCreateXML(string(xmlConfig), flags)
}

func (c *VirConnection) DomainCreateXML(xmlConfig string, flags VirDomainCreateFlags) (VirDomain, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virDomainCreateXML(c.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return VirDomain{}, GetLastError()
	}
	return VirDomain{ptr: ptr}, nil
}

func (c *VirConnection) DomainDefineXMLFromFile(xmlFile string) (VirDomain, error) {
	xmlConfig, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		return VirDomain{}, err
	}
	return c.DomainDefineXML(string(xmlConfig))
}

func (c *VirConnection) DomainDefineXML(xmlConfig string) (VirDomain, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virDomainDefineXML(c.ptr, cXml)
	if ptr == nil {
		return VirDomain{}, GetLastError()
	}
	return VirDomain{ptr: ptr}, nil
}

func (c *VirConnection) ListDefinedInterfaces() ([]string, error) {
	const maxIfaces = 1024
	var names [maxIfaces](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numIfaces := C.virConnectListDefinedInterfaces(
		c.ptr,
		(**C.char)(namesPtr),
		maxIfaces)
	if numIfaces == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numIfaces)
	for k := 0; k < int(numIfaces); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListDefinedNetworks() ([]string, error) {
	const maxNets = 1024
	var names [maxNets](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numNetworks := C.virConnectListDefinedNetworks(
		c.ptr,
		(**C.char)(namesPtr),
		maxNets)
	if numNetworks == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numNetworks)
	for k := 0; k < int(numNetworks); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) ListDefinedStoragePools() ([]string, error) {
	const maxPools = 1024
	var names [maxPools](*C.char)
	namesPtr := unsafe.Pointer(&names)
	numStoragePools := C.virConnectListDefinedStoragePools(
		c.ptr,
		(**C.char)(namesPtr),
		maxPools)
	if numStoragePools == -1 {
		return nil, GetLastError()
	}
	goNames := make([]string, numStoragePools)
	for k := 0; k < int(numStoragePools); k++ {
		goNames[k] = C.GoString(names[k])
		C.free(unsafe.Pointer(names[k]))
	}
	return goNames, nil
}

func (c *VirConnection) NumOfDefinedInterfaces() (int, error) {
	result := int(C.virConnectNumOfDefinedInterfaces(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfDefinedNetworks() (int, error) {
	result := int(C.virConnectNumOfDefinedNetworks(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfDefinedStoragePools() (int, error) {
	result := int(C.virConnectNumOfDefinedStoragePools(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfDomains() (int, error) {
	result := int(C.virConnectNumOfDomains(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfInterfaces() (int, error) {
	result := int(C.virConnectNumOfInterfaces(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfNetworks() (int, error) {
	result := int(C.virConnectNumOfNetworks(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfNWFilters() (int, error) {
	result := int(C.virConnectNumOfNWFilters(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NumOfSecrets() (int, error) {
	result := int(C.virConnectNumOfSecrets(c.ptr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) NetworkDefineXMLFromFile(xmlFile string) (VirNetwork, error) {
	xmlConfig, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		return VirNetwork{}, err
	}
	return c.NetworkDefineXML(string(xmlConfig))
}

func (c *VirConnection) NetworkDefineXML(xmlConfig string) (VirNetwork, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virNetworkDefineXML(c.ptr, cXml)
	if ptr == nil {
		return VirNetwork{}, GetLastError()
	}
	return VirNetwork{ptr: ptr}, nil
}

func (c *VirConnection) NetworkCreateXML(xmlConfig string) (VirNetwork, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virNetworkCreateXML(c.ptr, cXml)
	if ptr == nil {
		return VirNetwork{}, GetLastError()
	}
	return VirNetwork{ptr: ptr}, nil
}

func (c *VirConnection) LookupNetworkByName(name string) (VirNetwork, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virNetworkLookupByName(c.ptr, cName)
	if ptr == nil {
		return VirNetwork{}, GetLastError()
	}
	return VirNetwork{ptr: ptr}, nil
}

func (c *VirConnection) LookupNetworkByUUIDString(uuid string) (VirNetwork, error) {
	cUuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virNetworkLookupByUUIDString(c.ptr, cUuid)
	if ptr == nil {
		return VirNetwork{}, GetLastError()
	}
	return VirNetwork{ptr: ptr}, nil
}

func (c *VirConnection) SetKeepAlive(interval int, count uint) error {
	res := int(C.virConnectSetKeepAlive(c.ptr, C.int(interval), C.uint(count)))
	switch res {
	case 0:
		return nil
	default:
		return GetLastError()
	}
}

func (c *VirConnection) GetSysinfo(flags uint) (string, error) {
	cStr := C.virConnectGetSysinfo(c.ptr, C.uint(flags))
	if cStr == nil {
		return "", GetLastError()
	}
	info := C.GoString(cStr)
	C.free(unsafe.Pointer(cStr))
	return info, nil
}

func (c *VirConnection) GetURI() (string, error) {
	cStr := C.virConnectGetURI(c.ptr)
	if cStr == nil {
		return "", GetLastError()
	}
	uri := C.GoString(cStr)
	C.free(unsafe.Pointer(cStr))
	return uri, nil
}

func (c *VirConnection) GetMaxVcpus(typeAttr string) (int, error) {
	var cTypeAttr *C.char
	if typeAttr != "" {
		cTypeAttr = C.CString(typeAttr)
		defer C.free(unsafe.Pointer(cTypeAttr))
	}
	result := int(C.virConnectGetMaxVcpus(c.ptr, cTypeAttr))
	if result == -1 {
		return 0, GetLastError()
	}
	return result, nil
}

func (c *VirConnection) InterfaceDefineXMLFromFile(xmlFile string) (VirInterface, error) {
	xmlConfig, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		return VirInterface{}, err
	}
	return c.InterfaceDefineXML(string(xmlConfig), 0)
}

func (c *VirConnection) InterfaceDefineXML(xmlConfig string, flags uint32) (VirInterface, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virInterfaceDefineXML(c.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return VirInterface{}, GetLastError()
	}
	return VirInterface{ptr: ptr}, nil
}

func (c *VirConnection) LookupInterfaceByName(name string) (VirInterface, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virInterfaceLookupByName(c.ptr, cName)
	if ptr == nil {
		return VirInterface{}, GetLastError()
	}
	return VirInterface{ptr: ptr}, nil
}

func (c *VirConnection) LookupInterfaceByMACString(mac string) (VirInterface, error) {
	cName := C.CString(mac)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virInterfaceLookupByMACString(c.ptr, cName)
	if ptr == nil {
		return VirInterface{}, GetLastError()
	}
	return VirInterface{ptr: ptr}, nil
}

func (c *VirConnection) StoragePoolDefineXMLFromFile(xmlFile string) (VirStoragePool, error) {
	xmlConfig, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		return VirStoragePool{}, err
	}
	return c.StoragePoolDefineXML(string(xmlConfig), 0)
}

func (c *VirConnection) StoragePoolDefineXML(xmlConfig string, flags uint32) (VirStoragePool, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virStoragePoolDefineXML(c.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return VirStoragePool{}, GetLastError()
	}
	return VirStoragePool{ptr: ptr}, nil
}

func (c *VirConnection) LookupStoragePoolByName(name string) (VirStoragePool, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virStoragePoolLookupByName(c.ptr, cName)
	if ptr == nil {
		return VirStoragePool{}, GetLastError()
	}
	return VirStoragePool{ptr: ptr}, nil
}

func (c *VirConnection) LookupStoragePoolByUUIDString(uuid string) (VirStoragePool, error) {
	cUuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virStoragePoolLookupByUUIDString(c.ptr, cUuid)
	if ptr == nil {
		return VirStoragePool{}, GetLastError()
	}
	return VirStoragePool{ptr: ptr}, nil
}

func (c *VirConnection) NWFilterDefineXMLFromFile(xmlFile string) (VirNWFilter, error) {
	xmlConfig, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		return VirNWFilter{}, err
	}
	return c.NWFilterDefineXML(string(xmlConfig))
}

func (c *VirConnection) NWFilterDefineXML(xmlConfig string) (VirNWFilter, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virNWFilterDefineXML(c.ptr, cXml)
	if ptr == nil {
		return VirNWFilter{}, GetLastError()
	}
	return VirNWFilter{ptr: ptr}, nil
}

func (c *VirConnection) LookupNWFilterByName(name string) (VirNWFilter, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	ptr := C.virNWFilterLookupByName(c.ptr, cName)
	if ptr == nil {
		return VirNWFilter{}, GetLastError()
	}
	return VirNWFilter{ptr: ptr}, nil
}

func (c *VirConnection) LookupNWFilterByUUIDString(uuid string) (VirNWFilter, error) {
	cUuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virNWFilterLookupByUUIDString(c.ptr, cUuid)
	if ptr == nil {
		return VirNWFilter{}, GetLastError()
	}
	return VirNWFilter{ptr: ptr}, nil
}

func (c *VirConnection) LookupNWFilterByUUID(uuid []byte) (VirNWFilter, error) {
	if len(uuid) != C.VIR_UUID_BUFLEN {
		return VirNWFilter{}, fmt.Errorf("UUID must be exactly %d bytes in size",
			int(C.VIR_UUID_BUFLEN))
	}
	cUuid := C.CBytes(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virNWFilterLookupByUUID(c.ptr, (*C.uchar)(cUuid))
	if ptr == nil {
		return VirNWFilter{}, GetLastError()
	}
	return VirNWFilter{ptr: ptr}, nil
}

func (c *VirConnection) LookupStorageVolByKey(key string) (VirStorageVol, error) {
	cKey := C.CString(key)
	defer C.free(unsafe.Pointer(cKey))
	ptr := C.virStorageVolLookupByKey(c.ptr, cKey)
	if ptr == nil {
		return VirStorageVol{}, GetLastError()
	}
	return VirStorageVol{ptr: ptr}, nil
}

func (c *VirConnection) LookupStorageVolByPath(path string) (VirStorageVol, error) {
	cPath := C.CString(path)
	defer C.free(unsafe.Pointer(cPath))
	ptr := C.virStorageVolLookupByPath(c.ptr, cPath)
	if ptr == nil {
		return VirStorageVol{}, GetLastError()
	}
	return VirStorageVol{ptr: ptr}, nil
}

func (c *VirConnection) SecretDefineXMLFromFile(xmlFile string) (VirSecret, error) {
	xmlConfig, err := ioutil.ReadFile(xmlFile)
	if err != nil {
		return VirSecret{}, err
	}
	return c.SecretDefineXML(string(xmlConfig), 0)
}

func (c *VirConnection) SecretDefineXML(xmlConfig string, flags uint32) (VirSecret, error) {
	cXml := C.CString(string(xmlConfig))
	defer C.free(unsafe.Pointer(cXml))
	ptr := C.virSecretDefineXML(c.ptr, cXml, C.uint(flags))
	if ptr == nil {
		return VirSecret{}, GetLastError()
	}
	return VirSecret{ptr: ptr}, nil
}

func (c *VirConnection) LookupSecretByUUID(uuid []byte) (VirSecret, error) {
	if len(uuid) != C.VIR_UUID_BUFLEN {
		return VirSecret{}, fmt.Errorf("UUID must be exactly %d bytes in size",
			int(C.VIR_UUID_BUFLEN))
	}
	cUuid := C.CBytes(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virSecretLookupByUUID(c.ptr, (*C.uchar)(cUuid))
	if ptr == nil {
		return VirSecret{}, GetLastError()
	}
	return VirSecret{ptr: ptr}, nil
}

func (c *VirConnection) LookupSecretByUUIDString(uuid string) (VirSecret, error) {
	cUuid := C.CString(uuid)
	defer C.free(unsafe.Pointer(cUuid))
	ptr := C.virSecretLookupByUUIDString(c.ptr, cUuid)
	if ptr == nil {
		return VirSecret{}, GetLastError()
	}
	return VirSecret{ptr: ptr}, nil
}

func (c *VirConnection) LookupSecretByUsage(usageType VirSecretUsageType, usageID string) (VirSecret, error) {
	cUsageID := C.CString(usageID)
	defer C.free(unsafe.Pointer(cUsageID))
	ptr := C.virSecretLookupByUsage(c.ptr, C.int(usageType), cUsageID)
	if ptr == nil {
		return VirSecret{}, GetLastError()
	}
	return VirSecret{ptr: ptr}, nil
}

func (c *VirConnection) ListAllInterfaces(flags uint32) ([]VirInterface, error) {
	var cList *C.virInterfacePtr
	numIfaces := C.virConnectListAllInterfaces(c.ptr, (**C.virInterfacePtr)(&cList), C.uint(flags))
	if numIfaces == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numIfaces),
		Cap:  int(numIfaces),
	}
	var ifaces []VirInterface
	slice := *(*[]C.virInterfacePtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		ifaces = append(ifaces, VirInterface{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return ifaces, nil
}

func (c *VirConnection) ListAllNetworks(flags VirConnectListAllNetworksFlags) ([]VirNetwork, error) {
	var cList *C.virNetworkPtr
	numNets := C.virConnectListAllNetworks(c.ptr, (**C.virNetworkPtr)(&cList), C.uint(flags))
	if numNets == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numNets),
		Cap:  int(numNets),
	}
	var nets []VirNetwork
	slice := *(*[]C.virNetworkPtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		nets = append(nets, VirNetwork{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return nets, nil
}

func (c *VirConnection) ListAllDomains(flags VirConnectListAllDomainsFlags) ([]VirDomain, error) {
	var cList *C.virDomainPtr
	numDomains := C.virConnectListAllDomains(c.ptr, (**C.virDomainPtr)(&cList), C.uint(flags))
	if numDomains == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numDomains),
		Cap:  int(numDomains),
	}
	var domains []VirDomain
	slice := *(*[]C.virDomainPtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		domains = append(domains, VirDomain{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return domains, nil
}

func (c *VirConnection) ListAllNWFilters(flags uint32) ([]VirNWFilter, error) {
	var cList *C.virNWFilterPtr
	numNWFilters := C.virConnectListAllNWFilters(c.ptr, (**C.virNWFilterPtr)(&cList), C.uint(flags))
	if numNWFilters == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numNWFilters),
		Cap:  int(numNWFilters),
	}
	var filters []VirNWFilter
	slice := *(*[]C.virNWFilterPtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		filters = append(filters, VirNWFilter{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return filters, nil
}

func (c *VirConnection) ListAllStoragePools(flags VirConnectListAllStoragePoolsFlags) ([]VirStoragePool, error) {
	var cList *C.virStoragePoolPtr
	numPools := C.virConnectListAllStoragePools(c.ptr, (**C.virStoragePoolPtr)(&cList), C.uint(flags))
	if numPools == -1 {
		return nil, GetLastError()
	}
	hdr := reflect.SliceHeader{
		Data: uintptr(unsafe.Pointer(cList)),
		Len:  int(numPools),
		Cap:  int(numPools),
	}
	var pools []VirStoragePool
	slice := *(*[]C.virStoragePoolPtr)(unsafe.Pointer(&hdr))
	for _, ptr := range slice {
		pools = append(pools, VirStoragePool{ptr})
	}
	C.free(unsafe.Pointer(cList))
	return pools, nil
}

func (c *VirConnection) InterfaceChangeBegin(flags uint) error {
	ret := C.virInterfaceChangeBegin(c.ptr, C.uint(flags))
	if ret == -1 {
		return GetLastError()
	}
	return nil
}

func (c *VirConnection) InterfaceChangeCommit(flags uint) error {
	ret := C.virInterfaceChangeCommit(c.ptr, C.uint(flags))
	if ret == -1 {
		return GetLastError()
	}
	return nil
}

func (c *VirConnection) InterfaceChangeRollback(flags uint) error {
	ret := C.virInterfaceChangeRollback(c.ptr, C.uint(flags))
	if ret == -1 {
		return GetLastError()
	}
	return nil
}
