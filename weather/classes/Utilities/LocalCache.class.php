<?php

namespace Utilities;

class LocalCache
{
    protected $cacheAge;

    public function __construct($cacheAge)
    {
        $this->cacheAge = $cacheAge;
    }

    public function expired($filepath)
    {
        if (file_exists($filepath)) {
            return time() - filemtime($filepath) > $this->cacheAge;
        } else {
            return true; // file doesn't exist; we want to pull it
        }
    }
}
