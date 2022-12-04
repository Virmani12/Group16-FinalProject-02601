package main

import (
	"fmt"
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

	//changed this because timePoint[0] would have just been the initialized map which would have caused issues with drawing
	//now, if we're on the INITIAL map, we'll update the SAME INITIAL map to represent the first cycle being complete
	//now, timePoints[0] = the result from the first cycle, timePoints[1] = the result from the second cycle
	//When we draw, we'll take the initial map from main as the starting point, and then use timePoints to represent the completion of each cycle
	//also this helps with shortestTour because now shortestTour[0] corresponds to the shortest tour from cycle[0]
	for i := 0; i < numCycles; i++ {
		if i == 0 {
			timePoints[i] = UpdateMap(initialMap, numAnts, alpha, beta, rho, Q)

		} else {

			timePoints[i] = UpdateMap(timePoints[i-1], numAnts, alpha, beta, rho, Q)
		}

	}

	return timePoints
}

// UpdateMap takes the current Map and system parameters and runs the simulation one time
// (all ants go to all cities once), returning currentMap with updated values
func UpdateMap(currentMap Map, numAnts int, alpha, beta, rho, Q float64) Map {

	//set updatedAnts to current map and initialize ants here (instead of main), this way we can reset the ants tabu list and initialize the new set of ants in one step
	currentMap.ants = nil
	currentMap.ants = InitializeAnts(currentMap, numAnts)

	//need to move ants from town to town, storing in currentMap.ants
	for i := 0; i < len(currentMap.towns); i++ {

		if i < len(currentMap.towns)-1 {
			currentMap = UpdateAnts(currentMap, alpha, beta)
		} else {
			currentMap = ReturnTowns(currentMap)
		}

	}
	//updatedAnts is the same as currentMap but with the ants field updated

	//update the shortest tour found in this cycle
	shortestTour := FindShortestTour(currentMap)
	currentMap.shortestTours = append(currentMap.shortestTours, shortestTour)

	//need to update currentMap.pheromones based on the routes the ants took
	currentMap = UpdatePheromoneTable(currentMap, rho, Q) //updatedPheromones is the same as updatedAnts but with the pheromones field updated

	return currentMap
}

// UpdatePheromoneTable takes the currentMap and two constants, rho and Q as input and updates the pheromone
// table once the ants have completed one cycle and returns the currentMap with updated pheromone values.
func UpdatePheromoneTable(currentMap Map, rho float64, Q float64) Map {

	for i := range currentMap.ants {
		for j := 0; j < len(currentMap.ants[i].tabu)-1; j++ {
			a := currentMap.ants[i].tabu[j].label
			b := currentMap.ants[i].tabu[j+1].label

			dtk := Q / currentMap.ants[i].totalDistance
			currentMap.pheromones[a][b].deltaTrail += dtk
		}
	}

	/*
		for n := 0; n < len(currentMap.pheromones); n++ {
			for m := 0; m < len(currentMap.pheromones[n]); m++ {
				if m != n {
					for i := 0; i < len(currentMap.ants); i++ {
						for j := 0; j < len(currentMap.ants[i].tabu)-1; j++ {
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
				}

			}
		}
	*/
	for n := 0; n < len(currentMap.pheromones); n++ {
		for m := 0; m < len(currentMap.pheromones[n]); m++ {
			if m != n {
				//Updating the trail intensity(totalTrail) once the ant has completed one cycle
				//1-rho represents the evaporation of pheromones from the trail
				currentMap.pheromones[n][m].totalTrail += (rho * currentMap.pheromones[n][m].totalTrail) + currentMap.pheromones[n][m].deltaTrail

				//reset deltaTrail to zero
				currentMap.pheromones[n][m].deltaTrail = 0
			}
			//fmt.Print(currentMap.pheromones[n][m].totalTrail, " ")

		}
		//fmt.Println("")
	}

	return currentMap
}

// on the last iteration of this cycle, this function is called and it closes the tour each ant took by moving it to the town it started at and updating the distance
func ReturnTowns(currentMap Map) Map {

	//for each ant, set their next town to their original town, caluclate the distance from cur to next, add it to their total distance, and set cur to the first town
	for antIndex := range currentMap.ants {

		currentMap.ants[antIndex].next = currentMap.ants[antIndex].tabu[0]

		dist := currentMap.distanceMatrix[currentMap.ants[antIndex].cur.label][currentMap.ants[antIndex].next.label]

		currentMap.ants[antIndex].totalDistance += dist

		currentMap.ants[antIndex].cur = currentMap.ants[antIndex].next
	}

	return currentMap
}

// this function finds the shortest tour in this cycle based on each ants total distance
func FindShortestTour(currentMap Map) []*Town {
	//set shortest tour and distance to first tabu list
	shortestTour := currentMap.ants[0].tabu
	shortestDistance := currentMap.ants[0].totalDistance

	// if other ants distances are shorter, update the shortest tour to their tabu list
	for antIndex := 1; antIndex < len(currentMap.ants); antIndex++ {

		if currentMap.ants[antIndex].totalDistance < shortestDistance {

			shortestDistance = currentMap.ants[antIndex].totalDistance

			shortestTour = currentMap.ants[antIndex].tabu
		}
	}

	return shortestTour
}

// UpdateAnts takes currentMap and updates the ants field after moving to one more town
// it returns a map that's the same as currentMap but with the new ants values
func UpdateAnts(currentMap Map, alpha, beta float64) Map {
	//range through all ants and get what path each took to travel to each town
	for antIndex := range currentMap.ants {
		currentMap.ants[antIndex] = MoveAnt(currentMap.ants[antIndex], currentMap, alpha, beta)
	}

	return currentMap
}

// MoveAnt takes the current ant, the list of towns, the pheromone table, and system parameters and simulates one ant's
// journey across each town. It returns the updated ant object
func MoveAnt(currentAnt *Ant, currentMap Map, alpha, beta float64) *Ant {
	//set up loop so that ant keeps picking a next town until it's gone to every town

	//ant has starting town, need to take weighted probability of all other towns and pick next town

	currentAnt.next = PickNextTown(currentAnt, currentMap, alpha, beta)

	//calc dist from current town to next town and add to total distance
	dist := currentMap.distanceMatrix[currentAnt.cur.label][currentAnt.next.label]
	currentAnt.totalDistance += dist

	//move ant to next town
	currentAnt.cur = currentAnt.next

	//now add that town to tabu list
	currentAnt.tabu = append(currentAnt.tabu, currentAnt.cur)

	return currentAnt
}

// PickNextTown takes the pheromone table and system parameters to make a weighted probability decision about what town
// to travel to next. Returns the town that will be next
func PickNextTown(currentAnt *Ant, currentMap Map, alpha, beta float64) *Town {
	//use GitHub package randutil to do weighted random selection
	choices := make([]randutil.Choice, len(currentMap.towns))

	totalProbability := CalculateTotalProb(currentAnt, currentMap, alpha, beta)

	//range through all possibilities of towns, excluding current town, and calc prob for each
	for townIndex := range currentMap.towns {
		if currentAnt.cur.label != currentMap.towns[townIndex].label {
			choices[townIndex].Item = currentMap.towns[townIndex].label
			if InTabu(currentMap.towns[townIndex], currentAnt.tabu) {
				choices[townIndex].Weight = 0
			} else {
				trailProb := math.Pow(currentMap.pheromones[currentAnt.cur.label][currentMap.towns[townIndex].label].totalTrail, alpha)
				distProb := math.Pow(1/(currentMap.distanceMatrix[currentAnt.cur.label][currentMap.towns[townIndex].label]), beta)
				choices[townIndex].Weight = int(100 * ((trailProb * distProb) / totalProbability))
				//fmt.Println((100 * ((trailProb * distProb) / totalProbability)))
			}

		} else {
			choices[townIndex].Item = currentMap.towns[townIndex].label
			choices[townIndex].Weight = 0
		}

	}

	//now pick based on weighted random probability
	choice, err := randutil.WeightedChoice(choices)
	if err != nil {
		panic(err)
	}

	//take item from choice and determine town it corresponds to
	nextTown := currentMap.towns[choice.Item.(int)]
	//fmt.Println(nextTown)

	return nextTown
}

// this function calculates the sum of all transition probabilites from the current ant to the towns that aren't in the current ant's tabu list
func CalculateTotalProb(currentAnt *Ant, currentMap Map, alpha, beta float64) float64 {

	totalProb := 0.0

	for townIndex := range currentMap.towns {
		if currentAnt.cur.label != currentMap.towns[townIndex].label {
			if InTabu(currentMap.towns[townIndex], currentAnt.tabu) == false {
				trailProb := math.Pow(currentMap.pheromones[currentAnt.cur.label][currentMap.towns[townIndex].label].totalTrail, alpha)
				distProb := math.Pow(1/(currentMap.distanceMatrix[currentAnt.cur.label][currentMap.towns[townIndex].label]), beta)
				totalProb += (trailProb * distProb)
			}

		}

	}
	//fmt.Println(totalProb)
	return totalProb

}

// This function checks if the current town being checked is in the current ant's tabu list
// Input: current town and current ant's tabu list
// Output: True if the town is already in the list, false if it isn't
func InTabu(town *Town, currentAntTabu []*Town) bool {
	for i := range currentAntTabu {
		if currentAntTabu[i].label == town.label {
			return true
		}
	}
	return false

}

func ComputeDistance(currentMap Map, tabuList []*Town) float64 {

	totalDist := 0.0
	for i := 0; i < len(tabuList)-1; i++ {
		curTown := tabuList[i]
		nextTown := tabuList[i+1]
		curDist := currentMap.distanceMatrix[curTown.label][nextTown.label]
		totalDist += curDist
	}

	curTown := tabuList[len(tabuList)-1]
	nextTown := tabuList[0]
	curDist := currentMap.distanceMatrix[curTown.label][nextTown.label]
	totalDist += curDist

	return totalDist
}

func CalculateAvgDist(timePoints []Map, numCycles int) float64 {
	avgDist := 0.0
	for i := range timePoints[numCycles-1].shortestTours {

		curTour := timePoints[numCycles-1].shortestTours[i]
		for j := 0; j < len(curTour)-1; j++ {
			curTown := curTour[j].label
			nextTown := curTour[j+1].label
			dist := timePoints[numCycles-1].distanceMatrix[curTown][nextTown]
			avgDist += dist
		}

	}
	avgDist = avgDist / (float64(numCycles))
	return avgDist
}

func PrintTour(tour []*Town) {

	for i := 0; i < len(tour); i++ {
		if i < len(tour)-1 {
			fmt.Printf("Town %d (%f,%f) --> ", tour[i].label, tour[i].position.x, tour[i].position.y)
		} else {
			fmt.Printf("Town %d (%f,%f)", tour[i].label, tour[i].position.x, tour[i].position.y)
			fmt.Println()
		}

	}
}
