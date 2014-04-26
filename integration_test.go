// +build integration

package libvirt

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
	"time"
)

func defineTestLxcDomain(conn VirConnection, title string) (VirDomain, error) {
	if title == "" {
		title = time.Now().String()
	}
	xml := `<domain type='lxc'>
	  <name>` + title + `</name>
	  <title>` + title + `</title>
	  <memory>102400</memory>
	  <os>
	    <type>exe</type>
	    <init>/bin/sh</init>
	  </os>
	  <devices>
	    <console type='pty'/>
	  </devices>
	</domain>`
	dom, err := conn.DomainDefineXML(xml)
	return dom, err
}

// Integration tests are run against LXC using Libvirt 1.2.x
// on Debian Wheezy (libvirt from wheezy-backports)
//
// To run,
// 		go test -tags integration

func TestIntegrationGetMetadata(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	title := time.Now().String()
	dom, err := defineTestLxcDomain(conn, title)
	if err != nil {
		t.Error(err)
		return
	}
	defer dom.Free()
	if err := dom.Create(); err != nil {
		t.Error(err)
		return
	}
	v, err := dom.GetMetadata(VIR_DOMAIN_METADATA_TITLE, "", 0)
	dom.Destroy()
	if err != nil {
		t.Error(err)
		return
	}
	if v != title {
		t.Fatal("title didnt match: expected %s, got %s", title, v)
		return
	}
	if err := dom.Undefine(); err != nil {
		t.Error(err)
		return
	}
}

func TestIntegrationSetMetadata(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	dom, err := defineTestLxcDomain(conn, "")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		dom.Undefine()
		dom.Free()
	}()
	const domTitle = "newtitle"
	if err := dom.SetMetadata(VIR_DOMAIN_METADATA_TITLE, domTitle, "", "", 0); err != nil {
		t.Error(err)
		return
	}
	v, err := dom.GetMetadata(VIR_DOMAIN_METADATA_TITLE, "", 0)
	if err != nil {
		t.Error(err)
		return
	}
	if v != domTitle {
		t.Fatalf("VIR_DOMAIN_METADATA_TITLE should have been %s, not %s", domTitle, v)
		return
	}
}

func TestIntegrationGetSysinfo(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	info, err := conn.GetSysinfo(0)
	if err != nil {
		t.Error(err)
		return
	}
	if strings.Index(info, "<sysinfo") != 0 {
		t.Fatalf("Sysinfo not valid: %s", info)
		return
	}
}

func testNWFilterXML(name, chain string) string {
	defName := name
	if defName == "" {
		defName = time.Now().String()
	}
	return `<filter name='` + defName + `' chain='` + chain + `'>
            <rule action='drop' direction='out' priority='500'>
            <ip match='no' srcipaddr='$IP'/>
            </rule>
			</filter>`
}

func TestIntergrationDefineUndefineNWFilterXML(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML("", "ipv4"))
	if err != nil {
		t.Fatal(err)
	}
	defer func() {
		if err := filter.Undefine(); err != nil {
			t.Fatal(err)
		}
		filter.Free()
	}()
	_, err = conn.NWFilterDefineXML(testNWFilterXML("", "bad"))
	if err == nil {
		t.Fatal("Should have had an error")
	}
}

func TestIntegrationNWFilterGetName(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML("", "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()
	if _, err := filter.GetName(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationNWFilterGetUUID(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML("", "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()
	if _, err := filter.GetUUID(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationNWFilterGetUUIDString(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML("", "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()
	if _, err := filter.GetUUIDString(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationNWFilterGetXMLDesc(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML("", "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()
	if _, err := filter.GetXMLDesc(0); err != nil {
		t.Error(err)
	}
}

func TestIntegrationLookupNWFilterByName(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	origName := time.Now().String()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML(origName, "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()
	filter, err = conn.LookupNWFilterByName(origName)
	if err != nil {
		t.Error(err)
		return
	}
	var newName string
	newName, err = filter.GetName()
	if err != nil {
		t.Error(err)
		return
	}
	if newName != origName {
		t.Fatalf("expected filter name: %s ,got: %s", origName, newName)
	}
}

func TestIntegrationLookupNWFilterByUUIDString(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	origName := time.Now().String()
	filter, err := conn.NWFilterDefineXML(testNWFilterXML(origName, "ipv4"))
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		filter.Undefine()
		filter.Free()
	}()
	filter, err = conn.LookupNWFilterByName(origName)
	if err != nil {
		t.Error(err)
		return
	}
	var filterUUID string
	filterUUID, err = filter.GetUUIDString()
	if err != nil {
		t.Error(err)
		return
	}
	filter, err = conn.LookupNWFilterByUUIDString(filterUUID)
	if err != nil {
		t.Error(err)
		return
	}
	name, err := filter.GetName()
	if err != nil {
		t.Error(err)
		return
	}
	if name != origName {
		t.Fatalf("fetching by UUID: expected filter name: %s ,got: %s", name, origName)
	}
}

func TestIntegrationDomainAttachDetachDevice(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()

	dom, err := defineTestLxcDomain(conn, "")
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		dom.Undefine()
		dom.Free()
	}()
	const nwXml = `<interface type='network'>
		<mac address='52:54:00:37:aa:c7'/>
		<source network='default'/>
		<model type='virtio'/>
		</interface>`
	if err := dom.AttachDeviceFlags(nwXml, VIR_DOMAIN_DEVICE_MODIFY_CONFIG); err != nil {
		t.Error(err)
		return
	}
	if err := dom.DetachDeviceFlags(nwXml, VIR_DOMAIN_DEVICE_MODIFY_CONFIG); err != nil {
		t.Error(err)
		return
	}
}

func TestStorageVolResize(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()

	poolPath, err := ioutil.TempDir("", "default-pool-test-1")
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(poolPath)
	pool, err := conn.StoragePoolDefineXML(`<pool type='dir'>
                                          <name>default-pool-test-1</name>
                                          <target>
                                          <path>`+poolPath+`</path>
                                          </target>
                                          </pool>`, 0)
	defer func() {
		pool.Undefine()
		pool.Free()
	}()
	if err := pool.Create(0); err != nil {
		t.Error(err)
		return
	}
	defer pool.Destroy()
	vol, err := pool.StorageVolCreateXML(testStorageVolXML("", poolPath), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		vol.Delete(VIR_STORAGE_VOL_DELETE_NORMAL)
		vol.Free()
	}()
	const newCapacityInBytes = 12582912
	if err := vol.Resize(newCapacityInBytes, VIR_STORAGE_VOL_RESIZE_ALLOCATE); err != nil {
		t.Fatal(err)
	}
}

func TestStorageVolWipe(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()

	poolPath, err := ioutil.TempDir("", "default-pool-test-1")
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(poolPath)
	pool, err := conn.StoragePoolDefineXML(`<pool type='dir'>
                                          <name>default-pool-test-1</name>
                                          <target>
                                          <path>`+poolPath+`</path>
                                          </target>
                                          </pool>`, 0)
	defer func() {
		pool.Undefine()
		pool.Free()
	}()
	if err := pool.Create(0); err != nil {
		t.Error(err)
		return
	}
	defer pool.Destroy()
	vol, err := pool.StorageVolCreateXML(testStorageVolXML("", poolPath), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		vol.Delete(VIR_STORAGE_VOL_DELETE_NORMAL)
		vol.Free()
	}()
	if err := vol.Wipe(0); err != nil {
		t.Fatal(err)
	}
}

func TestStorageVolWipePattern(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()

	poolPath, err := ioutil.TempDir("", "default-pool-test-1")
	if err != nil {
		t.Error(err)
		return
	}
	defer os.RemoveAll(poolPath)
	pool, err := conn.StoragePoolDefineXML(`<pool type='dir'>
                                          <name>default-pool-test-1</name>
                                          <target>
                                          <path>`+poolPath+`</path>
                                          </target>
                                          </pool>`, 0)
	defer func() {
		pool.Undefine()
		pool.Free()
	}()
	if err := pool.Create(0); err != nil {
		t.Error(err)
		return
	}
	defer pool.Destroy()
	vol, err := pool.StorageVolCreateXML(testStorageVolXML("", poolPath), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		vol.Delete(VIR_STORAGE_VOL_DELETE_NORMAL)
		vol.Free()
	}()
	if err := vol.WipePattern(VIR_STORAGE_VOL_WIPE_ALG_ZERO, 0); err != nil {
		t.Fatal(err)
	}
}

func testSecretTypeCephFromXML(name string) string {
	var setName string
	if name == "" {
		setName = time.Now().String()
	} else {
		setName = name
	}
	return `<secret ephemeral='no' private='no'>
            <usage type='ceph'>
            <name>` + setName + `</name>
            </usage>
            </secret>`
}

func TestIntegrationSecretDefineUndefine(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(""), 0)
	if err != nil {
		t.Fatal(err)
	}
	defer sec.Free()

	if err := sec.Undefine(); err != nil {
		t.Fatal(err)
	}
}

func TestIntegrationSecretGetUUID(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(""), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		sec.Undefine()
		sec.Free()
	}()
	if _, err := sec.GetUUID(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationSecretGetUUIDString(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(""), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		sec.Undefine()
		sec.Free()
	}()
	if _, err := sec.GetUUIDString(); err != nil {
		t.Error(err)
	}
}

func TestIntegrationSecretGetXMLDesc(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(""), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		sec.Undefine()
		sec.Free()
	}()
	if _, err := sec.GetXMLDesc(0); err != nil {
		t.Error(err)
	}
}

func TestIntegrationSecretGetUsageType(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(""), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		sec.Undefine()
		sec.Free()
	}()
	uType, err := sec.GetUsageType()
	if err != nil {
		t.Error(err)
		return
	}
	if uType != VIR_SECRET_USAGE_TYPE_CEPH {
		t.Fatal("unexpected usage type.Expected usage type is Ceph")
	}
}

func TestIntegrationSecretGetUsageID(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	setUsageID := time.Now().String()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(setUsageID), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		sec.Undefine()
		sec.Free()
	}()
	recUsageID, err := sec.GetUsageID()
	if err != nil {
		t.Error(err)
		return
	}
	if recUsageID != setUsageID {
		t.Fatalf("exepected usage ID: %s, got: %s", setUsageID, recUsageID)
	}
}

func TestIntegrationLookupSecretByUsage(t *testing.T) {
	conn, err := NewVirConnection("lxc:///")
	if err != nil {
		t.Error(err)
		return
	}
	defer conn.CloseConnection()
	usageID := time.Now().String()
	sec, err := conn.SecretDefineXML(testSecretTypeCephFromXML(usageID), 0)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		sec.Undefine()
		sec.Free()
	}()
	sec, err = conn.LookupSecretByUsage(VIR_SECRET_USAGE_TYPE_CEPH, usageID)
	if err != nil {
		t.Fatal(err)
	}
}

