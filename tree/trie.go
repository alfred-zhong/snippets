package tree

type TrieNode struct {
	v        rune
	wFlag    bool
	children []*TrieNode
}

func (tn *TrieNode) IsLeaf() bool {
	return len(tn.children) == 0
}

func (tn *TrieNode) AddWord(word string) {
	runes := []rune(word)

	var nextNode *TrieNode
	var n int
	for ; n < len(tn.children); n++ {
		if tn.children[n].v == runes[0] {
			nextNode = tn.children[n]
			break
		} else if tn.children[n].v > runes[0] {
			break
		}
	}

	if nextNode == nil {
		nextNode = &TrieNode{
			v: runes[0],
		}

		var tmpChildren []*TrieNode
		tmpChildren = append(tmpChildren, tn.children[:n]...)
		tmpChildren = append(tmpChildren, nextNode)
		tmpChildren = append(tmpChildren, tn.children[n:]...)
		tn.children = tmpChildren
	}

	if len(runes) == 1 {
		nextNode.wFlag = true
	} else if len(runes) > 1 {
		nextNode.AddWord(string(runes[1:]))
	}
}

func (tn *TrieNode) FindWord(word string) bool {
	runes := []rune(word)

	var nextNode *TrieNode
	for n := 0; n < len(tn.children); n++ {
		if tn.children[n].v == runes[0] {
			nextNode = tn.children[n]
			break
		} else if tn.children[n].v > runes[0] {
			return false
		}
	}

	if nextNode == nil {
		return false
	}

	if len(runes) == 1 {
		return nextNode.wFlag
	}
	return nextNode.FindWord(string(runes[1:]))
}

func NewTrieNode() *TrieNode {
	return &TrieNode{
		v: rune('/'),
	}
}
