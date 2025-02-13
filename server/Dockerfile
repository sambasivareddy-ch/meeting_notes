FROM golang:latest AS builder

# Setting up working directory
WORKDIR /app 

# At first copy dependency files and download them
COPY go.mod go.sum ./ 

RUN go mod download && go mod verify 

# Copying the remaining code
COPY . ./ 

EXPOSE 8080

CMD [ "go", "run", "." ]