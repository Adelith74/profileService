FROM golang:latest

WORKDIR /profileService

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN go build -o profile-service /profileService/cmd/main.go

RUN chmod +x profile-service

EXPOSE 8080:8080

CMD ["./profile-service"]