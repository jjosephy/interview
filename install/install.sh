#!/bin/bash

path=$1
run=$2
cd $path

if [ $run -eq 1 ] 
then
    openssl genrsa -out private_key 2048
    openssl req -new -x509 -key private_key -out cert.pem -days 365 -subj "//C=US\ST=Washington\L=Seattle\O=Comp\CN=localhost"
    [ ! -d logs/ ] && mkdir logs/
    [ ! -d config/ ] && mkdir config/
    cp $GOPATH\\src\\github.com\\jjosephy\\interview\\config\\env.json config/env.json
else 
    echo 'running tests'
fi