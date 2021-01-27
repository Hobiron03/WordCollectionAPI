FROM golang:latest
RUN mkdir /go/src/wordcollection
WORKDIR /go/src/wordcollection

RUN apt-get update && apt-get install build-essential -y
ENV CGO_ENABLED=1 \
    GOOS=linux \
    GOARCH=amd64 \
    GO111MODULE=on
RUN go get -u github.com/dgrijalva/jwt-go \
    && go get -u github.com/gorilla/mux \
    && go get -u github.com/rs/cors \
    && go get -u github.com/google/uuid \
    && go get -u github.com/mattn/go-sqlite3 \
    && go get -u gopkg.in/go-ini/ini.v1

COPY . /go/src/wordcollection
CMD ["go", "run", "main.go" "$PORT"]
