A CA file has been bootstrapped using certificates from (brew --prefix mysql)/bin/mysqladmin -u root password NEWPASS

 mysqld --init-file=Users/jjacob151/mysql-init &


kill `cat /mysql-data-directory/host_name.pid`


For compilers to find openssl you may need to set:
  export LDFLAGS="-L/usr/local/opt/openssl/lib"
  export CPPFLAGS="-I/usr/local/opt/openssl/include"

For pkg-config to find openssl you may need to set:
  export PKG_CONFIG_PATH="/usr/local/opt/openssl/lib/pkgconfig"

==> mysql
We've installed your MySQL database without a root password. To secure it run:
    mysql_secure_installation

MySQL is configured to only allow connections from localhost by default

To connect run:
    mysql -uroot
<hr>
to set the password
mysqladmin -u root password 'sorrie123'
Password:
mysqladmin: connect to server at 'localhost' failed

mysql -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 10
Server version: 8.0.11 MySQL Community Server - GPL

Copyright (c) 2000, 2019, Oracle and/or its affiliates. All rights reserved.

Oracle is a registered trademark of Oracle Corporation and/or its
affiliates. Other names may be trademarks of their respective
owners.

Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 

After all that the easy way is to download installation package here:
https://dev.mysql.com/get/Downloads/MySQL-8.0/mysql-8.0.17-macos10.14-x86_64.dmg

