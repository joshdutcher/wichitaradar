<?php

ini_set('display_errors',true);

if (!defined('ROOT')) {
	define('ROOT', dirname(dirname(__FILE__)));
}

$imgpath = (strpos($_SERVER['HTTP_HOST'], 'wx.joshdutcher.com') !== FALSE) ? '/img/' : '/weather/img';

?>