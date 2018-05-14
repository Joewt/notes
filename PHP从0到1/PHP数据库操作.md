# PHP数据库操作

### mysql
mysql  -基本废弃  
连接数据库  
mysql_connect($server, $username, $pwd)

选择数据库  
mysql_select_db($database_name)  

设置字符集  
mysql_set_charset($charset)  

执行sql语句  
mysql_query($query)

获取结果  
mysql_fetch_***

### mysqli

面向过程方式连接数据库  
$connect = mysqli_connect('host','username','pwd','database')  


执行sql语句  
$result = mysqli_query($connect, $sql)  

获取结果集  

mysqli_fetch_all($result)``

### pdo  

$dbh = new PDO('mysql:host=localhost;dbname=test', $user, $pass);
