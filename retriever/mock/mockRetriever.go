package mock

type Retriever struct {
	Contents string
}

// 值接收者
func (r Retriever) Get(url string) string {
	//TODO implement me
	return r.Contents
	panic("implement me")
}
