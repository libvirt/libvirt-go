package main

import (
	"github.com/alexzorin/libvirt"
	"log"
)

func main() {
	vir, err := libvirt.NewVirConnection("vbox:///session")
	if err != nil {
		log.Fatalln(err.Error())
	}

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
}
