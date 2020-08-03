package trie

import (
	"container/heap"
	"encoding/json"
	"sync"
)

var instance *Trie
var once sync.Once

const topKNumber = 5

func MultiStringSearch(bigString string, smallStrings []*TextItem) []bool {
	// Write your code here.
	trie := GetInstance()

	b, e := json.MarshalIndent(trie.topKItems, "", " ")
	if e == nil {
		print(string(b))
	}

	return nil
}

type TextItem struct {
	Value  string
	Weight int
}

type Trie struct {
	char      byte
	children  map[byte]*Trie
	topKItems []*TextItem
}

func GetInstance() *Trie {
	once.Do(func() {
		instance = &Trie{
			children:  map[byte]*Trie{},
			topKItems: []*TextItem{},
		}
	})
	return instance
}

func (t *Trie) Add(textItem *TextItem) {
	trieStack := []*Trie{}

	current := t
	for i := range textItem.Value {
		trieStack = append(trieStack, current)

		letter := textItem.Value[i]
		if _, found := (*current).children[letter]; !found {
			(*current).children[letter] = &Trie{
				char:      letter,
				children:  map[byte]*Trie{},
				topKItems: []*TextItem{},
			}
		}

		trie := (*current).children[letter]
		current = trie
	}

	trieStack = append(trieStack, current)
	(*current).children['*'] = &Trie{
		topKItems: []*TextItem{textItem},
	}

	for i := len(trieStack) - 1; i >= 0; i-- {
		trie := trieStack[i]

		topKItemsCollection := [][]*TextItem{}

		for _, v := range (*trie).children {
			topKItemsCollection = append(topKItemsCollection, v.topKItems)
		}

		trie.topKItems = t.getTopKItems(topKItemsCollection, topKNumber)
	}
}

func (t *Trie) Search(text string) []*TextItem {
	node := t

	for i := 0; i < len(text); i++ {
		char := text[i]
		val, found := node.children[char]

		if !found {
			return []*TextItem{}
		}

		node = val
	}

	return node.topKItems
}

func (t Trie) getTopKItems(topKItemsCollection [][]*TextItem, k int) []*TextItem {
	pq := priorityQueue(topKItemsCollection)
	heap.Init(&pq)

	result := []*TextItem{}

	for k > 0 {
		textItems := heap.Pop(&pq).([]*TextItem)
		if len(textItems) == 0 {
			break
		}
		textItem := textItems[0]
		textItems = textItems[1:]
		heap.Push(&pq, textItems)
		result = append(result, textItem)
		k--
	}

	return result
}

type priorityQueue [][]*TextItem

func (pq priorityQueue) Len() int { return len(pq) }
func (pq priorityQueue) Less(i, j int) bool {
	len1 := len(pq[i])
	len2 := len(pq[j])
	if len1 == 0 || len2 == 0 {
		return len1 > len2
	}

	itemI, itemJ := pq[i][0], pq[j][0]
	return itemI.Weight > itemJ.Weight
}
func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	items := x.([]*TextItem)
	*pq = append(*pq, items)
}

func (pq *priorityQueue) Pop() interface{} {
	length := len(*pq)
	items := (*pq)[length-1]
	*pq = (*pq)[:length-1]
	return items
}
