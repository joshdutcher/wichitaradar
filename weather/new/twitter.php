<?php require_once('includes/init.php'); ?>
<?php require_once('includes/header.php'); ?>

<body>

<div id="layout">
    <?php require_once('includes/menu.php'); ?>

	<script src="http://widgets.twimg.com/j/2/widget.js"></script>

	<p>This is a live stream of weather related tweets. Inspired by <a href="http://twitter.com/#!/myz06vette" target="_blank">@myz06vette</a>'s page, <a href="http://www.ksstorms.com/">ksstorms.com</a>.</p>
    <div class="pure-g" id="mainbody">
        <div class="pure-u pure-u-sm-1 pure-u-lg-1-2 pure-u-xl-1-3">
			<span class="twtr_header">Wichita Weather</span><br/>
			<!-- my ICTWX list-->
			<a class="twitter-timeline"
				href="https://twitter.com/joshdutcher/ictwx"
				data-widget-id="337783900774486016"
				data-chrome="nofooter">
				Wichita Weather
			</a>
			<script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?'http':'https';if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src=p+"://platform.twitter.com/widgets.js";fjs.parentNode.insertBefore(js,fjs);}}(document,"script","twitter-wjs");</script>
        </div>

        <div class="pure-u pure-u-sm-1 pure-u-lg-1-2 pure-u-xl-1-3">
			<span class="twtr_header">#ksstorms/#kswx/#ictwx hashtags</span><br/>
			<!-- #ksstorms / #kswx / #ictwx -->
			<a class="twitter-timeline"
				href="https://twitter.com/search?q=%23kswx+OR+%23ictwx+OR+%23ksstorms"
				data-widget-id="337785602101637120"
				data-chrome="nofooter">
				#ksstorms/#kswx/#ictwx hashtags
			</a>
			<script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?'http':'https';if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src=p+"://platform.twitter.com/widgets.js";fjs.parentNode.insertBefore(js,fjs);}}(document,"script","twitter-wjs");</script>
        </div>

        <div class="pure-u pure-u-sm-1 pure-u-lg-1-2 pure-u-xl-1-3">
			<span class="twtr_header">US Weather</span><br/>
			<!-- my uswx list -->
			<a class="twitter-timeline"
				href="https://twitter.com/joshdutcher/uswx"
				data-widget-id="337784188226912256"
				data-chrome="nofooter">
				US Weather
			</a>
			<script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?'http':'https';if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src=p+"://platform.twitter.com/widgets.js";fjs.parentNode.insertBefore(js,fjs);}}(document,"script","twitter-wjs");</script>
        </div>
    </div>

    <?php require_once('includes/footer.php'); ?>
</div>

<script src="js/ui.js"></script>

</body>
</html>
