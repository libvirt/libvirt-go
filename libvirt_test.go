package libvirt

import (
	"testing"
)

func buildTestConnection() VirConnection {
	conn, _ := NewVirConnection("test:///default")
	return conn
}

func TestConnection(t *testing.T) {
	conn, err := NewVirConnection("test:///default")
	if err != nil {
		t.Error(err)
	}
	err = conn.CloseConnection()
	if err != nil {
		t.Error(err)
	}
}

func TestInvalidConnection(t *testing.T) {
	_, err := NewVirConnection("invalid_transport:///default")
	if err == nil {
		t.Error("Non-existent transport works")
	}
}

func TestCapabilities(t *testing.T) {
	conn := buildTestConnection()
	capabilities, err := conn.GetCapabilities()
	if err != nil {
		t.Error(err)
	}
	if capabilities == "" {
		t.Error("Capabilities was empty")
	}
}

func TestGetNodeInfo(t *testing.T) {
	conn := buildTestConnection()
	ni, err := conn.GetNodeInfo()
	if err != nil {
		t.Error(err)
	}
	if ni.GetModel() != "i686" {
		t.Error("Expected i686 model in test transport")
	}
}

func TestHostname(t *testing.T) {
	conn := buildTestConnection()
	hostname, err := conn.GetHostname()
	if err != nil {
		t.Error(err)
	}
	if hostname == "" {
		t.Error("Hostname was empty")
	}
}

func TestListDefinedDomains(t *testing.T) {
	conn := buildTestConnection()
	doms, err := conn.ListDefinedDomains()
	if err != nil {
		t.Error(err)
	}
	if len(doms) != 0 {
		t.Error("Defined domains should be zero in test transport")
	}
}

func TestListDomains(t *testing.T) {
	conn := buildTestConnection()
	doms, err := conn.ListDomains()
	if err != nil {
		t.Error(err)
	}
	if len(doms) != 1 {
		t.Error("Length of active domains should be 1 in test transport")
	}
	if doms[0] != 1 {
		t.Error("Active domain should have ID #1 in test transport")
	}
}

func TestLookupDomainById(t *testing.T) {
	conn := buildTestConnection()
	_, err := conn.LookupDomainById(1)
	if err != nil {
		t.Error(err)
	}
}

func TestLookupInvalidDomainById(t *testing.T) {
	conn := buildTestConnection()
	_, err := conn.LookupDomainById(2)
	if err == nil {
		t.Error("Domain #2 shouldn't exist in test transport")
	}
}

func TestLookupDomainByName(t *testing.T) {
	conn := buildTestConnection()
	_, err := conn.LookupDomainByName("test")
	if err != nil {
		t.Error(err)
	}
}

func TestLookupInvalidDomainByName(t *testing.T) {
	conn := buildTestConnection()
	_, err := conn.LookupDomainByName("non_existent_domain")
	if err == nil {
		t.Error("Could find non-existent domain by name")
	}
}
