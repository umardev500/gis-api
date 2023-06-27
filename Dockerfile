FROM golang as dev

WORKDIR /app

COPY . .

EXPOSE 5000

CMD air

FROM golang as prod

WORKDIR /app

COPY . .

RUN go mod tidy

RUN CGO_ENABLED=0 go build -o main

EXPOSE 5000

CMD [ "./main" ]
