package libvirt

import "testing"

func TestGlobalErrorCallback(t *testing.T) {
	var nbErrors int
	errors := make([]VirError, 0, 10)
	callback := ErrorCallback(func(err VirError, f func()) {
		errors = append(errors, err)
		f()
	})
	SetErrorFunc(callback, func() {
		nbErrors++
	})
	NewVirConnection("invalid_transport:///default")
	if len(errors) == 0 {
		t.Errorf("No errors were captured")
	}
	if len(errors) != nbErrors {
		t.Errorf("Captured %d errors (%+v) but counted only %d errors",
			len(errors), errors, nbErrors)
	}
	errors = make([]VirError, 0, 10)
	SetErrorFunc(nil, nil)
	NewVirConnection("invalid_transport:///default")
	if len(errors) != 0 {
		t.Errorf("More errors have been captured: %+v", errors)
	}
}
