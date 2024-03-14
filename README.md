## Run Locally

1. Clone the project

```bash
  git clone https://link-to-project
```

2. Go to the project directory

```bash
  cd my-project
```

3. Copy `.env.example` to `.env`

```bash
  cp .env.example .env
```

4. Add `publics/documents` directory

```bash
  mkdir publics/documents
```

5. Install dependencies

```bash
  go mod tidy
```

6. Start the server

```bash
  go run main.go
```

## Using Docker Compose Watch

1. Build the project

```bash
  docker compose up --build -d
```

2. Start compose watch

```bash
  docker compose watch
```
