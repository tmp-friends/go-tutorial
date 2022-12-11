FROM golang:1.19-alpine

# アップデートとgitのインストール
RUN apk update && apk add git

WORKDIR /usr/src

