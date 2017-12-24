#!/bin/bash
apt update
apt install wget nano
#cd /home/nur/Documents/bash/roundcube
cp roundcubemail-1.1.3-complete.tar.gz /opt/
cd /opt/
tar xzf roundcubemail-1.1.3-complete.tar.gz
mv roundcubemail-1.1.3 roundcube
rm roundcubemail-1.1.3-complete.tar.gz
chown -R www-data:www-data /opt/roundcube

## Database
pass=$(whiptail --passwordbox "Input Your Password Database !" 8 78 --title "Database Password" 3>&1 1>&2 2>&3)
if [ $? = 0 ]; then

mysql --defaults-file=/etc/mysql/debian.cnf -e "CREATE DATABASE roundcubemail;"
mysql --defaults-file=/etc/mysql/debian.cnf -e "flush privileges;"
mysql --defaults-file=/etc/mysql/debian.cnf roundcubemail < /opt/roundcube/SQL/mysql.initial.sql

## roundcube == apache2
cd /opt/roundcube/config
cp -pf config.inc.php.sample config.inc.php
perl -pi -e "s/pass/$pass/g" config.inc.php
#cd /home/nur/Documents/bash/roundcube/
cp roundcube.conf /etc/apache2/conf-available/
a2enconf roundcube
service apache2 reload

### Result 
echo "Roundcube Successfully Installed" >> ini.txt
echo "You Must login from ipserver/webmail" >> ini.txt
echo "THANK FOR USE MY SCRIPT BASH" >> ini.txt
echo "NUR ROKHAINI" >> ini.txt
whiptail --textbox ini.txt 12 80
rm ini.txt
clear


else
echo "Abort Installation"
fi
