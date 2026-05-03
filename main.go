package main

import (
	"flag"
	"fmt"
	"os"
	"recon-bot/recon"
	"regexp"
)

// isValidDomain verifica que el dominio tenga un formato válido y no incluya
// protocolos (http://), prefijos 'www.' o caracteres extraños.
func isValidDomain(domain string) bool {
	// Expresión regular para validar el formato de dominio:
	// - Debe contener etiquetas de caracteres alfanuméricos o guiones separadas por puntos.
	// - Debe terminar con un TLD de al menos 2 caracteres.
	// - Al no incluir ':' ni '/', rechazamos automáticamente URLs con protocolo.
	reFormat := regexp.MustCompile(`^([a-zA-Z0-9-]+\.)+[a-zA-Z]{2,}$`)
	if !reFormat.MatchString(domain) {
		return false
	}

	// El requerimiento especifica rechazar dominios que empiecen con 'www.'
	reWWW := regexp.MustCompile(`^www\.`)
	if reWWW.MatchString(domain) {
		return false
	}

	return true
}

func main() {
	// Configurar bandera -d
	domain := flag.String("d", "", "Dominio para realizar el reconocimiento (ej: google.com)")
	flag.Parse()

	// Validar que se haya provisto el dominio
	if *domain == "" {
		fmt.Println("Uso: recon-bot -d <dominio>")
		os.Exit(1)
	}

	// Validar el formato del dominio
	if !isValidDomain(*domain) {
		fmt.Printf("[-] Error: El dominio '%s' es inválido.\n", *domain)
		fmt.Println("Asegúrate de ingresar un dominio limpio (ej: empresa.com) sin http:// o www.")
		os.Exit(1)
	}

	// Imprimir banner y mensaje de inicio
	fmt.Println("#########################################")
	fmt.Println("#               RECON-BOT               #")
	fmt.Println("#      Herramienta de Reconocimiento    #")
	fmt.Println("#########################################")
	fmt.Printf("[+] Iniciando reconocimiento en: %s\n", *domain)

	// Obtener información de DNS
	recon.GetDNSInfo(*domain)
}
