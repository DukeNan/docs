```python
from collections import OrderedDict
from weakref import proxy as _proxy


class Kls():
    def public(self):
        print('hello public world!')

    def __private(self):
        print('hello private world!')

    def call_private(self):
        self.__private()


class _Link(object):
    __slots__ = 'prev', 'next', 'key', '__weakref__'


class DLinkList:
    def __init__(self):
        self.__hardroot = _Link()
        self.__root = root = _proxy(self.__hardroot)
        root.prev = root.next = root

    def add(self, key):
        link = _Link()
        root = self.__root
        # print(getattr(root, 'key', None))
        print(id(root))
        last = root.prev
        link.prev, link.next, link.key = last, root, key
        last.next = link

        root.prev = _proxy(link)


    def travel(self):
        cur = self.__root
        while cur.next is not self.__root:
            cur = cur.next
            print(cur.key)
            print(cur.next)


    def __iter__(self):
        'od.__iter__() <==> iter(od)'
        # Traverse the linked list in order.
        root = self.__root
        curr = root.next
        while curr is not root:
            yield curr.key
            curr = curr.next




if __name__ == '__main__':
    # dl = DLinkList()
    # dl = DLinkList()
    # dl.add('name')
    # dl.add('age')
    # dl.add('sex')
    # dl.add('score')
    # del dl['age']
    # for i in dl:
    #     print(i)
    d = OrderedDict(name='laowang', age=18, gender='male', score=90)
    d.popitem(last=False)
    print(d)


```

