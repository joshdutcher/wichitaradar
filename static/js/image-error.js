// Initialize Sentry
Sentry.init({
  environment:
    window.location.hostname === 'localhost' || window.location.hostname === '127.0.0.1' ? 'development' : 'production',
});

// Track image loading errors
function attachImageErrorListener(img) {
  // Only attach if not already attached
  if (img.dataset.errorListenerAttached) return;

  img.addEventListener('error', function (e) {
    // Add a small delay to allow for potential recovery
    setTimeout(() => {
      // Check if the image is still in an error state
      if (img.naturalWidth === 0) {
        const errorData = {
          src: this.src,
          alt: this.alt || 'No alt text',
          page: window.location.pathname,
          timestamp: new Date().toISOString(),
          naturalWidth: img.naturalWidth,
          naturalHeight: img.naturalHeight,
          complete: img.complete,
          crossOrigin: img.crossOrigin || 'none',
        };

        // Capture error in Sentry
        Sentry.captureMessage('Image failed to load', {
          level: 'error',
          tags: {
            page: errorData.page,
            image: errorData.src,
            crossOrigin: errorData.crossOrigin,
          },
          extra: errorData,
        });
      }
    }, 1000); // 1 second delay
  });

  img.dataset.errorListenerAttached = 'true';
}

// Handle initial images
document.addEventListener('DOMContentLoaded', function () {
  const images = document.getElementsByTagName('img');
  for (let img of images) {
    attachImageErrorListener(img);
  }
});

// Handle dynamically added images
const observer = new MutationObserver(mutations => {
  mutations.forEach(mutation => {
    mutation.addedNodes.forEach(node => {
      if (node.nodeName === 'IMG') {
        attachImageErrorListener(node);
      }
      // Also check for images within added nodes
      if (node.getElementsByTagName) {
        const images = node.getElementsByTagName('img');
        for (let img of images) {
          attachImageErrorListener(img);
        }
      }
    });
  });
});

observer.observe(document.body, {
  childList: true,
  subtree: true,
});
