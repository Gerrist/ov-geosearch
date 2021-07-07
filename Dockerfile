#FROM golang:1.14 as builder
#
#RUN mkdir /src
#WORKDIR /src
#RUN go get
#RUN CGO_ENABLED=0 go build -o /src/ov-geosearch
#
#FROM alpine
#COPY --from=builder /src/ovgs /app/ovgs
#WORKDIR /
#CMD ["/app/ovgs"]


FROM debian:latest AS build
WORKDIR src
COPY . .
RUN apt-get update -y && apt-get -y install git golang libzmq3-dev
RUN go build -o /dist/ov-geosearch .
#FROM scratch AS bin
#COPY --from=build /dist/ov-geosearch .
#WORKDIR /dist
RUN chmod +x /dist/ov-geosearch
CMD ['/dist/ov-geosearch']