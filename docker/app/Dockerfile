FROM golang:1.15

WORKDIR /go/src/app
COPY . .
ENV GOPATH $GOPATH:/go
ENV POSTGRESQL_URL $POSTGRESQL_URL:postgres://postgres:root@postgres:5432/skilldb?sslmode=disable

RUN go get -d -v ./...
RUN go get -tags 'postgres' -u github.com/golang-migrate/migrate
RUN go get github.com/go-delve/delve/cmd/dlv
RUN go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv
#RUN go build -gcflags "all=-N -l" -o /server ./server.go
#RUN GO111MODULE=off go get github.com/oxequa/realize

#CMD ["dlv", "--listen=:2345", "--headless=true", "--api-version=2", "exec", "/server"]
#CMD ["realize","start"]
CMD ["app"]