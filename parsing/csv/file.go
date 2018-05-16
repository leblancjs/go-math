package csv

type File [][]float64

func (f *File) Row(idx int) []float64 {
	return (*f)[idx]
}

func (f *File) Column(idx int) []float64 {
	values := make([]float64, 0, 0)
	records := *f
	for _, r := range records {
		values = append(values, r[idx])
	}
	return values
}

func (f *File) RowSize() int {
	records := *f

	rows := len(records)
	if rows <= 0 {
		return 0
	}

	rowSize := len(records[0])

	return rowSize
}
