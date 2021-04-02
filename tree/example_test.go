package tree

import "fmt"

func ExampleTrieNode_basic() {
	tn := NewTrieNode()
	tn.AddWord("cat")
	tn.AddWord("dog")

	fmt.Println(tn.FindWord("cat"))
	fmt.Println(tn.FindWord("dog"))
	fmt.Println(tn.FindWord("cate"))
	fmt.Println(tn.FindWord("food"))
	// Output:
	// true
	// true
	// false
	// false
}
