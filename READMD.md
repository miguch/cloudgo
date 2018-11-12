本应用是一个简单的登录注册应用，使用 Go 编写Web 服务，浏览器前端代码使用了一年前 Web开发课程的代码。应用中使用的 RESTful API 如下:

| API                 | 说明                                                         |
| ------------------- | ------------------------------------------------------------ |
| GET /?username=name | 若用户名有效，进入用户信息页，否则进入登录页                 |
| POST /signin        | 参数为username 与 password，若登陆成功返回200状态码          |
| GET /regist         | 进入注册页面                                                 |
| POST /regist        | 参数为 register，若为 true 则需提供用户所有数据字段进行注册，否则提供某个字段用于检测是否存在重复，若无重复返回200 状态码 |

应用中使用gorilla/mux作为路由，kataras/go-template 用于将 pug 模板生成 HTML，urfave/negroni 构建中间件处理链。

## 使用 Curl 测试

1. 获取默认界面：

```shell
$ curl 127.0.0.1:8080/

<!DOCTYPE html>
<html lang='en'>
    <head>
        <meta charset='UTF-8'>
        <title>User Sign Up</title>
        <link rel='stylesheet' type='text/css' href='style.css'>
        <script type='text/javascript' src='https://cdn.jsdelivr.net/npm/jquery@3.3.1/dist/jquery.min.js'>
        </script>
        <script type='text/javascript' src='signin.js'>
        </script>
    </head>
    <body>
        <h1>登陆</h1>
        <div id="main-area">
            <h2>用户登陆</h2>
            <button id="signup" class="button" href='/regist'>注册</button>
            <form method='post'>
                <p>
                    <label>用户名:</label>
                    <input name='username' type='text' placeholder='用户名'>
                </p>
                <p id="username-checker" class="checker"></p>
                <p>
                    <label>密码:</label>
                    <input name='password' type='password' placeholder='密码'>
                </p>
                <p id="password-checker" class="checker"></p>
            </form>
            <button id="submit" class="button">登陆</button>
            <button id="reset" class="button">重置</button>
            <p id="msg" class="invalid"></p>
        </div>
    </body>
</html>
```

2. 进入注册页面：

   ```shell
   $ curl 127.0.0.1:8080/regist
   
   <!DOCTYPE html>
   <html lang='en'>
       <head>
           <meta charset='UTF-8'>
           <title>User Sign Up</title>
           <link rel='stylesheet' type='text/css' href='style.css'>
           <script type='text/javascript' src='https://cdn.jsdelivr.net/npm/jquery@3.3.1/dist/jquery.min.js'>
           </script>
           <script type='text/javascript' src='signup.js'>
           </script>
       </head>
       <body>
           <h1>注册</h1>
           <div id="main-area">
               <h2>用户注册</h2>
               <form method='post'>
                   <p>
                       <label>用户名:</label>
                       <input name='username' type='text' placeholder='用户名'>
                   </p>
                   <p id="username-checker" class="checker valid"></p>
                   <p>
                       <label>学号:</label>
                       <input name='number' type='text' placeholder='学号'>
                   </p>
                   <p id="number-checker" class="checker valid"></p>
                   <p>
                       <label>密码:</label>
                       <input name='password' type='password' placeholder='密码'>
                   </p>
                   <p>
                       <label>确认:</label>
                       <input name='password-sec' type='password' placeholder='再次输入密码'>
                   </p>
                   <p id="password-checker" class="checker valid"></p>
                   <p>
                       <label>电话:</label>
                       <input name='phone' type='text' placeholder='电话'>
                   </p>
                   <p id="phone-checker" class="checker valid"></p>
                   <p>
                       <label>邮箱:</label>
                       <input name='mail' type='text' placeholder='邮箱'>
                   </p>
                   <p id="email-checker" class="checker valid"></p>
               </form>
               <input id="submit" class="button" type='submit' value='注册'>
               <button id="reset" class="button">重置</button>
               <p id="msg" class="invalid"></p>
           </div>
       </body>
   </html>
   ```

3. 提交注册用户

   ```shell
   $ curl -v -X POST --data-urlencode "register=true" --data-urlencode "username=qwert" --data-urlencode "password=123123" --data-urlencode "phone=12312332112" --data-urlencode "mail=12341213@asa.com" --data-urlencode "regist=true" --data-urlencode "number=16340000" 127.0.0.1:8080/regist
   
   Note: Unnecessary use of -X or --request, POST is already inferred.
   *   Trying 127.0.0.1...
   * TCP_NODELAY set
   * Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
   > POST /regist HTTP/1.1
   > Host: 127.0.0.1:8080
   > User-Agent: curl/7.54.0
   > Accept: */*
   > Content-Length: 114
   > Content-Type: application/x-www-form-urlencoded
   >
   * upload completely sent off: 114 out of 114 bytes
   < HTTP/1.1 200 OK
   < Date: Mon, 12 Nov 2018 03:01:47 GMT
   < Content-Length: 0
   <
   * Connection #0 to host 127.0.0.1 left intact
   ```

   返回状态码为200，注册成功

   4. 检查用户名重复

      ```shell
      $ curl -v -X POST --data-urlencode "register=false" --data-urlencode "username=aaaaaa" 127.0.0.1:8080/regist
      
      Note: Unnecessary use of -X or --request, POST is already inferred.
      *   Trying 127.0.0.1...
      * TCP_NODELAY set
      * Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
      > POST /regist HTTP/1.1
      > Host: 127.0.0.1:8080
      > User-Agent: curl/7.54.0
      > Accept: */*
      > Content-Length: 30
      > Content-Type: application/x-www-form-urlencoded
      >
      * upload completely sent off: 30 out of 30 bytes
      < HTTP/1.1 200 OK
      < Date: Mon, 12 Nov 2018 03:03:23 GMT
      < Content-Length: 0
      <
      * Connection #0 to host 127.0.0.1 left intact
      ```

      无重复，返回200

      ```shell
      $ curl -v -X POST --data-urlencode "register=false" --data-urlencode "username=qwert" 127.0.0.1:8080/regist
      Note: Unnecessary use of -X or --request, POST is already inferred.
      *   Trying 127.0.0.1...
      * TCP_NODELAY set
      * Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
      > POST /regist HTTP/1.1
      > Host: 127.0.0.1:8080
      > User-Agent: curl/7.54.0
      > Accept: */*
      > Content-Length: 29
      > Content-Type: application/x-www-form-urlencoded
      >
      * upload completely sent off: 29 out of 29 bytes
      < HTTP/1.1 400 Bad Request
      < Date: Mon, 12 Nov 2018 03:04:06 GMT
      < Content-Length: 0
      <
      * Connection #0 to host 127.0.0.1 left intact
      ```

      存在重复，返回400

   5. 用户登录测试

      ```shell
      $ curl -v -X POST --data-urlencode "password=123123" --data-urlencode "username=qwert" 127.0.0.1:8080/signin
      Note: Unnecessary use of -X or --request, POST is already inferred.
      *   Trying 127.0.0.1...
      * TCP_NODELAY set
      * Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
      > POST /signin HTTP/1.1
      > Host: 127.0.0.1:8080
      > User-Agent: curl/7.54.0
      > Accept: */*
      > Content-Length: 30
      > Content-Type: application/x-www-form-urlencoded
      >
      * upload completely sent off: 30 out of 30 bytes
      < HTTP/1.1 200 OK
      < Date: Mon, 12 Nov 2018 03:05:21 GMT
      < Content-Length: 0
      <
      * Connection #0 to host 127.0.0.1 left intact
      ```

      登录信息正确，返回200

      ```shell
      $ curl -v -X POST --data-urlencode "password=123123" --data-urlencode "username=qwerty" 127.0.0.1:8080/signin
      Note: Unnecessary use of -X or --request, POST is already inferred.
      *   Trying 127.0.0.1...
      * TCP_NODELAY set
      * Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
      > POST /signin HTTP/1.1
      > Host: 127.0.0.1:8080
      > User-Agent: curl/7.54.0
      > Accept: */*
      > Content-Length: 31
      > Content-Type: application/x-www-form-urlencoded
      >
      * upload completely sent off: 31 out of 31 bytes
      < HTTP/1.1 400 Bad Request
      < Date: Mon, 12 Nov 2018 03:06:03 GMT
      < Content-Length: 1
      < Content-Type: text/plain; charset=utf-8
      <
      * Connection #0 to host 127.0.0.1 left intact
      2
      ```

      登录信息错误，返回400，内容为2表示用户名不存在

      ```shell
      $ curl -v -X POST --data-urlencode "password=123122" --data-urlencode "username=qwert" 127.0.0.1:8080/signin
      Note: Unnecessary use of -X or --request, POST is already inferred.
      *   Trying 127.0.0.1...
      * TCP_NODELAY set
      * Connected to 127.0.0.1 (127.0.0.1) port 8080 (#0)
      > POST /signin HTTP/1.1
      > Host: 127.0.0.1:8080
      > User-Agent: curl/7.54.0
      > Accept: */*
      > Content-Length: 30
      > Content-Type: application/x-www-form-urlencoded
      >
      * upload completely sent off: 30 out of 30 bytes
      < HTTP/1.1 400 Bad Request
      < Date: Mon, 12 Nov 2018 03:07:16 GMT
      < Content-Length: 1
      < Content-Type: text/plain; charset=utf-8
      <
      * Connection #0 to host 127.0.0.1 left intact
      1
      ```

      登录信息错误，返回400 ，内容为1表示密码错误

      6. 进入用户信息页面

         ```shell
         $ curl "127.0.0.1:8080/?username=qwert"
         
         <!DOCTYPE html>
         <head>
             <meta charset='UTF-8'>
             <title>User Info</title>
             <link rel='stylesheet' type='text/css' href='style.css'>
             <script type='text/javascript' src='https://cdn.jsdelivr.net/npm/jquery@3.3.1/dist/jquery.min.js'>
             </script>
             <script type='text/javascript' src='signup.js'>
             </script>
         </head>
         <body>
             <h1>详情</h1>
             <div id="main-area">
                 <h2>用户详情</h2>
                 <p>用户名: qwert</p>
                 <p>学号:  16340000</p>
                 <p>电话:  12312332112</p>
                 <p>邮箱:  12341213@asa.com</p>
                 <button id="go-back" class="button">退出</button>
                 <p id="msg" class="valid"></p>
             </div>
         </body>
         ```

         用户存在，返回用户详情页的 html

         ```shell
         $ curl "127.0.0.1:8080/?username=qwerty"
         
         <!DOCTYPE html>
         <html lang='en'>
             <head>
                 <meta charset='UTF-8'>
                 <title>User Sign Up</title>
                 <link rel='stylesheet' type='text/css' href='style.css'>
                 <script type='text/javascript' src='https://cdn.jsdelivr.net/npm/jquery@3.3.1/dist/jquery.min.js'>
                 </script>
                 <script type='text/javascript' src='signin.js'>
                 </script>
             </head>
             <body>
                 <h1>登陆</h1>
                 <div id="main-area">
                     <h2>用户登陆</h2>
                     <button id="signup" class="button" href='/regist'>注册</button>
                     <form method='post'>
                         <p>
                             <label>用户名:</label>
                             <input name='username' type='text' placeholder='用户名'>
                         </p>
                         <p id="username-checker" class="checker"></p>
                         <p>
                             <label>密码:</label>
                             <input name='password' type='password' placeholder='密码'>
                         </p>
                         <p id="password-checker" class="checker"></p>
                     </form>
                     <button id="submit" class="button">登陆</button>
                     <button id="reset" class="button">重置</button>
                     <p id="msg" class="invalid">Queried user does not exist!</p>
                 </div>
             </body>
         </html>
         ```

         查找的用户不存在，返回登录页面和用户不存在的信息。

         ## 使用 ab 工具进行测试

         ab工具全称 Apache Benchmark，用于HTTP 性能测试。

         对于本次的应用，使用如下命令进行测试：

         ```shell
         ab -n 10000 -c 2000 -r http://127.0.0.1:8080/
         ```

         其中`-n`参数表示测试的总请求数，`-c` 表示并发数，`-r` 表示当 socket 出现错误时程序不直接退出。

         执行测试的结果如下：

         ```shell
         $ ab -n 10000 -c 2000 -r http://127.0.0.1:8080/
         This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
         Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
         Licensed to The Apache Software Foundation, http://www.apache.org/
         
         Benchmarking 127.0.0.1 (be patient)
         Completed 1000 requests
         Completed 2000 requests
         Completed 3000 requests
         Completed 4000 requests
         Completed 5000 requests
         Completed 6000 requests
         Completed 7000 requests
         Completed 8000 requests
         Completed 9000 requests
         Completed 10000 requests
         Finished 10000 requests
         
         
         Server Software:
         Server Hostname:        127.0.0.1
         Server Port:            8080
         
         Document Path:          /
         Document Length:        1277 bytes
         
         Concurrency Level:      2000
         Time taken for tests:   4.373 seconds
         Complete requests:      10000
         Failed requests:        0
         Total transferred:      13950000 bytes
         HTML transferred:       12770000 bytes
         Requests per second:    2286.95 [#/sec] (mean)
         Time per request:       874.525 [ms] (mean)
         Time per request:       0.437 [ms] (mean, across all concurrent requests)
         Transfer rate:          3115.53 [Kbytes/sec] received
         
         Connection Times (ms)
                       min  mean[+/-sd] median   max
         Connect:        0  131 262.9     39    1176
         Processing:     2  566 213.3    609    1588
         Waiting:        2  536 208.7    591    1587
         Total:          2  696 323.9    671    1777
         
         Percentage of the requests served within a certain time (ms)
           50%    671
           66%    721
           75%    763
           80%    807
           90%    995
           95%   1619
           98%   1646
           99%   1676
          100%   1777 (longest request)
         ```

         结果中各个信息的意义如下：

         ```
         #请求的文档路径
         Document Path:          /
         #请求的文档内容长度
         Document Length:        1277 bytes
         ```

         ```
         #完成的请求数
         Complete requests:      10000
         #失败的请求数
         Failed requests:        0
         #传输的总数据量
         Total transferred:      13950000 bytes
         #传输的 HTML 内容数据量
         HTML transferred:       12770000 bytes
         #每秒完成的请求数量
         Requests per second:    2286.95 [#/sec] (mean)
         #平均的请求等待时间
         Time per request:       874.525 [ms] (mean)
         #服务器对每个请求的平均处理时间
         Time per request:       0.437 [ms] (mean, across all concurrent requests)
         #平均网络流量
         Transfer rate:          3115.53 [Kbytes/sec] received
         ```

         ```
         #测试中网络在各个阶段消耗的时间分布
         Connection Times (ms)
                       min  mean[+/-sd] median   max
         Connect:        0  131 262.9     39    1176
         Processing:     2  566 213.3    609    1588
         Waiting:        2  536 208.7    591    1587
         Total:          2  696 323.9    671    1777
         ```

         ```
         #请求处理时间的分布，如下表90%的请求在 995ms 内处理完成
         Percentage of the requests served within a certain time (ms)
           50%    671
           66%    721
           75%    763
           80%    807
           90%    995
           95%   1619
           98%   1646
           99%   1676
          100%   1777 (longest request)
         ```

         **使用 POST 方法运行 ab 进行测试：**

         在 postdata.txt文件中输入如下内容：

         ```
         username=qwerty&password=123123
         ```

         执行如下命令进行测试：

         ```shell
         ab -n 10000 -c 2000 -r  -p postdata.txt -T "application/x-www-form-urlencoded" http://127.0.0.1:8080/signin
         ```

         其中`-p` 表示POST 传输的数据，-T 表示content-type头内容。

         执行结果如下：

         ```shell
         $ ab -n 10000 -c 2000 -r  -p postdata.txt -T "application/x-www-form-urlencoded" http://127.0.0.1:8080/signin
         This is ApacheBench, Version 2.3 <$Revision: 1807734 $>
         Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
         Licensed to The Apache Software Foundation, http://www.apache.org/
         
         Benchmarking 127.0.0.1 (be patient)
         Completed 1000 requests
         Completed 2000 requests
         Completed 3000 requests
         Completed 4000 requests
         Completed 5000 requests
         Completed 6000 requests
         Completed 7000 requests
         Completed 8000 requests
         Completed 9000 requests
         Completed 10000 requests
         Finished 10000 requests
         
         
         Server Software:
         Server Hostname:        127.0.0.1
         Server Port:            8080
         
         Document Path:          /signin
         Document Length:        1 bytes
         
         Concurrency Level:      2000
         Time taken for tests:   6.746 seconds
         Complete requests:      10000
         Failed requests:        0
         Non-2xx responses:      10000
         Total transferred:      1260000 bytes
         Total body sent:        1900000
         HTML transferred:       10000 bytes
         Requests per second:    1482.29 [#/sec] (mean)
         Time per request:       1349.266 [ms] (mean)
         Time per request:       0.675 [ms] (mean, across all concurrent requests)
         Transfer rate:          182.39 [Kbytes/sec] received
                                 275.03 kb/s sent
                                 457.42 kb/s total
         
         Connection Times (ms)
                       min  mean[+/-sd] median   max
         Connect:        0  186 329.0     86    3080
         Processing:     4  997 435.4   1012    2805
         Waiting:        4  969 434.8   1003    2802
         Total:          4 1183 515.9   1106    3938
         
         Percentage of the requests served within a certain time (ms)
           50%   1106
           66%   1323
           75%   1409
           80%   1535
           90%   1999
           95%   2177
           98%   2434
           99%   2472
          100%   3938 (longest request)
         ```

         ## 使用 Docker并进行部署

         在 Docker Hub 中连接 GitHub 账号并对该 Repo 启动 automated build即可令 Docker Hub 自动构建应用的 Docker 映像。编写如下的 Dockerfile 对构建进行配置：

         ```dockerfile
         FROM golang:1.11
         
         WORKDIR /go/src/github.com/miguch/cloudgo
         
         COPY . .
         
         RUN go get -d -v .
         RUN go build -v .
         ```

         在docker hub 的build details 中显示完成构建后，在装有 docker 服务的服务器上执行如下命令完成镜像的部署：

         ```shell
         #从 Docker Hub 获取镜像
         sudo docker pull miguch/cloudgo
         #创建容器运行镜像，-d 表示容器在后台运行，--name 表示容器名称，-p 表示将容器的8080端口转发到宿主机的8081端口，使容器可以从外部访问
         sudo docker run -d -p 0.0.0.0:8081:8080 --name cloudgo miguch/cloudgo ./cloudgo -a 0.0.0.0
         ```

         执行后从浏览器访问`<服务器 ip>:8081` 即可访问应用。