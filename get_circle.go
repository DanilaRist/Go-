package sprint

// Circle represents a circle with various properties
type Circle struct {
    Radius    float32
    Diameter  float32
    Area      float32
    Perimeter float32
}

// GetCircle returns a new Circle struct with fields filled based on the given radius
func GetCircle(r float32) Circle {
    diameter := 2 * r
    area := 3.14 * r * r
    perimeter := 2 * 3.14 * r

    return Circle{
        Radius:    r,
        Diameter:  diameter,
        Area:      area,
        Perimeter: perimeter,
    }
}

func main() {
    // Example usage
    // You can use the GetCircle function and assign its result to a variable like this:
    // c := GetCircle(3.0)
    // Or you can simply call the function without assigning the result to any variable:
    // GetCircle(3.0)
}
