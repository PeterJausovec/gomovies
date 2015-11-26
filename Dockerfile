FROM golang:1.4
COPY . /go/src/github.com/pj3677/gomovies
RUN go get -v github.com/pj3677/gomovies
ENTRYPOINT [ "gomovies" ]