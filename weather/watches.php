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
<?php /*
      			<table width='150' border='1' cellspacing='0' cellpadding='0'>

       				<tr>
        				<td valign='top' width='2'></td>
        				<td valign='top' width='125'><img src='http://forecast.weather.gov/wwamap/gif/spacer.gif' width='125' height='3' border='0'></td>
        				<td valign='top' width='20'></td>
        				<td valign='top' width='2'></td>
       				</tr>
       				<tr>
        				<td valign='top' align='left' width='2'></td>
        				<td valign='top' align='left' width='125'><a href="http://forecast.weather.gov/wwamap/wwatxtget.php?cwa=ict&wwa=tornado warning">Tornado Warning</a></td>
        				<td valign='top' align='right' width='20'><img src='http://forecast.weather.gov/wwamap/gif/FF0000.gif' width='20' height='12' border='1'></td><td valign='top' align='left' width='2'></td>
       				</tr>

       				<tr>
        				<td valign='top' width='2'></td>
        				<td valign='top' width='125'><img src='http://forecast.weather.gov/wwamap/gif/spacer.gif' width='125' height='3' border='0'></td>
        				<td valign='top' width='20'></td>
        				<td valign='top' width='2'></td>
       				</tr>
       				<tr>
        				<td valign='top' align='left' width='2'></td>
        				<td valign='top' align='left' width='125'><a href="http://forecast.weather.gov/wwamap/wwatxtget.php?cwa=ict&wwa=flash flood warning">Flash Flood Warning</a></td>
        				<td valign='top' align='right' width='20'><img src='http://forecast.weather.gov/wwamap/gif/8B0000.gif' width='20' height='12' border='1'></td><td valign='top' align='left' width='2'></td>
       				</tr>

       				<tr>
        				<td valign='top' width='2'></td>
        				<td valign='top' width='125'><img src='http://forecast.weather.gov/wwamap/gif/spacer.gif' width='125' height='3' border='0'></td>
        				<td valign='top' width='20'></td>
        				<td valign='top' width='2'></td>
       				</tr>
       				<tr>
        				<td valign='top' align='left' width='2'></td>
        				<td valign='top' align='left' width='125'><a href="http://forecast.weather.gov/wwamap/wwatxtget.php?cwa=ict&wwa=severe weather statement">Severe Weather Statement</a></td>
        				<td valign='top' align='right' width='20'><img src='http://forecast.weather.gov/wwamap/gif/00FFFF.gif' width='20' height='12' border='1'></td><td valign='top' align='left' width='2'></td>
       				</tr>

       				<tr>
        				<td valign='top' width='2'></td>
        				<td valign='top' width='125'><img src='http://forecast.weather.gov/wwamap/gif/spacer.gif' width='125' height='3' border='0'></td>
        				<td valign='top' width='20'></td>
        				<td valign='top' width='2'></td>
       				</tr>
       				<tr>
        				<td valign='top' align='left' width='2'></td>
        				<td valign='top' align='left' width='125'><a href="http://forecast.weather.gov/wwamap/wwatxtget.php?cwa=ict&wwa=flood warning">Flood Warning</a></td>
        				<td valign='top' align='right' width='20'><img src='http://forecast.weather.gov/wwamap/gif/00FF00.gif' width='20' height='12' border='1'></td><td valign='top' align='left' width='2'></td>
       				</tr>

       				<tr>
        				<td valign='top' width='2'></td>
        				<td valign='top' width='125'><img src='http://forecast.weather.gov/wwamap/gif/spacer.gif' width='125' height='3' border='0'></td>
        				<td valign='top' width='20'></td>
        				<td valign='top' width='2'></td>
       				</tr>
       				<tr>
        				<td valign='top' align='left' width='2'></td>
        				<td valign='top' align='left' width='125'><a href="http://forecast.weather.gov/wwamap/wwatxtget.php?cwa=ict&wwa=tornado watch">Tornado Watch</a></td>
        				<td valign='top' align='right' width='20'><img src='http://forecast.weather.gov/wwamap/gif/FFFF00.gif' width='20' height='12' border='1'></td><td valign='top' align='left' width='2'></td>
       				</tr>

       				<tr>
        				<td valign='top' width='2'></td>
        				<td valign='top' width='125'><img src='http://forecast.weather.gov/wwamap/gif/spacer.gif' width='125' height='3' border='0'></td>
        				<td valign='top' width='20'></td>
        				<td valign='top' width='2'></td>
       				</tr>
       				<tr>
        				<td valign='top' align='left' width='2'></td>
        				<td valign='top' align='left' width='125'><a href="http://forecast.weather.gov/wwamap/wwatxtget.php?cwa=ict&wwa=flash flood watch">Flash Flood Watch</a></td>
        				<td valign='top' align='right' width='20'><img src='http://forecast.weather.gov/wwamap/gif/2E8B57.gif' width='20' height='12' border='1'></td><td valign='top' align='left' width='2'></td>
       				</tr>

       				<tr>
        				<td valign='top' width='2'></td>
        				<td valign='top' width='125'><img src='http://forecast.weather.gov/wwamap/gif/spacer.gif' width='125' height='3' border='0'></td>
        				<td valign='top' width='20'></td>
        				<td valign='top' width='2'></td>
       				</tr>
       				<tr>
        				<td valign='top' align='left' width='2'></td>
        				<td valign='top' align='left' width='125'><a href="http://forecast.weather.gov/wwamap/wwatxtget.php?cwa=ict&wwa=flood advisory">Flood Advisory</a></td>
        				<td valign='top' align='right' width='20'><img src='http://forecast.weather.gov/wwamap/gif/00FF7F.gif' width='20' height='12' border='1'></td><td valign='top' align='left' width='2'></td>
       				</tr>

       				<tr>
        				<td valign='top' width='2'></td>
        				<td valign='top' width='125'><img src='http://forecast.weather.gov/wwamap/gif/spacer.gif' width='125' height='3' border='0'></td>
        				<td valign='top' width='20'></td>
        				<td valign='top' width='2'></td>
       				</tr>
       				<tr>
        				<td valign='top' align='left' width='2'></td>
        				<td valign='top' align='left' width='125'><a href="http://forecast.weather.gov/wwamap/wwatxtget.php?cwa=ict&wwa=hazardous weather outlook">Hazardous Weather Outlook</a></td>
        				<td valign='top' align='right' width='20'><img src='http://forecast.weather.gov/wwamap/gif/EEE8AA.gif' width='20' height='12' border='1'></td><td valign='top' align='left' width='2'></td>
       				</tr>

       				<tr>
        				<td valign='top' width='2'></td>
        				<td valign='top' width='125'><img src='http://forecast.weather.gov/wwamap/gif/spacer.gif' width='125' height='3' border='0'></td>
        				<td valign='top' width='20'></td>
        				<td valign='top' width='2'></td>
       				</tr>
       				<tr>
        				<td valign='top' align='left' width='2'></td>
        				<td valign='top' align='left' width='125'><a href="http://forecast.weather.gov/wwamap/wwatxtget.php?cwa=ict&wwa=short term forecast">Short Term Forecast</a></td>
        				<td valign='top' align='right' width='20'><img src='http://forecast.weather.gov/wwamap/gif/98FB98.gif' width='20' height='12' border='1'></td><td valign='top' align='left' width='2'></td>
       				</tr>

     			<tr>
      				<td colspan='4'><img src='http://forecast.weather.gov/wwamap/gif/spacer.gif' width='150' height='3' border='0'></td>
     			</tr>
    			</table>
    			*/ ?>
        </div>
    </div>

    <?php require_once('includes/footer.php'); ?>
</div>

<script src="js/ui.js"></script>

</body>
</html>
