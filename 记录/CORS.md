
##服务器
####响应头
* Access-Control-Allow-Origin `字段必须（支持跨域的话），*，表示接受任意域名的请求。设置为*不能使用Cookie`
* Access-Control-Allow-Credentials `是否允许发送Cookie。CORS默认不发送。这个值只能设置为true。前端需要设置withCredentials属性为true`
* Access-Control-Expose-Headers `前端只能拿到六个基本字段Cache-Control、Content-Language、Content-Type、Expires、Last-Modified、Pragma。如果需要其他字段就需要在这个字段中指定`

> 开启Cookie的情况下,Cookie依然遵循同源政策，只有用服务器域名设置的Cookie才会上传，其他域名的Cookie并不会上传，且（跨源）原网页代码中的document.cookie也无法读取服务器域名下的Cookie

##前端

###两种请求
####1、简单请求
> 同时满足一下条件
```
1、请求方式是一下三种之一
HEAD
GET
POST
2、HTTP头信息不超出一下几个字段
Accept
Accept-Language
Content-Language
Last-Event-ID
Content-Type  //只限于三个值application/x-www-form-urlencoded、multipart/form-data、text/plain
```
#####请求流程
1. 发出CORS请求`在头信息中增加一个origin字段`
2. 服务端判断`origin`值否在许可范围，如果允许跨域，则设置响应头`Access-Control-Allow-Origin`
3. 浏览器判断响应头是否包含字段，如果没有包含就标识不支持跨域请求`不能通过HTTP状态码判断`

####2、非简单请求
在正式通信前，增加一次HTTP查询请求`预检请求preflight`
非简单请求是那种对服务器有特殊要求的请求，比如请求方法是PUT或DELETE，或者Content-Type字段的类型是application/json
得到服务器肯定的大幅才会发起正式请求
"预检"请求用的请求方法是OPTIONS

（1）Access-Control-Request-Method

该字段是必须的，用来列出浏览器的CORS请求会用到哪些HTTP方法，上例是PUT。

（2）Access-Control-Request-Headers

该字段是一个逗号分隔的字符串，指定浏览器CORS请求会额外发送的头信息字段

######预检请求响应
服务器没有返回CORS相关的头信息字段，浏览器则认为服务器不同意预检请求，触发错误

服务器通过了预检请求，以后每次CORS请求，就跟简单请求一样，会带有Origin


（1）Access-Control-Allow-Methods

该字段必需，它的值是逗号分隔的一个字符串，表明服务器支持的所有跨域请求的方法。注意，返回的是所有支持的方法，而不单是浏览器请求的那个方法。这是为了避免多次"预检"请求。

（2）Access-Control-Allow-Headers

如果浏览器请求包括Access-Control-Request-Headers字段，则Access-Control-Allow-Headers字段是必需的。它也是一个逗号分隔的字符串，表明服务器支持的所有头信息字段，不限于浏览器在"预检"中请求的字段。

（3）Access-Control-Allow-Credentials

该字段与简单请求时的含义相同。

（4）Access-Control-Max-Age

该字段可选，用来指定本次预检请求的有效期，单位为秒。上面结果中，有效期是20天（1728000秒），即允许缓存该条回应1728000秒（即20天），在此期间，不用发出另一条预检请求。



Origin字段是浏览器自动添加的