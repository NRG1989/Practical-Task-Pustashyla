package massconvert

func PoundToKilogram(p Pound) Kilogram { return Kilogram(float64(p) / PoundPerKilogram) }

func KilogramToPound(k Kilogram) Pound { return Pound(float64(k) * PoundPerKilogram) }
