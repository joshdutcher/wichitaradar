<?php require_once('includes/init.php'); ?>
<?php require_once('includes/header.php'); ?>

<body>

<div id="layout">
    <?php require_once('includes/menu.php'); ?>

	<script src="http://widgets.twimg.com/j/2/widget.js"></script>


	<div id="main">
        <div class="header">
            <h1>Twitter</h1>
            <h2>Twitter feeds with real time weather info</h2>
        </div>

        <div style="">
			<div class="pure-g">
		        <div class="pure-u pure-u-sm-1 pure-u-lg-1-2 pure-u-xl-1-3">
					<h2 class="content-subhead">Wichita Weather</h2>
					<!-- my ICTWX list-->
					<a class="twitter-timeline"
						data-height="1200"
						data-link-color="#87C2ED"
						data-theme="dark"
						href="https://twitter.com/joshdutcher/lists/ictwx1?ref_src=twsrc%5Etfw"
						data-chrome="nofooter">
						Wichita, KS area weather reports and storm chasers
					</a>
					<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>
		        </div>

		        <div class="pure-u pure-u-sm-1 pure-u-lg-1-2 pure-u-xl-1-3">
					<h2 class="content-subhead">#ksstorms/#kswx/#ictwx hashtags</h2>
					<!-- #ksstorms / #kswx / #ictwx -->
					<a class="twitter-timeline"
						data-height="1200"
						data-link-color="#87C2ED"
						data-theme="dark"
						href="https://twitter.com/search?q=%23kswx+OR+%23ictwx+OR+%23ksstorms"
						data-widget-id="337785602101637120"
						data-chrome="nofooter">
						#ksstorms/#kswx/#ictwx hashtags
					</a>
					<script>!function(d,s,id){var js,fjs=d.getElementsByTagName(s)[0],p=/^http:/.test(d.location)?'http':'https';if(!d.getElementById(id)){js=d.createElement(s);js.id=id;js.src=p+"://platform.twitter.com/widgets.js";fjs.parentNode.insertBefore(js,fjs);}}(document,"script","twitter-wjs");</script>
		        </div>

		        <div class="pure-u pure-u-sm-1 pure-u-lg-1-2 pure-u-xl-1-3">
					<h2 class="content-subhead">US Weather</h2>
					<!-- my uswx list -->
					<a class="twitter-timeline"
						data-height="1200"
						data-link-color="#87C2ED"
						data-theme="dark"
						href="https://twitter.com/joshdutcher/lists/uswx?ref_src=twsrc%5Etfw"
						data-chrome="nofooter">
						Nationwide weather information and storm chasers
					</a>
					<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>
		        </div>
		    </div>
		</div>


    <?php require_once('includes/footer.php'); ?>
</div>

<script src="js/ui.js"></script>

</body>
</html>
