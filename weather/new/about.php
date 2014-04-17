<!DOCTYPE HTML PUBLIC "-//W3C//DTD HTML 4.01 Transitional//EN">
<html>
<head>
<title>Josh's Weather Station</title>
<meta http-equiv="Content-Type" content="text/html; charset=iso-8859-1" />
<meta http-equiv="pragma" content="no-cache" />
<meta http-equiv="refresh" content="300" />
<meta name="keywords" content="Weather, Radar, Satellite, Animation, Map, Watches, Warnings, Thunderstorms, Tornado, Storm warning, tornado warning, tornado watch" />
<meta name="description" content="Your one-stop shop for animated radar and satellite maps of area weather, including satellite, radar, watches and warnings, and power outages." />

<style type="text/css">
	body *{
		font-family: tahoma, verdana, arial, sans-serif;
		font-size: 12px;
		color: #999999;
	}

	a{
		text-decoration: none;
		color: #FFFFFF;
	}
	
	a:hover{
		text-decoration: underline;
	}
</style>
</head>

<?php
	switch($_GET['f']) {
		case 'ict':
			$url = 'http://wx.joshdutcher.com/ict/';
			break;
		case 'kc':
			$url = 'http://wx.joshdutcher.com/kc/';
			break;
		default:
			$url = 'http://wx.joshdutcher.com/';
	}
?>

<body bgcolor="#000000">
<table width="1000" border="0" cellspacing="0" cellpadding="1">
	<tr>
		<td align="right" style="font-family: tahoma, verdana, arial, sans-serif; color: #CCCCCC; font-size: 10px;"><a href="<?=$url?>" style="font-size: 10px;">back to weather</a></td>
	</tr>
	<tr>
		<td style="padding: 20px;">
			<p>My name is Josh Dutcher.  I originally put this weather page together so I could keep track of the weather if my satellite dish went out.
			I'm a web developer, so I just built it by harvesting source code from other sites.  I don't have permission
			to use any of these maps, but whatever.  I never
			anticipated it being used by anyone other than myself, so I just tossed it up on my server.  That's why it's not especially pretty - 
			it's just meant to be functional. I certainly don't make any money off it or anything like that - it's just a resource for
			people who find it useful.</p>
			
			<p>The site automatically refreshes every five minutes, which resets the tab you're viewing to the default radar tab. It does this
			because those radar animations get cached and need to be reloaded periodically in order to stay current. Auto refresh allows this
			to happen automatically without you having to hit F5 all the time. If you'd like it not to do that, like if you want to watch
			the twitter tab instead, just click the link at the top of the page that says "make it stop."</p>
			
			<p>As a general rule, most sites including this one work better in Firefox or Google Chrome than they do in Internet Explorer.</p>
			
			<p>Bookmark the site at <?=$url?></p>
			
			<p>I can be contacted at busychild424 &#97;&#116; &#103;&#109;&#97;&#105;&#108; dot com.</p>
		</td>
	</tr>
</table>

<script type="text/javascript">
	var gaJsHost = (("https:" == document.location.protocol) ? "https://ssl." : "http://www.");
	document.write(unescape("%3Cscript src='" + gaJsHost + "google-analytics.com/ga.js' type='text/javascript'%3E%3C/script%3E"));
</script>
<script type="text/javascript">
	try {
		var pageTracker = _gat._getTracker("UA-6815318-1");
		pageTracker._trackPageview();
	}
	catch(err) {}
</script>
</body>
</html>