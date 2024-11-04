# HypeBid

HypeBid is a platform that integrates with Twitch to track and analyze streamer metrics, enabling users to participate in prediction markets based on channel performance.

## Features

- Twitch OAuth Integration
- Real-time Follower Count Tracking
- Batch Processing for Multiple Channels
- Market Creation and Management
- User Authentication and Authorization
- Metric History and Analysis

## Getting Started

### Prerequisites

- Go 1.21+
- Docker
- PostgreSQL
- Twitch Developer Account

### Environment Setup
1. Create a Twitch Application:

You will need to create a Twitch Application at https://dev.twitch.tv/console/apps to get the `TWITCH_CLIENT_ID` and `TWITCH_CLIENT_SECRET`.

2. Clone the repository
```bash
git clone https://github.com/hypebid/hypebid-app.git
cd hype-bid-prototype
```

3. Create a `.env` file with the following variables:
```env
ENVIRONMENT=production
HOST_URL=http://localhost:8080
TWITCH_CLIENT_ID=your_twitch_client_id
TWITCH_CLIENT_SECRET=your_twitch_client_secret
DB_HOST=db
DB_USER=postgres
DB_PASSWORD=your_db_password
DB_NAME=hypebid
DB_PORT=5432
SERVER_PORT=8080
FOLLOWER_UPDATE_INTERVAL=15
TRACKED_LOGINS=twitch_login1,twitch_login2,twitch_login3...
HOST_URL=call_back_url_for_twitch_oauth
```

### Running the Application

Right now, the application can be run with Docker Compose. It will create a PostgreSQL database container and then connect the HypeBid server to the database.

1. Build the application:
```bash
docker build -t hypebid/hypebid:v1.0 .
```

2. Start the application:
```bash
docker compose up
```

## API Endpoints

### Authentication
- `GET /api/v1/auth/twitch/login` - Initiate Twitch OAuth flow
- `GET /api/v1/auth/twitch/callback` - OAuth callback handler

### Twitch Integration
- `GET /api/v1/twitch/followers?login={login}` - Get follower count for a channel
- `GET /api/v1/twitch/bulk-followers?login={login1}&login={login2}` - Get follower counts for multiple channels
- `GET /api/v1/twitch/users?login={login}` - Get Twitch user information

## Development

### Generate Mocks
```bash
mockery --all
```

### Running Tests
```bash
go test ./interanl/services
```

### Database Access
```bash
psql -h localhost -p 5432 -U postgres -d hypebid
```

## Architecture

The application follows a clean architecture pattern with the following structure:

- `cmd/` - Application entry points
- `internal/` - Private application code
  - `auth/` - Authentication and OAuth
  - `handlers/` - HTTP request handlers
  - `middleware/` - HTTP middleware
  - `services/` - Business logic
  - `repositories/` - Data access layer
  - `models/` - Domain models
  - `config/` - Configuration management

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Twitch API Documentation
- Go OAuth2 Package
- GORM