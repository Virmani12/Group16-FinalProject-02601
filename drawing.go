package main

import (
	"canvas"
	"image"
	"math"
)

func AnimateSystem(timePoints []Map, canvasWidth int) []image.Image {
	images := make([]image.Image, 0)

	if len(timePoints) == 0 {
		panic("Error: no map objects present in AnimateSystem")
	}
	/*
		for i := range timePoints {
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth))
		}
	*/
	for i := 0; i < len(timePoints); i++ {
		if i < len(timePoints)-1 {
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth, 0))
		} else {
			images = append(images, DrawToCanvas(timePoints[i], canvasWidth, 1))
		}
	}

	return images
}

func DrawToCanvas(currentMap Map, canvasWidth int, lastPos int) image.Image {

	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	c.SetFillColor(canvas.MakeColor(255, 255, 255))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	for _, a := range currentMap.towns {
		for _, b := range currentMap.towns {
			c.MoveTo(a.position.x, a.position.y)
			if a.label != b.label {
				pheromoneValue := math.Log10(currentMap.pheromones[a.label][b.label].totalTrail) / 17
				//fmt.Println(pheromoneValue)
				c.LineTo(b.position.x, b.position.y)
				c.SetStrokeColor(canvas.MakeColor(0, 0, 255))
				c.SetFillColor(canvas.MakeColor(0, 0, 255))
				c.SetLineWidth(pheromoneValue)
				c.Stroke()
				c.FillStroke()

			}
		}
	}

	shortestTour := currentMap.shortestTours[len(currentMap.shortestTours)-1]

	for i := 0; i < len(shortestTour)-1; i++ {
		c.MoveTo(shortestTour[i].position.x, shortestTour[i].position.y)
		c.LineTo(shortestTour[i+1].position.x, shortestTour[i+1].position.y)
		c.SetStrokeColor(canvas.MakeColor(0, 255, 0))
		c.SetFillColor(canvas.MakeColor(0, 255, 0))
		c.SetLineWidth(2.0)
		c.Stroke()
		c.FillStroke()
	}

	c.MoveTo(shortestTour[len(shortestTour)-1].position.x, shortestTour[len(shortestTour)-1].position.y)
	c.LineTo(shortestTour[0].position.x, shortestTour[0].position.y)
	c.SetStrokeColor(canvas.MakeColor(0, 255, 0))
	c.SetFillColor(canvas.MakeColor(0, 255, 0))
	c.SetLineWidth(2.0)
	c.Stroke()
	c.FillStroke()

	for _, b := range currentMap.towns {
		c.SetFillColor(canvas.MakeColor(0, 0, 0))
		c.Circle(b.position.x, b.position.y, 5.0)
		c.Fill()
	}

	return c.GetImage()

}
