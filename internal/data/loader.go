package data

import (
	"github.com/bymarkone/timelanguage/internal/language"
	"github.com/bymarkone/timelanguage/internal/planning"
	"github.com/bymarkone/timelanguage/internal/purpose"
	"github.com/bymarkone/timelanguage/internal/schedule"
	"log"
	"os"
	"strings"
)

type Loader struct {
	BaseFolder string
	loaded     map[string]string
}

func (l *Loader) Load() {
	planning.CreateRepository()
	schedule.CreateRepository()
	purpose.CreateRepository()
	l.loaded = make(map[string]string)

	filesInfo, err := os.ReadDir(l.BaseFolder)
	if err != nil {
		log.Fatal(err)
		return
	}
	for _, file := range filesInfo {
		if !strings.HasSuffix(file.Name(), ".gr") {
			continue
		}
		fileAddress := l.BaseFolder + "/" + file.Name()
		content, err := os.ReadFile(fileAddress)
		if err != nil {
			log.Fatal(err)
			return
		}
		context := strings.ReplaceAll(file.Name(), ".gr", "")
		l.loaded[context] = fileAddress
		text := string(content)
		l := language.NewLexer(text)
		p := language.NewParser(file.Name(), l)
		categories, items := p.Parse()
		language.Eval(context, categories, items)
	}
}
