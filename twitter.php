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
		        <div class="pure-u pure-u-sm-1 pure-u-lg-1-2 pure-u-xl-1-2">
					<h2 class="content-subhead">Wichita Weather</h2>
					<!-- my KSWX list-->

					<a class="twitter-timeline"
						data-height="1200"
						data-width="600"
						data-theme="dark"
						data-link-color="#87C2ED"
						data-theme="dark"
						href="https://twitter.com/joshdutcher/lists/kswx?ref_src=twsrc%5Etfw"
						data-chrome="nofooter">
						Wichita weather information
					</a>
					<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>
		        </div>

		        <div class="pure-u pure-u-sm-1 pure-u-lg-1-2 pure-u-xl-1-2">
					<h2 class="content-subhead">US Weather</h2>
					<!-- my uswx list -->
					<a class="twitter-timeline"
						data-height="1200"
						data-width="600"
						data-link-color="#87C2ED"
						data-theme="dark"
						href="https://twitter.com/joshdutcher/lists/uswx?ref_Src=twsrc%5Etfw"
						data-chrome="nofooter">
						Nationwide weather information
					</a>
					<script async src="https://platform.twitter.com/widgets.js" charset="utf-8"></script>
		        </div>
		    </div>
		</div>
</div>

<?php require_once('includes/footer.php'); ?>

</body>
</html>
