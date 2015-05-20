<?php require_once('includes/init.php'); ?>
<?php require_once('includes/header.php'); ?>
<body>

<div id="layout">
    <?php require_once('includes/menu.php'); ?>

    <div class="pure-g" id="mainbody">
        <div class="pure-u-1 pure-u-md-1-2 pure-u-lg-1-3">
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
    </div>

    <?php require_once('includes/footer.php'); ?>
</div>

<script src="js/ui.js"></script>

</body>
</html>
