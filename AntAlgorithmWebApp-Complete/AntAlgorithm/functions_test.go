package AntAlgorithm

import (
	"fmt"
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

	useValidation := false

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

	outcome := InitializeMap(testMap.initialTrail, testMap.numTowns, testMap.width, useValidation)

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

// Write this function after PickNextTown is done
func TestPickNextTown(t *testing.T) {
	alpha := 1.
	beta := 1.

	//need to initialize a pheromone table, list of towns, distance matrix, and list of ants
	numTowns := 4
	initialPheromone := 1.
	numAnts := 10
	width := 20.
	useValidation := false
	testPheromoneTable := InitializeTrail(numTowns, initialPheromone)
	currentMap := InitializeMap(testPheromoneTable, numTowns, width, useValidation)
	currentMap.distanceMatrix = InitializeDistanceMatrix(currentMap)
	currentMap.ants = InitializeAnts(currentMap, numAnts)

	//can modify pheromone table to have different values if wanted

	//now call PickNextTown
	//need to decide what to test/evaluate to determine if working correctly
	//maybe run multiple times and look at probability it chooses each next town for an ant

	//get list of where each ant would go next
	nextTownList := make([]int, numAnts)
	numCycles := 20
	for cycleNum := 0; cycleNum < numCycles; cycleNum++ {
		nextTown := PickNextTown(currentMap.ants[0], currentMap, alpha, beta)
		nextTownList[cycleNum] = nextTown.label
	}

	//now calc experimental prob of choosing each town
	townCounterList := make([]int, numTowns) //list where each index corresponds to a town label; values are the number of times that town was picked
	for _, val := range nextTownList {
		for i := range townCounterList {
			if val == i {
				townCounterList[i]++
			}
		}
	}

	//now turn counts into prob of picking each town
	townProbList := make([]float64, numTowns)
	for townIndex, count := range townCounterList {
		townProbList[townIndex] = float64(count / numCycles)
	}

	fmt.Println(townProbList)

	//would need to get currentMap.ants[0].cur so can do probability calc by hand and verify probabilities of choosing towns are correct
	//not sure if there'd be a way to have Go do this or make a test set b/c ant location and town position are randomized

}

func TestUpdatePheromoneTable(t *testing.T) {

	type test struct {
		initialTrail PheromoneTable
		numTowns     int
		width        float64
		answer       Map
	}

	var testMap test

	testMap.answer.width = 50
	testMap.numTowns = 3
	// might have to test if InitializeTrail is working
	initialTrailIntensity := 1.0
	testMap.answer.pheromones = InitializeTrail(testMap.numTowns, initialTrailIntensity)

	// panic out of range issue below, but why?

	testMap.answer.towns = make([]*Town, 3)

	var town1, town2, town3 Town

	town1.label = 0
	town1.position.x = 10
	town1.position.y = 10
	testMap.answer.towns[0] = &town1

	town2.label = 1
	town2.position.x = 10
	town2.position.y = 40
	testMap.answer.towns[1] = &town2

	town3.label = 2
	town3.position.x = 30
	town3.position.y = 10
	testMap.answer.towns[2] = &town3

	testMap.answer.ants = make([]*Ant, 3)

	var ant1, ant2, ant3 Ant

	ant1.tabu = make([]*Town, 3)
	ant2.tabu = make([]*Town, 3)
	ant3.tabu = make([]*Town, 3)

	ant1.tabu[0] = &town1
	ant1.tabu[1] = &town2
	ant1.tabu[2] = &town3
	// wait Adonis, total distance would be 20+30+approx36.06 for each ant because there are only 3 options, so maybe just have 1 ant update the table for the test
	ant1.totalDistance = 86.06
	// assuming ant[0] started in town 1
	testMap.answer.ants[0] = &ant1

	ant2.tabu[0] = &town1
	ant2.tabu[1] = &town2
	ant2.tabu[2] = &town3

	ant2.totalDistance = 86.06
	testMap.answer.ants[1] = &ant2

	ant3.tabu[0] = &town1
	ant3.tabu[1] = &town2
	ant3.tabu[2] = &town3
	ant3.totalDistance = 86.06
	testMap.answer.ants[2] = &ant3

	rho := 0.50
	Q := 50.0
	// Q/L = 0.5431
	// since rho = 0.5, then the trail intensity will be half of what it was (in this case 0.5)
	// then add 0.5431, which will have an overall addition of 0.0431
	// at the end of this single trial run, the pheromone table should read a 1.0431 for each trail since this ant travelled each route

	outcomePheromone := UpdatePheromoneTable(testMap.answer, rho, Q)

	for i := range outcomePheromone {
		for j := range outcomePheromone[i] {
			fmt.Println(outcomePheromone[i][j].totalTrail)
		}
	}

	fmt.Println(outcomePheromone[0])

}
