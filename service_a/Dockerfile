FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

#COPY *.go ./
COPY . ./

RUN go build -o /service_a

EXPOSE 8062

CMD [ "/service_a" ]