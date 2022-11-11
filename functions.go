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

func AntColony(initialMap Map, numCycles, numAnts int) []Map {

	// in this case each timePoint is after a complete ant cycle (visiting all the towns)
	timePoints := make([]Map, numCycles)
	timePoints[0] = initialMap

	for i := 1; i <= numCycles; i++ {
		timePoints[i] = UpdateMap(timePoints[i-1], numAnts)
	}

	/*
		// I am confused on what the shortestMaps are supposed to give us, should we use timePoints to represent the updated map after each cycle?
		var shortestMaps []Map
	*/

	//run subroutines from here

	return timePoints
}

// need to make row and col
func DistanceMatrix(initialMap Map) [][]float64 {
	// set the rows and cols to be equal to numTowns
	numRows := len(initialMap.towns)
	numCols := len(initialMap.towns)
	distMatrix := InitializeDistMatrix(numRows, numCols)

	// range over the towns
	for i := range initialMap.towns {
		// compare town[i] to town[j]
		for j := range initialMap.towns {
			// do not compare town distance to self since that will be zero anyways
			if initialMap.towns[i] != initialMap.towns[j] {
				distMatrix[i][j] = Distance(initialMap.towns[i].position, initialMap.towns[j].position)
			}
		}
	}
	return distMatrix
}

// InitializeBoard takes a number of rows and columns as inputs and it returns a gameboard with an appropriate number of rows and columns, where all values = 0
func InitializeDistMatrix(numRows, numCols int) [][]float64 {
	// make a 2-D slice
	distMatrix := make([][]float64, numRows)
	// now we need to make the the rows (as in having the length of each row)
	for r := range distMatrix {
		distMatrix[r] = make([]float64, numCols)
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
