## 适配器模式

### 意图

**适配器模式**是一种结构型设计模式， 它能使接口不兼容的对象能够相互合作。

![adapter01](../_media/images/design_patterns/adapter01.png)

### 模式结构

适配器模式（Adapter）包含以下主要角色：

1. 目标（Target）接口：当前系统业务所期待的接口，它可以是抽象类或接口。
2. 适配者（Adaptee）类：它是被访问和适配的现存组件库中的组件接口。
3. 适配器（Adapter）类：它是一个转换器，通过继承或引用适配者的对象，把适配者接口转换成目标接口，让客户按目标接口的格式访问适配者。
4. 客户类



适配器模式有对象适配器和类适配器两种实现：

对象适配器：

![adapter02](../_media/images/design_patterns/adapter02.jpg)

类适配器：

![adapter03](../_media/images/design_patterns/adapter03.jpg)

### 实现方式

1. 确保至少有两个类的接口不兼容：
   * 一个无法修改 （通常是第三方、 遗留系统或者存在众多已有依赖的类） 的功能性*服务*类。
   * 一个或多个将受益于使用服务类的*客户端*类。
2. 声明客户端接口， 描述客户端如何与服务交互。
3. 创建遵循客户端接口的适配器类。 所有方法暂时都为空。
4. 在适配器类中添加一个成员变量用于保存对于服务对象的引用。 通常情况下会通过构造函数对该成员变量进行初始化， 但有时在调用其方法时将该变量传递给适配器会更方便。
5. 依次实现适配器类客户端接口的所有方法。 适配器会将实际工作委派给服务对象， 自身只负责接口或数据格式的转换。
6. 客户端必须通过客户端接口使用适配器。 这样一来， 你就可以在不影响客户端代码的情况下修改或扩展适配器。

```python
class Target:
    """
    The Target defines the domain-specific interface used by the client code.
    Target 定义了客户端代码使用的特定于域的接口。
    """

    def request(self):
        return "Target: The default target's behavior."


class Adaptee:
    """
    The Adaptee contains some useful behavior, but its interface is incompatible
    with the existing client code. The Adaptee needs some adaptation before the
    client code can use it.

    Adaptee包含一些有用的行为，但是其接口与现有的客户端代码不兼容。 在客户端代码可以使用它之前，Adaptee需要进行一些调整。
    """

    def specific_request(self) -> str:
        return ".eetpadA eht fo roivaheb laicepS"


class Adapter(Target, Adaptee):
    """
    The Adapter makes the Adaptee's interface compatible with the Target's
    interface via multiple inheritance.

    适配器通过多次继承使Adaptee的接口与Target的接口兼容。
    """

    def request(self):
        return f"Adapter: (TRANSLATED) {self.specific_request()[::-1]}"


def client_code(target: Target) -> None:
    """
    The client code supports all classes that follow the Target interface.
    客户端代码支持遵循Target接口的所有类。
    """
    print(target.request(), end="")


if __name__ == '__main__':
    print("Client: I can work just fine with the Target objects:")
    target = Target()
    client_code(target)
    print("\n")

    adaptee = Adaptee()
    print("Client: The Adaptee class has a weird interface. "
          "See, I don't understand it:")
    print(f"Adaptee: {adaptee.specific_request()}", end="\n\n")

    print("Client: But I can work with it via the Adapter:")
    adapter = Adapter()
    client_code(adapter)

```

### 优缺点

:red_circle:  **单一职责原则**你可以将接口或数据转换代码从程序主要业务逻辑中分离。

:red_circle:  **开闭原则**。 只要客户端代码通过客户端接口与适配器进行交互， 你就能在不修改现有客户端代码的情况下在程序中添加新类型的适配器。

:black_circle:  代码整体复杂度增加， 因为你需要新增一系列接口和类。 有时直接更改服务类使其与其他代码兼容会更简单。
