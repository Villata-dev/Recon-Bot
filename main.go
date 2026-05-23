package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	// Ajusta "recon-bot/recon" si el módulo de tu go.mod se llama diferente
	"recon-bot/recon"
)

// isValidDomain valida que la entrada tenga estrictamente estructura de dominio
func isValidDomain(domain string) bool {
	regex := regexp.MustCompile(`^[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(domain)
}

func main() {
	// Configuración de banderas de consola
	domainPtr := flag.String("d", "", "Dominio objetivo para realizar el reconocimiento")
	outPtr := flag.String("o", "", "Archivo de salida opcional para guardar el reporte (.txt)")
	flag.Parse()

	// Validación 1: Verificar que se ingresó un dominio
	if *domainPtr == "" {
		fmt.Println("[-] Error: Debes especificar un dominio objetivo.")
		fmt.Println("Uso: go run main.go -d <dominio.com> [-o reporte.txt]")
		os.Exit(1)
	}

	// Validación 2: Verificar el formato correcto (sin http/www)
	if !isValidDomain(*domainPtr) {
		fmt.Println("[-] Error: Formato de dominio inválido. No uses 'https://' ni 'www'.")
		fmt.Println("Ejemplo correcto: go run main.go -d aiep.cl")
		os.Exit(1)
	}

	// BANNER
	fmt.Println("#########################################")
	fmt.Println("#                RECON-BOT              #")
	fmt.Println("#      Herramienta de Reconocimiento    #")
	fmt.Println("#########################################")
	fmt.Printf("[+] Iniciando reconocimiento en: %s\n", *domainPtr)

	// FASE 1: Enumeración DNS Básica
	ips, err := recon.GetDNSInfo(*domainPtr)
	if err != nil {
		fmt.Printf("[-] Error crítico en la fase DNS: %v\n", err)
		os.Exit(1)
	}

	// FASE 2 y 3: Geolocalización y Escaneo de Puertos Concurrente
	var puertosAbiertos []int
	if len(ips) > 0 {
		targetIP := ips[0]
		recon.GetGeoIP(targetIP)
		// Ejecuta de forma concurrente y nos devuelve el slice de abiertos
		puertosAbiertos = recon.ScanPorts(targetIP)
	} else {
		fmt.Println("[-] No se encontraron IPs válidas para Geolocalización o Escaneo de Puertos.")
	}

	// FASE 4: Banner Grabbing (HTTP Headers)
	recon.GetHTTPHeaders(*domainPtr)

	// FASE 5: Enumeración de Subdominios (Certificate Transparency)
	recon.GetSubdomains(*domainPtr)

	// FASE 6: Persistencia (Guardar reporte si se usó la bandera -o)
	if *outPtr != "" {
		recon.SaveReport(*outPtr, *domainPtr, ips, puertosAbiertos)
	}
}
