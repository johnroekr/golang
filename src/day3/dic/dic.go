package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strings"
	"sync"
	"time"
)

func typeof(v interface{}) string {
	return fmt.Sprintf("%T", v)
}

func main() {
	start := time.Now()

	// dir, err := os.Getwd()
	// fmt.Println(dir)

	/* read a file*/
	f, err := ioutil.ReadFile("C:\\GoApp\\src\\day2\\dic\\test_word.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		// fmt.Println("file len :", len(f))
	}

	var s string = string(f)
	var words []string

	/* remove non-Alphanumeric */
	// s = strings.ReplaceAll(s, "\"", "")
	// s = strings.ReplaceAll(s, "\r, "")
	reg, err := regexp.Compile("[^a-zA-Z0-9 ]+")
	if err != nil {
		log.Fatal(err)
	}
	s2 := reg.ReplaceAllString(s, "")
	words = strings.Split(s2, " ")
	// fmt.Println(len(words))

	/* build a map (word, count) */
	word_cnt := make(map[string]int)

	var wait sync.WaitGroup
	var mutex = new(sync.Mutex)

	const THREAD_NUM int = 3
	thread_unit := len(words) / THREAD_NUM
	// fmt.Println("base", thread_unit)

	for i := 0; i < THREAD_NUM; i++ {
		wait.Add(1)

		go func(n int) {
			defer func() {
				wait.Done()
				fmt.Printf("[*] Thread %d : %d ~ %d done\n", n, thread_unit*n, thread_unit*(n+1))
			}()

			for i := thread_unit * n; i < thread_unit*(n+1); i++ {
				/* w/o mutex error could occur since go dynamically allocates memory */
				// fmt.Printf("[thread %d] words[%d]\n", n, i)
				mutex.Lock()
				word_cnt[words[i]]++
				mutex.Unlock()
			}
		}(i)
	}
	// fmt.Println(word_cnt)
	wait.Wait()

	/* build a map (count, []word) */
	cnt_word := map[int][]string{}
	for k, v := range word_cnt {
		cnt_word[v] = append(cnt_word[v], k)
	}
	// fmt.Println("cnt_word", cnt_word)

	/* sorting */
	var counts []int

	for k := range cnt_word {
		counts = append(counts, k)
	}

	sort.Sort(sort.Reverse(sort.IntSlice(counts)))

	/* extract top 5 the most frequent words */
	i := 5
	for _, cnt := range counts {
		for _, word := range cnt_word[cnt] {
			fmt.Printf("%s : %d\n", word, cnt)
		}
		i--
		if i == 0 {
			break
		}
	}

	/* keyword count */
	// var input string

	// fmt.Print("[*] Enter a word to count: ")
	// fmt.Scanln(&input)
	// fmt.Println(word_cnt[input])
	elapsed := time.Since(start)
	log.Printf("Binomial took %s", elapsed)
}
