<?php require_once('includes/init.php'); ?>
<?php require_once('includes/header.php'); ?>
<body>

<style>
	#noaa_watches {
		background-color: #FFF;
		text-align: center;
	}
</style>

<div id="layout">
    <?php require_once('includes/menu.php'); ?>

    <div class="pure-g" id="mainbody">
        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
			<!-- ****** Accuweather ************ -->
			<img class="pure-img-responsive" src="http://sirocco.accuweather.com/adc_images2/english/current/svrwx/400x300/isvrwxNE_.gif"><br/>
			<img class="pure-img-responsive" src="http://sirocco.accuweather.com/web_images/svrwx/key/swskeys.gif">
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
        	<div id="noaa_watches">
		        <img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/watch/validww.png"><br/>
		        <img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/watch/wwlegend.png">
		    </div>
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
        	<img class="pure-img-responsive" src="http://www.weather.gov/wwamap/png/ict.png">
        </div>
    </div>

    <?php require_once('includes/footer.php'); ?>
</div>

<script src="js/ui.js"></script>

</body>
</html>
