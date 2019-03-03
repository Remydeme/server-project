#!/bin/bash 




echo "Build image for server-rest-api"

rv=$(docker build --tag=falconr/back-app server-rest-api/.)
echo "Steps 0 / 4"
if [ $rv -ne 0 ]
then
    echo "Failed to build back-app"
    return 
fi

echo "Pull image for mangodb"

echo "Steps 1 / 4"

rv=$(docker pull bitnami/mongodb:latest)

if [ $rv -ne 0 ]
then 
    echo "Failed to pull mangoDB "
    return 
fi

echo "Steps 2 / 4"

rv=$(docker build --tag=falconr/nginx-http2)


if [ $rv -ne 0 ] 
then 
    echo "Failed to build the nginx image"
    return 
fi

echo "Steps 3 / 4"

echo "Docker compose : Create the service"





