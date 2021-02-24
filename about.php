<?php require_once 'includes/init.php';?>

<!doctype html>
<html lang="en">
<?php require_once 'includes/header.php';?>
<body>

<div id="layout">
    <?php require_once 'includes/menu.php';?>

    <div id="main">
        <div class="header">
            <h1>About this site</h1>
            <h2>wichitaradar.com</h2>
        </div>

        <div class="content">
            <h2 class="content-subhead">Welcome to wichitaradar.com!</h2>

		    <p>My name is Josh Dutcher.  I originally put this weather page together so I could keep track of the weather if my satellite dish went out.
		    I'm a web developer, so I just built it by harvesting source code from other sites.  I don't have permission
		    to use any of these maps, but whatever.  I never
		    anticipated it being used by anyone other than myself, so I just tossed it up on my server.  I certainly don't make any
		    money from it or anything like that - it's just a resource for people who find it useful.</p>

		    <p>The main home page with all the radars and the satellite page automatically refresh every five minutes. They do this
		    because those animations only load once, when you load the page, and after a while they become outdated. Auto-refreshing the page
            reloads those animations automatically without you having to hit F5 all the time.</p>

		    <p>As a general rule, most sites including this one work better in Google Chrome or Firefox than they do in Internet Explorer or Edge.</p>

		    <p>Bookmark the site at <a href="http://wichitaradar.com/">wichitaradar.com</a></p>

		    <p>I can be contacted at josh.dutcher at joshdutcher dot com.</p>
        </div>
    </div>

    <?php require_once 'includes/footer.php';?>
</div>


<script src="js/ui.js"></script>


</body>
</html>
