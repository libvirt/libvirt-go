package libvirt

import (
	"crypto/rand"
	"fmt"
	"testing"
)

func buildTestInterface(mac string) (VirInterface, VirConnection) {
	conn := buildTestConnection()
	xml := `<interface type='ethernet' name='ethTest0'><mac address='` + mac + `'/></interface>`
	iface, err := conn.InterfaceDefineXML(xml, 0)
	if err != nil {
		panic(err)
	}
	return iface, conn
}

func generateRandomMac() string {
	macBuf := make([]byte, 3)
	if _, err := rand.Read(macBuf); err != nil {
		panic(err)
	}
	return fmt.Sprintf("aa:bb:cc:%02x:%02x:%02x", macBuf[0], macBuf[1], macBuf[2])
}
