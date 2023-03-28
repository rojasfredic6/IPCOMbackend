# Prueba técnica para Backend Developer en IPCOM

Este repositorio contiene una prueba técnica para el puesto de Backend Developer en la empresa IPCOM. La prueba consiste en desarrollar una API RESTful en el lenguaje de programación Go utilizando el framework Echo.

## Requisitos

Para ejecutar la aplicación, necesitas tener instalado Go 1.16 o una versión posterior.

## Instalación y ejecución

1. Clona este repositorio en tu computadora.
2. Navega al directorio raíz del proyecto en tu terminal.
3. Instala dependencias
4. Ejecuta el comando `go run main.go` para iniciar el servidor web.
5. La aplicación estará disponible en `http://localhost:1323`.

## API

La aplicación implementa los siguientes endpoints:

- `GET /`: Devuelve un mensaje para verificar que se ha conectado a la api correctamente
- `GET /resume/2019-12-01`: Devuelve un array de transacciones realizadas en formato JSON, a partir de la fecha '2019-12-01'. Recibe como parametro query una propiedad 'dias' con valor de tipo string con la cantidad de dias consecuente a la fecha antes descrita.
- `POST /csv`: Devuelve un array de organizaciones con usuarios y roles en formato JSON a partir de un archivo que el usuario debe cargar con una propiedad 'file' que tiene la siguiente estructura.
  | organizacion | usuario | rol   |
  |--------------|---------|-------|
  | org1         | jperez  | admin | 
  | org1         | jperez  | superadmin | 
  | org1         | asosa  | writer | 
  | org2         | jperez  | admin | 
  | org2         | rrodriguez  | writer | 
  | org2         | rrodriguez  | editor | 
  
