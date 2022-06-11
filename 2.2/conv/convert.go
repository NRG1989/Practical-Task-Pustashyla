package conv

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"converting/lengthconvert"
	"converting/massconvert"
	"converting/tempconvert"
)

func Parse(str string) (float64, string, error) {
	r := regexp.MustCompile("^([-+]?\\d*\\.?\\d+(?:[eE][-+]?\\d+)?)(.+)$")
	result := r.FindStringSubmatch(str)

	if len(result) != 3 {
		return 0, "", fmt.Errorf("invalid format: %s", str)
	}

	value, err := strconv.ParseFloat(result[1], 64)
	unit := result[2]

	if err != nil {
		return 0, "", err
	}

	return value, unit, nil
}

func Convert(value float64, unit string) (string, string, error) {
	lowerUnit := strings.ToLower(unit)

	if lowerUnit == "ft" {
		f := lengthconvert.Foot(value)
		return f.String(), lengthconvert.FootToMeter(f).String(), nil
	} else if lowerUnit == "m" {
		m := lengthconvert.Meter(value)
		return m.String(), lengthconvert.MeterToFoot(m).String(), nil
	} else if lowerUnit == "lb" {
		p := massconvert.Pound(value)
		return p.String(), massconvert.PoundToKilogram(p).String(), nil
	} else if lowerUnit == "kg" {
		k := massconvert.Kilogram(value)
		return k.String(), massconvert.KilogramToPound(k).String(), nil
	} else if lowerUnit == "c" || lowerUnit == "°c" || lowerUnit == "℃" {
		c := tempconvert.Celsius(value)
		return c.String(), tempconvert.CToF(c).String(), nil
	} else if lowerUnit == "f" || lowerUnit == "°f" || lowerUnit == "℉" {
		f := tempconvert.Fahrenheit(value)
		return f.String(), tempconvert.FToC(f).String(), nil
	} else {
		return "", "", fmt.Errorf("unknown unit: %s", unit)
	}
}
