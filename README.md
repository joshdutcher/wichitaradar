# Wichita Radar

A production-ready Go web application providing real-time weather radar and satellite imagery for the Wichita, Kansas area. Built with enterprise-grade engineering practices, featuring intelligent caching, background monitoring, and comprehensive error handling.

**🌐 Live Site**: [wichitaradar.com](https://wichitaradar.com)

## Core Features

- **Real-time Data Processing**: Animated radar loops with live weather data aggregation
- **Multi-Source Integration**: Satellite imagery, NWS weather stories, and precipitation maps
- **Intelligent Caching**: TTL-managed file caching with automatic expiration
- **Background Monitoring**: Goroutine-based health checks and failure detection
- **Mobile-Responsive**: Progressive enhancement with mobile-first design
- **Auto-refresh**: 5-minute intervals to maintain current animations

## Technical Architecture

**Backend**: Go 1.21 with standard library `net/http` and custom middleware stack  
**Frontend**: Server-side HTML templating with PureCSS framework  
**Data Processing**: XML parsing of NWS feeds with temporal filtering  
**Caching**: Custom file-based cache with configurable TTL  
**Error Handling**: Multi-layer error middleware with environment-aware responses  
**Monitoring**: Sentry integration with intelligent error aggregation and background failure detection  
**Testing**: Comprehensive test coverage (29 Go files, 12 test files) with race condition detection

## Development Workflow

### Prerequisites
- Go 1.21+
- Make (for build automation)

### Local Development
```bash
# Clone and setup
git clone https://github.com/joshdutcher/wichitaradar.git
cd wichitaradar

# Run the full test suite
make test

# Run with race condition detection  
make test-race

# Generate coverage report
make coverage

# Start development server
go run cmd/server/main.go
```

### CI/CD Pipeline
- **GitHub Actions**: Automated testing and deployment
- **Testing**: Unit tests with race detection (`go test -race`)  
- **Coverage**: Test coverage analysis and reporting
- **Deployment**: Automatic Railway deployment via nixpacks on successful builds
- **Monitoring**: UptimeRobot availability monitoring with automatic wake-up

## Production Architecture

### Infrastructure
- **Hosting**: Railway with nixpacks build configuration
- **Build**: Native Go compilation with embedded timezone data
- **Monitoring**: Multi-layer health monitoring (UptimeRobot + Sentry + custom health endpoints)
- **Performance**: Sub-3-second page loads with intelligent caching
- **Reliability**: 99.9%+ uptime with automatic instance wake-up

### Error Handling & Monitoring
- **Custom Middleware**: Centralized error handling with environment-aware messaging
- **Sentry Integration**: Advanced error tracking with contextual tagging and scoping
- **Background Monitoring**: Goroutine-based persistent failure detection
- **Graceful Degradation**: Robust fallback mechanisms for external service failures

## Project Structure

```
wichitaradar/
├── cmd/server/              # Application entry point and server setup
├── internal/                # Private application packages
│   ├── handlers/           # HTTP handlers with weather-specific logic
│   ├── middleware/         # Custom error handling middleware  
│   ├── cache/             # File-based caching system with TTL
│   └── config/            # Configuration management
├── pkg/templates/          # Reusable template system
├── menu/                  # Navigation menu with timezone handling
├── static/                # Static assets (CSS, JS, images)
├── templates/             # HTML templates with layout inheritance
├── nixpacks.toml         # Railway deployment configuration
└── .claude/              # Project documentation and development guides
```

## Development Roadmap

### Completed ✅
- [x] Production-ready error handling with middleware stack
- [x] Intelligent caching system with TTL management  
- [x] Background monitoring with persistent failure detection
- [x] Comprehensive test coverage with race detection
- [x] Railway deployment with automated CI/CD
- [x] Enhanced error tracking with Sentry integration

### Planned 🚧
- [ ] Integration tests in CI/CD pipeline
- [ ] Code linting automation (golangci-lint)
- [ ] Performance benchmarking and optimization
- [ ] SonarCloud code quality analysis

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This site aggregates weather data from various sources. The creator does not claim ownership of any of the radar images or data displayed. This is a personal project made available for public use.

## Contact

For questions or concerns, contact Josh Dutcher at josh.dutcher@joshdutcher.com
