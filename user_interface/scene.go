package user_interface

type Scene struct {
	UI []UIInterface
}

func (S *Scene) Add(a ...UIInterface) {
	S.UI = append(S.UI, a...)
}
