
# Personal API - HNG DevOps Stage 1

A minimal HTTP API built with Go, deployed behind Nginx as a reverse proxy on a Linux server. This project demonstrates basic API development, process management with systemd, and reverse proxy configuration.

## Live Deployment

**Base URL:** `https://personal-profile-api.vercel.app/`

## API Endpoints

All endpoints return `Content-Type: application/json` with HTTP status `200 OK`.

| Method | Endpoint | Description | Example Response |
|--------|----------|-------------|------------------|
| GET | `/` | Health check | `{"message":"API is running"}` |
| GET | `/health` | Detailed health status | `{"message":"healthy"}` |
| GET | `/me` | Developer information | `{"name":"Your Full Name","email":"you@example.com","github":"https://github.com/LunarKhord"}` |

### Example Request

```bash
curl http://your-server-ip/me
```

### Example Response

```json
{
  "name": "Your Full Name",
  "email": "you@example.com",
  "github": "https://github.com/yourusername"
}
```

## Tech Stack

- **Language:** Go 1.21+
- **Deployment:** Railway


## Running Locally

### Prerequisites
- Go 1.21+

### Steps
1. Clone the repository:
   ```bash
   git clone https://github.com/LunarKhord/personal-profile-api 
   cd personal-profile-api
   ```

2. Run the API:
   ```bash
   go run main.go
   ```

3. Test the endpoints:
   ```bash
   curl http://localhost:8080/
   curl http://localhost:8080/health
   curl http://localhost:8080/me
   ```

3. **Test the endpoints**
   ```bash
   curl http://localhost:8080/
   curl http://localhost:8080/health
   curl http://localhost:8080/me
   ```

The server listens on port `8080` by default.




## Project Structure

```
.
├── main.go          # API source code
├── go.mod           # Go module definition
├── README.md        # This file
└── .gitignore       # Git ignore rules
```

---

Built for HNG DevOps Stage 1.
