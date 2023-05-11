<?php

namespace Utilities;

use \Utilities\LocalCache;

class Scraper
{
    protected $remoteFilePath;
    protected $localFilePath;
    protected $localCache;

    public function __construct($remoteFilePath, $localFilePath, $cacheAge)
    {
        $this->remoteFilePath = $remoteFilePath;
        $this->localFilePath  = $localFilePath;
        $this->localCache     = new LocalCache($cacheAge);
    }

    public function getXMLFromFile()
    {
        if ($this->localCache->expired($this->localFilePath)) {
            $this->downloadFile();
        }

        return simplexml_load_file($this->localFilePath);
    }

    private function downloadFile() {
        $curl = curl_init($this->remoteFilePath);
        curl_setopt($curl, CURLOPT_RETURNTRANSFER, true);
        curl_setopt($curl, CURLOPT_USERAGENT, 'Mozilla/5.0 (Windows; U; Windows NT 5.1; en-US; rv:1.8.1.13) Gecko/20080311 Firefox/2.0.0.13');
        $data = curl_exec($curl);
        curl_close($curl);

        // Write data to file
        $file = fopen($this->localFilePath, 'wb');
        fwrite($file, $data);
        fclose($file);
    }
}
