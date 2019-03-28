package tree

import "fmt"

type BitNode struct {
	Data   int
	lchild *BitNode
	rchild *BitNode
}

// 二叉树顺序存储改链式存储
func CreateTree(list []int, pos int) *BitNode {
	if list == nil || pos > len(list) {
		return nil
	}
	return &BitNode{
		Data:   list[pos-1],
		lchild: CreateTree(list, 2*pos),
		rchild: CreateTree(list, 2*pos+1),
	}
}

// 二叉树的前序遍历递归算法
func PreOrderTraverse(node *BitNode) {
	if node == nil {
		return
	}

	fmt.Printf("%d\t", node.Data)
	PreOrderTraverse(node.lchild)
	PreOrderTraverse(node.rchild)
}

// 二叉树的中序遍历递归算法
func InOrderTraverse(node *BitNode) {
	if node == nil {
		return
	}

	InOrderTraverse(node.lchild)
	fmt.Printf("%d\t", node.Data)
	InOrderTraverse(node.rchild)
}

// 二叉树的后序遍历递归算法
func PostOrderTraverse(node *BitNode) {
	if node == nil {
		return
	}

	PostOrderTraverse(node.lchild)
	PostOrderTraverse(node.rchild)
	fmt.Printf("%d\t", node.Data)
}
