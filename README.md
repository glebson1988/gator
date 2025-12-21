# gator

gator is a small RSS aggregation CLI built with Go and Postgres.

## Requirements

- Go (for building/installing the CLI)
- Postgres (for storage)

## Install

Use `go install` to install the CLI binary:

```bash
go install github.com/glebson1988/gator@latest
```

This installs the `gator` binary into your `$GOBIN` (or `$GOPATH/bin`).

## Configure

Create a config file at `~/.gatorconfig.json`:

```json
{
  "db_url": "postgres://username:password@localhost:5432/gator?sslmode=disable",
  "current_user_name": "alice"
}
```

Ensure your database exists and has the schema applied. The schema lives in `sql/schema`.

### Migrations (goose)

If you use goose:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
goose -dir sql/schema postgres "$DATABASE_URL" up
```

### Query codegen (sqlc)

If you change SQL in `sql/queries`, regenerate the query code:

```bash
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
sqlc generate
```

## Run

Once installed, use the `gator` binary:

```bash
gator register <name>
gator login <name>
gator addfeed <name> <url>
gator follow <url>
gator following
gator unfollow <url>
gator feeds
gator browse
gator agg <interval>
gator reset
```

Notes:
- `register` creates a user.
- `login` sets the active user in the config file.
- `addfeed` creates a feed and automatically follows it.
- `follow` subscribes the current user to a feed by URL.
- `agg` fetches and stores posts on a schedule, e.g. `gator agg 10s`.
- `reset` clears all users and related data.

## Development

For local development, you can run:

```bash
go run . <command> [args...]
```

For production usage, prefer the compiled `gator` binary from `go install` or `go build`.
