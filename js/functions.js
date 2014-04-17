$(document).ready(function(){
	// Cache the Window object
	$window = $(window);




	/*
	----------------------------------------------------------
	**     Parallax Scrolling                               **
	**     Author: Mohiuddin Parekh                         **
	**     http://www.mohi.me                               **
	**     @mohiuddinparekh                                 **
	----------------------------------------------------------
	*/

	$('section[data-type="background"]').each(function(){
		var $bgobj = $(this); // assigning the object

		$(window).scroll(function() {

			// Scroll the background at var speed
			// the yPos is a negative value because we're scrolling it UP!								
			// var yPos = -($window.scrollTop() / $bgobj.data('speed'));
			var yPos = -( ($window.scrollTop() - $bgobj.offset().top) / $bgobj.data('speed'));

			// Put together our final background position
			var coords = '50% '+ yPos + 'px';

			// Move the background
			$bgobj.css({ backgroundPosition: coords });

		}); // window scroll Ends
	});	

	// Create HTML5 elements for IE's sake

	// document.createElement("article");
	document.createElement("section");







	// init the typewriter text
	setOffset();
	$('.type-me').typewriter({
		auto: false,
		interval: 100,
		caret: {
			visible: true,
			blink: true,
			interval: 1000
		}
	});

	$(window).scroll(function() {
		checkType();
	})



	/*
	**********************************************************
	***    stuff for loading Treehouse badges              ***
	**********************************************************
	*/
	var e = "joshdutcher",

	// Treehouse Json 
	t = "http://teamtreehouse.com/" + e + ".json",

	// Badges JQuery Identifier    
	n = $("#badges"),

	// Badges Array    
	r = [],

	// Badges Count
	i = 0;

	// Json Parse Treehouse User Badges Info
	$.getJSON(t, function (e) {
		// User Json Parse Select Badges Info
		var t = e.badges;

		// Format Each badge's HTML
		$.each(t, function (e, t) {
			r += '<li><a href="' + t.url + '" target="_blank"><img src="' + t.icon_url + '" alt="' + t.name + '" title="' + t.name + '"/></a></li>';
			i++
		});

		// Append Badge to #badges
		n.append(r);

		// Header Badges count generator
		$("#treehouse-count").append('I have earned ' + i + ' badges at Treehouse!');
	});










}); 


function checkType() {
	var $nextEl = $(".type-me[data-status='ready']:first")
	if ( $nextEl.data('offset') <= $(window).scrollTop() + ($nextEl.data('offset-padding') || 0) && $nextEl.data('status') == 'ready' ) {
		$nextEl.attr( 'data-status', 'typed' ).typewriter('type');
	}
}

function setOffset() {
	$('.type-me').each(function() {
		var offSet = $(this).offset().top; 
		$(this).data('offset', offSet).attr({'data-status': 'ready'}); 
	})
}
