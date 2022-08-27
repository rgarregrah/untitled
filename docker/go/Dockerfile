FROM golang:1.19-bullseye as builder
ENV CGO_ENABLED=0
COPY ./backend /app
WORKDIR /app
RUN go build -ldflags="-w -s" -o ./build ./main.go

FROM scratch as prod
COPY --from=builder /app/build /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /usr/share/zoneinfo/Asia/Tokyo /usr/share/zoneinfo/Asia/Tokyo
CMD ["/app"]