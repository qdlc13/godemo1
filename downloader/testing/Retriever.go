package testing

type Retriever struct {
}

func (r Retriever) Get(url string) string {
	return "fake content"
}
