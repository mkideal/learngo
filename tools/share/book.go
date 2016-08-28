package share

//------
// File
//------

type File struct {
	Name     string
	Fullpath string
	Content  []byte
	CodeList []Code
}

func (file File) Run(errList *ErrorList) {
	for _, code := range file.CodeList {
		if _, err := code.Run(); err != nil {
			errList.Push(err)
		}
	}
}

//---------
// Section
//---------

type Section struct {
	Id          int
	Name        string
	Title       string
	Prefix      string
	SectionFile File
}

func (section Section) Run(errList *ErrorList) {
	section.SectionFile.Run(errList)
}

//---------
// Chapter
//---------

type Chapter struct {
	Id       int
	Name     string
	Fullpath string
	TOCFile  File
	Sections []Section
}

func (chapter Chapter) Run(errList *ErrorList) {
	chapter.TOCFile.Run(errList)
	for _, sec := range chapter.Sections {
		sec.Run(errList)
	}
}

//------
// Book
//------

type Book struct {
	Fullpath string
	Chapters []Chapter
}

func (book Book) Run() error {
	errList := &ErrorList{}
	for _, ch := range book.Chapters {
		ch.Run(errList)
	}
	return errList.Done()
}
