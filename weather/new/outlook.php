<?php require_once('includes/init.php'); ?>
<?php require_once('includes/header.php'); ?>
<body>

<div id="layout">
    <?php require_once('includes/menu.php'); ?>

    <div class="pure-g" id="mainbody">
        <div class="pure-u-1 pure-u-md-1-2 pure-u-lg-1-3">
        	<div id="accuweather_sat">
				<img class="pure-img-responsive" src="http://www.weather.gov/images/ict/WxStory/FileL.png" />
			</div>
        </div>

        <div class="pure-u-1 pure-u-md-1-2 pure-u-lg-1-3">
        	convective outlook
        	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day1otlk_1300.gif" />
        	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day2otlk_0600.gif" />
        	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day3otlk_0730.gif" />
        </div>
    </div>

    <?php require_once('includes/footer.php'); ?>
</div>

<script src="js/ui.js"></script>

</body>
</html>
