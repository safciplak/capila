# Specify hash to prevent hijacking attacks
FROM golang:1.21.5-alpine as development

# The build will be done in the build folder which isn't present in the root image
WORKDIR /app

# This is done so that if the mod or sum file are unchanged the docker cache is used to build the rest of your application
COPY ./go.mod ./go.sum /app/
RUN go mod download

# Copy the source of the application into the docker, ignoring everything else
COPY ./ /app