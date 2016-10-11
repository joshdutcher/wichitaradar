<?php
/**
 * The base configuration for WordPress
 *
 * The wp-config.php creation script uses this file during the
 * installation. You don't have to use the web site, you can
 * copy this file to "wp-config.php" and fill in the values.
 *
 * This file contains the following configurations:
 *
 * * MySQL settings
 * * Secret keys
 * * Database table prefix
 * * ABSPATH
 *
 * @link https://codex.wordpress.org/Editing_wp-config.php
 *
 * @package WordPress
 */

// ** MySQL settings - You can get this info from your web host ** //
/** The name of the database for WordPress */
define('DB_NAME', 'joshznjd_wp');

/** MySQL database username */
define('DB_USER', 'joshznjd_admin');

/** MySQL database password */
define('DB_PASSWORD', 'G5L|/&:XXA`Xn@f8e5!x');

/** MySQL hostname */
define('DB_HOST', 'localhost');

/** Database Charset to use in creating database tables. */
define('DB_CHARSET', 'utf8');

/** The Database Collate type. Don't change this if in doubt. */
define('DB_COLLATE', '');

/**#@+
 * Authentication Unique Keys and Salts.
 *
 * Change these to different unique phrases!
 * You can generate these using the {@link https://api.wordpress.org/secret-key/1.1/salt/ WordPress.org secret-key service}
 * You can change these at any point in time to invalidate all existing cookies. This will force all users to have to log in again.
 *
 * @since 2.6.0
 */
define('AUTH_KEY',         'Vu9v|X?,|tq_D7b6DibGuS%8uc=*lfsNgqSL],y)?v@/To>P;6wY&z<ynkiGG L7');
define('SECURE_AUTH_KEY',  '/}eu{_Qxd[.AzE/1:-:3M`<{G+}puTT`@Fg~}uE,`}KO8RB) ~|TU3x@D(rGJ9dx');
define('LOGGED_IN_KEY',    'oO5uh5E4u|HC,h0=S*XZg QHR*&0=*}[0_`L>+OfB{b1[9--f1?&8(ZWd9 f>>0g');
define('NONCE_KEY',        'Cr-jB%R/R}hQa1Af+hc|oG{EM.boydCdd,DXJlW(M_@Qh,U;CwJSFunsQ*&I-T%f');
define('AUTH_SALT',        'X8t!lJ=f:dvTKF&{:cr Mf*Z}jQ9xBAc@h,O]hz8%9EZk`_oG/uc|Xp+}vQy[98 ');
define('SECURE_AUTH_SALT', 'V=S$p E10r[yYdnk~q-0/xUtaOjC5%hF#F+e _4)#|OhEQ @gdPI:*Rq[P:B1uV@');
define('LOGGED_IN_SALT',   'Y:b<b{mEvlfhq,XeQ+O<lvS.`0*nRS-D!SojG$JKWg[Qrz+3;-r@L~roqsqeMKfQ');
define('NONCE_SALT',       'ikv@N*!B<-J<jX>Q|qht6W7o72++fNvGHxi-7z+eNPnZ:!{qc`YEyN,^@LJP!9OS');

/**#@-*/

/**
 * WordPress Database Table prefix.
 *
 * You can have multiple installations in one database if you give each
 * a unique prefix. Only numbers, letters, and underscores please!
 */
$table_prefix  = 'wp_';

/**
 * For developers: WordPress debugging mode.
 *
 * Change this to true to enable the display of notices during development.
 * It is strongly recommended that plugin and theme developers use WP_DEBUG
 * in their development environments.
 *
 * For information on other constants that can be used for debugging,
 * visit the Codex.
 *
 * @link https://codex.wordpress.org/Debugging_in_WordPress
 */
define('WP_DEBUG', false);

/* That's all, stop editing! Happy blogging. */

/** Absolute path to the WordPress directory. */
if ( !defined('ABSPATH') )
    define('ABSPATH', dirname(__FILE__) . '/');

/** Sets up WordPress vars and included files. */
require_once(ABSPATH . 'wp-settings.php');
