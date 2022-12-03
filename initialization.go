package main

import (
	"math"
	"math/rand"
	"time"
)

// InitializeTrail creates a 2x2 table representing the pheromone intensity at every "edge" between  pairs of towns in our simulation
// Input: number of towns in this simulation, the initial trail intensity to set to every "edge"
// Output: Initial pheromone table with total trail intensity set to initialTrailIntensity
func InitializeTrail(numTowns int, initialTrailIntensity float64) PheromoneTable {

	//initialize pheromone table
	initialTrail := make(PheromoneTable, numTowns)

	for i := range initialTrail {
		initialTrail[i] = make([]*Trail, numTowns)
		for j := range initialTrail[i] {
			//set each edge to initialTrailIntensity
			if i == j {
				var curTrail Trail
				curTrail.totalTrail = 0.0
				curTrail.deltaTrail = 0.0
				initialTrail[i][j] = &curTrail
			} else {
				var curTrail Trail
				curTrail.totalTrail = initialTrailIntensity
				curTrail.deltaTrail = 0.0
				initialTrail[i][j] = &curTrail
			}

		}
	}

	return initialTrail
}

//InitializeMap creates initial map object before the simulation runs
//Input: initialized pheromone table, number of towns, and the width of the map
//Output: Map object with these fields set, shortest distance set to longest possible distance????
//func InitializeMap(initialTrailTable PheromoneTable, numTowns, width int) Map

// initialize the positions of each town in the map
func InitializeMap(initialTrail PheromoneTable, numTowns int, width float64) Map {
	var initialMap Map
	initialMap.towns = make([]*Town, numTowns)
	initialMap.pheromones = initialTrail
	initialMap.width = width

	// range over the towns and create a random position within the map
	for i := range initialMap.towns {
		var curTown Town
		curTown.label = i
		curTown.position.x = rand.Float64() * initialMap.width
		curTown.position.y = rand.Float64() * initialMap.width
		initialMap.towns[i] = &curTown
	}

	return initialMap
}

// InitializeDistanceMatrix initializes a distance matrix for all pairs of towns in the map
// Input: initial map to access towns
// Output: 2x2 slice representing the distance between every pair of towns
func InitializeDistanceMatrix(initialMap Map) [][]float64 {

	//initialize distance matrix
	distMatrix := make([][]float64, len(initialMap.towns))

	for i := range distMatrix {
		distMatrix[i] = make([]float64, len(initialMap.towns))
		for j := range distMatrix[i] {
			distMatrix[i][j] = Distance(initialMap.towns[i].position, initialMap.towns[j].position)
		}
	}
	return distMatrix
}

// InitializeAnts places an ant onto the location of one of the towns and adds that town to the ant's tabu list
// Input: initialMap to access info about towns
// Output: Ants with new added tabu list
func InitializeAnts(initialMap Map, numAnts int) []*Ant {
	// create an array of ants
	ants := make([]*Ant, numAnts)

	// seed generator for rand.Intn
	rand.Seed(time.Now().UnixNano())

	// loop through total number of ants
	for i := 0; i < numAnts; i++ {
		// randomly assign a town as a current town of the ant

		randTown := rand.Intn(len(initialMap.towns)) // spit out a number from 0 - #of towns noninclusive
		var curAnt Ant
		curAnt.cur = initialMap.towns[randTown]
		curAnt.tabu = append(curAnt.tabu, curAnt.cur)
		curAnt.totalDistance = 0.0
		ants[i] = &curAnt

	}
	return ants
}

// Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
// This is used to calculate the distance between each town and all other towns, appending the value into a table
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
