<?php

require_once('init.php');
require_once('functions.php');

/************************************
TO DO:
add something about @TornadoAlertApp
*************************************/

$thispage = 'http://' . $_SERVER['HTTP_HOST'] . $_SERVER['SCRIPT_NAME'];

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
if (isset($_REQUEST['refresh']) && $_REQUEST['refresh'] != 'no') {
	echo "<meta http-equiv=\"refresh\" content=\"300\" />";
}
?>
<meta name="keywords" content="Wichita Weather, Wichita, Weather, Radar, Satellite, Animation, Map, Watches, Warnings, Thunderstorms, Tornado, Storm warning, tornado warning, tornado watch" />
<meta name="description" content="Your one-stop shop for animated radar and satellite maps of Wichita area weather, including satellite, radar, watches and warnings, and power outages." />
<link rel="stylesheet" type="text/css" href="css/style.css" />

<link rel="apple-touch-icon-precomposed" sizes="57x57" href="apple-touch-icon-57x57.png" />
<link rel="apple-touch-icon-precomposed" sizes="114x114" href="apple-touch-icon-114x114.png" />
<link rel="apple-touch-icon-precomposed" sizes="72x72" href="apple-touch-icon-72x72.png" />
<link rel="apple-touch-icon-precomposed" sizes="144x144" href="apple-touch-icon-144x144.png" />
<link rel="apple-touch-icon-precomposed" sizes="60x60" href="apple-touch-icon-60x60.png" />
<link rel="apple-touch-icon-precomposed" sizes="120x120" href="apple-touch-icon-120x120.png" />
<link rel="apple-touch-icon-precomposed" sizes="76x76" href="apple-touch-icon-76x76.png" />
<link rel="apple-touch-icon-precomposed" sizes="152x152" href="apple-touch-icon-152x152.png" />
<link rel="icon" type="image/png" href="favicon-196x196.png" sizes="196x196" />
<link rel="icon" type="image/png" href="favicon-96x96.png" sizes="96x96" />
<link rel="icon" type="image/png" href="favicon-32x32.png" sizes="32x32" />
<link rel="icon" type="image/png" href="favicon-16x16.png" sizes="16x16" />
<link rel="icon" type="image/png" href="favicon-128.png" sizes="128x128" />
<meta name="application-name" content="Josh's Weather Station"/>
<meta name="msapplication-TileColor" content="#313131" />
<meta name="msapplication-TileImage" content="mstile-144x144.png" />
<meta name="msapplication-square70x70logo" content="mstile-70x70.png" />
<meta name="msapplication-square150x150logo" content="mstile-150x150.png" />
<meta name="msapplication-wide310x150logo" content="mstile-310x150.png" />
<meta name="msapplication-square310x310logo" content="mstile-310x310.png" />

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
			<?php if (isset($_REQUEST['refresh']) && $_REQUEST['refresh'] == 'no') { ?>
				<a href="<?=$thispage?>">start refreshing again</a>
			<?php } else { ?>
				this page automatically refreshes every five minutes. <a href="<?=$thispage?>?refresh=no">make it stop</a>
			<?php } ?>
		</td>
		<td align="right" style="font-family: tahoma, verdana, arial, sans-serif; color: #CCCCCC; font-size: 10px;"><a href="/about.php">about this page</a></td>
	</tr>
</table>

<?php
	$maps = array(
		'warnings' => false,
		'accuweather_radar' => array(
			'width' => '450',
			'height' => '450'
		),
		'kake_radar' => false
	);

	$maps2 = array(
		'twc_interactive' => false,
		'twc' => array(
			'width' => '592',
			'height' => '405',
			'note' => 'This map doesn\'t work in Internet Explorer (you should be using <a href="http://www.google.com/chrome" target="_blank">chrome</a> anyhow)'
		),
		'ksn_radar' => false,
		'noaabaseloop' => false,
		'twc_doppler' => array(
			'width' => '600',
			'height' => '405',
			'note' => 'This map somehow guesses your location, so sorry if it\'s wrong (visiting <a href="http://www.weather.com/" target="_blank">weather.com</a> might fix it)'
		)
	);
?>

<div id="Radar">
	<?php getMenu("Radar"); ?>
	<table width="1000" border="0" cellspacing="1" cellpadding="1">
		<tr>
			<td align="center" valign="top" nowrap>
				<?php
					foreach ($maps as $map => $iframe) {
						echo '<div id="' . $map . '" class="map">';
						if (is_array($iframe)) {
							if (isset($iframe['note']) && $iframe['note'] != '') {
								echo '<span style="font-size: 10px;">' . $iframe['note'] . '</span>';
							}
							echo '<iframe name="' . $map . '" id="' . $map . '" width="'. $iframe['width'] . '" height="' . $iframe['height'] . '" src="maps/' . $map . '.php" frameborder="0" marginheight="0" marginwidth="0" scrolling="no"></iframe>';
						} else {
							require('maps/' . $map . '.php');
						}
						echo '</div>';
					}
				?>
			</td>
			<td align="left" valign="top">
				<?php
					foreach ($maps2 as $map => $iframe) {
						echo '<div id="' . $map . '" class="map">';
						if (is_array($iframe)) {
							if (isset($iframe['note']) && $iframe['note'] != '') {
								echo '<span style="font-size: 10px;">' . $iframe['note'] . '</span>';
							}
							echo '<iframe name="' . $map . '" id="' . $map . '" width="'. $iframe['width'] . '" height="' . $iframe['height'] . '" src="maps/' . $map . '.php" frameborder="0" marginheight="0" marginwidth="0" scrolling="no"></iframe>';
						} else {
							require('maps/' . $map . '.php');
						}
						echo '</div>';
					}
				?>
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
	<?php getMenu("Satellite"); ?>
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
	<?php getMenu("Current Watches"); ?>
	These are convective watches, so this will be relevant for things like thunderstorms and tornadoes. Not so much for winter storms.<br/>
	<iframe width="900" height="600" scrolling="auto" src="http://www.spc.noaa.gov/products/watch/"></iframe>
</div>
<div id="Reflectivity" style="display: none;">
	<?php getMenu("Reflectivity"); ?>
	<iframe width="900" height="600" scrolling="auto" src="http://radar.weather.gov/radar.php?product=N0R&rid=ICT&loop=yes"></iframe>
</div>
<div id="Outages" style="display: none;">
	<?php getMenu("Outages"); ?>
	<iframe name="outage" id="outage" width="1000" height="600" src="http://outagemap.westarenergy.com/external/default.html" frameborder=0 marginheight=0 marginwidth=0 scrolling="auto" style="border; 1px solid black; background-color: #FFF; resize: both;"></iframe>
</div>

<div id="Resources" style="display: none;">
	<?php getMenu("Resources"); ?>

	<!--h2>Forums</h2>
	<ul>
		<li><a href="http://www.stormtrack.org/forum/forumdisplay.php?f=7" target="_blank">stormtrack.org</a> (look for a thread titled with today's date and the word "FCST" for forecasts, "NOW" for current reports)</li>
	</ul-->

	<h2>Streaming Video from Storm Chasers</h2>
	<ul>
		<li><a href="http://www.chasertv.com/" target="_blank">ChaserTV</a> "Live Weather Video On Demand"</li>
		<li><a href="http://www.severestudios.com/livechase" target="_blank">Severestudios.com</a> "The Leader in Live Severe Weather Streaming"</li>
		<li><a href="https://tvnweather.com/live" target="_blank">TVNWeather/Tornadovideos.net</a> "Live Storm Chasing"</li>
	</ul>

	<h2>Live Audio/Radio</h2>
	<ul>
		<!--li><a href="http://audiostream.wunderground.com/njenslin/wichita.mp3.m3u">NOAA weather radio for Wichita, from wunderground.com</a> live stream opens in your external audio player like windows media player or winamp</li-->
		<li><a href="http://www.radioreference.com/apps/audio/?ctid=970" target="_blank">RadioReference</a> Excellent collection of scanner feeds from Sedgwick County & Wichita area LEO agencies, emergency services, etc</li>
		<li><a href="javascript:void(window.open('http://player.streamtheworld.com/liveplayer.php?callsign=KFDIFM',%20'KFDI',%20'width=737,height=625,status=no,resizable=no,scrollbars=yes'))">KFDI radio</a> during storms, hands down the absolute best local radio storm coverage, with mobile spotters. Country music the rest of the time. Opens in a popup window.</li>
		<li><a href="http://scanwichita.com/listen.php" target="_blank">ScanWichita.com</a> Another great site with local area LEO and Emergency Services feeds, including Mid-Continent Approach Control</li>
	</ul>

	<h2>Other Stuff</h2>
	<ul>
		<li><a href="http://wichway.org/wichway/CameraTours" target="_blank">KanDrive Camera Tours</a> Multiple webcam views of traffic and road conditions along a Wichita area route of your chice, on a single web page. Choose from US-54, I-235, I-35, K-96, and I-135</li>
	</ul>
</div>

<div id="Twitter" style="display: none;">
	<?php getMenu("Twitter"); ?>
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