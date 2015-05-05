function rightNow() {
	if (window['performance'] && window['performance']['now']) {
		return window['performance']['now']();
	} else {
		return +(new Date());
	}
}


function animateLoop(animFrames, delayFrames, animSpeed, targetElm) {
	// animFrames = 12
	// delayFrames = 6
	// animSpeed = 150ms
	// targetElm = $('#myImage')
	totalFrames = animFrames + delayFrames;
	for (i=1;i<totalFrames;i++) {

	}

}


function goAnimate() {
	var fps          = 3,
		currentFrame = 0,
		totalFrames  = 11,
		img          = document.getElementById("myImage"),
		currentTime  = rightNow();

	(function animloop(time){
		var delta = (time - currentTime) / 1000;

		currentFrame += (delta * fps);

		var frameNum = Math.floor(currentFrame);

		if (frameNum >= totalFrames) {
			currentFrame = frameNum = 0;
		}

		requestAnimationFrame(animloop);

		//img.src = "http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_" +

		img.src = "http://wx.ksn.com/weather/images/ksn_ks_radar_" +
		(frameNum < 10 ? "0" : "") + frameNum + ".jpg";

		currentTime = time;
	})(currentTime);
}