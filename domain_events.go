package libvirt

import (
	"fmt"
	"unsafe"
)

/*
#cgo pkg-config: libvirt
#include <libvirt/libvirt.h>

void domainEventLifecycleCallback_cgo(virConnectPtr c, virDomainPtr d,
                                     int event, int detail, void* data);

void domainEventGenericCallback_cgo(virConnectPtr c, virDomainPtr d, void* data);

void domainEventRTCChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                     long long utcoffset, void* data);

void domainEventWatchdogCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    int action, void* data);

void domainEventIOErrorCallback_cgo(virConnectPtr c, virDomainPtr d,
                                   const char *srcPath, const char *devAlias,
                                   int action, void* data);

void domainEventGraphicsCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    int phase, const virDomainEventGraphicsAddress *local,
                                    const virDomainEventGraphicsAddress *remote,
                                    const char *authScheme,
                                    const virDomainEventGraphicsSubject *subject, void* data);

void domainEventIOErrorReasonCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         const char *srcPath, const char *devAlias,
                                         int action, const char *reason, void* data);

void domainEventBlockJobCallback_cgo(virConnectPtr c, virDomainPtr d,
                                    const char *disk, int type, int status, void* data);

void domainEventDiskChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                      const char *oldSrcPath, const char *newSrcPath,
                                      const char *devAlias, int reason, void* data);

void domainEventTrayChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                      const char *devAlias, int reason, void* data);

void domainEventReasonCallback_cgo(virConnectPtr c, virDomainPtr d,
                                  int reason, void* data);

void domainEventBalloonChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         unsigned long long actual, void* data);

void domainEventDeviceRemovedCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         const char *devAlias, void* data);

int virConnectDomainEventRegisterAny_cgo(virConnectPtr c,  virDomainPtr d,
                                         int eventID, virConnectDomainEventGenericCallback cb,
                                         long goCallbackId);
*/
import "C"

type DomainLifecycleEvent struct {
	Event DomainEventType
	// TODO: we can make Detail typesafe somehow ?
	Detail int
}

type DomainRTCChangeEvent struct {
	Utcoffset int64
}

type DomainWatchdogEvent struct {
	Action DomainEventWatchdogAction
}

type DomainIOErrorEvent struct {
	SrcPath  string
	DevAlias string
	Action   DomainEventIOErrorAction
}

type DomainEventGraphicsAddress struct {
	Family  DomainEventGraphicsAddressType
	Node    string
	Service string
}

type DomainEventGraphicsSubjectIdentity struct {
	Type string
	Name string
}

type DomainGraphicsEvent struct {
	Phase      DomainEventGraphicsPhase
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
	Type   DomainBlockJobType
	Status ConnectDomainEventBlockJobStatus
}

type DomainDiskChangeEvent struct {
	OldSrcPath string
	NewSrcPath string
	DevAlias   string
	Reason     ConnectDomainEventDiskChangeReason
}

type DomainTrayChangeEvent struct {
	DevAlias string
	Reason   ConnectDomainEventTrayChangeReason
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
	opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

	eventDetails := DomainLifecycleEvent{
		Event:  DomainEventType(event),
		Detail: detail,
	}

	callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventGenericCallback
func domainEventGenericCallback(c C.virConnectPtr, d C.virDomainPtr,
	opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

	callDomainCallbackId(opaque, &connection, &domain, nil)
}

//export domainEventRTCChangeCallback
func domainEventRTCChangeCallback(c C.virConnectPtr, d C.virDomainPtr,
	utcoffset int64, opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

	eventDetails := DomainRTCChangeEvent{
		Utcoffset: utcoffset,
	}

	callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventWatchdogCallback
func domainEventWatchdogCallback(c C.virConnectPtr, d C.virDomainPtr,
	action int, opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

	eventDetails := DomainWatchdogEvent{
		Action: DomainEventWatchdogAction(action),
	}

	callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventIOErrorCallback
func domainEventIOErrorCallback(c C.virConnectPtr, d C.virDomainPtr,
	srcPath *C.char, devAlias *C.char, action int, opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

	eventDetails := DomainIOErrorEvent{
		SrcPath:  C.GoString(srcPath),
		DevAlias: C.GoString(devAlias),
		Action:   DomainEventIOErrorAction(action),
	}

	callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventGraphicsCallback
func domainEventGraphicsCallback(c C.virConnectPtr, d C.virDomainPtr,
	phase int,
	local C.virDomainEventGraphicsAddressPtr,
	remote C.virDomainEventGraphicsAddressPtr,
	authScheme *C.char,
	subject C.virDomainEventGraphicsSubjectPtr,
	opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

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
		Phase: DomainEventGraphicsPhase(phase),
		Local: DomainEventGraphicsAddress{
			Family:  DomainEventGraphicsAddressType(local.family),
			Node:    C.GoString(local.node),
			Service: C.GoString(local.service),
		},
		Remote: DomainEventGraphicsAddress{
			Family:  DomainEventGraphicsAddressType(remote.family),
			Node:    C.GoString(remote.node),
			Service: C.GoString(remote.service),
		},
		AuthScheme: C.GoString(authScheme),
		Subject:    subjectGo,
	}

	callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventIOErrorReasonCallback
func domainEventIOErrorReasonCallback(c C.virConnectPtr, d C.virDomainPtr,
	srcPath *C.char, devAlias *C.char, action int, reason *C.char,
	opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

	eventDetails := DomainIOErrorReasonEvent{
		DomainIOErrorEvent: DomainIOErrorEvent{
			SrcPath:  C.GoString(srcPath),
			DevAlias: C.GoString(devAlias),
			Action:   DomainEventIOErrorAction(action),
		},
		Reason: C.GoString(reason),
	}

	callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventBlockJobCallback
func domainEventBlockJobCallback(c C.virConnectPtr, d C.virDomainPtr,
	disk *C.char, _type int, status int, opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

	eventDetails := DomainBlockJobEvent{
		Disk:   C.GoString(disk),
		Type:   DomainBlockJobType(_type),
		Status: ConnectDomainEventBlockJobStatus(status),
	}

	callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventDiskChangeCallback
func domainEventDiskChangeCallback(c C.virConnectPtr, d C.virDomainPtr,
	oldSrcPath *C.char, newSrcPath *C.char, devAlias *C.char,
	reason int, opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

	eventDetails := DomainDiskChangeEvent{
		OldSrcPath: C.GoString(oldSrcPath),
		NewSrcPath: C.GoString(newSrcPath),
		DevAlias:   C.GoString(devAlias),
		Reason:     ConnectDomainEventDiskChangeReason(reason),
	}

	callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventTrayChangeCallback
func domainEventTrayChangeCallback(c C.virConnectPtr, d C.virDomainPtr,
	devAlias *C.char, reason int, opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

	eventDetails := DomainTrayChangeEvent{
		DevAlias: C.GoString(devAlias),
		Reason:   ConnectDomainEventTrayChangeReason(reason),
	}

	callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventReasonCallback
func domainEventReasonCallback(c C.virConnectPtr, d C.virDomainPtr,
	reason int, opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

	eventDetails := DomainReasonEvent{
		Reason: reason,
	}

	callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventBalloonChangeCallback
func domainEventBalloonChangeCallback(c C.virConnectPtr, d C.virDomainPtr,
	actual uint64, opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

	eventDetails := DomainBalloonChangeEvent{
		Actual: actual,
	}

	callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

//export domainEventDeviceRemovedCallback
func domainEventDeviceRemovedCallback(c C.virConnectPtr, d C.virDomainPtr,
	devAlias *C.char, opaque int) {

	domain := Domain{ptr: d}
	connection := Connect{ptr: c}

	eventDetails := DomainDeviceRemovedEvent{
		DevAlias: C.GoString(devAlias),
	}
	callDomainCallbackId(opaque, &connection, &domain, eventDetails)
}

// BUG(vincentbernat): The returned value of DomainEventCallback is
// ignored and should be removed from the signature.

// DomainEventCallback is the signature of functions that can be
// registered as a domain event callback. The event parameter should
// be casted to the more specific event structure
// (eg. DomainLifecycleEvent). The return code is ignored.
type DomainEventCallback func(c *Connect, d *Domain,
	event interface{}, f func()) int

type domainCallbackContext struct {
	cb *DomainEventCallback
	f  func()
}

func callDomainCallbackId(goCallbackId int, c *Connect, d *Domain,
	event interface{}) {
	ctx := getCallbackId(goCallbackId)
	switch cctx := ctx.(type) {
	case *domainCallbackContext:
		(*cctx.cb)(c, d, event, cctx.f)
	default:
		panic("Inappropriate callback type called")
	}
}

// BUG(vincentbernat): The returned value of DomainEventRegister, should be an
// error instead of an int, for uniformity with other functions.

func (c *Connect) DomainEventRegister(dom Domain,
	eventId DomainEventID,
	callback *DomainEventCallback,
	opaque func()) int {
	var callbackPtr unsafe.Pointer
	context := &domainCallbackContext{
		cb: callback,
		f:  opaque,
	}
	goCallBackId := registerCallbackId(context)

	switch eventId {
	case DOMAIN_EVENT_ID_LIFECYCLE:
		callbackPtr = unsafe.Pointer(C.domainEventLifecycleCallback_cgo)
	case DOMAIN_EVENT_ID_REBOOT, DOMAIN_EVENT_ID_CONTROL_ERROR:
		callbackPtr = unsafe.Pointer(C.domainEventGenericCallback_cgo)
	case DOMAIN_EVENT_ID_RTC_CHANGE:
		callbackPtr = unsafe.Pointer(C.domainEventRTCChangeCallback_cgo)
	case DOMAIN_EVENT_ID_WATCHDOG:
		callbackPtr = unsafe.Pointer(C.domainEventWatchdogCallback_cgo)
	case DOMAIN_EVENT_ID_IO_ERROR:
		callbackPtr = unsafe.Pointer(C.domainEventIOErrorCallback_cgo)
	case DOMAIN_EVENT_ID_GRAPHICS:
		callbackPtr = unsafe.Pointer(C.domainEventGraphicsCallback_cgo)
	case DOMAIN_EVENT_ID_IO_ERROR_REASON:
		callbackPtr = unsafe.Pointer(C.domainEventIOErrorReasonCallback_cgo)
	case DOMAIN_EVENT_ID_BLOCK_JOB:
		// TODO Post 1.2.4, uncomment later
		// case DOMAIN_EVENT_ID_BLOCK_JOB_2:
		callbackPtr = unsafe.Pointer(C.domainEventBlockJobCallback_cgo)
	case DOMAIN_EVENT_ID_DISK_CHANGE:
		callbackPtr = unsafe.Pointer(C.domainEventDiskChangeCallback_cgo)
	case DOMAIN_EVENT_ID_TRAY_CHANGE:
		callbackPtr = unsafe.Pointer(C.domainEventTrayChangeCallback_cgo)
	case DOMAIN_EVENT_ID_PMWAKEUP, DOMAIN_EVENT_ID_PMSUSPEND, DOMAIN_EVENT_ID_PMSUSPEND_DISK:
		callbackPtr = unsafe.Pointer(C.domainEventReasonCallback_cgo)
	case DOMAIN_EVENT_ID_BALLOON_CHANGE:
		callbackPtr = unsafe.Pointer(C.domainEventBalloonChangeCallback_cgo)
	case DOMAIN_EVENT_ID_DEVICE_REMOVED:
		callbackPtr = unsafe.Pointer(C.domainEventDeviceRemovedCallback_cgo)
	default:
	}
	ret := C.virConnectDomainEventRegisterAny_cgo(c.ptr, dom.ptr, C.int(eventId),
		C.virConnectDomainEventGenericCallback(callbackPtr),
		C.long(goCallBackId))
	if ret == -1 {
		freeCallbackId(goCallBackId)
		return -1
	}
	return int(ret)
}

func (c *Connect) DomainEventDeregister(callbackId int) error {
	// Deregister the callback
	if i := int(C.virConnectDomainEventDeregisterAny(c.ptr, C.int(callbackId))); i != 0 {
		return GetLastError()
	}
	return nil
}

func (e DomainLifecycleEvent) String() string {
	var detail, event string
	switch e.Event {
	case DOMAIN_EVENT_DEFINED:
		event = "defined"
		switch DomainEventDefinedDetailType(e.Detail) {
		case DOMAIN_EVENT_DEFINED_ADDED:
			detail = "added"
		case DOMAIN_EVENT_DEFINED_UPDATED:
			detail = "updated"
		default:
			detail = "unknown"
		}

	case DOMAIN_EVENT_UNDEFINED:
		event = "undefined"
		switch DomainEventUndefinedDetailType(e.Detail) {
		case DOMAIN_EVENT_UNDEFINED_REMOVED:
			detail = "removed"
		default:
			detail = "unknown"
		}

	case DOMAIN_EVENT_STARTED:
		event = "started"
		switch DomainEventStartedDetailType(e.Detail) {
		case DOMAIN_EVENT_STARTED_BOOTED:
			detail = "booted"
		case DOMAIN_EVENT_STARTED_MIGRATED:
			detail = "migrated"
		case DOMAIN_EVENT_STARTED_RESTORED:
			detail = "restored"
		case DOMAIN_EVENT_STARTED_FROM_SNAPSHOT:
			detail = "snapshot"
		default:
			detail = "unknown"
		}

	case DOMAIN_EVENT_SUSPENDED:
		event = "suspended"
		switch DomainEventSuspendedDetailType(e.Detail) {
		case DOMAIN_EVENT_SUSPENDED_PAUSED:
			detail = "paused"
		case DOMAIN_EVENT_SUSPENDED_MIGRATED:
			detail = "migrated"
		case DOMAIN_EVENT_SUSPENDED_IOERROR:
			detail = "I/O error"
		case DOMAIN_EVENT_SUSPENDED_WATCHDOG:
			detail = "watchdog"
		case DOMAIN_EVENT_SUSPENDED_RESTORED:
			detail = "restored"
		case DOMAIN_EVENT_SUSPENDED_FROM_SNAPSHOT:
			detail = "snapshot"
		default:
			detail = "unknown"
		}

	case DOMAIN_EVENT_RESUMED:
		event = "resumed"
		switch DomainEventResumedDetailType(e.Detail) {
		case DOMAIN_EVENT_RESUMED_UNPAUSED:
			detail = "unpaused"
		case DOMAIN_EVENT_RESUMED_MIGRATED:
			detail = "migrated"
		case DOMAIN_EVENT_RESUMED_FROM_SNAPSHOT:
			detail = "snapshot"
		default:
			detail = "unknown"
		}

	case DOMAIN_EVENT_STOPPED:
		event = "stopped"
		switch DomainEventStoppedDetailType(e.Detail) {
		case DOMAIN_EVENT_STOPPED_SHUTDOWN:
			detail = "shutdown"
		case DOMAIN_EVENT_STOPPED_DESTROYED:
			detail = "destroyed"
		case DOMAIN_EVENT_STOPPED_CRASHED:
			detail = "crashed"
		case DOMAIN_EVENT_STOPPED_MIGRATED:
			detail = "migrated"
		case DOMAIN_EVENT_STOPPED_SAVED:
			detail = "saved"
		case DOMAIN_EVENT_STOPPED_FAILED:
			detail = "failed"
		case DOMAIN_EVENT_STOPPED_FROM_SNAPSHOT:
			detail = "snapshot"
		default:
			detail = "unknown"
		}

	case DOMAIN_EVENT_SHUTDOWN:
		event = "shutdown"
		switch DomainEventShutdownDetailType(e.Detail) {
		case DOMAIN_EVENT_SHUTDOWN_FINISHED:
			detail = "finished"
		default:
			detail = "unknown"
		}

	default:
		event = "unknown"
	}

	return fmt.Sprintf("Domain event=%q detail=%q", event, detail)
}

func (e DomainRTCChangeEvent) String() string {
	return fmt.Sprintf("RTC change offset=%d", e.Utcoffset)
}

func (e DomainWatchdogEvent) String() string {
	return fmt.Sprintf("Watchdog action=%d", e.Action)
}

func (e DomainIOErrorEvent) String() string {
	return fmt.Sprintf("I/O error path=%q alias=%q action=%d",
		e.SrcPath, e.DevAlias, e.Action)
}

func (e DomainGraphicsEvent) String() string {
	var phase string
	switch e.Phase {
	case DOMAIN_EVENT_GRAPHICS_CONNECT:
		phase = "connected"
	case DOMAIN_EVENT_GRAPHICS_INITIALIZE:
		phase = "initialized"
	case DOMAIN_EVENT_GRAPHICS_DISCONNECT:
		phase = "disconnected"
	default:
		phase = "unknown"
	}

	return fmt.Sprintf("Graphics phase=%q", phase)
}

func (e DomainIOErrorReasonEvent) String() string {
	return fmt.Sprintf("IO error path=%q alias=%q action=%d reason=%q",
		e.SrcPath, e.DevAlias, e.Action, e.Reason)
}

func (e DomainBlockJobEvent) String() string {
	return fmt.Sprintf("Block job disk=%q status=%d type=%d",
		e.Disk, e.Status, e.Type)
}

func (e DomainDiskChangeEvent) String() string {
	return fmt.Sprintf("Disk change old=%q new=%q alias=%q reason=%d",
		e.OldSrcPath, e.NewSrcPath, e.DevAlias, e.Reason)
}

func (e DomainTrayChangeEvent) String() string {
	return fmt.Sprintf("Tray change dev=%q reason=%d",
		e.DevAlias, e.Reason)
}

func (e DomainBalloonChangeEvent) String() string {
	return fmt.Sprintf("Ballon change %d", e.Actual)
}

func (e DomainDeviceRemovedEvent) String() string {
	return fmt.Sprintf("Device %q removed ", e.DevAlias)
}
