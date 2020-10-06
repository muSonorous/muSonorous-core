# muSonorous core music APIs
- This service contains core audio transformation components and helper methods

### To build Dockerfile
Execute the following command to build and start containarized service

```
docker build --tag muSonorousCore:1.0 . up
```

### Build without docker
#### To build without docker use 
```
go build
```

### To run the Microservice 
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