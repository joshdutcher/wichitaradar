// Populate .wu-slot elements with Weather Underground temperature images.
//
// A slot can declare either:
//   data-wu-src = direct URL (load immediately)
//   data-wu-key = fetch the live URL from /api/wu-temperature-images, keyed by this name
//
// The slot acts as the visual container for the loading / failed state while the
// URL is being resolved. Once an <img> is injected into the slot, image-loading.js
// takes over responsibility for the image's own load / failure lifecycle.

(function () {
  const slots = document.querySelectorAll('.wu-slot');
  if (slots.length === 0) {
    return;
  }

  const LOADING_DELAY_MS = 500;
  const REQUEST_TIMEOUT_MS = 5000;

  const apiSlots = [];
  slots.forEach((slot) => {
    if (slot.dataset.wuSrc) {
      mountImage(slot, slot.dataset.wuSrc);
    } else if (slot.dataset.wuKey) {
      apiSlots.push(slot);
    }
  });

  if (apiSlots.length > 0) {
    resolveUrlsAndMount(apiSlots);
  }

  function resolveUrlsAndMount(targetSlots) {
    const loadingTimer = setTimeout(() => {
      targetSlots.forEach((s) => {
        if (!s.dataset.state) {
          s.classList.add('img-loading', 'img-wrapper-responsive');
        }
      });
    }, LOADING_DELAY_MS);

    const controller = new AbortController();
    const abortTimer = setTimeout(() => {
      controller.abort();
    }, REQUEST_TIMEOUT_MS);

    fetch('/api/wu-temperature-images', { signal: controller.signal })
      .then((res) => {
        if (!res.ok) {
          throw new Error('bad status ' + res.status);
        }
        return res.json();
      })
      .then((data) => {
        clearTimeout(loadingTimer);
        clearTimeout(abortTimer);
        targetSlots.forEach((s) => {
          const url = data[s.dataset.wuKey];
          if (url) {
            mountImage(s, url);
          } else {
            markFailed(s);
          }
        });
      })
      .catch(() => {
        clearTimeout(loadingTimer);
        clearTimeout(abortTimer);
        targetSlots.forEach((s) => {
          if (s.dataset.state !== 'loaded') {
            markFailed(s);
          }
        });
      });
  }

  function mountImage(slot, url) {
    slot.dataset.state = 'loaded';
    slot.classList.remove('img-loading', 'img-failed');
    const href = slot.dataset.wuHref || '#';
    const link = document.createElement('a');
    link.href = href;
    const img = document.createElement('img');
    img.className = 'pure-img-responsive';
    if (slot.dataset.wuAlt) {
      img.alt = slot.dataset.wuAlt;
    }
    img.src = url;
    link.appendChild(img);
    slot.replaceChildren(link);
  }

  function markFailed(slot) {
    slot.dataset.state = 'failed';
    slot.classList.remove('img-loading');
    slot.dataset.failedLabel = buildFailedLabel(slot.dataset.wuAlt);
    slot.classList.add('img-failed', 'img-wrapper-responsive');
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
})();
