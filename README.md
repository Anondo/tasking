# Tasking

A golang application where tasks are published to a queue(rabbitmq) & they are
consumed by worker(s)

## Tasks

1. ```add``` ```returns the sum of the two arguments provided```
1. ```div``` ```returns the quotient of the two arguments provided```
1. ```sub``` ```returns the difference of the two arguments provided```
1. ```mul``` ```returns the product of the two arguments provided```

## Running the app

Go to the terminal &

1. ```docker-compose up -d```
Starting & running the docker containers required

1. ```./publish add 12 5```
publishing the add task with arguments 12 & 5

1. ```./consume```
starts the worker making it ready to consume tasks from the queue

Its better to run the ```publish``` & ```consume``` shell scripts from two separate
terminal windows/tabs to get a better view.

I have used the machinery package for asynchronous task queue/job queue based on distributed message passing.
Get it by ```go get github.com/RichardKnop/machinery/v1```.

## Requirements

1. Docker / Docker compose

or if you want to run everything locally

1. ```Rabbitmq``` ```as the broker```
1. ```Redis``` ```as the backend result storage```
