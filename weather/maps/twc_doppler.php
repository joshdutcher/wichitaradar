<!-- ******************************* -->
<!-- ****** TWC doppler ************ -->
<!-- ******************************* -->

<!-- version 07.15.04 modified by Tina Allen -->
<HTML>
<HEAD>
	<TITLE>weather.com - Map Navigator - Doppler Radar 600 Mile</TITLE>
	<!-- externals version 3.20.2006 @ 9:26 AM modified by WBF -->
<LINK REL="stylesheet" TYPE="text/css" HREF="http://ima.weather.com/web/common/header/stylesheet/style_sheet.css?032006" />
<LINK REL="alternate" TYPE="application/rss+xml" TITLE="The Weather Channel:  National Weather Outlook [RSS]"  HREF="http://www.weather.com/rss/national/rss_nwf_rss.xml" />
<LINK REL="alternate" TYPE="application/rss+xml" TITLE="The Weather Channel Blog [RSS]" HREF="http://blogs.weather.com/blog/weather/index.xml?from=autodiscovery" />
<META HTTP-EQUIV="MSThemeCompatible" Content="Yes" />
<meta http-equiv="imagetoolbar" content="no" />


<SCRIPT LANGUAGE="JavaScript1.2">
<!--
isMinNS4 = (document.layers) ? 1 : 0;
isMinIE4 = (document.all) ? 1 : 0;
isMinIE5 = (document.getElementById&&document.all) ? 1 : 0;
isNS6 = (document.getElementById&&!document.all) ? 1 : 0;
var popup;
var dateNow=new Date();
var haton=0;
function initialize_ad_array(adS)
{
  adS['uk.weather.com']=new Array('uk.weather.com','http://www.weather.com/RealMedia/a'+'ds/');
  adS['br.weather.com']=new Array('br.weather.com','http://www.weather.com/RealMedia/a'+'ds/');
  adS['fr.weather.com']=new Array('fr.weather.com','http://www.weather.com/RealMedia/a'+'ds/');
  adS['de.weather.com']=new Array('de.weather.com','http://www.weather.com/RealMedia/a'+'ds/');
  adS['espanol.weather.com']=new Array('espanol.weather.com','http://www.weather.com/RealMedia/a'+'ds/');
  adS['desktop3.weather.com']=new Array('desktop3.weather.com','http://www.weather.com/RealMedia/a'+'ds/');
  adS['desktop.weather.com']=new Array('desktop.weather.com','http://www.weather.com/RealMedia/a'+'ds/');
  adS['adstest.weather.com']=new Array('adstest.weather.com','http://adstest.weather.com/RealMedia/a'+'ds/');
  adS['photo.weather.com']=new Array('www.weather.com','http://www.weather.com/RealMedia/a'+'ds/');
  adS['www.w3.weather.com']=new Array('www.weather.com','http://www.w3.weather.com/RealMedia/a'+'ds/');
  adS['registration.weather.com']=new Array('registration.weather.com','https://registration.weather.com/RealMedia/a'+'ds/');
  adS['desktopfw.weather.com']=new Array('desktopfw.weather.com','http://www.weather.com/RealMedia/a'+'ds/');
}
// -->
</script>
<SCRIPT LANGUAGE="JavaScript1.2" SRC="http://ima.weather.com/common/header/javascript/ext.js" ></SCRIPT>
<SCRIPT LANGUAGE="JavaScript1.2" SRC="http://ima.weather.com/common/header/javascript/triggerParams.js" ></SCRIPT>
<SCRIPT LANGUAGE="JavaScript1.2" SRC="http://ima.weather.com/common/header/javascript/stdLauncher.js" ></SCRIPT>
<SCRIPT LANGUAGE="JavaScript1.2" SRC="http://ima.weather.com/common/header/javascript/divtools.js" ></SCRIPT><SCRIPT LANGUAGE="JavaScript1.2">
<!--

//
// <% /**
var remoteAddr="172.16.69.30";
// **/ %>

var adTest=GetCookie("oas_host_cookie");

function regenerate2() {
	return;
}

function initialize_dom_severe_scroll() {
	return;
}

function onPageStart() {
    regenerate2();
    if (isMinNS4 && document.images["holdspace"]) {
	    thisX = document.images["holdspace"].x;
	    thisY = document.images["holdspace"].y;
	    thisElement = makeObjectNS4();
    }
    if (document.getElementById) initialize_dom_severe_scroll();
}

window.onload = onPageStart;

var adS=new Array();
initialize_ad_array(adS);

OAS_url ='http://www.weather.com/RealMedia/a'+'ds/';
OAS_host = window.location.hostname;

if (OAS_host.indexOf('w3')>0)
{
	OAS_host = 'www.weather.com';
	OAS_url='http://www.w3.weather.com/RealMedia/a'+'ds/';
}else if (adS[OAS_host]){
     OAS_url=adS[OAS_host][1];
     OAS_host = adS[OAS_host][0];
}else{
   OAS_host='www.weather.com';
}

// special ads test code
if (adTest)
{
    if ((remoteAddr.indexOf("10.") == 0)||
        (remoteAddr.indexOf("169.254.") == 0)||
        (remoteAddr.indexOf("192.168.") == 0)||
        (remoteAddr.indexOf("216.133.140.1") == 0)||
		(remoteAddr.indexOf("216.133.140.2") == 0))
       {
	 OAS_host=adTest;
       }
  }

OAS_target = "_top";
OAS_version = 10;
OAS_rn = '001234567890';
OAS_rns = '1234567890';
OAS_rn = new String (Math.random()); OAS_rns = OAS_rn.substring(2,11);
function OAS_NORMAL(pos) {
if (OAS_MJX_on){
  document.write('<A HREF="' + OAS_url
+ 'click_nx.a'+'ds/nx.weather.com/noads/1' + OAS_rns + '@' + pos + '?keywords=no" TARGET=' + OAS_target
+ '>');
  document.write('<IMG SRC="' + OAS_url
+ 'adstream_nx.a'+'ds/nx.weather.com/noads/1' + OAS_rns + '@' + pos + '?keywords=no" BORDER=0></A>');
  }
}



// this'll get overwritten later on the page render,
// but if it doesn't, we're still OK

function OAS_RICH(pos) {
  OAS_NORMAL(pos);
}

//This was added to "touch" the UserPreferences and RMID cookies with every
//page view so that their expiration date changes.  If the cookie exists....
	var queryString = new Object;
	function parseParameter() {
	var temp_query = new RegExp ('^[^\\?]+\\?(.*)$');
	if ( ! temp_query.test(location) ) return false;
	var array = temp_query.exec(location);
	queryString.QUERY_STRING = array[1];
	var params = queryString.QUERY_STRING.split(/&/);
	for ( var i = 0; i < params.length; i++ ) {
		var keys = params[i].split(/=/);
		queryString[ keys[0] ] = unescape(keys[1]);
		}
	}
	function paramValue(key) {
	if ( key == null ) {
		alert("param() function has been used incorrectly.\nUSAGE: param(key)");
		return false;
	}
		return queryString[key];
	}
	parseParameter();

//This passed age and gender as OAS parameters
if(typeof(OAS_query) == 'undefined') {
// no OAS Query defined..cannot put age+gender in the query string
} else {
//partner cookie information
var partnerCookie = GetCookie("partner");
var idType = searchTermType();
// get age and gender if they are available...
var byear = getUserPreferences("13");
var age = 0;
var gender = getUserPreferences("14");
if(byear.length > 0 && gender.length > 0) {
var oq = OAS_query;
var now = new Date();
var thisYear = now.getFullYear();
age = thisYear - Number(byear);
if(oq == "") {
oq = "age=" + age + "&gender=" + gender.toLowerCase();
} else {
oq = oq + "&age=" + age + "&gender=" + gender.toLowerCase();
}
if (partnerCookie != ''){
	oq = oq + "&cobrand="+partnerCookie;
}

if(idType != ''){
	oq = oq + "&idtype="+idType;
}

OAS_query = oq;
}
}



	var vsearch = paramValue('search');
	var upcookie = GetCookie("UserPreferences");
	if(upcookie > 0 && vsearch != "search"){
		updateCookieExpDate("UserPreferences");
		updateCookieExpDate("RMID");
	}

var customization_pathname = (window.location.pathname.indexOf("/weather/my")>=0)?1:0;

//The below logic should only process on non-Customization pages.
if (customization_pathname == 0){
	var myPrefsCookie = GetCookie("MyPrefs");
	if(myPrefsCookie.length > 1){
		updateCookieUnescape("MyPrefs");
	}
}
//-->
</SCRIPT>
<SCRIPT LANGUAGE="JavaScript1.1">
<!--
OAS_version = 11;
if (!isMinIE4 && !isMinIE5 && !isNS6 && !isMinNS4) OAS_version = 10;
if (OAS_MJX_on){
if (OAS_version >= 11) document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.1 SRC="' + OAS_url + 'adstream_mjx.a'+'ds/' + OAS_host + OAS_spoof + '/1' + OAS_rns + '@' + OAS_listpos + '?' + OAS_query + '"><\/SCRIPT>');
}
//-->
</SCRIPT>
<SCRIPT LANGUAGE="JavaScript">
<!--
function OAS_AD(pos) {
		(OAS_version >= 11) ? OAS_RICH(pos) : OAS_NORMAL(pos);
}
//-->
</SCRIPT>
<SCRIPT LANGUAGE="JavaScript1.2" SRC="http://ima.weather.com/common/header/javascript/eluminate.js" ></SCRIPT>
<SCRIPT LANGUAGE="JavaScript1.2" SRC="http://ima.weather.com/common/header/javascript/techprops.js" ></SCRIPT>
<SCRIPT LANGUAGE="JavaScript1.2">
/* LAST UPDATED:  1/24/2006  By WBF */
cm_TrackLink = "A";
var cmSampleList = "|DRV_INTST_PIF_36HR|GEN_INTERACT_CONT_PHOTO|GEN_MAP_LRG|GEN_MAP_MPRM|HLTH_ALGY_PIF_36HR|HNG_G_PIF_36HR|HNG_H_PIF_36HR|HP_HOMEPAGE|HP_PERS_36HR|NEWS_FAM_VID_BREAKING|NEWS_FAM_VID_LOCAL|NEWS_TROP_CONT_TROP|RECR_GOLF_PIF_36HR|SEARCH|TRVL_BTRAV_PIF_10DAY|TRVL_BTRAV_PIF_36HR|TRVL_BTRAV_PIF_SIXTEN|TRVL_VAC_PIF_36HR|UNDCL_PIF_10DAY|UNDCL_PIF_36HR|UNDCL_PIF_DETAIL|UNDCL_PIF_HRLY|UNDCL_PIF_HRLY|UNDCL_PIF_MAP|UNDCL_PIF_SIXTEN|UNDCL_PIF_WKEND|UNKNOWN PAGE NAME|";
var cmSample = cmCheckSampleClick();
var cmSampleRate = 20;

function C9(e){
	cGI="";
	cGJ="";
	cGK="";
	var type=e.tagName.toUpperCase();
	if(type=="AREA"){
		cGJ=e.href?e.href:"";
		var p=e.parentElement?e.parentElement:e.parentNode;
		if(p!=null)
			cGI=p.name?p.name:"";
	}
	else{
		while(type!="A"&&type!="HTML"){
			if(!e.parentElement)
				e=e.parentNode;
			else
				e=e.parentElement;
		if(e)
			type=e.tagName.toUpperCase();
	}
		if(type=="A"){
			cGJ=e.href?e.href:"";
			cGI=e.name?e.name:"";
		}
	}
	cGJ=cG7.normalizeURL(cGJ,true);
	if(cV(cGJ)==true){
		var dt=new Date();
		cGK=dt.getTime();
		if (cmSample) {
			cM(cm_ClientTS,cGK,cGI,cGJ,false);
		}
	}
	else{
		cGJ="";
	}
}

function cmCheckSampleClick() {

	// check if URL in list
	var cm_pgid = c1(cm_ClientID);
	cm_pgid = cm_pgid.toUpperCase();
	if (cmSampleList.indexOf("|" + cm_pgid + "|") > -1) {

		// check if should sample
		var tempCookie = cI("RMID");
		var tempNumber = parseInt(tempCookie, 16);

		if ((tempNumber - Math.floor(tempNumber / cmSampleRate) * cmSampleRate) == 0) {
			return true;
		}
		return false;
	}
	return true;
}


var _pattern = /^[0-9]{5}$/;
var cmJv = "1.2";
var CM_page_id = "";
var CM_cat_id = "";
var CM_country = "";
var CM_state = "";
var CM_dma = "";
var CM_what_search = "";
var CM_where_search = "";
var OAS_spoof;
var CM_default_tag = "_not_yet_identified";
var CM_urs_id = getUserPreferences(2);
var CM_rmid = GetCookie('RMID');
var CM_partner = GetCookie('partner');
var CM_topProdID = "";
var CM_detailProdID = "";
var CM_prodViewTag = "";
var CM_shop5Tag = "";
var CM_shop9Tag = "";
var hstname = new String(window.location.hostname);

var query = new Object;
function parse() {
	var pat_query = new RegExp ('^[^\\?]+\\?(.*)$');
	if ( ! pat_query.test(location) ) return false;
	var array = pat_query.exec(location);
	query.QUERY_STRING = array[1];
	var params = query.QUERY_STRING.split(/&/);
	for ( var i = 0; i < params.length; i++ ) {
		var keys = params[i].split(/=/);
		query[ keys[0] ] = unescape(keys[1]);}}
function param(key) {
	if ( key == null ) {
		alert("param() function has been used incorrectly.\nUSAGE: param(key)");
		return false;}
	return query[key];}
parse();

var cm_OSK = "";

/* Coremetrics Tag v3.1, 2/28/2002
 * COPYRIGHT 1999-2002 COREMETRICS, INC.
 * ALL RIGHTS RESERVED. U.S.PATENT PENDING
 * The following functions aid in the creation of Coremetrics data tags. */
/* Creates a Pageview tag with the given Page ID */
function cmCreatePageviewTag(pageID, categoryID, rmID, ursID,whereSearchString,partner, dma, country, state, topProdID, detailProdID, prodViewTag,shop5Tag,shop9Tag){
      var cm;
      var sendTP = cI("tp");
      if (!sendTP) {cm = new _cm("tid", "6", "vn2", "e3.1");
        cm.addTP();
        document.cookie = "tp=Y";
      } else {cm = new _cm("tid", "1", "vn2", "e3.1"); }
if (GetCookie("partner") == "beta2"){cm.pi = "beta_" + pageID;
} else {cm.pi = pageID;}
      cm.cg = categoryID;
      cm.pv1 = rmID;
      cm.pv2 = ursID;
      cm.pv3 = partner;
      cm.pv4 = dma;
      cm.pv5 = country;
      cm.pv6 = state;
      cm.pv7 = topProdID;
      cm.pv8 = detailProdID;
      cm.pc = "Y";
      cm.se = whereSearchString;
      cm.writeImg();
      if (prodViewTag != "") {cmCreateProductviewTag(prodViewTag,pageID);}
      if (shop5Tag != "") {cmCreateShopAction5Tag(shop5Tag, shop5Tag, 1, 0,shop5Tag);}
      if (shop9Tag != "") {
        var currDay = new Date();
		var orderID = rmID+currDay.getYear()+(currDay.getMonth()+1)+currDay.getDate()+currDay.getHours()+currDay.getMinutes();
	  	cmCreateShopAction9Tag(shop9Tag, 1, 0, rmID,orderID);}
}
/* Creates a Productview Tag
 * Also creates a Pageview Tag by setting pc="Y"
 * Format of Page ID is "PRODUCT: <Product Name> (<Product ID>)"
 * productID      : required. Product ID to set on this Productview tag
 * productName    : required. Product Name to set on this Productview tag
 * categoryID     : optional. Category ID to set on this Productview tag
 * returns nothing, causes a document.write of an image request for this tag. */
function cmCreateProductviewTag(productID,productName) {
      var cm = new _cm("tid", "5", "vn2", "e3.1");
      if (productName == null) {productName = "";}
      cm.pr = productID;
      cm.pm = cm.pr;
      cm.cg = cm.pr;
      cm.pc = "Y";
      cm.pi = "PRODUCT: " + productName;
      cm.writeImg();
}
/* Creates a Shop tag with Action 5 (Shopping Cart)
 * productID      : required. Product ID to set on this Shop tag
 * quantity : required. Quantity to set on this Shop tag
 * productPrice   : required. Price of one unit of this product
 * categoryID     : optional. Category to set on this Shop tag
 * returns nothing, causes a document.write of an image request for this tag. */
function cmCreateShopAction5Tag(productID, productName, productQuantity,productPrice, categoryID){
      var cm = new _cm("tid", "4", "vn2", "e3.1");
      cm.at = "5";
      cm.pr = productID;
      cm.pm = productName;
      cm.qt = productQuantity;
      cm.bp = productPrice;
      if (categoryID) {cm.cg = categoryID;}
      cm.writeImg();
}
/* Creates a Shop tag with Action 9 (Order Receipt / Confirmed)
 * productID      : required. Product ID to set on this Shop tag
 * productName    : required. Product Name to set on this Shop tag
 * quantity       : required. Quantity to set on this Shop tag
 * productPrice   : required. Price of one unit of this product
 * customerID     : required. ID of customer making the purchase
 * orderID        : required. ID of order this lineitem belongs to
 * orderTotal     : required. Total price of order this lineitem belongs to
 * categoryID     : optional. Category to set on this Shop tag
 * returns nothing, causes a document.write of an image request for this tag.*/
function cmCreateShopAction9Tag(productID, productQuantity, productPrice, customerID, orderID) {
      var cm = new _cm("tid", "4", "vn2", "e3.1");
      cm.at = "9";
      cm.pr = productID;
      cm.qt = productQuantity;
      cm.bp = productPrice;
      cm.cd = customerID;
      cm.on = orderID;
      cm.cg = cm.pr;
      cm.tr = cm.bp;
      cm.writeImg();
      cm = new _cm("tid", "3", "vn2", "e3.1");
      cm.on = orderID;
      cm.tr = productPrice;
      cm.osk = "|" + productID + "|" + productPrice + "|1|";
      cm.sg = "0";
      cm.cd = customerID;
      cm.writeImg();
}
</SCRIPT>

<SCRIPT LANGUAGE="JavaScript1.2" SRC="http://ima.weather.com/common/header/javascript/cm_dma.js" ></SCRIPT>
<SCRIPT LANGUAGE="JavaScript1.2">
var CM_tag;
var pageName = new String(window.location.pathname);
var abStr = pageName.substring(0, 6);
if ((hstname == "registration.weather.com") ||(hstname == "notify.weather.com") || (hstname == "alerts.weather.com")){
 document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.2 SRC="http://ima.weather.com/common/header/javascript/cm_urs.js" ><\/SCRIPT>');
} else if (hstname == "photo.weather.com"){
 document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.2 SRC="http://ima.weather.com/common/header/javascript/cm_misc.js" ><\/SCRIPT>');
} else if ((abStr == "/outlo") || (abStr == "/activ")){
          abStr = pageName.substring(0, 15);
          if ((abStr == "/activities/rec") || (abStr == "/outlook/recrea")){document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.2 SRC="http://ima.weather.com/common/header/javascript/cm_recreation.js" ><\/SCRIPT>');}
          else if ((abStr == "/activities/hea") || (abStr == "/outlook/health")){document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.2 SRC="http://ima.weather.com/common/header/javascript/cm_health.js" ><\/SCRIPT>');}
          else if ((abStr == "/activities/tra") || (abStr == "/outlook/travel")){document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.2 SRC="http://ima.weather.com/common/header/javascript/cm_travel.js" ><\/SCRIPT>');}
          else if ((abStr == "/activities/dri") || (abStr == "/outlook/drivin")){document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.2 SRC="http://ima.weather.com/common/header/javascript/cm_driving.js" ><\/SCRIPT>');}
          else if ((abStr == "/activities/hom") || (abStr == "/outlook/homean")){document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.2 SRC="http://ima.weather.com/common/header/javascript/cm_homeandgarden.js" ><\/SCRIPT>');}
          else if ((abStr == "/activities/eve") || (abStr == "/outlook/events")){document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.2 SRC="http://ima.weather.com/common/header/javascript/cm_events.js" ><\/SCRIPT>');}
          else {document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.2 SRC="http://ima.weather.com/common/header/javascript/cm_undcld.js" ><\/SCRIPT>');}
} else if (abStr == "/servi"){
          document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.2 SRC="http://ima.weather.com/common/header/javascript/cm_urs.js" ><\/SCRIPT>');
} else if (abStr == "/weath"){
          document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.2 SRC="http://ima.weather.com/common/header/javascript/cm_undcld.js" ><\/SCRIPT>');
} else {
          document.write('<SCR' + 'IPT LANGUAGE=JavaScript1.2 SRC="http://ima.weather.com/common/header/javascript/cm_misc.js" ><\/SCRIPT>');
}
</SCRIPT>

<script language="JavaScript1.2">
addEvent(window,'load',mapInitLoad);
var images = new Array();
var thisMap = ['/looper/archive/us_ddc_closeradar_large_usen/'];
var loaded = 0;
var timerRunning = false;
var animSpeed = 400;
var animCounter = 0;
var imagenames = new Array();
var currentImage = 0;
var uniquej = new Date();
var uniquei = "?"+uniquej.getTime();
var thisloc = document.location;
setTimeout('reloadpage()' ,599000);
function reloadpage() {
	stopMap();
	var op = "The images you are viewing may be outdated.<p>\nPlease use the link below to restart the animation with current images.<p>\n<a href=\"javascript:{document.location='"+thisloc+"'}\">Restart Animation</a><p>\n";
	document.open();
	document.write(op);
	document.close();
}
function stopMap(){
	if(timerRunning) clearInterval(timerID);
	timerRunning = false;
}
function startMap(){
	stopMap();
	timerRunning = true;
	timerID = setInterval('animMap()',animSpeed);
}
function imageLoaded() {
	document.images['severeMap'].myComplete = true;
}
function loadImages(thisNum){
	imagenames = new Array( 'http://image.weather.com'+thisMap[thisNum]+'1L.jpg'+uniquei,'http://image.weather.com'+thisMap[thisNum]+'2L.jpg'+uniquei,'http://image.weather.com'+thisMap[thisNum]+'3L.jpg'+uniquei,'http://image.weather.com'+thisMap[thisNum]+'4L.jpg'+uniquei,'http://image.weather.com'+thisMap[thisNum]+'5L.jpg'+uniquei);
	loaded = 0;
	document.images['severeMap'].myComplete = true;
	document.images['severeMap'].onload = imageLoaded;
	for(n=0;n<imagenames.length;n++){
		images[n]=new Image();
		images[n].src=imagenames[n];
	}
	loaded = imagenames.length;
	startMap();
}
function mapInitLoad() {
	loadImages(0);
}
function animMap() {
	if (!document.images['severeMap'].myComplete) {
		return;
	}
	if (animCounter < 6) animCounter++;
	else animCounter = 0;
	var thisCounter = 0;
	if (animCounter == 5) thisCounter = 4;
	else if (animCounter == 6) thisCounter = 0;
	else thisCounter = animCounter;
	if (thisCounter != currentImage) {
		document.images['severeMap'].myComplete = false;
		document.images['severeMap'].src = images[thisCounter].src;
		currentImage = thisCounter;
	}
}
function animRMap() {
	if (!document.images['severeMap'].myComplete) return;
	if (animCounter > 0) animCounter--;
	else animCounter = 6;
	var thisCounter = 0;
	if (animCounter == 5) thisCounter = 4;
	else if (animCounter == 6) thisCounter = 0;
	else thisCounter = animCounter;
	if (thisCounter != currentImage) {
		document.images['severeMap'].myComplete = false;
		document.images['severeMap'].src = images[thisCounter].src;
		currentImage = thisCounter;
	}
}
function nextMap() {
	if (animCounter < 4) animCounter += 1;
	else animCounter = 0;
	document.images['severeMap'].src = images[animCounter].src;
	currentImage = animCounter;
}
function prevMap() {
	if (animCounter > 4) animCounter = 4;
	if (animCounter > 0) animCounter -= 1;
	else animCounter = 4;
	document.images['severeMap'].src = images[animCounter].src;
	currentImage = animCounter;
}
function startRMap(){
	stopMap();
	timerRunning = true;
	timerID = setInterval('animRMap()',animSpeed);
}
</script>
</HEAD>

<BODY BGCOLOR="#FFFFFF" LEFTMARGIN="0" TOPMARGIN="0" MARGINHEIGHT="0" MARGINWIDTH="0" LINK="#004371" VLINK="#004371">
<TABLE WIDTH=600 BORDER=0 CELLPADDING=0 CELLSPACING=0><TR>
<TD WIDTH="100%">

<img src="http://image.weather.com/web/legends/autoload_600x405.gif" alt="" width="600" height="405" border="0" id="severeMap" name="severeMap" /></TD></TR><FORM METHOD="post" ACTION="<!-- none specified -->"><TR><TD WIDTH="100%" CLASS="largeMap2">
<input type="button" value="<<" onClick="if(loaded!=0){animSpeed=400;startRMap();}" />&nbsp;&nbsp;&nbsp;

<input type="button" value=" < " onClick="if(loaded!=0){stopMap();prevMap();}" />&nbsp;&nbsp;&nbsp;
<input type="button" value=" stop " onClick="if(loaded!=0)stopMap();" />&nbsp;&nbsp;&nbsp;
<input type="button" value=" > " onClick="if(loaded!=0){stopMap();nextMap();}" />&nbsp;&nbsp;&nbsp;
<input type="button" value=">>" onClick="if(loaded!=0){animSpeed=400;startMap();}" /><p>
</p></TD></TR></FORM></TABLE>
</BODY>
</HTML>