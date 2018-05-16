package csv

type File [][]float64

func (file File) Row(idx int) []float64 {
	return file[idx]
}

func (file File) Column(idx int) []float64 {
	values := make([]float64, 0, 0)
	for _, r := range file {
		values = append(values, r[idx])
	}
	return values
}

func (file File) RowSize() int {
	rows := len(file)
	if rows <= 0 {
		return 0
	}

	rowSize := len(file[0])

	return rowSize
}
