package task

type Sequence struct {
	Num int
}

func (seq *Sequence) OddSequence() []int {
	var result []int
	for i := 1; i <= seq.Num; i += 2 {
		result = append(result, i)
	}
	return result
}

func (seq *Sequence) EvenSequence() []int {
	var result []int
	for i := 2; i <= seq.Num; i += 2 {
		result = append(result, i)
	}
	return result
}

func (seq *Sequence) PrimeSequence() []int {
	var result []int
	for i := 1; i <= seq.Num; i++ {
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
	return result
}

func (seq *Sequence) FiboSequence() []int {
	a, b := 0, 1
	var fibo int
	var result []int
	for a < seq.Num {
		fibo = a
		a = b
		b = fibo + a
		result = append(result, fibo)
	}
	return result
}
