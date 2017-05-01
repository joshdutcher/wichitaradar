<?php

set_include_path(dirname(dirname(__FILE__)) . DIRECTORY_SEPARATOR . "classes");
spl_autoload_extensions('.class.php');
spl_autoload_register();

$displayErrors = strpos($_SERVER['HTTP_HOST'], 'wx.joshdutcher.com') !== true;
ini_set('display_errors', $displayErrors);

if (!defined('ROOT')) {
    define('ROOT', dirname(dirname(__FILE__)));
}

$imgpath = (strpos($_SERVER['HTTP_HOST'], 'wx.joshdutcher.com') !== false) ? '/img/' : '/weather/img';
