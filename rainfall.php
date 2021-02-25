<?php require_once('includes/init.php'); ?>
<?php require_once('includes/header.php'); ?>

<body>
<div id="layout">
    <?php require_once('includes/menu.php'); ?>

    <div class="pure-g" id="mainbody">
    	<div class="pure-u pure-u-1 pure-u-md-1-1 pure-u-lg-1-2">
        	<div class="pure-u textbox">
        		Past 24 hours
        	</div>

        	<div class="accuweather">
                <a href="https://www.accuweather.com/en/us/kansas/weather-radar-24hr">
                    <img class="pure-img-responsive" src="http://sirocco.accuweather.com/nx_mosaic_640x480c/24hr/inm24hrks_.gif" /><br/>
                    <img class="pure-img-responsive" src="http://vortex.accuweather.com/adc2010/images/keys/radar24hrf.png" />
                </a>
	        </div>
			<a href="https://www.wunderground.com/maps/precipitation/daily/sln">
        		<img class="pure-img-responsive" src="https://s.w-x.co/staticmaps/wu/pbs/preday/sln/20210223/1200z.gif" />
			</a>
			<a href="https://www.wunderground.com/maps/precipitation/daily">
        		<img class="pure-img-responsive" src="https://s.w-x.co/staticmaps/wu/pbs/preday/usa/20210223/1200z.gif" />
			</a>
    	</div>
    	<div class="pure-u pure-u-1 pure-u-md-1-1 pure-u-lg-1-2">
    	    <div class="pure-u textbox">
        		Past week
        	</div>
			<a href="https://www.wunderground.com/maps/precipitation/weekly/sln">
        		<img class="pure-img-responsive" src="https://s.w-x.co/staticmaps/wu/pbs/preweek/sln/20210223/1200z.gif" />
			</a>
			<a href="https://www.wunderground.com/maps/precipitation/weekly">
				<img class="pure-img-responsive" src="https://s.w-x.co/staticmaps/wu/pbs/preweek/usa/20210223/1200z.gif" />
			</a>
    	</div>
    </div>
</div>

<?php require_once('includes/footer.php'); ?>

</body>
</html>
