FROM golang:buster as build

WORKDIR /app
COPY . .

RUN go build -o /hello-server server/main.go

FROM gcr.io/distroless/base-debian10

COPY --from=build /hello-server /

EXPOSE 5443 5444

CMD ["/hello-server"]