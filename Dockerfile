FROM golang:1.10.0

ENV PKG github.com/yuuki/mftracer
WORKDIR /go/src/$PKG
COPY . /go/src/$PKG
RUN make deps build
RUN go install $PKG/cmd/...