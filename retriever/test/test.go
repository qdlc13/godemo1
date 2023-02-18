package test

type Rr struct {
	Contents string
}

func (r *Rr) Get(url string) string {
	return r.Contents
}
