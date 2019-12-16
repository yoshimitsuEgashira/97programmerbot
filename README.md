# 97programmerbot

## Overview
A bot program that regularly tweeted "97 things programmers should know".<br>
[@97programmerbot](https://twitter.com/97programmerbot)

## Setup

### Dependencies
- Docker version 19.03.5
- docker-compose version 1.24.1
- go version go1.12.14 linux/amd64
- Twitter API

### Environment variables
You must apply to Twitter in advance for API usage. And set each acquired key as follows in `.env` . ( Please refer to the sample provided `.env.example` )

```env
CONSUMER_KEY=
CONSUMER_SECRET=
ACCESS_TOKEN=
ACCESS_TOKEN_SECRET=
```

## Usage

### Test
You can test Go by executing the following command.
```
❯ docker-compose run app go test -v ./test
```

### Run
Execute the command for the `app` service defined in docker-compose.yml.

```
❯ docker-compose run app go run main.go
```
