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

func AntCycle(initialMap Map, numCycles, numAnts int, alpha, beta, rho, Q float64) []Map {

	simulation := "cycle"

	// initialize slice of maps
	timePoints := make([]Map, numCycles)

	//Loop through the number of cycles and update the map
	//The first iteration will update the same initial map such that index 0 represents the end of the first cycle
	//This helps with keeping track of the shortest tours because each shortest tour will match the index of the cycle it pertains to
	for i := 0; i < numCycles; i++ {
		if i == 0 {
			timePoints[i] = UpdateMap(initialMap, numAnts, alpha, beta, rho, Q, simulation)

		} else {

			timePoints[i] = UpdateMap(timePoints[i-1], numAnts, alpha, beta, rho, Q, simulation)
		}

	}

	return timePoints
}

func AntDensity(initialMap Map, numCycles, numAnts int, alpha, beta, rho, Q float64) []Map {

	simulation := "density"

	// initialize slice of maps
	timePoints := make([]Map, numCycles)

	//Loop through the number of cycles and update the map
	//The first iteration will update the same initial map such that index 0 represents the end of the first cycle
	//This helps with keeping track of the shortest tours because each shortest tour will match the index of the cycle it pertains to
	for i := 0; i < numCycles; i++ {
		if i == 0 {
			timePoints[i] = UpdateMap(initialMap, numAnts, alpha, beta, rho, Q, simulation)

		} else {

			timePoints[i] = UpdateMap(timePoints[i-1], numAnts, alpha, beta, rho, Q, simulation)
		}

	}

	return timePoints
}

func AntQuantity(initialMap Map, numCycles, numAnts int, alpha, beta, rho, Q float64) []Map {

	simulation := "quantity"

	// initialize slice of maps
	timePoints := make([]Map, numCycles)

	//Loop through the number of cycles and update the map
	//The first iteration will update the same initial map such that index 0 represents the end of the first cycle
	//This helps with keeping track of the shortest tours because each shortest tour will match the index of the cycle it pertains to
	for i := 0; i < numCycles; i++ {
		if i == 0 {
			timePoints[i] = UpdateMap(initialMap, numAnts, alpha, beta, rho, Q, simulation)

		} else {

			timePoints[i] = UpdateMap(timePoints[i-1], numAnts, alpha, beta, rho, Q, simulation)
		}

	}

	return timePoints
}

// UpdateMap takes the current Map and system parameters and runs the simulation one time
// (all ants go to all cities once), returning currentMap with updated values
//Input: current map, number of ants, alpha, beta, rho, Q
//Output: New map representing one complete cycle of ants visiting every town once
func UpdateMap(currentMap Map, numAnts int, alpha, beta, rho, Q float64, simulation string) Map {

	//set updatedAnts to current map and initialize ants here (instead of main),
	//this way we can reset the ants tabu list and initialize the new set of ants in one step
	var updatedMap Map
	updatedMap = currentMap
	updatedMap.ants = nil
	updatedMap.ants = InitializeAnts(updatedMap, numAnts)

	//need to move ants from town to town, storing in updatedMap.ants
	//Once each ant's tabu list is full, we will return each ant to their starting town and update the distance
	for i := 0; i < len(updatedMap.towns); i++ {

		if i < len(updatedMap.towns)-1 {
			updatedMap.ants = UpdateAnts(updatedMap, alpha, beta, Q, simulation)
		} else {
			updatedMap.ants = ReturnTowns(updatedMap, Q, simulation)
		}

	}
	//updatedAnts is the same as currentMap but with the ants field changed

	//find and update the shortest tour found in this cycle
	shortestTour := FindShortestTour(updatedMap)
	updatedMap.shortestTours = append(updatedMap.shortestTours, shortestTour)

	if simulation == "cycle" {

		//need to update currentMap.pheromones based on the routes the ants took
		updatedMap.pheromones = UpdatePheromoneTable(updatedMap, rho, Q) //updatedPheromones is the same as updatedAnts but with the pheromones field updated
	}

	return updatedMap
}

// UpdatePheromoneTable takes the currentMap and two constants, rho and Q as input and updates the pheromone
// table once the ants have completed one cycle and returns the currentMap with updated pheromone values.
//Input: current map, rho, Q
//Output: updated pheromone table for the current cycle's map
func UpdatePheromoneTable(currentMap Map, rho float64, Q float64) PheromoneTable {

	//loop through each ant's tabu list and update their tabu list
	for i := range currentMap.ants {
		for j := 0; j < len(currentMap.ants[i].tabu)-1; j++ {
			//a and b is the current edge who's trail is being updated
			a := currentMap.ants[i].tabu[j].label
			b := currentMap.ants[i].tabu[j+1].label

			//increase in delta edge is the ratio of the constant number of trail each ant can drop divided by the total distance traveled by the ant
			dtk := Q / currentMap.ants[i].totalDistance
			currentMap.pheromones[a][b].deltaTrail += dtk
			currentMap.pheromones[b][a].deltaTrail += dtk
		}

		//Manually update the trail from the last unique town visited to the original town because the original town is only in the tabu list once
		a := currentMap.ants[i].tabu[len(currentMap.ants[i].tabu)-1].label
		b := currentMap.ants[i].tabu[0].label

		dtk := Q / currentMap.ants[i].totalDistance
		currentMap.pheromones[a][b].deltaTrail += dtk
		currentMap.pheromones[b][a].deltaTrail += dtk

	}

	//loop through the pheromone table and update each edges total trail using the evaporation constant and the curent delta trail
	for n := 0; n < len(currentMap.pheromones); n++ {
		for m := 0; m < len(currentMap.pheromones[n]); m++ {
			if m != n {
				//Updating the trail intensity(totalTrail) once the ant has completed one cycle
				//1-rho represents the evaporation of pheromones from the trail
				currentMap.pheromones[n][m].totalTrail = (rho * currentMap.pheromones[n][m].totalTrail) + currentMap.pheromones[n][m].deltaTrail

				//reset deltaTrail to zero
				currentMap.pheromones[n][m].deltaTrail = 0
			}

		}
	}

	return currentMap.pheromones
}

// on the last iteration of this cycle, this function is called and it closes the tour each ant took by moving it to the town it started at and updating the distance
//Input: current map
//Output: slice of pointers to ants with updated cur,next, and distances after traveling back to first town
func ReturnTowns(currentMap Map, Q float64, simulation string) []*Ant {

	//for each ant, set their next town to their original town, caluclate the distance from cur to next, add it to their total distance, and set cur to the first town
	for antIndex := range currentMap.ants {

		currentMap.ants[antIndex].next = currentMap.ants[antIndex].tabu[0]

		dist := currentMap.distanceMatrix[currentMap.ants[antIndex].cur.label][currentMap.ants[antIndex].next.label]

		currentMap.ants[antIndex].totalDistance += dist

		if simulation == "density" {
			currentMap.pheromones[currentMap.ants[antIndex].cur.label][currentMap.ants[antIndex].next.label].totalTrail += Q
			currentMap.pheromones[currentMap.ants[antIndex].next.label][currentMap.ants[antIndex].cur.label].totalTrail += Q
		} else if simulation == "quantity" {
			currentMap.pheromones[currentMap.ants[antIndex].cur.label][currentMap.ants[antIndex].next.label].totalTrail += (Q / dist)
			currentMap.pheromones[currentMap.ants[antIndex].next.label][currentMap.ants[antIndex].cur.label].totalTrail += (Q / dist)
		}

		currentMap.ants[antIndex].cur = currentMap.ants[antIndex].next
	}

	return currentMap.ants
}

// this function finds the shortest tour in this cycle based on each ants total distance
//Input: current map
//Output: tabu list of the ant who traveled the shortest distance
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
//Input: current map, alpha, beta
//Output: slice of pointers to ants after each ant moved to one new town
func UpdateAnts(currentMap Map, alpha, beta, Q float64, simulation string) []*Ant {
	//range through all ants and get what path each took to travel to each town
	for antIndex := range currentMap.ants {
		currentMap.ants[antIndex] = MoveAnt(currentMap.ants[antIndex], currentMap, alpha, beta, Q, simulation)
	}

	return currentMap.ants
}

// MoveAnt takes the current ant, the list of towns, the pheromone table, and system parameters and simulates one ant's
// journey across each town. It returns the updated ant object
//Input: current ant being updated, curent map, alpha, and beta
//Output: updated ant with new current and next position, updated tabu list, and updated distance traveled
func MoveAnt(currentAnt *Ant, currentMap Map, alpha, beta, Q float64, simulation string) *Ant {

	//current ant will pick the next town to visit based on probability
	currentAnt.next = PickNextTown(currentAnt, currentMap, alpha, beta)

	//calc dist from current town to next town and add to total distance
	dist := currentMap.distanceMatrix[currentAnt.cur.label][currentAnt.next.label]
	currentAnt.totalDistance += dist

	//these are used if we are testing the ant-density or the ant-quantity models
	//density adds a constant amount to each edge visited which isn't affected by the distance at all
	if simulation == "density" {
		currentMap.pheromones[currentAnt.cur.label][currentAnt.next.label].totalTrail += Q
		currentMap.pheromones[currentAnt.next.label][currentAnt.cur.label].totalTrail += Q
		//quantity adds a trail amount proportional to the length of the edge where shorter edges will have larger values
	} else if simulation == "quantity" {
		currentMap.pheromones[currentAnt.cur.label][currentAnt.next.label].totalTrail += (Q / dist)
		currentMap.pheromones[currentAnt.next.label][currentAnt.cur.label].totalTrail += (Q / dist)
	}

	//move ant to next town
	currentAnt.cur = currentAnt.next

	//now add that town to tabu list
	currentAnt.tabu = append(currentAnt.tabu, currentAnt.cur)

	return currentAnt
}

// PickNextTown takes the pheromone table and system parameters to make a weighted probability decision about what town
// to travel to next. Returns the town that will be next
//Input: Current ant, current map, alpha, beta
//Output: Pointer to the town that this ant picks based on weighted probability
func PickNextTown(currentAnt *Ant, currentMap Map, alpha, beta float64) *Town {
	//use GitHub package randutil to do weighted random selection
	choices := make([]randutil.Choice, len(currentMap.towns))

	//calculate the total weight as denominator
	totalProbability := CalculateTotalProb(currentAnt, currentMap, alpha, beta)

	//range through all possibilities of towns, excluding current town and towns already visited, and calculate probability for each
	for townIndex := range currentMap.towns {
		if currentAnt.cur.label != currentMap.towns[townIndex].label {
			choices[townIndex].Item = currentMap.towns[townIndex].label
			//if the town is in the tabu list, make it's weight 0 so it isn't chosen
			if InTabu(currentMap.towns[townIndex], currentAnt.tabu) {
				choices[townIndex].Weight = 0
			} else { //otherwise, calculate the weighted probabilty of choosing this town
				//trail value = current trail ^alpha
				trailProb := math.Pow(currentMap.pheromones[currentAnt.cur.label][currentMap.towns[townIndex].label].totalTrail, alpha)
				//visibility value = 1/distance ^beta
				distProb := math.Pow(1/(currentMap.distanceMatrix[currentAnt.cur.label][currentMap.towns[townIndex].label]), beta)
				//fmt.Println(100 * ((trailProb * distProb) / totalProbability))
				//weight = (trail weight * visibility weight) / total weight
				choices[townIndex].Weight = int(100 * ((trailProb * distProb) / totalProbability))
			}

		} else {
			//if the town is the current town, then set the weight to 0 so it isn't chosen
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

	return nextTown
}

// this function calculates the sum of all transition probabilites from the current ant to the towns that aren't in the current ant's tabu list
//Input: current ant, current map, alpha and beta
//Output: total weight of all towns that can be visited (serves as the denominator in the weighted probability formula)
func CalculateTotalProb(currentAnt *Ant, currentMap Map, alpha, beta float64) float64 {

	totalProb := 0.0

	for townIndex := range currentMap.towns {
		if currentAnt.cur.label != currentMap.towns[townIndex].label {
			if InTabu(currentMap.towns[townIndex], currentAnt.tabu) == false {
				//if the town is not in the tabu list, then calculate the trail and visibility weight and add it to the sum
				trailProb := math.Pow(currentMap.pheromones[currentAnt.cur.label][currentMap.towns[townIndex].label].totalTrail, alpha)
				distProb := math.Pow(1/(currentMap.distanceMatrix[currentAnt.cur.label][currentMap.towns[townIndex].label]), beta)
				totalProb += (trailProb * distProb)
			}

		}

	}
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

//This function computes the total distance of a tour from a slice of towns in order
//Input: current map, tabu list of towns
//Output: total distance covered by this tour
func ComputeDistance(currentMap Map, tabuList []*Town) float64 {

	totalDist := 0.0
	for i := 0; i < len(tabuList)-1; i++ {
		curTown := tabuList[i]
		nextTown := tabuList[i+1]
		//use distance matrix to quickly grab distance
		curDist := currentMap.distanceMatrix[curTown.label][nextTown.label]
		totalDist += curDist
	}

	//calculate the distance between the last town in the tabu list and the first town visited
	curTown := tabuList[len(tabuList)-1]
	nextTown := tabuList[0]
	curDist := currentMap.distanceMatrix[curTown.label][nextTown.label]
	totalDist += curDist

	return totalDist
}

//This function calculates the average shortest distances across all of the cycles
//Input: the list of timepoints, the number of cycles
//Output: average distance of all of the shortest paths
func ShortestTourAvgDist(timePoints []Map, numCycles int) float64 {
	avgDist := 0.0
	//loop through the final timepoints list of shortest tours (contains all of the shortest tours)
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

//This function calculates the average distance of all tours from one cycle
//Input: Current cycle
//Output: average tour length of all ants during this cycle
func AvgDistOfTour(timePoint Map) float64 {
	avgDist := 0.0
	for i := range timePoint.ants {
		avgDist += timePoint.ants[i].totalDistance
	}
	return avgDist / float64(len(timePoint.ants))
}

//This function prints out a tour from the starting town to the last town visited
//Input: slice of towns representing one tour
//Output: Prints out "Town __ (x position, y position) -->" for each town in the tour
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
