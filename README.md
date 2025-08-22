
# Chirpy

Chirpy is a mock-up Twitter backend I made as part of the Boot.dev course [Learn HTTP Servers in Go](https://www.boot.dev/courses/learn-http-servers-golang). It isn't intended for real use. It's just an exercise I used to familiarize myself with HTTP.



## Features
- PostgreSQL database built with SQLc, Goose and psql
- JWT and refresh token authentication
- Webhook support (for fictional provider, Polka)
- Authorization checks for (some) admin requests

## API Endpoints

### User Management

#### POST `/api/users`
Create a new user account.

**Request JSON:**
```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```
**Response JSON:**
```json
{
  "id": "uuid-string",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "email": "user@example.com",
  "is_chirpy_red": false
}
```

---

#### PUT `/api/users`
Update an existing user's information (requires authentication).

**Request Header:**
`Authorization: Bearer <access_token>`

**Request JSON:**
```json
{
  "email": "newemail@example.com",
  "password": "newpassword"
}
```
**Response JSON:** Same as user creation.

---

#### POST `/api/login`
Log in with email and password to receive an access token and refresh token.

**Request JSON:**
```json
{
  "email": "user@example.com",
  "password": "yourpassword"
}
```
**Response JSON:**
```json
{
  "id": "uuid-string",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "email": "user@example.com",
  "is_chirpy_red": false,
  "token": "<access_token>",
  "refresh_token": "<refresh_token>"
}
```

---

### Chirp Management

#### POST `/api/chirps`
Create a new chirp (requires authentication).

**Request Header:**
`Authorization: Bearer <access_token>`

**Request JSON:**
```json
{
  "body": "Hello, world!"
}
```
**Response JSON:**
```json
{
  "id": "uuid-string",
  "created_at": "timestamp",
  "updated_at": "timestamp",
  "user_id": "uuid-string",
  "body": "Hello, world!"
}
```

---

#### GET `/api/chirps`
Retrieve a list of all chirps.

**Query Parameters (optional):**
- `author_id`: Filter by user ID
- `sort`: `asc` or `desc` (default: asc)

**Response JSON:**
```json
[
  {
    "id": "uuid-string",
    "created_at": "timestamp",
    "updated_at": "timestamp",
    "user_id": "uuid-string",
    "body": "Hello, world!"
  }
]
```

---

#### GET `/api/chirps/{chirpID}`
Retrieve a single chirp by its ID.

**Response JSON:** Same as above, but for one chirp.

---

#### DELETE `/api/chirps/{chirpID}`
Delete a chirp by its ID (requires authentication and ownership).

**Request Header:**
`Authorization: Bearer <access_token>`

**Response:**
HTTP 204 No Content on success.

---

### Authentication & Session

#### POST `/api/refresh`
Exchange a valid refresh token for a new access token.

**Request Header:**
`Authorization: Bearer <refresh_token>`

**Response JSON:**
```json
{
  "token": "<new_access_token>"
}
```

---

#### POST `/api/revoke`
Revoke a refresh token, logging the user out of that session.

**Request Header:**
`Authorization: Bearer <refresh_token>`

**Response:**
HTTP 204 No Content on success.

---

### Health & Admin

#### GET `/api/healthz`
Health check endpoint for monitoring.

**Response:**
Plain text: `OK`

---

#### POST `/api/polka/webhooks`
Endpoint for receiving webhooks from the Polka payment service.

**Request Header:**
`Authorization: ApiKey <polka_key>`

**Request JSON:**
```json
{
  "event": "user.upgraded",
  "data": {
    "user_id": "uuid-string"
  }
}
```
**Response:**
HTTP 204 No Content.

---

#### GET `/admin/metrics`
Retrieve server metrics (for admin/monitoring).

**Response:**
HTML page with metrics.

---

#### POST `/admin/reset`
Reset server state (only available in development mode).

**Response:**
Plain text: `Reset data`

---

## Environment Variables (.env)

Requires a `.env` file in your project root with the following variables:

```env
DB_URL= <...>
PLATFORM= <...>
SECRET= <...>
POLKA_KEY= <...>
```

**Variable descriptions:**
- `DB_URL`: PostgreSQL connection string for Chirpy database.
- `PLATFORM`: Set to `dev` for development mode (enables admin reset endpoint).
- `SECRET`: Secret key used for signing JWTs (keep this safe!).
- `POLKA_KEY`: API key for authenticating Polka webhook requests.

---

Thanks for checking out my little project :) stay tuned for when I use this knowledge to make something cool!
