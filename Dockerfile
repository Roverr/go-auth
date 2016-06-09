FROM golang:1.6

ADD . /go/src/github.com/Roverr/go-auth
RUN cd /go/src/github.com/Roverr/go-auth && go get -d -v
RUN cd /go/src/github.com/Roverr/go-auth && go install -v
RUN chmod +x /go/src/github.com/Roverr/go-auth/scripts/wait-for-it.sh

# Run the wait-for-it-script which is waiting for MySQL container
# to start up and after that run go-auth application which now
# can listen to 3306 port.
ENTRYPOINT ["/go/src/github.com/Roverr/go-auth/scripts/wait-for-it.sh", "database:3306", "-s", "-t", "120", "--", "go-auth"]

# Document that the service listens on port 8080.
EXPOSE 8080
