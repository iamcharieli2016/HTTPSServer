# HTTPS Server - Database Metadata API

[![Go Version](https://img.shields.io/badge/Go-1.21+-blue.svg)](https://golang.org)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)
[![Build Status](https://img.shields.io/badge/build-passing-brightgreen.svg)]()

A secure HTTPS-only backend service providing database table structure metadata query functionality.

## ✨ Features

- 🔒 **HTTPS Only** - Secure encrypted communication
- 🗃️ **MySQL Support** - Full MySQL database integration
- 🔍 **Flexible Search** - Multiple search conditions support
- 📄 **Pagination** - Efficient data pagination
- 🔐 **Authentication** - Client-based authentication
- 📝 **Detailed Logging** - Comprehensive request/response logging
- 🚀 **High Performance** - Built with Gin framework
- 📦 **Easy Deployment** - Cross-platform binary support

## 🏗️ Project Structure

```
├── cmd/
│   └── server/           # Application entry point
├── internal/
│   ├── auth/            # Authentication logic
│   ├── config/          # Configuration management
│   ├── database/        # Database operations
│   ├── handler/         # HTTP request handlers
│   ├── model/           # Data models
│   └── utils/           # Utility functions
├── pkg/
│   └── response/        # Response structures
├── configs/             # Configuration files
├── sql/                 # Database scripts
├── scripts/             # Build and deployment scripts
├── docs/                # Documentation
└── deployments/         # Deployment configurations
```

## 🚀 Quick Start

### Prerequisites

- Go 1.21+
- MySQL 5.7+
- OpenSSL (for SSL certificates)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/yourusername/httpsserver.git
   cd httpsserver
   ```

2. **Install dependencies**
   ```bash
   make deps
   ```

3. **Setup database**
   ```bash
   mysql -u root -p < sql/database.sql
   ```

4. **Configure the application**
   ```bash
   cp configs/config.env.example config.env
   # Edit config.env with your database credentials
   ```

5. **Generate SSL certificates**
   ```bash
   make cert
   ```

6. **Build and run**
   ```bash
   make build
   make run
   ```

The server will start on `https://localhost:18443`

## 🔧 Configuration

Create a `config.env` file based on `configs/config.env.example`:

```env
# Database Configuration
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASSWORD=your_password
DB_DATABASE=metadata_db
DB_CHARSET=utf8mb4

# Server Configuration
SERVER_PORT=18443
SSL_CERT_FILE=server.crt
SSL_KEY_FILE=server.key

# Authentication Configuration
CLIENT_ID=eplat
CLIENT_SECRET=eplat2019111214440
```

## 📚 API Documentation

### Endpoint

```
POST https://localhost:18443/service/D_A_BSPDMETA
```

### Request Format

```json
{
    "params": {
        "columnComment": "",
        "columnType": "",
        "columnName": "",
        "tableComment": "",
        "tableName": "",
        "tableSchema": "",
        "dbType": ""
    },
    "serviceId": "D_A_BSPDMETA",
    "showCount": "true",
    "offset": 0,
    "limit": 10,
    "userId": "171635",
    "clientId": "eplat",
    "clientSecret": "eplat2019111214440"
}
```

### Response Format

```json
{
    "success": true,
    "data": [
        {
            "tableSchema": "test_db",
            "tableName": "users",
            "tableComment": "用户表",
            "columnName": "id",
            "columnType": "INT AUTO_INCREMENT",
            "columnComment": "用户ID",
            "dbType": "mysql"
        }
    ],
    "total": 100
}
```

### Example Usage

```bash
curl -k -X POST https://localhost:18443/service/D_A_BSPDMETA \
  -H 'Content-Type: application/json' \
  -d '{
    "params": {"tableName": "users"},
    "serviceId": "D_A_BSPDMETA",
    "showCount": "true",
    "offset": 0,
    "limit": 10,
    "userId": "171635",
    "clientId": "eplat",
    "clientSecret": "eplat2019111214440"
  }'
```

## 🛠️ Development

### Build Commands

```bash
# Build for current platform
make build

# Build for specific platforms
make build-linux
make build-windows
make build-darwin

# Build for all platforms
make build-all

# Run application
make run

# Run tests
make test

# Format code
make fmt

# Clean build files
make clean
```

### Testing

```bash
# Run all tests
make test

# Run tests with coverage
make test-coverage
```

## 🚀 Deployment

### Option 1: Using Makefile

```bash
# Create deployment package
make package
```

### Option 2: Manual Deployment

1. **Build for target platform**
   ```bash
   make build-linux
   ```

2. **Upload to server**
   ```bash
   scp build/httpsserver-linux user@server:/opt/httpsserver/
   scp -r configs/ sql/ scripts/ user@server:/opt/httpsserver/
   ```

3. **Setup on server**
   ```bash
   # On the server
   cd /opt/httpsserver
   chmod +x httpsserver-linux
   chmod +x scripts/*.sh
   
   # Initialize database
   mysql -u root -p < sql/database.sql
   
   # Configure
   cp configs/config.env.example config.env
   # Edit config.env
   
   # Start service
   ./scripts/start-linux.sh
   ```

### Systemd Service

```bash
# Install systemd service
sudo cp deployments/systemd/httpsserver.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable httpsserver
sudo systemctl start httpsserver
```

## 📋 Database Schema

The application uses a `table_metadata` table to store database metadata:

| Column | Type | Description |
|--------|------|-------------|
| id | INT | Primary key |
| table_schema | VARCHAR(255) | Database name |
| table_name | VARCHAR(255) | Table name |
| table_comment | TEXT | Table comment |
| column_name | VARCHAR(255) | Column name |
| column_type | VARCHAR(255) | Column type |
| column_comment | TEXT | Column comment |
| db_type | VARCHAR(50) | Database type |
| created_at | TIMESTAMP | Created time |
| updated_at | TIMESTAMP | Updated time |

## 🔒 Security

- **HTTPS Only**: All communications are encrypted
- **Client Authentication**: Validates client credentials
- **Input Validation**: Prevents SQL injection
- **CORS Support**: Configurable cross-origin requests

## 📝 Logging

The application provides detailed logging including:
- Request/Response details
- Authentication attempts
- Database operations
- Error tracking

Logs are written to `logs/server.log` by default.

## 🤝 Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 📄 License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## 🆘 Support

If you encounter any issues:

1. Check the [troubleshooting guide](docs/TROUBLESHOOTING.md)
2. Review the logs in `logs/server.log`
3. Open an issue on GitHub

## 🏷️ Versioning

We use [SemVer](http://semver.org/) for versioning. For the versions available, see the [tags on this repository](https://github.com/yourusername/httpsserver/tags).

## 📞 Contact

- Project Link: [https://github.com/yourusername/httpsserver](https://github.com/yourusername/httpsserver)
- Documentation: [https://yourusername.github.io/httpsserver](https://yourusername.github.io/httpsserver) 