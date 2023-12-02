package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type game struct {
	id int
	draws []map[string]int
}

type TooManyInt struct{}

func (m *TooManyInt) Error() string {
	return "Attempted to parse to many ints make sure input string only has one int"
}

func main () { 
	const redCube = 12
	const greenCube = 13
	const blueCube = 14 

	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("Unable to load input file: %s", err)
	}

	var correctIdsSum int
	var powerSum int
	for _, line := range lines {
		game, err := parseGame(line)
		if err != nil {
			log.Fatalf("error parsing game: %s", err)
		}
		gameCorrect := true
		highestRed := 0
		highestGreen := 0
		highestBlue := 0 
		for _, set := range game.draws {
			for key, value := range set {
				if key == "blue" {
					if blueCube < value {
						gameCorrect = false
					}
					if value > highestBlue {
						highestBlue = value
					} 
				}
				if key == "red" {
					if redCube < value {
						gameCorrect = false
					}
					if value > highestRed {
						highestRed = value
					} 
				}
				if key == "green" {
					if greenCube < value {
						gameCorrect = false
					}
					if value > highestGreen {
						highestGreen = value
					} 
				}
			}
		}
		fmt.Println(line)
		fmt.Println("Game Status:", gameCorrect)
		if gameCorrect {
			correctIdsSum += game.id
		}
		fmt.Println("Ids sum:", correctIdsSum)

		// calculate the power
		fmt.Println("HRed:",highestRed,"HBlue:", highestBlue, "HGreen:", highestGreen)
		power := highestBlue * highestGreen * highestRed
		fmt.Println("Power:", power)
		powerSum += power
		fmt.Println("Power Sum:", powerSum)
	}
}


//Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func parseGame(line string) (game, error) {
	var gameReturn game
	gameLine := strings.FieldsFunc(line, func(r rune) bool {
		return r == ':'
	})
	var err error
	gameReturn.id, err = getSingleInt(gameLine[0])
	if err!= nil {
		return gameReturn, err
	}

	gameSet := strings.FieldsFunc(gameLine[1], func(r rune) bool {
		return r == ';'
	})

	for _, value := range gameSet {
		var set map[string]int
		set = make(map[string]int)
		cubes := strings.FieldsFunc(value, func(r rune) bool {
			return r == ','
		})
		for _, cube := range cubes {
			cubeStats := strings.FieldsFunc(cube, func(r rune) bool {
				return r == ' '
			})
			cubeNum, err := strconv.Atoi(cubeStats[0])
			if err != nil {
				return gameReturn, err
			}
			set[cubeStats[1]] = cubeNum
		}
		gameReturn.draws = append(gameReturn.draws, set)
	}

	return gameReturn, err
}

func getSingleInt(input string) (int, error) {
	re := regexp.MustCompile("[0-9]+")
	found := re.FindAllString(input, -1)
	if len(found) > 1{
		return -1, &TooManyInt{}
	}
	returnInt, err := strconv.Atoi(found[0])
	if err != nil {
		return -1, err
	}

	return returnInt, nil
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