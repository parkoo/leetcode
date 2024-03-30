# 生成golang题解模板文件
mkdir -p $1/go && touch $1/go/solution_$2.go

# 模板定义
template="package main\n\n// 思路: \n\n// 时间复杂度: O(?)    空间复杂度: O(?)\n\n"

# 向模版文件中写入模版
echo "$template" > $1/go/solution_$2.go  # 在变量的引用的时候加入双引号echo "$my_variable", 可以避免变量中的多个空格被合并为一个