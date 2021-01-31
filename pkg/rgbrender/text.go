package rgbrender

import (
	"fmt"
	"image"
	"image/color"
	"io/ioutil"
	"math"

	"github.com/golang/freetype"
	"github.com/golang/freetype/truetype"

	"github.com/markbates/pkger"
	"golang.org/x/image/font"
	"golang.org/x/image/math/fixed"

	rgb "github.com/robbydyer/sports/pkg/rgbmatrix-rpi"
)

// TextWriter ...
type TextWriter struct {
	context          *freetype.Context
	font             *truetype.Font
	XStartCorrection int
	YStartCorrection int
	FontSize         float64
	LineSpace        float64
}

// DefaultTextWriter ...
func DefaultTextWriter() (*TextWriter, error) {
	fnt, err := DefaultFont()
	if err != nil {
		return nil, err
	}

	t := NewTextWriter(fnt, 8)
	t.YStartCorrection = -2

	return t, nil
}

// NewTextWriter ...
func NewTextWriter(font *truetype.Font, fontSize float64) *TextWriter {
	cntx := freetype.NewContext()
	cntx.SetFont(font)
	cntx.SetFontSize(fontSize)

	return &TextWriter{
		context:   cntx,
		font:      font,
		FontSize:  fontSize,
		LineSpace: 0.5,
	}
}

// DefaultFont ...
func DefaultFont() (*truetype.Font, error) {
	return FontFromAsset("github.com/robbydyer/sports:/assets/fonts/04b24.ttf")
}

// FontFromAsset ...
func FontFromAsset(asset string) (*truetype.Font, error) {
	f, err := pkger.Open(asset)
	if err != nil {
		return nil, fmt.Errorf("failed to open asset %s with pkger: %w", asset, err)
	}
	defer f.Close()

	fBytes, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, fmt.Errorf("failed to read font bytes: %w", err)
	}

	fnt, err := freetype.ParseFont(fBytes)
	if err != nil {
		return nil, fmt.Errorf("failed to parse font: %w", err)
	}

	return fnt, nil
}

// SetFont ...
func (t *TextWriter) SetFont(fnt *truetype.Font) {
	t.font = fnt
	if t.context == nil {
		t.context = freetype.NewContext()
	}
	t.context.SetFont(fnt)
}

// Write ...
func (t *TextWriter) Write(canvas *rgb.Canvas, bounds image.Rectangle, str []string, clr color.Color) error {
	startX := bounds.Min.X + t.XStartCorrection
	drawer := &font.Drawer{
		Dst: canvas,
		Src: image.NewUniform(clr),
		Face: truetype.NewFace(t.font,
			&truetype.Options{
				Size:    t.FontSize,
				Hinting: font.HintingFull,
			},
		),
	}

	// lineY represents how much space to add for the newline
	lineY := int(math.Floor(t.FontSize + t.LineSpace))

	y := int(math.Floor(t.FontSize)) + bounds.Min.Y + t.YStartCorrection
	drawer.Dot = fixed.P(startX, y)

	for _, s := range str {
		drawer.DrawString(s)
		y += lineY + t.YStartCorrection
		drawer.Dot = fixed.P(startX, y)
	}

	return nil
}
