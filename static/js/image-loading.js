// Image lifecycle handler: applied to every <img> on the page.
//
// Behaviors:
//   * Show a shimmer placeholder if the image hasn't loaded within LOADING_DELAY_MS.
//   * Show a "failed" placeholder and report an error to the server after LOAD_TIMEOUT_MS,
//     even if the browser's own `error` event hasn't fired (covers hung/unresponsive hosts).
//   * Report success / failure to /api/image-success and /api/image-error.
//   * Allowlist gating for reports stays identical to the prior image-error.js behavior.
//
// Opt-outs:
//   * Images inside #menu keep reporting but skip the visual wrapping.
//   * Images with class "no-loading-indicator" keep reporting but skip the visual wrapping.

const ALLOWED_HOSTS = new Set([
  'wichitaradar.com',
  'www.wichitaradar.com',
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

const LOADING_DELAY_MS = 500;
const LOAD_TIMEOUT_MS = 5000;

function shouldReportImage(img) {
  try {
    const srcUrl = img.currentSrc || img.src;
    if (!srcUrl) {
      return false;
    }
    const u = new URL(srcUrl, window.location.href);
    if (!ALLOWED_HOSTS.has(u.hostname)) {
      return false;
    }
    if (img.naturalWidth === 1 && img.naturalHeight === 1) {
      return false;
    }
    return true;
  } catch (err) {
    return false;
  }
}

function reportImageError(img, reason) {
  if (!shouldReportImage(img)) {
    return;
  }
  const payload = {
    src: img.src,
    alt: img.alt || 'No alt text',
    page: window.location.pathname,
    timestamp: new Date().toISOString(),
    width: img.naturalWidth,
    height: img.naturalHeight,
    error: reason || '',
  };
  fetch('/api/image-error', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify(payload),
  }).catch(err => {
    console.error('Failed to report image error:', err);
  });
}

function reportSuccess(img) {
  if (!shouldReportImage(img)) {
    return;
  }
  fetch('/api/image-success', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: JSON.stringify({ src: img.src }),
  }).catch(err => {
    console.error('Failed to report image success:', err);
  });
}

function shouldEnhanceVisual(img) {
  if (img.classList.contains('no-loading-indicator')) {
    return false;
  }
  if (img.closest('#menu')) {
    return false;
  }
  return true;
}

function wrapImage(img) {
  if (img.parentElement && img.parentElement.classList.contains('img-wrapper')) {
    return img.parentElement;
  }
  const wrapper = document.createElement('span');
  wrapper.className = 'img-wrapper';
  if (img.classList.contains('pure-img-responsive')) {
    wrapper.classList.add('img-wrapper-responsive');
  }
  const parent = img.parentNode;
  if (!parent) {
    return null;
  }
  parent.insertBefore(wrapper, img);
  wrapper.appendChild(img);
  return wrapper;
}

function setLoaded(img, wrapper) {
  img.dataset.imgState = 'loaded';
  if (wrapper) {
    wrapper.classList.remove('img-loading', 'img-failed');
  }
}

function buildFailedLabel(description) {
  const base = 'Image unavailable';
  if (!description) {
    return base;
  }
  const trimmed = String(description).trim();
  if (!trimmed) {
    return base;
  }
  return base + ': ' + trimmed;
}

function setFailed(img, wrapper) {
  img.dataset.imgState = 'failed';
  if (wrapper) {
    wrapper.classList.remove('img-loading');
    wrapper.dataset.failedLabel = buildFailedLabel(img.alt);
    wrapper.classList.add('img-failed');
  }
}

function attach(img) {
  if (img.dataset.imgEnhanced) {
    return;
  }
  img.dataset.imgEnhanced = 'true';

  const enhanceVisual = shouldEnhanceVisual(img);
  const wrapper = enhanceVisual ? wrapImage(img) : null;

  // If the image already finished loading (e.g. cached) before we got here,
  // emit the appropriate terminal state once and move on.
  if (img.complete) {
    if (img.naturalWidth > 0) {
      setLoaded(img, wrapper);
      reportSuccess(img);
    } else {
      setFailed(img, wrapper);
      reportImageError(img, 'complete-with-no-natural-dimensions');
    }
    return;
  }

  let loadingTimer = null;
  let failTimer = null;

  if (enhanceVisual && wrapper) {
    loadingTimer = setTimeout(() => {
      if (img.dataset.imgState) {
        return;
      }
      img.dataset.imgState = 'loading';
      wrapper.classList.add('img-loading');
    }, LOADING_DELAY_MS);
  }

  failTimer = setTimeout(() => {
    if (img.dataset.imgState === 'loaded') {
      return;
    }
    setFailed(img, wrapper);
    reportImageError(img, 'timeout');
  }, LOAD_TIMEOUT_MS);

  img.addEventListener('load', () => {
    if (loadingTimer) {
      clearTimeout(loadingTimer);
    }
    if (failTimer) {
      clearTimeout(failTimer);
    }
    setLoaded(img, wrapper);
    // Animated frame-swapping (animate.js) fires load on every frame.
    // Report success only when the image isn't actively animating, so we
    // don't flood /api/image-success with one request per frame swap.
    if (!img.dataset.animating) {
      reportSuccess(img);
    }
  });

  img.addEventListener('error', () => {
    if (loadingTimer) {
      clearTimeout(loadingTimer);
    }
    if (failTimer) {
      clearTimeout(failTimer);
    }
    if (img.dataset.imgState === 'failed') {
      return;
    }
    setFailed(img, wrapper);
    reportImageError(img, 'load-error');
  });
}

function attachToSubtree(root) {
  if (root.nodeName === 'IMG') {
    attach(root);
    return;
  }
  if (!root.getElementsByTagName) {
    return;
  }
  const images = root.getElementsByTagName('img');
  for (const img of images) {
    attach(img);
  }
}

document.addEventListener('DOMContentLoaded', () => {
  attachToSubtree(document);

  const observer = new MutationObserver(mutations => {
    mutations.forEach(mutation => {
      mutation.addedNodes.forEach(node => {
        attachToSubtree(node);
      });
    });
  });
  observer.observe(document.body, { childList: true, subtree: true });
});
