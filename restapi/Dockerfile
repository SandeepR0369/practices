FROM golang:1.12.9-alpine

RUN mkdir /restapi

#ENV SOURCES /go/src/restapi/

COPY . /restapi

RUN cd /restapi  && CGO_ENABLED=0 go build -o restapi .

#ENV PORT 8080
#EXPOSE 8080

#ENTRYPOINT restapi

CMD ["/restapi/restapi"]





