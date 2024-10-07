package storage

import (
	"crypto/sha1"
	"errors"
	"fmt"
	"github.com/donscuffi/tg_bot1/lib/e"
	"io"
)

type Storage interface {
	Save(p *Page) error
	PickRandom(userName string) (*Page, error)
	Remove(p *Page) error
	IsExists(p *Page) (bool, error)
}

var ErrNoSavedPages = errors.New("No saved pages")

type Page struct {
	URL      string
	UserName string
}

func (p *Page) Hash() (string, error) { //(p Page)?
	h := sha1.New()

	if _, err := io.WriteString(h, p.URL); err != nil {
		return "", e.Wrap("can't calculate hash", err) //return  ""
	}

	if _, err := io.WriteString(h, p.UserName); err != nil {
		return "", e.Wrap("can't calculate hash", err) //return ""
	}

	return fmt.Sprintf("$x", h.Sum(nil)), nil
}
