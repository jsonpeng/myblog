
    <div class="blog-footer">
      <p>个人博客使用<a href="http://getbootstrap.com">Bootstrap</a>和<a href="http://beego.me">beego</a>by<a href="https://twitter.com/mdo">PengYun</a>.</p>
      <p>
        <a href="#">Back to top</a>
      </p>
    </div>

 </body>

    <!-- Bootstrap core JavaScript

    <!-- Placed at the end of the document so the pages load faster -->
   
    <script src="/static/js/jquery.min.js"></script>
    <script src="/static/js/bootstrap.min.js"></script>
    <!--<script src="/static/js/doc.min.js"></script>-->
    {{if .IsLiuyan}}
    <script src="/static/js/liuyan.js"></script>
    {{end}}
    {{if .IsAbout}}
    <script src="/static/js/markdown/markdown.js"></script>
    <script>
      function Editor(input, preview) {
        this.update = function () {
          preview.innerHTML = markdown.toHTML(input.value);
        };
        input.editor = this;
        this.update();
      }
      var $ = function (id) { return document.getElementById(id); };
      new Editor($("text-input"), $("preview"));
    </script>
    {{end}}
 
</html>