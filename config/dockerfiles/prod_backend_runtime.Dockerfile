# Runtime del backend: sin Go, sin air y sin código fuente. El binario entra por
# volumen desde /srv/catalina-support (ver prod.yml).
FROM postgres:18-alpine

# El ejecutor usa las herramientas oficiales para crear y verificar el respaldo previo.
RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
