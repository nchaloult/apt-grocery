FROM golang:1.13
COPY . /src
WORKDIR /src

RUN GOOS=linux go build -o bin/apt-grocery .

FROM heroku/heroku:18
WORKDIR /app
COPY --from=0 /src/bin/apt-grocery /app
COPY storage/list.json /app
COPY storage/prices.json /app
CMD ["./apt-grocery"]
