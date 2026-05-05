package recon

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

// Estructura para mapear el JSON que nos devuelve crt.sh
type CrtShEntry struct {
	NameValue string `json:"name_value"`
}

// GetSubdomains busca subdominios usando los logs de Certificate Transparency
func GetSubdomains(domain string) {
	fmt.Printf("\n[*] Buscando subdominios pasivos en crt.sh para: %s\n", domain)

	// Usamos %25. para que busque todo lo que esté antes del dominio
	url := fmt.Sprintf("https://crt.sh/?q=%%25.%s&output=json", domain)

	client := &http.Client{Timeout: 15 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[-] Error conectando a crt.sh: %v\n", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		fmt.Printf("[-] La API de crt.sh no está disponible en este momento (Status: %d)\n", resp.StatusCode)
		return
	}

	body, _ := io.ReadAll(resp.Body)

	var entries []CrtShEntry
	if err := json.Unmarshal(body, &entries); err != nil {
		fmt.Printf("[-] Error procesando el JSON: %v\n", err)
		return
	}

	// Usamos un mapa para eliminar los subdominios duplicados
	uniqueSubs := make(map[string]bool)
	for _, entry := range entries {
		// Limpiamos los saltos de línea y los comodines (*)
		cleanName := strings.ToLower(entry.NameValue)
		cleanName = strings.ReplaceAll(cleanName, "*.", "")

		// Separamos si vienen varios dominios en la misma línea
		parts := strings.Split(cleanName, "\n")
		for _, p := range parts {
			if p != domain && p != "" && !uniqueSubs[p] {
				uniqueSubs[p] = true
			}
		}
	}

	if len(uniqueSubs) == 0 {
		fmt.Println("\t- No se encontraron subdominios en los registros.")
		return
	}

	fmt.Printf("[+] Se encontraron %d subdominios (mostrando hasta 15):\n", len(uniqueSubs))
	count := 0
	for sub := range uniqueSubs {
		if count >= 15 {
			fmt.Println("\t- ... (y muchos más)")
			break
		}
		fmt.Printf("\t- %s\n", sub)
		count++
	}
}
