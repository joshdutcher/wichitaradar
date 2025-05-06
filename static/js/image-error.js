function attachImageErrorListener(img) {
  // Only attach if not already attached
  if (img.dataset.errorListenerAttached) return;

  img.addEventListener('error', function (e) {
    const errorData = {
      src: this.src,
      alt: this.alt || 'No alt text',
      page: window.location.pathname,
      timestamp: new Date().toISOString(),
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

  img.dataset.errorListenerAttached = 'true';
}

// Handle initial images and set up observer
document.addEventListener('DOMContentLoaded', function () {
  // Handle initial images
  const images = document.getElementsByTagName('img');
  for (let img of images) {
    attachImageErrorListener(img);
  }

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

  // Now we can safely observe document.body
  observer.observe(document.body, {
    childList: true,
    subtree: true,
  });
});
