#FROM debian:latest AS build
#WORKDIR src
#COPY . .
#RUN apt-get update -y && apt-get -y install git golang libzmq3-dev
#RUN go build -o /dist/ov-geosearch .
##FROM scratch AS bin
##COPY --from=build /dist/ov-geosearch .
##WORKDIR /dist
#RUN chmod +x /dist/ov-geosearch
#CMD ['/dist/ov-geosearch']

FROM debian:latest AS build
RUN apt-get update -y && apt-get -y install git golang libzmq3-dev
RUN git clone https://github.com/Gerrist/ov-geosearch
WORKDIR /ov-geosearch
CMD go get
CMD go run .