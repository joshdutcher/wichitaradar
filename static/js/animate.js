// Function is called from inline scripts in HTML templates
// eslint-disable-next-line no-unused-vars
function animateFrames(fileInfo, pauseFrames, frameDelay, imgDomId, reverse) {
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
      const url = fileInfo.urlPrefix + frameNum + fileInfo.urlSuffix;
      document.querySelector(imgDomId).src = url;

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
