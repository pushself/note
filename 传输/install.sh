make clean
./configure --with-php-config=/web/app/php7/bin/php-config
make
make install-all
pkill -f "php-f"
/web/app/php7/sbin/php-fpm
