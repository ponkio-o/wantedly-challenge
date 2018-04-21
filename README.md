# WantedlyChallengeWebApp
”Hello world”というメッセージがJSON形式で返ってくるAPI

# Setup
## Install Docker

```sh
$ brew cask install docker
$ open /Applications/Docker.app
```

## Download&Build

```sh
$ git clone https://github.com/ponkio-o/wantedly-challenge.git
$ cd wantedly-challenge
$ docker build -t api-server .
```

# Run
```sh
$ docker run -p 8080:9000 [REPOSITORY]:[TAG] or [IMAGE]
```

# Usage
Request
```sh
$ curl -XGET -H 'Content-Type:application/json' http://localhost:8080/
```

Response
```sh
{"message":"Hello Wolrd!!"}
```

# Stop
```sh
$ docker ps
CONTAINER ID        IMAGE               COMMAND                  CREATED             STATUS              PORTS                    NAMES
dad40377f81a        api-server:latest   "/bin/sh -c 'go run …"   2 minutes ago       Up 2 minutes        0.0.0.0:8080->9000/tcp   unruffled_beaver
$ docker stop [CONTAINER ID]
```

# Heroku

Request
```sh
$ curl -XGET -H 'Content-Type:application/json' http://wantedly-webapp.herokuapp.com/
```

Response
```sh
{"message":"Hello Wolrd!!"}
```
