<?php require_once 'includes/init.php'; ?>
<?php require_once 'includes/header.php';?>

<body>

<div id="layout">
    <?php require_once 'includes/menu.php';?>

    <div class="pure-g">
        <div class="pure-u pure-u-1 pure-u-md-1-1 pure-u-lg-1-2">
        	<a href="https://www.star.nesdis.noaa.gov/GOES/GOES16_CONUS_Band.php?band=GEOCOLOR&length=24">
                <?php /*<img class="pure-img-responsive" src="<?php echo $goesUrlArray[0]; ?>" id="goes16" /> */ ?>
                <img class="pure-img-responsive" src="https://cdn.star.nesdis.noaa.gov/GOES16/ABI/GIFS/GOES16-CONUS-GEOCOLOR-625x375.gif" />
            </a>
            <a href="https://www.star.nesdis.noaa.gov/GOES/GOES16_sector_band.php?sector=umv&band=GEOCOLOR&length=24">
                <?php /*<img class="pure-img-responsive" src="<?php echo $goesUrlUMVArray[0]; ?>" id="goes16-umv" />*/ ?>
                <img class="pure-img-responsive" src="https://cdn.star.nesdis.noaa.gov/GOES16/ABI/GIFS/GOES16-UMV-GEOCOLOR-600x600.gif" />
            </a>
        </div>

        <div class="pure-u pure-u-1 pure-u-md-1-1 pure-u-lg-1-2">
            <a href="https://www.wunderground.com/maps/satellite/regional-infrared/usace">
                <img class="pure-img-responsive" src="https://s.w-x.co/staticmaps/wu/wu/satir1200_cur/usace/animate.png"  />
            </a>
            <a href="https://www.accuweather.com/en/us/kansas/satellite">
                <img class="pure-img-responsive" src="http://sirocco.accuweather.com/sat_mosaic_640x480_public/ei/isaeks_.gif" />
            </a>
        </div>
    </div>
</div>

<?php require_once 'includes/footer.php';?>

<script>
	$(function() {
		// weather.gov
        var wgov = {
            numImages: 6,
            urlPrefix: 'http://www.weather.gov/images/nws/satellite_images/',
            urlSuffix: '.jpg',
            leadingZero: false,
            startingFrame: 6
        }
        $.when(
            preloadImages(wgov)
        ).then(
            animateFrames(wgov, 5, 400, '#wgovsat', true)
        );
	});
</script>

</body>
</html>
