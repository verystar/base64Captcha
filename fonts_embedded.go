package base64Captcha

import (
	"embed"

	"golang.org/x/image/font/opentype"
)

type EmbeddedFontsStorage struct {
	fs embed.FS
}

func (s *EmbeddedFontsStorage) LoadFontByName(name string) *opentype.Font {
	fontBytes, err := s.fs.ReadFile(name)
	if err != nil {
		panic(err)
	}

	//font file bytes to opentype font
	trueTypeFont, err := opentype.Parse(fontBytes)
	if err != nil {
		panic(err)
	}

	return trueTypeFont
}

// LoadFontsByNames import fonts from dir, preserving the given order.
func (s *EmbeddedFontsStorage) LoadFontsByNames(assetFontNames []string) []*opentype.Font {
	fonts := make([]*opentype.Font, 0)
	for _, assetName := range assetFontNames {
		f := s.LoadFontByName(assetName)
		fonts = append(fonts, f)
	}
	return fonts
}

func NewEmbeddedFontsStorage(fs embed.FS) *EmbeddedFontsStorage {
	return &EmbeddedFontsStorage{
		fs: fs,
	}
}
