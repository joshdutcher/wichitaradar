<?php
// oh, turns out I don't need this, they make an animated gif available. neat.
namespace Utilities;

class GetGoesImages
{
    public static function getImages($directoryURL, $imageDimension, $numImages) {
        $html = file_get_contents($directoryURL);
        $goesUrlArray = explode("\n", $html);

        // filter the array of html lines to only include the image dimension we want

        $goesUrlArray = array_filter($goesUrlArray, function ($line) use ($imageDimension){
            return strpos($line, $imageDimension);
        });

        // parse each line to only return the url of the image
        foreach ($goesUrlArray as &$line) {
            $line = str_replace("</a>", "", substr($line, strpos($line, ">")+1)); // remove the <a href=""> and </a> tags from around the filename
            $line = substr($line, 0, strpos($line, ".jpg")+4); // remove all the trailing text so only the filename remains
            $line = $directoryURL.$line;
        }

        // get rid of the base/latest image (i.e. "625x375.jpg");
        if (($key = array_search($directoryURL . $imageDimension . ".jpg", $goesUrlArray)) !== false) {
            unset($goesUrlArray[$key]);
            $goesUrlArray = array_values($goesUrlArray);
        }

        // sort array by filename (this sorts it by timestamp). it's probably already sorted this way but this makes sure
        usort($goesUrlArray, function($a, $b) {
            return strcmp($a, $b);
        });

        // use only the most recent 36 images
        $goesUrlArray = array_slice($goesUrlArray, -36);

        return $goesUrlArray;
    }
}
