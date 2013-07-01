<!--
  User: guoyao
  Time: 07-01-2013 17:25
  Blog: http://www.guoyao.me
-->

<html>
<head>
    <title></title>
</head>
<body>
<form action="/login?username=guoyao" method="post">
    用户名:<input type="text" name="username">
    密码:<input type="password" name="password">
    <input type="hidden" name="token" value="{{.}}">
    <input type="submit" value="登陆">
</form>
</body>
</html>