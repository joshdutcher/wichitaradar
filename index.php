<?php

require_once 'includes/init.php';

/************************************
css & responsive framework:
purecss.io
color palette:
http://paletton.com/#uid=73y0u0k+bFooMYfHB+5YtpFVIew
favicon generator:
http://www.favicomatic.com/

TODO:
get all assets the same, either local or remote
https://github.com/emartinez-usgs/earthquake-widget

 *************************************/

?>
<?php require_once 'includes/header.php';?>
<body>

<div id="layout">
    <?php require_once 'includes/menu.php';?>

    <div class="pure-g">
        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
        	<!-- ****** Wunderground (old Intellicast) radar ****** -->
            <a href="https://www.wunderground.com/maps/radar/current/sln">
                <img class="pure-img-responsive" src="https://s.w-x.co/staticmaps/wu/wxtype/county_loc/sln/animate.png" border="0" />
            </a>

			<!-- ****** Accuweather radar ****** -->
        	<a href="https://www.accuweather.com/en/us/kansas/weather-radar?play=1">
                <img class="pure-img-responsive" src="http://sirocco.accuweather.com/nx_mosaic_640x480_public/sir/inmasirks_.gif" />
            </a>

            <!-- ****** Wunderground radar ****** -->
            <a href="https://www.wunderground.com/radar/us/ks/salina">
                <img class="pure-img-responsive" src="https://s.w-x.co/staticmaps/wu/wu/wxtype1200_cur/ussln/animate.png" />
            </a>
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
            <!-- ****** KSN KS radar ***** -->
            <a href="http://www.ksn.com/weather/images/kansas-radar">
                <img class="pure-img-responsive" src="https://media.psg.nexstardigital.net/ksnw/weather/images/ksn_ks_radar_8.jpg" id="ksnKSLoop" />
            </a>

        	<!-- ****** KSN Southcentral radar **** -->
        	<a href="http://www.ksn.com/weather/images/southcentral-kansas-radar">
                <img class="pure-img-responsive" src="https://media.psg.nexstardigital.net/ksnw/weather/images/ksn_sc_radar_8.jpg" id="ksnSCKSLoop" />
            </a>

            <!-- ****** KSN Wichita radar **** -->
            <a href="http://www.ksn.com/weather/images/wichita-radar">
                <img class="pure-img-responsive" src="https://media.psg.nexstardigital.net/ksnw/weather/images/ksn_wichita_radar_8.jpg" id="ksnWichitaLoop" />
            </a>
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
            <!-- ****** weather.gov radar loop ***** -->
            <img class="pure-img-responsive" src="https://radar.weather.gov/ridge/standard/KICT_loop.gif" />
        </div>
    </div>
</div>

<?php require_once 'includes/footer.php';?>

<script>
	$(function() {
        // KSN KS RADAR
        var newksnks = {
            numImages: 8,
            urlPrefix: 'https://media.psg.nexstardigital.net/ksnw/weather/images/ksn_ks_radar_',
            urlSuffix: '.jpg',
            leadingZero: false,
            startingFrame: 1
        }
        $.when(
            preloadImages(newksnks)
        ).then(
            // function animateFrames(fileInfo, pauseFrames, frameDelay, imgDomId, reverse=false) {
            animateFrames(newksnks, 5, 300, '#ksnKSLoop', true)
        );

        // KSN SC KS RADAR
        var newksn = {
            numImages: 8,
            urlPrefix: 'https://media.psg.nexstardigital.net/ksnw/weather/images/ksn_sc_radar_',
            urlSuffix: '.jpg',
            leadingZero: false,
            startingFrame: 1
        }
        $.when(
            preloadImages(newksn)
        ).then(
            animateFrames(newksn, 5, 300, '#ksnSCKSLoop', true)
        );

        // KSN WICHITA LOOP
        var ksnWichita = {
            numImages: 8,
            urlPrefix: 'https://media.psg.nexstardigital.net/ksnw/weather/images/ksn_wichita_radar_',
            urlSuffix: '.jpg',
            leadingZero: false,
            startingFrame: 1
        }
        $.when(
            preloadImages(ksnWichita)
        ).then(
            animateFrames(ksnWichita, 5, 300, '#ksnWichitaLoop', true)
        );
	});
</script>

</body>
</html>
