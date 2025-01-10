package matrix

import "matrices/arraytools"

// Normalize erwartet eine Spaltennummer.
// Falls das Diagonalelement [col][col] nicht 0 ist, wird die Zeile durch das Diagonalelement normiert.
// D.h. die gesamte Zeile col wird durch das Diagonalelement geteilt.
func (m Matrix) Normalize(col int) {

	//catch error
	if m[col][col] == 0 {
		return
	}

	//do division, write result into row
	m[col] = arraytools.ScalarMult_return(m[col], 1/m[col][col])

}

// EliminateBelow erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen unter der Zeile row mit -1/Matrix[row][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element unter dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateBelow(row int) {

	col := row
	//catchen falls nicht normiert
	if m[row][row] != 1 {

		return

	}

	//durch restliche zeilen iterieren
	for i := row + 1; i < len(m); i++ {

		m[i] = arraytools.Add_return(m[row], arraytools.ScalarMult_return(m[i], -1/m[i][col]))

	}

}

// EliminateAbove erwartet eine Zeilennummer `row`.
// Multipliziert alle Zeilen über der Zeile row mit -1/Matrix[row][row] und addiert sie zur Zeile row.
// Dadurch wird jeweils das Element über dem Diagonalelement 0.
// Voraussetzung: Die Zeile row ist bereits normiert, d.h. das Diagonalelement ist 1.
func (m Matrix) EliminateAbove(row int) {

	col := row
	//catchen falls nicht normiert
	if m[row][row] != 1 {

		return

	}

	//durch restliche zeilen iterieren
	for i := row - 1; i >= 0; i-- {

		m[i] = arraytools.Add_return(m[row], arraytools.ScalarMult_return(m[i], -1/m[i][col]))

	}

}

// UpperTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine obere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) UpperTriangular() {

	for i := range m {

		//normalize diagonal element
		if m[i][i] != 1 {

			m.Normalize(i)

		}

		m.EliminateBelow(i)

	}

}

// LowerTriangular führt die Gauß-Elimination für alle Zeilen der Matrix durch.
// So entsteht im linken Bereich eine untere Dreiecksmatrix, bei der die Diagonalelemente 1 sind.
func (m Matrix) LowerTriangular() {
	// TODO
}

// Gauss transformiert die Matrix im linken Bereich in die Einheitsmatrix.
func (m Matrix) Gauss() {
	// TODO
}
