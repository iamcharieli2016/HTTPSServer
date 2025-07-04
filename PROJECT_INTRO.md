# HTTPS Server - Enterprise Go Microservice

## Project Overview

A production-ready HTTPS server built with Go that provides secure database metadata querying services. This project demonstrates the complete transformation from a single-file application (358 lines) into an enterprise-grade microservice following modern Go best practices.

## Key Features

**🔐 Security & Performance**
- HTTPS/TLS encryption with self-signed certificate generation
- Client authentication system with secure credential validation
- High-performance architecture: <5s startup, <50MB memory usage, <100ms response time

**🗃️ Database Integration**
- MySQL connectivity with advanced metadata querying capabilities
- Flexible filtering by schema, table, and column parameters
- Efficient pagination and connection pooling

**🏗️ Modern Architecture**
- Clean, modular design with dependency injection
- Standard Go project layout with clear separation of concerns
- Comprehensive unit testing (70%+ coverage) with CI/CD integration

## Technology Stack

- **Backend**: Go 1.21+ with Gin web framework
- **Database**: MySQL with native driver
- **DevOps**: Docker containerization, GitHub Actions CI/CD
- **Infrastructure**: Nginx reverse proxy, multi-platform builds
- **Quality**: Automated testing, security scanning (Gosec, Nancy)

## Production Ready

The application is enterprise-ready with:
- **Containerization**: Multi-stage Docker builds (15MB images)
- **Monitoring**: Health check endpoints and structured logging
- **Deployment**: Support for traditional, Docker, and Kubernetes deployments
- **Security**: Vulnerability scanning and secure coding practices
- **Scalability**: Stateless design for horizontal scaling

## API Capabilities

Primary endpoint `/service/D_A_BSPDMETA` provides database table structure and column metadata with flexible filtering options. Health monitoring available via `/health` endpoint for load balancer integration.

**Result**: A maintainable, scalable, and secure microservice that showcases modern Go development practices and enterprise-grade architecture patterns. 