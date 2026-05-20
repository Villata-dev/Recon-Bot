package recon

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

// Estructura para mapear la respuesta de ip-api.com
type GeoIPResponse struct {
	Status  string `json:"status"`
	Country string `json:"country"`
	City    string `json:"city"`
	ISP     string `json:"isp"`
	Org     string `json:"org"`
	AS      string `json:"as"`
}

// GetGeoIP obtiene la ubicación física y el proveedor de internet de una IP
func GetGeoIP(ip string) {
	fmt.Printf("\n[*] Geolocalizando la IP: %s\n", ip)

	url := fmt.Sprintf("http://ip-api.com/json/%s", ip)

	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		fmt.Printf("[-] Error conectando al servicio de GeoIP: %v\n", err)
		return
	}
	defer resp.Body.Close()

	var geoData GeoIPResponse
	if err := json.NewDecoder(resp.Body).Decode(&geoData); err != nil {
		fmt.Printf("[-] Error decodificando datos geográficos: %v\n", err)
		return
	}

	if geoData.Status != "success" {
		fmt.Println("[-] No se pudo geolocalizar esta IP (puede ser una IP privada o local).")
		return
	}

	fmt.Println("[+] Información Geográfica y de Red:")
	fmt.Printf("\t- País: %s\n", geoData.Country)
	fmt.Printf("\t- Ciudad: %s\n", geoData.City)
	fmt.Printf("\t- ISP (Proveedor): %s\n", geoData.ISP)
	fmt.Printf("\t- Organización: %s\n", geoData.Org)
	fmt.Printf("\t- ASN: %s\n", geoData.AS)
}
