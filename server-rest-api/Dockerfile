FROM golang:alpine

MAINTAINER Rémy DEME <demeremy@gmail.com>


WORKDIR $GOPATH/src/server

# Install git.
# Git is required for fetching the dependencies.
RUN apk update && apk add --no-cache git

# Copy the current directory contents into the container at /app
COPY  . .

ADD . .

# Install any needed packages
RUN go get "github.com/dgrijalva/jwt-go" 
RUN go get "github.com/iris-contrib/middleware/jwt"
RUN go get "github.com/kataras/iris"   

# Make port 3000 available to the world outside this container
EXPOSE 3000

# Define environment variable
ENV NAME World

# Set proxy server, replace host:port with values for your servers
ENV http_proxy host:port
ENV https_proxy host:port


# Run app.py when the container launches
CMD ["go" , "run",  "server.go"]
