## 单例模式



### 意图

**单例模式**是一种创建型设计模式， 让你能够保证一个类只有一个实例， 并提供一个访问该实例的全局节点。

![singleton01](../_media/images/design_patterns/singleton01.png)

### 模式结构

单例模式的主要角色如下。

- 单例类：包含一个实例且能自行创建这个实例的类。
- 访问类：使用单例的类。

![singleton02](../_media/images/design_patterns/singleton02.png)

### 实现方式

1. 在类中添加一个私有静态成员变量用于保存单例实例。
2. 声明一个公有静态构建方法用于获取单例实例。
3. 在静态方法中实现"延迟初始化"。 该方法会在首次被调用时创建一个新对象， 并将其存储在静态成员变量中。 此后该方法每次被调用时都返回该实例。
4. 将类的构造函数设为私有。 类的静态方法仍能调用构造函数， 但是其他对象不能调用。
5. 检查客户端代码， 将对单例的构造函数的调用替换为对其静态构建方法的调用。

### 实例

#### 使用模块生成单例

**Python 的模块就是天然的单例模式**，因为模块在第一次导入时，会生成 `.pyc` 文件，当第二次导入时，就会直接加载 `.pyc` 文件，而不会再次执行模块代码。因此，我们只需把相关的函数和数据定义在一个模块中，就可以获得一个单例对象了。如果我们真的想要一个单例类，可以考虑这样做：

```python
class Singleton(object):
    def foo(self):
        pass
singleton = Singleton()
```

将上面的代码保存在文件 `mysingleton.py` 中，要使用时，直接在其他文件中导入此文件中的对象，这个对象即是单例模式的对象

```python
from a import singleton
```

#### 使用装饰器

```python
def singleton(cls, *args, **kwargs):
    """类装饰器"""
    instances = {}

    def _singleton():
        if cls not in instances:
            instances[cls] = cls(*args, **kwargs)
        return instances[cls]

    return _singleton


@singleton
class A:
    a = 1

    def __init__(self, x=0):
        self.x = x


if __name__ == '__main__':
    one = A()
    two = A()

    two.a = 3
    print(one.a)
    print(id(one))
    print(id(two))
    print(one == two)
    print(two.x)
    one.x = 1
    print(one.x)
    print(one is two)

    
[out]:
3
4436525648
4436525648
True
0
1
True
```

#### 使用类

```python
class Singleton:

    def __init__(self, *args, **kwargs):
        pass

    @classmethod
    def instance(cls, *args, **kwargs):
        if not hasattr(Singleton, '_instance'):
            Singleton._instance = Singleton(*args, **kwargs)
        return Singleton._instance


if __name__ == '__main__':
    s = Singleton().instance()
    s1 = Singleton().instance()
    print(id(s))
    print(id(s1))
    
[out]:
4303832144
4303832144
```

**注意：** 此方法不支持多线程

支持多线程:

```python
from threading import Lock


class Singleton:
    _lock = Lock()

    def __init__(self, *args, **kwargs):
        pass

    @classmethod
    def instance(cls, *args, **kwargs):
        if not hasattr(Singleton, '_instance'):
            with Singleton._lock:
                if not hasattr(Singleton, '_instance'):
                    Singleton._instance = Singleton(*args, **kwargs)
        return Singleton._instance


if __name__ == '__main__':
    s = Singleton().instance()
    s1 = Singleton().instance()
    print(id(s))
    print(id(s1))

```



#### 基于`__new__`实现

```python
from threading import Lock


class Singleton:
    _lock = Lock()

    def __init__(self):
        pass

    def __new__(cls, *args, **kwargs):
        """
        为对象分配内存，返回对象的引用
        “new”决定是否要使用该类的”init”方法，
        因为”new” 可以调用其他类的构造方法或者直接返回别的类创建的对象来作为本类的实例。
        """
        if not hasattr(Singleton, '_instance'):
            with cls._lock:
                if not hasattr(Singleton, '_instance'):
                    cls._instance = super().__new__(cls)
        return cls._instance


if __name__ == '__main__':
    obj1 = Singleton()
    obj2 = Singleton()

    print(id(obj1))
    print(id(obj2))
```

### 优缺点

:red_circle:  可以保证一个类只有一个实例。

:red_circle:  获得了一个指向该实例的全局访问节点。

:red_circle:   仅在首次请求单例对象时对其进行初始化。

:black_circle:  违反了**单一职责原则**。 该模式同时解决了两个问题。

:black_circle:  单例模式可能掩盖不良设计， 比如程序各组件之间相互了解过多等。

:black_circle:  该模式在多线程环境下需要进行特殊处理， 避免多个线程多次创建单例对象。

:black_circle:  单例的客户端代码单元测试可能会比较困难， 因为许多测试框架以基于继承的方式创建模拟对象。 由于单例类的构造函数是私有的， 而且绝大部分语言无法重写静态方法， 所以你需要想出仔细考虑模拟单例的方法。 要么干脆不编写测试代码， 或者不使用单例模式。



   

