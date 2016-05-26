function preloadImages(frames, prefix, suffix, leadingZero) {
	// preload images
	isn=new Array();
	for (i=1;i<frames+1;i++) {
		isn[i]=new Image();
		if (i<frames){
			// Use 0 placeholder in filename when i is less than 10
			if (i<10 && leadingZero){
				isn[i].src=prefix + '0' + i + suffix;
			} else {
				isn[i].src=prefix + i + suffix;
			}
		}
	}
}

function animateFrames(frames, pauseFrames, delay, prefix, suffix, img, leadingZero) {
	var currentFrame = 1,
		totalFrames = frames + pauseFrames;

	setInterval(function() {
		if (currentFrame <= totalFrames) {
			if (currentFrame <= frames) {
				if (leadingZero) {
					showFrame = (currentFrame < 10 ? "0" : "") + currentFrame;
				} else {
					showFrame = currentFrame;
				}
			} else {
				showFrame = frames;
			}
		}
		currentFrame = currentFrame % totalFrames + 1;

		var url = prefix + showFrame + suffix;
		$(img).attr('src', url);
	}, delay);
}

function animateFramesReverse(frames, pauseFrames, delay, prefix, suffix, img, leadingZero) {
	var totalFrames = frames + pauseFrames;
	var currentFrame = frames;
	var showFrame = frames;

	setInterval(function() {
		if (currentFrame > 0) {
			if (currentFrame <= frames) {
				if (leadingZero) {
					showFrame = (currentFrame < 10 ? "0" : "") + currentFrame;
				} else {
					showFrame = currentFrame;
				}
			} else {
				showFrame = 1;
			}
			currentFrame--;
		}
		if (currentFrame == 0) {
			currentFrame = totalFrames;
		}

		var url = prefix + showFrame + suffix;
		$(img).attr('src', url);
	}, delay);
}
