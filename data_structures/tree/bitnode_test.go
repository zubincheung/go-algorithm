package tree

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {

	t.Run("test", func(t *testing.T) {
		list := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
		node := CreateTree(list, 1)

		PreOrderTraverse(node)
		fmt.Printf("\n")

		InOrderTraverse(node)
		fmt.Printf("\n")

		PostOrderTraverse(node)
		fmt.Printf("\n")
	})
}
