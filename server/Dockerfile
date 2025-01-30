FROM golang:latest

# Setting up working directory
WORKDIR /app 

# At first copy dependency files and download them
COPY go.mod go.sum ./ 

RUN go mod download && go mod verify 

# Copying the remaining code
COPY . ./ 

EXPOSE 8080

# Building the app and placing executable file 
# at /usr/local/bin/ as "goserver" executable
RUN go build -v -o /usr/local/bin/goserver ./...

CMD [ "goserver" ]