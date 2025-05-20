# 🚖 Ride Booking API

A scalable, modular backend service for a ride-hailing application (like Uber or Ola), written in Go. This service allows users to register, login, drivers to update availability/location, and both drivers and riders to manage trips.

---

## 📁 Project Structure

```bash
ridebooking/
├── cmd/
│   └── server/                  # 📦 Entry point (main.go)
├── internal/                    # 🔒 Internal modules
│   ├── db/                      # 📂 MongoDB configuration
│   ├── handler/                 # 🧩 Route handlers (users, drivers, trips)
│   ├── middleware/              # 🔐 JWT and other middlewares
│   ├── model/                   # 📜 Structs and enums
│   ├── repository/              # 💾 DB CRUD logic
│   ├── route/                   # 🌐 Route registration
│   └── service/                 # 💡 Business logic
├── docs/                        # 📚 Swagger auto-generated docs
├── pkg/                         # 🛠️ Shared packages/utilities
├── scripts/                     # 📝 Setup/migration scripts
├── .env                         # ⚙️ Environment variables
├── go.mod / go.sum              # 📦 Dependency management
└── README.md                    # 📘 You're here!

🚀 Getting Started
✅ Prerequisites
Go 1.19+

MongoDB

swag CLI for generating Swagger docs:
go install github.com/swaggo/swag/cmd/swag@latest

⚙️ Setup
# Clone the repo
git clone https://github.com/yourusername/ridebooking.git
cd ridebooking

# Load environment variables
cp .env.example .env
# Update .env with DB connection, JWT secrets etc.

# Install dependencies
go mod tidy

# Generate Swagger docs
swag init -g cmd/server/main.go --parseInternal --parseDependency

# Run the app
go run cmd/server/main.go

📬 API Endpoints
🔐 Authentication
| Method | Endpoint    | Description         |
| ------ | ----------- | ------------------- |
| `POST` | `/register` | Register a new user |
| `POST` | `/login`    | Login and get JWT   |

👤 User APIs (/api/ridebooking/user)
| Method   | Endpoint        | Description             |
| -------- | --------------- | ----------------------- |
| `GET`    | `/user/emailId` | Get user by email ID    |
| `GET`    | `/user/id`      | Get user by user ID     |
| `PUT`    | `/user`         | Update user details     |
| `DELETE` | `/user/emailId` | Delete user by email ID |

🚗 Driver APIs (/api/ridebooking/driver)
| Method | Endpoint                   | Description                      |
| ------ | -------------------------- | -------------------------------- |
| `PUT`  | `/driver/location`         | Update driver location           |
| `PUT`  | `/driver/availability`     | Update driver availability       |
| `GET`  | `/driver/available`        | Get all available drivers        |
| `GET`  | `/driver/nearby?x=10&y=20` | Get nearby drivers to a location |

📍 Trip APIs (/api/ridebooking/trip)
| Method | Endpoint           | Description                |
| ------ | ------------------ | -------------------------- |
| `POST` | `/trip`            | Create a trip              |
| `PUT`  | `/trip`            | Update a trip              |
| `GET`  | `/trip?tripId=123` | Get trip by ID             |
| `GET`  | `/trip/driver`     | Get all trips by driver ID |
| `GET`  | `/trip/rider`      | Get all trips by rider ID  |

📑 Swagger UI
Swagger documentation is available once the server is running:

http://localhost:8090/swagger/index.html

Make sure docs/ is generated correctly using:
swag init -g cmd/server/main.go --parseInternal --parseDependency

🛡️ Middleware
✅ JWT Authentication: All /api/ridebooking/ routes are protected using JWT middleware.

🔐 Add your bearer token to access authenticated endpoints.

🧪 Sample Request Payloads
🔐 Login
POST /login
{
  "email": "driver@example.com",
  "password": "123456"
}

🚗 Create Trip

POST /api/ridebooking/trip
{
  "tripId": "T123",
  "riderId": "U100",
  "driverId": "D200",
  "startTime": "2025-05-20T08:00:00Z",
  "endTime": "2025-05-20T08:30:00Z",
  "status": "pending",
  "startLocation": { "x": 10.5, "y": 20.7 },
  "endLocation": { "x": 15.2, "y": 25.1 },
  "totalDistance": 7.2
}
