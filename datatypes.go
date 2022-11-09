package main

//Ants contains a tabu list to keep track of all towns met so far along the ant's tour
// Also contains a current and next town to represent each step taken
// Lastly contains a float keeping track of the distance covered so far along the ant's tour
type Ant struct {
	tabu          []*Town
	cur, next     Town
	totalDistance float64
}

//Towns have an integer label and an x and y coordinate to be represented on a map
type Town struct {
	label int
	x, y  float64
}

//Trails keep track of the total trail intensity as well as the change in trail intensity after one cycle between a pair of towns
type Trail struct {
	totalTrail, deltaTrail float64
}

//Map contains all the towns, the width of the map, the pheromone trail between every pair of towns, and the shortest distance found
//Used for visual representation
type Map struct {
	towns      []*Town
	pheromones PheromoneTable
	width      float64
	antArray   []*Ant
}

//PheromoneTable is a matrix containing the pheromone trail intensity between every pair of towns
type PheromoneTable [][]*Trail
