FROM golang:1.12
COPY . /src
WORKDIR /src

RUN GOOS=linux go build -o bin/apt-grocery .

FROM heroku/heroku:18
WORKDIR /app
COPY --from=0 /src/bin/apt-grocery /app
COPY storage/list.json /app
CMD ["./apt-grocery"]
