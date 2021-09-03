FROM golang:1.16-alpine

WORKDIR /app

RUN go mod init github.com/gospodinzerkalo/crime_city_api

RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["make","rest"]