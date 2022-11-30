package main

import "fmt"

// main will run our optimization problem with varying sets of parameters
func main() {

	//Below are the various parameters and initial values (all subject to change)
	alpha := 1.0
	beta := 1.0
	rho := 0.99
	Q := 100.0
	initialIntensity := 1.0 //should be scaled based on number of towns and Q
	numTowns := 50
	numAnts := 50
	numCycles := 1000
	width := 500.0

	//initialize pheromone trail from number of towns and intitial intensity
	initialTrail := InitializeTrail(numTowns, initialIntensity)

	//initialize map of towns with random x,y coordinate, width, and pheromone table
	// I inputted the initial values which can be read from the command line
	initialMap := InitializeMap(initialTrail, numTowns, width)

	//create distance matrix from initialMap based on town positions
	initialMap.distanceMatrix = InitializeDistanceMatrix(initialMap)

	//Simulate AntColony
	//Input: alpha, beta, rho, initialMap, numCycles
	//Output: Array of Maps showing the best route after each cycle (only keeping an array for visualization purposes)
	timePoints := AntColony(initialMap, numCycles, numAnts, alpha, beta, rho, Q)

	//animate shortest maps
	fmt.Println(timePoints)

}
