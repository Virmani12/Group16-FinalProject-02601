package main

import (
	"canvas"
	"image"
)

//This function creates a slice of images visually representing each cycles output
//Input: slice of map timepoints, canvas width
//Output: slice of images for each timepoint
func AnimateSystem(timePoints []Map, canvasWidth int) []image.Image {
	images := make([]image.Image, 0)

	if len(timePoints) == 0 {
		panic("Error: no map objects present in AnimateSystem")
	}

	for i := range timePoints {
		images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
	}

	return images
}

//This function draws each map to an images by setting town points, edges between every two towns, and highlights the shortest tour in the map
//Input: current map, canvas wideth
//Output: image of the map
func DrawToCanvas(currentMap Map, canvasWidth int) image.Image {

	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	c.SetFillColor(canvas.MakeColor(255, 255, 255))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	//Set up all the edges between towns in the color black
	//Also scales the pheromone value to change the edge width based off of the amount of pheromone along the edge
	//STILL NEED TO FIGURE OUT HOW TO SCALE THIS VALUE CORRECTLY
	for _, a := range currentMap.towns {
		for _, b := range currentMap.towns {
			c.MoveTo(a.position.x, a.position.y)
			if a.label != b.label {
				//pheromoneValue := math.Log10(currentMap.pheromones[a.label][b.label].totalTrail) / 17
				c.LineTo(b.position.x, b.position.y)
				c.SetStrokeColor(canvas.MakeColor(0, 0, 0))
				c.SetFillColor(canvas.MakeColor(0, 0, 0))
				c.SetLineWidth(0.5)
				c.Stroke()
				c.FillStroke()

			}
		}
	}

	//Extract the shortest tour from this timepoint
	shortestTour := currentMap.shortestTours[len(currentMap.shortestTours)-1]

	//Highlight the edges in the shortest tour slice in green
	for i := 0; i < len(shortestTour)-1; i++ {
		c.MoveTo(shortestTour[i].position.x, shortestTour[i].position.y)
		c.LineTo(shortestTour[i+1].position.x, shortestTour[i+1].position.y)
		c.SetStrokeColor(canvas.MakeColor(0, 255, 0))
		c.SetFillColor(canvas.MakeColor(0, 255, 0))
		c.SetLineWidth(2.0)
		c.Stroke()
		c.FillStroke()
	}
	//Highlight the last edge from the last town in the slice to the first town in the slice to connect the tour
	c.MoveTo(shortestTour[len(shortestTour)-1].position.x, shortestTour[len(shortestTour)-1].position.y)
	c.LineTo(shortestTour[0].position.x, shortestTour[0].position.y)
	c.SetStrokeColor(canvas.MakeColor(0, 255, 0))
	c.SetFillColor(canvas.MakeColor(0, 255, 0))
	c.SetLineWidth(2.0)
	c.Stroke()
	c.FillStroke()

	//Draw the towns in this map in red with a radius of 5
	for _, b := range currentMap.towns {
		c.SetFillColor(canvas.MakeColor(255, 0, 0))
		c.Circle(b.position.x, b.position.y, 5.0)
		c.Fill()
	}

	return c.GetImage()

}
