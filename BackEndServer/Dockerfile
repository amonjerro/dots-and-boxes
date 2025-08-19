FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download
COPY . ./
RUN ls

RUN CGO_ENABLED=0 GOOS=linux go build -v -o /dab-server 

CMD ["/dab-server"]