package sprint

// Point represents text at a specific point
type Point struct {
    X     float32
    Y     float32
    Text  string
}

// MakePoint creates a new Point struct with the given parameters
func MakePoint(x, y float32, text string) Point {
    return Point{
        X:    x,
        Y:    y,
        Text: text,
    }
}
