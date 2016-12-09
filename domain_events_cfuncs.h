#ifndef DOMAIN_EVENTS_CFUNCS_H__
#define DOMAIN_EVENTS_CFUNCS_H__

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

void domainEventPMSuspendCallback_cgo(virConnectPtr c, virDomainPtr d,
				      int reason, void* data);

void domainEventPMWakeupCallback_cgo(virConnectPtr c, virDomainPtr d,
				     int reason, void* data);

void domainEventPMSuspendDiskCallback_cgo(virConnectPtr c, virDomainPtr d,
					  int reason, void* data);

void domainEventBalloonChangeCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         unsigned long long actual, void* data);

void domainEventDeviceRemovedCallback_cgo(virConnectPtr c, virDomainPtr d,
                                         const char *devAlias, void* data);
void domainEventTunableCallback_cgo(virConnectPtr conn,
				    virDomainPtr dom,
				    virTypedParameterPtr params,
				    int nparams,
				    void *opaque);
void domainEventAgentLifecycleCallback_cgo(virConnectPtr conn,
					   virDomainPtr dom,
					   int state,
					   int reason,
					   void *opaque);
void domainEventDeviceAddedCallback_cgo(virConnectPtr conn,
					virDomainPtr dom,
					const char *devAlias,
					void *opaque);
void domainEventMigrationIterationCallback_cgo(virConnectPtr conn,
					       virDomainPtr dom,
					       int iteration,
					       void *opaque);
void domainEventJobCompletedCallback_cgo(virConnectPtr conn,
					 virDomainPtr dom,
					 virTypedParameterPtr params,
					 int nparams,
					 void *opaque);
void domainEventDeviceRemovalFailedCallback_cgo(virConnectPtr conn,
						virDomainPtr dom,
						const char *devAlias,
						void *opaque);

int virConnectDomainEventRegisterAny_cgo(virConnectPtr c,  virDomainPtr d,
                                         int eventID, virConnectDomainEventGenericCallback cb,
                                         long goCallbackId);


#endif /* DOMAIN_EVENTS_CFUNCS_H__ */
