package recon

import (
	"fmt"
	"os"
	"time"
)

// SaveReport guarda un resumen de los hallazgos en un archivo de texto
func SaveReport(filename, domain string, ips []string, openPorts []int) {
	fmt.Printf("\n[*] Guardando reporte en: %s...\n", filename)

	// Creamos el archivo (si ya existe, lo sobrescribe)
	file, err := os.Create(filename)
	if err != nil {
		fmt.Printf("[-] Error al crear el archivo de reporte: %v\n", err)
		return
	}
	defer file.Close() // Aseguramos que el archivo se cierre al terminar

	// Escribimos el contenido en el archivo
	fecha := time.Now().Format("2006-01-02 15:04:05")
	fmt.Fprintf(file, "========================================\n")
	fmt.Fprintf(file, "REPORTE DE RECONOCIMIENTO - %s\n", domain)
	fmt.Fprintf(file, "Fecha del escaneo: %s\n", fecha)
	fmt.Fprintf(file, "========================================\n\n")

	fmt.Fprintf(file, "[+] Direcciones IP Encontradas:\n")
	for _, ip := range ips {
		fmt.Fprintf(file, "    - %s\n", ip)
	}

	fmt.Fprintf(file, "\n[+] Puertos Abiertos:\n")
	if len(openPorts) > 0 {
		for _, port := range openPorts {
			fmt.Fprintf(file, "    - %d\n", port)
		}
	} else {
		fmt.Fprintf(file, "    - Ninguno de los puertos comunes estaba abierto.\n")
	}

	fmt.Fprintf(file, "\n========================================\n")
	fmt.Fprintf(file, "Fin del reporte.\n")

	fmt.Println("[+] Reporte guardado con éxito.")
}
