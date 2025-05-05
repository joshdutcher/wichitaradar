function preloadImages(fileInfo) {
  var promises = [];
  for (var i = fileInfo.startingFrame; i <= fileInfo.numImages; i++) {
    var frameNum = fileInfo.leadingZero ? ('00' + i).slice(-2) : i;
    var url = fileInfo.urlPrefix + frameNum + fileInfo.urlSuffix;
    promises.push(loadImage(url));
  }
  return Promise.all(promises);
}

function loadImage(url) {
  return new Promise(function (resolve, reject) {
    var img = new Image();
    img.onload = function () {
      resolve(img);
    };
    img.onerror = function () {
      reject(new Error('Failed to load image: ' + url));
    };
    img.src = url;
  });
}

function animateFrames(fileInfo, pauseFrames, frameDelay, imgDomId, reverse) {
  var currentFrame = reverse ? fileInfo.numImages : fileInfo.startingFrame;
  var direction = reverse ? -1 : 1;
  var pauseCounter = 0;
  var isPaused = false;

  function updateFrame() {
    if (isPaused) {
      pauseCounter++;
      if (pauseCounter >= pauseFrames) {
        isPaused = false;
        pauseCounter = 0;
      }
    } else {
      var frameNum = fileInfo.leadingZero ? ('00' + currentFrame).slice(-2) : currentFrame;
      var url = fileInfo.urlPrefix + frameNum + fileInfo.urlSuffix;
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

function reloadAnimatedImage(imgId, interval) {
  const img = document.getElementById(imgId);
  if (!img) return;

  // Store the original relative URL without any timestamps
  const baseUrl = img.getAttribute('src').split('?')[0];

  setInterval(() => {
    // Force reload by replacing the timestamp in the URL
    const timestamp = new Date().getTime();
    img.src = baseUrl + '?t=' + timestamp;
  }, interval);
}
