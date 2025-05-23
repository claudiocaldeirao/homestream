#!/bin/bash

# load .env
if [ -f .env ]; then
  echo "📄 Loading environment variables from .env..."
  export $(grep -v '^#' .env | xargs)
fi

# default values
CONTAINER_NAME=${MONGO_CONTAINER_NAME:-homestream-mongo}
PORT=${MONGO_PORT:-27017}
DB_NAME=${MONGODB_DATABASE:-homestreamdb}

if [ "$(docker ps -aq -f name=$CONTAINER_NAME)" ]; then
  echo "🛑 Stopping and removing existing container..."
  docker stop $CONTAINER_NAME
  docker rm $CONTAINER_NAME
fi

echo "🚀 Starting MongoDB container..."
docker run -d \
  --name $CONTAINER_NAME \
  -p $PORT:27017 \
  mongo:latest

echo "✅ MongoDB is running at $MONGODB_URI"
