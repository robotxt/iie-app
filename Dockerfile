FROM golang:1.16

WORKDIR /go/src/app

COPY . .

# Copy firebase Credential
RUN mkdir /opt/firebase_cred
COPY ./src/repo/firebase/credentials.json /opt/firebase_cred/.

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Download CompileDaemon for auto reload when saved
RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon --build="go build -o executable" --command=./executable
