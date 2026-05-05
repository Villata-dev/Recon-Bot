package recon

import (
	"fmt"
	"net/http"
	"time"
)

// GetHTTPHeaders extrae cabeceras del servidor web para identificar tecnologías
func GetHTTPHeaders(domain string) {
	fmt.Printf("\n[*] Extrayendo cabeceras HTTP (Banner Grabbing) para: %s\n", domain)

	// Armamos la URL asumiendo HTTPS (ya que vimos el 443 abierto)
	url := "https://" + domain

	// Configuramos un cliente con un timeout para no quedarnos pegados
	client := &http.Client{
		Timeout: 5 * time.Second,
	}

	// Usamos HEAD en lugar de GET para ahorrar ancho de banda
	req, err := http.NewRequest("HEAD", url, nil)
	if err != nil {
		fmt.Printf("[-] Error creando la petición HTTP: %v\n", err)
		return
	}

	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("[-] Error conectando al servidor web: %v\n", err)
		return
	}
	defer resp.Body.Close()

	fmt.Println("[+] Cabeceras reveladoras encontradas:")

	// Lista de cabeceras que suelen filtrar información valiosa
	interestingHeaders := []string{
		"Server",
		"X-Powered-By",
		"X-AspNet-Version",
		"Via",
		"X-Frame-Options",
		"Cf-Ray", // Nos dice si usa Cloudflare
	}

	foundAny := false
	for _, h := range interestingHeaders {
		if val := resp.Header.Get(h); val != "" {
			fmt.Printf("\t- %s: %s\n", h, val)
			foundAny = true
		}
	}

	if !foundAny {
		fmt.Println("\t- No se filtró información en las cabeceras comunes. (¡Buen trabajo del sysadmin!)")
	}
}
