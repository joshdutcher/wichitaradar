<?php

require_once 'includes/init.php';
/*
use Utilities\GetGoesImages;

// get GOES-16 images
$getGoesImages = new GetGoesImages();

// all U.S.
// (image dimension options: 416x250, 625x375, 1250x750, 2500x1500, 5000x3000)
$directoryURL   = 'https://cdn.star.nesdis.noaa.gov/GOES16/ABI/CONUS/GEOCOLOR/';
$imageDimension = '625x375';
$numImages      = 36;

$goesUrlArray = $getGoesImages->getImages($directoryURL, $imageDimension, $numImages);

// Upper Mississippi Valley
// (image dimension options: 300x300, 600x600, 1200x1200)
$directoryURL   = 'https://cdn.star.nesdis.noaa.gov/GOES16/ABI/SECTOR/umv/GEOCOLOR/';
$imageDimension = '600x600';
$numImages      = 36;

$goesUrlUMVArray = $getGoesImages->getImages($directoryURL, $imageDimension, $numImages);
*/
?>
<?php require_once 'includes/header.php';?>

<body>

<div id="layout">
    <?php require_once 'includes/menu.php';?>

    <div class="pure-g">
        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
        	<a href="https://www.star.nesdis.noaa.gov/GOES/GOES16_CONUS_Band.php?band=GEOCOLOR&length=24">
                <?php /*<img class="pure-img-responsive" src="<?php echo $goesUrlArray[0]; ?>" id="goes16" /> */ ?>
                <img class="pure-img-responsive" src="https://cdn.star.nesdis.noaa.gov/GOES16/ABI/GIFS/GOES16-CONUS-GEOCOLOR-625x375.gif" />
            </a>
            <a href="https://www.star.nesdis.noaa.gov/GOES/GOES16_sector_band.php?sector=umv&band=GEOCOLOR&length=24">
                <?php /*<img class="pure-img-responsive" src="<?php echo $goesUrlUMVArray[0]; ?>" id="goes16-umv" />*/ ?>
                <img class="pure-img-responsive" src="https://cdn.star.nesdis.noaa.gov/GOES16/ABI/GIFS/GOES16-UMV-GEOCOLOR-600x600.gif" />
            </a>
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
            <a href="https://www.accuweather.com/en/us/kansas/satellite">
                <img class="pure-img-responsive" src="http://sirocco.accuweather.com/sat_mosaic_640x480_public/ei/isaeks_.gif" />
            </a>
            <a href="http://mp1.met.psu.edu/~fxg1/SAT_SC/anim8vis.html">
                <img class="pure-img-responsive" src="http://mp1.met.psu.edu/~fxg1/SAT_SC/satir_1.gif" id="psu" />
            </a>
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
            <a href="http://www.intellicast.com/National/Satellite/Infrared.aspx">
                <img class="pure-img-responsive" src="http://images.intellicast.com/WxImages/SatelliteLoop/hiusa_None_anim.gif"  />
            </a>
        </div>
    </div>

    <?php require_once 'includes/footer.php';?>
</div>

<script src="js/ui.js"></script>

<script>
	$(function() {
		// mp1.met.psu.edu
		var psu = {
			numImages: 24,
			urlPrefix: 'http://mp1.met.psu.edu/~fxg1/SAT_SC/satir_',
			urlSuffix: '.gif',
            leadingZero: false,
            startingFrame: 1
		}
		$.when(
			preloadImages(psu)
		).then(
			animateFrames(psu, 7, 200, '#psu')
		);

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

        <?php /*
        // GOES East
        var goesUrls = new Array();
        goesUrls.push("<?php echo implode($goesUrlArray, '","'); ?>");
        var goes16 = {
            URLs: goesUrls
        }
        goes16.numImages = goes16.URLs.length;
        $.when(
            preloadImages(goes16)
        ).then(
            animateFrames(goes16, 10, 70, '#goes16')
        );

        // GOES East UMV
        var goesUMVUrls = new Array();
        goesUMVUrls.push("<?php echo implode($goesUrlUMVArray, '","'); ?>");
        var goes16UMV = {
            URLs: goesUMVUrls
        }
        goes16UMV.numImages = goes16UMV.URLs.length;
        $.when(
            preloadImages(goes16UMV)
        ).then(
            animateFrames(goes16UMV, 10, 70, '#goes16-umv')
        );

        // var wgov = {
		// 	frames: 6,
		// 	prefix: 'http://www.weather.gov/images/nws/satellite_images/',
		// 	suffix: '.jpg'
		// }
		// $.when(
		// 	preloadImages(wgov.frames, wgov.prefix, wgov.suffix, false)
		// ).then(
		// 	animateFramesReverse(wgov.frames, 5, 250, wgov.prefix, wgov.suffix, '#wgovsat', false)
		// );
        */ ?>
	});
</script>

</body>
</html>
