<?php

require_once('includes/init.php');

/************************************
css & responsive framework:
	purecss.io
color palette:
	http://paletton.com/#uid=73y0u0k+bFooMYfHB+5YtpFVIew
favicon generator:
	http://www.favicomatic.com/

TO DO:
	get all assets the same, either local or remote
	hourly forecast
	http://www.spc.noaa.gov/products/activity_loop.gif

*************************************/

?>

<?php require_once('includes/header.php'); ?>
<body>

<div id="layout">
    <?php require_once('includes/menu.php'); ?>

    <div class="pure-g">
        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
        	<!-- ****** Intellicast radar ****** -->
        	<img class="pure-img-responsive" src="http://images.intellicast.com/WxImages/RadarLoop/sln_None_anim.gif" />

			<!-- ****** Accuweather radar ****** -->
        	<img class="pure-img-responsive" src="http://sirocco.accuweather.com/nx_mosaic_640x480_public/sir/inmasirks_.gif" />

        	<!-- ****** TWC Doppler ************ -->
        	<img class="pure-img-responsive" src="http://image.weather.com/looper/archive/us_ddc_closeradar_large_usen/1L.jpg" id="weathercom" />
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
			<!-- ****** KAKE Doppler SCKS ****** -->
        	<img class="pure-img-responsive" src="http://gray.ftp.clickability.com/kakewebftp/wx-radar-Zone-SC.gif" />

			<!-- ****** KAKE Doppler ICT ******* -->
        	<img class="pure-img-responsive" src="http://gray.ftp.clickability.com/kakewebftp/wx-radar-Wichita.gif" />

        	<!-- ****** KSN Pinpoint ********** -->
        	<img class="pure-img-responsive" src="http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_01.jpg" id="ksnLoop" />

        	<!-- ****** KSN Southcentral radar * -->
        	<img class="pure-img-responsive" src="http://wx.ksn.com/weather/images/ksn_sc_radar_01.jpg" id="newKsnLoop" />
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
			<!-- ****** Wunderground radar ***** -->
            <img class="pure-img-responsive" src="http://radblast.wunderground.com/cgi-bin/radar/WUNIDS_map?station=ICT&brand=wui&num=6&delay=15&type=N0R&frame=0&scale=1.000&noclutter=0&showstorms=0&mapx=400&mapy=240&centerx=400&centery=240&transx=0&transy=0&showlabels=1&severe=0&rainsnow=0&lightning=0&smooth=0&rand=23867719&lat=0&lon=0&label=you">

			<!-- * NOAA base reflectivity loop * -->
			<img class="pure-img-responsive" src="http://radar.weather.gov/lite/NCR/ICT_loop.gif">
        </div>
    </div>

    <?php require_once('includes/footer.php'); ?>
</div>

<script src="js/ui.js"></script>

<script>
	$(function() {
		// KSN RADAR
		var ksn = {
			frames: 12,
			prefix: 'http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_',
			suffix: '.jpg'
		}
		$.when(
			preloadImages(ksn.frames, ksn.prefix, ksn.suffix, true)
		).then(
			animateFrames(ksn.frames, 5, 150, ksn.prefix, ksn.suffix, '#ksnLoop', true)
		);

		// NEW KSN RADAR
		var newksn = {
			frames: 12,
			prefix: 'http://wx.ksn.com/weather/images/ksn_sc_radar_',
			suffix: '.jpg'
		}
		$.when(
			preloadImages(newksn.frames, newksn.prefix, newksn.suffix, true)
		).then(
			animateFrames(newksn.frames, 5, 200, newksn.prefix, newksn.suffix, '#newKsnLoop', true)
		);

		// WEATHER.COM
		var weathercom = {
			frames: 5,
			prefix: 'http://image.weather.com/looper/archive/us_ddc_closeradar_large_usen/',
			suffix: 'L.jpg'
		}
		$.when(
			preloadImages(weathercom.frames, weathercom.prefix, weathercom.suffix, false)
		).then(
			animateFrames(weathercom.frames, 2, 300, weathercom.prefix, weathercom.suffix, '#weathercom', false)
		);
	});
</script>

</body>
</html>
