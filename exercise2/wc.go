package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// strings are composed of runes
	vowels := make(map[rune]int)
	consonants := make(map[rune]int)
	total_chars := 0

	vowels_set := map[string]struct{}{}
	vowels_set["a"] = struct{}{}
	vowels_set["e"] = struct{}{}
	vowels_set["i"] = struct{}{}
	vowels_set["o"] = struct{}{}
	vowels_set["u"] = struct{}{}

	for scanner.Scan() {
		text := scanner.Text()
		text = strings.ToLower(text)
		// iterate over every rune of text

		for _, c := range text {

			_, has_vowel := vowels_set[string(c)]
			if has_vowel {
				vowels[c]++
			} else {
				consonants[c]++
			}
			total_chars++
		}

	}

	fmt.Println("statistics:")

	max_vowel_k := "a"
	max_value := 0
	min_vowel_k := "a"
	min_vowel_value := total_chars

	for k, v := range vowels {
		if v >= max_value {
			max_value = v
			max_vowel_k = string(k)
		}
		if v <= min_vowel_value {
			min_vowel_value = v
			min_vowel_k = string(k)
		}
		fmt.Printf("vowel %c appears %d times\n", k, v)
	}
	fmt.Printf("Total characters: %d\n", total_chars)

	max_consonant_k := "a"
	max_consonant_value := 0
	min_consonant_k := "a"
	min_consonant_value := total_chars

	for k, v := range consonants {
		if v >= max_consonant_value {
			max_consonant_value = v
			max_consonant_k = string(k)
		}
		if v <= min_consonant_value {
			min_consonant_value = v
			min_consonant_k = string(k)
		}
		fmt.Printf("consonant %c appears %d times\n", k, v)
	}

	if max_consonant_value > max_value {
		fmt.Printf("Max appereance %c appears %d times\n", max_consonant_k, max_consonant_value)
	} else {
		fmt.Printf("Max appereance %c appears %d times\n", max_vowel_k, max_value)
	}

	if min_consonant_value > min_vowel_value {
		fmt.Printf("Min appereance %s appears %d times\n", string(min_consonant_k), min_consonant_value)
	} else {
		fmt.Printf("Min appereance %s appears %d times\n", string(min_vowel_k), min_vowel_value)
	}

}
