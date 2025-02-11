package ttf

import (
	"image/color"
	"unsafe"

	"github.com/jupiterrider/purego-sdl3/internal/convert"
	"github.com/jupiterrider/purego-sdl3/sdl"
)

const (
	PropFontCreateFilenameString           = "SDL_ttf.font.create.filename"
	PropFontCreateIostreamPointer          = "SDL_ttf.font.create.iostream"
	PropFontCreateIostreamOffsetNumber     = "SDL_ttf.font.create.iostream.offset"
	PropFontCreateIostreamAutocloseBoolean = "SDL_ttf.font.create.iostream.autoclose"
	PropFontCreateSizeFloat                = "SDL_ttf.font.create.size"
	PropFontCreateFaceNumber               = "SDL_ttf.font.create.face"
	PropFontCreateHorizontalDpiNumber      = "SDL_ttf.font.create.hdpi"
	PropFontCreateVerticalDpiNumber        = "SDL_ttf.font.create.vdpi"
	PropFontCreateExistingFont             = "SDL_ttf.font.create.existing_font"

	PropFontOutlineLineCapNumber    = "SDL_ttf.font.outline.line_cap"
	PropFontOutlineLineJoinNumber   = "SDL_ttf.font.outline.line_join"
	PropFontOutlineMiterLimitNumber = "SDL_ttf.font.outline.miter_limit"

	PropRendererTextEngineRenderer         = "SDL_ttf.renderer_text_engine.create.renderer"
	PropRendererTextEngineAtlasTextureSize = "SDL_ttf.renderer_text_engine.create.atlas_texture_size"

	PropGPUTextEngineDevice           = "SDL_ttf.gpu_text_engine.create.device"
	PropGPUTextEngineAtlasTextureSize = "SDL_ttf.gpu_text_engine.create.atlas_texture_size"
)

type FontStyleFlags uint32

const (
	StyleNormal        = 0x00
	StyleBold          = 0x01
	StyleItalic        = 0x02
	StyleUnderline     = 0x04
	StyleStrikethrough = 0x08
)

type SubStringFlags uint32

const (
	SubStringDirectionMask = 0x000000FF
	SubStringTextStart     = 0x00000100
	SubStringLineStart     = 0x00000200
	SubStringLineEnd       = 0x00000400
	SubStringTextEnd       = 0x00000800
)

type Direction uint32

const (
	DirectionInvalid Direction = 0
	DirectionLtr     Direction = iota + 3 // Left to Right
	DirectionRtl                          // Right to Left
	DirectionTtb                          // Top to Bottom
	DirectionBtt                          // Bottom to Top
)

type HintingFlags uint32

const (
	HintingNormal HintingFlags = iota
	HintingLight
	HintingMono
	HintingNone
	HintingLightSubpixel
)

type HorizontalAlignment int32

const (
	HorizontalAlignInvalid HorizontalAlignment = iota - 1
	HorizontalAlignLeft
	HorizontalAlignCenter
	HorizontalAlignRight
)

type ImageType uint32

const (
	ImageInvalid ImageType = iota
	ImageAlpha
	ImageColor
	ImageSDF
)

type GPUTextEngineWinding int32

const (
	GPUTextEngineWindingInvalid GPUTextEngineWinding = iota - 1
	GPUTextEngineWindingClockwise
	GPUTextEngineWindingCounterClockwise
)

type Font struct{}

type TextData struct{}

type TextEngine struct{}

type Text struct {
	text     *byte
	NumLines int32
	Refcount int32
	Internal *TextData
}

func (t *Text) Text() string {
	return convert.ToString(t.text)
}

type SubString struct {
	Flags        SubStringFlags
	Offset       int32
	Length       int32
	LineIndex    int32
	ClusterIndex int32
	Rect         sdl.Rect
}

// Version gets the version of the dynamically linked SDL_ttf library.
func Version() (major, minor, patch int32) {
	version := ttfVersion()
	major = version / 1000000
	minor = version / 1000 % 1000
	patch = version % 1000
	return
}

// GetFreeTypeVersion queries the version of the FreeType library in use.
func GetFreeTypeVersion() (major, minor, patch int32) {
	ttfGetFreeTypeVersion(&major, &minor, &patch)
	return
}

// Init initializes SDL_ttf.
func Init() bool {
	return ttfInit()
}

// GetHarfBuzzVersion queries the version of the HarfBuzz library in use.
func GetHarfBuzzVersion() (major, minor, patch int32) {
	ttfGetHarfBuzzVersion(&major, &minor, &patch)
	return
}

func AddFallbackFont(font *Font, fallback *Font) bool {
	return ttfAddFallbackFont(font, fallback)
}

func AppendTextString(text *Text, string string, length uint64) bool {
	return ttfAppendTextString(text, string, length)
}

func ClearFallbackFonts(font *Font) {
	ttfClearFallbackFonts(font)
}

func CloseFont(font *Font) {
	ttfCloseFont(font)
}

func CopyFont(existingFont *Font) *Font {
	return ttfCopyFont(existingFont)
}

func CreateGPUTextEngine(device *sdl.GPUDevice) *TextEngine {
	return ttfCreateGPUTextEngine(device)
}

func CreateGPUTextEngineWithProperties(props sdl.PropertiesID) *TextEngine {
	return ttfCreateGPUTextEngineWithProperties(props)
}

func CreateRendererTextEngine(renderer *sdl.Renderer) *TextEngine {
	return ttfCreateRendererTextEngine(renderer)
}

func CreateRendererTextEngineWithProperties(props sdl.PropertiesID) *TextEngine {
	return ttfCreateRendererTextEngineWithProperties(props)
}

func CreateSurfaceTextEngine() *TextEngine {
	return ttfCreateSurfaceTextEngine()
}

func CreateText(engine *TextEngine, font *Font, text string, length uint64) *Text {
	return ttfCreateText(engine, font, text, length)
}

func DeleteTextString(text *Text, offset int32, length int32) bool {
	return ttfDeleteTextString(text, offset, length)
}

func DestroyGPUTextEngine(engine *TextEngine) {
	ttfDestroyGPUTextEngine(engine)
}

func DestroyRendererTextEngine(engine *TextEngine) {
	ttfDestroyRendererTextEngine(engine)
}

func DestroySurfaceTextEngine(engine *TextEngine) {
	ttfDestroySurfaceTextEngine(engine)
}

func DestroyText(text *Text) {
	ttfDestroyText(text)
}

func DrawRendererText(text *Text, x float32, y float32) bool {
	return ttfDrawRendererText(text, x, y)
}

func DrawSurfaceText(text *Text, x int32, y int32, surface *sdl.Surface) bool {
	return ttfDrawSurfaceText(text, x, y, surface)
}

func FontHasGlyph(font *Font, ch rune) bool {
	return ttfFontHasGlyph(font, ch)
}

func FontIsFixedWidth(font *Font) bool {
	return ttfFontIsFixedWidth(font)
}

func FontIsScalable(font *Font) bool {
	return ttfFontIsScalable(font)
}

func GetFontAscent(font *Font) int32 {
	return ttfGetFontAscent(font)
}

func GetFontDescent(font *Font) int32 {
	return ttfGetFontDescent(font)
}

func GetFontDirection(font *Font) Direction {
	return ttfGetFontDirection(font)
}

func GetFontDPI(font *Font, hdpi *int32, vdpi *int32) bool {
	return ttfGetFontDPI(font, hdpi, vdpi)
}

func GetFontFamilyName(font *Font) string {
	return ttfGetFontFamilyName(font)
}

func GetFontGeneration(font *Font) uint32 {
	return ttfGetFontGeneration(font)
}

func GetFontHeight(font *Font) int32 {
	return ttfGetFontHeight(font)
}

func GetFontHinting(font *Font) HintingFlags {
	return ttfGetFontHinting(font)
}

func GetFontKerning(font *Font) bool {
	return ttfGetFontKerning(font)
}

func GetFontLineSkip(font *Font) int32 {
	return ttfGetFontLineSkip(font)
}

func GetFontOutline(font *Font) int32 {
	return ttfGetFontOutline(font)
}

func GetFontProperties(font *Font) sdl.PropertiesID {
	return ttfGetFontProperties(font)
}

func GetFontScript(font *Font) uint32 {
	return ttfGetFontScript(font)
}

func GetFontSDF(font *Font) bool {
	return ttfGetFontSDF(font)
}

func GetFontSize(font *Font) float32 {
	return ttfGetFontSize(font)
}

func GetFontStyle(font *Font) FontStyleFlags {
	return ttfGetFontStyle(font)
}

func GetFontStyleName(font *Font) string {
	return ttfGetFontStyleName(font)
}

func GetFontWrapAlignment(font *Font) HorizontalAlignment {
	return ttfGetFontWrapAlignment(font)
}

func GetGlyphImage(font *Font, ch uint32, image_type *ImageType) *sdl.Surface {
	return ttfGetGlyphImage(font, ch, image_type)
}

func GetGlyphImageForIndex(font *Font, glyph_index uint32, image_type *ImageType) *sdl.Surface {
	return ttfGetGlyphImageForIndex(font, glyph_index, image_type)
}

func GetGlyphKerning(font *Font, previous_ch uint32, ch uint32, kerning *int32) bool {
	return ttfGetGlyphKerning(font, previous_ch, ch, kerning)
}

func GetGlyphMetrics(font *Font, ch uint32, minx *int32, maxx *int32, miny *int32, maxy *int32, advance *int32) bool {
	return ttfGetGlyphMetrics(font, ch, minx, maxx, miny, maxy, advance)
}

func GetGlyphScript(ch uint32) uint32 {
	return ttfGetGlyphScript(ch)
}

func GetGPUTextDrawData(text *Text) {
	ttfGetGPUTextDrawData(text)
}

func GetGPUTextEngineWinding(engine *TextEngine) GPUTextEngineWinding {
	return ttfGetGPUTextEngineWinding(engine)
}

func GetNextTextSubString(text *Text, substring *SubString, next *SubString) bool {
	return ttfGetNextTextSubString(text, substring, next)
}

func GetNumFontFaces(font *Font) int32 {
	return ttfGetNumFontFaces(font)
}

// func GetPreviousTextSubString(text *Text, substring *SubString, previous *SubString) bool {
//	return ttfGetPreviousTextSubString(text, substring, previous)
// }

// func GetStringSize(font *Font, text string, length uint64, w *int32, h *int32) bool {
//	return ttfGetStringSize(font, text, length, w, h)
// }

// func GetStringSizeWrapped(font *Font, text string, length uint64, wrap_width int32, w *int32, h *int32) bool {
//	return ttfGetStringSizeWrapped(font, text, length, wrap_width, w, h)
// }

// func GetTextColor(text *Text, r *uint8, g *uint8, b *uint8, a *uint8) bool {
//	return ttfGetTextColor(text, r, g, b, a)
// }

// func GetTextColorFloat(text *Text, r *float32, g *float32, b *float32, a *float32) bool {
//	return ttfGetTextColorFloat(text, r, g, b, a)
// }

// func GetTextDirection(text *Text) Direction {
//	return ttfGetTextDirection(text)
// }

// func GetTextEngine(text *Text) *TextEngine {
//	return ttfGetTextEngine(text)
// }

// func GetTextFont(text *Text) *Font {
//	return ttfGetTextFont(text)
// }

// func GetTextPosition(text *Text, x *int32, y *int32) bool {
//	return ttfGetTextPosition(text, x, y)
// }

// func GetTextProperties(text *Text) PropertiesID {
//	return ttfGetTextProperties(text)
// }

// func GetTextScript(text *Text) uint32 {
//	return ttfGetTextScript(text)
// }

// func GetTextSize(text *Text, w *int32, h *int32) bool {
//	return ttfGetTextSize(text, w, h)
// }

// func GetTextSubString(text *Text, offset int32, substring *SubString) bool {
//	return ttfGetTextSubString(text, offset, substring)
// }

// func GetTextSubStringForLine(text *Text, line int32, substring *SubString) bool {
//	return ttfGetTextSubStringForLine(text, line, substring)
// }

// func GetTextSubStringForPoint(text *Text, x int32, y int32, substring *SubString) bool {
//	return ttfGetTextSubStringForPoint(text, x, y, substring)
// }

// func GetTextSubStringsForRange(text *Text, offset int32, length int32, count *int32)  {
//	ttfGetTextSubStringsForRange(text, offset, length, count)
// }

// func GetTextWrapWidth(text *Text, wrap_width *int32) bool {
//	return ttfGetTextWrapWidth(text, wrap_width)
// }

// func InsertTextString(text *Text, offset int32, string string, length uint64) bool {
//	return ttfInsertTextString(text, offset, string, length)
// }

// func MeasureString(font *Font, text string, length uint64, max_width int32, measured_width *int32, measured_length *uint64) bool {
//	return ttfMeasureString(font, text, length, max_width, measured_width, measured_length)
// }

func OpenFont(file string, ptsize float32) *Font {
	return ttfOpenFont(file, ptsize)
}

func OpenFontIO(src *sdl.IOStream, closeio bool, ptsize float32) *Font {
	return ttfOpenFontIO(src, closeio, ptsize)
}

// func OpenFontWithProperties(props PropertiesID) *Font {
//	return ttfOpenFontWithProperties(props)
// }

func Quit() {
	ttfQuit()
}

// func RemoveFallbackFont(font *Font, fallback *Font)  {
//	ttfRemoveFallbackFont(font, fallback)
// }

// func RenderGlyph_Blended(font *Font, ch uint32, fg color.RGBA) *Surface {
//	return ttfRenderGlyph_Blended(font, ch, fg)
// }

// func RenderGlyph_LCD(font *Font, ch uint32, fg color.RGBA, bg color.RGBA) *Surface {
//	return ttfRenderGlyph_LCD(font, ch, fg, bg)
// }

// func RenderGlyph_Shaded(font *Font, ch uint32, fg color.RGBA, bg color.RGBA) *Surface {
//	return ttfRenderGlyph_Shaded(font, ch, fg, bg)
// }

// func RenderGlyph_Solid(font *Font, ch uint32, fg color.RGBA) *Surface {
//	return ttfRenderGlyph_Solid(font, ch, fg)
// }

// func RenderText_Blended(font *Font, text string, length uint64, fg color.RGBA) *Surface {
//	return ttfRenderText_Blended(font, text, length, fg)
// }

// func RenderText_Blended_Wrapped(font *Font, text string, length uint64, fg color.RGBA, wrap_width int32) *Surface {
//	return ttfRenderText_Blended_Wrapped(font, text, length, fg, wrap_width)
// }

// func RenderText_LCD(font *Font, text string, length uint64, fg color.RGBA, bg color.RGBA) *Surface {
//	return ttfRenderText_LCD(font, text, length, fg, bg)
// }

// func RenderText_LCD_Wrapped(font *Font, text string, length uint64, fg color.RGBA, bg color.RGBA, wrap_width int32) *Surface {
//	return ttfRenderText_LCD_Wrapped(font, text, length, fg, bg, wrap_width)
// }

// func RenderText_Shaded(font *Font, text string, length uint64, fg color.RGBA, bg color.RGBA) *Surface {
//	return ttfRenderText_Shaded(font, text, length, fg, bg)
// }

// func RenderText_Shaded_Wrapped(font *Font, text string, length uint64, fg color.RGBA, bg color.RGBA, wrap_width int32) *Surface {
//	return ttfRenderText_Shaded_Wrapped(font, text, length, fg, bg, wrap_width)
// }

func RenderTextSolid(font *Font, text string, length uint64, fg color.RGBA) *sdl.Surface {
	return ttfRenderTextSolid(font, text, length, *(*uintptr)(unsafe.Pointer(&fg)))
}

// func RenderText_Solid_Wrapped(font *Font, text string, length uint64, fg color.RGBA, wrapLength int32) *Surface {
//	return ttfRenderText_Solid_Wrapped(font, text, length, fg, wrapLength)
// }

// func SetFontDirection(font *Font, direction Direction) bool {
//	return ttfSetFontDirection(font, direction)
// }

// func SetFontHinting(font *Font, hinting HintingFlags)  {
//	ttfSetFontHinting(font, hinting)
// }

// func SetFontKerning(font *Font, enabled bool)  {
//	ttfSetFontKerning(font, enabled)
// }

// func SetFontLanguage(font *Font, language_bcp47 string) bool {
//	return ttfSetFontLanguage(font, language_bcp47)
// }

// func SetFontLineSkip(font *Font, lineskip int32)  {
//	ttfSetFontLineSkip(font, lineskip)
// }

// func SetFontOutline(font *Font, outline int32) bool {
//	return ttfSetFontOutline(font, outline)
// }

// func SetFontScript(font *Font, script uint32) bool {
//	return ttfSetFontScript(font, script)
// }

// func SetFontSDF(font *Font, enabled bool) bool {
//	return ttfSetFontSDF(font, enabled)
// }

// func SetFontSize(font *Font, ptsize float32) bool {
//	return ttfSetFontSize(font, ptsize)
// }

// func SetFontSizeDPI(font *Font, ptsize float32, hdpi int32, vdpi int32) bool {
//	return ttfSetFontSizeDPI(font, ptsize, hdpi, vdpi)
// }

// func SetFontStyle(font *Font, style FontStyleFlags)  {
//	ttfSetFontStyle(font, style)
// }

// func SetFontWrapAlignment(font *Font, align HorizontalAlignment)  {
//	ttfSetFontWrapAlignment(font, align)
// }

// func SetGPUTextEngineWinding(engine *TextEngine, winding GPUTextEngineWinding)  {
//	ttfSetGPUTextEngineWinding(engine, winding)
// }

// func SetTextColor(text *Text, r uint8, g uint8, b uint8, a uint8) bool {
//	return ttfSetTextColor(text, r, g, b, a)
// }

// func SetTextColorFloat(text *Text, r float32, g float32, b float32, a float32) bool {
//	return ttfSetTextColorFloat(text, r, g, b, a)
// }

// func SetTextDirection(text *Text, direction Direction) bool {
//	return ttfSetTextDirection(text, direction)
// }

// func SetTextEngine(text *Text, engine *TextEngine) bool {
//	return ttfSetTextEngine(text, engine)
// }

// func SetTextFont(text *Text, font *Font) bool {
//	return ttfSetTextFont(text, font)
// }

// func SetTextPosition(text *Text, x int32, y int32) bool {
//	return ttfSetTextPosition(text, x, y)
// }

// func SetTextScript(text *Text, script uint32) bool {
//	return ttfSetTextScript(text, script)
// }

// func SetTextString(text *Text, string string, length uint64) bool {
//	return ttfSetTextString(text, string, length)
// }

// func SetTextWrapWhitespaceVisible(text *Text, visible bool) bool {
//	return ttfSetTextWrapWhitespaceVisible(text, visible)
// }

// func SetTextWrapWidth(text *Text, wrap_width int32) bool {
//	return ttfSetTextWrapWidth(text, wrap_width)
// }

// func TextWrapWhitespaceVisible(text *Text) bool {
//	return ttfTextWrapWhitespaceVisible(text)
// }

// func UpdateText(text *Text) bool {
//	return ttfUpdateText(text)
// }

// func WasInit() int32 {
//	return ttfWasInit()
// }
