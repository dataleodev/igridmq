FROM golang:1.16.2-alpine3.13
MAINTAINER "Pius Alfred" "me.pius1102@gmail.com"
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go mod tidy
RUN cd cmd \
 && go build -o igridmq
EXPOSE 1883
EXPOSE 8080
CMD ["/app/cmd/igridmq"]