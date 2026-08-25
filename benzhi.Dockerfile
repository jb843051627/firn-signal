FROM golang:1.22-bookworm

WORKDIR /workspace
ENV GOPROXY=https://goproxy.cn,direct
ENV GOTOOLCHAIN=local

COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build ./...
RUN go build -o /usr/local/bin/firn-signal .
ENTRYPOINT ["/usr/local/bin/firn-signal"]
