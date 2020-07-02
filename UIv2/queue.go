package UIv2

type SpriteQueue struct {
	arr []basicInterface
}

func (Q *SpriteQueue) B() bool {
	return len(Q.arr) != 0
}

func (Q *SpriteQueue) Add(v ...basicInterface) {
	Q.arr = append(Q.arr, v...)
}

func (Q *SpriteQueue) Pop(i int) {
	l := len(Q.arr) - 1
	Q.arr[i] = Q.arr[l]
	Q.arr = Q.arr[:l]
}
