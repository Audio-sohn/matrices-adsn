package matrixtools

import "matrices/arraytools"

// GetRow liefert die i-te Zeile der Matrix m.
func GetRow(m [][]float64, i int) []float64 {

	return m[i]
}

// GetCol liefert die j-te Spalte der Matrix m.
func GetCol(m [][]float64, j int) []float64 {
	col := make([]float64, len(m))

	for i, _ := range m {

		col[i] = m[i][j]

	}
	// TODO
	return col
}

// AddRows erwartet eine Matrix und zwei Zeilennummern.
// Addiert die beiden Zeilen paarweise und speichert das Ergebnis in der ersten Zeile.
func AddRows(m [][]float64, i, j int) {

	m[0] = arraytools.Add_return(m[i], m[j])

}

// ScalarMultRow erwartet eine Matrix, eine Zeilennummer und einen skalaren Faktor.
// Multipliziert die Zeile mit dem Faktor und speichert das Ergebnis in der Zeile.
func ScalarMultRow(m [][]float64, i int, factor float64) {

	m[i] = arraytools.ScalarMult_return(m[i], factor)

}

// Transpose erwartet eine Matrix und liefert ihre Transponierte.
// D.h. alle Zeilen der ersten Matrix werden zu Spalten der Transponierten und umgekehrt.
func Transposed(m [][]float64) [][]float64 {

	transposed := make([][]float64, len(m[0]))

	for i := range transposed {

		transposed[i] = make([]float64, len(m))

	}

	for i := range transposed {

		for j := range transposed[0] {

			transposed[i][j] = m[j][i]

		}

	}

	return transposed
}
