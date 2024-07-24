FROM golang:1.22-alpine as base
RUN apk --no-cache update

FROM base as ci
WORKDIR /app/
COPY . .
RUN go mod tidy

FROM ci as build
WORKDIR /app/
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o entrypoint

FROM scratch
WORKDIR /
COPY --from=base /usr/local/share/ca-certificates /usr/local/share/ca-certificates
COPY --from=base /etc/ssl/certs /etc/ssl/certs/
COPY --from=builder /app/entrypoint .
COPY --from=builder /app/rev.txt .

ENTRYPOINT [ "/entrypoint" ]
