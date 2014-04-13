package libvirt

import (
	"testing"
	"time"
)

func buildTestConnection() VirConnection {
	conn, err := NewVirConnection("test:///default")
	if err != nil {
		panic(err)
	}
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

func TestGetType(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	tp, err := conn.GetType()
	if err != nil {
		t.Error(err)
		return
	}
	if tp != "Test" {
		t.Fatalf("type should have been test: %s", tp)
		return
	}
}

func TestIsAlive(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	alive, err := conn.IsAlive()
	if err != nil {
		t.Error(err)
		return
	}
	if !alive {
		t.Fatal("Connection should be alive")
		return
	}
}

func TestIsEncryptedAndSecure(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	secure, err := conn.IsSecure()
	if err != nil {
		t.Log(err)
		return
	}
	enc, err := conn.IsEncrypted()
	if err != nil {
		t.Error(err)
		return
	}
	if !secure {
		t.Fatal("Test driver should be secure")
		return
	}
	if enc {
		t.Fatal("Test driver should not be encrypted")
		return
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
	if doms == nil {
		t.Fatal("ListDefinedDomains shouldn't be nil")
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
	if doms == nil {
		t.Fatal("ListDomains shouldn't be nil")
		return
	}
}

func TestListInterfaces(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	_, err := conn.ListInterfaces()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestListNetworks(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	_, err := conn.ListNetworks()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestListStoragePools(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	_, err := conn.ListStoragePools()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestLookupDomainById(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	ids, err := conn.ListDomains()
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(ids)
	if len(ids) == 0 {
		t.Fatal("Length of ListDomains shouldn't be zero")
		return
	}
	if _, err := conn.LookupDomainById(ids[0]); err != nil {
		t.Error(err)
		return
	}
}

func TestLookupInvalidDomainById(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	_, err := conn.LookupDomainById(12345)
	if err == nil {
		t.Error("Domain #12345 shouldn't exist in test transport")
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
	defName := time.Now().String()
	xml := `<domain type="test">
		<name>` + defName + `</name>
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
	defer dom.Undefine()
	name, err := dom.GetName()
	if err != nil {
		t.Error(err)
		return
	}
	if name != defName {
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

func TestListDefinedInterfaces(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	_, err := conn.ListDefinedInterfaces()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestListDefinedNetworks(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	_, err := conn.ListDefinedNetworks()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestListDefinedStoragePools(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	_, err := conn.ListDefinedStoragePools()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestNumOfDefinedInterfaces(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	if _, err := conn.NumOfDefinedInterfaces(); err != nil {
		t.Error(err)
		return
	}
}

func TestNumOfDefinedNetworks(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	if _, err := conn.NumOfDefinedNetworks(); err != nil {
		t.Error(err)
		return
	}
}

func TestNumOfDefinedStoragePools(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	if _, err := conn.NumOfDefinedStoragePools(); err != nil {
		t.Error(err)
		return
	}
}

func TestNumOfDomains(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	if _, err := conn.NumOfDomains(); err != nil {
		t.Error(err)
		return
	}
}

func TestNumOfInterfaces(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	if _, err := conn.NumOfInterfaces(); err != nil {
		t.Error(err)
		return
	}
}

func TestNumOfNetworks(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	if _, err := conn.NumOfNetworks(); err != nil {
		t.Error(err)
		return
	}
}

func TestNumOfNWFilters(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	if _, err := conn.NumOfNWFilters(); err == nil {
		t.Fatalf("NumOfNWFilters should fail due to no support on test driver")
		return
	}
}

func TestGetURI(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
	uri, err := conn.GetURI()
	if err != nil {
		t.Error(err)
	}
	origUri := "test:///default"
	if uri != origUri {
		t.Fatalf("should be %s but got %s", origUri, uri)
	}
}
