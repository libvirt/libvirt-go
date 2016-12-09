package libvirt

import (
	"fmt"
	"testing"
	"time"
)

func init() {
	EventRegisterDefaultImpl()
}

func TestStoragePoolEventRegister(t *testing.T) {

	if true {
		return
	}
	callbackId := -1

	conn := buildTestConnection()
	defer func() {
		if callbackId >= 0 {
			if err := conn.StoragePoolEventDeregister(callbackId); err != nil {
				t.Errorf("got `%v` on StoragePoolEventDeregister instead of nil", err)
			}
		}
		if res, _ := conn.CloseConnection(); res != 0 {
			t.Errorf("CloseConnection() == %d, expected 0", res)
		}
	}()

	defName := time.Now().String()

	nbEvents := 0

	callback := func(c *Connect, d *StoragePool, event *StoragePoolEventLifecycle) {
		if event.Event == STORAGE_POOL_EVENT_STARTED {
			domName, _ := d.GetName()
			if defName != domName {
				t.Fatalf("Name was not '%s': %s", defName, domName)
			}
		}
		eventString := fmt.Sprintf("%s", event)
		expected := "StoragePool event=\"started\" detail=\"booted\""
		if eventString != expected {
			t.Errorf("event == %q, expected %q", eventString, expected)
		}
		nbEvents++
	}

	callbackId, err := conn.StoragePoolEventLifecycleRegister(nil, callback)
	if err != nil {
		t.Error(err)
		return
	}

	// Test a minimally valid xml
	xml := `<pool>
		<name>` + defName + `</name>
	</pool>`
	dom, err := conn.StoragePoolCreateXML(xml, 0)
	if err != nil {
		t.Error(err)
		return
	}

	// This is blocking as long as there is no message
	EventRunDefaultImpl()
	if nbEvents == 0 {
		t.Fatal("At least one event was expected")
	}

	defer func() {
		dom.Destroy()
		dom.Free()
	}()

	// Check that the internal context entry was added, and that there only is
	// one.
	goCallbackLock.Lock()
	if len(goCallbacks) != 1 {
		t.Errorf("goCallbacks should hold one entry, got %+v", goCallbacks)
	}
	goCallbackLock.Unlock()

	// Deregister the event
	if err := conn.StoragePoolEventDeregister(callbackId); err != nil {
		t.Fatal("Event deregistration failed with: %v", err)
	}
	callbackId = -1 // Don't deregister twice

	// Check that the internal context entries was removed
	goCallbackLock.Lock()
	if len(goCallbacks) > 0 {
		t.Errorf("goCallbacks entry wasn't removed: %+v", goCallbacks)
	}
	goCallbackLock.Unlock()
}
