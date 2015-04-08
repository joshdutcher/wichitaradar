<?php

require_once('includes/init.php');
require_once('includes/functions.php');

/************************************
TO DO:
add something about @TornadoAlertApp
*************************************/

?>

<!doctype html>
<html lang="en">
<?php require_once('includes/header.php'); ?>
<body>

<div id="layout">
    <?php require_once('includes/menu.php'); ?>

    <p>My name is Josh Dutcher.  I originally put this weather page together so I could keep track of the weather if my satellite dish went out.
    I'm a web developer, so I just built it by harvesting source code from other sites.  I don't have permission
    to use any of these maps, but whatever.  I never
    anticipated it being used by anyone other than myself, so I just tossed it up on my server.  That's why it's not especially pretty -
    it's just meant to be functional. I certainly don't make any money off it or anything like that - it's just a resource for
    people who find it useful.</p>

    <p>The site automatically refreshes every five minutes, which resets the tab you're viewing to the default radar tab. It does this
    because those radar animations get cached and need to be reloaded periodically in order to stay current. Auto refresh allows this
    to happen automatically without you having to hit F5 all the time. If you'd like it not to do that, like if you want to watch
    the twitter tab instead, just click the link at the top of the page that says "make it stop."</p>

    <p>As a general rule, most sites including this one work better in Firefox or Google Chrome than they do in Internet Explorer.</p>

    <p>Bookmark the site at <?=$url?></p>

    <p>I can be contacted at busychild424 &#97;&#116; &#103;&#109;&#97;&#105;&#108; dot com.</p>

    <?php require_once('includes/footer.php'); ?>
</div>


<script src="js/ui.js"></script>


</body>
</html>
