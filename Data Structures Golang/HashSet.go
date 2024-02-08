package main

import (
	"fmt"
)

const ArraySize int = 5

// HashSet
type HashSet struct {
	arr [ArraySize]*Bucket
}

// Bucket
type Bucket struct {
	bucketnode *BucketNode
}

// BucketNode
type BucketNode struct {
	Key  string
	next *BucketNode
}

func BuildHashSet() HashSet {
	myHashSet := HashSet{}
	for i := 0; i < ArraySize; i++ {
		myHashSet.arr[i] = &Bucket{}
	}
	return myHashSet
}
func Hash(s string) int {
	ans := 0
	for i := 0; i < len(s); i++ {
		ans += int(s[i])
	}
	return ans % ArraySize
}

// Insert
func (hashset *HashSet) Insert(s string) {
	hashValue := Hash(s)
	hashset.arr[hashValue].InsertBucket(s)
}

// Search
func (hashset *HashSet) Search(s string) bool {
	hashValue := Hash(s)
	return hashset.arr[hashValue].SearchBucket(s)
}

// Delete
func (hashset *HashSet) Delete(s string) {
	hashValue := Hash(s)
	hashset.arr[hashValue].DeleteBucket(s)
}

// insertbucket
func (bucket *Bucket) InsertBucket(s string) {
	newBucketNode := &BucketNode{Key: s, next: nil}
	newBucketNode.next = bucket.bucketnode
	bucket.bucketnode = newBucketNode
}

// searchbucket
func (bucket *Bucket) SearchBucket(s string) bool {
	curBucketNode := bucket.bucketnode
	for curBucketNode != nil {
		if curBucketNode.Key == s {
			return true
		}
		curBucketNode = curBucketNode.next
	}
	return false
}

// deletebucket
func (bucket *Bucket) DeleteBucket(s string) {
	prevBucketNode := bucket.bucketnode
	if prevBucketNode != nil && prevBucketNode.Key == s {
		bucket.bucketnode = prevBucketNode.next
		return
	}
	curBucketNode := prevBucketNode.next
	for curBucketNode != nil {
		if curBucketNode.Key == s {
			prevBucketNode.next = curBucketNode.next
			return
		}
		prevBucketNode = curBucketNode
		curBucketNode = curBucketNode.next
	}
}
func main() {
	hashSet := BuildHashSet()
	fmt.Println(hashSet)
	input := []string{
		"Github",
		"Golang",
		"Data",
		"Profile",
		"Share",
		"Tools",
		"Coding",
	}

	for i := 0; i < len(input); i++ {
		hashSet.Insert(input[i])
	}
	hashSet.Delete("Golang")
	hashSet.Delete("Coding")
	fmt.Println(hashSet.Search("Golang"))
	fmt.Println(hashSet.Search("Coding"))
	fmt.Println(hashSet.Search("Samartha"))
}
