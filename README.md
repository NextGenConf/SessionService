# Session Service

The session service provides an endpoint to store and retreive sessions.

## Development

Prerequisites:

* Go 1.12 minimum
* Docker

The session service uses [MongoDB](https://www.mongodb.com) to store
the sessions. To develop the session service, a MongoDB server should be
running.

Run these commands:

```sh
DOCKER_NETWORK="session-dev"

MONGO_SERVER_CONTAINER="mongo-server"

# Create a docker network
docker network create $DOCKER_NETWORK

# Start the mongo server and connect it to the previously created network and open on port 27017 on host
docker run -d --network $DOCKER_NETWORK --name $MONGO_SERVER_CONTAINER -p 27017:27017 mongo
```

This will run a mongo server container called `mongo-server` listening at [mongodb://localhost:27017](mongodb://localhost:27017)
atached to the `session-dev` docker bridge network.

To populate the mongo server with test data, assuming that the environment
variables `DOCKER_NETWORK=session-dev` and `MONGO_SERVER_CONTAINER=mongo-server`,
and that the current working directory is at the root of the `SessionService` repo,
run these commands:

```sh
MONGO_CLIENT_CONTAINER="mongo-client"

docker run --rm --network $DOCKER_NETWORK --name $MONGO_CLIENT_CONTAINER -v `pwd`/testdata:/testdata mongo \
    mongoimport \
    --host $MONGO_SERVER_CONTAINER \
    --db sessions \
    --collection sessions \
    --drop \
    --file /testdata/sessions.json
```

To check that MongDB has been populated by the fake data, run:

```sh
docker run --rm --network $DOCKER_NETWORK --name $MONGO_CLIENT_CONTAINER -it mongo \
    mongo --host $MONGO_SERVER_CONTAINER sessions
```

This should take you to the Mongo Shell. Run:

```sh
> db.sessions.find();
```

You should see this output:

```sh
> db.sessions.find();
{ "_id" : ObjectId("5dfe7e24baecce7d224fb7a3"), "uniqueName" : "my-session",
"title" : "My Session", "subtitle" : "I'm gonna talk about stuff",
"description" : "There  is some stuff that I'm going to talk about", "presenterId" : "me123",
"slideDeckUrl" : "https://stuff" }
{ "_id" : ObjectId("5dfe7e24baecce7d224fb7a4"), "uniqueName" : "my-session-2",
"title" : "My Session 2", "subtitle" : "I'm gonna talk about more stuff",
"description" : "There is more stuff that I'm going to talk about", "presenterId" : "me123",
"slideDeckUrl" : "https://more-stuff" }
>
```

To run the session service, run

```sh
go run main.go
```
