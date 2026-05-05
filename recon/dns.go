package recon

import (
	"fmt"
	"net"
)

// GetDNSInfo ahora retorna las IPs encontradas para que otros módulos las usen
func GetDNSInfo(domain string) ([]string, error) {
	fmt.Printf("\n[*] Extrayendo información de DNS para: %s\n", domain)

	ips, err := net.LookupIP(domain)
	if err != nil {
		return nil, fmt.Errorf("error resolviendo IP: %v", err)
	}

	var ipStrings []string
	fmt.Println("[+] Direcciones IP:")
	for _, ip := range ips {
		ipStr := ip.String()
		ipStrings = append(ipStrings, ipStr)
		fmt.Printf("\t- %s\n", ipStr)
	}

	nsRecords, err := net.LookupNS(domain)
	if err == nil {
		fmt.Println("[+] Servidores de Nombres (NS):")
		for _, ns := range nsRecords {
			fmt.Printf("\t- %s\n", ns.Host)
		}
	}

	return ipStrings, nil
}
