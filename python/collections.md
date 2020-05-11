

## 容器数据类型

这个模块实现了特定目标的容器，以提供Python标准内建容器 dict , list , set , 和 tuple 的替代选择。

| 类型         | 说明                                                         |
| :----------- | :----------------------------------------------------------- |
| namedtuple() | 创建命名元组子类的工厂函数                                   |
| deque        | 类似列表(list)的容器，实现了在两端快速添加(append)和弹出(pop) |
| ChainMap     | 类似字典(dict)的容器类，将多个映射集合到一个视图里面         |
| Counter      | 字典的子类，提供了可哈希对象的计数功能                       |
| OrderedDict  | 字典的子类，保存了他们被添加的顺序                           |
| defaultdict  | 字典的子类，提供了一个工厂函数，为字典查询提供一个默认值     |
| UserDict     | 封装了字典对象，简化了字典子类化                             |
| UserList     | 封装了列表对象，简化了列表子类化                             |
| UserString   | 封装了列表对象，简化了字符串子类化                           |



## Counter 对象

一个计数器工具提供快速和方便的计数。比如

```python
>>> # Tally occurrences of words in a list
>>> cnt = Counter()
>>> for word in ['red', 'blue', 'red', 'green', 'blue', 'blue']:
...     cnt[word] += 1
>>> cnt
Counter({'blue': 3, 'red': 2, 'green': 1})

>>> # Find the ten most common words in Hamlet
>>> import re
>>> words = re.findall(r'\w+', open('hamlet.txt').read().lower())
>>> Counter(words).most_common(10)
[('the', 1143), ('and', 966), ('to', 762), ('of', 669), ('i', 631),
 ('you', 554),  ('a', 546), ('my', 514), ('hamlet', 471), ('in', 451)]

```

一个 Counter 是一个 dict 的子类，用于计数可哈希对象。它是一个集合，元素像字典键(key)一样存储，它们的计数存储为值。计数可以是任何整数值，包括0和负数。 Counter 类有点像其他语言中的 bags或multisets。

元素从一个 *iterable* 被计数或从其他的 *mapping* (or counter)初始化：

```python
c = Counter()                           # a new, empty counter
c = Counter('gallahad')                 # a new counter from an iterable
c = Counter({'red': 4, 'blue': 2})      # a new counter from a mapping
c = Counter(cats=4, dogs=8)             # a new counter from keyword args
```

Counter对象有一个字典接口，如果引用的键没有任何记录，就返回一个0，而不是弹出一个 `KeyError` :

```python
c = Counter(['eggs', 'ham'])
>>> c['bacon']                              # count of a missing element is zero
0
```

设置一个计数为0不会从计数器中移去一个元素。使用 `del` 来删除它:

```python
c['sausage'] = 0                        # counter entry with a zero count
>>> del c['sausage'] 
```

### elements()

返回一个迭代器，其中每个元素将重复出现计数值所指定次。 元素会按首次出现的顺序返回。 如果一个元素的计数值小于1，elements() 将会忽略它。

```python
c = Counter(a=4, b=2, c=0, d=-2)
>>> sorted(c.elements())
['a', 'a', 'a', 'a', 'b', 'b']
```

### most_common([n])

返回一个列表，其中包含 n 个最常见的元素及出现次数，按常见程度由高到低排序。 如果 n 被省略或为 None，most_common() 将返回计数器中的 所有 元素。 计数值相等的元素按首次出现的顺序排序：

```python
Counter('abracadabra').most_common(3)
[('a', 5), ('b', 2), ('r', 2)]
```

### `subtract`([*iterable-or-mapping*])

从 迭代对象 或 映射对象 减去元素。像 dict.update() 但是是减去，而不是替换。输入和输出都可以是0或者负数。

```python
c = Counter(a=4, b=2, c=0, d=-2)
>>> d = Counter(a=1, b=2, c=3, d=4)
>>> c.subtract(d)
>>> c
Counter({'a': 3, 'b': 0, 'c': -3, 'd': -6})
```

### `update`([*iterable-or-mapping*])

从 迭代对象 计数元素或者 从另一个 映射对象 (或计数器) 添加。 像 dict.update() 但是是加上，而不是替换。另外，迭代对象 应该是序列元素，而不是一个 (key, value) 对。

```python
>>> c = Counter(a=1, b=1, c=1)
>>> c.update(a=1, b=1, c=1)
>>> c
Counter({'a': 2, 'b': 2, 'c': 2})

```

### 常用案例

提供了几个数学操作，可以结合 Counter 对象，以生产 multisets (计数器中大于0的元素）。 加和减，结合计数器，通过加上或者减去元素的相应计数。交集和并集返回相应计数的最小或最大值。每种操作都可以接受带符号的计数，但是输出会忽略掉结果为零或者小于零的计数。

```python
c = Counter(a=3, b=1)
>>> d = Counter(a=1, b=2)
>>> c + d                       # add two counters together:  c[x] + d[x]
Counter({'a': 4, 'b': 3})
>>> c - d                       # subtract (keeping only positive counts)
Counter({'a': 2})
>>> c & d                       # intersection:  min(c[x], d[x]) 
Counter({'a': 1, 'b': 1})
>>> c | d                       # union:  max(c[x], d[x])
Counter({'a': 3, 'b': 2})
```

单目加和减（一元操作符）意思是从空计数器加或者减去。

```python
c = Counter(a=2, b=-4)
>>> +c
Counter({'a': 2})
>>> -c
Counter({'b': 4})
```



## deque 对象

- 返回一个新的双向队列对象，从左到右初始化(用方法 append()) ，从 iterable （迭代对象) 数据创建。如果 iterable 没有指定，新队列为空。
- Deque队列是由栈或者queue队列生成的。Deque 支持线程安全，内存高效添加(append)和弹出(pop)，从两端都可以，两个方向的大概开销都是 O(1) 复杂度。
- 虽然 list 对象也支持类似操作，不过这里优化了定长操作和 pop(0) 和 insert(0, v) 的开销。它们引起 O(n) 内存移动的操作，改变底层数据表达的大小和位置。
- 如果 maxlen 没有指定或者是 None ，deques 可以增长到任意长度。否则，deque就限定到指定最大长度。一旦限定长度的deque满了，当新项加入时，同样数量的项就从另一端弹出。限定长度deque提供类似Unix filter tail 的功能。它们同样可以用与追踪最近的交换和其他数据池活动。

### 方法

- append(x)
  	- 添加 *x* 到右端。
- appendleft(*x*)
   - 添加 *x* 到左端。
- clear()
   - 移除所有元素，使其长度为0.
- copy()
   - 创建一份浅拷贝。
- count(x)
   - 计算 deque 中元素等于 *x* 的个数。
- extend(iterable)
   - 扩展deque的右侧，通过添加iterable参数中的元素。
- extendleft(iterable)
   - 扩展deque的左侧，通过添加iterable参数中的元素。注意，左添加时，在结果中iterable参数中的顺序将被反过来添加。
- index(x[, start[, stop]])
   - 返回 x 在 deque 中的位置（在索引 start 之后，索引 stop 之前）。 返回第一个匹配项，如果未找到则引发` ValueError`。
- insert(i, x)
   - 在位置 *i* 插入 *x* 。
   - 如果插入会导致一个限长 deque 超出长度 maxlen 的话，就引发一个 `IndexError`。
- pop()
   - 移去并且返回一个元素，deque 最右侧的那一个。 如果没有元素的话，就引发一个 `IndexError`。
- popleft()
   - 移去并且返回一个元素，deque 最左侧的那一个。 如果没有元素的话，就引发 `IndexError`。
- remove(value)
   - 移除找到的第一个 value。 如果没有的话就引发 `ValueError`。
- reverse()
   - 将deque逆序排列。返回 `None` 。
- rotate(n=1)
   - 向右循环移动 *n* 步。 如果 *n* 是负数，就向左循环。
   - 如果deque不是空的，向右循环移动一步就等价于 `d.appendleft(d.pop())` ， 向左循环一步就等价于 `d.append(d.popleft())` 。
- maxlen
   - Deque的最大尺寸，如果没有限定的话就是 `None` 。



## defaultdict 对象

- 返回一个新的类似字典的对象。 defaultdict 是内置 dict 类的子类。它重载了一个方法并添加了一个可写的实例变量。其余的功能与 dict 类相同，此处不再重复说明。
- 本对象包含一个名为 default_factory 的属性，构造时，第一个参数用于为该属性提供初始值，默认为 None。所有其他参数（包括关键字参数）都相当于传递给 dict 的构造函数。

defaultdict 对象除了支持标准 dict 的操作，还支持以下方法作为扩展：

- ` __missing__`(*key*)
  - 如果 default_factory 属性为 None，则调用本方法会抛出 `KeyError` 异常，附带参数 key。
  - 如果 default_factory 不为 `None`，则它会被（不带参数地）调用来为 key 提供一个默认值，这个值和 key 作为一对键值对被插入到字典中，并作为本方法的返回值返回。
  - 如果调用 default_factory 时抛出了异常，这个异常会原封不动地向外层传递。
  - 在无法找到所需键值时，本方法会被 dict 中的 `__getitem__()` 方法调用。无论本方法返回了值还是抛出了异常，都会被 `__getitem__()` 传递。
- **default_factory**
  - 本属性由 `__missing__()`方法来调用。如果构造对象时提供了第一个参数，则本属性会被初始化成那个参数，如果未提供第一个参数，则本属性为 `None`。

### 实例

使用 list 作为 default_factory，很轻松地将（键-值对组成的）序列转换为（键-列表组成的）字典：

```python
s = [('yellow', 1), ('blue', 2), ('yellow', 3), ('blue', 4), ('red', 1)]
>>> d = defaultdict(list)
>>> for k, v in s:
...     d[k].append(v)
...
>>> sorted(d.items())
[('blue', [2, 4]), ('red', [1]), ('yellow', [1, 3])]
```

当每个键第一次遇见时，它还没有在字典里面，所以自动创建该条目，即调用 default_factory 方法，返回一个空的 list。 list.append() 操作添加值到这个新的列表里。当再次存取该键时，就正常操作，list.append() 添加另一个值到列表中。这个计数比它的等价方法 dict.setdefault() 要快速和简单：

```python
d = {}
>>> for k, v in s:
...     d.setdefault(k, []).append(v)
...
>>> sorted(d.items())
[('blue', [2, 4]), ('red', [1]), ('yellow', [1, 3])]
```

设置 default_factory 为 int，使 defaultdict 用于计数:

```python
s = 'mississippi'
>>> d = defaultdict(int)
>>> for k in s:
...     d[k] += 1
...
>>> sorted(d.items())
[('i', 4), ('m', 1), ('p', 2), ('s', 4)]
```

设置 default_factory 为 set 使 defaultdict 用于构建 set 集合：

```python
s = [('red', 1), ('blue', 2), ('red', 3), ('blue', 4), ('red', 1), ('blue', 4)]
>>> d = defaultdict(set)
>>> for k, v in s:
...     d[k].add(v)
...
>>> sorted(d.items())
[('blue', {2, 4}), ('red', {1, 3})]
```

 **注意**

 defaultdict["name"] 方式会创建一个缺省值

defaultdict.get("name")方式不会新创建一个值



## namedtuple() 命名元组的工厂函数

```python
>>> # Basic example
>>> Point = namedtuple('Point', ['x', 'y'])
>>> p = Point(11, y=22)     # instantiate with positional or keyword arguments
>>> p[0] + p[1]             # indexable like the plain tuple (11, 22)
33
>>> x, y = p                # unpack like a regular tuple
>>> x, y
(11, 22)
>>> p.x + p.y               # fields also accessible by name
33
>>> p                       # readable __repr__ with a name=value style
```

除了继承元组的方法，命名元组还支持三个额外的方法和两个属性。为了防止域名冲突，方法和属性以下划线开始。

- `_make(iterable)`

  ```python
  from collections import namedtuple
  
  Point = namedtuple('Point', ['x', 'y'])
  t = [11, 22]
  Point._make(t) # Point(x=11, y=22)
  ```

- `_asdict()`

  返回一个新的 dict ，它将字段名称映射到它们对应的值：
  
  ```python
  >>> p = Point(x=11, y=22)
  >>> p._asdict()
  {'x': 11, 'y': 22}
  ```
  
- `_replace(**kwargs)`

  返回一个新的命名元组实例，并将指定域替换为新的值

  ```python
  >>> p = Point(x=11, y=22)
  >>> p._replace(x=33)
  Point(x=33, y=22)
  
  >>> for partnum, record in inventory.items():
    ...     inventory[partnum] = record._replace(price=newprices[partnum],timestamp=time.now())
  ```

- `_fields`

  字符串元组列出了域名。用于提醒和从现有元组创建一个新的命名元组类型。

  ```python
  >>> p._fields            # view the field names
  ('x', 'y')
  
  >>> Color = namedtuple('Color', 'red green blue')
  >>> Pixel = namedtuple('Pixel', Point._fields + Color._fields)
  >>> Pixel(11, 22, 128, 255, 0)
  Pixel(x=11, y=22, red=128, green=255, blue=0)
  ```

- `_field_defaults`

  默认值的字典。

  ```python
  >>> Account = namedtuple('Account', ['type', 'balance'], defaults=[0])
  >>> Account._field_defaults
  {'balance': 0}
  >>> Account('premium')
  Account(type='premium', balance=0)
  ```



**注意：**

Python > 3.5，使用 `typing.NamedTuple`



## OrderedDict

### 方法

- popitem(last=True)

  有序字典的 `popitem()` 方法移除并返回一个 (key, value) 键值对。 如果 last 值为真，则按 LIFO 后进先出的顺序返回键值对，否则就按 FIFO 先进先出的顺序返回键值对。

- move_to_end(key, last=True)

  将现有 key 移动到有序字典的任一端。 如果 last 为真值（默认）则将元素移至末尾；如果 last 为假值则将元素移至开头。如果 key 不存在则会触发 `KeyError`:

  ```python
  >>> d = OrderedDict.fromkeys('abcde')
  >>> d.move_to_end('b')
  >>> ''.join(d.keys())
  'acdeb'
  >>> d.move_to_end('b', last=False)
  >>> ''.join(d.keys())
  'bacde'
  ```

### OrderedDict实例

创建记住键值 最后 插入顺序的有序字典变体很简单。 如果新条目覆盖现有条目，则原始插入位置将更改并移至末尾:

```python
class LastUpdatedOrderedDict(OrderedDict):
    'Store items in the order the keys were last added'

    def __setitem__(self, key, value):
        super().__setitem__(key, value)
        self.move_to_end(key)
```

一个 OrderedDict 对于实现 functools.lru_cache() 的变体也很有用:

```python
class LRU(OrderedDict):
    'Limit size, evicting the least recently looked-up key when full'

    def __init__(self, maxsize=128, /, *args, **kwds):
        self.maxsize = maxsize
        super().__init__(*args, **kwds)

    def __getitem__(self, key):
        value = super().__getitem__(key)
        self.move_to_end(key)
        return value

    def __setitem__(self, key, value):
        super().__setitem__(key, value)
        if len(self) > self.maxsize:
            oldest = next(iter(self))
            del self[oldest]
```

