package libvirt

import (
	"testing"
)

func buildTestDomain() (VirDomain, VirConnection) {
	conn := buildTestConnection()
	dom, _ := conn.LookupDomainById(1)
	return dom, conn
}

func TestGetDomainName(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	name, err := dom.GetName()
	if err != nil {
		t.Error(err)
	}
	if name != "test" {
		t.Error("Name of active domain in test transport should be 'test'")
	}
}

func TestGetDomainState(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
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
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	_, err := dom.GetUUID()
	// how to test uuid validity?
	if err != nil {
		t.Error(err)
	}
}

func TestGetDomainUUIDString(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	_, err := dom.GetUUIDString()
	if err != nil {
		t.Error(err)
	}
}

func TestGetDomainInfo(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	_, err := dom.GetInfo()
	if err != nil {
		t.Error(err)
	}
}

func TestGetDomainXMLDesc(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	_, err := dom.GetXMLDesc(0)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateDomainSnapshotXML(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	_, err := dom.CreateSnapshotXML(`
		<domainsnapshot>
			<description>Test snapshot that will fail because its unsupported</description>
		</domainsnapshot>
	`, 0)
	if err == nil {
		t.Error("Snapshot should have failed due to being unsupported on test transport")
	}
}

func TestSaveDomain(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	err := dom.Save("/tmp/libvirt-go-test.tmp")
	if err != nil {
		t.Error(err)
	}
}

func TestSaveDomainFlags(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	err := dom.SaveFlags("/tmp/libvirt-go-test.tmp", "", 0)
	if err == nil {
		t.Error("Excected xml modification unsupported")
	}
}

func TestCreateDestroyDomain(t *testing.T) {
	conn := buildTestConnection()
	defer conn.CloseConnection()
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
	if err = dom.Create(); err != nil {
		t.Error(err)
		return
	}
	state, err := dom.GetState()
	if err != nil {
		t.Error(err)
		return
	}
	if state[0] != VIR_DOMAIN_RUNNING {
		t.Fatal("Domain should be running")
		return
	}
	if err = dom.Destroy(); err != nil {
		t.Error(err)
		return
	}
	state, err = dom.GetState()
	if err != nil {
		t.Error(err)
		return
	}
	if state[0] != VIR_DOMAIN_SHUTOFF {
		t.Fatal("Domain should be destroyed")
		return
	}
}
