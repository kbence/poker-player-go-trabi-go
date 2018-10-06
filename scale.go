package main

type FloatScale struct {
	DomainMin float64
	DomainMax float64
	RangeMin  float64
	RangeMax  float64
}

func (s FloatScale) Domain(min, max float64) FloatScale {
	s.DomainMin = min
	s.DomainMax = max

	return s
}

func (s FloatScale) Range(min, max float64) FloatScale {
	s.RangeMin = min
	s.RangeMax = max

	return s
}

func (s FloatScale) Scale(value float64) float64 {
	return s.RangeMin + (s.RangeMax-s.RangeMin)*(value-s.DomainMin)/(s.DomainMax-s.DomainMin)
}
