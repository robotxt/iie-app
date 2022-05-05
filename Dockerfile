FROM golang:1.18

WORKDIR /go/src/app

COPY . .

COPY cmd/init-user-db.sh /docker-entrypoint-initdb.d/init-user-db.sh

# Copy firebase Credential
RUN mkdir /opt/firebase_cred
COPY ./src/repo/firebase/credentials.json /opt/firebase_cred/.

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Download CompileDaemon for auto reload when saved
RUN go install github.com/githubnemo/CompileDaemon@latest

ENTRYPOINT CompileDaemon --build="go build -o executable" --command=./executable
