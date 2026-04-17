// Build an animated loop out of a numbered series of still images.
// Used when the source provides individual frames rather than an animated GIF/APNG.
//
// All frames are preloaded in parallel before animation begins so the first cycle
// doesn't stutter through partially-loaded frames. Once animation starts, the
// displayed <img> is marked with data-animating so image-loading.js stops POSTing
// a /api/image-success per frame swap.
//
// If any frame fails to preload, we dispatch an "error" event on the displayed
// <img> so image-loading.js renders the failed state with the image's alt text.

// eslint-disable-next-line no-unused-vars
function animateFrames(fileInfo, pauseFrames, frameDelay, imgDomId, reverse) {
  const img = document.querySelector(imgDomId);
  if (!img) {
    return;
  }

  const frameUrls = [];
  for (let i = fileInfo.startingFrame; i <= fileInfo.numImages; i++) {
    const frameNum = fileInfo.leadingZero ? ('00' + i).slice(-2) : i;
    frameUrls.push(fileInfo.urlPrefix + frameNum + fileInfo.urlSuffix);
  }

  const preloads = frameUrls.map((url) => {
    return new Promise((resolve, reject) => {
      const preloader = new Image();
      preloader.addEventListener('load', () => resolve(url));
      preloader.addEventListener('error', () => reject(new Error('failed to preload ' + url)));
      preloader.src = url;
    });
  });

  Promise.all(preloads)
    .then(() => {
      img.dataset.animating = 'true';
      startAnimation();
    })
    .catch(() => {
      img.dispatchEvent(new Event('error'));
    });

  function startAnimation() {
    let currentFrame = reverse ? fileInfo.numImages : fileInfo.startingFrame;
    const direction = reverse ? -1 : 1;
    let pauseCounter = 0;
    let isPaused = false;

    function updateFrame() {
      if (isPaused) {
        pauseCounter++;
        if (pauseCounter >= pauseFrames) {
          isPaused = false;
          pauseCounter = 0;
        }
      } else {
        const frameNum = fileInfo.leadingZero ? ('00' + currentFrame).slice(-2) : currentFrame;
        img.src = fileInfo.urlPrefix + frameNum + fileInfo.urlSuffix;

        currentFrame += direction;

        if (reverse) {
          if (currentFrame < fileInfo.startingFrame) {
            currentFrame = fileInfo.numImages;
            isPaused = true;
          }
        } else {
          if (currentFrame > fileInfo.numImages) {
            currentFrame = fileInfo.startingFrame;
            isPaused = true;
          }
        }
      }

      setTimeout(updateFrame, frameDelay);
    }

    updateFrame();
  }
}
