FROM golang



# Copy the local package files to the container's workspace.
ADD . /go/src/github.com/jjosephy/go/interview

RUN apt-get update
RUN apt-get install -y libldap2-dev
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/mqu/openldap
RUN go get labix.org/v2/mgo
RUN go get labix.org/v2/mgo/bson

RUN go install github.com/jjosephy/go/interview
RUN openssl genrsa -out /go/bin/private_key 2048
RUN openssl req -new -x509 -key /go/bin/private_key -out /go/bin/cert.pem -days 365 -subj "/C=US/ST=Washington/L=SeattleO=Nord/CN=nord.interivew.com"

# Run the outyet command by default when the container starts.
ENTRYPOINT /go/bin/interview

# Document that the service listens on port 8080.
EXPOSE 8080

#FROM busybox
#ENV foo /bar
#WORKDIR ${foo}   # WORKDIR /bar
#ADD . $foo       # ADD . /bar
#COPY \$foo /quux # COPY $foo /quux
