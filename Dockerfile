FROM golang
ADD . /go/src/github.com/skim1420/k8ssimple
RUN go install github.com/skim1420/k8ssimple
ENTRYPOINT /go/bin/k8ssimple
