# WichitaRadar

A weather information website serving the Wichita, Kansas area. This project provides real-time weather data, including radar, satellite, temperature maps, and weather stories.

## Project Structure

The project follows standard Go project layout conventions:

- `/cmd/server` - Main web server application
- `/internal/handlers` - HTTP handlers specific to this weather site
- `/pkg/templates` - Reusable template system
- `/static` - Static assets (CSS, JavaScript, images)
- `/templates` - HTML templates for the website

## Features

- Real-time weather radar
- Satellite imagery
- Temperature maps
- Weather stories and outlooks
- Severe weather watches
- Rainfall data
- Weather resources and links

## Technical Details

- Built with Go
- Uses standard library `net/http` for web serving
- HTML templates with layout inheritance
- Responsive design using Pure CSS
- Docker containerization

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Disclaimer

This site aggregates weather data from various sources. The creator does not claim ownership of any of the radar images or data displayed. This is a personal project made available for public use.

## Contact

For questions or concerns, contact Josh Dutcher at josh.dutcher@joshdutcher.com

