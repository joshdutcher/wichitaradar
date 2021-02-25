<?php require_once('includes/init.php'); ?>
<?php require_once('includes/header.php'); ?>
<body>

<style>
	#noaa_watches {
		background-color: #FFF;
		text-align: center;
		display: table-cell;
	}

	#weathergov_legend {
		background-color: #4C484A;
		color: #FFF;
		text-align: left;
		font-size: 0.6em;
		margin: 0 auto;
		border-spacing: 2px;
		border-collapse: separate;
	}

	#weathergov_legend td {
		padding: 1px 4px;
	}

	.legend_cell {
		border: 2px solid black;
		width: 15px;
	}
</style>

<div id="layout">
    <?php require_once('includes/menu.php'); ?>

    <div class="pure-g" id="mainbody">
        <div class="pure-u pure-u-1 pure-u-md-1-1 pure-u-lg-1-2">
			<!-- ****** Accuweather ************ -->
			<a href="https://www.accuweather.com/en/us/severe-weather">
				<img class="pure-img-responsive" src="http://sirocco.accuweather.com/adc_images2/english/current/svrwx/400x300/isvrwxNE_.gif" /><br/>
				<img class="pure-img-responsive" src="http://sirocco.accuweather.com/web_images/svrwx/key/swskeys.gif" />
			</a>
			<br/>
			<a href="https://www.weather.gov/ict/">
				<img class="pure-img-responsive" src="http://www.weather.gov/wwamap/png/ict.png"><br/>
			</a>
			<table id="weathergov_legend">
				<tr>
					<td class="legend_cell" style="background-color: #FF0000;"></td>
					<td>Tornado Warning</td>
					<td class="legend_cell" style="background-color: #8B0000;"></td>
					<td>Flash Flood Warning</td>
				</tr>
				<tr>
					<td class="legend_cell" style="background-color: #FFFF00;"></td>
					<td>Tornado Watch</td>
					<td class="legend_cell" style="background-color: #2E8B57;"></td>
					<td>Flash Flood Watch</td>

				</tr>
				<tr>
					<td class="legend_cell" style="background-color: #FFA500;"></td>
					<td>Severe Thunderstorm Warning</td>
					<td class="legend_cell" style="background-color: #00FF00;"></td>
					<td>Flood Warning</td>
				</tr>
				<tr>
					<td class="legend_cell" style="background-color: #DB7093;"></td>
					<td>Severe Thunderstorm Watch</td>
					<td class="legend_cell" style="background-color: #00FF7F;"></td>
					<td>Flood Advisory</td>
				</tr>
				<tr>
					<td class="legend_cell" style="background-color: #00FFFF;"></td>
					<td>Severe Weather Statement</td>
					<td class="legend_cell" style="background-color: #D2B48C;"></td>
					<td>Wind Advisory</td>
				</tr>
				<tr>
					<td class="legend_cell" style="background-color: #EEE8AA;"></td>
					<td>Hazardous Weather Outlook</td>
					<td class="legend_cell" style="background-color: #98FB98;"></td>
					<td>Short Term Forecast</td>
				</tr>
				<tr>
					<td class="legend_cell" style="background-color: #FF69B4;"></td>
					<td>Winter Storm Warning</td>
					<td class="legend_cell" style="background-color: #7B68EE;"></td>
					<td>Winter Weather Advisory</td>
				</tr>
			</table>
        </div>

        <div class="pure-u pure-u-1 pure-u-md-1-1 pure-u-lg-1-2">
        	<div id="noaa_watches">
        		<a href="http://www.spc.noaa.gov/classic.html">
			        <img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/activity_loop.gif"><br/>
			    </a>
			    <a href="http://www.spc.noaa.gov/products/watch/">
			        <img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/watch/validww.png"><br/>
			        <img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/watch/wwlegend.png">
			    </a>
		    </div>
        </div>
    </div>
</div>

<?php require_once('includes/footer.php'); ?>

</body>
</html>
