package commons

type BuiltinTypes struct{}

func (BuiltinTypes) Map(i int, reply *map[int]int) error {
	(*reply)[i] = i
	return nil
}
func (BuiltinTypes) Slice(i int, reply *[]int) error {
	*reply = append(*reply, i)
	for i, v := range *reply {
		println(i, v)
	}
	return nil
}
func (BuiltinTypes) Array(i int, reply *[1]int) error {
	(*reply)[0] = 1
	return nil
}
