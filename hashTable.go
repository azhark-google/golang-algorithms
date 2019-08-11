// Hash table implementation in golang
// Using linkedlist for handling hash collisions

package main

import (
	"fmt"
	"strconv"
)

type hashTable struct {
  keyStore  []*keyVal
  bucketNum int  

}

type dataType interface {

}

type keyVal struct {
  Key dataType
  Val dataType
  Next *keyVal
}


// function takes the total of ascii of string chars
// divide it by bucketNum
	
func hashKey (key dataType, bucketNum int) int {
  asciiTotal := 0
  hashKey := ""
  switch k := key.(type) {
    case int:
      hashKey = strconv.Itoa(k)
      case string:
        hashKey = k
      default:
        fmt.Printf("Invalid datatype used for hash key")
        return 0
  }


  for _,s:= range hashKey {
    asciiTotal =+ int(s)
  }
  return (asciiTotal + len(hashKey))%bucketNum
  
}

func(ht *hashTable) insertHashTab(key, val dataType) {
  if (ht.keyStore == nil ) {
    kv := &keyVal{Key:key, Val:val}
    hash := hashKey(key, ht.bucketNum)
    ht.keyStore = make([]*keyVal, ht.bucketNum)
    ht.keyStore[hash] = kv
    return
  }

  hash := hashKey(key, ht.bucketNum)

  kv := ht.keyStore[hash]
  if kv==nil {
    ht.keyStore[hash] = &keyVal{Key:key, Val:val}
    return
  }
  current := kv
  // If there is hash collision use linked list to store next key,val
  for kv != nil {
    if kv.Key != key {
      current = kv
      kv = kv.Next
      continue
    }
    kv.Val=val
    return
  }
  
  kv=&keyVal{Key:key, Val:val}
  current.Next = kv
  
}

func(ht *hashTable) searchHashTab(key dataType) (dataType) {
  hash := hashKey(key, ht.bucketNum)
  kv := ht.keyStore[hash]
  for kv != nil {
    if kv.Key == key {
        return kv.Val
    }
    kv = kv.Next
  }
  return "None"
}

func (ht *hashTable) deleteHashTab(key dataType) {
  hash := hashKey(key, ht.bucketNum)
  kv := ht.keyStore[hash]
  current := kv
  for kv != nil {
    
    if kv.Key == key {
       if kv.Next == nil {
         current.Next = nil
         kv = nil
         return
       }
       current.Next = kv.Next
       kv = nil
       return
    }
    kv = kv.Next 
  }    
  return
}


func main() {
  ht := &hashTable{bucketNum: 10}
  ht.insertHashTab("query01", "13-July-2019")
  ht.insertHashTab("query02", "14-July-2019")
  ht.insertHashTab("query01", "15-July-2019")
  ht.insertHashTab("query03", "16-July-2019")
  ht.insertHashTab("query04", "17-July-2019")
  ht.insertHashTab("query05", "18-July-2019")
  ht.insertHashTab("query06", "19-July-2019")
  ht.insertHashTab("query07", "20-July-2019")
  ht.insertHashTab("query08", "21-July-2019")
  ht.insertHashTab("query09", "22-July-2019")
  ht.insertHashTab("query10", "23-July-2019")
  ht.insertHashTab("query11", "24-July-2019")
  ht.insertHashTab("query11", "25-July-2019")
  

  fmt.Println(ht)

  fmt.Println(ht.searchHashTab("query11"))
  ht.deleteHashTab("query11")
  fmt.Println(ht.searchHashTab("query11"))
  
}
