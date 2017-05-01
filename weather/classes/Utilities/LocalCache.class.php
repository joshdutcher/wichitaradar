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
        return time() - filemtime($filepath) > $this->cacheAge;
    }

}
