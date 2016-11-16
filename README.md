# 基于**go**语言框架beego设计的博客系统

------
## 1.0 beego简介
采用[beego](http://beego.me)以及[bootstrap](http://www.bootcss.com/)搭建:  
主要核心:  

> * angular.js双向数据绑定
> * ajax绑定
> * MVC设计模式


![cmd-markdown-logo](http://45.32.248.21/static/img/architecture.png)  
beego 是基于八大独立的模块构建的，是一个高度解耦的框架。当初设计 beego 的时候就是考虑功能模块化，用户即使不使用 beego 的 HTTP 逻辑，也依旧可以使用这些独立模块，例如：你可以使用 cache 模块来做你的缓存逻辑；使用日志模块来记录你的操作信息；使用 config 模块来解析你各种格式的文件。所以 beego 不仅可以用于 `HTTP` 类的应用开发，在你的 socket 游戏开发中也是很有用的模块，这也是 beego 为什么受欢迎的一个原因。大家如果玩过乐高的话，应该知道很多高级的东西都是一块一块的积木搭建出来的，而设计 beego 的时候，这些模块就是积木，高级机器人就是 beego。

## 1.1 执行逻辑
beego 是基于这些模块构建的，那么它的执行逻辑是怎么样的呢？beego 是一个典型的`MVC`架构，它的执行逻辑如下图所示：
![cmd-markdown-logo](http://45.32.248.21/static/img/flow.png) 

## 1.2 项目结构
该博客项目结构为:  
  
  
![cmd-markdown-logo](http://45.32.248.21/static/img/project.png)  
从上面的目录结构我们可以看出来 M（models 目录）、V（views 目录）和 C（controllers 目录）的结构， `main.go` 是入口文件。

## 2.0 博客开发流程
依照`view`->`model`-> `controller`的开发流程开发。

## 2.1 前端页面及静态文件设计
前端页面采用html5+bootstrap开发，数据交互采用angular.js来绑定`ajax`请求，绑定按钮的回调函数来执行`ajax`请求

## 2.2 使用beego框架搭建后台

## 2.3  使用angular. js实现前后台数据对接  

angular.js的`$http`包来实现前后台的数据对接,双向数据绑定实现表单验证

## 2.4 使用layer.js实现简单的页面弹窗效果 

## 3.0 写在后面  

![cmd-markdown-logo](http://45.32.248.21/static/img/golang.jpg)  
go语言是一门简洁高效的编程语言，通过使用go语言以及beego框架使我对orm数据映射以及mvc的设计模式更加深刻，也在未来的工作中更加得心应手!
