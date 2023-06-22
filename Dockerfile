FROM golang as dev

WORKDIR /app

COPY . .

EXPOSE 5000

CMD air