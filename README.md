# 97programmer

[![Go](https://img.shields.io/github/go-mod/go-version/yossiee/97programmerbot?style=plastic)](go.mod)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat)](LICENSE)
[![](https://img.shields.io/docker/stars/yossiee/97programmer?style=plastic)](https://hub.docker.com/r/yossiee/97programmer)
[![](https://img.shields.io/docker/pulls/yossiee/97programmer?style=plastic)](https://hub.docker.com/r/yossiee/97programmer)
[![](https://img.shields.io/docker/cloud/build/yossiee/97programmer?style=plastic)](https://hub.docker.com/r/yossiee/97programmer)
[![Docker Cloud Automated build](https://img.shields.io/docker/cloud/automated/yossiee/97programmer?style=plastic)](https://hub.docker.com/r/yossiee/97programmer)
![build](https://github.com/yossiee/97programmer/workflows/build/badge.svg)
[![@97programmerbot](https://img.shields.io/twitter/follow/97programmerbot?label=follow%20me&style=social)](https://twitter.com/97programmerbot)

## Overview
This is a bot regularly tweeted "97 things programmers should know".

![511RPej0BNL _SY291_BO1,204,203,200_QL40_ML2_](https://user-images.githubusercontent.com/38056766/119687988-9339d980-be82-11eb-87f7-a4286b2b1d00.jpg)

[97 programmer bot \(@97programmerbot\) / Twitter](https://twitter.com/97programmerbot)

## Setup
### Dependencies
| software | description |
| :---: | --- |
| _Docker_ | <li>If you want to manage Docker image in the registry like me, create an account for the [docker hub](https://hub.docker.com).</li><li>By the way, this is [the docker image](https://hub.docker.com/r/yossiee/97programmer) of this project.</li> |
| _Twitter_ | <li>Visit [https://twitter.com/signup](https://twitter.com/signup) to create an account.</li><li>Register developer account and get API keys from [https://developer.twitter.com](https://developer.twitter.com)</li> |

### Environment variables
You must apply to Twitter in advance for API usage. And set each acquired key as follows in `.env`. ( Please refer to the sample provided `.env.example`. )

```.env
CONSUMER_KEY=
CONSUMER_SECRET=
ACCESS_TOKEN=
ACCESS_TOKEN_SECRET=
```

## Usage
```sh
# Clone this repository
$ git clone git@github.com:yossiee/97programmer.git

$ cd 97programmer && cp .env.example .env

# Build
$ docker build --no-cache=true -t 97programmer:latest .

# Test
$ docker run -it --env-file=.env 97programmer:latest go test -v -cover ./...

# Tweet
$ docker run -it --env-file=.env 97programmer:latest
```

## License

MIT
