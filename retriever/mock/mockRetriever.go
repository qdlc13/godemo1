package mock

type Retriever struct {
	Contents string
}

func (r Retriever) Get(url string) string {
	//TODO implement me
	return r.Contents
	panic("implement me")
}
