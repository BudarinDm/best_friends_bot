FROM golang:1.20

WORKDIR /src
COPY . .

RUN go mod tidy
RUN go mod vendor
RUN CGO_ENABLED=0 GOOS=linux go build -mod=vendor -a -installsuffix cgo -o app -ldflags "-X 'main.version=${VERSION}'" ./cmd/main.go

EXPOSE 4000

CMD ["./app"]
