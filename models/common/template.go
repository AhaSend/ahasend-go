package common

// Template editors. A template's editor is set when it is created and never
// changes.
//
//   - TemplateEditorAdvanced: designed in MJML.
//   - TemplateEditorSimple: designed in the dashboard's simple editor. The API
//     reads and writes its design as AhaSend Markup (AhaML), and makes its HTML
//     from the design.
//   - TemplateEditorHTML: an HTML template.
const (
	TemplateEditorAdvanced = "advanced"
	TemplateEditorSimple   = "simple"
	TemplateEditorHTML     = "html"
)

// MarkupWarningJinjaInText is the code of the only AhaML warning: text holds
// {{, which stays text. Use a <Var> for a variable.
const MarkupWarningJinjaInText = "jinja_in_text"

// MarkupProblem is one problem of AhaSend Markup (AhaML) that a write refused,
// or one warning about AhaML that a write stored.
type MarkupProblem struct {
	// Code says what is wrong, as a code that stays the same. The AhaML
	// reference lists every code. A warning's code is
	// MarkupWarningJinjaInText.
	Code string `json:"code"`
	// Message says what is wrong, in words for a person.
	Message string `json:"message"`
	// Line is the line where the problem starts, counted from 1. It is 0 when
	// the problem has no place in the AhaML, as for a size limit.
	Line int `json:"line,omitempty"`
	// Column is the column where the problem starts, counted from 1 in
	// Unicode characters. It is 0 when Line is.
	Column int `json:"column,omitempty"`
	// Path names the element of the problem, each tag with its position among
	// its siblings of the same tag, such as "Email > Section[2] > Paragraph[1]".
	// It is empty when Line is 0.
	Path string `json:"path,omitempty"`
	// Hint says how to fix the problem, when a fix is known. Warnings have
	// none.
	Hint string `json:"hint,omitempty"`
}
