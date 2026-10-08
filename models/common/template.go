package common

// Template editors. A template's editor is set when it is created and never
// changes.
//
//   - TemplateEditorAdvanced: designed in MJML.
//   - TemplateEditorSimple: designed in the dashboard's simple editor. Its HTML
//     is made from the design and can be read but not changed through the API.
//   - TemplateEditorHTML: an HTML template.
const (
	TemplateEditorAdvanced = "advanced"
	TemplateEditorSimple   = "simple"
	TemplateEditorHTML     = "html"
)
