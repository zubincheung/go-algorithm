// Create by Zubin on 2018-08-28 15:30:33

package sort

import "reflect"

// 交换list中下标为i,j的值
func swap(list []int, i int, j int) {
	tmp := list[i]
	list[i] = list[j]
	list[j] = tmp
}

// SliceEquals 用以比较两个Slice(基础数据类型,如[]int)内含值是否相等
func SliceEquals(a, b interface{}) bool {
	// a,有任意一个不是slice返回false
	_a := reflect.ValueOf(a)
	if _a.Kind() != reflect.Slice {
		panic("param a must be a slice")
	}
	_b := reflect.ValueOf(b)
	if _b.Kind() != reflect.Slice {
		panic("param a must be a slice")
	}
	// 长度不等则两个slice不同
	if _a.Len() != _b.Len() {
		return false
	}
	// 依次比较每个值
	for i := 0; i < _a.Len(); i++ {
		if _a.Index(i).Interface() != _b.Index(i).Interface() {
			return false
		}
	}
	return true
}
