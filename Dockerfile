FROM golang:1.20
WORKDIR /go/src/maria-prototype
COPY . /go/src/maria-prototype/
RUN go mod tidy

# Airをインストール
RUN go install github.com/cosmtrek/air@v1.27.3

# airコマンドでGoファイルを起動
CMD ["air"]