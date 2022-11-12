package main

import "math/rand"

// InitializeTrail creates a 2x2 table representing the pheromone intensity at every "edge" between  pairs of towns in our simulation
// Input: number of towns in this simulation, the initial trail intensity to set to every "edge"
// Output: Initial pheromone table with total trail intensity set to initialTrailIntensity
func InitializeTrail(numTowns int, initialTrailIntensity float64) PheromoneTable {

	//initialize pheromone table
	initialTrail := make(PheromoneTable, numTowns)

	for i := range initialTrail {
		initialTrail[i] = make([]*Trail, numTowns)
		for j := range initialTrail[i] {
			//set each edge to initialTrailIntensity
			initialTrail[i][j].totalTrail = initialTrailIntensity
		}
	}

	return initialTrail
}

//InitializeMap creates initial map object before the simulation runs
//Input: initialized pheromone table, number of towns, and the width of the map
//Output: Map object with these fields set, shortest distance set to longest possible distance????
//func InitializeMap(initialTrailTable PheromoneTable, numTowns, width int) Map

// initialize the positions of each town in the map
func InitializeMap(initialTrail PheromoneTable, numTowns int, width float64) Map {
	var initialMap Map
	initialMap.towns = make([]*Town, numTowns)
	// initialMap.pheromone = ??
	initialMap.width = width

	// range over the towns and create a random position within the map
	for i := range initialMap.towns {
		initialMap.towns[i].label = i
		initialMap.towns[i].position.x = rand.Float64() * initialMap.width
		initialMap.towns[i].position.y = rand.Float64() * initialMap.width
	}

	initialMap.pheromones = initialTrail
	return initialMap
}
