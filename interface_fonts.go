package base64Captcha

import "golang.org/x/image/font/opentype"

// FontsStorage interface for working with fonts
type FontsStorage interface {
	// LoadFontByName returns the font from the storage
	LoadFontByName(name string) *opentype.Font

	// LoadFontsByNames returns multiple fonts from storage
	LoadFontsByNames(assetFontNames []string) []*opentype.Font
}
