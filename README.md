
# GoBlog

A simple blogging application API built with Go.

## Features

- [x] User registration
- [x] JWT authentication
- [ ] Create and publish blog posts
- [ ] Clean and minimal design
- [ ] Fast performance


## Quick Setup

### Prerequisites

- Go 1.16 or higher
- Git

### Installation

1. Clone the repository:
```bash
git clone https://github.com/yourusername/goblog.git
cd goblog
```

2. Install dependencies:
```bash
go mod download
```

3. Build the application:
```bash
go build -o goblog
```

4. Run the application:
```bash
./goblog
```

The application will start on `http://localhost:5000`

### Docker Setup (Optional)

If you prefer to run the database using Docker:

```bash
docker-compose up -d
```

To stop the containers:
```bash
docker-compose down
```

To reset volumes:
```bash
docker-compose down -v
```

The application will be accessible at `http://localhost:5000`

## Project Structure

```
goblog/
├── main.go
├── handlers/
├── models/
├── templates/
└── README.md
```

## Configuration

Set environment variables as needed:

```bash
JWT_SECRET=supersecretkey

POSTGRES_HOST=localhost
POSTGRES_USER=bootrix
POSTGRES_PASSWORD=p@ssw0rd
POSTGRES_DB=goblog
POSTGRES_PORT=5432

SERVER_PORT=:5000
```


## Usage

Visit `http://localhost:5000` in your browser to access the blog.

## Contributing

Pull requests are welcome. For major changes, please open an issue first.

## License

MIT License

Copyright (c) 2024 Solanoize

Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated documentation files (the "Software"), to deal in the Software without restriction, including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all copies or substantial portions of the Software.




