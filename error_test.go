package libvirt

import (
	"reflect"
	"testing"
)

func TestGetLastError(t *testing.T) {
	_, err := NewConnect("invalid_transport:///default")
	if err == nil {
		t.Fatalf("Expected an error when creating invalid connection")
	}
	got := GetLastError()
	expected := Error{0, 0, "", 0}
	if !reflect.DeepEqual(got, expected) {
		t.Errorf("Expected error %+v, got %+v", expected, got)
	}
	if got != ErrNoError {
		t.Errorf("Expected error to be ErrNoError")
	}
}
