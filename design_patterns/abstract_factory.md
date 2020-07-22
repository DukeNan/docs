## 抽象工厂模式

### 模式动机

**抽象工厂模式**是一种创建型设计模式， 它能创建一系列相关的对象， 而无需指定其具体类。

![abstract-factory01](../_media/images/design_patterns/abstract-factory01.png)

### 模式结构

抽象工厂模式包含如下角色：

- AbstractFactory：抽象工厂
- ConcreteFactory：具体工厂
- AbstractProduct：抽象产品
- Product：具体产品

![abatract-factory02](../_media/images/design_patterns/abatract-factory02.jpg)

### 实现方法

1. 以不同的产品类型与产品变体为维度绘制矩阵。
2. 为所有产品声明抽象产品接口。 然后让所有具体产品类实现这些接口。
3. 声明抽象工厂接口， 并且在接口中为所有抽象产品提供一组构建方法。
4. 为每种产品变体实现一个具体工厂类。
5. 在应用程序中开发初始化代码。 该代码根据应用程序配置或当前环境， 对特定具体工厂类进行初始化。 然后将该工厂对象传递给所有需要创建产品的类。
6. 找出代码中所有对产品构造函数的直接调用， 将其替换为对工厂对象中相应构建方法的调用。

### 实例

```python
from abc import ABCMeta, abstractmethod


class PizzaFactory(metaclass=ABCMeta):

    @abstractmethod
    def createVegPizza(self):
        """创建蔬菜披萨"""
        pass

    @abstractmethod
    def createNonVegPizza(self):
        """创建无蔬菜披萨"""
        pass


class IndianPizzaFactory(PizzaFactory):
    """印度风味披萨工厂"""

    def createVegPizza(self):
        return DeluxeVeggiePizza()

    def createNonVegPizza(self):
        return ChickenPizza()


class USPizzaFactory(PizzaFactory):
    """美式披萨工厂"""

    def createVegPizza(self):
        return MexicanVegPizza()

    def createNonVegPizza(self):
        return HamPizza()


class VegPizza(metaclass=ABCMeta):
    """抽象产品：素食披萨"""

    @abstractmethod
    def prepare(self, VegPizza):
        """开始准备"""
        pass


class NonVegPizza(metaclass=ABCMeta):
    """抽象产品：非素食披萨"""

    @abstractmethod
    def serve(self, VegPizza):
        pass


class DeluxeVeggiePizza(VegPizza):
    """具体产品：豪华素食披萨"""

    def prepare(self):
        print('Prepare', type(self).__name__)


class ChickenPizza(NonVegPizza):
    """具体产品： 鸡肉披萨"""

    def serve(self, VegPizza):
        print(type(self).__name__, ' is served with Chicken on ', type(VegPizza).__name__)


class MexicanVegPizza(VegPizza):
    """具体产品：墨西哥素食披萨"""

    def prepare(self):
        print('Prepare', type(self).__name__)


class HamPizza(NonVegPizza):
    """具体产品：火腿披萨"""

    def serve(self, VegPizza):
        print(type(self).__name__, ' is served with Ham on ', type(VegPizza).__name__)


# 当用户来到PizzaStore并要一份美式非素食披萨的时候，USPizzaFactory负责准备素食，然后在上面加上火腿，马上就变成非素食披萨了。

class PizzaStore:
    def __init__(self):
        pass

    def makePizzas(self):
        for factory in [IndianPizzaFactory(), USPizzaFactory()]:
            self.factory = factory
            self.NonVegPizza = self.factory.createNonVegPizza()
            self.VegPizza = self.factory.createVegPizza()
            self.VegPizza.prepare()
            self.NonVegPizza.serve(self.VegPizza)


if __name__ == '__main__':
    pizza = PizzaStore()
    pizza.makePizzas()
    
[out]: Prepare DeluxeVeggiePizza
			 ChickenPizza  is served with Chicken on  DeluxeVeggiePizza
       Prepare MexicanVegPizza
       HamPizza  is served with Ham on  MexicanVegPizza
 

```

### 优缺点

:red_circle:  你可以确保同一工厂生成的产品相互匹配。

:red_circle:  你可以避免客户端和具体产品代码的耦合。

:red_circle:  **单一职责原则**。 你可以将产品生成代码抽取到同一位置， 使得代码易于维护。

:red_circle:  **开闭原则**。 向应用程序中引入新产品变体时， 你无需修改客户端代码。

:black_circle:  由于采用该模式需要向应用中引入众多接口和类， 代码可能会比之前更加复杂。


