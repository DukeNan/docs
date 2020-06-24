## YML文件

### 释义

YAML（/ˈjæməl/，尾音类似 camel 骆驼）是一个可读性高，用来表达数据序列化的格式。YAML 参考了其他多种语言，包括： C 语言、 Python、Perl，并从 XML、电子邮件的数据格式（RFC 2822）中获得灵感。Clark Evans 在 2001 年首次发表了这种语言 ，另外 Ingy döt Net 与 Oren Ben-Kiki 也是这语言的共同设计者 。当前已经有数种编程语言或脚本语言支持（或者说解析）这种语言。

YAML 是 "YAML Ain't a Markup Language"（YAML 不是一种标记语言）的递归缩写。在开发的这种语言时， YAML 的意思其实是："Yet Another Markup Language"（仍是一种标记语言，但为了强调这种语言以数据做为中心，而不是以标记语言为重点，而用反向缩略语重命名。

官网：https://yaml.org/

## 功能

YAML 的语法和其他高级语言类似，并且可以简单表达清单、散列表，标量等数据形态。 它使用空白符号缩进和大量依赖外观的特色，特别适合用来表达或编辑数据结构、各种配置文件、倾印调试内容、文件大纲（例如：许多电子邮件标题格式和 YAML 非常接近）。尽管它比较适合用来表达层次结构式（hierarchical model）的数据结构，不过也有精致的语法可以表示关系性（relational model）的数据。由于 YAML 使用空白字符和分行来分隔数据，使得它特别适合用 grep／Python／Perl／Ruby 操作。其让人最容易上手的特色是巧妙避开各种封闭符号，如：引号、各种括号等，这些符号在嵌套结构时会变得复杂而难以辨认。

## 示例

```yaml
receipt:     Oz-Ware Purchase Invoice
date:        2012-08-06
customer:
    given:   Dorothy
    family:  Gale
   
items:
    - part_no:   A4786
      descrip:   Water Bucket (Filled)
      price:     1.47
      quantity:  4

    - part_no:   E1628
      descrip:   High Heeled "Ruby" Slippers
      size:      8
      price:     133.7
      quantity:  1

bill-to:  &id001
    street: | 
            123 Tornado Alley
            Suite 16
    city:   East Centerville
    state:  KS

ship-to:  *id001   

specialDelivery:  >
    Follow the Yellow Brick
    Road to the Emerald City.
    Pay no attention to the
    man behind the curtain.
```

**注意:** 在 YAML 中，字符串不一定要用双引号标示。另外，在缩进中空白字符的数目并不是非常重要，只要相同层次结构的元素左侧对齐就可以了（不过不能使用 TAB 字符）。这个文件的顶层由七个键值组成：其中一个键值 "items"，是两个元素构成的数组（或称清单），这清单中的两个元素同时也是包含了四个键值的散列表。文件中重复的部分用这个方法处理：使用锚点（ &）和引用（*）标签将 "bill-to" 散列表的内容复制到 "ship-to" 散列表。也可以在文件中加入选择性的空行，以增加可读性。在一个文件中，可同时包含多个文件，并用 "---"分隔。选择性的符号"..." 可以用来表示文件结尾（在利用流的通信中，这非常有用，可以在不关闭流的情况下，发送结束信号）。

### 构成元素

#### YAML 的基本组件

YAML 提供缩进／区块以及内置（inline）两种格式，来表示清单和散列表。以下展示几种 YAML 的基本原件。

##### 清单（数组）

```yaml
 --- # 最喜爱的电影
 - Casablanca
 - North by Northwest
 - Notorious
```

另外还有一种内置格式（inline format）可以选择──用方括号围住，并用逗号 + 空白区隔（类似 JSON 的语法）

```yaml
 --- # 购物清单
 [milk, pumpkin pie, eggs, juice]
```

##### 散列表(Map|Dict)

键值和数据由冒号及空白字符分开。区块形式（常使用与 YAML 数据文档中）使用缩进和换行符分隔 key: value 对。内置形式（常使用与 YAML 数据流中）在大括号中使用逗号 + 空白字符分隔 key: value 对。

```yaml
 --- # 区块形式
   name: John Smith
   age: 33
 --- # 内置形式
 {name: John Smith, age: 33}
```

##### 区块的字符

再次强调，字符串不需要包在引号之内。有两种方法书写多行文字（multi-line strings），一种可以保存新行（使用 | 字符），另一种可以折叠新行（使用 > 字符）

##### 保存新行 (Newlines preserved)

```yaml
data: |                                     # 译著者：這是一首著名的五行民谣(limerick)
   There once was a man from Darjeeling     # 这里有个人来自大吉岭
   Who got on a bus bound for Ealing        # 他搭上一班往伊灵的公車
       It said on the door                  # 门上这样说的
       "Please don't spit on the floor"     # "请勿在地上吐痰"
   So he carefully spat on the ceiling      # 说以他小心翼翼的吐在天花板上
```

根据设置，前方的引领空白符号（leading white space）必须排成条状，以便和其他数据或是行为（如示例中的缩进）明显区分。

##### 折叠新行 (Newlines folded)

```yaml
data: >
   Wrapped text         # 折叠的文字
   will be folded       # 将会被收
   into a single        # 进单一一个
   paragraph            # 段落
   
   Blank lines denote   # 空白的行代表
   paragraph breaks     # 段落之间的区隔
```

和保存新行不同的是，换行字符会被转换成空白字符。而引领空白字符则会被自动消去。

#### 层次结构化的元素

##### 于清单中使用散列表

```yaml
- {name: John Smith, age: 33}
- name: Mary Smith
  age: 27
```

##### 于散列表中使用清单

```yaml
men: [John Smith, Bill Jones]
women:
  - Mary Smith
  - Susan Williams
```

### YAML 的高级组件

这部分算是一个后续的讨论，在比较各种数数据列语言时，YAML 最常被提到的特色有两个：关系树和数据形态。

#### 树状结构之间的交互引用

##### 数据合并和参考

为了维持文件的简洁，并避免数据输入的错误，YAML 提供了结点参考（*）和散列合并（<<）参考到其他结点标签的锚点标记（&）。参考会将树状结构加入锚点标记的内容，并可以在所有数据结构中运作（可以参考上面 "ship-to" 的示例）合并只有散列表可以使用，可以将键值自锚点标记复制到指定的散列表中。

当数据被 instantiate 合并和参考会被剖析器自动展开。

```yaml
#眼部镭射手术之标准程序
---
- step:  &id001                  # 定义锚标签 &id001
    instrument:      Lasik 2000
    pulseEnergy:     5.4
    pulseDuration:   12
    repetition:      1000
    spotSize:        1mm

- step:
     <<: *id001                  # 合并键值：使用在锚标签定义的内容
     spotSize:       2mm         # 复写“spotSize”键值

- step:
     <<: *id001                  # 合并键值：使用在锚标签定义的内容
     pulseEnergy:    500.0       # 复写键值
     alert: >                    # 加入其他鍵值
           warn patient of 
           audible pop
```

##### 数据形态

由于自动判定数据形态的功能，严格类型（也就是用户有宣告的数据形态）很难在大部分的 YAML 文件中看到。数据类型可以被区分成三大类：原码（core），定义（defined），用户定义（user-defined）。原码可以自动被解析器分析（例如：浮点数，整数，字符串，清单，映射，...）。有一些高级的数据形态──例如比特数据──在 YAML 中有被 “定义”，但不是每一种解析器都有支持。最后，YAML 支持用户自定的区域变量，包括：自定义的类别，结构或基本类型（例如：四倍精度的浮点数）。

##### 强制转型

YAML 的自动判定数据形态是哪一种实体。但有时用户会想要将数据强制转型成自定的某种类型。最常见的状况是字符串，有时候可能看起来像数字或布尔值，这种时候可以使用双引号，或是使用严格类型标签。

```yaml
---
a: 123                     # 整数
b: "123"                   # 字串（使用双括号）
c: 123.0                   # 浮点数
d: !!float 123             # 浮点数，使用!!表达的严格型态
e: !!str 123               # 字串，使用严格形态
f: !!str Yes               # 字串，使用严格形态
g: Yes                     # 布林值"真"
h: Yes we have No bananas  # 字串（包含"Yes"和"No"）
```

##### 其他特殊数据形态

除了一般的数据形态之外，用户也可以使用一些较为高级的类型，但不保证可被每种解析器分析。使用时和强制转型类似，要在形态名称之前加上两个惊叹号（!!）。有几种重要的形态在本篇没有讨论，包括集合（sets），有序映照（ordered maps），时间戳记（timestamps）以及十六进制数据（hexadecimal）。下面这个示例则是比特数据（binary）

```yaml
---
picture: !!binary |
 R0lGODlhDAAMAIQAAP//9/X
 17unp5WZmZgAAAOfn515eXv
 Pz7Y6OjuDg4J+fn5OTk6enp
 56enmleECcgggoBADs=mZmE
```

##### 用户自行扩展的数据形态

许多 YAML 的实现允许用户自定义数据形态。在将一个对象序列化时，这个方法还颇方便的。某些区域数据形态可能不存在默认的数据形态中，不过这种类型在特定的 YAML 应用程序中是有定义的。这种区域数据形态用惊叹号（ !）表示。

```yaml
myObject:  !myClass { name: Joe, age: 15}
```

### 语法

在 yaml.org（英文）可以找到轻巧而好用的 小抄（亦是用 YAML 表示）及格式说明。下面的内容，是关于基本组件的摘要。

YAML 使用可打印的 Unicode 字符，可使用 UTF-8 或 UTF-16。
 使用空白字符为文件缩进来表示结构；不过不能使用跳格字符 (TAB)。
 注解由井字号（ # ）开始，可以出现在一行中的任何位置，而且范围只有一行（也就是一般所谓的单行注解）
 每个清单成员以单行表示，并用短杠 + 空白（ -   ）起始。或使用方括号（ [ ] ），并用逗号 + 空白（ ,   ）分开成员。
 每个散列表的成员用冒号 + 空白（ :   ）分开键值和内容。或使用大括号（ {   } ），并用逗号 + 空白（ ,   ）分开。
 散列表的键值可以用问号 ( ? ) 起始，用来明确的表示多个词汇组成的键值。
 字符串平常并不使用引号，但必要的时候可以用双引号 ( " ) 或单引号 ( ' ) 框住。
 使用双引号表示字符串时，可用倒斜线（ \ ）开始的转义字符（这跟 C 语言类似）表示特殊字符。
 区块的字符串用缩进和修饰符（非必要）来和其他数据分隔，有新行保留（preserve）（使用符号 | ）或新行折叠（flod）（使用符号 > ）两种方式。
 在单一文件中，可用连续三个连字号（---）区分多个文件。
 另外，还有选择性的连续三个点号（ ... ）用来表示文件结尾。
 重复的内容可使从参考标记星号 ( * ) 复制到锚点标记（ & ）。
 指定格式可以使用两个惊叹号 ( !! )，后面接上名称。
 文件中的单一文件可以使用指导指令，使用方法是百分比符号 ( % )。有两个指导指令在 YAML1.1 版中被定义：
 % YAML 指导指令，用来识别文件的 YAML 版本。
 % TAG 指导指令，被用在 URI 的前缀标记。这个方法在标记节点的类型时相当有用。
 YAML 在使用逗号及冒号时，后面都必须接一个空白字符，所以可以在字符串或数值中自由加入分隔符号（例如：5,280 或 [http://www.wikipedia.org](https://links.jianshu.com/go?to=http%3A%2F%2Fwww.wikipedia.org)）而不需要使用引号。

另外还有两个特殊符号在 YAML 中被保留，有可能在未来的版本被使用 --（ @ ）和（ ` ）。

### PyYAML

创建data.yml文件

```yaml
receipt:     Oz-Ware Purchase Invoice
date:        2012-08-06
customer:
    given:   Dorothy
    family:  Gale
   
items:
    - part_no:   A4786
      descrip:   Water Bucket (Filled)
      price:     1.47
      quantity:  4

    - part_no:   E1628
      descrip:   High Heeled "Ruby" Slippers
      size:      8
      price:     133.7
      quantity:  1

bill-to:  &id001
    street: | 
            123 Tornado Alley
            Suite 16
    city:   East Centerville
    state:  KS

ship-to:  *id001   

specialDelivery:  >
    Follow the Yellow Brick
    Road to the Emerald City.
    Pay no attention to the
    man behind the curtain.
```

读取data.yml

```python
import yaml
from pprint import pprint


with open('data.yml', 'r', encoding='utf-8') as f:
    yaml_data = f.read()

data = yaml.safe_load(yaml_data)
pprint(data)

# -------------------------------------------------------------------------------------
[out]:
{'bill-to': {'city': 'East Centerville',
             'state': 'KS',
             'street': '123 Tornado Alley\nSuite 16\n'},
 'customer': {'family': 'Gale', 'given': 'Dorothy'},
 'date': datetime.date(2012, 8, 6),
 'items': [{'descrip': 'Water Bucket (Filled)',
            'part_no': 'A4786',
            'price': 1.47,
            'quantity': 4},
           {'descrip': 'High Heeled "Ruby" Slippers',
            'part_no': 'E1628',
            'price': 133.7,
            'quantity': 1,
            'size': 8}],
 'receipt': 'Oz-Ware Purchase Invoice',
 'ship-to': {'city': 'East Centerville',
             'state': 'KS',
             'street': '123 Tornado Alley\nSuite 16\n'},
 'specialDelivery': 'Follow the Yellow Brick Road to the Emerald City. Pay no '
                    'attention to the man behind the curtain.'}
```

