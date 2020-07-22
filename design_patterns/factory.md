## 工厂方法

### 定义

**工厂方法模式**是一种创建型设计模式， 其在父类中提供一个创建对象的方法， 允许子类决定实例化对象的类型。

在简单工厂模式中，可以根据传递的参数不同，返回不同类的实例。简单工厂模式定义了一个类，这个类专门用于创建其他类的实例，这些被创建的类都有一个共同的父类。

### 应用场景

工厂类负责创建的对象比较少。 简单工厂模式最大的优点在于实现对象的创建和对象的使用分离，但是如果产品过多时，会导致工厂代码非常复杂。简单工厂模式的要点就在于当你需要什么，只需要传入一个正确的参数，就可以获取你所需要的对象，而无须知道其创建细节。

### 模式结构

简单工厂模式包含如下角色：

- **Factory：工厂角色**

  工厂角色负责实现创建所有实例的内部逻辑

- **Product：抽象产品角色**

  抽象产品角色是所创建的所有对象的父类，负责描述所有实例所共有的公共接口

- **ConcreteProduct：具体产品角色**

  具体产品角色是创建目标，所有创建的对象都充当这个角色的某个具体类的实例。

![SimpleFactory](../_media/images/design_patterns/SimpleFactory.jpg)

### 实现方式

1. 让所有产品都遵循同一接口。 该接口必须声明对所有产品都有意义的方法。
2. 在创建类中添加一个空的工厂方法。 该方法的返回类型必须遵循通用的产品接口。
3. 在创建者代码中找到对于产品构造函数的所有引用。 将它们依次替换为对于工厂方法的调用， 同时将创建产品的代码移入工厂方法。 你可能需要在工厂方法中添加临时参数来控制返回的产品类型。
4. 为工厂方法中的每种产品编写一个创建者子类， 然后在子类中重写工厂方法， 并将基本方法中的相关创建代码移动到工厂方法中。
5. 如果应用中的产品类型太多， 那么为每个产品创建子类并无太大必要， 这时你也可以在子类中复用基类中的控制参数。
6. 如果代码经过上述移动后， 基础工厂方法中已经没有任何代码， 你可以将其转变为抽象类。 如果基础工厂方法中还有其他语句， 你可以将其设置为该方法的默认行为。



### 代码示例

```python
from abc import ABCMeta, abstractmethod


class Payment(metaclass=ABCMeta):

    @abstractmethod
    def pay(self, money):
        pass


class AliPay(Payment):

    def __init__(self, yu_e_bao=False):
        self.yu_e_bao = yu_e_bao

    def pay(self, money):

        if self.yu_e_bao:
            print(f'use yu_e_bao pay {money}')

        else:
            print(f'use zhifubao pay {money}')


class WeChat(Payment):

    def pay(self, money):
        print(f'use WeChat pay {money}')


class PayMethod():

    def create_payment(self, method):

        if method == 'yu_e_bao':
            return AliPay(True)

        elif method == 'zhifubao':
            return AliPay(False)

        elif method == 'Wechat':
            return WeChat()
        else:
            raise ValueError('method error')


if __name__ == '__main__':
    p = PayMethod()
    f = p.create_payment('zhifubao')
    f.pay(30)
    f = p.create_payment('Wechat')
    f.pay(300)

```

此模式还有三个角色：

* 抽象产品角色：Payment类，定义产品的必要功能。
* 具体产品角色：Alipay，Wechat类，具体实例化出来的对象。
* 工厂角色：Paymethod类，根据参数输出具体产品。

**注意：**

- 我们定义一个接口创建对象，但是工厂本身并不负责创建对象，而是将这一任务交由子类来完成，即子类决定了要实例化哪些类。
- Factory方法的创建是通过继承而不是通过实例化来完成的。
- 工厂方法使设计更加具有可定制性。它可以返回相同的实例或子类要，而不是某种类型的对象（就像在简单工厂方法那样）。

### 优缺点

优点：

* 你可以避免创建者和具体产品之间的紧密耦合。

* 单一职责原则。 你可以将产品创建代码放在程序的单一位置， 从而使得代码更容易维护。

* 开闭原则。 无需更改现有客户端代码， 你就可以在程序中引入新的产品类型。

缺点：

* 应用工厂方法模式需要引入许多新的子类， 代码可能会因此变得更复杂。 最好的情况是将该模式引入创建者类的现有层次结构中。

