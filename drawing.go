package main

import (
	"canvas"
	"image"
)

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

func DrawToCanvas(currentMap Map, canvasWidth int) image.Image {

	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)

	c.SetFillColor(canvas.MakeColor(255, 255, 255))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	for _, b := range currentMap.towns {
		c.SetFillColor(canvas.MakeColor(0, 0, 0))
		c.Circle(b.position.x, b.position.y, 3.0)
		c.Fill()
	}

	for _, a := range currentMap.towns {
		for _, b := range currentMap.towns {
			c.MoveTo(a.position.x, a.position.y)
			if a.label != b.label {
				//pheromoneValue := currentMap.pheromones[a.label][b.label].totalTrail
				//fmt.Println(pheromoneValue)
				c.LineTo(b.position.x, b.position.y)
				c.SetStrokeColor(canvas.MakeColor(0, 0, 0))
				c.SetFillColor(canvas.MakeColor(0, 0, 0))
				c.SetLineWidth(0.5)
				c.Stroke()
				c.FillStroke()

			}
		}
	}

	return c.GetImage()

}
