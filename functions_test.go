package main

import (
	"fmt"
	"testing"
)

// Testing functions here
func TestInitializeAnts(t *testing.T) {
	var testMap Map

	testMap.towns = make([]*Town, 3)

	var towns1, towns2, towns3 Town

	towns1.label = 0
	towns1.position.x = 10
	towns1.position.y = 10
	testMap.towns[0] = &towns1

	towns2.label = 1
	towns2.position.x = 10
	towns2.position.y = 40
	testMap.towns[1] = &towns2

	towns3.label = 2
	towns3.position.x = 30
	towns3.position.y = 20
	testMap.towns[2] = &towns3

	outcome := InitializeAnts(testMap, 2)

	// ERROR: getting an error saying there is an improper pointer or dreference being used
	//outcome := InitializeAnts(testMap, 2)

	if outcome[0].cur.label != outcome[0].tabu[0].label {
		t.Error("The current town and first town on tabu list do not match.")
	}

	//fmt.Println("First ant located at town:", ant1_town, "\nFirst ant tabu:", ant1_tabu, "\nSecond ant located:", ant2_town, "\nSecond ant located:", ant2_tabu)
}

func TestInitializeDistanceMatrix(t *testing.T) {
	var testMap Map

	testMap.towns = make([]*Town, 3)

	var t1, t2, t3 Town
	t1.label = 0
	t1.position.x = 3
	t1.position.y = 5
	testMap.towns[0] = &t1

	t2.label = 1
	t2.position.x = 6
	t2.position.y = 5
	testMap.towns[1] = &t2
	// once the pointer has had an address , in go, you can use the pointer to access the original and make changes
	// instead of making a copy with the changes! , example : use the following to "edit" t2.position.x "testMap.towns[1].position.x = 0"

	t3.label = 0
	t3.position.x = 9
	t3.position.y = 5
	testMap.towns[2] = &t3

	answer := [][]float64{
		{0, 3, 6},
		{3, 0, 3},
		{6, 3, 0},
	}

	outcome := InitializeDistanceMatrix(testMap)
	for i := range answer {
		for j := range answer[i] {
			if outcome[i][j] != answer[i][j] {
				t.Errorf("Error! for input distance matrix, your code does not match answer key at row[%d] and col[%d])", i, j)
			}
		}
	}
	fmt.Println("Success! Correct Distance Matrix")
	fmt.Println("Outcome:", outcome, "\nAnswer:", answer)
}
