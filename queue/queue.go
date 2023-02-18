package queue

type Queue []interface{} //interface{}表示支持任何类型

func (q *Queue) Push(v interface{}) {
	*q = append(*q, v)
}
func (q *Queue) Push2(v int) { //限定他只能输入int
	*q = append(*q, v)
}
func (q *Queue) Push3(v interface{}) { //任何类型
	*q = append(*q, v.(int)) //q.push3("等等") 编译不报错，运行时报错
	//s,ok:= v.(int) 判断v的类型是int则赋值给s,ok为true，否者ok为false
}

func (q *Queue) Pop() interface{} {
	head := (*q)[0]
	*q = (*q)[1:]
	return head
}
func (q *Queue) Pop2() int { //限定只能pop int类型
	head := (*q)[0]
	*q = (*q)[1:]
	return head.(int) ////限定只能pop int类型
}

func (q *Queue) IsEmpty() bool {
	return len(*q) == 0
}
