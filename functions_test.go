package main

import (
	"testing"
)

// Testing functions here
func TestInitializeAnts(t *testing.T) {

	type test struct {
		initialMap Map
		numAnts    int
		answer     int
	}

	var testMap test

	testMap.initialMap.towns = make([]*Town, 3)

	var towns1, towns2, towns3 Town

	towns1.label = 0
	towns1.position.x = 10
	towns1.position.y = 10
	testMap.initialMap.towns[0] = &towns1

	towns2.label = 1
	towns2.position.x = 10
	towns2.position.y = 40
	testMap.initialMap.towns[1] = &towns2

	towns3.label = 2
	towns3.position.x = 30
	towns3.position.y = 20
	testMap.initialMap.towns[2] = &towns3

	testMap.answer = 3

	outcome := InitializeAnts(testMap.initialMap, 3)

	if len(outcome) != testMap.answer {
		t.Errorf("Error. InitializeAnts results in array of length %d and real answer is %d", len(outcome), testMap.answer)
	}

	//fmt.Println("First ant located at town:", ant1_town, "\nFirst ant tabu:", ant1_tabu, "\nSecond ant located:", ant2_town, "\nSecond ant located:", ant2_tabu)
}

func TestInitializeDistanceMatrix(t *testing.T) {

	type test struct {
		initialMap Map
		answer     [][]float64
	}

	var testMap test

	testMap.initialMap.towns = make([]*Town, 3)

	var t1, t2, t3 Town
	t1.label = 0
	t1.position.x = 3
	t1.position.y = 5
	testMap.initialMap.towns[0] = &t1

	t2.label = 1
	t2.position.x = 6
	t2.position.y = 5
	testMap.initialMap.towns[1] = &t2
	// once the pointer has had an address , in go, you can use the pointer to access the original and make changes
	// instead of making a copy with the changes! , example : use the following to "edit" t2.position.x "testMap.towns[1].position.x = 0"

	t3.label = 0
	t3.position.x = 9
	t3.position.y = 5
	testMap.initialMap.towns[2] = &t3

	testMap.answer = [][]float64{
		{0, 3, 6},
		{3, 0, 3},
		{6, 3, 0},
	}

	outcome := InitializeDistanceMatrix(testMap.initialMap)
	for i := range outcome {
		for j := range outcome[i] {
			if outcome[i][j] != testMap.answer[i][j] {
				t.Errorf("Error! At position %d,%d, function resulted %f, real answer is %f", i, j, outcome[i][j], testMap.answer[i][j])
			}
		}
	}
}

func TestInitializeTrail(t *testing.T) {

	type test struct {
		numTowns              int
		initialTrailIntensity float64
		answer                PheromoneTable
	}

	var testTable test
	testTable.numTowns = 3
	testTable.initialTrailIntensity = 0.2

	var testTrail Trail
	testTrail.totalTrail = 0.2
	testTrail.deltaTrail = 0.0

	var testNoTrail Trail
	testNoTrail.totalTrail = 0.0
	testNoTrail.deltaTrail = 0.0

	testTable.answer = [][]*Trail{
		{&testNoTrail, &testTrail, &testTrail},
		{&testTrail, &testNoTrail, &testTrail},
		{&testTrail, &testTrail, &testNoTrail}}

	outcome := InitializeTrail(testTable.numTowns, testTable.initialTrailIntensity)

	for i := range outcome {
		for j := range outcome[i] {
			if outcome[i][j].totalTrail != testTable.answer[i][j].totalTrail {
				t.Errorf("Error at position %d,%d, the function returned total trail %f and the real value is %f", i, j, outcome[i][j].totalTrail, testTable.answer[i][j].totalTrail)
			} else if outcome[i][j].deltaTrail != testTable.answer[i][j].deltaTrail {
				t.Errorf("Error at position %d,%d, the function returned delta trail %f and the real value is %f", i, j, outcome[i][j].deltaTrail, testTable.answer[i][j].deltaTrail)
			}
		}
	}

}

func TestInitializeMap(t *testing.T) {

	type test struct {
		initialTrail PheromoneTable
		numTowns     int
		width        float64
		answer       Map
	}

	var testMap test
	var testTrail Trail
	testTrail.totalTrail = 0.2
	testTrail.deltaTrail = 0.0

	var testNoTrail Trail
	testNoTrail.totalTrail = 0.0
	testNoTrail.deltaTrail = 0.0

	testMap.initialTrail = [][]*Trail{
		{&testNoTrail, &testTrail, &testTrail},
		{&testTrail, &testNoTrail, &testTrail},
		{&testTrail, &testTrail, &testNoTrail}}

	testMap.numTowns = 3
	testMap.width = 100

	testMap.answer.width = 100
	testMap.answer.pheromones = [][]*Trail{
		{&testNoTrail, &testTrail, &testTrail},
		{&testTrail, &testNoTrail, &testTrail},
		{&testTrail, &testTrail, &testNoTrail}}

	testMap.answer.towns = make([]*Town, 3)

	var towns1, towns2, towns3 Town

	towns1.label = 0
	towns1.position.x = 10
	towns1.position.y = 10
	testMap.answer.towns[0] = &towns1

	towns2.label = 1
	towns2.position.x = 10
	towns2.position.y = 40
	testMap.answer.towns[1] = &towns2

	towns3.label = 2
	towns3.position.x = 30
	towns3.position.y = 20
	testMap.answer.towns[2] = &towns3

	outcome := InitializeMap(testMap.initialTrail, testMap.numTowns, testMap.width)

	for i := range outcome.pheromones {
		for j := range outcome.pheromones[i] {
			if outcome.pheromones[i][j].totalTrail != testMap.answer.pheromones[i][j].totalTrail {
				t.Errorf("Error at position %d,%d, the function returned total trail %f and the real value is %f", i, j, outcome.pheromones[i][j].totalTrail, testMap.answer.pheromones[i][j].totalTrail)
			} else if outcome.pheromones[i][j].deltaTrail != testMap.answer.pheromones[i][j].deltaTrail {
				t.Errorf("Error at position %d,%d, the function returned delta trail %f and the real value is %f", i, j, outcome.pheromones[i][j].deltaTrail, testMap.answer.pheromones[i][j].deltaTrail)
			}
		}
	}

	if outcome.width != testMap.answer.width {
		t.Errorf("Error. Function returned width of %f, and real width is %f", outcome.width, testMap.answer.width)
	}

	for i := range outcome.towns {
		if outcome.towns[i].label != testMap.answer.towns[i].label {
			t.Errorf("Error at position %d. Function returned label of %d and real label is %d", i, outcome.towns[i].label, testMap.answer.towns[i].label)
		}
	}

}
