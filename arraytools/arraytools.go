package arraytools

// Mult erwartet ein Array und einen skalaren Faktor.
// Multipliziert jedes Element des Arrays mit dem Faktor.
func ScalarMult(a []float64, factor float64) {

	for i, speci := range a {

		a[i] = speci * factor

	}
}

func ScalarMult_return(a []float64, factor float64) []float64 {

	temp := make([]float64, len(a))

	for i, speci := range a {

		temp[i] = speci * factor

	}

	return temp
}

// Add erwartet zwei Arrays der gleichen Länge.
// Addiert die Elemente der Arrays paarweise.
func Add(a, b []float64) []float64 {

	for i, _ := range a {

		a[i] += b[i]

	}

	return a
}

func Add_return(a, b []float64) []float64 {

	temp := make([]float64, len(a))

	for i, _ := range a {

		temp[i] = a[i] + b[i]

	}

	return temp

}
