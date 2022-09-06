package task

type Sequence struct {
	Limit int
}

func (seq *Sequence) OddSequence() (result []int) {
	for i := 1; i <= seq.Limit; i += 2 {
		result = append(result, i)
	}
	return
}

func (seq *Sequence) EvenSequence() (result []int) {
	for i := 2; i <= seq.Limit; i += 2 {
		result = append(result, i)
	}
	return
}

func (seq *Sequence) PrimeSequence() (result []int) {
	for i := 1; i <= seq.Limit; i++ {
		check := 0
		for divider := 1; divider <= i; divider++ {
			if i%divider == 0 {
				check++
			}
		}
		if check == 2 {
			result = append(result, i)
		}
	}
	return
}

func (seq *Sequence) FiboSequence() (result []int) {
	a, b := 0, 1
	var fibo int
	for a < seq.Limit {
		fibo = a
		a = b
		b = fibo + a
		result = append(result, fibo)
	}
	return
}
