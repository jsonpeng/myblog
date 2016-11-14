function sendMess(){ 
var infor=$('#information');//获取要显示留言部分的div
var username=$('#usename').html();//获取留言者的名字 
var message=document.getElementById("message").value;//获得留言的内容 
var showTime=document.createElement('h4');//创建显示时间的节点 
var con= document.createElement('div');//创建一个div 
con.setAttribute('class', 'div');//给div加上我们写好class 
var showName=document.createElement('h4')//创建显示留言名字的节点 
showName.setAttribute('class', 'h4');//加上样式 
var mess=document.createElement('p');//创建显示留言内容的节点 
mess.setAttribute('class', 'mess');//加上样式 
var mydate=new Date(); //新建一个时间对象 v
var time=(mydate.getFullYear()+'-'+(mydate.getMonth()+1)+'-'+mydate.getDay()+' '+mydate.getHours()+':'+mydate.getMinutes()+':'+mydate.getSeconds());//时间格式化 
showName.innerHTML=('留言人：'+username);//把获取的名字显示为我们创建的h4标签的内容 
mess.innerHTML=('留言内容：'+'<br/>'+message);//把获取的内容显示为我们创建的p标签的内容 
showTime.innerHTML=time;//显示时间 
con.appendChild(showName);//把我们创建的节点加上子节点 
con.appendChild(mess); 
con.appendChild(showTime); 
infor.append(con); 
}

