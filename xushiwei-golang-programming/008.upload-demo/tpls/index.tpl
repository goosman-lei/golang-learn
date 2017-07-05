<html>
<head>
    <script type="text/javascript" src="/static/jquery-2.2.3.js"></script>
</head>
<body>
    <form action="/index" method="POST" enctype="multipart/form-data">
        <input type="file" name="upload_file" /><br />
        <input type="submit" value="上传"/>
    </form>
    {%upload_success%}
    <br />
    <h1>FileList</h1>
    {%file_list%}
</body>
</html>
