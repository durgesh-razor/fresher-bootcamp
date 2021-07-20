package main

import "fmt"

type Node struct {
	val string
	left *Node
	right *Node
}

func inorder(root *Node){
	if root.left != nil{
		inorder(root.left)
	}
	fmt.Printf("%s ", root.val)
	if root.right != nil {
		inorder(root.right)
	}
}
func main(){
	tree := [5]Node{}
	tree[0].val = "+"

	tree[1].val = "a"
	tree[2].val = "-"

	tree[0].left = &tree[1]
	tree[0].right = &tree[2]

	tree[3].val = "b"
	tree[4].val = "c"
	tree[2].left = &tree[3]
	tree[2].right = &tree[4]

	inorder(&tree[0])
}