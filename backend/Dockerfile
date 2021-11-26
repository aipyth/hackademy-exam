FROM golang:1.16-alpine

WORKDIR /app
COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY *.go ./
RUN go build -o /backend

EXPOSE 8080

CMD ["/backend"]

##############################

# FROM golang:1.16 AS builder

# WORKDIR /app
# COPY . .

# ENV GO111MODULE=on

# RUN CGO_ENABLED=0 GOOS=linux go build -installsuffix cgo -o main .

##############################

# FROM scratch

# COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
# COPY --from=builder /app/main /

# CMD ["/main"]

