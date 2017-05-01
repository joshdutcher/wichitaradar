<?php

namespace Utilities;

use \Utilities\LocalCache;

class GetXML
{
    protected $localPath;
    protected $dataURL;
    protected $localCache;

    public function __construct($dataURL, $cacheAge)
    {
        $this->localPath  = dirname(dirname(dirname(__FILE__))) . '/scraped/xml/';
        $this->dataURL    = $dataURL;
        $this->localCache = new LocalCache($cacheAge);
    }

    public function getXML($filename)
    {
        $filepath = $this->localPath . $filename;

        if ($this->localCache->expired($filepath)) {
            // create curl resource
            $ch = curl_init();

            // set url
            curl_setopt($ch, CURLOPT_URL, $this->dataURL);

            //return the transfer as a string
            curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
            curl_setopt($ch, CURLOPT_USERAGENT, 'Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.13) Gecko/20080311 Firefox/2.0.0.13');

            // $output contains the output string
            $content = curl_exec($ch);

            // close curl resource to free up system resources
            curl_close($ch);

            $this->writeFile($content, $filename);
        } else {
            $content = file_get_contents($filepath);
        }

        return $content;
    }

    protected function writeFile($content, $filename)
    {
        $file = fopen($this->localPath . $filename, 'w');
        fwrite($file, $content);
        fclose($file);
    }
}
