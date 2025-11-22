package piscine

func SortIntegerTable(table []int) {
	// Sorting using bubble sort algorithm
	n := len(table)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1; j++ {
			if table[j] > table[j+1] {
				swap(&table[j], &table[j+1])
			}
		}
	}
}

func swap(i, j *int) {
	*i, *j = *j, *i
}
