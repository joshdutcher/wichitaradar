// Allowlist of hosts we actually care about
const ALLOWED_HOSTS = new Set([
  'wichitaradar.com',
  'www.wichitaradar.com',
  // External providers embedded in the app
  'img.shields.io',
  'www.spc.noaa.gov',
  's.w-x.co',
  'sirocco.accuweather.com',
  'media.psg.nexstardigital.net',
  'graphical.weather.gov',
  'www.weather.gov',
  'weather.gov',
  'radar.weather.gov',
  'x-hv1.pivotalweather.com',
  'cdn.star.nesdis.noaa.gov',
]);

// Validate that an image is from our domain and not a tracking pixel
function shouldReportImage(img) {
  try {
    // Use currentSrc (actual loaded src) or fall back to src attribute
    const srcUrl = img.currentSrc || img.src;
    if (!srcUrl) {
      return false;
    }

    // Parse URL to get hostname
    const u = new URL(srcUrl);

    // Check if hostname is in our allowlist
    if (!ALLOWED_HOSTS.has(u.hostname)) {
      return false; // Not our domain, ignore it
    }

    // Optional: Skip 1×1 pixels (often tracking pixels)
    if (img.naturalWidth === 1 && img.naturalHeight === 1) {
      return false;
    }

    return true; // Passed all checks
  } catch (err) {
    // Ignore bad URLs (invalid, relative, etc.)
    return false;
  }
}

function attachImageListeners(img) {
  // Only attach if not already attached
  if (img.dataset.listenersAttached) {
    return;
  }

  // Track errors
  img.addEventListener('error', function () {
    // Validate image before reporting
    if (!shouldReportImage(this)) {
      return; // Ignore third-party images, tracking pixels, etc.
    }

    const errorData = {
      src: this.src,
      alt: this.alt || 'No alt text',
      page: window.location.pathname,
      timestamp: new Date().toISOString(),
      width: this.naturalWidth,
      height: this.naturalHeight,
    };

    // Report to server
    fetch('/api/image-error', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(errorData),
    }).catch(err => {
      console.error('Failed to report image error:', err);
    });
  });

  // Track successful loads
  img.addEventListener('load', function () {
    // Validate image before reporting
    if (!shouldReportImage(this)) {
      return; // Ignore third-party images, tracking pixels, etc.
    }

    const successData = {
      src: this.src,
    };

    // Report success to server (server tracks which images were previously failing)
    fetch('/api/image-success', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(successData),
    }).catch(err => {
      console.error('Failed to report image success:', err);
    });
  });

  img.dataset.listenersAttached = 'true';
}

// Handle initial images and set up observer
document.addEventListener('DOMContentLoaded', () => {
  // Handle initial images
  const images = document.getElementsByTagName('img');
  for (const img of images) {
    attachImageListeners(img);
  }

  // Handle dynamically added images
  const observer = new MutationObserver(mutations => {
    mutations.forEach(mutation => {
      mutation.addedNodes.forEach(node => {
        if (node.nodeName === 'IMG') {
          attachImageListeners(node);
        }
        // Also check for images within added nodes
        if (node.getElementsByTagName) {
          const images = node.getElementsByTagName('img');
          for (const img of images) {
            attachImageListeners(img);
          }
        }
      });
    });
  });

  // Now we can safely observe document.body
  observer.observe(document.body, {
    childList: true,
    subtree: true,
  });
});
