<?php

namespace Utilities;

use \Utilities\LocalCache;

class GetImage
{
    protected $localPath;
    protected $dataURL;
    protected $localCache;
    protected $referer;

    public function __construct($dataURL, $cacheAge, $referer)
    {
        $this->localPath  = dirname(dirname(dirname(__FILE__))) . '/scraped/images/';
        $this->dataURL    = $dataURL;
        $this->localCache = new LocalCache($cacheAge);
        $this->referer    = $referer;
    }

    public function getImage($filename)
    {
        $filepath = $this->localPath . $filename;

        if ($this->localCache->expired($filepath)) {
            // create curl resource
            $ch = curl_init();

            // set url
            curl_setopt($ch, CURLOPT_URL, $this->dataURL);

            // set headers and options
            curl_setopt($ch, CURLOPT_HEADER, 0);
            curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
            curl_setopt($ch, CURLOPT_BINARYTRANSFER,1);
            curl_setopt($ch, CURLOPT_USERAGENT, 'Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.13) Gecko/20080311 Firefox/2.0.0.13');

            // set referer
            curl_setopt($ch, CURLOPT_REFERER, $this->referer);

            // $output contains the output string
            $content = curl_exec($ch);

            // close curl resource to free up system resources
            curl_close($ch);

            $this->writeFile($content, $filepath);
        }
    }

    protected function writeFile($content, $filepath)
    {
        if(file_exists($filepath)){
            unlink($filepath);
        }
        $file = fopen($filepath, 'w');
        fwrite($file, $content);
        fclose($file);
    }
}
