package libvirt

import (
	"testing"
)

func buildTestConnection() VirConnection {
	conn, _ := NewVirConnection("test:///default")
	return conn
}

func buildTestDomain() VirDomain {
	conn := buildTestConnection()
	dom, _ := conn.LookupDomainById(1)
	return dom
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
		t.Error(err)
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

func TestGetDomainName(t *testing.T) {
	dom := buildTestDomain()
	name, err := dom.GetName()
	if err != nil {
		t.Error(err)
	}
	if name != "test" {
		t.Error("Name of active domain in test transport should be 'test'")
	}
}

func TestGetDomainState(t *testing.T) {
	dom := buildTestDomain()
	state, err := dom.GetState()
	if err != nil {
		t.Error(err)
	}
	if len(state) != 2 {
		t.Error("Length of domain state should be 2")
	}
	if state[0] != 1 || state[1] != 1 {
		t.Error("Domain state in test transport should be [1 1]")
	}
}

func TestGetDomainUUID(t *testing.T) {
	dom := buildTestDomain()
	_, err := dom.GetUUID()
	// how to test uuid validity?
	if err != nil {
		t.Error(err)
	}
}

func TestGetDomainUUIDString(t *testing.T) {
	dom := buildTestDomain()
	uuid, err := dom.GetUUIDString()
	if err != nil {
		t.Error(uuid)
	}
}
