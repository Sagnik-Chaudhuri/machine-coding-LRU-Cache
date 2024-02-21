package service

import (
	"errors"
	"log"
	"machine-coding-lru-cache/internal/model"
	"sync"
)

type LRUCacheServiceImpl struct {
	KVPair   map[string]*model.Node
	Head     *model.Node
	Tail     *model.Node
	Capacity int
}

var cacheServiceInstance CacheService
var cacheServiceOnce sync.Once

func GetCacheService(capacity int) CacheService {
	log.Println("init GetCacheService")
	cacheServiceOnce.Do(func() {
		cacheServiceInstance = &LRUCacheServiceImpl{
			KVPair:   map[string]*model.Node{},
			Head:     nil,
			Tail:     nil,
			Capacity: capacity,
		}
	})
	return cacheServiceInstance
}

// GetCacheServiceSimplifiedInstance without sync.Once call
func GetCacheServiceSimplifiedInstance(capacity int) CacheService {
	log.Println("init GetCacheServiceSimplifiedInstance")
	return &LRUCacheServiceImpl{
		KVPair:   map[string]*model.Node{},
		Head:     nil,
		Tail:     nil,
		Capacity: capacity,
	}
}

func (l *LRUCacheServiceImpl) addNodeToHead(newNode *model.Node) {
	existingHead := l.Head
	if existingHead == nil {
		l.Head = newNode
		l.Tail = newNode
		newNode.Next = nil
		newNode.Prev = nil
	} else {
		newNode.Next = existingHead
		newNode.Prev = nil
		existingHead.Prev = newNode
		l.Head = newNode
	}
	return
}

func (l *LRUCacheServiceImpl) deleteNodeFromList(delNode *model.Node) {
	if l.Head == nil {
		return // List is empty, nothing to delete
	}

	if delNode == nil {
		return // Node to delete is nil, nothing to delete
	}

	// If the node to delete is the head
	if delNode == l.Head {
		l.Head = delNode.Next
		if l.Head != nil {
			l.Head.Prev = nil
		} else {
			// The list is now empty
			l.Tail = nil
		}
		return
	}

	// If the node to delete is the tail
	if delNode == l.Tail {
		l.Tail = delNode.Prev
		if l.Tail != nil {
			l.Tail.Next = nil
		} else {
			// The list is now empty
			l.Head = nil
		}
		return
	}

	// If the node to delete is in the middle
	if delNode.Prev != nil {
		delNode.Prev.Next = delNode.Next
	}
	if delNode.Next != nil {
		delNode.Next.Prev = delNode.Prev
	}
}

func (l *LRUCacheServiceImpl) GetValue(key string) (string, error) {
	if node, ok := l.KVPair[key]; !ok {
		return "", errors.New("value not found in cache")
	} else {
		l.deleteNodeFromList(node)
		l.addNodeToHead(node)
		return node.Val, nil
	}
}

func (l *LRUCacheServiceImpl) PutValue(key string, value string) {
	if node, ok := l.KVPair[key]; ok {
		node.Val = value
		l.deleteNodeFromList(node)
		l.addNodeToHead(node)
	} else {
		if len(l.KVPair) >= l.Capacity {
			tail := l.Tail
			delete(l.KVPair, tail.Key)
			l.deleteNodeFromList(tail)
		}
		newNode := &model.Node{
			Key:  key,
			Val:  value,
			Prev: nil,
			Next: nil,
		}
		l.addNodeToHead(newNode)
		l.KVPair[key] = newNode
	}
}
