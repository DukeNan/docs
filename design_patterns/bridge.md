## 桥接模式

### 意图

**桥接模式**是一种结构型设计模式， 可将一个大类或一系列紧密相关的类拆分为抽象和实现两个独立的层次结构， 从而能在开发时分别使用。

![bridge01](../_media/images/design_patterns/bridge01.png)

### 模式结构

桥接（Bridge）模式包含以下主要角色：

1. 抽象化（Abstraction）角色：

   定义抽象类，并包含一个对实现化对象的引用。

2. 扩展抽象化（Refined  Abstraction）角色：

   是抽象化角色的子类，实现父类中的业务方法，并通过组合关系调用实现化角色中的业务方法。

3. 实现化（Implementor）角色：

   定义实现化角色的接口，供扩展抽象化角色调用。

4. 具体实现化（Concrete Implementor）角色：

   给出实现化角色接口的具体实现。

![bridge02](../_media/images/design_patterns/bridge02.gif)

### 实现方式

1. 明确类中独立的维度。 独立的概念可能是： 抽象/平台， 域/基础设施， 前端/后端或接口/实现。
2. 了解客户端的业务需求， 并在抽象基类中定义它们。
3. 确定在所有平台上都可执行的业务。 并在通用实现接口中声明抽象部分所需的业务。
4. 为你域内的所有平台创建实现类， 但需确保它们遵循实现部分的接口。
5. 在抽象类中添加指向实现类型的引用成员变量。 抽象部分会将大部分工作委派给该成员变量所指向的实现对象。
6. 如果你的高层逻辑有多个变体， 则可通过扩展抽象基类为每个变体创建一个精确抽象。
7. 客户端代码必须将实现对象传递给抽象部分的构造函数才能使其能够相互关联。 此后， 客户端只需与抽象对象进行交互， 无需和实现对象打交道。

```python
from __future__ import annotations
from abc import ABC, abstractmethod


class Abstraction:
    """
    The Abstraction defines the interface for the "control" part of the two
    class hierarchies. It maintains a reference to an object of the
    Implementation hierarchy and delegates all of the real work to this object.

    '抽象部分'定义了两个类层次结构中'控制'部分的接口。
    它管理着一个指向'实现部分'层次结构中对象的引用，并会将所有真实工作委派给该对象。
    """

    def __init__(self, implementation: Implementation) -> None:
        self.implementation = implementation

    def operation(self) -> str:
        return (f"Abstraction: Base operation with:\n"
                f"{self.implementation.operation_implementation()}")


class ExtendedAbstraction(Abstraction):
    """
    You can extend the Abstraction without changing the Implementation classes.

    您可以扩展Abstraction，而无需更改实现类。
    """

    def operation(self) -> str:
        return (f"ExtendedAbstraction: Extended operation with:\n"
                f"{self.implementation.operation_implementation()}")


class Implementation(ABC):
    """
    The Implementation defines the interface for all implementation classes. It
    doesn't have to match the Abstraction's interface. In fact, the two
    interfaces can be entirely different. Typically the Implementation interface
    provides only primitive operations, while the Abstraction defines higher-
    level operations based on those primitives.

    “实现部分”接口声明了在所有具体实现类中通用的方法。它不需要与抽象接口相匹配。
    实际上，这两个接口可以完全不一样。通常实现接口只提供原语操作，而抽象接口则会基于这些操作定义较高层次的操作。
    """

    @abstractmethod
    def operation_implementation(self) -> str:
        pass


"""
Each Concrete Implementation corresponds to a specific platform and implements
the Implementation interface using that platform's API.

每个具体实现都对应于一个特定平台，并使用该平台的API来实现实现接口。
"""


class ConcreteImplementationA(Implementation):

    def operation_implementation(self) -> str:
        return "ConcreteImplementationA: Here's the result on the platform A."


class ConcreteImplementationB(Implementation):
    def operation_implementation(self) -> str:
        return "ConcreteImplementationB: Here's the result on the platform B."


def client_code(abstraction: Abstraction) -> None:
    """
    Except for the initialization phase, where an Abstraction object gets linked
    with a specific Implementation object, the client code should only depend on
    the Abstraction class. This way the client code can support any abstraction-
    implementation combination.
    """
    print(abstraction.operation(), end="")


if __name__ == '__main__':
    """
    The client code should be able to work with any pre-configured abstraction-
    implementation combination.
    
    客户端代码应该能够与任何预配置的抽象实现组合一起使用。
    """
    implementation = ConcreteImplementationA()
    abstraction = Abstraction(implementation)
    client_code(abstraction)

    print('\n')

    implementation = ConcreteImplementationB()
    abstraction = Abstraction(implementation)
    client_code(abstraction)

```

[out]:

```python
Abstraction: Base operation with:
ConcreteImplementationA: Here's the result on the platform A.

Abstraction: Base operation with:
ConcreteImplementationB: Here's the result on the platform B.
```

### 优缺点

:red_circle:  你可以创建与平台无关的类和程序。

:red_circle:  客户端代码仅与高层抽象部分进行互动， 不会接触到平台的详细信息。

:red_circle:  **开闭原则**。 你可以新增抽象部分和实现部分， 且它们之间不会相互影响。

:red_circle:  **单一职责原则**。 抽象部分专注于处理高层逻辑， 实现部分处理平台细节。

:black_circle:  对高内聚的类使用该模式可能会让代码更加复杂。

