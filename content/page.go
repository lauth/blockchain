package content

import (
	"fmt"

	"github.com/ponzu-cms/ponzu/management/editor"
	"github.com/ponzu-cms/ponzu/system/item"
)

type Page struct {
	item.Item

	Title string `json:"title"`
	Body  string `json:"body"`
}

// MarshalEditor writes a buffer of html to edit a Page within the CMS
// and implements editor.Editable
func (p *Page) MarshalEditor() ([]byte, error) {
	view, err := editor.Form(p,
		// Take note that the first argument to these Input-like functions
		// is the string version of each Page field, and must follow
		// this pattern for auto-decoding and auto-encoding reasons:
		editor.Field{
			View: editor.Input("Title", p, map[string]string{
				"label":       "Title",
				"type":        "text",
				"placeholder": "Enter the Title here",
			}),
		},
		editor.Field{
			View: editor.Richtext("Body", p, map[string]string{
				"label":       "Body",
				"placeholder": "Enter the Body here",
			}),
		},
	)

	if err != nil {
		return nil, fmt.Errorf("Failed to render Page editor view: %s", err.Error())
	}

	return view, nil
}

func init() {
	item.Types["Page"] = func() interface{} { return new(Page) }
}

// String defines how a Page is printed. Update it using more descriptive
// fields from the Page struct type
func (p *Page) String() string {
	return fmt.Sprintf("Page: %s", p.UUID)
}
