# GO GORM REST API

API Rest que trabaja con PostgreSQL para manejar una lista de usuarios con sus respectivas tareas. Es un proyecto que sirve como ejemplo sobre cómo crear una API con Go, Gorilla/Mux y Gorm ORM.

## Cómo se preparó el proyecto para el desarrollo

1. Ejecutar `go mod init github.com/FranciscoJSB12/go-gorm-rest-api`, es importante que sea una URL única, por lo general se escriben las URL's de Github como se muestra en el ejemplo.

2. Para la creación del servidor se puede usar el módulo de Go `net/http`, Frameworks de Go o módulos que están encima de `net/http`, siendo el caso de `gorilla/mux`, el cual se instala con el comando `go get -u github.com/gorilla/mux`

3. Instalar un recargador en vivo para aplicaciones de go, para este proyecto se usó `air` se puede encontar en google buscando `golang air`. Con el siguiente comando se instala `go install github.com/air-verse/air@latest`. Es importante crear un alias si se está trabajando con linux para poder usar el paquete, se puede hacer `alias air='$(go env GOPATH)/bin/air'`.

4. Correr `air init` para crear el archivo de configuración de air y luego el comando `air` para encender el servidor y que este se reinice automáticamente según los cambios.

5. Instalar `GORM`, el cual es un ORM para manejar las consultas SQL. Para ello usar el comando `go get -u gorm.io/gorm`.

6. También se debe instalar el módulo para Postgres con `go get -u gorm.io/driver/postgres`.
