// Create by Zubin on 2018-06-26 14:17:37

package list

import (
	"fmt"
	"testing"
)

func TestCircularLinkList_GetElem(t *testing.T) {
	list := NewCircularLinkList()
	if !list.ListEmpty() {
		t.Error("list not empty")
	}

	list.AddToFirst("T")
	list.ClearList()

	list.AddToLast("D")
	list.AddToLast("E")
	list.AddToFirst("C")
	list.AddToFirst("B")
	list.AddToFirst("A")
	list.Add(4, "G")
	list.Add(4, "F")
	list.RemoveFirst()
	list.RemoveLast()
	list.Remove(2)
	for i := 1; i < list.Size()+1; i++ {
		fmt.Printf("%d : %s\n", i, list.GetElem(i))
	}
	fmt.Printf("Index of 'F': %d\n", list.LocateElem("F"))
	fmt.Println("Reversing...")
	list.Reverse()
	for i := 1; i < list.Size()+1; i++ {
		fmt.Printf("%d : %s\n", i, list.GetElem(i))
	}
	fmt.Printf("Index of 'F': %d\n", list.LocateElem("F"))

	list.ClearList()
	fmt.Println(list.Size())
}
