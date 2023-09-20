FROM golang:1.19 AS build

WORKDIR /app
RUN go env -w GO111MODULE=on
RUN go env -w GOPROXY=https://goproxy.cn,direct


COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN go build -mod=mod --tags netgo -v -o app .
#CMD ["app"]

FROM alpine:3.16
WORKDIR "/app"
EXPOSE 8080

COPY --from=build app .
CMD ["./app"]
