# 结构体排序
- 线性化
  - 首先要站成一排，按照某种规则排序（有索引）
  - 遍历，只要是容器，必须能增加、删除
- 线性表
  - Len() int               为了便于计算有多少个，线性数据结构的长度，确定索引范围
  - Less(i, j int) bool     任意2个索引的元素比较大小，返回bool，true表示符合比较，false表示比较失败
  - Swap(i, j int) false    false 交换2个索引的元素
- 排序
  - 有算法（不同算法之间会有差异），有策略
- 方法

```txt
        sort.Ints([]int)
```

    - sort.Ints([]int{4, 1, 3})内部使用了sort.Sort(sort.Interface())，type IntSlice []int就可以扩展3个方法了

```txt
        sort.strings([]string)
```

    - type StringSlice []string 扩展Interface3个方法了
  - sort.Sort(sort.Interface)   所有容器的元素排序，map不是线性排序

- []any 可以进行排序吗？排序的序列，元素类型不一致可以排序吗？
  - 可以，但是一般没有业务意义，需要找到比较的维度
    - less func (i, j int) {

```txt
                return fmt.Sprintf("%v", x[i]) < fmt.Sprintf("%v", x[j])
            }
```

  - 场景
    - n个不同类型的商品放到购物车里面 Cart[]
      - 单价排序
      - 总价排序
      - 降价排序
      - 临期排序
