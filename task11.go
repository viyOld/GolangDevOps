//What is an anagram? Well, two words are anagrams of each other if they both contain the same letters. For example:
//'abba' & 'baab' == true
//'abba' & 'bbaa' == true
//'abba' & 'abbba' == false
//'abba' & 'abca' == false
//Write a function that will find all the anagrams of a word from a list. You will be given two inputs a word and an array with words.
//You should return an array of all the anagrams or an empty array if there are none. For example:
//anagrams('abba', ['aabb', 'abcd', 'bbaa', 'dada']) => ['aabb', 'bbaa']
//anagrams('racer', ['crazer', 'carer', 'racar', 'caers', 'racer']) => ['carer', 'racer']
//anagrams('laser', ['lazing', 'lazy',  'lacer']) => []

package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println(Anagrams("abba", []string{"aabb", "abcd", "bbaa", "dada"})) //[]string{"aabb", "bbaa"}
	fmt.Println(Anagrams("laser", []string{"lazing", "lazy", "lacer"}))     //, nil)
}

func Anagrams(word string, words []string) []string {
	var res []string
	wordArr := []rune(word)
	wordMap := make(map[rune]int)
	for _, val := range wordArr {
		wordMap[val] += 1
	}
	for i, val := range words {
		optArr := []rune(val)
		optMap := make(map[rune]int)
		for _, vv := range optArr {
			optMap[vv] += 1
		}
		if reflect.DeepEqual(wordMap, optMap) == true {
			//fmt.Println(i)
			res = append(res, words[i])
		}

	}
	return res
}
