package main

import (
	"github.com/alexzorin/libvirt-go"
	"log"
)

func main() {
	vir, err := libvirt.NewVirConnection("vbox:///session")
	if err != nil {
		log.Fatalln(err.Error())
	}

	hostname, _ := vir.GetHostname()
	log.Printf("Connected to hypervisor at %s\n", hostname)

	// capabilities, err := vir.GetCapabilities()
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	domains, err := vir.ListDomains()
	if err != nil {
		log.Fatalln(err.Error())
	}

	for k := range domains {
		log.Printf("Domain %d\n", domains[k])
		dom, err := vir.LookupDomainById(domains[k])
		if err != nil {
			log.Fatalln(err.Error())
		}
		name, _ := dom.GetName()
		log.Println(name)
	}

	vir.CloseConnection()
}
