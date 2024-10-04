package main

import (
	"chat-raft/pkg/node"
	"fmt"
)

func main() {
	fmt.Println("chatNode server init...")
	var nodeNum uint8 = node.Initiliaze()
	fmt.Printf("Node number is: %v\n", nodeNum)

}
