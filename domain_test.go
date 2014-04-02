package libvirt

import (
	"testing"
	"time"
)

func buildTestDomain() (VirDomain, VirConnection) {
	conn := buildTestConnection()
	dom, err := conn.DomainDefineXML(`<domain type="test">
		<name>` + time.Now().String() + `</name>
		<memory unit="KiB">8192</memory>
		<os>
			<type>hvm</type>
		</os>
	</domain>`)
	if err != nil {
		panic(err)
	}
	return dom, conn
}

func TestGetDomainName(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	if _, err := dom.GetName(); err != nil {
		t.Error(err)
		return
	}
}

func TestGetDomainState(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	state, err := dom.GetState()
	if err != nil {
		t.Error(err)
		return
	}
	if len(state) != 2 {
		t.Error("Length of domain state should be 2")
		return
	}
	if state[0] != 5 || state[1] != 0 {
		t.Error("Domain state in test transport should be [5 0]")
		return
	}
}

func TestGetDomainUUID(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	_, err := dom.GetUUID()
	// how to test uuid validity?
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetDomainUUIDString(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	_, err := dom.GetUUIDString()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetDomainInfo(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	_, err := dom.GetInfo()
	if err != nil {
		t.Error(err)
		return
	}
}

func TestGetDomainXMLDesc(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	_, err := dom.GetXMLDesc(0)
	if err != nil {
		t.Error(err)
		return
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
	if err != nil {
		t.Error(err)
		return
	}
}

func TestSaveDomain(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	// get the name so we can get a handle on it later
	domName, err := dom.GetName()
	if err != nil {
		t.Error(err)
		return
	}
	const tmpFile = "/tmp/libvirt-go-test.tmp"
	if err := dom.Save(tmpFile); err != nil {
		t.Error(err)
		return
	}
	if err := conn.Restore(tmpFile); err != nil {
		t.Error(err)
		return
	}
	if _, err = conn.LookupDomainByName(domName); err != nil {
		t.Error(err)
		return
	}
}

func TestSaveDomainFlags(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	const srcFile = "/tmp/libvirt-go-test.tmp"
	if err := dom.SaveFlags(srcFile, "", 0); err == nil {
		t.Fatal("expected xml modification unsupported")
		return
	}
}

func TestCreateDestroyDomain(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	if err := dom.Create(); err != nil {
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

func TestShutdownDomain(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	if err := dom.Create(); err != nil {
		t.Error(err)
		return
	}
	if err := dom.Shutdown(); err != nil {
		t.Error(err)
		return
	}
	state, err := dom.GetState()
	if err != nil {
		t.Error(err)
		return
	}
	if state[0] != 5 || state[1] != 1 {
		t.Fatal("state should be [5 1]")
		return
	}
}

func TestShutdownReboot(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	if err := dom.Reboot(0); err != nil {
		t.Error(err)
		return
	}
}

func TestAutostart(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	as, err := dom.GetAutostart()
	if err != nil {
		t.Error(err)
		return
	}
	if as {
		t.Fatal("autostart should be false")
		return
	}
	if err := dom.SetAutostart(true); err != nil {
		t.Error(err)
		return
	}
	as, err = dom.GetAutostart()
	if err != nil {
		t.Error(err)
		return
	}
	if !as {
		t.Fatal("autostart should be true")
		return
	}
}

func TestDomainIsActive(t *testing.T) {
	dom, conn := buildTestDomain()
	defer conn.CloseConnection()
	if err := dom.Create(); err != nil {
		t.Log(err)
		return
	}
	active, err := dom.IsActive()
	if err != nil {
		t.Error(err)
		return
	}
	if !active {
		t.Fatal("Domain should be active")
		return
	}
	if err := dom.Destroy(); err != nil {
		t.Error(err)
		return
	}
	active, err = dom.IsActive()
	if err != nil {
		t.Error(err)
		return
	}
	if active {
		t.Fatal("Domain should be inactive")
		return
	}
}
