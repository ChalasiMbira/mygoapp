# Golang Full-Stack App

A simple full-stack application built with **Golang** for the backend and a static **HTML/JavaScript** frontend. The backend serves a REST API, and the frontend fetches and displays data from it. Ideal for junior developers learning Go and full-stack development.

## Features
- **Backend**: Golang HTTP server with a GET '/api/greet' endpoint that returns a JSON greeting.
- **Frontend**: Static HTML page that fetches and displays the greeting using JavaScript.
- **No Dependencies**: Uses Go's standard library and vanilla JavaScript.

## Project Structure

myapp/
├── main.go               # Backend server code
├── static/
│   └── index.html        # Frontend HTML page
├── .gitignore            # Git ignore file
├── README.md             # Project documentation


## Prerequisites
- [Go](https://go.dev/doc/install) (verify with 'go version')
- [Git](https://git-scm.com/downloads) (verify with 'git --version')
- A web browser

## Setup
1. Clone the repository:
      bash
   git clone https://github.com/ChalasiMbira/mygoapp.git
   cd myapp