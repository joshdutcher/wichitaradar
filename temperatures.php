<?php

require_once 'includes/init.php';

use Utilities\GetImage;
use Utilities\SWXCOFiles;

// figure out the date and timestamps for the wunderground files and make sure they exist
$swxco = new SWXCOFiles();
$swxcofiles = $swxco->getImagePaths();

?>
<?php require_once 'includes/header.php';?>
<body>

<div id="layout">
    <?php require_once 'includes/menu.php';?>

    <div class="pure-g">
        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
			<!-- ****** KSNW ***** -->
            <a href="https://www.ksn.com/weather/images/kansas-temps/">
                <img class="pure-img-responsive" src="https://media.psg.nexstardigital.net/ksnw/weather/images/ksnow_full.jpg" border="0" />
            </a>

            <!-- ****** Weather.gov ***** -->
            <a href="https://graphical.weather.gov/sectors/kansas.php#tabs">
                <img class="pure-img-responsive" src="https://graphical.weather.gov/images/kansas/MaxT1_kansas.png" border="0" />
            </a>
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
			<?php if ($swxcofiles['ddc'] != '') { ?>
            <!-- ****** wunderground dodge city ****** -->
        	<a href="https://www.wunderground.com/maps/temperature/us-current/ddc">
                <img class="pure-img-responsive" src="<?php echo $swxcofiles['ddc']; ?>" border="0" />
            </a>
            <?php } ?>

			<!-- ****** Weather.gov ***** -->
            <a href="https://graphical.weather.gov/sectors/centplains.php#tabs">
                <img class="pure-img-responsive" src="https://graphical.weather.gov/images/centplains/MaxT1_centplains.png" border="0" />
            </a>
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
        	<!-- ****** Weather.com **** -->
        	<a href="https://weather.com/maps/ustemperaturemap">
                <img class="pure-img-responsive" src="https://s.w-x.co/staticmaps/acttemp_1280x720.jpg" border="0" />
            </a>

            <!-- ****** wunderground ****** -->
            <?php if ($swxcofiles['usa'] != '') { ?>
            <a href="https://www.wunderground.com/maps/temperature/us-current/usa">
                <img class="pure-img-responsive" src="<?php echo $swxcofiles['usa']; ?>" border="0" />
            </a>
            <?php } ?>

			<!-- ****** Weather.gov ***** -->
            <a href="https://graphical.weather.gov/sectors/conus.php#tabs">
                <img class="pure-img-responsive" src="https://graphical.weather.gov/images/conus/MaxT1_conus.png" border="0" />
            </a>
        </div>
    </div>
</div>

<?php require_once 'includes/footer.php';?>

</body>
</html>
