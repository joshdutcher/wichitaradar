// Initialize Sentry
Sentry.init({
  environment: 'production',
  tracesSampleRate: 1.0,
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
