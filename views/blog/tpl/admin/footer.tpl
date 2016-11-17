</body>
<script src="/static/js/jquery.min.js"></script>
<script src="/static/js/aj.js"></script>
<script src="/static/js/layer-2.4/layer.js"></script>
<script type="text/javascript">
  var app=angular.module('myblog',[]) 
  app.config(['$interpolateProvider', function($interpolateProvider) {
  $interpolateProvider.startSymbol('[[');
  $interpolateProvider.endSymbol(']]');
}]);

  app.controller('addtopic', function($scope) {
     $scope.alert=function(){
      layer.open({
  type:2,
  title: '添加文章',
  shadeClose: true,
  shade: 0.5,
  area: ['600px', '75%'],
  content: 'http://localhost:2020/admin/addtopic' //content的url
}); 
}
});
app.controller('login',['$scope','$http',function($scope,$http){
$scope.asave=function(){
  $http({
    method:'POST',
    url:'/login',
    data:$.param($scope.formData),
    headers:{'Content-Type':'application/x-www-form-urlencoded; charset=UTF-8'}
  }).success(function (data){
if(data.errno==0x00){
  layer.alert("登录成功!",{icon:3});
  location.href="/";
}else if(data.errno==0x12){
  layer.alert("用户名不存在!",{icon:5});
}else if(data.errno==0x14){
  layer.alert("密码不正确!",{icon:5});

}else if(data.errno==0x10){
  layer.alert("管理员你好!");
  location.href="/admin/index";

}

  });
};
}]);

app.controller('register',['$scope','$http',function($scope,$http){
 $scope.asave= function(){
         $http({
             method: 'POST',
             url:'/register',
             data:$.param($scope.formData), // pass in data as strings
             headers: {'Content-Type':'application/x-www-form-urlencoded; charset=UTF-8'}  
         }).success(function (data) {
    if(data.errno==0x00){
      layer.msg('注册成功', {
               offset: 0,
               shift: 6
});
      location.href="/login";
    }else if(data.errno==0x16){
      layer.alert("用户已存在,请重新注册",{icon:5});
    }
 
});
  };
  }]);

  

</script>
</html>