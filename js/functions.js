function preloadImages(fileInfo) {
	// get image file names
	var imageFileNames = getImageFileNames(fileInfo);

	// preload images
	var imageArray = new Array();
	for (i=0;i<imageFileNames.length;i++) {
		imageArray[i]     = new Image();
		imageArray[i].src = imageFileNames[i];
	}
}

function getImageFileNames(fileInfo) {
    if (("URLs" in fileInfo) && fileInfo.URLs.length > 0) {
        return fileInfo.URLs;
    }
	var lastImage  = fileInfo.startingFrame + fileInfo.numImages - 1;
	var fileNames  = new Array();
	var imgCounter = 0;
	for (i=fileInfo.startingFrame;i<=lastImage;i++) {
		if (fileInfo.leadingZero && i <= 9) {
			fileNames[imgCounter] = fileInfo.urlPrefix + '0' + i + fileInfo.urlSuffix;
		} else {
			fileNames[imgCounter] = fileInfo.urlPrefix + i + fileInfo.urlSuffix;
		}
		imgCounter++;
	}
	return fileNames;
}

function animateFrames(fileInfo, pauseFrames, frameDelay, imgDomId, reverse=false) {
    var imageFileNames = getImageFileNames(fileInfo);
    if (imageFileNames.length == 0) {
        return;
    }
	if (reverse) {
		imageFileNames = imageFileNames.reverse();
	}

	var totalFrames = fileInfo.numImages + pauseFrames;
	var frameTimer = 0;
	var lastImageIndex = imageFileNames.length-1;

    if (imageFileNames.length == 1) {
        // if there's only one image there's no need to animate/rotate them
        $(imgDomId).attr('src', imageFileNames[0]);
        return;
    } else {
        // if we got here, imageFileNames.length is > 1
        setInterval(function() {
    		if (frameTimer <= lastImageIndex) {
    			showFrame = frameTimer;
    		} else {
    			showFrame = lastImageIndex;
    		}

    		// $(imgDomId+'_frametimer').html(frameTimer);
    		// $(imgDomId+'_showframe').html(showFrame);
    		// $(imgDomId+'_url').html(imageFileNames[showFrame]);

    		$(imgDomId).attr('src', imageFileNames[showFrame]);

    		if (frameTimer++ == totalFrames-1) {
    			frameTimer = 0;
    		}
    	}, frameDelay);
    }
}
