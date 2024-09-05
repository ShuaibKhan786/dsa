// Implement a time-based key-value data structure that supports:

// Storing multiple values for the same key at specified time stamps
// Retrieving the key's value at a specified timestamp
// Implement the TimeMap class:

// TimeMap() Initializes the object.
// void set(String key, String value, int timestamp) Stores the key key with the value value at the given time timestamp.
// String get(String key, int timestamp) Returns the most recent value of key if set was previously called on it and the most recent timestamp for that key prev_timestamp is less than or equal to the given timestamp (prev_timestamp <= timestamp). If there are no values, it returns "".
// Note: For all calls to set, the timestamps are in strictly increasing order.

// Example 1:
// Input:
// ["TimeMap", "set", ["alice", "happy", 1], "get", ["alice", 1], "get", ["alice", 2], "set", ["alice", "sad", 3], "get", ["alice", 3]]

// Output:
// [null, null, "happy", "happy", null, "sad"]

// Explanation:
// TimeMap timeMap = new TimeMap();
// timeMap.set("alice", "happy", 1);  // store the key "alice" and value "happy" along with timestamp = 1.
// timeMap.get("alice", 1);           // return "happy"
// timeMap.get("alice", 2);           // return "happy", there is no value stored for timestamp 2, thus we return the value at timestamp 1.
// timeMap.set("alice", "sad", 3);    // store the key "alice" and value "sad" along with timestamp = 3.
// timeMap.get("alice", 3);           // return "sad"
// 
// Constraints:

// 1 <= key.length, value.length <= 100
// key and value only include lowercase English letters and digits.
// 1 <= timestamp <= 1000

package main

import "fmt"

type Pair struct {
	Value     string
	TimeStamp int
}

type TimeMap struct {
	Store map[string][]Pair
}

func new() *TimeMap {
	return &TimeMap{
		make(map[string][]Pair),
	}
}

func (t *TimeMap) set(key, value string, timestamp int) {
	t.Store[key] = append(t.Store[key], Pair{Value: value, TimeStamp: timestamp})
}

func (t *TimeMap) get(key string, timestamp int) string {
	s, ok := t.Store[key]
	if !ok {
		return ""
	}

	left := 0
	right := len(s) - 1

	for left <= right {
		mid := (left + right) / 2

		switch {
		case timestamp < s[mid].TimeStamp:
			right = mid - 1
		case timestamp > s[mid].TimeStamp:
			left = mid + 1
		default:
			return s[mid].Value
		}
	}

	if right >= 0 {
		return s[right].Value
	}

	return ""
}

func main() {
	obj := new()
	obj.set("alice", "happy", 1)
	first := obj.get("alice", 1) //happy
	if first != "happy" {
		fmt.Println("Expected: happy but Got: %v", first)
		return
	}
	fmt.Println(first)
	second := obj.get("alice", 2) //happy
	if second != "happy" {
		fmt.Println("Expected: happy but Got: %v", second)
		return
	}
	fmt.Println(second)
	obj.set("alice", "sad", 3)
	third := obj.get("alice", 2) //happy
	if third != "happy" {
		fmt.Println("Expected: happy but Got: %v", third)
		return
	}
	fmt.Println(third)
	fourth := obj.get("alice", 3) // sad
	if fourth != "sad" {
		fmt.Println("Expected: happy but Got: %v", fourth)
		return
	}
	fmt.Println(fourth)
	fmt.Println("Test Passed")
}
