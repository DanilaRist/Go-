package sprint

// Point represents text at a specific point
type Point struct {
	X    float64
	Y    float64
	Text string
}

// SquaredDistance calculates the squared distance of a point from the origin
func SquaredDistance(p Point) float64 {
	return p.X*p.X + p.Y*p.Y
}

// PointDiff returns the Point structure that is further ahead in the X or Y coordinates
func PointDiff(p1, p2 Point) Point {
	// Calculate the squared distances from the origin for each point
	dist1 := SquaredDistance(p1)
	dist2 := SquaredDistance(p2)

	// Compare the squared distances and return the point with greater distance
	if dist1 > dist2 {
		return p1
	} else if dist2 > dist1 {
		return p2
	}
	// If distances are equal, return either point
	return p1
}

// The main function is not needed for this implementation
