FROM golang:1.20
WORKDIR /go/src/intern-assignment
COPY . /go/src/intern-assignment
RUN go mod tidy

# Airをインストール
RUN go install github.com/cosmtrek/air@v1.27.3
