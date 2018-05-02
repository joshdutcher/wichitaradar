<?php

set_include_path(dirname(dirname(__FILE__)) . DIRECTORY_SEPARATOR . "classes");
spl_autoload_extensions('.class.php');
spl_autoload_register();

$displayErrors = strpos($_SERVER['HTTP_HOST'], 'wichitaradar.com') !== true;
ini_set('display_errors', $displayErrors);

if (!defined('ROOT')) {
    define('ROOT', dirname(dirname(__FILE__)));
}

function debug($content, $comments=false) {
    $backtrace = debug_backtrace();
    $calling_file = $backtrace[0]['file'];
    $calling_line = $backtrace[0]['line'];

    $vartype = gettype($content);

    if ($comments) {
        // prints the content inside an html comment, requiring user to view source to read it
        $prefix = "<!--" . "\r\n" . $calling_file . " line " . $calling_line . "\r\n";
        $suffix = "\r\n" . "-->" . "\r\n";
    } else {
        $prefix = '<div style="width: 100%; text-align:left;" align="left">' .
            '<pre style="align: left; font-size: 0.8em;  font-family: \'Courier New\', Courier, monospace; background-color: #111; padding: 5px; color: #CCF; border: 1px solid #999;">' .
            '<span style="font-weight: bold; color: #FFA">' . $calling_file . '</span> line <span style="font-weight: bold; color: #FFA">' . $calling_line . '</span>:<br/><br/>';
        $suffix = '</pre></div>';
    }

    echo $prefix;

    switch($vartype) {
        case 'array':
            print_r($content);
            break;
        case 'object':
            var_dump($content);
            break;
        default:
            echo $content;
    }

    echo $suffix;
}
