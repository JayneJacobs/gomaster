# Go Database


# SQL Read Modify Crete Delete
http://go-database-sql.org/index.html
https://golang.org/pkg/database/sql/
https://github.com/go-sql-driver/mysql


- Generic layer for SQL DBs in go
- Must be used with database driver
- DB/SQL common interfaces/methods to access SQL
    - includes basic connection pool

``` go
import "database/sql"
import _ "github.com/go-sql-driver/mysql"

db, err := sql.Open("mysql", "user:password@/dbname")
[username[:password]@][protocol[(address)]]/dbname[?param1=value1&...&paramN=valueN]
```

//Use the default protocol (tcp) and host (localhost:3306):

```sql
db, err = user:password@/dbname

mysql> CREATE USER 'hydra'@'%' IDENTIFIED BY 'hydraisme';
Query OK, 0 rows affected (0.01 sec)

mysql> GRANT All PRIVILEGES ON Hydra.* TO 'hydra'@'%';
Query OK, 0 rows affected (0.01 sec)

mysql> FLUSH PRIVILEGES;
Query OK, 0 rows affected (0.01 sec)

mysql> SELECT * FROM Hydra.Personnel where id = 2;
+----+------+-------------------+----------+
| id | Name | SecurityClearance | Position |
+----+------+-------------------+----------+
|  2 | Jim  |                 2 | Bass     |
+----+------+-------------------+----------+
1 row in set (0.00 sec)
```


1. Un install using rm *mysql in path
2. reinstall using the mysql.dmg see [Notes from Install](NotesFromInstall.md)
3. set path variables to /usr/local/mysql


Example querys: 

```sql
INSERT INTO `Hydra`.`Personnel` (`id`, `Name`, `SecurityClearance`, `Position`) VALUES ('5', 'Rich', '1', 'Guitar');
CREATE USER 'hydra'@'%' IDENTIFIED BY 'hydraisme';
GRANT All PRIVILEGES ON Hydra.* TO 'hydra'@'%';
FLUSH PRIVILEGES;
```
