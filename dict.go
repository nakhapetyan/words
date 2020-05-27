package main

import "github.com/buger/jsonparser"

type Dict struct {
	Words        map[string]uint64
	lettersFlags map[rune]uint64
}

func NewDict(words []string) *Dict {
	dict := new(Dict)
	dict.lettersFlags = map[rune]uint64{
		'a': 1, 'b': 4, 'c': 16, 'd': 64, 'e': 256, 'f': 1024, 'g': 4096, 'h': 16384, 'i': 65536, 'j': 262144, 'k': 1048576, 'l': 4194304, 'm': 16777216, 'n': 67108864, 'o': 268435456, 'p': 1073741824, 'q': 4294967296, 'r': 17179869184, 's': 68719476736, 't': 274877906944, 'u': 1099511627776, 'v': 4398046511104, 'w': 17592186044416, 'x': 70368744177664, 'y': 281474976710656, 'z': 1125899906842624,
		'A': 1, 'B': 4, 'C': 16, 'D': 64, 'E': 256, 'F': 1024, 'G': 4096, 'H': 16384, 'I': 65536, 'J': 262144, 'K': 1048576, 'L': 4194304, 'M': 16777216, 'N': 67108864, 'O': 268435456, 'P': 1073741824, 'Q': 4294967296, 'R': 17179869184, 'S': 68719476736, 'T': 274877906944, 'U': 1099511627776, 'V': 4398046511104, 'W': 17592186044416, 'X': 70368744177664, 'Y': 281474976710656, 'Z': 1125899906842624,
		'а': 1, 'б': 4, 'в': 16, 'г': 64, 'д': 256, 'е': 1024, 'ж': 4096, 'з': 16384, 'и': 65536, 'й': 262144, 'к': 1048576, 'л': 4194304, 'м': 16777216, 'н': 67108864, 'о': 268435456, 'п': 1073741824, 'р': 4294967296, 'с': 17179869184, 'т': 68719476736, 'у': 274877906944, 'ф': 1099511627776, 'х': 4398046511104, 'ц': 17592186044416, 'ч': 70368744177664, 'ш': 281474976710656, 'щ': 1125899906842624, 'ы': 4503599627370496, 'ь': 18014398509481984, 'э': 72057594037927936, 'ю': 288230376151711744, 'я': 1152921504606846976, 'ъ': 4611686018427387904,
		'А': 1, 'Б': 4, 'В': 16, 'Г': 64, 'Д': 256, 'Е': 1024, 'Ж': 4096, 'З': 16384, 'И': 65536, 'Й': 262144, 'К': 1048576, 'Л': 4194304, 'М': 16777216, 'Н': 67108864, 'О': 268435456, 'П': 1073741824, 'Р': 4294967296, 'С': 17179869184, 'Т': 68719476736, 'У': 274877906944, 'Ф': 1099511627776, 'Х': 4398046511104, 'Ц': 17592186044416, 'Ч': 70368744177664, 'Ш': 281474976710656, 'Щ': 1125899906842624, 'Ы': 4503599627370496, 'Ь': 18014398509481984, 'Э': 72057594037927936, 'Ю': 288230376151711744, 'Я': 1152921504606846976, 'Ъ': 4611686018427387904,
	}
	dict.loadWords(words)
	return dict
}

func (dict *Dict) getWordHash(word string) uint64 {
	a := []rune(word)
	var hash uint64
	for _, letter := range a {
		hash += dict.lettersFlags[letter]
	}

	return hash
}

func (dict *Dict) getNeedWordHash(word string) uint64 {
	a := []rune(word)
	var hash uint64
	for _, letter := range a {
		if (hash & dict.lettersFlags[letter]) == 0 {
			hash += dict.lettersFlags[letter]
		} else {
			hash += dict.lettersFlags[letter] * 2
		}
	}

	return hash
}

func (dict *Dict) parseWords(data []byte) []string {
	var words []string
	jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		words = append(words, string(value))
	})

	return words
}

func (dict *Dict) loadWords(words []string) {
	dict.Words = map[string]uint64{}
	for _, word := range words {
		dict.Words[word] = dict.getWordHash(word)
	}
}

func (dict *Dict) findWordsByLetters(needWord string) []string {
	var result []string
	needHash := dict.getNeedWordHash(needWord)

	for word, hash := range dict.Words {
		if (hash ^ (needHash & hash)) == 0 {
			result = append(result, word)
		}
	}

	return result
}
