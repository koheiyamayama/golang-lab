FROM golang:1.23.6
COPY go.mod ./
RUN go mod download
COPY . .
CMD [ "go", "run", "." ]

