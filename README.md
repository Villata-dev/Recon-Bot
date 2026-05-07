# Recon-Bot 🚀

Recon-Bot es una herramienta de línea de comandos (CLI) desarrollada en Golang, diseñada para realizar tareas de reconocimiento pasivo y activo de forma rápida y eficiente. Es ideal para profesionales de ciberseguridad y entusiastas del OSINT que buscan automatizar la fase de recolección de información sobre un dominio objetivo.

## ✨ Características (Features)

La herramienta integra cuatro funcionalidades clave para el reconocimiento:

1.  **Enumeración DNS:** Extrae registros A (direcciones IP) y registros NS (Servidores de Nombres) asociados al dominio.
2.  **Escaneo de Puertos TCP:** Identifica puertos comunes abiertos (21, 22, 80, 443, etc.) mediante escaneo activo.
3.  **Banner Grabbing HTTP:** Analiza las cabeceras de respuesta del servidor web para identificar tecnologías, versiones de software y configuraciones de seguridad.
4.  **Búsqueda de Subdominios:** Descubre subdominios asociados utilizando los registros públicos de Certificate Transparency a través de la API de `crt.sh`.

## 🛠️ Instalación

Para compilar y ejecutar Recon-Bot, asegúrate de tener instalado [Go](https://go.dev/dl/) (versión 1.24.3 o superior recomendada).

1.  **Clona el repositorio:**
    ```bash
    git clone <url-del-repositorio>
    cd recon-bot
    ```

2.  **Compila el binario:**
    ```bash
    go build -o recon-bot main.go
    ```

## 🚀 Uso

Ejecutar Recon-Bot es sencillo. Solo necesitas proporcionar el dominio objetivo mediante la bandera `-d`.

> **Nota:** El dominio debe ingresarse en formato limpio (ej: `ejemplo.com`). No incluyas `http://`, `https://` ni el prefijo `www.`.

```bash
./recon-bot -d google.com
```

### Ejemplo de salida:
```text
#########################################
#               RECON-BOT               #
#      Herramienta de Reconocimiento    #
#########################################
[+] Iniciando reconocimiento en: google.com

[*] Extrayendo información de DNS para: google.com
...
[*] Escaneando puertos comunes en la IP: 142.250.189.14
...
[*] Extrayendo cabeceras HTTP (Banner Grabbing) para: google.com
...
[*] Buscando subdominios pasivos en crt.sh para: google.com
...
```

## ⚖️ Aviso Legal (Disclaimer)

**ESTA HERRAMIENTA ES SOLO PARA FINES EDUCATIVOS Y AUDITORÍAS DE SEGURIDAD AUTORIZADAS.**

El uso de Recon-Bot contra objetivos sin el consentimiento previo por escrito es ilegal. Es responsabilidad del usuario final obedecer todas las leyes locales, estatales y federales aplicables. El autor no asume ninguna responsabilidad y no es responsable de ningún mal uso o daño causado por este programa. Al utilizar esta herramienta, aceptas usarla de manera ética y legal.
