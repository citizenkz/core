# Citizen REST API

A comprehensive REST API for managing citizen benefits, categories, filters, and user authentication with email notifications.

## Features

- 🔐 **Authentication & Authorization** - JWT-based auth with password reset via email OTP
- 📧 **Email Notifications** - SMTP integration for account activities
- 🎁 **Benefits Management** - Create, read, update, delete benefits with filters and categories
- 🏷️ **Categories** - Organize benefits by categories
- 🔍 **Filters** - Advanced filtering with range support (string, number, date)
- 👤 **User Profile** - Complete profile management including email/password updates

## Technology Stack

- **Language**: Go 1.21+
- **Framework**: Chi Router
- **Database**: PostgreSQL
- **ORM**: Ent
- **Authentication**: JWT
- **Email**: SMTP (Gmail)

## Quick Start

### Prerequisites

- Go 1.21 or higher
- PostgreSQL
- Gmail account with App Password for SMTP

### Installation

1. Clone the repository:
```bash
git clone https://github.com/citizenkz/core.git
cd core
```

2. Install dependencies:
```bash
go mod download
```

3. Configure the application by creating `local.yaml`:
```yaml
env: "local"
port: 8080
database:
  user: "your_db_user"
  password: "your_db_password"
  host: "localhost"
  port: 5432
  name: "citizen"
  sslmode: "disable"
jwtsecret: "your-secret-key-here"
smtp:
  host: "smtp.gmail.com"
  port: 587
  username: "your-email@gmail.com"
  password: "your-app-password"
  from: "noreply@citizen.com"
```

4. Run the application:
```bash
go run main.go
```

The server will start on `http://localhost:8080`

## API Documentation

Base URL: `http://localhost:8080/api/v1`

Complete API documentation is available in [`api-endpoints.json`](./api-endpoints.json)

### Authentication Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/auth/register` | Register new user | No |
| POST | `/auth/login` | Login user | No |
| GET | `/auth/profile` | Get user profile | Yes |
| PUT | `/auth/password` | Update password | Yes |
| PUT | `/auth/email` | Update email | Yes |
| DELETE | `/auth/profile` | Delete account | Yes |
| POST | `/auth/forget-password` | Request password reset OTP | No |
| POST | `/auth/forget-password/confirm` | Confirm OTP & reset password | No |

### Category Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/category/` | Create category | No |
| POST | `/category/list` | List categories | No |
| GET | `/category/{id}` | Get category | No |
| PUT | `/category/{id}` | Update category | No |
| DELETE | `/category/{id}` | Delete category | No |

### Filter Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/filter/` | Create filter | No |
| GET | `/filter/` | List filters | No |
| POST | `/filter/save` | Save user filters | No |

### Benefit Endpoints

| Method | Endpoint | Description | Auth Required |
|--------|----------|-------------|---------------|
| POST | `/benefit/` | Create benefit | No |
| POST | `/benefit/list` | List benefits (with filters) | No |
| GET | `/benefit/{id}` | Get benefit | No |
| PUT | `/benefit/{id}` | Update benefit | No |
| DELETE | `/benefit/{id}` | Delete benefit | No |

## Filter Types

The API supports three types of filters:

1. **STRING_RANGE** - Single value filter (e.g., status, category)
2. **NUMBER_RANGE** - Numeric range filter (e.g., age 18-25)
3. **DATE_RANGE** - Date range filter (e.g., valid dates)

## Testing

Run the test script to verify all endpoints:

```bash
chmod +x test.sh
./test.sh
```

The script will test all endpoints using `aidosg65@gmail.com` as the test email.

## Email Notifications

The system sends email notifications for:

- Password reset OTP (6-digit code, expires in 10 minutes)
- Password successfully changed
- Email address changed
- Account deleted

## Benefit Filtering Logic

When listing benefits with filters:
- Benefits **without** a specific filter are shown (pass through)
- Benefits **with** matching filter values are shown
- Benefits with non-matching filter values are **excluded**

Example: If searching for age 18-25:
- Benefit with no age filter: **Shown** ✓
- Benefit with age 20-30: **Shown** ✓ (overlaps)
- Benefit with age 30-40: **Hidden** ✗ (no overlap)

## Project Structure

```
core/
├── app/                 # Application setup
├── config/              # Configuration management
├── ent/                 # Ent ORM schemas and generated code
│   └── schema/         # Database schemas
├── services/           # Business logic by domain
│   ├── auth/          # Authentication & user management
│   ├── benefit/       # Benefit management
│   ├── category/      # Category management
│   └── filter/        # Filter management
├── utils/             # Utility functions
│   ├── email/        # Email service
│   ├── gen/          # ID generation
│   ├── json/         # JSON helpers
│   └── jwt/          # JWT token handling
├── api-endpoints.json # Complete API documentation
├── test.sh           # API testing script
└── main.go           # Application entry point
```

## Development

### Adding a New Service

1. Create service directory: `services/your-service/`
2. Add entity models in `entity/`
3. Implement storage layer in `storage/`
4. Add business logic in `usecase/`
5. Create HTTP handlers in `server/`
6. Register routes in `app/app.go`

### Database Migrations

Ent automatically creates/updates the schema when the application starts. For production, consider using migration files.

## License

This project is licensed under the MIT License.

## Contact

For questions or support, please open an issue on GitHub.
