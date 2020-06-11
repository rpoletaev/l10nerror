package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v2"
)

type Localizer interface {
	Localize(id string, lang language.Tag) string
	LocalizeParams(id string, lang language.Tag, params interface{}) string
}

type Config struct {
	// Format format of translation files
	Format string
	// Dir directory with translation files
	Dir string
}
type MyLocalizer struct {
	*Config
	locmap map[language.Tag]*i18n.Localizer
}

func NewLocalizer(c *Config) *MyLocalizer {
	l := &MyLocalizer{
		Config: c,
	}

	b := i18n.NewBundle(language.English)
	b.RegisterUnmarshalFunc(c.Format, yaml.Unmarshal)

	files, err := ioutil.ReadDir(c.Dir)
	if err != nil {
		panic(fmt.Sprint("can't read translation directory:", err))
	}

	lm := make(map[language.Tag]*i18n.Localizer, len(files))
	for _, f := range files {
		nparts := strings.Split(f.Name(), ".")
		if len(nparts) < 3 || nparts[0] != "active" {
			continue
		}

		tag, err := language.Parse(nparts[1])
		if err != nil {
			continue
		}

		b.MustLoadMessageFile(path.Join(c.Dir, f.Name()))
		lm[tag] = i18n.NewLocalizer(b, nparts[1])
	}
	l.locmap = lm
	return l
}

func (l *MyLocalizer) LocalizeParams(lang language.Tag, params TranslatableHTTPError) (string, error) {
	return l.locmap[lang].Localize(&i18n.LocalizeConfig{
		MessageID:    params.TranslateID(),
		TemplateData: params,
	})
}

func (l *MyLocalizer) Localize(id string, lang language.Tag) (string, error) {
	return l.locmap[lang].Localize(&i18n.LocalizeConfig{MessageID: id})
}

func main() {
	c := &Config{
		Dir:    "out",
		Format: "yaml",
	}

	loczr := NewLocalizer(c)
	ru, err := loczr.LocalizeParams(language.Russian, AlreadyExistsError{
		Err:    ErrAlreadyExists,
		Entity: "additional services",
	})

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(ru)

	fil, err := loczr.LocalizeParams(language.Filipino, AlreadyExistsError{
		Err:    ErrAlreadyExists,
		Entity: "additional services",
	})
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(fil)
}

var ErrAlreadyExists = errors.New("ErrAlreadyExists")

type TranslatableHTTPError interface {
	TranslateID() string
	Error() string
	Code() int
}

type AlreadyExistsError struct {
	Err    error
	Entity string
}

func (err AlreadyExistsError) Error() string {
	return err.Err.Error()
}

func (err AlreadyExistsError) Code() int {
	return http.StatusNotFound
}

func (err AlreadyExistsError) TranslateID() string {
	return err.Err.Error()
}
