FROM golang:1.22.2-alpine


RUN go version
ENV GOPATH=/

COPY ./ ./

RUN go mod download
RUN go build -o denet-api ./cmd/app/main.go

CMD [ "./denet-api" ]