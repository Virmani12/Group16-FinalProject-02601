package main

import (
	"fmt"
	"gifhelper"
)

// main will run our optimization problem with varying sets of parameters
func main() {

	//Below are the various parameters and initial values (all subject to change)
	alpha := 1.0
	beta := 5.0
	rho := 0.50
	Q := 50.0
	initialIntensity := 0.01 //should be scaled based on number of towns and Q
	numTowns := 30
	numAnts := 30
	numCycles := 100
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

	avgDist := CalculateAvgDist(timePoints, numCycles)
	fmt.Println("average dist: ", avgDist)

	lastTourLength := ComputeDistance(timePoints[numCycles-1], timePoints[numCycles-1].shortestTours[numCycles-1])
	fmt.Println("last tour distance: ", lastTourLength)

	PrintTour(timePoints[numCycles-1].shortestTours[numCycles-1])

	imageList := AnimateSystem(timePoints, int(width))

	fmt.Println("drawing images")
	gifhelper.ImagesToGIF(imageList, "ants")
	fmt.Println("GIF drawn")

}
