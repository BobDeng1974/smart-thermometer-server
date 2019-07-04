Smart thermometer server
=============================

This is part of the wather predictor project that consists in a smart thermometer, a server and a dashboard

## Description
This is a go server that is meant to run with as little resources as possible, it will have a file database, since it's meant as a DIY project.

## Commands
it needs go >= 1.12


bootstrap the application (set dependency)

```sh
.ci/bootstrap
```

run
```sh
ENVIRONMENT=lololol SUSERNAME=lol SPASSWORD=lol go run main.go
```
where SUSERNAME and SPASSWORD are environment variables for the basic auth


## Deploy with Docker compose
You will need to have docker and docker compose

You will need to create a .env.production file, you can use mock-env as a reference
You will need also to create a /data on your home folder
```sh
mkdir $HOME/data
```

in order to start the server in production you will need to run the following code

```sh
sudo docker-compose -f docker-compose-production.yml up -d
```

## Licence
https://www.mozilla.org/media/MPL/2.0/index.815ca599c9df.txt