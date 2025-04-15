(function () {
  // Get menu elements
  var layout = document.getElementById('layout');
  var menu = document.getElementById('menu');
  var menuLink = document.getElementById('menuLink');

  function toggleClass(element, className) {
    var classes = element.className.split(/\s+/);
    var length = classes.length;
    var i = 0;

    for (; i < length; i++) {
      if (classes[i] === className) {
        classes.splice(i, 1);
        break;
      }
    }
    // The className is not found
    if (length === classes.length) {
      classes.push(className);
    }

    element.className = classes.join(' ');
  }

  function toggleMenu() {
    // Add 'active' class to both toggleMenu and menu
    toggleClass(layout, 'active');
    toggleClass(menu, 'active');
    toggleClass(menuLink, 'active');
  }

  // Event handler for menu clicks
  menuLink.addEventListener('click', function (e) {
    e.preventDefault();
    toggleMenu();
  });

  // Close menu when clicking outside
  document.addEventListener('click', function (e) {
    // If menu is active and click is outside menu and not on menu toggle
    if (
      layout.className.indexOf('active') !== -1 &&
      !menu.contains(e.target) &&
      e.target !== menuLink &&
      !menuLink.contains(e.target)
    ) {
      toggleMenu();
    }
  });
})();
