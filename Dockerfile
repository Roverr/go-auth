FROM golang:1.6

ADD . /go/src/go-auth
RUN cd /go/src/go-auth && go get -d -v
RUN cd /go/src/go-auth && go install -v
RUN chmod +x /go/src/go-auth/scripts/wait-for-it.sh

# Run the go-auth command by default when the container starts.
ENTRYPOINT ["/go/src/go-auth/scripts/wait-for-it.sh", "database:3306", "-s", "-t", "60", "--", "go-auth"]

# Document that the service listens on port 8080.
EXPOSE 8080
