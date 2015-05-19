<?php

$menu = array(
	'Radar' => 'index.php',
	'Satellite' => 'satellite.php',
	'Current Watches' => 'watches.php',
	'Reflectivity' => 'reflectivity.php',
	'Outages' => 'outages.php',
	'Resources' => 'resources.php',
	'Twitter' => 'twitter.php',
	'About' => 'about.php'
);

?>

<!-- Menu toggle -->
<a href="#menu" id="menuLink" class="menu-link">
    <!-- Hamburger icon -->
    <span></span>
</a>

<div id="menu">
    <div class="pure-menu">
        <a class="pure-menu-heading" href="index.php">Home</a>

        <ul class="pure-menu-list">
        <?php
        	foreach ($menu as $item => $file) {
        		echo '<li class="pure-menu-item"><a href="' . $file . '" class="pure-menu-link">' . $item . '</a></li>';
        	}
        ?>
        </ul>
    </div>
</div>