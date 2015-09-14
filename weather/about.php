<?php require_once('includes/init.php'); ?>

<!doctype html>
<html lang="en">
<?php require_once('includes/header.php'); ?>
<body>

<div id="layout">
    <?php require_once('includes/menu.php'); ?>


    <div id="main">
        <div class="header">
            <h1>About this site</h1>
            <h2>wx.joshdutcher.com</h2>
        </div>

        <div class="content">
            <h2 class="content-subhead">Welcome to Version 2.0 of wx.joshdutcher.com!</h2>

		    <p>Updates include a visual overhaul and most importantly, responsive design. This means the site should look
		    100% better on mobile devices - smartphones, iPads, etc. I also added some new maps and removed some crappy ones.</p>

		    <h2 class="content-subhead">But what is this for?</h2>

		    <p>My name is Josh Dutcher.  I originally put this weather page together so I could keep track of the weather if my satellite dish went out.
		    I'm a web developer, so I just built it by harvesting source code from other sites.  I don't have permission
		    to use any of these maps, but whatever.  I never
		    anticipated it being used by anyone other than myself, so I just tossed it up on my server.  I certainly don't make any
		    money from it or anything like that - it's just a resource for people who find it useful.</p>

		    <p>The main home page with all the radars automatically refreshes every five minutes. It does this
		    because those radar animations get cached and need to be reloaded periodically in order to stay current. Auto refresh allows this
		    to happen automatically without you having to hit F5 all the time.</p>

		    <p>As a general rule, most sites including this one work better in Google Chrome or Firefox than they do in Internet Explorer.</p>

		    <p>Bookmark the site at <a href="http://wx.joshdutcher.com/">wx.joshdutcher.com</a></p>

		    <p>I can be contacted at josh.dutcher at joshdutcher dot com.</p>
        </div>
    </div>

    <?php require_once('includes/footer.php'); ?>
</div>


<script src="js/ui.js"></script>


</body>
</html>
