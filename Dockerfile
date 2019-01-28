# For demo don't using in production
FROM golang:1.9.2-alpine3.6
RUN mkdir /go/src/project/
WORKDIR /go/src/project/

ADD . /go/src/project/
# RUN go mod download

CMD ["go", "run", "main.go"]
# EXPOSE 1324