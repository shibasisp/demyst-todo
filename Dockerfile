FROM golang:latest

COPY . .

RUN make build

ENV FLAGS=""

CMD ["make", "run"]
