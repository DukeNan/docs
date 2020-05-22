

![微信支付](https://ss0.bdstatic.com/94oJfD_bAAcT8t7mm9GUKT-xh_/timg?image&quality=100&size=b4000_4000&sec=1586502050&di=240106366d9ef7b5d10d42f9bca782dc&src=http://dmimg.5054399.com/allimg/pkm/pk/22.jpg ':size=300x300')

***

![logo1](https://docsify.js.org/_media/icon.svg ':size=WIDTHxHEIGHT')
![logo1](https://docsify.js.org/_media/icon.svg ':size=WIDTHxHEIGHT')
![logo1](https://docsify.js.org/_media/icon.svg ':size=WIDTHxHEIGHT')



<div style="align: center">
<img src="https://docsify.js.org/_media/icon.svg"/>
</div>

***


![logo](https://docsify.js.org/_media/icon.svg ':size=10%')


***


<img src="https://docsify.js.org/_media/icon.svg" width = "100" height = "100" div align=right/>


<img src="https://docsify.js.org/_media/icon.svg" width = "100" height = "100" div align=right/>








```python
import gevent
from gevent.local import local

stash = local()

def f1():
    stash.x = 1
    print(stash.x)

def f2():
    stash.y = 2
    print(stash.y)

    try:
        stash.x
    except AttributeError:
        print("x is not local to f2")

g1 = gevent.spawn(f1)
g2 = gevent.spawn(f2)

gevent.joinall([g1, g2])

```

