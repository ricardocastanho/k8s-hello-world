FROM golang:1.19-alpine

WORKDIR /app

COPY ./main.go ./main.go
COPY ./go.mod ./go.mod

RUN ["go", "build", "-o", "gobernetes", "."]
RUN ["chmod", "+x", "./gobernetes"]

CMD ["./gobernetes"]
