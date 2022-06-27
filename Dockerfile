FROM golang:alpine as build

WORKDIR /app
COPY . .

RUN go build -o /app/hello-server server/main.go

FROM alpine

WORKDIR /app
COPY --from=build /app/hello-server /app/hello-server
COPY step-hello-run.sh /app/step-hello-run.sh
RUN chmod a+rx /app/step-hello-run.sh
RUN sh -c 'echo http://dl-cdn.alpinelinux.org/alpine/edge/testing >>  /etc/apk/repositories'
RUN apk add step-cli step-certificates

EXPOSE 5443 5444

#CMD ["/app/hello-server"]
CMD ["/app/step-hello-run.sh"]