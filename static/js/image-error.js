// Initialize Sentry
Sentry.init({
  environment:
    window.location.hostname === 'localhost' || window.location.hostname === '127.0.0.1' ? 'development' : 'production',
});

// Track image loading errors
document.addEventListener('DOMContentLoaded', function () {
  const images = document.getElementsByTagName('img');

  for (let img of images) {
    img.addEventListener('error', function (e) {
      const errorData = {
        src: this.src,
        alt: this.alt || 'No alt text',
        page: window.location.pathname,
        timestamp: new Date().toISOString(),
      };

      // Capture error in Sentry
      Sentry.captureMessage('Image failed to load', {
        level: 'error',
        tags: {
          page: errorData.page,
          image: errorData.src,
        },
        extra: errorData,
      });
    });
  }
});
