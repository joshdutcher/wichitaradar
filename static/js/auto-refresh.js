// Periodically reloads the page so radar/satellite imagery stays current.
//
// Unlike <meta http-equiv="refresh">, this never reloads a hidden tab. If the
// interval elapses while the tab is in the background, the reload waits until
// the user comes back to it.
//
// Timer-driven reloads add ?refresh=1 so the Matomo snippet in
// base.layout.html sends a ping instead of counting another pageview. A reload
// triggered by the user returning to a stale tab is left unflagged: that's a
// genuine return visit (and the Matomo visit may have expired, which would
// make a ping a no-op).
(function () {
  const script = document.currentScript;
  const intervalMs = parseInt(script && script.dataset.interval, 10) * 1000;
  if (!(intervalMs > 0)) {
    return;
  }

  const loadedAt = Date.now();

  function reload(flagAsAutoRefresh) {
    const url = new URL(window.location.href);
    if (flagAsAutoRefresh) {
      url.searchParams.set('refresh', '1');
    }
    window.location.replace(url.toString());
  }

  function isDue() {
    return !document.hidden && Date.now() - loadedAt >= intervalMs;
  }

  setTimeout(() => {
    if (isDue()) {
      reload(true);
    }
  }, intervalMs);

  document.addEventListener('visibilitychange', () => {
    if (isDue()) {
      reload(false);
    }
  });
})();
