package recon

import (
	"fmt"
	"net"
	"sync"
	"time"
)

// ScanPorts escanea una lista de puertos TCP comunes usando concurrencia
func ScanPorts(ip string) {
	fmt.Printf("\n[*] Escaneando puertos comunes en la IP: %s (Modo Concurrente)\n", ip)

	commonPorts := []int{21, 22, 23, 25, 53, 80, 110, 139, 443, 445, 3306, 3389, 8080, 8443}
	timeout := 2 * time.Second

	var wg sync.WaitGroup
	var mu sync.Mutex
	var openPorts []int

	// Lanzamos una Goroutine por cada puerto
	for _, port := range commonPorts {
		wg.Add(1) // Sumamos 1 al contador de tareas pendientes

		// La palabra clave 'go' lanza esto en un hilo concurrente (Goroutine)
		go func(p int) {
			defer wg.Done() // Restamos 1 al contador cuando esta función termine

			target := fmt.Sprintf("%s:%d", ip, p)
			conn, err := net.DialTimeout("tcp", target, timeout)

			if err == nil {
				// Mutex evita "Race Conditions" (que 2 hilos choquen al guardar datos)
				mu.Lock()
				openPorts = append(openPorts, p)
				mu.Unlock()

				conn.Close()
			}
		}(port) // Pasamos 'port' como argumento (p) para evitar bugs de memoria en el loop
	}

	// Esperamos a que todas las goroutines llamen a wg.Done()
	wg.Wait()

	// Imprimimos los resultados recolectados
	if len(openPorts) > 0 {
		fmt.Println("[+] Resultados de puertos abiertos:")
		for _, port := range openPorts {
			fmt.Printf("\t- Puerto %d [ABIERTO]\n", port)
		}
	} else {
		fmt.Println("\t- No se encontraron puertos comunes abiertos.")
	}
}
