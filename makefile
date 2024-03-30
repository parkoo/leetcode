# 生成golang题解模版
# eg: make gen n={题目名称} o={解法序号}  >>>  ./题目名称/go/solution_题目序号.go
gen:
	sh scripts/gen.sh $(n) $(o)
