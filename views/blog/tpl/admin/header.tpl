<!DOCTYPE html>
<html>
<head>
{{if .IsLogin}}
<title>登录页面</title>
{{else if .Reg}}
<title>注册页面</title>
{{else}}
<title>博客后台管理</title>
{{end}}
    <!--图标-->
    <link rel="icon" href="/static/img/windows-live-writer.ico">
    <link rel="stylesheet" type="text/css" href="/static/css/bootstrap.min.css">
</head>
<body ng-app="myblog">
{{if .IsModify}}
{{else}}
<div class="navbar navbar-default">
  <div class="container">
 <!--  <a class="navbar-brand" href="/admin/index">博客后台</a> -->
<a class="navbar-brand" href="/">博客前端页面</a>
  <ul class="nav navbar-nav">
   <li {{if .IsHome}}class="active"{{end}}><a href="/admin/index">后台首页</a></li>
   <li {{if .IsCate}}class="active"{{end}}><a href="/admin/category">文章分类管理</a></li>
   <li  {{if .IsTopic}}class="active"{{end}}><a href="/admin/topic">文章发布管理</a></li>
   <li {{if .IsLiu}}class="active"{{end}}><a href="/admin/liuyan">博客留言管理</a></li>
  </ul>
  </div>  
</div>
{{end}}