package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main () { 
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("failed to read input.txt: %s", err)
	}
	fullNum := 0 
	for i, line := range lines {
		fmt.Println(i, line)
		// remove digit words left corner parser
		cleanedLineLeft := replaceStringDigitLeft(line)
		leftString := getTrebuchetValue(cleanedLineLeft)
		// remove digit words right corner parser 
		cleanedLineRight := replaceStringDigitRight(line)
		rightString := getTrebuchetValue(reverse(cleanedLineRight))
		fmt.Println(i, cleanedLineLeft, cleanedLineRight)	
		num := leftString + rightString
		lineNum, err := strconv.Atoi(num)
		if err != nil {
			log.Fatalf("failed to convert string to int: %s", err)
		}
		fullNum += lineNum
		fmt.Println(fullNum, lineNum)
	}
}

// function, which takes a string as 
// argument and return the reverse of string. 
func reverse(s string) string { 
    rns := []rune(s) // convert to rune 
    for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 { 
  
        // swap the letters of the string, 
        // like first with last and so on. 
        rns[i], rns[j] = rns[j], rns[i] 
    } 
  
    // return the reversed string. 
    return string(rns) 
} 

func replaceStringDigitLeft(inputString string) string {
	result := inputString
	words := map[string]string{ "two":"2", "one": "1", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	i := 0
	for true {
		if i >= len(result) {
			break
		}
		for word, num := range words {
			if i+len(word) > len(result) {
				continue
			}
			substr := result[i:i+len(word)]
			if word == substr {
				result = strings.Replace(result, word, num, 1)
				break
			}
		}
		i += 1
	}
	return result
}

func replaceStringDigitRight(inputString string) string {
	result := inputString
	words := map[string]string{ "two":"2", "one": "1", "three": "3", "four": "4", "five": "5", "six": "6", "seven": "7", "eight": "8", "nine": "9"}
	i := len(result)-1
	for true {
		if i == 0 {
			break
		}
		for word, num := range words {
			if i+len(word) > len(result) {
				continue
			}
			substr := result[i:i+len(word)]
			if word == substr {
				result = strings.Replace(result, word, num, 1)
				break
			}
		}
		i -= 1
	}
	return result
}

func getTrebuchetValue(inputString string) (string) {
	left := 0
	leftDigit := 0
	for true {
		if leftDigit != 0 {
			break
		}
		// check left index then check right for if we found the left and right
		if leftDigit == 0 {
			leftD, _ := getDigit(inputString, left)
			left += 1
			leftDigit = leftD
		}
	}
	return strconv.Itoa(leftDigit)
}

//checks if the index in the string is a diget if it is it returns it
func getDigit(inputString string, index int) (int, error) {
	pointToCheck := inputString[index]
	strPoint := string(pointToCheck)
	result, err := strconv.Atoi(strPoint)
	if err != nil {
		return 0, err
	}
	return result, nil
}

// readLines reads a whole file into memory
// and returns a slice of its lines.
func readLines(path string) ([]string, error) {
    file, err := os.Open(path)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    var lines []string
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }
    return lines, scanner.Err()
}