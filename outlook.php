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

use Utilities\GetFileContents;

function getGraphicasts() {
    $dataURL  = 'https://www.weather.gov/source/ict/wxstory/wxstory.xml';
    $cacheAge = '900'; // in seconds. 3600 = 1 hour, 1800 = 30 minutes, etc
    $filename = 'wxstory.xml';

    $getXML     = new GetFileContents($dataURL, $cacheAge);
    $xmlContent = $getXML->curlFile($filename);

    $xml         = new SimpleXMLElement($xmlContent);
    $graphicasts = $xml->xpath('//*/graphicasts/graphicast');

    $wxstoryImgArray = [];
    $wxstoryCount = 0;
    foreach ($graphicasts as $graphicast) {
        $timeNow   = time();
        $startTime = $graphicast->StartTime->__toString();
        $endTime   = $graphicast->EndTime->__toString();
        $radar     = (boolean) $graphicast->radar->__toString();
        $imageUrl  = $graphicast->SmallImage->__toString();
        if ($timeNow < $endTime && $timeNow >= $startTime && !$radar) {
            $wxstoryImgArray[$wxstoryCount]['url'] = 'http://www.weather.gov' . $imageUrl . '?' . rand(100000,999999);
            $wxstoryImgArray[$wxstoryCount]['alt'] = $graphicast->description->__toString();
            $wxstoryImgArray[$wxstoryCount]['order'] = (int) $graphicast->order->__toString();
        }
        $wxstoryCount++;
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

// this assumes there is only one <img> tag in each description. If there
// are more than one, it only returns the first one.
function parseDescForImg($description) {
    $doc = new DOMDocument();
    $doc->loadHTML($description);
    $xpath = new DOMXPath($doc);
    $img = $xpath->evaluate("string(//img/@src)");
    return $img;
}

function getConvectiveOutlooks() {
    $dataURL = 'https://www.spc.noaa.gov/products/spcacrss.xml';
    $cacheAge = '900';
    $filename = 'spcacrss.xml';

    $getXML = new getFileContents($dataURL, $cacheAge);
    $xmlContent = $getXML->curlFile($filename);
    $xml = new SimpleXMLElement($xmlContent);

    // $fakexml = dirname(__FILE__) . "/fakeoutlookrss.xml";
    // $xml = simplexml_load_file($fakexml);

    $convectiveOutlooks = $xml->xpath("//*/item");

    // sort them by pubDate, most recent to least recent
    function sortNodeByDate($a, $b) {
        return (strtotime($a->pubDate) < strtotime($b->pubDate)) ? -1 : 1;
    }
    usort($convectiveOutlooks, "sortNodeByDate");

    // if there is more than one outlook for a given day, this will
    // only keep the most recent one for a given day
    // resulting in one convective outlook per day
    $outlooks = [];
    foreach ($convectiveOutlooks as $node) {
        $key = substr($node->title, strpos($node->title, "UTC"));
        $outlooks[$key]['link'] = $node->link->__toString();
        $outlooks[$key]['title'] = $node->title->__toString();
        $outlooks[$key]['img'] = parseDescForImg($node->description->__toString());
        $outlooks[$key]['pubDate'] = $node->pubDate->__toString();
    }

    return $outlooks;
}

$wxstoryImgArray = getGraphicasts();
// $convectiveOutlooks = getConvectiveOutlooks();
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
<?php
foreach ($convectiveOutlooks as $outlook) {
    echo "<a href=\"{$outlook['link']}\">";
    echo "<img class=\"pure-img-responsive\" src=\"{$outlook['img']}\" border=\"0\" alt=\"{$outlook['title']}\" />";
    echo '</a>';
}
            // <a href="http://www.spc.noaa.gov/products/outlook/">
            // 	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day1otlk_1300.gif" />
            // 	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day2otlk_0600.gif" />
            // 	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day3otlk_0730.gif" />
            // </a>
?>
            </div>
    </div>
</div>

<?php require_once 'includes/footer.php';?>

</body>
</html>
