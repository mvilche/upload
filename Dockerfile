FROM golang:1.16.4-alpine3.13
WORKDIR /app
RUN go get github.com/revel/revel && go get github.com/revel/cmd/revel
COPY . /app
RUN revel clean . && revel build -v -a /app -t /tmp/app
EXPOSE 9000
CMD revel run -m prod