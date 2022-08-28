FROM golang:1.19

# Don't use go get because /go mounted volume and hide behind.

# Will attach project_root/src/.
WORKDIR /work

RUN cd /usr/local && \
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

RUN go install github.com/schemalex/schemalex/cmd/schemadiff@latest && \
  go install github.com/schemalex/schemalex/cmd/schemalex@latest && \
  go install github.com/schemalex/schemalex/cmd/schemalint@latest

CMD ["air"]