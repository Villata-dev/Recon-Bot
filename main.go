package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"

	// Reemplaza "recon-bot" por el nombre exacto de tu módulo si es diferente
	"recon-bot/recon"
)

// isValidDomain verifica usando una expresión regular si el formato del dominio es correcto
func isValidDomain(domain string) bool {
	// Expresión regular para dominios estándar (ej: google.com, sub.dominio.cl)
	regex := regexp.MustCompile(`^[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	return regex.MatchString(domain)
}

func main() {
	// Configuración de la bandera de consola -d
	domainPtr := flag.String("d", "", "Dominio objetivo para realizar el reconocimiento")
	flag.Parse()

	// Validación: Si no se pasa el argumento, mostramos cómo se usa
	if *domainPtr == "" {
		fmt.Println("[-] Error: Debes especificar un dominio.")
		fmt.Println("Uso: go run main.go -d <dominio.com>")
		os.Exit(1)
	}

	// Validación: Verificar que el dominio no traiga http:// o basura
	if !isValidDomain(*domainPtr) {
		fmt.Println("[-] Error: Formato de dominio inválido. No uses 'https://' ni 'www'.")
		fmt.Println("Ejemplo correcto: go run main.go -d aiep.cl")
		os.Exit(1)
	}

	// BANNER DE BIENVENIDA
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

	// FASE 2 y 3: Geolocalización y Escaneo de Puertos (Usando la primera IP resuelta)
	if len(ips) > 0 {
		targetIP := ips[0]
		recon.GetGeoIP(targetIP)
		recon.ScanPorts(targetIP)
	} else {
		fmt.Println("[-] No se encontraron IPs válidas para Geolocalización o Escaneo de Puertos.")
	}

	// FASE 4: Banner Grabbing (Análisis de Cabeceras HTTP)
	recon.GetHTTPHeaders(*domainPtr)

	// FASE 5: Enumeración Pasiva de Subdominios (crt.sh)
	recon.GetSubdomains(*domainPtr)
}
