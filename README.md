# High level overview about the idea (μSonorous)

- We all are aware about apps like StarMaker and Smule. We can sing any song with track and also get to know about the score based on how much we are in key. But these app lacks features like a good instant messaging, creating own track with inbuilt online music editor which contains rich features like Audacity (with free tempo and pitch correction), live Jamming sessions, creating own track, and cloud storage like soundcloud and so on.

- We’ll create a music creation platform for everybody where people will be able to search music, sing with karaoke, create own track, jam with others including features like instant messaging, cloud storage, .. .sounds very basic right ? Well well well! there is something more


# muSonorous core music APIs
- This service contains core audio transformation components and helper methods

## Instructions for running without Container Orchestrator
First move to the `core` directory which contains all the files of our service.
```
cd core
```
### To build Dockerfile
Execute the following command to build and start containarized service
- use ```docker build --tag muSonorousCore:1.0 . up```


### Build without docker
#### To build without docker use 
```
go build
```

#### To run the Microservice 
For UNIX Based OS
```
./core
```

For WINDOWS OS
```
core.exe
```
### Requirements without Docker 
You need atleast one instance of `MongoDB` and `Redis`. You can use `Docker` to run these.
If you don't want to install with `Docker` then install `Redis` and `MongoDB` manually. 

#### Running MongoDB using docker
```
docker run --name mongo -d mongo
```

#### Running Redis using docker
```
docker run --name redis -d redis
```

#### (Note : Don't forget to update `.env.yaml` file file with `hostname`, `username` and `password`)
#### Updating `.env.yaml` is not required if you are running the service using `Dockerfile`