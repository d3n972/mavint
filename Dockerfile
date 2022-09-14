FROM golang:1-alpine as build
RUN mkdir /app
WORKDIR /app
COPY go.* /app/
RUN go mod download
COPY ./ /app/
RUN go build -o /usr/bin/appsrv ./

FROM alpine
ENV GIN_MODE=release
COPY --from=build /usr/bin/appsrv /bin
CMD ["/bin/appsrv"]