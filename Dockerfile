ARG GO_VERSION=1.21.6

FROM golang:${GO_VERSION}-alpine AS build

# Definimos que no hay un modulo para instalar las dependencias y que vaya directo a GOPROXY
RUN go env -w GOPROXY=direct

# Instalar Git
RUN apk add --no-cache git

# Instalar certificados de seguridad para que el servidor pueda correr
RUN apk --no-cache add ca-certificates && update-ca-certificates


# Directorio donde se ejecutaran los comandos
WORKDIR /src

# Copiar moludos al directorio /src
COPY ./go.mod ./go.sum ./

# Descargar los modulos
RUN go mod download

COPY ./ ./

# -installsuffix 'static' -> para que el ejecutable funcione en el contenedor
# -o /omar-rest-ws -> como se va a llamar el archivo
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o /omar-rest-ws

# Otra imagen encargada de ejecutar la aplicacion
FROM scratch AS runner

# Copiamos los certificados creados en la imagen anterior
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

COPY .env ./

# Copiamos el ejecutable desarrollado en la linea anterior
COPY --from=builder /omar-rest-ws /omar-rest-ws

EXPOSE 5050

# Especificamos el comando que vamos a ejecutar
ENTRYPOINT [ "/omar-rest-ws" ]