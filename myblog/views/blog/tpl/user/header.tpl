<!DOCTYPE html>
<html lang="en">
  <head>
    <meta charset="utf-8">
    <meta http-equiv="X-UA-Compatible" content="IE=edge">
    <meta name="viewport" content="width=device-width, initial-scale=1">
    <meta name="description" content="">
    <meta name="author" content="">
    
<link rel="icon" href="/static/img/blog.ico">
    <title>Peng的博客</title>

    <!-- Bootstrap core CSS -->
    <link href="/static/css/bootstrap.min.css" rel="stylesheet">

    <!-- Custom styles for this template -->
    <link href="/static/css/blog.css" rel="stylesheet">
    {{if .IsShouji}}
   <link href="/static/css/shouji.css" rel="stylesheet">
  {{end}}
  </head>

  <body>

    <div class="blog-masthead">
      <div class="container">
        <nav class="blog-nav">
 {{if .IsIndex}}<a  class="blog-nav-item active" href="/">{{else}}<a  class="blog-nav-item" href="/"> 
 {{end}}博客主页</a>
{{if .IsCategory}}<a class="blog-nav-item active" href="/category">{{else}}<a class="blog-nav-item" href="/category">{{end}}文章分类</a>
{{if .IsShouji}} <a class="blog-nav-item active" href="/shouji">{{else}} <a class="blog-nav-item" href="/shouji">{{end}}手记</a>
{{if .IsLiuyan}}<a class="blog-nav-item active" href="/liuyan">{{else}}<a class="blog-nav-item" href="/liuyan">
{{end}}留言板</a>
{{if .IsAbout}}<a class="blog-nav-item active" href="/about">{{else}}<a class="blog-nav-item" href="/about">
{{end}}关于</a>
        </nav>
      </div>
    </div>