// Cache of currently failing images to minimize API calls
let failingImages = new Set();
let lastFailingImagesUpdate = 0;
const FAILING_IMAGES_CACHE_TTL = 60000; // 1 minute

// Fetch list of currently failing images
async function updateFailingImages() {
  const now = Date.now();
  if (now - lastFailingImagesUpdate < FAILING_IMAGES_CACHE_TTL) {
    return; // Use cached data
  }

  try {
    const response = await fetch('/api/image-failure-status');
    if (response.ok) {
      const failingImageUrls = await response.json();
      failingImages = new Set(failingImageUrls);
      lastFailingImagesUpdate = now;
    }
  } catch (err) {
    console.error('Failed to fetch failing images status:', err);
  }
}

function attachImageListeners(img) {
  // Only attach if not already attached
  if (img.dataset.listenersAttached) return;

  // Track errors
  img.addEventListener('error', function (e) {
    const errorData = {
      src: this.src,
      alt: this.alt || 'No alt text',
      page: window.location.pathname,
      timestamp: new Date().toISOString(),
    };

    // Add to local failing images cache
    failingImages.add(this.src);

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

  // Track successful loads - but only for images that were failing
  img.addEventListener('load', async function (e) {
    // Update failing images cache if it's stale
    await updateFailingImages();
    
    // Only report success if this image was in the failing list
    if (failingImages.has(this.src)) {
      const successData = {
        src: this.src,
        page: window.location.pathname,
        timestamp: new Date().toISOString(),
      };

      // Remove from local cache since it's now working
      failingImages.delete(this.src);

      // Report success to server
      fetch('/api/image-success', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify(successData),
      }).catch(err => {
        console.error('Failed to report image success:', err);
      });
    }
  });

  img.dataset.listenersAttached = 'true';
}

// Handle initial images and set up observer
document.addEventListener('DOMContentLoaded', async function () {
  // Fetch initial failing images status
  await updateFailingImages();
  
  // Handle initial images
  const images = document.getElementsByTagName('img');
  for (let img of images) {
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
          for (let img of images) {
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
