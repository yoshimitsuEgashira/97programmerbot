# 97programmerbot

![](https://img.shields.io/github/go-mod/go-version/yossiee/97programmerbot?style=plastic)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
![](https://img.shields.io/twitter/url?label=follow%20me%21&style=social&url=https%3A%2F%2Ftwitter.com%2F97programmerbot)

## Overview
This is a bot program that regularly tweeted "97 things programmers should know".

[97 programmer bot \(@97programmerbot\) / Twitter](https://twitter.com/97programmerbot)

## Setup
### Dependencies
- Git
- Docker
- Twitter ( create your account, auth developer and get keys )

### Environment variables
You must apply to Twitter in advance for API usage. And set each acquired key as follows in `.env` . ( Please refer to the sample provided `.env.example` )

```sh
CONSUMER_KEY=
CONSUMER_SECRET=
ACCESS_TOKEN=
ACCESS_TOKEN_SECRET=
```

## Usage
```sh
# Clone this repository
$ git clone git@github.com:yossiee/97programmerbot.git

$ cd 97programmerbot && cp .env.example .env

# Build
$ docker buld --no-cache=true -t 97programmerbot:v1 .

# Test
$ docker run -it 97programmerbot:v1 go run test -v ./test

# Tweet
$ docker run -it 97programmerbot:v1 go run main.go
```
