package lengthconvert

func FootToMeter(f Foot) Meter { return Meter(float64(f) / FootPerMeter) }

func MeterToFoot(m Meter) Foot { return Foot(float64(m) * FootPerMeter) }
