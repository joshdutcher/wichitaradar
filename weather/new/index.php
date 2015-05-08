<?php

require_once('includes/init.php');
require_once('includes/functions.php');

/************************************
TO DO:
add something about @TornadoAlertApp

http://forecast.weather.gov/MapClick.php?lat=37.6874&lon=-97.3427&unit=0&lg=english&FcstType=graphical

purecss.io

animation: http://awardwinningfjords.com/2012/03/08/image-sequences.html

bing maps like on KSNW?

hourly forecast
*************************************/

?>

<!doctype html>
<html lang="en">
<?php require_once('includes/header.php'); ?>
<body>

<style>
#layout {
    /*background-color: black;*/
}
</style>

<div id="layout">
    <?php require_once('includes/menu.php'); ?>

    <?php /*

    <div style="float: left;">
        <?php require('maps/warnings.php'); ?>
    </div>

    <div style="float: left;">
        <?php require('maps/kake_radar.php'); ?>
    </div>

    <div style="float: left;">
        <?php require('maps/ksn_radar.php'); ?>
    </div>

    <div style="float: left;">
        <?php require('maps/noaabaseloop.php'); ?>
    </div>

    <div style="float: left;">
        <?php require('maps/twc.php'); ?>
    </div>

    <br style="clear: both" />

    */ ?>


    <div class="pure-g">
        <div class="pure-u-1 pure-u-md-1-2 pure-u-lg-1-3">
            <div id="myImage"></div>
            <?php
                require('maps/ksn_radar.php');
                require('maps/kake_radar.php');
                require('maps/noaabaseloop.php');


                /*
                <a href="http://www.dillonmcintosh.tumblr.com/">
                    <img src="http://24.media.tumblr.com/d6b9403c704c3e5aa1725c106e8a9430/tumblr_mvyxd9PUpZ1st5lhmo1_1280.jpg"
                         alt="Beach">
                </a>

                <aside class="photo-box-caption">
                    <span>Watches &and; Warnings</span>
                </aside>
                */
            ?>
        </div>

        <div class="pure-u-1 pure-u-md-1-2 pure-u-lg-1-3">
            <?php
                require('maps/intellicast_wwa.php');

                /*
                <h1 class="text-box-head">Photos from around the world</h1>
                <p class="text-box-subhead">A collection of beautiful photos gathered from Unsplash.com.</p>
                */
            ?>
        </div>

        <div class="pure-u-1 pure-u-md-1-2 pure-u-lg-1-3">
            <?php
                require('maps/warnings.php');
                require('maps/weather_gov_wwa.php');


                /*
                <a href="http://ngkhanhlinh.dunked.com/">
                    <img src="http://31.media.tumblr.com/aa1779a718c2844969f23c4f5dec86b1/tumblr_mvyxhonf601st5lhmo1_1280.jpg"
                         alt="Meadow">
                </a>

                <aside class="photo-box-caption">
                    <span>
                        by <a href="http://ngkhanhlinh.dunked.com/">Linh Nguyen</a>
                    </span>
                </aside>
                */
            ?>
        </div>

        <?php /*

        <div class="photo-box u-1 u-med-1-2 u-lrg-1-3">
            <a href="http://www.nilssonlee.se/">
                <img src="http://24.media.tumblr.com/23e3f4bb271b8bdc415275fb7061f204/tumblr_mve3rvxwaP1st5lhmo1_1280.jpg"
                     alt="City">
            </a>

            <aside class="photo-box-caption">
                <span>
                    by <a href="http://www.nilssonlee.se/">Jonas Nilsson Lee</a>
                </span>
            </aside>
        </div>

        <div class="photo-box u-1 u-med-1-2 u-lrg-1-3">
            <a href="http://www.flickr.com/photos/rulasibai/">
                <img src="http://24.media.tumblr.com/ac840897b5f73fa6bc43f73996f02572/tumblr_mrraat0H431st5lhmo1_1280.jpg"
                     alt="Flowers">
            </a>

            <aside class="photo-box-caption">
                <span>
                    by <a href="http://www.flickr.com/photos/rulasibai/">Rula Sibai</a>
                </span>
            </aside>
        </div>

        <div class="photo-box u-1 u-med-1-2 u-lrg-1-3">
            <a href="http://www.flickr.com/photos/charliefoster/">
                <img src="http://24.media.tumblr.com/e100564a3e73c9456acddb9f62f96c79/tumblr_mufs8mix841st5lhmo1_1280.jpg"
                     alt="Bridge">
            </a>

            <aside class="photo-box-caption">
                <span>
                    by <a href="http://www.flickr.com/photos/charliefoster/">Charlie Foster</a>
                </span>
            </aside>
        </div>

        <div class="photo-box photo-box-thin u-1 u-lrg-2-3">
            <a href="http://ngkhanhlinh.dunked.com/">
                <img src="http://24.media.tumblr.com/c35afcc83e18ea7875160f64c039f471/tumblr_mwhdohfePJ1st5lhmo1_1280.jpg"
                     alt="Balloons">
            </a>

            <aside class="photo-box-caption">
                <span>
                    by <a href="http://ngkhanhlinh.dunked.com/">Linh Nguyen</a>
                </span>
            </aside>
        </div>

        <div class="photo-box photo-box-thin u-1 u-med-2-3">
            <a href="http://twitter.com/iBoZR">
                <img src="http://25.media.tumblr.com/95c842c76d60b7bc982d92c76216d037/tumblr_mx3tnm96k81st5lhmo1_1280.jpg"
                     alt="Rain Drops">
            </a>

            <aside class="photo-box-caption">
                <span>
                    by <a href="http://twitter.com/iBoZR">Thanun Buranapong</a>
                </span>
            </aside>
        </div>

        <div class="photo-box u-1 u-med-1-3">
            <a href="http://www.goodfreephotos.com/">
                <img src="http://25.media.tumblr.com/88b812f5f9c3d7b83560fd635435d538/tumblr_mx3tlblmY21st5lhmo1_1280.jpg"
                     alt="Port">
            </a>

            <aside class="photo-box-caption">
                <span>
                    by <a href="http://www.goodfreephotos.com/">Yinan Chen</a>
                </span>
            </aside>
        </div>

        <div class="u-1">
            <div class="l-box">
                <?php // <h2>Creating a Photo Gallery Layout</h2> ?>

                <p>
                    There is also a <a href="http://wx.joshdutcher.com/kc">northeast KS & KCMO area</a> version of this site.
                </p>
            </div>
        </div>
        */
        ?>
    </div>

    <?php require_once('includes/footer.php'); ?>
</div>


<script src="js/ui.js"></script>
<script type="text/javascript">
	//$(goAnimate);
</script>


</body>
</html>
