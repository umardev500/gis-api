FROM golang as dev

WORKDIR /app

COPY . .

EXPOSE 5000

CMD air

FROM golang as prod

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main

COPY main .

RUN ls -la

EXPOSE 5000

CMD [ "main" ]
