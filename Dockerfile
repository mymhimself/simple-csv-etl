FROM golang:latest as back-builder
COPY . /app
WORKDIR /app
RUN go build -o /bin/main ./main.go

FROM ubuntu:22.04 as host

RUN apt-get update \
  && apt-get install -y ca-certificates curl libxrender1 wget libfontconfig libxtst6 xz-utils unzip less  fontconfig libfontenc1 libjpeg-turbo8 xfonts-75dpi xfonts-base xfonts-encodings xfonts-utils
 
COPY --from=back-builder /bin/main /

WORKDIR /


