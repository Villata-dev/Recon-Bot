package recon

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// ScanPorts escanea una lista de puertos TCP comunes usando concurrencia y retorna los abiertos
func ScanPorts(ip string) []int {
	fmt.Printf("\n[*] Escaneando puertos comunes en la IP: %s (Modo Concurrente)\n", ip)

	commonPorts := []int{21, 22, 23, 25, 53, 80, 110, 139, 443, 445, 3306, 3389, 8080, 8443}
	timeout := 2 * time.Second

	var wg sync.WaitGroup
	var mu sync.Mutex
	var openPorts []int

	// Lanzamos una Goroutine por cada puerto de forma concurrente
	for _, port := range commonPorts {
		wg.Add(1)

		go func(p int) {
			defer wg.Done()

			target := fmt.Sprintf("%s:%d", ip, p)
			conn, err := net.DialTimeout("tcp", target, timeout)

			if err == nil {
				// Bloqueamos el canal con Mutex para evitar colisiones al escribir en el slice
				mu.Lock()
				openPorts = append(openPorts, p)
				mu.Unlock()

				conn.Close()
			}
		}(port)
	}

	// Esperamos a que terminen todas las consultas concurrentes
	wg.Wait()

	// Imprimimos el resultado en pantalla
	if len(openPorts) > 0 {
		fmt.Println("[+] Resultados de puertos abiertos:")
		for _, port := range openPorts {
			fmt.Printf("\t- Puerto %d [ABIERTO]\n", port)
		}
	} else {
		fmt.Println("\t- No se encontraron puertos comunes abiertos.")
	}

	// Retornamos el slice con los puertos para que el generador de reportes los use
	return openPorts
}
