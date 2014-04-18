// +build integration

package libvirt

import (
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
