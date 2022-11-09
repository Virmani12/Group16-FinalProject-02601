package main

import "fmt"

// main will run our optimization problem with varying sets of parameters
func main() {

	//Below are the various parameters and initial values (all subject to change)
	alpha := 1.0
	beta := 1.0
	rho := 0.99
	Q := 100
	initialIntensity := 1.0 //should be scaled based on number of towns and Q
	numTowns := 50
	numAnts := 50
	numCycles := 1000
	width := 500

	//initialize pheromone trail from number of towns and intitial intensity
	initialTrail := InitializeTrail(numTowns, initialIntensity)

	//

	//initialize map of towns with random x,y coordinate, width, and pheromone table
	initialMap := InitializeMap(initialTrail, numTowns, width)

	//Simulate AntColony
	//Input: alpha, beta, rho, initialMap, numCycles
	//Output: Array of Maps showing the best route after each cycle (only keeping an array for visualization purposes)
	shortestMaps := AntColony(initialMap, alpha, beta, rho, numCycles, numAnts, Q)

	//animate shortest maps
	fmt.Println(shortestMaps)

}
