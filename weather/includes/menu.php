<?php

$menu = array(
	'Radar' => array('url' => 'index.php', 'tooltip' => 'What you really came here for'),
	'Base Reflectivity' => array('url' => 'reflectivity.php', 'tooltip' => ''),
	'Satellite' => array('url' => 'satellite.php', 'tooltip' => ''),
	'Watches/Warnings' => array('url' => 'watches.php', 'tooltip' => ''),
	'Outlook' => array('url' => 'outlook.php', 'tooltip' => ''),
	'Resources' => array('url' => 'resources.php', 'tooltip' => ''),
	'Twitter' => array('url' => 'twitter.php', 'tooltip' => ''),
	'About' => array('url' => 'about.php', 'tooltip' => '')
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
        		echo '<li class="pure-menu-item"><a href="' . $file['url'] . '" class="pure-menu-link" title="' . $file['tooltip'] . '">' . $item . '</a></li>';
        	}
        ?>
        </ul>
    </div>
</div>