## 索引生命周期

### 基础啊流程

#### 1.创建索引策略

```json
PUT _ilm/policy/eryajf_policy 
{
  "policy": {                       
    "phases": {
      "hot": {                      
        "actions": {
              "rollover":{
                  "max_age":"30s"
              }
        }
      },
      "delete": {
        "min_age": "90s",           
        "actions": {
          "delete": {}              
        }
      }
    }
  }
}
```

#### 2.创建索引模板

````json
PUT _template/template_eryajf
{
  "index_patterns": ["eryajf-*"],                 
  "settings": {
    "number_of_shards": 1,
    "number_of_replicas": 1,
    "index.lifecycle.name": "eryajf_policy",      
    "index.lifecycle.rollover_alias": "eryajf"    
  }
}
````

#### 3.创建索引

```json
PUT eryajf-1
{
  "aliases": {
    "eryajf": {
      "is_write_index": true
    }
  }
}
```

#### 4. 添加数据

```
PUT /eryajf/_bulk
{"index":{}}
{"message":"hello-01"}
{"index":{}}
{"message":"hello-02"}
{"index":{}}
{"message":"hello-03"}
```



```python
from elasticsearch import Elasticsearch
from pprint import pprint
from common.clients import clients

es = Elasticsearch(['http://192.168.124.16:9200'], timeout=30, maxsize=10, http_compress=True)

# body = {
#   "index_patterns": ["individual_stock_search-*"],
#   "settings": {
#     "number_of_shards": 1,
#     "number_of_replicas": 1,
#     "index.lifecycle.name": "eryajf_policy",
#     "index.lifecycle.rollover_alias": "individual_stock_search"
#   }
# }

# data = es.search(index='company_event')
# data = es.indices.put_template('individual_stock_search', body=body)


# 1.创建 policy
# policy_body = {
#     "policy": {
#         "phases": {
#             "hot": {
#                 "actions": {
#                     "rollover": {
#                         "max_age": "12h"
#                     }
#                 }
#             },
#             "delete": {
#                 "min_age": "12h",
#                 "actions": {
#                     "delete": {}
#                 }
#             }
#         }
#     }
# }
# es.ilm.put_lifecycle('stock_cached_policy', body=policy_body)

# 2. 创建索引模板
# template_body = {
#   "index_patterns": ["stock_cached-*"],
#   "settings": {
#     "number_of_shards": 1,
#     "number_of_replicas": 1,
#     "index.lifecycle.name": "stock_cached_policy",
#     "index.lifecycle.rollover_alias": "stock_cached"
#   }
# }
# es.indices.put_template('stock_cached_template', body=template_body)


# 3.创建索引
# index_body = {
#   "aliases": {
#     "stock_cached": {
#       "is_write_index": "true"
#     }
#   }
# }
# es.indices.create('stock_cached-1', body=index_body)
# data = es.search(index='stock_cached')
# pprint(data)


# 4. 创建数据

```

