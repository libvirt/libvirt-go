package libvirt

import (
	"sync"
	"unsafe"
)

/*
#cgo LDFLAGS: -lvirt
#include <libvirt/libvirt.h>

int domainEventLifecycleCallback_cgo(virConnectPtr c, virDomainPtr d,
                                     int event, int detail, void* data);

int domainEventGenericCallback_cgo(virConnectPtr c, virDomainPtr d, void* data);

int domainEventRTCChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                     long long utcoffset, void* data);

int domainEventWatchdogCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    int action, void* data);

int domainEventIOErrorCallback_cgo(virConnectPtr c, virDomainPtr d,
                                   const char *srcPath, const char *devAlias,
                                   int action, void* data);

int domainEventGraphicsCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    int phase, const virDomainEventGraphicsAddress *local,
                                    const virDomainEventGraphicsAddress *remote,
                                    const char *authScheme,
                                    const virDomainEventGraphicsSubject *subject, void* data);

int domainEventIOErrorReasonCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         const char *srcPath, const char *devAlias,
                                         int action, const char *reason, void* data);

int domainEventBlockJobCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    const char *disk, int type, int status, void* data);

int domainEventDiskChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                      const char *oldSrcPath, const char *newSrcPath,
                                      const char *devAlias, int reason, void* data);

int domainEventTrayChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                      const char *devAlias, int reason, void* data);

int domainEventReasonCallback_cgo(virConnectPtr c, virDomainPtr d,
                                  int reason, void* data);

int domainEventBalloonChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         unsigned long long actual, void* data);

int domainEventDeviceRemovedCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         const char *devAlias, void* data);

int virConnectDomainEventRegisterAny_cgo(virConnectPtr c,  virDomainPtr d,
                                         int eventID, virConnectDomainEventGenericCallback cb,
                                         long goCallbackId);
*/
import "C"

type DomainLifecycleEvent struct {
	Event  int
	Detail int
}

type DomainRTCChangeEvent struct {
	Utcoffset int64
}

type DomainWatchdogEvent struct {
	Action int
}

type DomainIOErrorEvent struct {
	SrcPath  string
	DevAlias string
	Action   int
}

type DomainEventGraphicsAddress struct {
	Family  int
	Node    string
	Service string
}

type DomainEventGraphicsSubjectIdentity struct {
	Type string
	Name string
}

type DomainGraphicsEvent struct {
	Phase      int
	Local      DomainEventGraphicsAddress
	Remote     DomainEventGraphicsAddress
	AuthScheme string
	Subject    []DomainEventGraphicsSubjectIdentity
}

type DomainIOErrorReasonEvent struct {
	DomainIOErrorEvent
	Reason string
}

type DomainBlockJobEvent struct {
	Disk   string
	Type   int
	Status int
}

type DomainDiskChangeEvent struct {
	OldSrcPath string
	NewSrcPath string
	DevAlias   string
	Reason     int
}

type DomainTrayChangeEvent struct {
	DevAlias string
	Reason   int
}

type DomainReasonEvent struct {
	Reason int
}

type DomainBalloonChangeEvent struct {
	Actual uint64
}

type DomainDeviceRemovedEvent struct {
	DevAlias string
}

//export domainEventLifecycleCallback
func domainEventLifecycleCallback(c C.virConnectPtr, d C.virDomainPtr,
	event int, detail int,
	opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainLifecycleEvent{
		Event:  event,
		Detail: detail,
	}

	return callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventGenericCallback
func domainEventGenericCallback(c C.virConnectPtr, d C.virDomainPtr,
	opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	return callDomainCallbackId(opaque, &connection, &domain, nil)
}

//export domainEventRTCChangeCallback
func domainEventRTCChangeCallback(c C.virConnectPtr, d C.virDomainPtr,
	utcoffset int64, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainRTCChangeEvent{
		Utcoffset: utcoffset,
	}

	return callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventWatchdogCallback
func domainEventWatchdogCallback(c C.virConnectPtr, d C.virDomainPtr,
	action int, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainWatchdogEvent{
		Action: action,
	}

	return callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventIOErrorCallback
func domainEventIOErrorCallback(c C.virConnectPtr, d C.virDomainPtr,
	srcPath string, devAlias string, action int, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainIOErrorEvent{
		SrcPath:  srcPath,
		DevAlias: devAlias,
		Action:   action,
	}

	return callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventGraphicsCallback
func domainEventGraphicsCallback(c C.virConnectPtr, d C.virDomainPtr,
	phase int,
	local C.virDomainEventGraphicsAddressPtr,
	remote C.virDomainEventGraphicsAddressPtr,
	authScheme string,
	subject C.virDomainEventGraphicsSubjectPtr,
	opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	subjectGo := make([]DomainEventGraphicsSubjectIdentity, subject.nidentity)
	nidentities := int(subject.nidentity)
	identities := (*[1 << 30]C.virDomainEventGraphicsSubjectIdentity)(unsafe.Pointer(&subject.identities))[:nidentities:nidentities]
	for _, identity := range identities {
		subjectGo = append(subjectGo,
			DomainEventGraphicsSubjectIdentity{
				Type: C.GoString(identity._type),
				Name: C.GoString(identity.name),
			},
		)
	}

	eventDetails := DomainGraphicsEvent{
		Phase: phase,
		Local: DomainEventGraphicsAddress{
			Family:  int(local.family),
			Node:    C.GoString(local.node),
			Service: C.GoString(local.service),
		},
		Remote: DomainEventGraphicsAddress{
			Family:  int(remote.family),
			Node:    C.GoString(remote.node),
			Service: C.GoString(remote.service),
		},
		AuthScheme: authScheme,
		Subject:    subjectGo,
	}

	return callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventIOErrorReasonCallback
func domainEventIOErrorReasonCallback(c C.virConnectPtr, d C.virDomainPtr,
	srcPath string, devAlias string, action int, reason string,
	opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainIOErrorReasonEvent{
		DomainIOErrorEvent: DomainIOErrorEvent{
			SrcPath:  srcPath,
			DevAlias: devAlias,
			Action:   action,
		},
		Reason: reason,
	}

	return callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventBlockJobCallback
func domainEventBlockJobCallback(c C.virConnectPtr, d C.virDomainPtr,
	disk string, _type int, status int, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainBlockJobEvent{
		Disk:   disk,
		Type:   _type,
		Status: status,
	}

	return callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventDiskChangeCallback
func domainEventDiskChangeCallback(c C.virConnectPtr, d C.virDomainPtr,
	oldSrcPath string, newSrcPath string, devAlias string,
	reason int, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainDiskChangeEvent{
		OldSrcPath: oldSrcPath,
		NewSrcPath: newSrcPath,
		DevAlias:   devAlias,
		Reason:     reason,
	}

	return callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventTrayChangeCallback
func domainEventTrayChangeCallback(c C.virConnectPtr, d C.virDomainPtr,
	devAlias string, reason int, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainTrayChangeEvent{
		DevAlias: devAlias,
		Reason:   reason,
	}

	return callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventReasonCallback
func domainEventReasonCallback(c C.virConnectPtr, d C.virDomainPtr,
	reason int, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainReasonEvent{
		Reason: reason,
	}

	return callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventBalloonChangeCallback
func domainEventBalloonChangeCallback(c C.virConnectPtr, d C.virDomainPtr,
	actual uint64, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainBalloonChangeEvent{
		Actual: actual,
	}

	return callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventDeviceRemovedCallback
func domainEventDeviceRemovedCallback(c C.virConnectPtr, d C.virDomainPtr,
	devAlias string, opaque int) int {

	domain := VirDomain{ptr: d}
	connection := VirConnection{ptr: c}

	eventDetails := DomainDeviceRemovedEvent{
		DevAlias: devAlias,
	}
	return callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

type DomainEventCallback func(c *VirConnection, d *VirDomain,
	event interface{}, f func()) int

type domainCallbackContext struct {
	cb *DomainEventCallback
	f  func()
}

const firstGoCallbackId int = 100 // help catch some additional errors during test
var goCallbackLock sync.RWMutex
var goCallbacks = make(map[int]interface{})
var nextGoCallbackId int = firstGoCallbackId

//export freeCallbackId
func freeCallbackId(goCallbackId int) {
	goCallbackLock.Lock()
	delete(goCallbacks, goCallbackId)
	goCallbackLock.Unlock()
}

func getCallbackId(goCallbackId int) interface{} {
	goCallbackLock.RLock()
	ctx := goCallbacks[goCallbackId]
	goCallbackLock.RUnlock()
	if ctx == nil {
		// If this happens there must be a bug in libvirt
		panic("Callback arrived after freeCallbackId was called")
	}
	return ctx
}

func callDomainCallbackId(goCallbackId int, c *VirConnection, d *VirDomain,
	event interface{}) int {
	ctx := getCallbackId(goCallbackId)
	switch cctx := ctx.(type) {
	case *domainCallbackContext:
		return (*cctx.cb)(c, d, event, cctx.f)
	default:
		panic("Inappropriate callback type called")
	}
}

func registerCallbackId(ctx interface{}) int {
	goCallbackLock.Lock()
	goCallBackId := nextGoCallbackId
	nextGoCallbackId++
	for goCallbacks[nextGoCallbackId] != nil {
		nextGoCallbackId++
	}
	goCallbacks[goCallBackId] = ctx
	goCallbackLock.Unlock()
	return goCallBackId
}

func (c *VirConnection) DomainEventRegister(dom VirDomain,
	eventId int,
	callback *DomainEventCallback,
	opaque func()) int {
	var callbackPtr unsafe.Pointer
	context := &domainCallbackContext{
		cb: callback,
		f:  opaque,
	}
	goCallBackId := registerCallbackId(context)

	switch eventId {
	case VIR_DOMAIN_EVENT_ID_LIFECYCLE:
		callbackPtr = unsafe.Pointer(C.domainEventLifecycleCallback_cgo)
	case VIR_DOMAIN_EVENT_ID_REBOOT:
	case VIR_DOMAIN_EVENT_ID_CONTROL_ERROR:
		callbackPtr = unsafe.Pointer(C.domainEventGenericCallback_cgo)
	case VIR_DOMAIN_EVENT_ID_RTC_CHANGE:
		callbackPtr = unsafe.Pointer(C.domainEventRTCChangeCallback_cgo)
	case VIR_DOMAIN_EVENT_ID_WATCHDOG:
		callbackPtr = unsafe.Pointer(C.domainEventWatchdogCallback_cgo)
	case VIR_DOMAIN_EVENT_ID_IO_ERROR:
		callbackPtr = unsafe.Pointer(C.domainEventIOErrorCallback_cgo)
	case VIR_DOMAIN_EVENT_ID_GRAPHICS:
		callbackPtr = unsafe.Pointer(C.domainEventGraphicsCallback_cgo)
	case VIR_DOMAIN_EVENT_ID_IO_ERROR_REASON:
		callbackPtr = unsafe.Pointer(C.domainEventIOErrorReasonCallback_cgo)
	case VIR_DOMAIN_EVENT_ID_BLOCK_JOB:
		// TODO Post 1.2.4, uncomment later
		// case VIR_DOMAIN_EVENT_ID_BLOCK_JOB_2:
		callbackPtr = unsafe.Pointer(C.domainEventBlockJobCallback_cgo)
	case VIR_DOMAIN_EVENT_ID_DISK_CHANGE:
		callbackPtr = unsafe.Pointer(C.domainEventDiskChangeCallback_cgo)
	case VIR_DOMAIN_EVENT_ID_TRAY_CHANGE:
		callbackPtr = unsafe.Pointer(C.domainEventTrayChangeCallback_cgo)
	case VIR_DOMAIN_EVENT_ID_PMWAKEUP:
	case VIR_DOMAIN_EVENT_ID_PMSUSPEND:
	case VIR_DOMAIN_EVENT_ID_PMSUSPEND_DISK:
		callbackPtr = unsafe.Pointer(C.domainEventReasonCallback_cgo)
	case VIR_DOMAIN_EVENT_ID_BALLOON_CHANGE:
		callbackPtr = unsafe.Pointer(C.domainEventBalloonChangeCallback_cgo)
	case VIR_DOMAIN_EVENT_ID_DEVICE_REMOVED:
		callbackPtr = unsafe.Pointer(C.domainEventDeviceRemovedCallback_cgo)
	default:
	}
	ret := C.virConnectDomainEventRegisterAny_cgo(c.ptr, dom.ptr, C.VIR_DOMAIN_EVENT_ID_LIFECYCLE,
		C.virConnectDomainEventGenericCallback(callbackPtr),
		C.long(goCallBackId))
	if ret == -1 {
		freeCallbackId(goCallBackId)
		return -1
	}
	return int(ret)
}

func (c *VirConnection) DomainEventDeregister(callbackId int) int {
	// Deregister the callback
	return int(C.virConnectDomainEventDeregisterAny(c.ptr, C.int(callbackId)))
}

func EventRegisterDefaultImpl() int {
	return int(C.virEventRegisterDefaultImpl())
}

func EventRunDefaultImpl() int {
	return int(C.virEventRunDefaultImpl())
}
