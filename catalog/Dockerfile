FROM golang:latest 

WORKDIR /go/src
COPY . .
RUN GOOS=linux go build -ldflags="-s -w"
CMD [ "./catalog" ]
