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
		return
	}
	_, err = conn.CloseConnection()
	if err != nil {
		t.Error(err)
		return
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
	defer conn.CloseConnection()
	capabilities, err := conn.GetCapabilities()
	if err != nil {
		t.Error(err)
		return
	}
	if capabilities == "" {
		t.Error("Capabilities was empty")
		return
	}
}

func TestGetNodeInfo(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	ni, err := conn.GetNodeInfo()
	if err != nil {
		t.Error(err)
		return
	}
	if ni.GetModel() != "i686" {
		t.Error("Expected i686 model in test transport")
		return
	}
}

func TestHostname(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	hostname, err := conn.GetHostname()
	if err != nil {
		t.Error(err)
		return
	}
	if hostname == "" {
		t.Error("Hostname was empty")
		return
	}
}

func TestListDefinedDomains(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	doms, err := conn.ListDefinedDomains()
	if err != nil {
		t.Error(err)
		return
	}
	if len(doms) != 0 {
		t.Error("Defined domains should be zero in test transport")
		return
	}
}

func TestListDomains(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	doms, err := conn.ListDomains()
	if err != nil {
		t.Error(err)
		return
	}
	if len(doms) != 1 {
		t.Error("Length of active domains should be 1 in test transport")
		return
	}
	if doms[0] != 1 {
		t.Error("Active domain should have ID #1 in test transport")
		return
	}
}

func TestLookupDomainById(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	_, err := conn.LookupDomainById(1)
	if err != nil {
		t.Error(err)
		return
	}
}

func TestLookupInvalidDomainById(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	_, err := conn.LookupDomainById(2)
	if err == nil {
		t.Error("Domain #2 shouldn't exist in test transport")
		return
	}
}

func TestLookupDomainByName(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	_, err := conn.LookupDomainByName("test")
	if err != nil {
		t.Error(err)
		return
	}
}

func TestLookupInvalidDomainByName(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	_, err := conn.LookupDomainByName("non_existent_domain")
	if err == nil {
		t.Error("Could find non-existent domain by name")
		return
	}
}

func TestDomainDefineXML(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	// Test a minimally valid xml
	xml := `
	<domain type="test">
		<name>test domain</name>
		<memory unit="KiB">8192</memory>
		<os>
			<type>hvm</type>
		</os>
	</domain>`
	dom, err := conn.DomainDefineXML(xml)
	if err != nil {
		t.Error(err)
		return
	}
	name, err := dom.GetName()
	if err != nil {
		t.Error(err)
		return
	}
	if name != "test domain" {
		t.Fatalf("Name was not 'test': %s", name)
		return
	}
	// And an invalid one
	xml = `<domain type="test"></domain>`
	_, err = conn.DomainDefineXML(xml)
	if err == nil {
		t.Fatal("Should have had an error")
		return
	}
}
