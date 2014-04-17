<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html xmlns="http://www.w3.org/1999/xhtml">
<head>
<title>Josh's Weather Station - Mobile</title>
<meta http-equiv="Content-Type" content="text/html; charset=iso-8859-1" />
<meta name="keywords" content="Wichita Weather, Wichita, Weather, Radar, Satellite, Animation, Map, Watches, Warnings,
Thunderstorms, Tornado, Storm warning, tornado warning, tornado watch" />
<meta name="description" content="Your one-stop shop for animated radar and satellite maps of Wichita area weather, including
satellite, radar, watches and warnings, and power outages." />

<style type="text/css">
	body{
		color: #FFFFFF;
		background-color: #000000;
		font-family: tahoma, verdana, arial, sans-serif;
		font-size: 10px;
	}
	
	div{
		padding: 5px 0;
	}
	
	a:link{
		color: #6666FF;
	}
	
	a:visited{
		color: #FF6666;
	}
	
	a:active{
		color: #6666FF;
	}
	
	ul{
		margin-left: .5em;
		padding-left: .5em;
	}
</style>

</head>

<body>

<ul class="noIndent">
	<li>bookmark m.joshdutcher.com for easier access</li>
	<li>this page is graphics intensive and may take a long time to load</li>
</ul>

<div id="ksn_radar">
	<!-- Begin pkg: iwradar rev3 -->
	<script language="JavaScript" type="text/javascript">

		<!-- Hide from JavaScript-Impaired Browsers
		if (document.images) {
			//Total frames represented by 18 followed by 17, frames 13-18 are repeats of 12
			isn=new Array();
			for (i=1;i<18;i++) {
				isn[i]=new Image();
				if (i<17){
					// Use 0 placeholder in filename when i is less than 10
					if (i<10){
						isn[i].src="http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_0"+i+".jpg";
					} else {
						if (i<13){
							isn[i].src="http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_"+i+".jpg";
						} else {
							isn[i].src="http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_12.jpg";
						}
					}
				} else {
					isn[i].src="/images/inv.gif";
				}
			}
		}

		//dla = default speed or delay between frames
		ctr=1;
		dla=150;
		j=1;
		stpit=0;
		restart=0;
	
		function startIt1(param){
			if (restart==1){
				stpit=0;
			}
			if (stpit<1){
				setTimeout("prtIt1()",dla);
				restart=0;
			}
		}

		function prtIt1(){
			if (document.images){
				document.ani1.src=isn[j].src;
				j++;
				if (j>16){
					j=1
				}
				if (stpit==1){
					stpit=0;
				}
				else{
					startIt1();
					}
				}
			else{
				alert("You need a JavaScript 1.1 compatible browser. We recommmend Netscape 3+ "
				+"or MSIE4+.");
			}
		}

	// End Hiding -->
	</script>
	<script type="text/javascript" language="JavaScript1.1" event="onLoad">
		startIt1();
	</script>

	<!--- here is the actual map image -->
	<img name="ani1" border=0 src="http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_01.jpg" alt="Radar image" width="300" />
</div>

<div id="basereflectivity">
	<img src="http://radar.weather.gov/lite/NCR/ICT_loop.gif" width="300" />
</div>

<div id="accuweather_radar">
	<img border="0" src="http://sirocco.accuweather.com/nx_mosaic_400x300c/SIR/inmaSIRKS_.gif" width="300" />
</div>

<div id="accuweather_sat">
	<img border="0" src="http://sirocco.accuweather.com/sat_mosaic_640x480_public/ei/isaeks_.gif" width="300" />
</div>

<div id="warnings">
	<img border="0" src="http://sirocco.accuweather.com/adc_images2/english/current/svrwx/400x300/isvrwxNE_.gif" width="300"><br/>
	<img border="0" src="http://sirocco.accuweather.com/web_images/svrwx/key/swskeys.gif" width="300">
</div>

<p><a href="http://www.joshdutcher.com/weather/">go to the main site</a></p>

</body>
</html>