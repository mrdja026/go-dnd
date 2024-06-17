# syntax=docker/dockerfile:1

FROM golang:1.22

WORKDIR /app

COPY ./ ./

copy main/main.go ./

RUN if [ -f go.mod ]; then go mod download; fi

COPY pkg .

RUN go build -o /godocker

EXPOSE 3333

CMD [ "/godocker" ]

