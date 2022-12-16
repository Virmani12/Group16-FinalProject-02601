package main

import (
	"encoding/csv"
	"fmt"
	"gifhelper"
	"log"
	"os"
	"strconv"
	"strings"
)

// main will run our optimization problem with varying sets of parameters
func main() {

	var simulationType int
	var townSet int
	var alpha float64
	var beta float64
	var rho float64

	fmt.Print("Select which simulation you would like to run: '1' for Ant-Cycle, '2' for Ant-Density, '3' for Ant-Quantity... ")
	fmt.Scan(&simulationType)
	fmt.Println("")
	fmt.Print("Would you like to use a validation set or a random set of 30 towns? Type '1' for Oliver30 dataset, '2' for random town... ")
	fmt.Scan(&townSet)
	fmt.Println("")
	fmt.Print("Now time to select parameters to adjust the algorithm performance...")
	fmt.Print("Please provide a float64 value for the ALPHA parameter (importance towards pheromone): ")
	fmt.Scan(&alpha)
	fmt.Println("")
	fmt.Print("Please provide a float64 value for the BETA parameter (importance towards distance): ")
	fmt.Scan(&beta)
	fmt.Println("")
	fmt.Print("Finally, please provide a float64 value for the RHO parameter (pheromone decay rate): ")
	fmt.Scan(&rho)
	fmt.Println("")

	if townSet == 1 {
		fmt.Printf("Running simulation %d with the validation set...alpha=%f, beta=%f, rho=%f", simulationType, alpha, beta, rho)
	} else {
		fmt.Printf("Running simulation %d with a random set of towns...alpha=%f, beta=%f, rho=%f", simulationType, alpha, beta, rho)
	}
	fmt.Println("")

	var timePoints []Map
	numCycles := 1000
	numTowns := 30
	numAnts := 30
	initialIntensity := 0.01 //should be scaled based on number of towns and Q
	width := 500.0

	if simulationType == 1 {
		//Below are the various parameters and initial values (all subject to change)

		Q := 500.0
		var useOliver bool

		if townSet == 1 {
			useOliver = true
		} else {
			useOliver = false
		}

		//initialize pheromone trail from number of towns and intitial intensity
		initialTrail := InitializeTrail(numTowns, initialIntensity)

		//initialize map of towns with random x,y coordinate, width, and pheromone table
		// I inputted the initial values which can be read from the command line
		initialMap := InitializeMap(initialTrail, numTowns, width, useOliver)

		//create distance matrix from initialMap based on town positions
		initialMap.distanceMatrix = InitializeDistanceMatrix(initialMap)

		//Simulate AntColony
		//Input: alpha, beta, rho, initialMap, numCycles
		//Output: Array of Maps showing the best route after each cycle (only keeping an array for visualization purposes)
		timePoints = AntCycle(initialMap, numCycles, numAnts, alpha, beta, rho, Q)

	} else if simulationType == 2 {

		Q := 500.0

		var useOliver bool

		if townSet == 1 {
			useOliver = true
		} else {
			useOliver = false
		}

		//initialize pheromone trail from number of towns and intitial intensity
		initialTrail := InitializeTrail(numTowns, initialIntensity)

		//initialize map of towns with random x,y coordinate, width, and pheromone table
		// I inputted the initial values which can be read from the command line
		initialMap := InitializeMap(initialTrail, numTowns, width, useOliver)

		//create distance matrix from initialMap based on town positions
		initialMap.distanceMatrix = InitializeDistanceMatrix(initialMap)

		//Simulate AntColony
		//Input: alpha, beta, rho, initialMap, numCycles
		//Output: Array of Maps showing the best route after each cycle (only keeping an array for visualization purposes)
		timePoints = AntDensity(initialMap, numCycles, numAnts, alpha, beta, rho, Q)

	} else if simulationType == 3 {

		Q := 500.0

		var useOliver bool

		if townSet == 1 {
			useOliver = true
		} else {
			useOliver = false
		}

		//initialize pheromone trail from number of towns and intitial intensity
		initialTrail := InitializeTrail(numTowns, initialIntensity)

		//initialize map of towns with random x,y coordinate, width, and pheromone table
		// I inputted the initial values which can be read from the command line
		initialMap := InitializeMap(initialTrail, numTowns, width, useOliver)

		//create distance matrix from initialMap based on town positions
		initialMap.distanceMatrix = InitializeDistanceMatrix(initialMap)

		//Simulate AntColony
		//Input: alpha, beta, rho, initialMap, numCycles
		//Output: Array of Maps showing the best route after each cycle (only keeping an array for visualization purposes)
		timePoints = AntQuantity(initialMap, numCycles, numAnts, alpha, beta, rho, Q)

	}

	shortestDist := ComputeDistance(timePoints[numCycles-1], timePoints[numCycles-1].shortestTours[0])
	for i := 1; i < len(timePoints[numCycles-1].shortestTours); i++ {

		currDist := ComputeDistance(timePoints[numCycles-1], timePoints[numCycles-1].shortestTours[i])

		if currDist < shortestDist {
			shortestDist = currDist
		}
	}
	fmt.Println("Shortest Distance found: ", shortestDist)

	//calculate the average of the shortest distances found in each cycle
	//and print out
	avgDist := ShortestTourAvgDist(timePoints, numCycles)
	fmt.Println("average shortest distances: ", avgDist)

	//Printing out the distance of the last tour
	lastTourLength := ComputeDistance(timePoints[numCycles-1], timePoints[numCycles-1].shortestTours[numCycles-1])
	fmt.Println("last tour distance: ", lastTourLength)

	totalAverageLength := 0.0
	for i := 0; i < len(timePoints); i++ {
		currAvg := AvgDistOfTour(timePoints[i])
		totalAverageLength += currAvg
	}
	totalAverageLength = totalAverageLength / float64(len(timePoints))
	fmt.Println("Average length of tours: ", totalAverageLength)

	//printing out the last tour
	PrintTour(timePoints[numCycles-1].shortestTours[numCycles-1])

	//animate our timepoints
	imageList := AnimateSystem(timePoints, int(width))

	//convert to gif
	fmt.Println("drawing images")

	//setting gif file name based on algorithm type and parameters
	var gifName string

	//a,b,r = string version of alpha,beta,rho (0.9 --> "0_9")
	a := fmt.Sprintf("%.2f", alpha)
	a = strings.Replace(a, ".", "_", 1)

	b := fmt.Sprintf("%.2f", beta)
	b = strings.Replace(b, ".", "_", 1)

	r := fmt.Sprintf("%.2f", rho)
	r = strings.Replace(r, ".", "_", 1)

	if simulationType == 1 {
		gifName = fmt.Sprintf("ant-cycle-a%s-b%s-r%s", a, b, r)

	} else if simulationType == 2 {
		gifName = fmt.Sprintf("ant-density-a%s-b%s-r%s", a, b, r)
	} else {
		gifName = fmt.Sprintf("ant-quantity-a%s-b%s-r%s", a, b, r)
	}
	gifhelper.ImagesToGIF(imageList, gifName)

	fmt.Println("GIF drawn")

	//////////////////////////////////////////////////////////////
	//ANALYSIS////////////////////////
	fmt.Print("Would you like to save analysis data? Type '1' for yes and '2' for no: ")
	var analysis int
	fmt.Scan(&analysis)

	if analysis == 2 {
		fmt.Println("Exiting...")
		os.Exit(0)
	}

	//exporting shortest tours to csv for analysis in R
	shortestTour := timePoints[numCycles-1].shortestTours
	//create csv file for shortest tours
	var shortestTourFileName string
	if simulationType == 1 {
		shortestTourFileName = fmt.Sprintf("shortestToursAC-a%s-b%s-r%s.csv", a, b, r)
	} else if simulationType == 2 {
		shortestTourFileName = fmt.Sprintf("shortestToursAD-a%s-b%s-r%s.csv", a, b, r)
	} else {
		shortestTourFileName = fmt.Sprintf("shortestToursAQ-a%s-b%s-r%s.csv", a, b, r)
	}
	csvFile, err := os.Create(shortestTourFileName)

	if err != nil {
		log.Fatalf("failed creating shortest tours file: %s", err)
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
	var averageTourFileName string
	if simulationType == 1 {
		averageTourFileName = fmt.Sprintf("averageCycleTourLengthAC-a%s-b%s-r%s.csv", a, b, r)
	} else if simulationType == 2 {
		averageTourFileName = fmt.Sprintf("averageCycleTourLengthAD-a%s-b%s-r%s.csv", a, b, r)
	} else {
		averageTourFileName = fmt.Sprintf("averageCycleTourLengthAQ-a%s-b%s-r%s.csv", a, b, r)
	}
	csvFile, err = os.Create(averageTourFileName)

	if err != nil {
		log.Fatalf("failed creating average cycle tour length file: %s", err)
	}

	//start a new buffered writer
	csvwriter = csv.NewWriter(csvFile)

	//for each time point, find the average distance of all the tours in the cycle and append it to a csv with the index it pertains to
	for i := 0; i < len(timePoints); i++ {
		avgDist := AvgDistOfTour(timePoints[i])
		stringPos := strconv.Itoa(i)
		stringDist := fmt.Sprintf("%f", avgDist)
		var row []string
		row = append(row, stringPos)
		row = append(row, stringDist)

		csvwriter.Write(row)

	}

	//flush and close
	csvwriter.Flush()
	csvFile.Close()
}
