<?php

/*
maps to add

http://www.kctv5.com/wxmap/13384378/detail.html
http://www.nbcactionnews.com/subindex/weather/maps
http://www.kmbc.com/weather/16241998/iframe.html?qs=;longname=;shortname=Interactive;days=&ib_wxwidget=true
http://www.kmbc.com/weather/grid.html#HEARSTWX=http%3A//www.kmbc.com/weather/1677299/media.html%3Fqs%3D%3Bref%3D/weather/16809084/media.html%3Blongname%3DMap%2520Room
http://www.fox4kc.com/weather/

*/
$thispage = 'http://' . $_SERVER[HTTP_HOST] . $_SERVER[SCRIPT_NAME];

$tabs = array('Radar','Satellite','Watches','Reflectivity','Outages','Resources','Twitter');
$tabwidth = floor(1000/count($tabs));

function getMenu($currentTab) {
	global $tabs, $tabwidth;
	echo '<!-- currentTab: ' . $currentTab . ' -->';
	echo '<table border="0" cellspacing="2" cellpadding="3" width="1000" class="tabs">';
	echo '		<tr>';
	
	foreach ($tabs as $tab) {
		if ($currentTab == $tab) {
			echo "<td bgcolor=\"#666666\" width=\"$tabwidth\" align=\"center\">$tab</td>";
		}
		else {
			echo "<td bgcolor=\"#333333\" width=\"$tabwidth\" align=\"center\" onMouseOver=\"this.style.backgroundColor='#666666';this.style.cursor='pointer';\" onMouseOut=\"this.style.backgroundColor='#333333';this.style.cursor='default';\" onClick=\"switchTabs('$tab');\"><a href=\"javascript:switchTabs('$tab');\">$tab</a></td>";
		}
	}

	echo '		</tr>';
	echo '</table>';
}

?>


<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
<title>Josh's Weather Station - Kansas City, Topeka, Lawrence</title>
<meta http-equiv="Content-Type" content="text/html; charset=iso-8859-1" />
<meta http-equiv="pragma" content="no-cache" />
<?php
if ($_REQUEST['refresh'] != 'no') {
	echo "<meta http-equiv=\"refresh\" content=\"300\" />";
}
?>
<meta name="keywords" content="Kansas City, Topeka, Lawrence, Weather, Radar, Satellite, Animation, Map, Watches, Warnings, Thunderstorms, Tornado, Storm warning, tornado warning, tornado watch" />
<meta name="description" content="Your one-stop shop for animated radar and satellite maps of Kansas City, Topeka, and Lawrence area weather, including satellite, radar, watches and warnings, and power outages." />
<link rel="stylesheet" type="text/css" href="../style.css" />
</head>

<script language="javascript" type="text/javascript">
	<!--

	function switchTabs(tab) {
		var tabs = [];
		<?php
			// use the php array of tabs defined above so we only have to explicitly define it once, up there
			for ($i=0; $i<count($tabs); $i++) {
				echo "tabs[$i] = '$tabs[$i]';" . "\r\n";
			}
		?>
		
		for (i=0; i<tabs.length; i++) {
			if (tabs[i] == tab) {
				document.getElementById(tabs[i]).style.display = 'block';
			}
			else {
				document.getElementById(tabs[i]).style.display = 'none';
			}
		}
	}
	
	//-->
</script>

<body bgcolor="#000000">

<div class="orange">This is based on my original <a href="http://wx.joshdutcher.com/new/">Wichita-centered site</a>.</div>

<table width="1000" border="0" cellspacing="0" cellpadding="1">
	<!--tr>
		<td colspan="2" style="font-family: tahoma, verdana, arial, sans-serif; color: #CCCCCC; font-size: 10px;"><a href="http://m.joshdutcher.com/wxkc">view the mobile site</a></td>
	</tr-->
	<tr>
		<td style="font-family: tahoma, verdana, arial, sans-serif; color: #CCCCCC; font-size: 10px;">
			<?php if ($_REQUEST['refresh'] == 'no') { ?>
				<a href="<?=$thispage?>">start refreshing again</a>
			<?php } else { ?>
				this page automatically refreshes every five minutes. <a href="<?=$thispage?>?refresh=no">make it stop</a>
			<?php } ?>
		</td>
		<td align="right" style="font-family: tahoma, verdana, arial, sans-serif; color: #CCCCCC; font-size: 10px;"><a href="/about.php?f=kc">about this page</a></td>
	</tr>
</table>
<div id="Radar">
	<? getMenu("Radar"); ?>
	<table width="1000" border="0" cellspacing="1" cellpadding="1">
		<tr> 
			<td align="center" valign="top" nowrap>

				<!-- ******************************* -->
                <!-- ****** Watches & Warnings ***** -->
                <!-- ******************************* -->

				<div id="warnings" class="map">
					<img src="http://sirocco.accuweather.com/adc_images2/english/current/svrwx/400x300/isvrwxMO_.gif"><br/>
					<img src="http://sirocco.accuweather.com/web_images/svrwx/key/swskeys.gif">								
				</div>

				<!-- ******************************* -->
				<!-- ****** AccuWeather Radar ****** -->
				<!-- ******************************* -->
				
				<div id="accuweather_radar" class="map">
                    <img src="http://sirocco.accuweather.com/sat_mosaic_400x300_public/rs/isarmo_.gif" />
				</div>
			</td>
			<td align="left" valign="top">

                <!-- ******************************* -->
                <!-- ****** TWC interactive ******** -->
                <!-- ******************************* -->

				<div id="twc_interactive" class="map">
					<script src="http://www.weather.com/common/flash/wxgold/AC_OETags.js" language="javascript"></script>
					<script src="http://www.weather.com/common/flash8MapUtil.js" language="javascript"></script>
					<SCRIPT LANGUAGE="JavaScript1.2" SRC="http://www.weather.com/maps/interactive/config/beta/InteractiveMapConfig.js?2007011"></SCRIPT>
					<script language="JavaScript1.2">
						<!--
							if(isFlashAvailable()){
								if(isFlashVer8Available()){
									getFlashObjectTag("http://image.weather.com/web/flash/FMMain.swf",600,508, "flashid", "wxAnimateOnStart=true&viewPortWidth=600&viewPortHeight=405&lat=38.96&long=-95.27&initialWeatherLayerType=radar&trackingBaseURL=x.weather.com&initialZoomLevel=7&productID=9&panFrameAlpha=60&config="+freeSiteInteractiveMapXMLConfig,"#FFFFFF");
								}
								else{getFlashVerNotAvailableMessage();}
							}else {
								getFlashNotAvailableMessage();
							}
						// -->
					</script>
					<noscript>
						This content requires the Adobe Flash Player and a browser with JavaScript enabled.
						<a href="http://www.adobe.com/go/getflash/" target="_blank">Get Flash</a>
					</noscript>
				</div>

                <!-- ******************************************************************************************************** -->
                <!-- ****** Intellicast ************************************************************************************* -->
                <!-- ****** from http://www.intellicast.com/storm/severe/metro.aspx?location=USMO0460&animate=true ********** -->
                <!-- ******************************************************************************************************** -->

				<div id="intellicast" class="map">
					<div><img src="http://images.intellicast.com/images/legends/MetroStormWatch_700.gif" alt="Legend" width="600" /></div>
					<div style="position: relative;"><img src="http://images.intellicast.com/images/basemaps/Metro/mkc.jpg" width="600" />
						<div style="position: absolute; z-index: 1; top: 0; left: 0;"><img src="http://images.intellicast.com/WxImages/MetroStormWatchLoop/mkc_None_anim.gif" width="600"></div>
						<div style="position: absolute; z-index: 2; top: 0; left: 0;"><img src="http://images.intellicast.com/images/Overlays/Metro/counties/mkc.gif" width="600"></div>
						<div style="position: absolute; z-index: 3; top: 0; left: 0;"><img src="http://images.intellicast.com/images/Overlays/Metro/highways/mkc.gif" width="600"></div>
						<div style="position: absolute; z-index: 4; top: 0; left: 0;"><img src="http://images.intellicast.com/WeatherImages/MetroWatchBoxes/mkc.gif" width="600"></div>
						<div style="position: absolute; z-index: 5; top: 0; left: 0;"><img src="http://images.intellicast.com/images/Overlays/Metro/cities/mkc.gif" width="600"></div>
					</div>
				</div>

                <!-- ******************************* -->
                <!-- ****** TWC2 ******************* -->
                <!-- ******************************* -->
				
				<div id="twc" class="map">
					<span style="font-size: 10px;">This map doesn't work in Internet Explorer (you should be using <a href="http://www.google.com/chrome" target="_blank">chrome</a> anyhow):</span>
					<iframe name="twc" id="twc" width="592" height="405" src="maps/twc/map.html" frameborder="0" marginheight="0" marginwidth="0" scrolling="no"></iframe>
				</div>

                <!-- ******************************* -->
                <!-- * NOAA base reflectivity loop * -->
                <!-- ******************************* -->
				
				<div id="noaabaseloop" class="map">
					<img src="http://radar.weather.gov/lite/NCR/TWX_loop.gif">
				</div>

                <!-- ******************************* -->
                <!-- ****** TWC doppler ************ -->
                <!-- ******************************* -->

				<div id="twc_doppler" class="map">
					<span style="font-size: 10px; text-align: center;">This map uses magic to guess your location, so sorry if it's wrong (visiting <a href="http://www.weather.com/" target="_blank">weather.com</a> might fix it)</span><br/>
					<iframe name="mapI" id="mapI" width=600 height=405  src="maps/weather/map.html"	frameborder=0 marginheight=0 marginwidth=0 scrolling="no">
						<img src="http://image.weather.com/web/blank.gif" width=600 height=405 name="holdspace" border=0 alt="" />
						<script language="javascript1.2">
							<!--
								if (isMinNS4) var mapNURL = "/maps/local/local/us_close_ddc/1b/index_large_animated.html";
							// -->
						</script>
					</iframe>
				</div>
			</td>
		</tr>
	</table>
</div>
<div id="Satellite" style="display: none;">
	<!-- ******************************* -->
	<!-- ***** Accuweather satellite *** -->
	<!-- ******************************* -->
	<? getMenu("Satellite"); ?>
	<table border="0" cellspacing="0" cellpadding="5">	
		<tr>
			<td bgcolor="#E9E9E5" align="center">
				<img src="http://sirocco.accuweather.com/sat_mosaic_640x480_public/ei/isaemo_.gif" /><br/><br/>
				<img src="http://vortex.accuweather.com/adc2004/common/images/keys/400x40/sat_ei.gif" />
			</td>
		</tr>
	</table>
</div>
<div id="Watches" style="display: none;">
	<? getMenu("Watches"); ?>
	These are convective watches, so this will be relevant for things like thunderstorms and tornadoes. Not so much for winter storms.<br/>
	<iframe width="900" height="800" scrolling="auto" src="http://www.spc.noaa.gov/products/watch/"></iframe>
</div>
<div id="Reflectivity" style="display: none;">
	<? getMenu("Reflectivity"); ?>
	<iframe width="900" height="800" scrolling="auto" src="http://radar.weather.gov/radar.php?product=N0R&rid=TWX&loop=yes"></iframe>
</div>
<div id="Outages" style="display: none;">
	<? getMenu("Outages"); ?>
	<iframe name="outage" id="outage" width="800" height="600" src="maps/westar/map.html" frameborder=0 marginheight=0 marginwidth=0 scrolling="no"></iframe>
</div>

<div id="Resources" style="display: none;">
	<? getMenu("Resources"); ?>

	<!--h2>Forums</h2>
	<ul>
		<li><a href="http://www.stormtrack.org/forum/forumdisplay.php?f=7" target="_blank">stormtrack.org</a> (look for a thread titled with today's date and the word "FCST" for forecasts, "NOW" for current reports)</li>
	</ul-->
	
	<h2>Streaming Video from Storm Chasers</h2>
	<ul>
		<li><a href="http://www.chasertv.com/" target="_blank">ChaserTV</a> "Live Weather Video On Demand"</li>
		<li><a href="http://www.severestudios.com/livechase" target="_blank">Severestudios.com</a> "The Leader in Live Severe Weather Streaming"</li>
		<li><a href="http://www.tornadovideos.net/full-screen-chaser-video.php" target="_blank">Tornadovideos.net</a></li>
	</ul>
	
	<h2>Live Audio/Radio</h2>
	<ul>
		<li><a href="http://audioplayer.wunderground.com/OPwx/OverlandPark.mp3.m3u">NOAA weather radio for the Kansas City area, from wunderground.com</a> live stream opens in your external audio player like windows media player or winamp</li>
		<li><a href="http://www.radioreference.com/apps/audio/?coid=1" target="_blank">RadioReference</a> Click the state, then the county you're interested in, and you'll be presented with a list of scanner feeds from LEO agencies, emergency services, etc. Great site.</li>
		<!--li><a href="javascript:void(window.open('http://player.streamtheworld.com/jbroadcast/?CALLSIGN=KFDIFM',%20'KFDI',%20'width=737,height=625,status=no,resizable=no,scrollbars=yes'))">KFDI radio</a> during storms, the best local radio coverage, with mobile spotters. Otherwise, country music. Opens in a popup window.</li-->
	</ul>
</div>

<div id="Twitter" style="display: none;">
	<? getMenu("Twitter"); ?>
	<script src="http://widgets.twimg.com/j/2/widget.js"></script>
	<p>This is a live stream of weather related tweets. Inspired by <a href="http://twitter.com/#!/myz06vette" target="_blank">@myz06vette</a>'s page, <a href="http://www.ksstorms.com/">ksstorms.com</a>.</p>
	<p>Currently the "relevant hashtags" are #kswx, #ksstorms, #mowx, #mostorms, #kcstorms, and #kcwx. If you have KC area storm chasers I can add to another column, please let me know.</p>
	<p>Sometimes Twitter gets strange and these don't work right.</p>
	<table border="0" cellspacing="0" cellpadding="5">
		<tr>
			<td valign="top">
				<span class="twtr_header">KC Weather</span><br/>
				<!-- kcmowx -->
				<a class="twitter-timeline" href="https://twitter.com/joshdutcher/kcwx" data-widget-id="337784502908776448" width="323" height="600" data-chrome="nofooter">KC Weather</a>
				<script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?'http':'https';if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src=p+"://platform.twitter.com/widgets.js";fjs.parentNode.insertBefore(js,fjs);}}(document,"script","twitter-wjs");</script>
			</td>
			<td valign="top">
				<span class="twtr_header">relevant hashtags</span><br/>
				<!-- hashtags -->
				<a class="twitter-timeline" href="https://twitter.com/search?q=%23kswx+OR+%23ksstorms+OR+%23mowx+OR+%23mostorms+OR+%23kcstorms+OR+%23kcwx" data-widget-id="337785180095934464" width="323" height="600" data-chrome="nofooter">relevant hashtags</a>
				<script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?'http':'https';if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src=p+"://platform.twitter.com/widgets.js";fjs.parentNode.insertBefore(js,fjs);}}(document,"script","twitter-wjs");</script>
			</td>
			<td valign="top">
				<span class="twtr_header">US Weather</span><br/>
				<!-- my uswx list -->
				<a class="twitter-timeline" href="https://twitter.com/joshdutcher/uswx" data-widget-id="337784188226912256" width="323" height="600" data-chrome="nofooter">US Weather</a>
				<script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?'http':'https';if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src=p+"://platform.twitter.com/widgets.js";fjs.parentNode.insertBefore(js,fjs);}}(document,"script","twitter-wjs");</script>
			</td>
		</tr>
	</table>
</div>

<script type="text/javascript">
	var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
	document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
	try {
		var pageTracker = _gat._getTracker("UA-6815318-1");
		pageTracker._trackPageview();
	}
	catch(err) {}
</script>
</body>
</html>