# Utiliza una imagen base que incluya Go
FROM golang:latest

# Establece el directorio de trabajo dentro de la imagen
WORKDIR /go/src/app

# Copia el código fuente de tu proyecto al directorio de trabajo en la imagen
COPY . .

# Instala las dependencias de Go
RUN go get -d -v ./...

# Instala las dependencias de Python
RUN apt-get update && apt-get install -y python3 python3-pip
COPY requirements.txt .
RUN pip3 install -r requirements.txt

# Expone el puerto si tu proyecto necesita escuchar en un puerto específico
# EXPOSE 8080

# Comando por defecto para ejecutar tu proyecto Go
CMD go run src/*.go

