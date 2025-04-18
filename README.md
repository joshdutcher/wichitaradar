# Wichita Radar

A Go web application providing real-time weather radar and satellite imagery for the Wichita, Kansas area. The site features animated radar loops, satellite imagery, and weather stories from the National Weather Service.

## Features

- Real-time radar imagery from various local sources
- Satellite imagery with infrared and visible light options
- Weather stories and forecasts from the NWS
- Precipitation maps for 24-hour and 7-day periods
- Mobile-responsive design
- Automatic copyright year updates

## Technical Implementation

- Built with Go using the standard library `net/http` package
- Server-side HTML templating with Go's `html/template`
- Uses the "PureCSS" framework for styling with responsive design
- Client-side JavaScript for image error tracking with Sentry
- Comprehensive error handling with environment-aware responses
- Automated testing and deployment pipeline

## Development & Deployment

### CI/CD Pipeline
- GitHub Actions for continuous integration
- Automated unit testing with race detection
- Coverage reporting
- Automatic deployment to Render on successful builds
- UptimeRobot monitoring for downtime detection and automatic instance wake-up

### Local Development
```bash
# Run the test suite (matches CI pipeline)
make test

# Run tests with race detection
make test-race

# Generate coverage report
make coverage
```

## Monitoring & Reliability

- UptimeRobot monitors site availability
- Sentry integration for client-side error tracking
- Server-side error handling with detailed logging in development and user-friendly messages in production
- Automatic instance wake-up on Render to prevent cold starts
- Daily health checks to ensure service availability

## Project Structure

The project follows standard Go project layout conventions:

- `/cmd/server` - Main web server application
- `/internal/handlers` - HTTP handlers specific to this weather site
- `/pkg/templates` - Reusable template system
- `/static` - Static assets (CSS, JavaScript, images)
- `/templates` - HTML templates for the website

## TODO

- [ ] Add integration tests to CI/CD pipeline
- [ ] Implement code linting in CI process
- [ ] Add performance benchmarking
- [x] Enhance error tracking with more detailed metrics
- [ ] Integrate SonarCloud for code quality and coverage analysis

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This site aggregates weather data from various sources. The creator does not claim ownership of any of the radar images or data displayed. This is a personal project made available for public use.

## Contact

For questions or concerns, contact Josh Dutcher at josh.dutcher@joshdutcher.com
