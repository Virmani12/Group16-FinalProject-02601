package AntAlgorithm

import (
	"encoding/csv"
	"fmt"
	"gifhelper"
	"log"
	"os"
	"strconv"
)

// main will run our optimization problem with varying sets of parameters
func main() {

	//Below are the various parameters and initial values (all subject to change)
	alpha := 1.0
	beta := 5.0
	rho := 0.50
	Q := 500.0
	initialIntensity := 0.01 //should be scaled based on number of towns and Q
	numTowns := 30
	numAnts := 30
	numCycles := 1000
	width := 500.0
	useOliver := true

	//initialize pheromone trail from number of towns and intitial intensity
	initialTrail := InitializeTrail(numTowns, initialIntensity)

	//initialize map of towns with random x,y coordinate, width, and pheromone table
	// I inputted the initial values which can be read from the command line
	initialMap := InitializeMap(initialTrail, numTowns, width, useOliver)

	//create distance matrix from initialMap based on town positions
	// moved this initialization within the initializeMap function
	//initialMap.distanceMatrix = InitializeDistanceMatrix(initialMap)

	//Simulate AntColony
	//Input: alpha, beta, rho, initialMap, numCycles
	//Output: Array of Maps showing the best route after each cycle (only keeping an array for visualization purposes)
	timePoints := AntColony(initialMap, numCycles, numAnts, alpha, beta, rho, Q)

	//calculate the average of the shortest distances found in each cycle
	//and print out
	avgDist := ShortestTourAvgDist(timePoints, numCycles)
	fmt.Println("average shortest distances: ", avgDist)

	fmt.Println(len(timePoints[numCycles-1].shortestTours))

	//Printing out the distance of the last tour
	lastTourLength := ComputeDistance(timePoints[numCycles-1], timePoints[numCycles-1].shortestTours[numCycles-1])
	fmt.Println("last tour distance: ", lastTourLength)

	//printing out the last tour
	PrintTour(timePoints[numCycles-1].shortestTours[numCycles-1])

	//animate our timepoints
	imageList := AnimateSystem(timePoints, int(width))

	//convert to gif
	fmt.Println("drawing images")
	gifhelper.ImagesToGIF(imageList, "ants")
	fmt.Println("GIF drawn")

	//exporting shortest tours to csv for analysis in R
	shortestTour := timePoints[numCycles-1].shortestTours

	//create csv file for shortest tours
	csvFile, err := os.Create("shortestTours.csv")

	if err != nil {
		log.Fatalf("failed creaing file: %s", err)
	}

	//start a buffered writer
	csvwriter := csv.NewWriter(csvFile)

	//for each index in slice of shortest tours, find the distance of the tour and append it to a csv with the index it pertains to
	for i := 0; i < len(shortestTour); i++ {
		curShortestTour := ComputeDistance(timePoints[numCycles-1], shortestTour[i])
		stringPos := strconv.Itoa(i)
		stringDist := fmt.Sprintf("%f", curShortestTour)
		var row []string
		row = append(row, stringPos)
		row = append(row, stringDist)

		csvwriter.Write(row)

	}

	//flush and close
	csvwriter.Flush()
	csvFile.Close()

	//exporting each cycles average tour length to csv for analysis in R
	csvFile, err = os.Create("averageCycleTourLength.csv")

	if err != nil {
		log.Fatalf("failed creating file: %s", err)
	}

	csvwriter = csv.NewWriter(csvFile)

	for i := 0; i < len(timePoints); i++ {
		avgDist := AvgDistOfTour(timePoints[i])
		stringPos := strconv.Itoa(i)
		stringDist := fmt.Sprintf("%f", avgDist)
		var row []string
		row = append(row, stringPos)
		row = append(row, stringDist)

		csvwriter.Write(row)

	}

	csvwriter.Flush()
	csvFile.Close()
}
