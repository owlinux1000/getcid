FROM golang:latest
LABEL maintainer="Chihiro Hasegawa <encry1024@gmail.com>"
ADD main.go /go/src
ENTRYPOINT ["go"]
CMD ["run", "src/main.go"]
