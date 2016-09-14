FROM golang:1.6

COPY . /go/src/github.com/nacgarg/telegram-go-markov-bot/
WORKDIR /go/src/github.com/nacgarg/telegram-go-markov-bot/

RUN ls
RUN go build .

ENTRYPOINT ["/go/src/github.com/nacgarg/telegram-go-markov-bot/telegram-go-markov-bot"]
CMD ["--help"]