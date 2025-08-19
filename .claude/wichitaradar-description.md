 # Wichita Radar Website Description

  **wichitaradar.com** is a Go-based web application that provides real-time weather radar and satellite imagery
  specifically for the Wichita, Kansas area. Created by Josh Dutcher as a personal project, it aggregates weather data from
  various sources to offer a convenient, mobile-responsive weather monitoring solution.

  ## Primary Purpose

  The site serves as a weather monitoring dashboard, originally built so the creator could track weather conditions when his
   satellite dish went out. It has since become a public resource for anyone needing quick access to local weather radar and
   satellite data.

  ## Core Features

  - **Real-time Radar Imagery**: Displays animated radar loops from various local weather sources
  - **Satellite Imagery**: Provides both infrared and visible light satellite views
  - **Weather Stories & Forecasts**: Integrates National Weather Service (NWS) weather stories and forecasts
  - **Precipitation Maps**: Shows 24-hour and 7-day precipitation data
  - **Auto-refresh**: Main pages automatically refresh every 5 minutes to keep animations current
  - **Mobile-responsive Design**: Works across desktop and mobile devices

  ## Technical Architecture

  - **Backend**: Built in Go 1.21 using standard library `net/http` with custom middleware stack
  - **Templating**: Server-side HTML templating with Go's `html/template` and custom template management system
  - **Frontend**: Uses PureCSS framework for responsive styling with mobile-first design
  - **Error Tracking**: Multi-layer error handling with Sentry integration and custom AppError middleware
  - **Deployment**: Hosted on Railway with nixpacks configuration, embedded timezone data, and automated CI/CD
  - **Monitoring**: UptimeRobot availability monitoring with automatic instance wake-up and comprehensive health checks

  ## Advanced Technical Features

  ### 🏗️ **Intelligent Caching System**
  - **Custom File-Based Cache**: TTL-managed caching with automatic expiration and cleanup
  - **Weather Data Optimization**: Smart caching of XML weather feeds with configurable refresh intervals
  - **Performance Enhancement**: Reduces external API calls while maintaining data freshness

  ### 🌐 **Real-Time Data Processing**
  - **XML Feed Parsing**: Dynamic parsing of National Weather Service XML data with struct mapping
  - **Temporal Filtering**: Time-based validation using Unix timestamps to show only current weather stories
  - **Multi-Source Integration**: Seamless aggregation of radar imagery, satellite data, and forecast information

  ### 🚀 **Background Processing & Monitoring**
  - **Goroutine-Based Monitoring**: Background health checks with ticker-based persistent failure detection
  - **Intelligent Error Tracking**: Client-side image error detection with server-side aggregation and pattern analysis
  - **Automated Alerting**: Threshold-based failure reporting with automatic Sentry integration and contextual tagging

  ### 🛡️ **Production-Ready Error Handling**
  - **Custom Middleware Stack**: Centralized error handling with environment-aware messaging (detailed dev/friendly prod)
  - **Graceful Degradation**: Robust fallback mechanisms for external service failures
  - **Comprehensive Logging**: Structured logging with failure pattern detection and automatic log rotation

  ## Page Structure

  The site includes multiple specialized pages:

  - **Home**: Main radar dashboard with auto-refresh
  - **Satellite**: Dedicated satellite imagery page
  - **Rainfall**: Precipitation tracking and maps
  - **Temperatures**: Temperature data visualization
  - **Outlook**: Weather outlook and forecasts
  - **About**: Background information about the site and creator

  ## Performance & Reliability Engineering

  ### 🔧 **Development Excellence**
  - **Test-Driven Development**: Comprehensive unit test coverage with race condition detection using `go test -race`
  - **CI/CD Pipeline**: GitHub Actions integration with automated testing, coverage reporting, and zero-downtime deployments
  - **Clean Architecture**: Separation of concerns with dedicated handlers, middleware, caching, and configuration layers
  - **Code Quality**: Consistent error handling patterns, dependency injection, and interface-based design

  ### ⚙️ **Infrastructure & DevOps**
  - **Container Deployment**: Railway hosting with nixpacks build optimization and automatic scaling
  - **Health Monitoring**: Multi-layer monitoring with UptimeRobot, Sentry error tracking, and custom health endpoints
  - **Security**: Environment-aware configuration, secure error handling, and production secret management
  - **Performance**: Sub-3-second page loads with intelligent caching and optimized asset delivery

  ## Operational Characteristics

  - **High Availability**: 99.9%+ uptime with automatic instance wake-up and comprehensive health monitoring  
  - **Scalable Architecture**: Goroutine-based concurrent processing with efficient resource utilization
  - **Data Integrity**: Robust error handling with graceful degradation when external services are unavailable
  - **Cross-Browser Compatibility**: Optimized for modern browsers with responsive design and progressive enhancement
  - **Professional Engineering**: Non-commercial project demonstrating enterprise-grade development practices

  ## Project Evolution

  This weather aggregation platform demonstrates a full-stack Go application with production-ready engineering practices.
  Originally built as a personal solution for local weather monitoring, it showcases advanced technical capabilities including
  real-time data processing, intelligent caching strategies, background monitoring systems, and comprehensive error handling.
  The project exemplifies clean architecture principles, test-driven development, and modern DevOps practices while serving
  as a community resource for weather data access.