<?php

namespace Utilities;

// date_default_timezone_set('America/Chicago');

class SWXCOFiles
{
    protected $runtime;
    protected $date;
    protected $hour;
    protected $prefix = 'https://s.w-x.co/staticmaps/wu/fee4c/temp_cur';
    protected $images = ['usa','ddc'];
    protected $imagepaths = [];

    public function __construct()
    {
        $this->timezone = new \DateTimeZone('UTC');
        $this->runtime = new \DateTimeImmutable("now", $this->timezone);
        $this->runtime->setTimeZone($this->timezone);
        $this->maxHoursAgo = 12;
    }

    public function getImagePaths()
    {
        foreach ($this->images as $imagetype) {
            // If the image doesn't exist, it might not be built yet. Go back an hour and look for the previous one.
            // In fact, go back as far as $this->maxHoursAgo hours before giving up.
            $hoursago = 0;
            $imagefound = false;
            while ($hoursago < $this->maxHoursAgo) {
                $interval = 'PT' . $hoursago . 'H';
                $timeToCheck = $this->runtime->sub(new \DateInterval($interval));
                
                // https://s.w-x.co/staticmaps/wu/fee4c/temp_cur/ddc/20191002/1600z.jpg
                $imageUrlElements = [
                    $this->prefix,
                    $imagetype,
                    $timeToCheck->format('Ymd'),
                    $timeToCheck->format('H') . "00z.jpg"
                ];
                $imageUrl = implode("/", $imageUrlElements);

                if ($this->checkRemoteFile($imageUrl)) {
                    $this->imagepaths[$imagetype] = $imageUrl;
                    $imagefound = true;
                    break;
                }

                $hoursago++;
            }
            if (!$imagefound) {
                $this->imagepaths[$imagetype] = '';
            }
        }

        return $this->imagepaths;
    }

    protected function checkRemoteFile($url)
    {
        $ch = curl_init();
        curl_setopt($ch, CURLOPT_URL,$url);
        // don't download content
        curl_setopt($ch, CURLOPT_NOBODY, 1);
        curl_setopt($ch, CURLOPT_FAILONERROR, 1);
        curl_setopt($ch, CURLOPT_RETURNTRANSFER, 1);
    
        $result = curl_exec($ch);
        curl_close($ch);
        if($result !== FALSE)
        {
            return true;
        }
        else
        {
            return false;
        }
    }
}
