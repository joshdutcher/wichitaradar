<?php require_once('includes/init.php'); ?>
<?php require_once('includes/header.php'); ?>
<?php
    /*
    TODO:
    http://www.weather.gov/images/nws/satellite_images/6.jpg   6-1
    http://mp1.met.psu.edu/~fxg1/SAT_SC/satir_14.gif   1-24
    */
?>
<body>

<div id="layout">
    <?php require_once('includes/menu.php'); ?>

    <div class="pure-g">
        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
        	<div class="accuweather">
				<img class="pure-img-responsive" src="http://sirocco.accuweather.com/sat_mosaic_640x480_public/ei/isaeks_.gif" /><br/>
				<img class="pure-img-responsive" src="http://vortex.accuweather.com/adc2004/common/images/keys/400x40/sat_ei.gif" />
			</div>
            <img class="pure-img-responsive" src="http://weather.unisys.com/satellite/sat_ir_enh_cp_loop-12.gif" />
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
            <img class="pure-img-responsive" src="https://icons.wxug.com/data/640x480/2xsp_vi_anim.gif" />
        	<img class="pure-img-responsive" src="http://images.intellicast.com/WxImages/Satellite/ceusa.jpg" />
            <img class="pure-img-responsive" src="http://images.intellicast.com/WxImages/SatelliteIR/sln.jpg" />
        </div>

        <div class="pure-u pure-u-md-1 pure-u-lg-1-2 pure-u-xl-1-3">
            <img class="pure-img-responsive" src="http://images.intellicast.com/WxImages/SatelliteLoop/hiusa_None_anim.gif"  />
            <img class="pure-img-responsive" src="http://www.ssd.noaa.gov/PS/PCPN/DATA/RT/NA/IR4/20.jpg" />
        </div>
    </div>

    <?php require_once('includes/footer.php'); ?>
</div>

<script src="js/ui.js"></script>

</body>
</html>
