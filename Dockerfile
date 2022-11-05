FROM golang:1.17.0

RUN mkdir /sbercloud
WORKDIR /sbercloud

COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

COPY /cmd/main.go cmd/
COPY /config/.env config/
COPY /doc/weather.swagger.json doc/
COPY /driver/*.go driver/
COPY /gclient/*.go gclient/
COPY /internal/entity/*.go internal/entity/
COPY /pb/ pb/
COPY /pkg/handler/*.go pkg/handler/
COPY /pkg/repository/*.go pkg/repository/
COPY /pkg/service/*.go pkg/service/
COPY /proto/google/api/*.proto proto/google/api/
COPY /proto/protoc-gen-openapiv2/options/*.proto proto/protoc-gen-openapiv2/options/
COPY /proto/*.go proto/
COPY data.json .
COPY Makefile .
COPY server.go .

RUN go build -o sbercloud cmd/main.go

EXPOSE 80

CMD ["./sbercloud"]