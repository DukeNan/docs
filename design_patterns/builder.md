##  生成器模式

### 意图

**生成器模式（又称<font color="red">建造者模式</font>）**是一种创建型设计模式， 使你能够分步骤创建复杂对象。 该模式允许你使用相同的创建代码生成不同类型和形式的对象。

![builder01](../_media/images/design_patterns/builder01.png)

### 模式结构

建造者模式包含如下角色：

- Builder：抽象建造者

  它是一个包含创建产品各个子部件的抽象方法的接口，通常还包含一个返回复杂产品的方法 getResult()。

- ConcreteBuilder：具体建造者

  实现 Builder 接口，完成复杂产品的各个部件的具体创建方法。具体生成器也可以构造不遵循通用接口的产品。

- Director：指挥者

  它调用建造者对象中的部件构造与装配方法完成复杂对象的创建，在指挥者中不涉及具体产品的信息。

- Product：产品角色

  它是包含多个组成部件的复杂对象，由具体建造者来创建其各个部件。

![builder02](../_media/images/design_patterns/builder02.gif)

### 实现方法

步骤：

1. 清晰地定义通用步骤， 确保它们可以制造所有形式的产品。 否则你将无法进一步实施该模式。
2. 在基本生成器接口中声明这些步骤。
3. 为每个形式的产品创建具体生成器类， 并实现其构造步骤。（不要忘记实现获取构造结果对象的方法。 你不能在生成器接口中声明该方法， 因为不同生成器构造的产品可能没有公共接口， 因此你就不知道该方法返回的对象类型。 但是， 如果所有产品都位于单一类层次中， 你就可以安全地在基本接口中添加获取生成对象的方法。）
4. 考虑创建主管类。 它可以使用同一生成器对象来封装多种构造产品的方式。
5. 客户端代码会同时创建生成器和主管对象。 构造开始前， 客户端必须将生成器对象传递给主管对象。 通常情况下， 客户端只需调用主管类构造函数一次即可。 主管类使用生成器对象完成后续所有制造任务。 还有另一种方式， 那就是客户端可以将生成器对象直接传递给主管类的制造方法。
6. 只有在所有产品都遵循相同接口的情况下， 构造结果可以直接通过主管类获取。 否则， 客户端应当通过生成器获取构造结果。

### 实例

```python
from abc import ABCMeta, abstractmethod


class Car:
    """
    一辆汽车可能配备有 GPS 设备、行车电脑和几个座位。
    不同型号的汽车（运动型轿车、SUV 和敞篷车）可能会安装或启用不同的功能。
    """

    def __init__(self):
        self.parts = []

    def add(self, part) -> None:
        """组装零件"""
        self.parts.append(part)

    def list_parts(self) -> None:
        print(f"Car parts: {', '.join(self.parts)}", end="")


class Manual:
    """
    用户使用手册应该根据汽车配置进行编制，并介绍汽车的所有功能。
    """

    def __init__(self):
        self.parts = []

    def add(self, part) -> None:
        """封装零件说明项"""
        self.parts.append(part)

    def list_parts(self) -> None:
        print(f"Manual parts: {', '.join(self.parts)}", end="")


class Builder(metaclass=ABCMeta):
    """
    生成器接口声明了创建产品对象不同部件的方法。
    """

    @property
    @abstractmethod
    def product(self) -> None:  # 只读属性
        pass

    @abstractmethod
    def reset(self) -> None:
        pass

    @abstractmethod
    def set_seats(self, *args, **kwargs) -> None:
        pass

    @abstractmethod
    def set_engine(self, *args, **kwargs) -> None:
        pass

    @abstractmethod
    def set_trip_computer(self) -> None:
        pass

    @abstractmethod
    def set_gps(self) -> None:
        pass


class CarBuilder(Builder):
    """
    具体生成器类将遵循生成器接口并提供生成步骤的具体实现。
    你的程序中可能会有多个以不同方式实现的生成器变体。
    """

    def __init__(self):
        """一个新的生成器实例必须包含一个在后续组装过程中使用的空产品对象"""
        self.reset()

    def reset(self) -> None:
        """ reset（重置）方法可清除正在生成的对象。"""
        self._car = Car()

    """所有生成步骤都会与同一个产品实例进行交互"""

    def set_seats(self, num: int) -> None:
        """设置座椅数量"""
        self._car.add(f"{num}张桌椅")

    def set_engine(self, engine_type: str) -> None:
        """安装指定引擎"""
        self._car.add(engine_type)

    def set_trip_computer(self) -> None:
        """安装行车电脑"""
        self._car.add('行车电脑')

    def set_gps(self) -> None:
        """安装全球定位系统"""
        self._car.add('GPS')

    @property
    def product(self) -> Car:
        """
        具体生成器需要自行提供获取结果的方法。这是因为不同类型的生成器可能
        会创建不遵循相同接口的、完全不同的产品。所以也就无法在生成器接口中
        声明这些方法（至少在静态类型的编程语言中是这样的）。

        通常在生成器实例将结果返回给客户端后，它们应该做好生成另一个产品的
        准备。因此生成器实例通常会在 `getProduct（获取产品）`方法主体末尾
        调用重置方法。但是该行为并不是必需的，你也可让生成器等待客户端明确
        调用重置方法后再去处理之前的结果。
        """
        car = self._car
        self.reset()
        return car


class CarManualBuilder(Builder):
    """
    生成器与其他创建型模式的不同之处在于：它让你能创建不遵循相同接口的产品。
    """

    def __init__(self):
        self.reset()

    def reset(self):
        self._car_manual = Manual()

    def set_seats(self):
        self._car_manual.add('座椅介绍')

    def set_engine(self):
        self._car_manual.add('引擎介绍')

    def set_trip_computer(self):
        self._car_manual.add('行车电脑介绍')

    def set_gps(self, *args, **kwargs):
        self._car_manual.add('GPS介绍')

    @property
    def product(self) -> Manual:
        """
        返回使用手册并重置生成器。
        """
        manual = self._car_manual
        self.reset()
        return manual


class Director:
    def __init__(self) -> None:
        self._builder = None

    def set_builder(self, builder: Builder) -> None:
        self.builder = builder

    def construct_sport_car(self, builder: Builder):
        builder.reset()
        builder.set_seats(num=2)
        builder.set_engine('跑车引擎')
        builder.set_trip_computer()
        builder.set_gps()

    def constructSUV(self, builder: Builder):
        pass

    def constructManual(self, builder: Builder):
        builder.reset()
        builder.set_seats()
        builder.set_engine()
        builder.set_trip_computer()
        builder.set_gps()


if __name__ == '__main__':
    director = Director()
    car_builder = CarBuilder()

    print('生产超跑......')
    director.construct_sport_car(car_builder)
    car = car_builder.product
    car.list_parts()

    print('\n')

    print('生产说明书......')
    manual_builder = CarManualBuilder()
    director.set_builder(manual_builder)
    director.constructManual(manual_builder)
    manual = manual_builder.product
    manual.list_parts()

```

### 模式扩展

建造者模式的简化:

- 省略抽象建造者角色：如果系统中只需要一个具体建造者的话，可以省略掉抽象建造者。
- 省略指挥者角色：在具体建造者只有一个的情况下，如果抽象建造者角色已经被省略掉，那么还可以省略指挥者角色，让Builder角色扮演指挥者与建造者双重角色。

建造者模式:vs:抽象工厂模式：

* 与抽象工厂模式相比， **建造者模式返回一个组装好的完整产品** ，而 **抽象工厂模式返回一系列相关的产品，这些产品位于不同的产品等级结构，构成了一个产品族。**
* 在抽象工厂模式中，客户端实例化工厂类，然后调用工厂方法获取所需产品对象，而在建造者模式中，客户端可以不直接调用建造者的相关方法，而是通过指挥者类来指导如何生成对象，包括对象的组装过程和建造步骤，它侧重于一步步构造一个复杂对象，返回一个完整的对象。
* 如果将抽象工厂模式看成 **汽车配件生产工厂** ，生产一个产品族的产品，那么建造者模式就是一个 **汽车组装工厂** ，通过对部件的组装可以返回一辆完整的汽车。

### 优缺点

:red_circle:  你可以分步创建对象， 暂缓创建步骤或递归运行创建步骤。

:red_circle:  生成不同形式的产品时， 你可以复用相同的制造代码。

:red_circle:  **单一职责原则**。 你可以将复杂构造代码从产品的业务逻辑中分离出来。

:black_circle:  由于该模式需要新增多个类， 因此代码整体复杂程度会有所增加。




