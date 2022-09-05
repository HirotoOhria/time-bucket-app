FROM golang:1.19

# Don't use go get because /go mounted volume and hide behind.

# Attach /project_root/backend to /backend by docker compose.
# For run air with go.mod.
WORKDIR /backend/src

RUN cd /usr/local && \
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s

CMD ["air", "-c", "../.air.toml"]