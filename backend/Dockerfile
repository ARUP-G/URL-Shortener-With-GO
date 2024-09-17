FROM golang:1.23-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main .

# Next stage 
FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/main .

EXPOSE 8181

CMD [ "./main" ]