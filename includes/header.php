<!doctype html>
<html lang="en">

<head>
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
	<meta name="keywords" content="Wichita Weather, Wichita, Weather, Radar, Satellite, Animation, Map, Watches, Warnings, Thunderstorms, Tornado, Storm warning, tornado warning, tornado watch" />
	<meta name="description" content="Collection of animated radar and satellite maps of Wichita area weather, including radar, satellite, watches/warnings, and more info." />
    <title>Wichita Radar</title>
    <meta http-equiv="pragma" content="no-cache" />
    <meta http-equiv="cache-control" content="no-cache" />
<?php
// 300 seconds = 5 minutes
// 3600 seconds = 1 hour

$auto_refresh_pages = [
    'index'        => 300,
    'satellite'    => 300,
    'reflectivity' => 300,
    'outlook'      => 1800,
    'watches'      => 600,
    'twitter'      => 300,
];

$pagename = basename($_SERVER["PHP_SELF"], '.php');

if (array_key_exists($pagename, $auto_refresh_pages)) {
    echo "  <meta http-equiv=\"refresh\" content=\"{$auto_refresh_pages[$pagename]}\" />";
}
?>

	<link rel="apple-touch-icon-precomposed" sizes="57x57" href="apple-touch-icon-57x57.png" />
	<link rel="apple-touch-icon-precomposed" sizes="114x114" href="apple-touch-icon-114x114.png" />
	<link rel="apple-touch-icon-precomposed" sizes="72x72" href="apple-touch-icon-72x72.png" />
	<link rel="apple-touch-icon-precomposed" sizes="144x144" href="apple-touch-icon-144x144.png" />
	<link rel="apple-touch-icon-precomposed" sizes="60x60" href="apple-touch-icon-60x60.png" />
	<link rel="apple-touch-icon-precomposed" sizes="120x120" href="apple-touch-icon-120x120.png" />
	<link rel="apple-touch-icon-precomposed" sizes="76x76" href="apple-touch-icon-76x76.png" />
	<link rel="apple-touch-icon-precomposed" sizes="152x152" href="apple-touch-icon-152x152.png" />
	<link rel="icon" type="image/png" href="favicon-196x196.png" sizes="196x196" />
	<link rel="icon" type="image/png" href="favicon-96x96.png" sizes="96x96" />
	<link rel="icon" type="image/png" href="favicon-32x32.png" sizes="32x32" />
	<link rel="icon" type="image/png" href="favicon-16x16.png" sizes="16x16" />
	<link rel="icon" type="image/png" href="favicon-128.png" sizes="128x128" />
	<meta name="application-name" content="Wichita Radar"/>
	<meta name="msapplication-TileColor" content="#202020" />
	<meta name="msapplication-TileImage" content="mstile-144x144.png" />
	<meta name="msapplication-square70x70logo" content="mstile-70x70.png" />
	<meta name="msapplication-square150x150logo" content="mstile-150x150.png" />
	<meta name="msapplication-wide310x150logo" content="mstile-310x150.png" />
	<meta name="msapplication-square310x310logo" content="mstile-310x310.png" />

    <link rel="stylesheet" href="css/pure/0.6.0/pure-min.css">

    <!--[if lte IE 8]>
        <link rel="stylesheet" href="css/layouts/side-menu-old-ie.css">
        <link rel="stylesheet" href="css/pure/0.6.0/grids-responsive-old-ie-min.css">
    <![endif]-->
    <!--[if gt IE 8]><!-->
        <link rel="stylesheet" href="css/layouts/side-menu.css">
        <link rel="stylesheet" href="css/pure/0.6.0/grids-responsive-min.css">
    <!--<![endif]-->

    <link href='https://fonts.googleapis.com/css?family=Open+Sans:400italic,700italic,400,700' rel='stylesheet' type='text/css'>
    <link rel="stylesheet" href="css/wx.css">

    <script src="https://ajax.googleapis.com/ajax/libs/jquery/2.1.4/jquery.min.js"></script>
    <script type="text/javascript" src="js/functions.js"></script>
</head>
