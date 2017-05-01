<?php
/*
<iframe src="https://embed.windytv.com/?37.805,-97.339,6,temp,message,marker,metric.wind.mph,metric.temp.F" width="750" height="500" frameborder="0"></iframe>
 */

// die();

// $dataUrl = 'http://www.weather.gov/source/ict/wxstory/wxstory.xml';
// $xml     = Utilities::getExternalXML($dataUrl);
// die();

// $xml         = new SimpleXMLElement($content);
// $graphicasts = $xml->xpath('/*/graphicasts/graphicast');
// // var_dump($graphicasts);
// foreach ($graphicasts as $graphicast) {
//     // print_r($graphicast);
//     // var_dump($graphicast->title);
//     // var_dump($graphicast->description);
//     // var_dump($graphicast->ImageLoop);
//     // var_dump($graphicast->StartTime->__toString());
//     // $StartTime = date('Y/m/d H:i:s', $graphicast->StartTime->__toString());
//     $starttime = new \DateTime();
//     $starttime->setTimeStamp($graphicast->StartTime->__toString());
//     echo $starttime->format('Y-m-d H:i:s') . "<br>";
//     // var_dump($starttime);
//     // echo "StartTime: {$graphicast->StartTime->__toString()} ({$StartTime})<br/>";
// }
// // while (list(, $node) = each($result)) {
// //     var_dump($node);
// // }
// die();

// // $storyXML2 = simplexml_load_string($content);
// // print_r($storyXML2);

// // parsing begins here:
// libxml_use_internal_errors(true);
// $doc = new DOMDocument();
// if ($doc->loadXML($content)) {

// } else {
//     die('fail');
// }

// // use xpath
// $xpath = new DOMXPath($doc);
// // "//div[@id='StoryContent']//img[contains(@class, 'cq-dd-image')][1]/@src",

// //*[@id="image_tiles"]/div/div[1]/div/li[2]/a/img

// $graphicasts = $xpath->query("graphicasts");
// var_dump($graphicasts);
// foreach ($graphicasts as $graphicast) {
//     var_dump($graphicast);
//     echo $graphicast->StartTime . "<br/>";
//     echo $graphicast->EndTime . "<br/>";
// }
// die();

// if ($descriptions->length !== 0) {
//     foreach ($descriptions as $description) {
//         $metatags['description'][] = $this->cleanContent($description->value);
//     }
// }

// // get open graph meta tags
// $og_metas = $xpath->query('//*/meta[starts-with(@property, \'og:\')]');
// foreach ($og_metas as $og_meta) {
//     $property       = $og_meta->getAttribute('property');
//     $metatagContent = $this->cleanContent($og_meta->getAttribute('content'));
//     $tags_we_want   = [
//         'og:title',
//         'og:description',
//     ];

//     if (in_array($property, $tags_we_want)) {
//         $metatags[$property] = $metatagContent;
//     }
// }

// $ae               = new ArticleExtractor;
// $stripped_content = $ae->getContent($content);
// $stripped_content = $this->cleanContent($stripped_content);

// $return_data = [
//     'success'    => true,
//     'html_title' => $html_title,
//     'meta_tags'  => $metatags,
//     'content'    => $stripped_content,
// ];

// return $return_data;

require_once 'includes/init.php';

// sometimes http://www.weather.gov/ict/ has only one weather story image up, sometimes 3 or 4. the problem is
// that they leave old images up so we can't just grab whatever image files are there and toss them on the page.
// we therefore need to use the same xml document they are using.

use Utilities\GetXML;

$dataUrl  = 'http://www.weather.gov/source/ict/wxstory/wxstory.xml';
$cacheAge = '3600'; // 1 hour, in seconds
$getXML   = new GetXML($dataUrl, $cacheAge);

$filename = 'wxstory.xml';
$getXML->getAndWriteXML($filename);

require_once 'includes/header.php';
?>

<body>

<div id="layout">
    <?php require_once 'includes/menu.php';?>

    <div class="pure-g" id="mainbody">
         <div class="pure-u pure-u-1 pure-u-md-1-1 pure-u-lg-1-2">
            <a href="http://www.weather.gov/crh/weatherstory?sid=ict#.WCX0gvkrJhE">
                <img class="pure-img-responsive" src="http://www.weather.gov/images/ict/wxstory/Tab2FileL.png" border="0" id="wichitaWeatherStory" />
            </a>
    	</div>
<?php
// <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
// <img class="pure-img-responsive" src="http://gray.ftp.clickability.com/kakewebftp/wx-forecast-7day-SC.jpeg" />
// <img class="pure-img-responsive" src="http://wx.ksn.com/weather/images/wx_weekly_640.jpg" />
// <img class="pure-img-responsive" src="http://image.weather.com/images/maps/forecast/map_wkpln_day1_3uscn_enus_720x486.jpg" />
// </div>
;?>
         <div class="pure-u pure-u-1 pure-u-md-1-1 pure-u-lg-1-2">
        	<div class="pure-u textbox">
        		Convective outlook for the next 3 days
        	</div>
        	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day1otlk_1300.gif" />
        	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day2otlk_0600.gif" />
        	<img class="pure-img-responsive" src="http://www.spc.noaa.gov/products/outlook/day3otlk_0730.gif" />
    	</div>
    </div>

    <?php require_once 'includes/footer.php';?>
</div>

<script src="js/ui.js"></script>

<script>
    $(function() {
        // Wichita Weather Story
        var wichitaWeatherStory = {
            URLs: getWeatherStoryUrls()
        }
        wichitaWeatherStory.numImages = wichitaWeatherStory.URLs.length;
        $.when(
            preloadImages(wichitaWeatherStory)
        ).then(
            animateFrames(wichitaWeatherStory, 0, 5000, '#wichitaWeatherStory')
        );
    });

    function getWeatherStoryUrls() {
        var weatherStoryUrls = new Array();
        $.ajax({
            type: "GET",
            url: '/scraped/xml/wxstory.xml',
            cache: false,
            async: false,
            dataType: "xml",
            success: function(data) {
                $(data).find('graphicasts').each(function() {
                    $(this).children().each(function() {
                        var exp = $(this).find("EndTime").text()
                        var img = $(this).find("SmallImage").text()
                        var rad = $(this).find("radar").text()
                        milliseconds = (new Date).getTime();
                        if ((milliseconds/1000) < exp && rad == 0) {
                            weatherStoryUrls.push(img);
                        }
                    })
                });
            }
        });
        return weatherStoryUrls;
    }
</script>

</body>
</html>
