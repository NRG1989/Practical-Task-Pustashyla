package massconvert

import "fmt"

type Pound float64

type Kilogram float64

const (
	PoundPerKilogram float64 = 0.45359237
)

func (p Pound) String() string    { return fmt.Sprintf("%glb", p) }
func (k Kilogram) String() string { return fmt.Sprintf("%gkg", k) }
