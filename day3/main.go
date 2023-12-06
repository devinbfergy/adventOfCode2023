package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"unicode"
)

type engine_object struct {
	x int
	y int
	value string
	digit bool
	period bool
	hasSymbolTouching bool
	isStar bool
}

func main() {
	lines, err := readLines("input.txt")
	if err != nil {
		log.Fatalf("Unable to load input file: %s", err)
	}
	var engine_data [][]engine_object 
	var stars []engine_object
	for row, line := range lines {
		var line_data []engine_object
		for col, character := range line {
			digit_bool := isDigit(character)
			period_bool := isPeriod(character)
			star := isStar(character)
			line_data = append(line_data, engine_object{
				x : row,
				y : col,
				value : string(character),
				digit : digit_bool,
				period : period_bool,
				hasSymbolTouching: false,
				isStar: star,
				})
			if star {
				stars = append(stars, engine_object{
					x : row,
					y : col,
					value : string(character),
					digit : digit_bool,
					period : period_bool,
					hasSymbolTouching: false,
					isStar: star,
					})
			}
		}
		engine_data = append(engine_data, line_data)
	}
	var digits [][]engine_object
	// initial digit
	var digit []engine_object
	for _, row := range engine_data {
		for i, item := range row {
			log.Println(item.x, item.y, item.value)
			if item.period || !item.digit {
				if len(digit) > 0 {
					digits = append(digits, digit)
					digit = []engine_object{}
				}
			}
			if item.digit {
				hasSymbolTouching := touchingSymbol(item.x, item.y, engine_data)
				log.Println("Has Symbol:", hasSymbolTouching)
				item.hasSymbolTouching = hasSymbolTouching
				digit = append(digit, item)
			}
			// if the end of row append the digit above to digits
			if len(row)-1 == i && item.digit {
				digits = append(digits, digit)
			    digit = []engine_object{}
			}
		}
	}
	total_count := 0 
	for _, d := range digits {
		var num string
		touch := false
		for _, item := range d {
			num = num + item.value
			if item.hasSymbolTouching {
				touch = true
			}
		}
		log.Println(num, touch)
		if touch {
			i, err := strconv.Atoi(num)
			if err != nil {
				log.Fatalf("Unable to convert %s msg: %s", num, err)
			}
			total_count += i
		}
	}
	log.Println("Total Count from Touching:", total_count)

	total_star := 0
	for _, star := range stars {
		log.Println("star:", star.x, star.y, star.value)
		num1 := ""
		num2 := ""
		for _, dig := range digits {
			var num string 
			touchingStar := false
			
			for _, i := range dig {
				num = num + i.value
				if star.x - 1 == i.x && star.y == i.y {
					 touchingStar = true
				} else if star.x + 1 == i.x && star.y == i.y {
					touchingStar = true
				} else if star.x == i.x && star.y - 1 == i.y {
					touchingStar = true
				} else if star.x == i.x && star.y + 1 == i.y {
					touchingStar = true
				} else if star.x - 1 == i.x && star.y + 1 == i.y {
					touchingStar = true
				} else if star.x + 1 == i.x && star.y + 1 == i.y {
					touchingStar = true
				} else if star.x + 1 == i.x && star.y - 1 == i.y {
					touchingStar = true
				} else if star.x - 1 == i.x && star.y - 1 == i.y {
					touchingStar = true
				}
			}
			if num1 == "" && touchingStar {
				num1 = num
			} else if num2 == "" && touchingStar {
				num2 = num
				break
			} 
		}
		if num1 != "" && num2 != "" {
			log.Println("num1", num1, "num2", num2)
			int1, err := strconv.Atoi(num1)
			if err != nil {
				log.Fatalf("Unable to convert %s msg: %s", num1, err)
			}
			int2, err := strconv.Atoi(num2)
			if err != nil {
				log.Fatalf("Unable to convert %s msg: %s", num2, err)
			}
			total_star += int1 * int2
			num1 = ""
			num2 = ""
		}
	} 
	log.Println("Total gear ratio:", total_star)
}

func touchingSymbol(x int, y int, engine_matrix [][]engine_object) bool {
	// check above 
	north := checkForSymbol(x-1, y, engine_matrix)
	south := checkForSymbol(x+1, y, engine_matrix)
	east := checkForSymbol(x, y+1, engine_matrix)
	west := checkForSymbol(x, y-1, engine_matrix)
	n_east := checkForSymbol(x-1, y+1, engine_matrix)
	n_west := checkForSymbol(x-1, y-1, engine_matrix)
	s_east := checkForSymbol(x+1, y+1, engine_matrix)
	s_west := checkForSymbol(x+1, y-1, engine_matrix)

	return north || south || east || west || n_east || n_west || s_east || s_west
}

// looks into the engine matrix and sees if it is a symbol at that place
func checkForSymbol(x int, y int, engine_matrix [][]engine_object) bool {
	if x < 0 || y < 0 {
		return false
	} 
	if x >= len(engine_matrix) {
		return false
	}
	if y >= len(engine_matrix[x]){
		return false
	}
	item := engine_matrix[x][y]
	return !item.digit && !item.period 
}

func isStar(s rune) bool {
	return s == '*'
}

func isDigit(s rune) bool {
	return unicode.IsDigit(s)
}

func isPeriod(s rune) bool {
	return s == '.'
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