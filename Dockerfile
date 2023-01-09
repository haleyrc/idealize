FROM golang:1.19-alpine as builder

WORKDIR /app

RUN apk -U upgrade && \
    apk add --no-cache git

COPY . .

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go env -w "GOPRIVATE=github.com/haleyrc/*"
RUN go build -ldflags="-w -s" ./cmd/idealize

FROM scratch

EXPOSE 8080

COPY --from=builder /app/idealize /usr/bin/

ENTRYPOINT ["idealize"]