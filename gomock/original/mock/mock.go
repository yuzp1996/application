package mock

type MinimalTask struct {
	Name string
}

func NewMinimalTask() *MinimalTask {
	return &MinimalTask{}
}

func (mt *MinimalTask) Do(idx int) (string, error) {
	return "", nil
}
