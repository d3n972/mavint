FROM golang:1-alpine as build
RUN mkdir /app
WORKDIR /app
COPY go.* /app/
RUN go mod download
COPY ./ /app/
RUN go build -o /usr/bin/appsrv -buildvcs=true ./

FROM alpine
ENV GIN_MODE=release
RUN apk add tzdata
COPY --from=build /usr/bin/appsrv /bin
CMD ["/bin/appsrv"]