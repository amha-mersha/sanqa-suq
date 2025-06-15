# Sanqa Suq - E-commerce API

Sanqa Suq is a Go REST API for an e-commerce platform, handling users, products, and categories with JWT authentication and PostgreSQL.

## Quick Start
1. **Clone the Repository**:
   ```bash
   git clone https://github.com/amha-mersha/sanqa-suq.git
   cd sanqa-suq
   ```

2. **Set Up Environment Variables**:
   - Copy `example.env` to `.env` and update values if needed:
     ```bash
     cp example.env .env
     ```
   - Example `.env`:
     ```env
     POSTGRES_USER=sanqasuq
     POSTGRES_PASSWORD=sanqasuq
     POSTGRES_DB=sanqasuq
     POSTGRES_HOST=postgres
     POSTGRES_PORT=5432
     POSTGRES_URL=postgresql://sanqasuq:sanqasuq@postgres:5432/sanqasuq
     API_VERSION=v1
     JWT_SECRET=supersecretkey
     JWT_ISSUER=example.com
     ```

3. **Install Dependencies**:
   ```bash
   go mod download
   ```

## Running the Application
### Local Development (with Air)
1. Ensure PostgreSQL is running (via Docker Compose or locally).
2. Run with hot reloading (uses `.air.toml` for config):
   ```bash
   air
   ```
   API available at `http://localhost:8080/api/v1`.

### Docker Compose
1. Start the app and PostgreSQL:
   ```bash
   docker-compose up --build
   ```
   API available at `http://localhost:8080/api/v1`. Migrations are applied automatically on container startup.

### Makefile
- Run locally:
  ```bash
  make run
  ```
- Build Docker image:
  ```bash
  make build
  ```

## Database Migrations
Uses `go-migrate` for database schema changes. Migrations are applied automatically when starting the Docker container. For manual migrations (e.g., local development), ensure `.env` is set up.

- **Apply Migrations**:
  ```bash
  make migration_up
  ```
- **Revert Migrations**:
  ```bash
  make migration_down
  ```
- **Fix Migration Version** (replace `VERSION` with desired version):
  ```bash
  make migration_fix VERSION=1
  ```

## API Testing with Bruno
1. Install Bruno: https://www.usebruno.com/
2. Open `api-testing-sanqasuq` folder in Bruno.
3. Use `Development.bru` environment for local testing.
4. Run collections to test endpoints.

## Development Tools
- **Air**: Hot reloading (configured via `.air.toml`). Install:
  ```bash
  curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
  ```
- **Go Migrate**: Database migrations. Install:
  ```bash
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```
- **Bruno**: API testing tool.

## Project Structure
```
.
├── api-testing-sanqasuq/  # Bruno API testing collections
├── cmd/                   # Main application entry point
├── internal/              # Core logic (auth, handlers, services, etc.)
├── migrations/            # Database migration files
├── tmp/                   # Temporary build files
├── .air.toml              # Air hot reload config
├── dev.Dockerfile         # Dev Dockerfile
├── docker-compose.yaml    # Docker Compose setup
├── entrypoint.sh          # Docker entrypoint for migrations
├── example.env            # Example env variables
├── Makefile               # Automation scripts
└── go.mod                 # Go dependencies
```