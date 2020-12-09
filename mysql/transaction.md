##  MySQL事务

### 简介

事务是MySQL区别于NoSQL的重要特征，是保证关系型数据库数据一致性的关键技术。事务可看作是对数据库操作的基本执行单元，可能包含一个或者多个SQL语句。这些语句在执行时，要么都执行，要么都不执行。



![mysql事务](../_media/images/mysql/transaction_01.png)