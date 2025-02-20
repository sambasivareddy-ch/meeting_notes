# Meeting Notes

## Demonstration Video
[Application](https://drive.google.com/file/d/1b8haJ3x8g18WdpJb2K6QtBS1Qk0nQmLI/view?usp=drive_link)

## Overview
Meeting Notes is a full stack application designed to fetch all Google Meets/Hangouts from the Google API, display them, and allow users to take notes on them.

## Basic Information
1. User consents to use their email address, from which meetings/hangouts will be fetched.
2. A session is created and the user is stored in the database.
3. Using the Google Calendar API and the access token/userId provided by Google upon consent, meetings are fetched from the calendar API and stored in the meetings table.
4. All saved meetings are displayed to the user.
5. Users can add or edit notes, which are stored in the meeting notes table.

## Technologies Used
- React
- Go
- PostgreSQL
- Docker

## Directory Structure
```
meeting_notes
├── client          # React app created using "create-react-app"
├── server          # Go backend with go.mod & go.sum files
└── docker-compose.yml
```

## Installation

### Without Docker

#### Prerequisites
- Node.js and npm
- Go
- PostgreSQL

#### Steps
1. Clone the repository:
    ```bash
    git clone https://github.com/sambasivareddy-ch/meeting_notes.git
    cd meeting_notes
    ```
2. Set up the client:
    ```bash
    cd client
    npm install
    npm start
    ```
3. Set up the server:
    ```bash
    cd ../server
    go mod download
    go run main.go
    ```
4. Set up PostgreSQL:
    - Create a database and update the connection string in the server configuration.

### With Docker

#### Prerequisites
- Docker
- Docker Compose

#### Steps
1. Clone the repository:
    ```bash
    git clone https://github.com/sambasivareddy-ch/meeting_notes.git
    cd meeting_notes
    ```
2. Run Docker Compose:
    ```bash
    docker-compose up --build
    ```
