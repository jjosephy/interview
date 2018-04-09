#!/bin/bash

cd $1
[ -e "./cert.pem" ] && rm "cert.pem" || echo "cert.pem doesnt exist"
[ -e "./private_key" ] && rm "private_key" || echo "private_key doesnt exist"
[ -d "logs/" ] && rm -r "logs/" || echo "logs/ directory doesnt exist"
[ -d "config/" ] && rm -r "config/" || echo "config/ directory doesnt exist"