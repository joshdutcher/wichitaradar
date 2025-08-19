# PLANNING.md - Wichita Radar Architecture

## Vision
Create a reliable, fast, and simple weather radar aggregation service for the Wichita, Kansas area. Focus on minimal resource usage, maximum uptime, and clean user experience while aggregating real-time weather data from multiple sources.

## Architecture

### Design Pattern
- **Clean Architecture**: Separation of concerns with clear boundaries
- **Standard Library First**: Minimize external dependencies
- **Middleware Pattern**: Centralized error handling and request processing
- **Template-Driven**: Server-side rendering for optimal performance
- **File-Based Caching**: Simple, reliable data persistence

### Package Structure
```
wichitaradar/
├── cmd/server/                 # Application entry point
│   ├── main.go                 # Server setup and route configuration
│   └── main_test.go           # Integration tests
├── internal/                   # Private application code
│   ├── handlers/              # HTTP handlers for weather pages
│   │   ├── home.go            # Main radar page
│   │   ├── satellite.go       # Satellite imagery
│   │   ├── rainfall.go        # Precipitation maps
│   │   ├── temperatures.go    # Temperature data
│   │   ├── outlook.go         # Weather forecasts (with XML caching)
│   │   └── simple.go          # Static pages (about, resources, etc.)
│   ├── middleware/            # Request processing middleware
│   │   ├── error.go           # Error handling wrapper
│   │   └── error_test.go      # Error handling tests
│   ├── cache/                 # File-based caching system
│   │   ├── cache.go           # Core caching interface
│   │   ├── directories.go     # Cache directory management
│   │   └── cache_test.go      # Caching tests
│   ├── config/                # Configuration management
│   │   ├── config.go          # Environment and settings
│   │   └── config_test.go     # Configuration tests
│   └── testutils/             # Testing utilities
│       └── cache.go           # Mock cache for testing
├── pkg/templates/             # Reusable template system
│   └── templates.go           # Template initialization and rendering
├── static/                    # Static web assets
│   ├── css/                   # Stylesheets
│   ├── js/                    # JavaScript files
│   └── images/                # Static images
└── templates/                 # HTML templates
    ├── layout.html            # Base template
    ├── home.html              # Radar page template
    ├── satellite.html         # Satellite page template
    └── [other page templates]
```

### Data Architecture
- **Caching Strategy**: File-based with configurable TTL (5 minutes default)
- **Cache Types**: XML data cache for weather forecasts, animated cache for radar loops
- **Data Sources**: Multiple weather service APIs (NWS, radar providers)
- **Error Handling**: Graceful degradation when sources are unavailable

### Key Abstractions
- **FileCache**: Interface for file-based caching with TTL
- **ErrorHandler**: Middleware wrapper for consistent error handling
- **Template System**: Centralized template management and rendering
- **Handler Functions**: Clean HTTP handlers with error return values

## Tech Stack

### Core Technologies
- **Backend**: Go 1.21 with standard library (net/http, html/template)
- **Frontend**: Server-side rendered HTML with PureCSS framework
- **Templating**: Go's html/template with custom template system
- **Caching**: File-based caching with configurable TTL
- **Error Tracking**: Sentry for client-side and server-side error monitoring
- **Static Assets**: Direct file serving via http.FileServer

### Testing
- **Unit Tests**: Comprehensive test coverage with race detection
- **Test Utilities**: Mock cache and testing helpers
- **CI/CD**: GitHub Actions with automated testing
- **Coverage**: Coverage reporting and analysis

## Required Tools

### Development Environment
- Go 1.21+ installed and configured
- Make for build automation
- Git for version control
- Text editor with Go language support

### External Services Setup
- **Sentry Account**: For error tracking and monitoring
  - Create project and obtain DSN
  - Set SENTRY_DSN environment variable
- **Render Account**: For production deployment
  - Connected to GitHub repository
  - Auto-deploy on successful builds
- **UptimeRobot**: For monitoring and auto wake-up
  - Health check endpoint: `/health`

### API Keys & Configuration
- **SENTRY_DSN**: Required for error tracking
- **PORT**: Configurable port (default: 80)
- **ENV**: Environment flag (development/production)

## Design Specifications

### User Experience
- **Mobile-First**: Responsive design optimized for mobile devices
- **Performance**: Sub-3 second page loads, auto-refresh every 5 minutes
- **Accessibility**: Clean, readable interface with proper semantic markup
- **Error Handling**: User-friendly error messages in production

### Visual Design
- **Framework**: PureCSS for lightweight, responsive styling
- **Layout**: Clean, minimal interface focused on weather data
- **Colors**: Weather-appropriate color scheme
- **Typography**: Readable fonts optimized for data display

## Performance Targets

### Response Time Goals
- **Static Content**: <100ms
- **Dynamic Pages**: <500ms
- **Cached Data**: <200ms
- **Health Check**: <50ms

### Resource Usage
- **Memory**: <100MB steady state
- **CPU**: Minimal usage during normal operation
- **Disk**: Efficient cache cleanup and rotation
- **Network**: Optimized external API calls

## Security & Privacy

### Security Measures
- **Input Validation**: Proper sanitization of all inputs
- **Error Handling**: No sensitive information in error messages
- **HTTPS**: Production deployment with SSL/TLS
- **Content Security**: Proper content type headers

### Privacy Approach
- **No User Tracking**: Minimal client-side JavaScript
- **Public Data**: Only public weather data aggregation
- **Error Tracking**: Sentry for debugging, no personal data collection

## Implementation Status
*For current project status and active development progress, see CLAUDE.md and SESSION.md*

### Architecture Status
- ✅ Core server architecture implemented
- ✅ Middleware system with error handling
- ✅ File-based caching system
- ✅ Template system with layout inheritance
- ✅ Static file serving
- ✅ Health check endpoint
- ✅ Sentry integration
- ✅ Production deployment pipeline

### Current Technical Debt
- Integration tests could be expanded
- Code linting not yet integrated in CI
- Performance benchmarking not implemented
- SonarCloud integration planned but not implemented

## Future Considerations

### Previously Identified Features
- Integration tests in CI/CD pipeline
- Code linting automation
- Performance benchmarking
- Enhanced error tracking metrics
- SonarCloud code quality analysis

### Additional Future Enhancement Ideas
*Brainstorming notes for potential far-future development:*

#### Performance Enhancements
- Redis caching layer for high-traffic scenarios
- CDN integration for static assets
- HTTP/2 and compression optimizations
- WebP image format support

#### Feature Expansions
- Weather alerts and notifications
- Historical weather data archive
- Storm tracking and analysis
- Multi-location support
- API endpoints for programmatic access
- Mobile PWA capabilities

#### Monitoring & Analytics
- Detailed performance metrics
- User experience analytics (privacy-preserving)
- Weather data source reliability tracking
- Automated performance regression detection

#### Infrastructure Improvements
- Multi-region deployment
- Database integration for persistent data
- Kubernetes deployment options
- Automated backup and disaster recovery