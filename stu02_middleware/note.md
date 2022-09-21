# 中间件学习
## 基本概念
想要给路由增加一个中间件，需要2步：
1. 定义一个方法，接收一个参数是HandlerFunc，并返回一个HandlerFunc, 如下
func Middleware01(h echo.HandlerFunc) echo.HandlerFunc
(HandlerFunc是处理一个请求并设置响应的方法，定义如下：
type HandlerFunc func(c echo.Context) error
)
2. 在设置e.GET(path, handler, middlewares...) 时，将middleware注入进去。
3. 效果就是在没有改动handler代码的前提下，却额外执行了一些前置或者后置代码。

问题
1. 为什么要把middleware定义成这样的形式？func(h HandlerFunc) HandlerFunc?
2. middleware是怎么起作用的？
3. 多个middleware是怎么起作用的？

## middleware包裹比喻
applyMiddleware(h, middleware...)

func m1(h handler) handler {
  return func(c context) error{
    a();
    h();
    b();
  }
}

h = m1(h),
h原本是 func(c context) error { xxx; }
m1(h) 是 func(c context) error {
    a();
    h();
    b();
}

h原本是一个handler，
m1(h) 相当于包裹了h，加了一层外衣（一些额外的动作），依然还是一个handler。
m1中间件本身不是handler，而是给传入handler增加“外衣”的工具函数或完成加工过程的办法。
所以经过一系列中间件加持，最后传给路由的还是一个handler。

用加外衣比喻比较形象。
如果有3个m中间件，一个实际路由函数h。执行结果会是怎么样的？
m1 返回 func(c context) error {
    m1a();
    h();
    m1b();
}

m2 返回 func(c context) error {
    m2a();
    h();
    m2b();
}

m3 返回 func(c context) error {
    m3a();
    h();
    m3b();
}

内衣 m3a(); h(); m3b();
外衣 m2a(); (内衣) m2b();
外套 m1a(); (外衣) m1b();

try02m1_s
try02m2_s
try02m3_s
try02Check
try02m3_e
try02m2_e
try02m1_e

## group上的middleware
可以在group上注册middleware，
这样所有group下的路由都可以执行这些公共中间件。
r := e.Group("/aa", middleware...)
或者 r.Use(middleware...) 都会在group一级临时保存这些中间件。
1. 执行顺序
Add方法里，
middleware = append(middleware, group.middleware...)
middleware = append(middleware, middleware...)
所以group中间件在路由上中间件的外层。
2. 生效顺序
如果 
r.Add("/a", h)
r.Use(middleware...)
r.Add("/b", h)
/a是不会得到中间件的，因为中间件的作用在Add时就执行完了，完成了对h的包裹。
而不是在路由请求过来时才会执行中间件函数。
而此时取到的中间件还没有Use声明的那些。