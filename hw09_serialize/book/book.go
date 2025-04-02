package book

type Book struct {
	ID     int     `json:"id" xml:"ID" yaml:"id"`
	Title  string  `json:"title" xml:"Title" yaml:"title"`
	Author string  `json:"author" xml:"Author" yaml:"author"`
	Year   int     `json:"year" xml:"Year" yaml:"year"`
	Size   int     `json:"size" xml:"Size" yaml:"size"`
	Rate   float64 `json:"rate" xml:"Rate" yaml:"rate"`
	Sample []byte  `json:"sample" xml:"Sample" yaml:"sample"`
}

type List struct {
	Books []Book `xml:"Book"`
}
