# Use the official go docker image built on debian.
FROM golang:1.7

WORKDIR /go/src/github.com/rewiko/app/

# reload code
RUN go get github.com/codegangsta/gin

RUN go get github.com/gin-gonic/gin
RUN go get github.com/tools/godep
RUN go get github.com/spf13/viper 
RUN go get github.com/Sirupsen/logrus
RUN go get gopkg.in/mgo.v2
RUN go get github.com/gedex/inflector
RUN go get github.com/jinzhu/gorm
RUN go get github.com/mattn/go-sqlite3
RUN go get github.com/manyminds/api2go/jsonapi 
RUN go get github.com/julienschmidt/httprouter
RUN go get github.com/gocql/gocql

ENTRYPOINT gin -a 8081 -p 8080

# Grab the source code and add it to the workspace.
ADD ./src/ /go/src/github.com/rewiko/app

#RUN godep restore

# Open up the port where the app is running.
EXPOSE 8080
