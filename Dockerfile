FROM golang:1.5
ADD ./dockerprog.go /working/dockerprog.go

ENTRYPOINT ["go", "run", "/working/dockerprog.go"]

