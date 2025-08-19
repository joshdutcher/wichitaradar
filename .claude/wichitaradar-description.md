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

  - **Backend**: Built in Go using the standard library `net/http` package
  - **Templating**: Server-side HTML templating with Go's `html/template`
  - **Frontend**: Uses PureCSS framework for responsive styling
  - **Error Tracking**: Client-side JavaScript integration with Sentry for error monitoring
  - **Deployment**: Hosted on Railway with automated CI/CD via GitHub Actions
  - **Monitoring**: UptimeRobot monitors availability with automatic instance wake-up

  ## Page Structure

  The site includes multiple specialized pages:

  - **Home**: Main radar dashboard with auto-refresh
  - **Satellite**: Dedicated satellite imagery page
  - **Rainfall**: Precipitation tracking and maps
  - **Temperatures**: Temperature data visualization
  - **Outlook**: Weather outlook and forecasts
  - **About**: Background information about the site and creator

  ## Operational Characteristics

  - Aggregates data from multiple weather sources without claiming ownership
  - Designed for personal use but made publicly available
  - Non-commercial project with no monetization
  - Emphasizes reliability with comprehensive error handling and monitoring
  - Optimized for Chrome and Firefox browsers

  This is essentially a personal weather aggregation tool that has evolved into a community resource, providing a clean,
  focused interface for accessing local weather data without the complexity of larger weather services.