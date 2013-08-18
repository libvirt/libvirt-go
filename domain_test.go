package libvirt

import (
	"testing"
)

func buildTestDomain() VirDomain {
	conn := buildTestConnection()
	dom, _ := conn.LookupDomainById(1)
	return dom
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
	_, err := dom.GetUUIDString()
	if err != nil {
		t.Error(err)
	}
}

func TestGetDomainInfo(t *testing.T) {
	dom := buildTestDomain()
	_, err := dom.GetInfo()
	if err != nil {
		t.Error(err)
	}
}

func TestGetDomainXMLDesc(t *testing.T) {
	dom := buildTestDomain()
	_, err := dom.GetXMLDesc(0)
	if err != nil {
		t.Error(err)
	}
}

func TestCreateDomainSnapshotXML(t *testing.T) {
	dom := buildTestDomain()
	_, err := dom.CreateSnapshotXML(`
		<domainsnapshot>
			<description>Test snapshot that will fail because its unsupported</description>
		</domainsnapshot>
	`, 0)
	if err == nil {
		t.Error("Snapshot should have failed due to being unsupported on test transport")
	}
}
