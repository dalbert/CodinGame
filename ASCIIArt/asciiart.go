package main

import "fmt"
import "os"
import "bufio"
import "bytes"

//import "strings"
//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)

	var L int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &L)

	var H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &H)

	scanner.Scan()
	T := scanner.Text()
	var font [][]string
	for i := 0; i < H; i++ {
		scanner.Scan()
		ROW := scanner.Text()
		font = append(font, scanFontRow(ROW, L))
	}

	fmt.Println(compileString(T, font, L, H)) // Write answer to stdout
}

func compileString(input string, font [][]string, charWidth int, charHeight int) string {
	var output bytes.Buffer
	charMap := stringToIntArray(input)
	for h := 0; h < charHeight; h++ {
		for _, charInt := range charMap {
			output.WriteString(font[h][charInt])
		}
		output.WriteRune('\n')
	}
	return output.String()
}

func stringToIntArray(input string) []int {
	var charMap []int
	for _, char := range input {
		if rune(char) > 64 && rune(char) < 91 {
			charMap = append(charMap, int(rune(char)-65))
		} else if rune(char) > 96 && rune(char) < 123 {
			charMap = append(charMap, int(rune(char)-97))
		} else {
			charMap = append(charMap, 26)
		}
	}
	return charMap
}

func scanFontRow(row string, charWidth int) []string {
	var charRows []string
	for i := 0; i < 27; i++ {
		charRows = append(charRows, row[i*charWidth:i*charWidth+charWidth])
	}
	return charRows
}
