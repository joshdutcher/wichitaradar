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
		//$(div).html('<img src="' + url + '" class="pure-img-responsive" />');
		$(img).attr('src', url);
	}, delay);
}

function preloadImages(frames, prefix, suffix, leadingZero) {
	// preload images
	console.log('frames: ' + frames + ' prefix: ' + prefix + ' suffix: ' + suffix);
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