# Go Image Simple Editor App

using boilerplate from : [gofiber-boilerplate](https://github.com/gofiber/boilerplate.git)

### Demo

https://github.com/randhipp/gise/assets/25372728/db3b564d-c313-46bc-ae5b-803c03c4d926

## Development

### Start the application 


```bash
go run app.go
```

### Use local container

```
# Shows all commands
make help

# Clean packages
make clean-packages

# Generate go.mod & go.sum files
make requirements

# Generate docker image
make build

# Generate docker image with no cache
make build-no-cache

# Run the projec in a local container
make up

# Run local container in background
make up-silent

# Run local container in background with prefork
make up-silent-prefork

# Stop container
make stop

# Start container
make start
```

## Production

```bash
docker build -t gofiber .
docker run -d -p 3000:3000 gofiber ./app -prod
```

Go to http://localhost:3000
