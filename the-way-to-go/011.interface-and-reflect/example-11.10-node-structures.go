package main

import (
    "fmt"
    "encoding/json"
)

type Node struct {
    Le *Node
    Data interface{}
    Ri *Node
}

func NewNode(left, right *Node) *Node {
    return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
    n.Data = data
}

func main() {
    root := NewNode(nil, nil)
    root.SetData("root node")
    // make child (leaf nodes)
    rootLeft := NewNode(nil, nil)
    rootLeft.SetData("left node")
    rootRight := NewNode(nil, nil)
    rootRight.SetData("right node")
    root.Le = rootLeft
    root.Ri = rootRight

    fmt.Printf("%#v\n", root)
    if jsonStr, e := json.MarshalIndent(root, "", "  "); e == nil {
        fmt.Printf("%s\n", jsonStr)
    }
}