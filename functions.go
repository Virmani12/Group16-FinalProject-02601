package main

import (
	"math"
	"math/rand"
	"time"
)

//AntColony will simulate numCycles cycles of the traveling salesman problem
//Input: initial map with pheromone trail set to initial trail intensity, calculation parameters,
// the total number of cycles, and the number of ants being simulated at each cycle
//Output: an array of maps where each map contains the shortest distance found from that cycle

// removed alpha, beta, rho and Q from the input line so we can put as command line arguments

func AntColony(initialMap Map, numCycles, numAnts int, alpha, beta, rho, Q float64) []Map {

	// in this case each timePoint is after a complete ant cycle (visiting all the towns)
	timePoints := make([]Map, numCycles)
	timePoints[0] = initialMap

	for i := 1; i <= numCycles; i++ {
		timePoints[i] = UpdateMap(timePoints[i-1], numAnts, alpha, beta, rho, Q)
	}

	/*
		// I am confused on what the shortestMaps are supposed to give us, should we use timePoints to represent the updated map after each cycle?
		var shortestMaps []Map
	*/

	//run subroutines from here

	return timePoints
}

//InitializeDistanceMatrix initializes a distance matrix for all pairs of towns in the map
//Input: initial map to access towns
//Output: 2x2 slice representing the dsitance between every pair of towns
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
		ants[i].cur = *initialMap.towns[randTown]
		ants[i].tabu = append(ants[i].tabu, &ants[i].cur)
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
