FROM golang:1.13.3 as build

LABEL maintainer="Damien GOUYETTE <damien.gouyette@gmail.com>"

WORKDIR /build

COPY src/* /build/
COPY go.mod . 
RUN go mod download

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -tags netgo -ldflags '-w' -o /build/proxy  .


FROM scratch
COPY --from=build /build /build

ENTRYPOINT ["/build/proxy"]
    