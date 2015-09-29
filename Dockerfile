FROM golang:1.5
ADD ./dockerprog.go /working/dockerprog.go
RUN mkdir /app
COPY stats /app/.
ENTRYPOINT ["go", "run", "/working/dockerprog.go"]

