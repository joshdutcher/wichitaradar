# CLAUDE.md - Wichita Radar Development Guide

## Current Project Status (High-Level)

### ✅ Completed Milestones
- ✅ Core weather radar aggregation functionality
- ✅ Responsive web design with PureCSS framework
- ✅ Server-side HTML templating system
- ✅ Sentry error tracking integration
- ✅ Comprehensive middleware and error handling
- ✅ CI/CD pipeline with GitHub Actions
- ✅ Production deployment on Render
- ✅ UptimeRobot monitoring and auto wake-up
- ✅ File-based caching system for weather data
- ✅ Full test coverage with race detection

### 🎯 Current Phase
**Current Phase**: Production Maintenance & Enhancement
**Status**: Stable production service serving Wichita weather data
**Next Priority**: Performance optimization and feature enhancements
**Recent Achievement**: Solid CI/CD pipeline with automated testing and deployment

### 🛠️ Development Environment Status
- Go 1.21 with minimal external dependencies (only Sentry)
- Standard library HTTP server
- File-based caching with configurable TTL
- Environment-aware error handling (dev vs prod)

## Development Standards and Patterns

### Code Standards
- **Language**: Go with standard library emphasis
- **Architecture**: Clean separation of concerns (handlers, middleware, cache, config)
- **Error Handling**: Comprehensive error handling with middleware wrapper
- **Testing**: Unit tests with race detection, aim for high coverage
- **Dependencies**: Minimal external dependencies (currently only Sentry)

### File Organization
*See PLANNING.md for detailed package structure and architecture specifications.*

```
wichitaradar/
├── cmd/server/           # Main application entry point
├── internal/             # Private application code
│   ├── handlers/         # HTTP handlers for weather pages
│   ├── middleware/       # Error handling middleware
│   ├── cache/           # File-based caching system
│   └── config/          # Configuration management
├── pkg/templates/       # Reusable template system
├── static/              # CSS, JS, images
└── templates/           # HTML templates
```

### Implementation Guidelines
*See PLANNING.md for detailed implementation specifications, architecture patterns, and design requirements.*

## Development Environment

### System Requirements
- Go 1.21 or later
- Unix-like environment (Linux/macOS preferred)
- Network access for weather data sources

**Build Commands:**
```bash
# Development
go run cmd/server/main.go

# Testing
make test          # Run full test suite
make test-race     # Run with race detection
make coverage      # Generate coverage report

# Production build
go build -o wichitaradar cmd/server/main.go
```

**Requirements:**
- Port 80 for production, configurable via PORT env var
- SENTRY_DSN environment variable for error tracking
- ENV=production for production mode

**Benefits:**
- Minimal resource usage
- Fast startup times
- Simple deployment
- Reliable error tracking

## Testing Checklist

### Pre-deployment Validation
- [ ] All unit tests pass with race detection
- [ ] Error handling middleware properly wraps all handlers
- [ ] Template rendering works for all pages
- [ ] Static file serving functions correctly
- [ ] Cache directories are created and accessible
- [ ] Sentry integration initializes (with valid DSN)
- [ ] Health endpoint returns 200 OK
- [ ] All weather data sources respond correctly

### Performance Checks
- [ ] Response times under 500ms for cached content
- [ ] Memory usage stable over time
- [ ] No goroutine leaks
- [ ] File cache cleanup working properly

## Common Pitfalls to Avoid

### Development Pitfalls
- Don't bypass the error handling middleware
- Don't add unnecessary dependencies without justification
- Don't ignore race conditions in tests
- Don't hardcode file paths (use filepath.Join)
- Don't skip cache directory creation checks

### Production Pitfalls
- Always set SENTRY_DSN in production
- Ensure ENV=production is set for proper error handling
- Monitor disk space for cache directories
- Don't ignore UptimeRobot alerts

## Reference Information

### Key Dependencies
*See PLANNING.md for complete dependency list and version specifications.*

- **Core**: Go standard library (net/http, html/template)
- **External**: github.com/getsentry/sentry-go v0.32.0
- **Frontend**: PureCSS framework (CDN)

### Project Context
Personal weather aggregation project serving Wichita, Kansas area. Created by Josh Dutcher as a practical solution for local weather monitoring. Non-commercial, publicly available service with focus on reliability and simplicity.

## Questions to Ask User

### When Adding Features
1. Should this feature maintain the minimal dependency approach?
2. Does this require additional caching strategy?
3. How should errors be handled for this feature?
4. What's the performance impact on the single-threaded model?

### When Debugging Issues
1. Is this happening in development or production?
2. Are we seeing errors in Sentry?
3. Is the cache working correctly?
4. Are all weather data sources responding?

### When Optimizing
1. What are the current response times?
2. Is memory usage growing over time?
3. Are we hitting any rate limits on weather data sources?
4. Should we implement additional caching layers?