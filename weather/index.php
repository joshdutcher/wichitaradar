<?php

/************************************
TO DO:
add something about @TornadoAlertApp
*************************************/

$thispage = 'http://' . $_SERVER[HTTP_HOST] . $_SERVER[SCRIPT_NAME];

// easily just define which tabs I want to have on the site
$tabs = array(
	'Radar',
	'Satellite',
	'Current Watches',
	'Reflectivity',
	'Outages',
	'Resources',
	'Twitter'
);

// automatically make the tabs all the same width
$tabwidth = floor(1000/count($tabs));

// dynamically generate the tab menu based on which tab we're on
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
<title>Josh's Weather Station</title>
<meta http-equiv="Content-Type" content="text/html; charset=iso-8859-1" />
<meta http-equiv="pragma" content="no-cache" />
<?php
if ($_REQUEST['refresh'] != 'no') {
	echo "<meta http-equiv=\"refresh\" content=\"300\" />";
}
?>
<meta name="keywords" content="Wichita Weather, Wichita, Weather, Radar, Satellite, Animation, Map, Watches, Warnings, Thunderstorms, Tornado, Storm warning, tornado warning, tornado watch" />
<meta name="description" content="Your one-stop shop for animated radar and satellite maps of Wichita area weather, including satellite, radar, watches and warnings, and power outages." />
<link rel="stylesheet" type="text/css" href="css/style.css" />
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

<div class="orange">There is also a <a href="http://wx.joshdutcher.com/kc">northeast KS & KCMO area</a> version of this site.</div>

<table width="1000" border="0" cellspacing="0" cellpadding="1">
	<tr>
		<td colspan="2" style="font-family: tahoma, verdana, arial, sans-serif; color: #CCCCCC; font-size: 10px;"><a href="http://m.joshdutcher.com">view the mobile site</a></td>
	</tr>
	<tr>
		<td style="font-family: tahoma, verdana, arial, sans-serif; color: #CCCCCC; font-size: 10px;">
			<?php if ($_REQUEST['refresh'] == 'no') { ?>
				<a href="<?=$thispage?>">start refreshing again</a>
			<?php } else { ?>
				this page automatically refreshes every five minutes. <a href="<?=$thispage?>?refresh=no">make it stop</a>
			<?php } ?>
		</td>
		<td align="right" style="font-family: tahoma, verdana, arial, sans-serif; color: #CCCCCC; font-size: 10px;"><a href="/about.php">about this page</a></td>
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
					<img src="http://sirocco.accuweather.com/adc_images2/english/current/svrwx/400x300/isvrwxNE_.gif"><br/>
					<img src="http://sirocco.accuweather.com/web_images/svrwx/key/swskeys.gif">								
				</div>
				
				<!-- ******************************* -->
				<!-- ****** AccuWeather Radar ****** -->
				<!-- ******************************* -->
				
				<div id="accuweather_radar" class="map">
					<iframe name="mapII" id="mapII" width="450" height="450" src="maps/accuweather/map.html" frameborder="0" marginheight="0" marginwidth="0" scrolling="no"></iframe>
				</div>
				
                <!-- ******************************* -->
                <!-- ****** KAKE radar ************* -->
                <!-- ******************************* -->
				
				<div id="kake_radar" class="map">
					<img src="http://gray.ftp.clickability.com/kakewebftp/wx-radar-kakeland.gif" width="450">
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
									getFlashObjectTag("http://image.weather.com/web/flash/FMMain.swf",600,508, "flashid", "wxAnimateOnStart=true&viewPortWidth=600&viewPortHeight=405&lat=37.67&long=-97.36&initialWeatherLayerType=radar&trackingBaseURL=x.weather.com&initialZoomLevel=7&productID=9&panFrameAlpha=60&config="+freeSiteInteractiveMapXMLConfig,"#FFFFFF");
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
				
                <!-- ******************************* -->
                <!-- ****** TWC2 ******************* -->
                <!-- ******************************* -->
				
				<div id="twc" class="map">
					<span style="font-size: 10px;">This map doesn't work in Internet Explorer (you should be using <a href="http://www.google.com/chrome" target="_blank">chrome</a> anyhow):</span>
					<iframe name="twc" id="twc" width="592" height="405" src="maps/twc/map.html" frameborder="0" marginheight="0" marginwidth="0" scrolling="no"></iframe>
				</div>

                <!-- ******************************* -->
                <!-- ****** KSN Radar ************** -->
                <!-- ******************************* -->

				<div id="ksn_radar" class="map">
					<!-- Begin pkg: iwradar rev3 -->
					<script language="JavaScript" type="text/javascript">

						<!-- Hide from JavaScript-Impaired Browsers
						if (document.images) {
							//Total frames represented by 18 followed by 17, frames 13-18 are repeats of 12
							isn=new Array();
							for (i=1;i<18;i++) {
								isn[i]=new Image();
								if (i<17){
									// Use 0 placeholder in filename when i is less than 10
									if (i<10){
										isn[i].src="http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_0"+i+".jpg";
									} else {
										if (i<13){
											isn[i].src="http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_"+i+".jpg";
										} else {
											isn[i].src="http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_12.jpg";
										}
									}
								} else {
									isn[i].src="/images/inv.gif";
								}
							}
						}
			
						//dla = default speed or delay between frames
						ctr=1;
						dla=150;
						j=1;
						stpit=0;
						restart=0;
					
						function startIt1(param){
							if (restart==1){
								stpit=0;
							}
							if (stpit<1){
								setTimeout("prtIt1()",dla);
								restart=0;
							}
						}
			
						function prtIt1(){
							if (document.images){
								document.ani1.src=isn[j].src;
								j++;
								if (j>16){
									j=1
								}
								if (stpit==1){
									stpit=0;
								}
								else{
									startIt1();
									}
								}
							else{
								alert("You need a JavaScript 1.1 compatible browser. We recommmend Netscape 3+ "
								+"or MSIE4+.");
							}
						}
			
						function speedIt1(){
							if (stpit==0){
								if (dla>250){
									dla-=250;
									if (dla<150){
										dla=100;
									}
								}
								else{
									dla-=50;
									if (dla<50){
										dla-=25;
										if (dla<25){
											dla=25;
										}
									}
								}
							}
						}
			
						function slowIt1(){
							if (stpit == 0){
								if (dla<50){
									dla+=25;
								}
								else{
									if (dla<150){
										dla+=50;
									}
									else{
										if (dla<=2150){
											dla+=250;
											if (dla>2150){
												dla=2500;
											}
										}
									}
								}
							}
						}
						
						function testStp1(){
							if (stpit==0){
								stpit=1;
							}
						}
			
						function Back1(){
							stpit=1;
							restart=1;
							if (document.images){
								j--;
								if (j<1){
									j=12
								}
								else{
									if (j>12){
										j=1
									}
								}
							document.ani1.src=isn[j].src;
							}
						}
			
						function Forward1(){
							stpit=1;
							restart=1;
							if (document.images){
								j++;
								if (j>12){
									j=1
								}
							document.ani1.src=isn[j].src;
							}
						}
					// End Hiding -->
					</script>
					<script type="text/javascript" language="JavaScript1.1">
						startIt1();
					</script>

					<!--- here is the actual map image -->
					<img name="ani1" border=0 src="http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_01.jpg" alt="Radar image" width="600" height="450" />
				</div>

                <!-- ******************************* -->
                <!-- * NOAA base reflectivity loop * -->
                <!-- ******************************* -->
				
				<div id="noaabaseloop" class="map">
					<img src="http://radar.weather.gov/lite/NCR/ICT_loop.gif">
				</div>

                <!-- ******************************* -->
                <!-- ****** TWC doppler ************ -->
                <!-- ******************************* -->

				<div id="twc_doppler" class="map">
					<span style="font-size: 10px;">This map somehow guesses your location, so sorry if it's wrong (visiting <a href="http://www.weather.com/" target="_blank">weather.com</a> might fix it)</span><br/>
					<iframe name="mapI" id="mapI" width=600 height=405 src="maps/weather/map.html"	frameborder=0 marginheight=0 marginwidth=0 scrolling="no"></iframe>
				</div>
			</td>
		</tr>
	</table>
</div>
<!--div id="forecast" style="display: none;">
	<img src="http://030b577.netsolhost.com/kwch/images/weather/WEB_7Day_ICT_final.JPG">
</div-->
<div id="Satellite" style="display: none;">
	<!-- ******************************* -->
	<!-- ***** Accuweather satellite *** -->
	<!-- ******************************* -->
	<? getMenu("Satellite"); ?>
	<table border="0" cellspacing="0" cellpadding="5">	
		<tr>
			<td bgcolor="#E9E9E5" align="center">
				<img src="http://sirocco.accuweather.com/sat_mosaic_640x480_public/ei/isaeks_.gif" /><br/><br/>
				<img src="http://vortex.accuweather.com/adc2004/common/images/keys/400x40/sat_ei.gif" />
			</td>
		</tr>
	</table>
</div>
<div id="Current Watches" style="display: none;">
	<? getMenu("Current Watches"); ?>
	These are convective watches, so this will be relevant for things like thunderstorms and tornadoes. Not so much for winter storms.<br/>
	<iframe width="900" height="600" scrolling="auto" src="http://www.spc.noaa.gov/products/watch/"></iframe>
</div>
<div id="Reflectivity" style="display: none;">
	<? getMenu("Reflectivity"); ?>
	<iframe width="900" height="600" scrolling="auto" src="http://radar.weather.gov/radar.php?product=N0R&rid=ICT&loop=yes"></iframe>
</div>
<div id="Outages" style="display: none;">
	<? getMenu("Outages"); ?>
	<iframe name="outage" id="outage" width="1000" height="600" src="http://outagemap.westarenergy.com/external/default.html" frameborder=0 marginheight=0 marginwidth=0 scrolling="auto" style="border; 1px solid black; background-color: #FFF; resize: both;"></iframe>
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
		<!--li><a href="http://audiostream.wunderground.com/njenslin/wichita.mp3.m3u">NOAA weather radio for Wichita, from wunderground.com</a> live stream opens in your external audio player like windows media player or winamp</li-->
		<li><a href="http://www.radioreference.com/apps/audio/?ctid=970" target="_blank">RadioReference</a> Excellent collection of scanner feeds from Sedgwick County & Wichita area LEO agencies, emergency services, etc</li>
		<li><a href="javascript:void(window.open('http://player.streamtheworld.com/jbroadcast/?CALLSIGN=KFDIFM',%20'KFDI',%20'width=737,height=625,status=no,resizable=no,scrollbars=yes'))">KFDI radio</a> during storms, hands down the absolute best local radio storm coverage, with mobile spotters. Country music the rest of the time. Opens in a popup window.</li>
		<li><a href="http://scanwichita.com/listen.php" target="_blank">ScanWichita.com</a> Another great site with local area LEO and Emergency Services feeds, including Mid-Continent Approach Control</li>
	</ul>
	
	<h2>Other Stuff</h2>
	<ul>
		<li><a href="http://wichway.org/WichWay/CameraTour/CameraTour.aspx" target="_blank">KanDrive Camera Tours</a> Multiple webcam views of traffic and road conditions along a Wichita area route of your chice, on a single web page. Choose from US-54, I-235, I-35, K-96, and I-135</li>
	</ul>
</div>

<div id="Twitter" style="display: none;">
	<? getMenu("Twitter"); ?>
	<script src="http://widgets.twimg.com/j/2/widget.js"></script>
	<p>This is a live stream of weather related tweets. Inspired by <a href="http://twitter.com/#!/myz06vette" target="_blank">@myz06vette</a>'s page, <a href="http://www.ksstorms.com/">ksstorms.com</a>.</p>
	<p>Sometimes Twitter gets strange and these don't work right.</p>
	<table border="0" cellspacing="0" cellpadding="5">
		<tr>
			<td valign="top">
				<span class="twtr_header">Wichita Weather</span><br/>
				<!-- my ICTWX list-->
				<a class="twitter-timeline" href="https://twitter.com/joshdutcher/ictwx" data-widget-id="337783900774486016" width="323" height="600" data-chrome="nofooter">Wichita Weather</a>
				<script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?'http':'https';if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src=p+"://platform.twitter.com/widgets.js";fjs.parentNode.insertBefore(js,fjs);}}(document,"script","twitter-wjs");</script>
			</td>
			<td valign="top">
				<span class="twtr_header">#ksstorms/#kswx/#ictwx hashtags</span><br/>
				<!-- #ksstorms / #kswx / #ictwx-->
				<a class="twitter-timeline" href="https://twitter.com/search?q=%23kswx+OR+%23ictwx+OR+%23ksstorms" data-widget-id="337785602101637120" width="323" height="600" data-chrome="nofooter">relevant hashtags</a>
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

<script>
	(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
	(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
	m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
	})(window,document,'script','//www.google-analytics.com/analytics.js','ga');

	ga('create', 'UA-6815318-1', 'joshdutcher.com');
	ga('send', 'pageview');
</script>
</body>
</html>