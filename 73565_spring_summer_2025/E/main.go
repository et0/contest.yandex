package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type found struct {
	index int
	words []string
}

func main() {
	var in *bufio.Reader
	var out *bufio.Writer
	in = bufio.NewReader(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
	defer out.Flush()

	t9 := map[rune]string{
		'A': "2",
		'B': "22",
		'C': "222",
		'D': "3",
		'E': "33",
		'F': "333",
		'G': "4",
		'H': "44",
		'I': "444",
		'J': "5",
		'K': "55",
		'L': "555",
		'M': "6",
		'N': "66",
		'O': "666",
		'P': "7",
		'Q': "77",
		'R': "777",
		'S': "7777",
		'T': "8",
		'U': "88",
		'V': "888",
		'W': "9",
		'X': "99",
		'Y': "999",
		'Z': "9999",
	}

	var secret string
	fmt.Fscan(in, &secret)
	sizeSecret := len(secret)

	var n int
	fmt.Fscan(in, &n)

	dict := make(map[string]string, n)
	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(in, &word)

		tmp := make([]string, 128)
		for _, w := range word {
			tmp = append(tmp, t9[w])
		}

		key := strings.Join(tmp, "")
		dict[key] = word
	}

	founds := make([]found, 0, n)
	founds = append(founds, found{0, make([]string, 0, n)})

	for len(founds) > 0 {
		sizeF := len(founds)
		f := founds[sizeF-1]
		founds = founds[0 : sizeF-1]

		for k, d := range dict {
			lenK := len(k)
			if f.index+lenK >= sizeSecret || k != secret[f.index:f.index+lenK] {
				continue
			}

			// если последнее найденное слово, достигло конца секретной строки
			if f.index+lenK == sizeSecret {
				// добавляем рассшифрованное слово
				f.words = append(f.words, d)

				// распечатываем все найденые слова
				fmt.Fprint(out, strings.Join(f.words, " "))

				// обнуляем слайс, что бы выйти из главного цилка
				founds = nil

				break
			}

			// добавляем новую проверку
			new := make([]string, 0, len(f.words)+1)
			new = append(new, f.words...)
			founds = append(founds, found{
				f.index + lenK,
				append(new, d),
			})
		}
	}
}
