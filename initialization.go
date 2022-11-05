package main

//InitializeTrail creates a 2x2 table representing the pheromone intensity at every "edge" between  pairs of towns in our simulation
//Input: number of towns in this simulation, the initial trail intensity to set to every "edge"
//Output: Initial pheromone table with total trail intensity set to initialTrailIntensity
func InitializeTrail(numTowns int, initialTrailIntensity float64) PheromoneTable

//InitializeMap creates initial map object before the simulation runs
//Input: initialized pheromone table, number of towns, and the width of the map
//Output: Map object with these fields set, shortest distance set to longest possible distance????
func InitializeMap(initialTrailTable PheromoneTable, numTowns, width int) Map
