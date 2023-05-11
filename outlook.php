<?php
/*
<iframe src="https://embed.windytv.com/?37.805,-97.339,6,temp,message,marker,metric.wind.mph,metric.temp.F" width="750" height="500" frameborder="0"></iframe>
 */

require_once 'includes/init.php';

// sometimes http://www.weather.gov/ict/ has only one weather story image up, sometimes 3 or 4. the problem is
// that they leave old images up so we can't just grab whatever image files are there and toss them on the page.
// we therefore need to use the same xml document they are using.

// similarly the convective outlook doesn't always have the same number of images available so we need to
// parse the same xml document they are using and display images based on that.

use Utilities\Scraper;

function getGraphicasts() {
    $remoteFilePath = 'https://www.weather.gov/source/ict/wxstory/wxstory.xml';
    $localFilePath  = dirname(__FILE__) . '/scraped/xml/wxstory.xml';
    $cacheAge       = '900'; // in seconds. 3600 = 1 hour, 1800 = 30 minutes, etc

    $scrapedFile = new Scraper($remoteFilePath, $localFilePath, $cacheAge);
    $xml         = $scrapedFile->getXMLFromFile();

    $wxstoryImgArray = [];

    $timeNow = time();
    foreach ($xml->graphicasts->graphicast as $graphicast) {
        $startTime = (string) $graphicast->StartTime;
        $endTime   = (string) $graphicast->EndTime;
        $radar     = (boolean) $graphicast->radar->__toString();
        if ($timeNow < $endTime && $timeNow >= $startTime && !$radar) {
            $imageUrl = (string) $graphicast->SmallImage;
            $wxstoryImgArray[] = array(
                'url'   => 'http://weather.gov' . $imageUrl . '?' . rand(100000,999999),
                'alt'   => preg_replace('/\s+/', ' ', trim((string) $graphicast->description)),
                'order' => (int) $graphicast->order
            );
        }
    }
    
    if (empty($wxstoryImgArray)) {
        $wxstoryImgArray[0]['url']   = '/img/nostories.png';
        $wxstoryImgArray[0]['alt']   = 'No Weather Stories!';
        $wxstoryImgArray[0]['order'] = 0;
    }

    // make sure the images display in the intended order
    function cmp($a, $b) {
        return $a['order'] - $b['order'];
    }
    usort($wxstoryImgArray,"cmp");

    return $wxstoryImgArray;
}

$wxstoryImgArray = getGraphicasts();
require_once 'includes/header.php';
?>

<body>

<div id="layout">
    <?php require_once 'includes/menu.php'; ?>

    <div class="pure-g" id="mainbody">
         <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
            <a href="https://www.ksn.com/weather/">
                <img class="pure-img-responsive" src="https://media.psg.nexstardigital.net/ksnw/weather/images/wx_weekly_full.jpg" />
                <img class="pure-img-responsive" src="https://media.psg.nexstardigital.net/ksnw/weather/images/QEHighsToday.jpg" />
                <img class="pure-img-responsive" src="https://media.psg.nexstardigital.net/ksnw/weather/images/QELowsTonight.jpg" />
                <img class="pure-img-responsive" src="https://media.psg.nexstardigital.net/ksnw/weather/images/QEHighsTomorrow.jpg" />
            </a>
         </div>

         <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
<?php
foreach ($wxstoryImgArray as $story) {
    echo '<a href="http://www.weather.gov/crh/weatherstory?sid=ict#.WCX0gvkrJhE">';
    echo "<img class=\"pure-img-responsive\" src=\"{$story['url']}\" border=\"0\" id=\"wichitaWeatherStory\" alt=\"{$story['alt']}\" />";
    echo '</a>';
}
?>
	    </div>
         <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
        	<div class="pure-u textbox">
        		Convective outlook
        	</div>
            <a href="http://www.spc.noaa.gov/products/outlook/">
            	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day1otlk.gif" />
            	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day2otlk.gif" />
            	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day3otlk.gif" />
            </a>
        </div>
    </div>
</div>

<?php require_once 'includes/footer.php';?>

</body>
</html>
