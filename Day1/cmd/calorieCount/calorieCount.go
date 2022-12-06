package calorieCount

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

// Elf represents an Elf carrying a certain amount of food.
type Elf struct {
	ID    int
	Total int
}

func Init() {
	// Read the input data
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// Read the input data.
	scanner := bufio.NewScanner(file)
	elves := make(map[int]int)
	id := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			// New Elf
			id++
			continue
		}

		// Add Calories to the current Elf's total
		calories, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		elves[id] += calories
	}

	// Sort the Elves by their total number of Calories
	sortedElves := make([]Elf, 0, len(elves))
	for id, total := range elves {
		sortedElves = append(sortedElves, Elf{ID: id, Total: total})
	}
	sort.Slice(sortedElves, func(i, j int) bool {
		return sortedElves[i].Total > sortedElves[j].Total
	})

	// Print the top three Elves and their total Calories
	fmt.Println("Top three Elves:")
	total := 0
	for i := 0; i < 3; i++ {
		el := sortedElves[i]
		fmt.Printf("Elf %d: %d Calories\n", el.ID, el.Total)
		total += el.Total
	}

	fmt.Printf("Total Calories: %d\n", total)
}
