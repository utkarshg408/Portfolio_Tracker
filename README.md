# Portfolio Tracker

A Go/PostgreSQL backend foundation for tracking investment portfolios, holdings, snapshots, and activity. This repository was initially a model-only scaffold; it now includes database configuration, an initial schema, a user repository, health checks, and a server-side market quote endpoint.

## Current state

Completed:

- Domain models for users, portfolios, holdings, companies, snapshots, and activity logs.
- PostgreSQL schema migration with foreign keys, uniqueness rules, checks, and indexes.
- Config through environment variables, without committing secrets.
- Database startup/ping, `/health` and `/ready` endpoints.
- Finnhub quote client and `GET /api/v1/market/quote?symbol=AAPL` endpoint.
- User repository methods for creating a user and finding one by email.

Still to build:

- Authentication (password hashing, signup/login, JWT/session middleware).
- Portfolio and holding CRUD handlers, validation, authorization, and tests.
- Scheduled price refreshes and snapshot calculation.
- A frontend/client, rate limiting, CORS policy, logging/monitoring, CI, and deployment.

## Local setup

1. Install Go 1.26+ and Docker Desktop.
2. Start PostgreSQL from the repository root:

   ```powershell
   docker compose up -d
   ```

3. Apply the schema migration. With `psql` installed:

   ```powershell
   Get-Content Backend/migrations/001_initial_schema.sql | docker compose exec -T postgres psql -U portfolio -d portfolio_tracker
   ```

4. Copy `Backend/.env.example` to `Backend/.env`, then set a real `FINNHUB_API_KEY`. Do not commit this file.
5. From `Backend`, export the values (or configure them in your shell) and run:

   ```powershell
   go mod tidy
   go run .
   ```

6. Verify the service:

   ```powershell
   Invoke-RestMethod http://localhost:8080/health
   Invoke-RestMethod "http://localhost:8080/api/v1/market/quote?symbol=AAPL"
   ```

## API-key flow

1. Create a Finnhub account and generate an API key.
2. Store it only in the local `FINNHUB_API_KEY` environment variable or your deployment secret manager.
3. The backend appends the key when it calls Finnhub; browsers must never receive it.
4. Before production, add caching, a provider-rate-limit policy, retry/backoff for transient failures, and a second provider fallback.

## Delivery workflow

1. Start the database and run migrations.
2. Add authentication and protect every user/portfolio route.
3. Add CRUD endpoints and tests for portfolios and holdings.
4. Calculate portfolio totals from holdings plus cached market quotes, then store daily snapshots.
5. Add a scheduler/worker for quote refresh and snapshots.
6. Add a frontend, CI checks (`go test ./...`, lint), and production secret/database configuration.
