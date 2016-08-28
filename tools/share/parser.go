package share

import (
	"bufio"
	"bytes"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/mkideal/log"
)

type parser struct {
	book Book
}

func (p *parser) parseFile(filename string) File {
	file := File{
		Name:     filepath.Base(filename),
		Fullpath: filename,
	}
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return file
	}
	file.Content = content

	reader := bufio.NewReader(bytes.NewBuffer(content))
	// scan file
	const (
		gocodeStart = "```go"
		gocodeEnd   = "```"
	)
	var (
		lineno = 0
		// gocode state
		gocodeStartLineno = -1
		output            = new(bytes.Buffer)
		outputPrefix      = ""
		codeContent       = new(bytes.Buffer)
		readingCode       = false
		readingOutput     = false
		readedOutput      = false
		outputLineno      = -1
		runnable          = false
	)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		trimedLine := strings.TrimSpace(line)
		log.With(log.M{
			"readingCode":       readingCode,
			"readingOutput":     readingOutput,
			"lineno":            lineno,
			"gocodeStartLineno": gocodeStartLineno,
			"outputLineno":      outputLineno,
			"runnable":          runnable,
		}).Debug("read line: %s", trimedLine)
		if trimedLine == gocodeStart {
			// reset gocode state
			gocodeStartLineno = lineno
			output.Reset()
			codeContent.Reset()
			readingCode = true
			readingOutput = false
			readedOutput = false
			runnable = false
			outputPrefix = ""
		} else if trimedLine == gocodeEnd && readingCode {
			code := Code{
				Filename:  filename,
				StartLine: gocodeStartLineno + 1,
				EndLine:   lineno + 1,
				Content:   codeContent.Bytes(),
				Output:    output.String(),
				Runnable:  runnable,
			}
			readingCode = false
			file.CodeList = append(file.CodeList, code)
		} else if readingCode {
			if !readedOutput {
				codeContent.WriteString(line)
				if trimedLine == "// output:" {
					outputPrefix = " "
					readingOutput = true
					outputLineno = 0
					output.Reset()
					runnable = true
				} else if trimedLine == "//output:" {
					readingOutput = true
					outputLineno = 0
					output.Reset()
					runnable = true
				} else if readingOutput {
					if strings.HasPrefix(trimedLine, "//") {
						if outputLineno > 0 {
							output.WriteByte('\n')
						}
						s := strings.TrimPrefix(strings.TrimPrefix(trimedLine, "//"), outputPrefix)
						log.Debug("output: `%s`", s)
						output.WriteString(s)
						outputLineno++
					} else {
						readedOutput = true
						readingOutput = false
					}
				}
			}
		}
		lineno++
	}
	return file
}

func Parse(chapters ...int) Book {
	log.Debug("Parse: chapters=%v", chapters)
	p := new(parser)
	p.book.Fullpath = *flRootdir
	for _, chapter := range chapters {
		dir := ChapterDir(chapter)
		ch := Chapter{
			Id:       chapter,
			Name:     ChapterName(chapter),
			TOCFile:  p.parseFile(ChapterTOCFilepath(chapter)),
			Fullpath: dir,
		}
		sections, names := Sections(chapter)
		for i, section := range sections {
			ch.Sections = append(ch.Sections, Section{
				Id:          section,
				Name:        names[i],
				Prefix:      SectionFilenamePrefix(chapter, section),
				SectionFile: p.parseFile(filepath.Join(dir, names[i])),
			})
		}
		p.book.Chapters = append(p.book.Chapters, ch)
	}
	log.Trace("Parse: book=%v", p.book)
	return p.book
}
