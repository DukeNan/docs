

## 闭包

在函数内部再定义一个函数，并且这个函数用到了外边函数的变量，那么称里面的这个函数为闭包

内部函数对外部函数作用域里变量的引用（非全局变量），则称内部函数为闭包。

```python
def counter(start=0):
    count=[start]
    def incr():
        count[0] += 1
        return count[0]
    return incr
```



**总结：**

* 闭包似优化了变量，原来需要类对象完成的工作，闭包也可以完成
* 由于闭包引用了外部函数的局部变量，则外部函数的局部变量没有及时释放，消耗内存



##  装饰器

装饰器(Decorators)是Python的一个重要部分。简单地说：他们是修改其他函数的功能的函数。他们有助于让我们的代码更简短，也更Pythonic（Python范儿）。



### 使用场景

* 权限校验等场景
* 引入日志
* 函数执行时间统计
* 执行函数前预备处理
* 执行函数后清理功能
* 缓存
* 事务处理

……

### 实例

1.无参数函数:

```python
def outer(fn):
    def inner():
        fn()

    return inner
  
# from functools import wraps
# def outer(fn):
#     @wraps(fn)
#     def inner():
#         return fn()
# 
#     return inner

@outer
def fun():
    print('12345678')


if __name__ == '__main__':
    fun()  # 相当于：fun() == outer(fun)(),  输出：12345678
```

2.被装饰的函数有参数:

```python
def outer(fn):
    def inner(a, b):
        fn(a, b)

    return inner


@outer
def fun(a, b):
    print(a + b)


if __name__ == '__main__':
    fun(2, 3)  # 相当于：fun(2,3) == outer(fun)(2,3), # 输出5

```

3.被装饰的函数有不定长参数:

```python
def outer(fn):
    def inner(*args, **kwargs):
        fn(*args, **kwargs)

    return inner


@outer
def fun(a, b, c):
    print(a + b + c)


if __name__ == '__main__':
    fun(1, b=2, c=3)  # 输出

```

4.装饰器中的return:

```python
def outer(fn):
    def inner():
        print('我是%s'% fun.__name__)
        fn()
    return inner


@outer
def fun():
    print('hahahhah')
    return '王者'

if __name__ == '__main__':
    fun()  # 相当于：fun() == outer(fun)()
	>>> 我是inner
	>>> hahahhah

    print(fun())
	>>> 我是inner
	>>> hahahhah
	>>> None  # 添加return返回 ‘王者’
```

**注意：** 一般情况下为了让装饰器更通用，可以有`return`, 即：`fn()`改为` return fun()`

5：装饰器带参数,在原有装饰器的基础上，设置外部变量：

```python
def timefun_arg(pre="hello"):
    def timefun(func):
        def wrappedfunc():
            print("%s called at %s" % (func.__name__, pre))
            return func()

        return wrappedfunc

    return timefun


@timefun_arg("linux")
def foo():
    print("I am foo")


@timefun_arg("python")
def too():
    print("I am too")


if __name__ == '__main__':
    foo()  # foo()相当于：timefun_arg("itcast")(foo)()
```



### `functools.wraps`

使用装饰器极大地复用了代码，但是他有一个缺点就是原函数的元信息不见了，比如函数的docstring、__name__、参数列表:

```python
def use_logging(func):
    def _deco(*args, **kwargs):
        print("%s is running" % func.__name__)
        func(*args, **kwargs)

    return _deco


@use_logging
def bar():
    print('i am bar')
    print(bar.__name__)


if __name__ == '__main__':
    bar()
# bar is running
# i am bar
# _deco
# 函数名变为_deco而不是bar，这个情况在使用反射的特性的时候就会造成问题。因此引入了functools.wraps解决这个问题。

```

使用`functools.wraps`:

```python
from functools import wraps


def use_logging(func):
    @wraps(func)
    def _deco(*args, **kwargs):
        print("%s is running" % func.__name__)
        func(*args, **kwargs)

    return _deco


@use_logging
def bar():
    print('i am bar')
    print(bar.__name__)


bar()
# result:
# bar is running
# i am bar
# bar ,这个结果是我们想要的。OK啦

```

实现带参数和不带参数的装饰器自适应:

```python
import functools


def use_logging(arg):
    if callable(arg):  # 判断参入的参数是否是函数，不带参数的装饰器调用这个分支
        @functools.wraps(arg)
        def _deco(*args, **kwargs):
            print("%s is running" % arg.__name__)
            arg(*args, **kwargs)

        return _deco
    else:  # 带参数的装饰器调用这个分支
        def _deco(func):
            @functools.wraps(func)
            def __deco(*args, **kwargs):
                if arg == "warn":
                    print("warn: %s is running" % func.__name__)
                return func(*args, **kwargs)

            return __deco

        return _deco


@use_logging("warn")
# @use_logging
def bar():
    print('i am bar')
    print(bar.__name__)


bar()

```



### 多层装饰器

装饰顺序按靠近函数顺序执行，调用时由外而内，执行顺序和装饰顺序相反。



#### LEGB规则

locals -> enclosing function -> globals -> builtins

locals，当前所在命名空间（如函数、模块），函数的参数也属于命名空间内的变量

enclosing，外部嵌套函数的命名空间（闭包中常见）

globals，全局变量，函数定义所在模块的命名空间

builtins，内建模块的命名空间。