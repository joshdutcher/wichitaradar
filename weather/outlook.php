<?php require_once('includes/init.php'); ?>
<?php require_once('includes/header.php'); ?>
<body>

<div id="layout">
    <?php require_once('includes/menu.php'); ?>

    <div class="pure-g" id="mainbody">
    	<div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
			<img class="pure-img-responsive" src="http://www.weather.gov/images/ict/WxStory/FileL.png" />

			<img class="pure-img-responsive" src="http://www.weather.gov/images/ict/GraphiCast/FileL.png" />
    	</div>
    	<div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
			<img class="pure-img-responsive" src="http://gray.ftp.clickability.com/kakewebftp/wx-forecast-7day-SC.jpeg" />
			<img class="pure-img-responsive" src="http://wx.ksn.com/weather/images/wx_weekly_640.jpg" />
			<div class="pure-u textbox">
				The following are experimental
			</div>
			<img class="pure-img-responsive" src="http://www.kwch.com/image/view/-/23064696/highRes/475/-/maxh/480/maxw/640/-/g8m8l2z/-/AM-FORECAST-IMAGE.jpg" />
			<img class="pure-img-responsive" src="http://image.weather.com/images/maps/forecast/map_wkpln_day1_3uscn_enus_720x486.jpg" />
    	</div>
    	<div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
        	<div class="pure-u textbox">
        		Convective outlook for the next 3 days
        	</div>
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
