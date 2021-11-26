FROM golang:1.16-alpine
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY *.go ./
RUN go build -o /instances_of_and
EXPOSE 8000
CMD ["/instances_of_and"]