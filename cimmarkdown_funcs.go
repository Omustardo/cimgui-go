// Code generated by cmd/codegen from https://github.com/AllenDang/cimgui-go.
// DO NOT EDIT.

package imgui

// #include "extra_types.h"
// #include "cimmarkdown_structs_accessor.h"
// #include "cimmarkdown_wrapper.h"
import "C"

func IsCharInsideWord(c_ rune) bool {
	return C.IsCharInsideWord(C.char(c_)) == C.bool(true)
}

func Markdown(markdown_ string, markdownLength_ uint64, mdConfig_ MarkdownConfig) {
	markdown_Arg, markdown_Fin := WrapString(markdown_)
	mdConfig_Arg, mdConfig_Fin := mdConfig_.c()
	C.Markdown(markdown_Arg, C.xulong(markdownLength_), mdConfig_Arg)

	markdown_Fin()
	mdConfig_Fin()
}

func RenderLine(markdown_ string, line_ *Line, textRegion_ *TextRegion, mdConfig_ MarkdownConfig) {
	markdown_Arg, markdown_Fin := WrapString(markdown_)
	line_Arg, line_Fin := line_.handle()
	textRegion_Arg, textRegion_Fin := textRegion_.handle()
	mdConfig_Arg, mdConfig_Fin := mdConfig_.c()
	C.RenderLine(markdown_Arg, line_Arg, textRegion_Arg, mdConfig_Arg)

	markdown_Fin()
	line_Fin()
	textRegion_Fin()
	mdConfig_Fin()
}

func RenderLinkText(self *TextRegion, text_ string, link_ Link, markdown_ string, mdConfig_ MarkdownConfig, linkHoverStart_ []string) bool {
	selfArg, selfFin := self.handle()
	text_Arg, text_Fin := WrapString(text_)
	link_Arg, link_Fin := link_.c()
	markdown_Arg, markdown_Fin := WrapString(markdown_)
	mdConfig_Arg, mdConfig_Fin := mdConfig_.c()
	linkHoverStart_Arg, linkHoverStart_Fin := WrapStringList(linkHoverStart_)

	defer func() {
		selfFin()
		text_Fin()
		link_Fin()
		markdown_Fin()
		mdConfig_Fin()
		linkHoverStart_Fin()
	}()
	return C.wrap_RenderLinkText(selfArg, text_Arg, link_Arg, markdown_Arg, mdConfig_Arg, linkHoverStart_Arg) == C.bool(true)
}

// RenderLinkTextWrappedV parameter default value hint:
// bIndentToHere_: false
func RenderLinkTextWrappedV(self *TextRegion, text_ string, link_ Link, markdown_ string, mdConfig_ MarkdownConfig, linkHoverStart_ []string, bIndentToHere_ bool) {
	selfArg, selfFin := self.handle()
	text_Arg, text_Fin := WrapString(text_)
	link_Arg, link_Fin := link_.c()
	markdown_Arg, markdown_Fin := WrapString(markdown_)
	mdConfig_Arg, mdConfig_Fin := mdConfig_.c()
	linkHoverStart_Arg, linkHoverStart_Fin := WrapStringList(linkHoverStart_)
	C.wrap_RenderLinkTextWrappedV(selfArg, text_Arg, link_Arg, markdown_Arg, mdConfig_Arg, linkHoverStart_Arg, C.bool(bIndentToHere_))

	selfFin()
	text_Fin()
	link_Fin()
	markdown_Fin()
	mdConfig_Fin()
	linkHoverStart_Fin()
}

func RenderListTextWrapped(self *TextRegion, text_ string) {
	selfArg, selfFin := self.handle()
	text_Arg, text_Fin := WrapString(text_)
	C.wrap_RenderListTextWrapped(selfArg, text_Arg)

	selfFin()
	text_Fin()
}

// RenderTextWrappedV parameter default value hint:
// bIndentToHere_: false
func RenderTextWrappedV(self *TextRegion, text_ string, bIndentToHere_ bool) {
	selfArg, selfFin := self.handle()
	text_Arg, text_Fin := WrapString(text_)
	C.wrap_RenderTextWrappedV(selfArg, text_Arg, C.bool(bIndentToHere_))

	selfFin()
	text_Fin()
}

func ResetIndent(self *TextRegion) {
	selfArg, selfFin := self.handle()
	C.ResetIndent(selfArg)

	selfFin()
}

func NewTextRegion() *TextRegion {
	return newTextRegionFromC(C.TextRegion_TextRegion())
}

func (self *TextRegion) Destroy() {
	selfArg, selfFin := self.handle()
	C.TextRegion_destroy(selfArg)

	selfFin()
}

func UnderLine(col_ Color) {
	C.UnderLine(col_.toC())
}

func RenderLinkTextWrapped(self *TextRegion, text_ string, link_ Link, markdown_ string, mdConfig_ MarkdownConfig, linkHoverStart_ []string) {
	selfArg, selfFin := self.handle()
	text_Arg, text_Fin := WrapString(text_)
	link_Arg, link_Fin := link_.c()
	markdown_Arg, markdown_Fin := WrapString(markdown_)
	mdConfig_Arg, mdConfig_Fin := mdConfig_.c()
	linkHoverStart_Arg, linkHoverStart_Fin := WrapStringList(linkHoverStart_)
	C.wrap_RenderLinkTextWrapped(selfArg, text_Arg, link_Arg, markdown_Arg, mdConfig_Arg, linkHoverStart_Arg)

	selfFin()
	text_Fin()
	link_Fin()
	markdown_Fin()
	mdConfig_Fin()
	linkHoverStart_Fin()
}

func RenderTextWrapped(self *TextRegion, text_ string) {
	selfArg, selfFin := self.handle()
	text_Arg, text_Fin := WrapString(text_)
	C.wrap_RenderTextWrapped(selfArg, text_Arg)

	selfFin()
	text_Fin()
}

func (self Emphasis) SetState(v EmphasisState) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Emphasis_SetState(selfArg, C.EmphasisState(v))
}

func (self *Emphasis) State() EmphasisState {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return EmphasisState(C.wrap_Emphasis_GetState(selfArg))
}

func (self Emphasis) SetText(v TextBlock) {
	vArg, _ := v.c()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Emphasis_SetText(selfArg, vArg)
}

func (self *Emphasis) Text() TextBlock {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_Emphasis_GetText(selfArg)
	return *newTextBlockFromC(func() *C.TextBlock { result := result; return &result }())
}

func (self Emphasis) SetSym(v rune) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Emphasis_SetSym(selfArg, C.char(v))
}

func (self *Emphasis) Sym() rune {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return rune(C.wrap_Emphasis_GetSym(selfArg))
}

func (self Line) SetIsHeading(v bool) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Line_SetIsHeading(selfArg, C.bool(v))
}

func (self *Line) IsHeading() bool {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Line_GetIsHeading(selfArg) == C.bool(true)
}

func (self Line) SetIsEmphasis(v bool) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Line_SetIsEmphasis(selfArg, C.bool(v))
}

func (self *Line) IsEmphasis() bool {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Line_GetIsEmphasis(selfArg) == C.bool(true)
}

func (self Line) SetIsUnorderedListStart(v bool) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Line_SetIsUnorderedListStart(selfArg, C.bool(v))
}

func (self *Line) IsUnorderedListStart() bool {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Line_GetIsUnorderedListStart(selfArg) == C.bool(true)
}

func (self Line) SetIsLeadingSpace(v bool) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Line_SetIsLeadingSpace(selfArg, C.bool(v))
}

func (self *Line) IsLeadingSpace() bool {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Line_GetIsLeadingSpace(selfArg) == C.bool(true)
}

func (self Line) SetLeadSpaceCount(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Line_SetLeadSpaceCount(selfArg, C.int(v))
}

func (self *Line) LeadSpaceCount() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetLeadSpaceCount(selfArg))
}

func (self Line) SetHeadingCount(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Line_SetHeadingCount(selfArg, C.int(v))
}

func (self *Line) HeadingCount() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetHeadingCount(selfArg))
}

func (self Line) SetEmphasisCount(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Line_SetEmphasisCount(selfArg, C.int(v))
}

func (self *Line) EmphasisCount() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetEmphasisCount(selfArg))
}

func (self Line) SetLineStart(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Line_SetLineStart(selfArg, C.int(v))
}

func (self *Line) LineStart() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetLineStart(selfArg))
}

func (self Line) SetLineEnd(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Line_SetLineEnd(selfArg, C.int(v))
}

func (self *Line) LineEnd() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetLineEnd(selfArg))
}

func (self Line) SetLastRenderPosition(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Line_SetLastRenderPosition(selfArg, C.int(v))
}

func (self *Line) LastRenderPosition() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Line_GetLastRenderPosition(selfArg))
}

func (self Link) SetState(v LinkState) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Link_SetState(selfArg, C.LinkState(v))
}

func (self *Link) State() LinkState {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return LinkState(C.wrap_Link_GetState(selfArg))
}

func (self Link) SetText(v TextBlock) {
	vArg, _ := v.c()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Link_SetText(selfArg, vArg)
}

func (self *Link) Text() TextBlock {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_Link_GetText(selfArg)
	return *newTextBlockFromC(func() *C.TextBlock { result := result; return &result }())
}

func (self Link) SetUrl(v TextBlock) {
	vArg, _ := v.c()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Link_SetUrl(selfArg, vArg)
}

func (self *Link) Url() TextBlock {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_Link_GetUrl(selfArg)
	return *newTextBlockFromC(func() *C.TextBlock { result := result; return &result }())
}

func (self Link) SetIsImage(v bool) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Link_SetIsImage(selfArg, C.bool(v))
}

func (self *Link) IsImage() bool {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_Link_GetIsImage(selfArg) == C.bool(true)
}

func (self Link) SetNumbracketsopen(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_Link_SetNum_brackets_open(selfArg, C.int(v))
}

func (self *Link) Numbracketsopen() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_Link_GetNum_brackets_open(selfArg))
}

func (self MarkdownConfig) SetLinkIcon(v string) {
	vArg, _ := WrapString(v)

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetLinkIcon(selfArg, vArg)
}

func (self *MarkdownConfig) LinkIcon() string {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.GoString(C.wrap_MarkdownConfig_GetLinkIcon(selfArg))
}

func (self MarkdownConfig) SetHeadingFormats(v *[3]MarkdownHeadingFormat) {
	vArg := make([]MarkdownHeadingFormat, len(v))
	for i, vV := range v {
		vVArg, vVFin := vV.c()
		vArg[i] = vVArg
	}

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetHeadingFormats(selfArg, (**C.MarkdownHeadingFormat)(&vArg[0]))

	for i, vV := range vArg {
		(*v)[i] = *newMarkdownHeadingFormatFromC(func() *C.MarkdownHeadingFormat { result := vV; return &result }())
	}
}

func (self MarkdownConfig) SetUserData(v uintptr) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownConfig_SetUserData(selfArg, C.uintptr_t(v))
}

func (self *MarkdownConfig) UserData() uintptr {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return uintptr(C.wrap_MarkdownConfig_GetUserData(selfArg))
}

func (self MarkdownFormatInfo) SetType(v MarkdownFormatType) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownFormatInfo_SetType(selfArg, C.MarkdownFormatType(v))
}

func (self *MarkdownFormatInfo) Type() MarkdownFormatType {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return MarkdownFormatType(C.wrap_MarkdownFormatInfo_GetType(selfArg))
}

func (self MarkdownFormatInfo) SetItemHovered(v bool) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownFormatInfo_SetItemHovered(selfArg, C.bool(v))
}

func (self *MarkdownFormatInfo) ItemHovered() bool {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownFormatInfo_GetItemHovered(selfArg) == C.bool(true)
}

func (self MarkdownFormatInfo) SetConfig(v *MarkdownConfig) {
	vArg, _ := v.handle()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownFormatInfo_SetConfig(selfArg, vArg)
}

func (self *MarkdownFormatInfo) Config() *MarkdownConfig {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return newMarkdownConfigFromC(C.wrap_MarkdownFormatInfo_GetConfig(selfArg))
}

func (self MarkdownHeadingFormat) SetFont(v *Font) {
	vArg, _ := v.handle()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownHeadingFormat_SetFont(selfArg, vArg)
}

func (self *MarkdownHeadingFormat) Font() *Font {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return newFontFromC(C.wrap_MarkdownHeadingFormat_GetFont(selfArg))
}

func (self MarkdownHeadingFormat) SetSeparator(v bool) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownHeadingFormat_SetSeparator(selfArg, C.bool(v))
}

func (self *MarkdownHeadingFormat) Separator() bool {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownHeadingFormat_GetSeparator(selfArg) == C.bool(true)
}

func (self MarkdownImageData) SetIsValid(v bool) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetIsValid(selfArg, C.bool(v))
}

func (self *MarkdownImageData) IsValid() bool {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownImageData_GetIsValid(selfArg) == C.bool(true)
}

func (self MarkdownImageData) SetUseLinkCallback(v bool) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetUseLinkCallback(selfArg, C.bool(v))
}

func (self *MarkdownImageData) UseLinkCallback() bool {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownImageData_GetUseLinkCallback(selfArg) == C.bool(true)
}

func (self MarkdownImageData) SetUsertextureid(v TextureID) {
	vArg, _ := v.c()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetUser_texture_id(selfArg, vArg)
}

func (self MarkdownImageData) SetSize(v Vec2) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetSize(selfArg, v.toC())
}

func (self *MarkdownImageData) Size() Vec2 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *(&Vec2{}).fromC(C.wrap_MarkdownImageData_GetSize(selfArg))
}

func (self MarkdownImageData) SetUv0(v Vec2) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetUv0(selfArg, v.toC())
}

func (self *MarkdownImageData) Uv0() Vec2 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *(&Vec2{}).fromC(C.wrap_MarkdownImageData_GetUv0(selfArg))
}

func (self MarkdownImageData) SetUv1(v Vec2) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetUv1(selfArg, v.toC())
}

func (self *MarkdownImageData) Uv1() Vec2 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *(&Vec2{}).fromC(C.wrap_MarkdownImageData_GetUv1(selfArg))
}

func (self MarkdownImageData) SetTintcol(v Vec4) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetTint_col(selfArg, v.toC())
}

func (self *MarkdownImageData) Tintcol() Vec4 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *(&Vec4{}).fromC(C.wrap_MarkdownImageData_GetTint_col(selfArg))
}

func (self MarkdownImageData) SetBordercol(v Vec4) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownImageData_SetBorder_col(selfArg, v.toC())
}

func (self *MarkdownImageData) Bordercol() Vec4 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return *(&Vec4{}).fromC(C.wrap_MarkdownImageData_GetBorder_col(selfArg))
}

func (self MarkdownLinkCallbackData) SetText(v string) {
	vArg, _ := WrapString(v)

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetText(selfArg, vArg)
}

func (self *MarkdownLinkCallbackData) Text() string {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.GoString(C.wrap_MarkdownLinkCallbackData_GetText(selfArg))
}

func (self MarkdownLinkCallbackData) SetTextLength(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetTextLength(selfArg, C.int(v))
}

func (self *MarkdownLinkCallbackData) TextLength() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_MarkdownLinkCallbackData_GetTextLength(selfArg))
}

func (self MarkdownLinkCallbackData) SetLink(v string) {
	vArg, _ := WrapString(v)

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetLink(selfArg, vArg)
}

func (self *MarkdownLinkCallbackData) Link() string {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.GoString(C.wrap_MarkdownLinkCallbackData_GetLink(selfArg))
}

func (self MarkdownLinkCallbackData) SetLinkLength(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetLinkLength(selfArg, C.int(v))
}

func (self *MarkdownLinkCallbackData) LinkLength() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_MarkdownLinkCallbackData_GetLinkLength(selfArg))
}

func (self MarkdownLinkCallbackData) SetUserData(v uintptr) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetUserData(selfArg, C.uintptr_t(v))
}

func (self *MarkdownLinkCallbackData) UserData() uintptr {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return uintptr(C.wrap_MarkdownLinkCallbackData_GetUserData(selfArg))
}

func (self MarkdownLinkCallbackData) SetIsImage(v bool) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownLinkCallbackData_SetIsImage(selfArg, C.bool(v))
}

func (self *MarkdownLinkCallbackData) IsImage() bool {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.wrap_MarkdownLinkCallbackData_GetIsImage(selfArg) == C.bool(true)
}

func (self MarkdownTooltipCallbackData) SetLinkData(v MarkdownLinkCallbackData) {
	vArg, _ := v.c()

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownTooltipCallbackData_SetLinkData(selfArg, vArg)
}

func (self *MarkdownTooltipCallbackData) LinkData() MarkdownLinkCallbackData {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()

	result := C.wrap_MarkdownTooltipCallbackData_GetLinkData(selfArg)
	return *newMarkdownLinkCallbackDataFromC(func() *C.MarkdownLinkCallbackData { result := result; return &result }())
}

func (self MarkdownTooltipCallbackData) SetLinkIcon(v string) {
	vArg, _ := WrapString(v)

	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_MarkdownTooltipCallbackData_SetLinkIcon(selfArg, vArg)
}

func (self *MarkdownTooltipCallbackData) LinkIcon() string {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return C.GoString(C.wrap_MarkdownTooltipCallbackData_GetLinkIcon(selfArg))
}

func (self TextBlock) SetStart(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_TextBlock_SetStart(selfArg, C.int(v))
}

func (self *TextBlock) Start() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_TextBlock_GetStart(selfArg))
}

func (self TextBlock) SetStop(v int32) {
	selfArg, selfFin := self.handle()
	defer selfFin()
	C.wrap_TextBlock_SetStop(selfArg, C.int(v))
}

func (self *TextBlock) Stop() int32 {
	selfArg, selfFin := self.handle()

	defer func() {
		selfFin()
	}()
	return int32(C.wrap_TextBlock_GetStop(selfArg))
}
