package recon

import (
	"fmt"
	"net"
)

// GetDNSInfo busca y muestra información básica de DNS para un dominio dado.
func GetDNSInfo(domain string) {
	fmt.Printf("\n[*] Extrayendo información de DNS para: %s\n", domain)

	// Buscar direcciones IP
	ips, err := net.LookupIP(domain)
	if err != nil {
		fmt.Printf("[-] Error al buscar IPs: %v\n", err)
	} else {
		fmt.Println("[+] Direcciones IP:")
		for _, ip := range ips {
			fmt.Printf("\t- %s\n", ip.String())
		}
	}

	// Buscar Servidores de Nombres (NS)
	nss, err := net.LookupNS(domain)
	if err != nil {
		fmt.Printf("[-] Error al buscar Servidores de Nombres: %v\n", err)
	} else {
		fmt.Println("[+] Servidores de Nombres (NS):")
		for _, ns := range nss {
			fmt.Printf("\t- %s\n", ns.Host)
		}
	}
}
