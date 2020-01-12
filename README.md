# 97programmerbot

## Overview
This is a bot program that regularly tweeted "97 things programmers should know".<br>
[@97programmerbot](https://twitter.com/97programmerbot)

## Setup

### Dependencies
- Git
- Docker
- Twitter ( create your account, auth developer and get keys )

### Environment variables
You must apply to Twitter in advance for API usage. And set each acquired key as follows in `.env` . ( Please refer to the sample provided `.env.example` )

```.env
CONSUMER_KEY=
CONSUMER_SECRET=
ACCESS_TOKEN=
ACCESS_TOKEN_SECRET=
```

## Usage
```bash
# clone this repository
❯ git clone git@github.com:yoshimitsuEgashira/97programmerbot.git

# build
❯ docker buld --no-cache=true -t 97programmerbot:v1 .
...
Successfully built 1234abcd
Successfully tagged 97programmerbot:v1

# test
❯ docker run -it 97programmerbot:v1 go run test -v ./test

# run
❯ docker run -it 97programmerbot:v1 go run main.go
```
