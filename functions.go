package main

import (
	"math"

	"github.com/jmcvetta/randutil"
	//"math/rand"
	//"time"
)

//AntColony will simulate numCycles cycles of the traveling salesman problem
//Input: initial map with pheromone trail set to initial trail intensity, calculation parameters,
// the total number of cycles, and the number of ants being simulated at each cycle
//Output: an array of maps where each map contains the shortest distance found from that cycle

// removed alpha, beta, rho and Q from the input line so we can put as command line arguments

func AntColony(initialMap Map, numCycles, numAnts int, alpha, beta, rho, Q float64) []Map {

	// in this case each timePoint is after a complete cycle for all ants (visiting all the towns)
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

// UpdateMap takes the current Map and system parameters and runs the simulation one time
// (all ants go to all cities once), returning currentMap with updated values
func UpdateMap(currentMap Map, numAnts int, alpha, beta, rho, Q float64) Map {
	//need to move ants from town to town, storing in currentMap.ants
	updatedAnts := UpdateAnts(currentMap, alpha, beta) //updatedAnts is the same as currentMap but with the ants field updated

	//need to update currentMap.pheromones based on the routes the ants took
	updatedPheromones := UpdatePheromoneTable(updatedAnts, rho, Q) //updatedPheromones is the same as updatedAnts but with the pheromones field updated

	return updatedPheromones
}

// UpdatePheromone Table takes the currentMap and two constants, rho and Q as input and updates the pheromone
// table once the ants have completed one cycle and returns the currentMap with updated pheromone values.
func UpdatePheromoneTable(currentMap Map, rho float64, Q float64) Map {
	for n := 0; n < len(currentMap.pheromones); n++ {
		for m := 0; m < len(currentMap.pheromones[n]); m++ {
			for i := 0; i < len(currentMap.ants); i++ {
				for j := 0; j < len(currentMap.ants[i].tabu); j++ {
					a := currentMap.ants[i].tabu[j].label
					b := currentMap.ants[i].tabu[(j + 1)].label

					//checking to see if the matrix is symmetrical
					if a == n && b == m {

						// quantity per unit of length of trail substance
						dtk := Q / currentMap.ants[i].totalDistance

						currentMap.pheromones[n][m].deltaTrail += dtk
					}
				}
			}
			//Updating the trail intensity(totalTrail) once the ant has completed one cycle
			//1-rho represents the evaporation of pheromones from the trail
			currentMap.pheromones[n][m].totalTrail += ((1 - rho) * currentMap.pheromones[n][m].totalTrail) + currentMap.pheromones[n][m].deltaTrail

			//reset deltaTrail to zero
			currentMap.pheromones[n][m].deltaTrail = 0

		}
	}

	return currentMap
}

// UpdateAnts takes currentMap and updates the ants field so that they've all gone through the cycle once,
// it returns a map that's the same as currentMap but with the new ants values
func UpdateAnts(currentMap Map, alpha, beta float64) Map {
	//range through all ants and get what path each took to travel to each town
	for antIndex := range currentMap.ants {
		currentMap.ants[antIndex] = MoveAnt(currentMap.ants[antIndex], currentMap.towns, currentMap.pheromones, alpha, beta)
	}

	return currentMap
}

// MoveAnt takes the current ant, the list of towns, the pheromone table, and system parameters and simulates one ant's
// journey across each town. It returns the updated ant object
func MoveAnt(currentAnt *Ant, towns []*Town, pheromones PheromoneTable, alpha, beta float64) *Ant {
	//set up loop so that ant keeps picking a next town until it's gone to every town
	for len(currentAnt.tabu) < len(towns) {
		//ant has starting town, need to take weighted probability of all other towns and pick next town
		currentAnt.next = PickNextTown(currentAnt.cur, towns, pheromones, alpha, beta)

		//calc dist from current town to next town and add to total distance
		dist := Distance(currentAnt.cur.position, currentAnt.next.position)
		currentAnt.totalDistance += dist

		//now add that town to tabu list
		currentAnt.tabu = append(currentAnt.tabu, currentAnt.cur)

		//move ant to next town
		currentAnt.cur = currentAnt.next
	}

	return currentAnt
}

// PickNextTown takes the pheromone table and system parameters to make a weighted probability decision about what town
// to travel to next. Returns the town that will be next
func PickNextTown(currentTown *Town, towns []*Town, pheromones PheromoneTable, alpha, beta float64) *Town {
	//use GitHub package randutil to do weighted random selection
	choices := make([]randutil.Choice, len(towns)-1)

	//range through all possibilities of towns, excluding current town, and calc prob for each
	for townIndex := range towns {
		if currentTown.label != towns[townIndex].label {
			choices[townIndex].Item = towns[townIndex].label
			choices[townIndex].Weight = 1 //INSERT EQ FROM PAPER
		}

	}

	//now pick based on weighted random probability
	choice, err := randutil.WeightedChoice(choices)
	if err != nil {
		panic(err)
	}

	//take item from choice and determine town it corresponds to
	nextTown := towns[choice.Item.(int)]

	return nextTown
}

// Distance takes two position ordered pairs and it returns the distance between these two points in 2-D space.
// This is used to calculate the distance between each town and all other towns, appending the value into a table
func Distance(p1, p2 OrderedPair) float64 {
	// this is the distance formula from days of precalculus long ago ...
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}
