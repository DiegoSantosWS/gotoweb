FROM golang:alpine as build-stage
LABEL maintainer="diegosantosws1@gmail.com"
WORKDIR /go/src/github.com/DiegoSantosWS/gotoweb/projeto_api
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest
LABEL maintainer="diegosantosws1@gmail.com"
COPY --from=build-stage /go/src/github.com/DiegoSantosWS/gotoweb/projeto_api /
RUN apk --no-cache add ca-certificates && update-ca-certificates
CMD ["/main"]