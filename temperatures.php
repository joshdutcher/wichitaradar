<?php

require_once 'includes/init.php';

?>
<?php require_once 'includes/header.php';?>
<body>

<div id="layout">
    <?php require_once 'includes/menu.php';?>

    <div class="pure-g">
        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
        	<!-- ****** wunderground ****** -->
            <a href="https://www.wunderground.com/maps/temperature/us-current/usa">
                <img class="pure-img-responsive" src="https://s.w-x.co/staticmaps/wu/fee4c/temp_cur/usa/20191002/1500z.jpg" border="0" />
            </a>

			<!-- ****** wunderground dodge city ****** -->
        	<a href="https://www.wunderground.com/maps/temperature/us-current/ddc">
                <img class="pure-img-responsive" src="https://s.w-x.co/staticmaps/wu/fee4c/temp_cur/ddc/20191002/1600z.jpg" />
            </a>
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
			<!-- ****** Weather.gov ***** -->
            <a href="https://graphical.weather.gov/sectors/conus.php#tabs">
                <img class="pure-img-responsive" src="https://graphical.weather.gov/images/conus/MaxT1_conus.png">
            </a>

			<!-- ****** Weather.gov ***** -->
            <a href="https://graphical.weather.gov/sectors/centplains.php#tabs">
                <img class="pure-img-responsive" src="https://graphical.weather.gov/images/centplains/MaxT1_centplains.png">
            </a>

			<!-- ****** Weather.gov ***** -->
            <a href="https://graphical.weather.gov/sectors/kansas.php#tabs">
                <img class="pure-img-responsive" src="https://graphical.weather.gov/images/kansas/MaxT1_kansas.png">
            </a>
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
            <!-- ****** USAirNet ***** -->
            <a href="http://www.usairnet.com/weather/maps/current/current-temperature/">
                <img class="pure-img-responsive whitebg" src="http://www.usairnet.com/weather/images/current-temperature.png" />
            </a>

        	<!-- ****** Weather.com **** -->
        	<a href="https://weather.com/maps/ustemperaturemap">
                <img class="pure-img-responsive" src="https://s.w-x.co/staticmaps/acttemp_1280x720.jpg" />
            </a>
        </div>
    </div>

    <?php require_once 'includes/footer.php';?>
</div>

<script src="js/ui.js"></script>

</body>
</html>
