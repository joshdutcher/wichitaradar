<?php
/*
<iframe src="https://embed.windytv.com/?37.805,-97.339,6,temp,message,marker,metric.wind.mph,metric.temp.F" width="750" height="500" frameborder="0"></iframe>
 */

require_once 'includes/init.php';

// sometimes http://www.weather.gov/ict/ has only one weather story image up, sometimes 3 or 4. the problem is
// that they leave old images up so we can't just grab whatever image files are there and toss them on the page.
// we therefore need to use the same xml document they are using.

use Utilities\GetXML;

$dataURL  = 'https://www.weather.gov/source/ict/wxstory/wxstory.xml';
$cacheAge = '900'; // in seconds. 3600 = 1 hour, 1800 = 30 minutes, etc
$filename = 'wxstory.xml';

$getXML     = new GetXML($dataURL, $cacheAge);
$xmlContent = $getXML->getXML($filename);

$xml         = new SimpleXMLElement($xmlContent);
$graphicasts = $xml->xpath('//*/graphicasts/graphicast');

$wxstoryImgArray = [];
foreach ($graphicasts as $graphicast) {
    $timeNow  = time();
    $endTime  = $graphicast->EndTime->__toString();
    $radar    = (boolean) $graphicast->radar->__toString();
    $imageUrl = $graphicast->SmallImage->__toString();
    if ($timeNow < $endTime && !$radar) {
        $wxstoryImgArray[] = 'http://www.weather.gov' . $imageUrl;
    }
}

if (empty($wxstoryImgArray)) {
    $wxstoryImgArray[] = '/img/nostories.png';
}

require_once 'includes/header.php';
?>

<body>

<div id="layout">
    <?php require_once 'includes/menu.php';?>

    <div class="pure-g" id="mainbody">
         <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
         <img class="pure-img-responsive" src="https://media.ksn.com/nxs-ksnwtv-media-us-east-1/weather/images/wx_weekly_full.jpg" />
            <img class="pure-img-responsive" src="https://media.ksn.com/nxs-ksnwtv-media-us-east-1/weather/images/QEHighsToday.jpg" />
            <img class="pure-img-responsive" src="https://media.ksn.com/nxs-ksnwtv-media-us-east-1/weather/images/QELowsTonight.jpg" />
            <img class="pure-img-responsive" src="https://media.ksn.com/nxs-ksnwtv-media-us-east-1/weather/images/QEHighsTomorrow.jpg" />
         </div>

         <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
<?php
foreach ($wxstoryImgArray as $story) {
    echo '<a href="http://www.weather.gov/crh/weatherstory?sid=ict#.WCX0gvkrJhE">';
    echo "<img class=\"pure-img-responsive\" src=\"{$story}\" border=\"0\" id=\"wichitaWeatherStory\" />";
    echo '</a>';
}
?>
	</div>

         <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
        	<div class="pure-u textbox">
        		Convective outlook for the next 3 days
        	</div>
            <a href="http://www.spc.noaa.gov/products/outlook/">
            	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day1otlk_1300.gif" />
            	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day2otlk_0600.gif" />
            	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day3otlk_0730.gif" />
            </a>
    	</div>
    </div>

    <?php require_once 'includes/footer.php';?>
</div>

<script src="js/ui.js"></script>

<script>
    // $(function() {
    //     // Wichita Weather Story
    //     var wichitaWeatherStory = {
    //         URLs: getWeatherStoryUrls()
    //     }
    //     wichitaWeatherStory.numImages = wichitaWeatherStory.URLs.length;
    //     $.when(
    //         preloadImages(wichitaWeatherStory)
    //     ).then(
    //         animateFrames(wichitaWeatherStory, 0, 5000, '#wichitaWeatherStory')
    //     );
    // });

    // function getWeatherStoryUrls() {
    //     var weatherStoryUrls = new Array();
    //     weatherStoryUrls.push("<?php echo implode($wxstoryImgArray, '","'); ?>");

    //     // $.ajax({
    //     //     type: "GET",
    //     //     url: '/scraped/xml/wxstory.xml',
    //     //     cache: false,
    //     //     async: false,
    //     //     dataType: "xml",
    //     //     success: function(data) {
    //     //         $(data).find('graphicasts').each(function() {
    //     //             $(this).children().each(function() {
    //     //                 var exp = $(this).find("EndTime").text()
    //     //                 var img = $(this).find("SmallImage").text()
    //     //                 var rad = $(this).find("radar").text()
    //     //                 milliseconds = (new Date).getTime();
    //     //                 if ((milliseconds/1000) < exp && rad == 0) {
    //     //                     weatherStoryUrls.push(img);
    //     //                 }
    //     //             })
    //     //         });
    //     //     }
    //     // });
    //     // console.log(weatherStoryUrls);
    //     return weatherStoryUrls;
    // }
</script>

</body>
</html>
