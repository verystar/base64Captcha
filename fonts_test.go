package base64Captcha

import (
	"testing"
)

// sources:
// fonts/Flim-Flam.ttf (140.576kB)
// fonts/actionj.ttf (34.944kB)
// fonts/readme.md (162B)

func Test_loadFontByName(t *testing.T) {
	f := DefaultEmbeddedFonts.LoadFontByName("fonts/actionj.ttf")
	if f == nil {
		t.Error("failed")
	}

	defer recoverPanic(t)
	f = DefaultEmbeddedFonts.LoadFontByName("fonts/readme.md")

}
func recoverPanic(t *testing.T) {
	r := recover()
	if r == nil {
		t.Error("not trigger panic")
	}
}

func Test_loadFontsByNames(t *testing.T) {

	fs := DefaultEmbeddedFonts.LoadFontsByNames([]string{"fonts/actionj.ttf", "fonts/Flim-Flam.ttf"})
	if len(fs) != 2 {
		t.Error("failed")
	}
	defer recoverPanic(t)
	DefaultEmbeddedFonts.LoadFontsByNames([]string{"fonts/actionj.txxxxxtf"})
}

func Test_randFontFrom(t *testing.T) {
	f := randFontFrom(fontsAll)
	if f == nil {
		t.Error("failed")
	}
}
