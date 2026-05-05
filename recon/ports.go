package recon

import (
	"fmt"
	"net"
	"time"
)

// ScanPorts escanea una lista de puertos TCP comunes en una IP dada
func ScanPorts(ip string) {
	fmt.Printf("\n[*] Escaneando puertos comunes en la IP: %s\n", ip)

	// Lista de puertos "top" a escanear
	commonPorts := []int{21, 22, 23, 25, 53, 80, 110, 139, 443, 445, 3306, 3389, 8080, 8443}

	// Timeout corto (2 segundos) para no quedarnos pegados si el firewall bloquea
	timeout := 2 * time.Second

	fmt.Println("[+] Resultados de puertos abiertos:")
	openPortsFound := false

	for _, port := range commonPorts {
		target := fmt.Sprintf("%s:%d", ip, port)
		conn, err := net.DialTimeout("tcp", target, timeout)

		if err == nil {
			// Si no hay error (err == nil), la puerta está abierta
			fmt.Printf("\t- Puerto %d [ABIERTO]\n", port)
			conn.Close() // ¡Siempre hay que cerrar la conexión!
			openPortsFound = true
		}
	}

	if !openPortsFound {
		fmt.Println("\t- No se encontraron puertos comunes abiertos.")
	}
}
