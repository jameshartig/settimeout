FROM golang:latest
RUN mkdir /build
ADD settimeout.go /build
ADD assets /build/assets
WORKDIR /build
RUN go build -o settimeout .
CMD ["/build/settimeout"]
EXPOSE 51004
