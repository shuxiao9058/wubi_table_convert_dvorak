package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strings"
)

func readLines(path string) ([]string, error) {
	var lines []string
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func convertLine(line string) (string, error) {
	splits := strings.Split(line, " ")
	newSplits := []string{}
	for _, s := range splits {
		s = strings.TrimSpace(s)
		if s != "" {
			newSplits = append(newSplits, s)
		}
	}

	if len(newSplits) < 2 {
		return "", errors.New("qqq")
	}

	// out := ""
	// for i := 1; i < len(newSplits); i++ {
	//	// out = out + fmt.Sprintf("%s\t%s\n", newSplits[i], tableConvert(newSplits[0]))
	//	// out = out + fmt.Sprintf("%s\t%s\n", newSplits[i], newSplits[0])
	//	// out = out + fmt.Sprintf("%s\t%s\n", newSplits[i], newSplits[0])
	//	out = out + fmt.Sprintf("%s\t%s\n", newSplits[i], tableConvert(newSplits[0]))
	// }

	// return out, nil

	// 0是中文；1,3是编码；2是频度
	// if len(newSplits) < 3 {
	//	log.Fatalf("len(newSplits) is %d", len(newSplits))
	//	return "", errors.New("error len of newSplits")
	// }

	// if len(newSplits) > 3 {
	//	return fmt.Sprintf("%-10s%-10s%-10s%-10s",
	//		newSplits[0], tableConvert(newSplits[1]), newSplits[2], tableConvert(newSplits[3])), nil
	// } else {
	//	return fmt.Sprintf("%-10s%-10s%-10s",
	//		newSplits[0], tableConvert(newSplits[1]), newSplits[2]), nil
	// }

	// if len(newSplits) > 3 {
	//	return fmt.Sprintf("%s %s %s %s",
	//		newSplits[0], tableConvert(newSplits[1]), newSplits[2], tableConvert(newSplits[3])), nil
	// } else {
	//	return fmt.Sprintf("%s %s %s",
	//		newSplits[0], tableConvert(newSplits[1]), newSplits[2]), nil
	// }

	if len(newSplits) > 3 {
		return fmt.Sprintf("%s\t%s\t%s\t%s\n",
			newSplits[0], tableConvert(newSplits[1]), newSplits[2], tableConvert(newSplits[3])), nil
	} else {
		return fmt.Sprintf("%s\t%s\t%s\n",
			newSplits[0], tableConvert(newSplits[1]), newSplits[2]), nil
	}
}

// var tableMap = map[string]string{}
var tableMap = map[string]string{
	// 一区
	"g": "i",
	"f": "u",
	"d": "e",
	"s": "o",
	"a": "a",

	// 二区
	"h": "d",
	"j": "h",
	"k": "t",
	"l": "n",
	"m": "s",

	// 三区
	"t": "x",
	"r": "k",
	"e": "j",
	"w": "q",
	"q": "p",

	// // 四区
	// "y": "f",
	// "u": "g",
	// "i": "c",
	// "o": "r",
	// "p": "l",

	// // 五区
	// "n": "b",
	// "b": "m",
	// "v": "w",
	// "c": "v", // v
	// "x": "z",

	// 四区
	"y": "b",
	"u": "m",
	"i": "w",
	"o": "v",
	"p": "z",

	// 五区
	"n": "f",
	"b": "g",
	"v": "c",
	"c": "r", // v
	"x": "l",

	// 学习键
	"z": "y",
}

func tableConvert(code string) string {
	var out string
	for i, _ := range code {
		out += tableMap[string(code[i])]
	}

	return out
}

func main() {
	f, err := os.OpenFile("out.txt", os.O_RDWR|os.O_CREATE, 0777)
	if err != nil {
		log.Fatal(err.Error())
		return
	}

	// dicts, err := readLines("1.06c_16016.txt")
	dicts, err := readLines("wubi06.txt")
	if err != nil {
		log.Println("error load dicts")
		return
	}

	for _, line := range dicts {
		out, err := convertLine(line)
		if err != nil {
		} else {
			// log.Println(out)
			// f.Write(b []byte)
			f.WriteString(out)
			// f.Write([]byte("\n"))
		}
	}

	f.Close()
	// log.Printf("len(dicts) is: %d", len(dicts))
}
