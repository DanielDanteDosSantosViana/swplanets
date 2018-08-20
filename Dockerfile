
FROM golang

ADD  . /go/src/github.com/DanielDanteDosSantosViana/swplanets
WORKDIR /go/src/github.com/DanielDanteDosSantosViana/swplanets

RUN go install github.com/DanielDanteDosSantosViana/swplanets/cmd/swplanetsd

ENV PORT_ENV=8081

ENTRYPOINT /go/bin/swplanetsd

EXPOSE 8081