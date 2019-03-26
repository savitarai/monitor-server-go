# ***savitar-monitor数据库设计***

## 建表规则
* 按服务器分库
example:mon_127_0_0_1('.'替换成下划线)

* 按属性类型和属性id分表
expamle:commattr_<attr_id>

## 通用监控表结构
|列名|类型|备注|
|----|----|----|
|id|int|自增长|
|attr_id|int|属性ID|
|time_min|int|分钟|
|val|bigint|该分钟的值|
## 统计监控表结构
|列名|类型|备注|
|----|----|----|
|id|int|自增长|
|attr_id|int|属性ID|
|time_min|int|分钟|
|min|bigint|最小值|
|max|bigint|最大值|
|count|bigint|计数|
|sum|bigint|总值|
|first|bigint|第一个值|
|end|bigint|最后一个值|
## 秒级监控表结构
|列名|类型|备注|
|----|----|----|
|id|int|自增长|
|attr_id|int|属性ID|
|time_min|int|分钟|
|sum|bigint|总值|
|sec_0|bigint|第0秒|
|sec_N|bigint|第N秒(0<=N<=59)|
## 属性信息表结构
|列名|类型|备注|
|----|----|----|
|id|int|属性ID|
|name|varchar(128)|属性ID名字|
|description|varchar(256)|属性ID备注|
