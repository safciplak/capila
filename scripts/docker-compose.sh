#!/usr/bin/env bash

# Spins the builder image up with the application, ignoring most the the landscape
docker-compose -f deployment/docker-compose.yml up --build "capila-watcher"
