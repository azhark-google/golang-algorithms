package main

import (
	"fmt"
)

type Tree struct {
root *Node
}

type Node struct {
  Left *Node
  Right *Node
  Val int
}

var values = make([]int, 0)

 func traverse(t *Node) []int {
    if t.Right != nil {
       traverse(t.Right)
    }  
    values = append(values, t.Val)
    if t.Left != nil {
       traverse(t.Left)
    }  
  return values
}

func(t *Tree) insert(val int) {
  var NodeType string

  if t.root == nil {
    t.root = &Node{Left: nil, Right: nil, Val: val}
    return
  }

  root := t.root
  child := root
  
  for (child != nil)   {
    
    if (val > child.Val) {
      root = child
      child = root.Right
      NodeType = "Right"
    } else {
      root = child
      child = root.Left
      NodeType = "Left"
    }
        
  }

  if NodeType == "Right" {
    root.Right=&Node{Left: nil, Right: nil, Val: val}
//    fmt.Println(root.Right.Val)
    return 
  } else if NodeType == "Left" {
    root.Left=&Node{Left: nil, Right: nil, Val: val}
//    fmt.Println(root.Left.Val)
    return
  } 

}

func (t *Tree) delete (val int) string{
  node := t.root

// if there is single node, delete that node
  if node.Val == val && node.Left == nil && node.Right == nil {
    node = nil
//    return fmt.Sprintf("deleted", val)
  } 

  parent := &Node{}
  nodeType := ""
  if node.Val == val  && (node.Left != nil || node.Right != nil) {
     nodeType = "Root"
  }
// traverse until val is found
  
  for val != node.Val {
    if val > node.Val {
      parent = node
      node = node.Right
      nodeType = "Right"
    } else if val < node.Val {
      parent = node
      node = node.Left
      nodeType = "Left"
    } else if val == node.Val {
      break;
    }

//    return "Not found"
  }

// find inorder left or right predecessor of the node
// inorder left predecessor -> right most child of the parent's left node
// inorder right predecessor -> left most child of the parent's right node

// if  node is a leaf node, delete that element
  if node.Left == nil && node.Right == nil {
    if nodeType == "Left" {
       parent.Left = nil
    } else {
      parent.Right = nil
    }
   return fmt.Sprintf("deleted %v", val)
  }  

// node left child then traverse until left predecessor is found  
  if (node.Right == nil && node.Left != nil) || (node.Right != nil && node.Left != nil)  {
    leftNode := node.Left
    RightmostParent := leftNode
    Rightmost := leftNode.Right
    if Rightmost == nil {
      if nodeType == "Left" {
        parent.Left = leftNode
        node=nil
      } else if nodeType == "Right" {
        parent.Right = leftNode
        node=nil
      } else if nodeType == "Root" {
        node = leftNode
      }
    return fmt.Sprintf("deleted %v", val)
    }
    for Rightmost.Right != nil {
      RightmostParent = Rightmost
      Rightmost = Rightmost.Right
    }
    
    if Rightmost.Left != nil {
      RightmostParent.Right=Rightmost.Left
    } else {
      RightmostParent.Right=nil
    }
    if nodeType == "Left" {
      parent.Left = Rightmost
      Rightmost.Left=node.Left
      node=nil
    } else if nodeType == "Right" {
      parent.Right = Rightmost
      Rightmost.Left=node.Left
      node=nil
    } else if nodeType == "Root" {
      Rightmost.Left=node.Left
      Rightmost.Right=node.Right
      t.root=Rightmost
    }
    fmt.Println(Rightmost.Val)
    
    return fmt.Sprintf("deleted %v", val)
  }

if node.Left == nil && node.Right != nil {
    RightNode := node.Right
    LeftmostParent := RightNode
    Leftmost := RightNode.Left
    if Leftmost == nil {
      if nodeType == "Left" {
         parent.Left = RightNode
      } else {
        parent.Right = RightNode
      }
    node=nil
    return fmt.Sprintf("deleted %v", val)
    }
    for Leftmost.Left != nil {
      LeftmostParent = Leftmost
      Leftmost = Leftmost.Left
    }
    if Leftmost.Right != nil {
      LeftmostParent.Left=Leftmost.Right
    } else {
      LeftmostParent.Left=nil
    }
    if nodeType == "Left" {
       parent.Left = Leftmost
    } else {
      parent.Right = Leftmost
    }
    Leftmost.Right=node.Right
    node=nil
  return fmt.Sprintf("deleted %v", val)
  }

//---> go left
//---> go right....most
//---->replace rightmostParent.right - rightmost.left
//---->parent=rightmost

return fmt.Sprint("not found")
}

//      10
// 8          14
//5 9     12        19
//   7  11  13    17   20
// 6            15                 
//4                16


//10
//     18
//  12    20
//11   13
//       17
// descleft -> leftChild-->rightMost
// descLeft = leftChildNode

func main () {
t := &Tree{}
t.insert(10)
t.insert(15)
t.insert(5)
t.insert(18)
t.insert(12)
t.insert(8)
t.insert(9)
fmt.Println(traverse(t.root))
t.delete(10)
values=nil
fmt.Println(traverse(t.root))
t.delete(12)

values=nil
fmt.Println(traverse(t.root))
t.delete(9)

values=nil
fmt.Println(traverse(t.root))
t.delete(15)

values=nil
fmt.Println(traverse(t.root))


}

