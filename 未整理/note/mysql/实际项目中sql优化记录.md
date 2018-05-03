####判断"select"查询是否使用了索引
```
explain SELECT * FROM table
```
* type
	* ALL `全表扫描`
	* const `索引查找`
* key
	* NULL `没有使用索引`
	* PRIMARY `使用了主键`
	* 其他 `索引的名字`

####php循环查询
```
foreach($data as $k => $v)
{
	$sql = 'SELECT * FROM 表名 WHERE suoyin = 1 AND id = $v['id']';
}
```
#####结论：
*  1、以上查询耗费时间会根据循环次数增长
> 原因：每次循环都需要通过索引去筛选符合数据，而使用"in"语句则只需要一次索引查找，后面的条件可以根据索引筛选出来的数据中匹配
#####解决：
```
$sql = 'SELECT * FROM 表名 WHERE suoyin = 1 AND id IN ($v['id']');
```


####总结
#####1、
#####2、
