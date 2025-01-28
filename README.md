# Gotiny

Gotiny is a Url Shortner tool written in golang along with the generated Short URL analytics and frontend served using go templates.

## Overview

The URL Shortner shortens the passed URL based on a combination of `Base62 Encoding` and `MD5 hash algorithm` to reduce the generated shorturl collisions. When the short url is called, it redirects to the original long url. The tool also tracks how many times the shorturl is clicked. This tool uses `Postgres` as a persistent storage and `Redis` as a cache to store the url and analytics. it uses  `html/template` to render the frontend on server side.

## Table of Content

- [Features](#features)
- [Structure](#structure)
- [Environment Variables](#environment-variables)
- [Installation](#installation)
- [Usage](#usage)
- [Next steps](#next-steps)

## Features

- Generate shorturl using combination of Base62 and MD5 Hash algorithms.
- Track the generated short urls.
- Renders the frontend on server side.
- Developed completely using golang(backend and frontend).
- Uses Redis to serve the requests faster.
- Logs all the requests andd responses using `slog`.

## Structure

```bash
.
├── cmd             # Main Application(github.com/heisenberg8055/gotiny)
│
├── config          # Reads and Sets Evnironment variables from .env(github.com/joho/godotenv)
│
├── internal        # Internal Packages
│   │
│   ├── api         # Http Server (net/http)           
│   │ 
│   ├── log         # Logging client and utils (slog)
│   │   
│   ├── postgres    # Posgtres Client and utils (pgx)
│   │   
│   ├── redis-client # Redis Client and utils (redis-go official client)
│   │   
│   └── templates   # Go templates for server side rendering (html/templates)
│
├── static          # All static files for template rendering
│
└── .env            # Store all environment variables
```

## Environment Variables

```bash
POSTGRES_USER

POSTGRES_PASSWORD

POSTGRES_DATABASE

POSTGRES_PORT

POSTGRES_HOST

REDIS_URL
```

Installation

```bash
# Clone the repository
got clone https://github.com/heisenberg8055/gotiny.git

cd gotiny

# Install Dependencies
go mod download

cd cmd/gotiny

# Build
go build -o gotiny

# Run
./gotiny
```

## Usage

```bash

GET     /               # renders home

POST    /               # create shorturl

GET     /{shortURL}     # redirects shorturl

GET     /count          # fetches count of shorturl

```

## Next steps

- [ ] Currently the shortner doesn't supports users and authorized short urls.
- [ ] Add many more analytics options

