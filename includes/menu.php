<?php

$menu = array(
    'Radar'             => array('url' => 'index.php', 'tooltip' => 'What you really came here for'),
    'Satellite'         => array('url' => 'satellite.php', 'tooltip' => ''),
    'Watches/Warnings'  => array('url' => 'watches.php', 'tooltip' => ''),
    'Current Temps'     => array('url' => 'temperatures.php', 'tooltip' => ''),
    'Outlook'           => array('url' => 'outlook.php', 'tooltip' => ''),
    'Rainfall Amounts'  => array('url' => 'rainfall.php', 'tooltip' => ''),
    // 'Earthquakes'       => array('url' => 'earthquakes.php', 'tooltip' => ''),
    'Resources'         => array('url' => 'resources.php', 'tooltip' => ''),
    'Twitter'           => array('url' => 'twitter.php', 'tooltip' => ''),
    'About'             => array('url' => 'about.php', 'tooltip' => ''),
    'Donate'            => array('url' => 'donate.php', 'tooltip' => 'Buy me a beer!', 'new' => true)
);

$loadtime = new DateTime('now');
$loadtime->setTimezone(new DateTimeZone('America/Chicago'));

?>

<!-- Menu toggle -->
<a href="#menu" id="menuLink" class="menu-link">
    <!-- Hamburger icon -->
    <span></span>
</a>

<div id="menu">
    <div class="pure-menu">
        <ul class="pure-menu-list">
        <?php
foreach ($menu as $item => $file) {
    $class = 'pure-menu-item';
    if (basename($_SERVER['PHP_SELF']) == $file['url']) {
        $class .= ' selected';
    }
    echo '<li class="' . $class . '">';
    if (isset($file['new'])) {
        echo '<div class="new-menu-link"><img id="new" src="/img/new.png"></div>';
    }
    echo '<a href="' . $file['url'] . '" class="pure-menu-link" title="' . $file['tooltip'] . '">' . $item . '</a>';
    echo '</li>';
}
?>

        </ul>
    </div>
    <div class="menu-footer">
        <?php echo $loadtime->format('H:i:s'); ?><br/>
        <?php echo $loadtime->format('Y-m-d'); ?>
    </div>
</div>
