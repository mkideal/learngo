package share

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/mkideal/log"
	"github.com/mkideal/log/logger"
)

var (
	flRootdir  = flag.String("dir", "./", "Root directory")
	flLogLevel = flag.String("log-level", "info", "Log level")

	charset = []byte("ABCDEFGHIJKLMNOPQRSATUVWXYZabcdefghijklmnopqrsatuvwxyz")

	ErrChapterOutOfRange = errors.New("chapter out of range: must be in interval [1,99]")
	ErrSectionOutOfRange = errors.New("section out of range: must be in interval [1,99]")
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RandString(n int) string {
	s := ""
	for i := 0; i < n; i++ {
		c := charset[rand.Intn(len(charset))]
		s += string(c)
	}
	return s
}

func RootDir() string {
	return *flRootdir
}

func LogLevel() logger.Level {
	level, _ := logger.ParseLevel(*flLogLevel)
	return level
}

func CheckChapterNumber(chapter int) error {
	if chapter <= 0 || chapter >= 100 {
		return ErrChapterOutOfRange
	}
	return nil
}

func CheckSectionNumber(section int) error {
	if section <= 0 || section >= 100 {
		return ErrSectionOutOfRange
	}
	return nil
}

func ChapterName(chapter int) string {
	return fmt.Sprintf("ch%02d", chapter)
}

func ChapterDir(chapter int) string {
	return filepath.Join(RootDir(), ChapterName(chapter))
}

func ChapterTOCFilepath(chapter int) string {
	chapterName := ChapterName(chapter)
	return filepath.Join(RootDir(), chapterName, chapterName+".md")
}

func SectionFilenamePrefix(chapter, section int) string {
	chapterName := ChapterName(chapter)
	return fmt.Sprintf("%s-%2d-", chapterName, section)
}

func SectionFilepath(chapter, section int, sectionName string) string {
	chapterName := ChapterName(chapter)
	return filepath.Join(RootDir(), chapterName, fmt.Sprintf("%s-%2d-%s.md", chapterName, section, sectionName))
}

func ChapterExist(chapter int) bool {
	_, err := os.Stat(ChapterDir(chapter))
	return err == nil
}

func SectionExist(chapter, section int) bool {
	chapterName := ChapterName(chapter)
	pattner := filepath.Join(RootDir(), chapterName, fmt.Sprintf("%s-%2d-*.md", chapterName, section))
	filepath.Glob(pattner)
	return false
}

func CreateChapter(chapter int) (exist bool, err error) {
	if err := CheckChapterNumber(chapter); err != nil {
		return false, err
	}
	if ChapterExist(chapter) {
		return true, nil
	}
	return false, os.MkdirAll(ChapterDir(chapter), 0755)
}

type ErrorList struct {
	errs []error
}

func (e *ErrorList) Done() error {
	if len(e.errs) == 0 {
		return nil
	}
	return e
}

func (e *ErrorList) Push(err error) {
	if err != nil {
		e.errs = append(e.errs, err)
	}
}

func (e *ErrorList) Error() string {
	b := new(bytes.Buffer)
	for i, err := range e.errs {
		if i != 0 {
			b.WriteByte('\n')
		}
		b.WriteString(err.Error())
	}
	return b.String()
}

func parseChapterFromDirname(dirname string) (chapter int, ok bool) {
	if len(dirname) == 4 && dirname[:2] == "ch" {
		v, err := strconv.ParseInt(dirname[2:], 10, 64)
		if err == nil && v > 0 && v < 100 {
			return int(v), true
		}
	}
	return 0, false
}

var sectionFilenameRegexp = regexp.MustCompile("^([0-9]{2})-.*\\.md$")

func parseSectionFromFilename(filename string) (section int, ok bool) {
	if filename == "README.md" {
		return 0, true
	}
	ret := sectionFilenameRegexp.FindStringSubmatch(filename)
	if len(ret) == 2 {
		sectionId, err := strconv.Atoi(ret[1])
		if err == nil && sectionId > 0 && sectionId < 100 {
			return int(sectionId), true
		}
	}
	return 0, false
}

func Chapters() []int {
	chapters := []int{}
	filepath.Walk(RootDir(), func(path string, info os.FileInfo, err error) error {
		if info == nil || err != nil {
			return filepath.SkipDir
		}
		if info.IsDir() {
			chapter, ok := parseChapterFromDirname(info.Name())
			if ok {
				chapters = append(chapters, chapter)
			}
			if RootDir() != path {
				return filepath.SkipDir
			}
		}
		return nil
	})
	return chapters
}

func Sections(chapter int) (sections []int, names []string) {
	chapterDir := ChapterDir(chapter)
	filepath.Walk(chapterDir, func(path string, info os.FileInfo, err error) error {
		log.Debug("Sections: path=%v", path)
		if info == nil || err != nil {
			return filepath.SkipDir
		}
		if info.IsDir() && path != chapterDir {
			return filepath.SkipDir
		}
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}
		section, ok := parseSectionFromFilename(info.Name())
		if ok {
			sections = append(sections, section)
			names = append(names, info.Name())
		}
		return nil
	})
	log.Debug("Sections: sections=%v,names=%v", sections, names)
	return
}
