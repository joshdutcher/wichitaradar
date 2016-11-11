<?php
/*
<iframe src="https://embed.windytv.com/?37.805,-97.339,6,temp,message,marker,metric.wind.mph,metric.temp.F" width="750" height="500" frameborder="0"></iframe>
 */
;?>

<?php require_once 'includes/init.php';?>
<?php require_once 'includes/header.php';?>
<body>

<div id="layout">
    <?php require_once 'includes/menu.php';?>

    <div class="pure-g" id="mainbody">
        <div class="pure-u pure-u-1 pure-u-md-1-1 pure-u-lg-1-2">
            <a href="http://www.weather.gov/crh/weatherstory?sid=ict#.WCX0gvkrJhE">
                <img class="pure-img-responsive" src="http://www.weather.gov/images/ict/WxStory/FileL.png" border="0" />
            </a>
			<img class="pure-img-responsive" src="http://www.weather.gov/images/ict/GraphiCast/FileL.png" />
    	</div>
<?php
// <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
// <img class="pure-img-responsive" src="http://gray.ftp.clickability.com/kakewebftp/wx-forecast-7day-SC.jpeg" />
// <img class="pure-img-responsive" src="http://wx.ksn.com/weather/images/wx_weekly_640.jpg" />
// <img class="pure-img-responsive" src="http://image.weather.com/images/maps/forecast/map_wkpln_day1_3uscn_enus_720x486.jpg" />
// </div>
;?>
        <div class="pure-u pure-u-1 pure-u-md-1-1 pure-u-lg-1-2">
        	<div class="pure-u textbox">
        		Convective outlook for the next 3 days
        	</div>
        	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day1otlk_1300.gif" />
        	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day2otlk_0600.gif" />
        	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day3otlk_0730.gif" />
    	</div>
    </div>

    <?php require_once 'includes/footer.php';?>
</div>

<script src="js/ui.js"></script>

</body>
</html>
