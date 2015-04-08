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

	function speedIt1(){
		if (stpit==0){
			if (dla>250){
				dla-=250;
				if (dla<150){
					dla=100;
				}
			}
			else{
				dla-=50;
				if (dla<50){
					dla-=25;
					if (dla<25){
						dla=25;
					}
				}
			}
		}
	}

	function slowIt1(){
		if (stpit == 0){
			if (dla<50){
				dla+=25;
			}
			else{
				if (dla<150){
					dla+=50;
				}
				else{
					if (dla<=2150){
						dla+=250;
						if (dla>2150){
							dla=2500;
						}
					}
				}
			}
		}
	}

	function testStp1(){
		if (stpit==0){
			stpit=1;
		}
	}

	function Back1(){
		stpit=1;
		restart=1;
		if (document.images){
			j--;
			if (j<1){
				j=12
			}
			else{
				if (j>12){
					j=1
				}
			}
		document.ani1.src=isn[j].src;
		}
	}

	function Forward1(){
		stpit=1;
		restart=1;
		if (document.images){
			j++;
			if (j>12){
				j=1
			}
		document.ani1.src=isn[j].src;
		}
	}
// End Hiding -->
</script>
<script type="text/javascript" language="JavaScript1.1">
	startIt1();
</script>

<!--- here is the actual map image -->
<img name="ani1" border=0 src="http://cache1.intelliweather.net/imagery/KSNW/rad_ks_wichita_640x480_01.jpg" alt="Radar image" width="600" height="450" />
