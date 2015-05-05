<!-- ******************************* -->
<!-- ****** KSN Radar ************** -->
<!-- ******************************* -->

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
<script type="text/javascript">
	setTimeout("prtIt1()",150);
</script>

<!--- here is the actual map image -->
<img name="ani1" border=0 class="pure-img-responsive" src="http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_01.jpg" alt="Radar image" />
