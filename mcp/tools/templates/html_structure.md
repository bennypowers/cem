<{{.Element.TagName}}{{range .RequiredAttributes}} {{.Name}}="{{.Value}}"{{end}}{{range .SortedAttributes}} {{.Name}}="{{.Value}}"{{end}}>{{if gt (len .Slots) 0}}
{{range .Slots}}{{if .Name}}  <span slot="{{.Name}}">{{.ExampleContent}}</span>
{{else}}  {{.DefaultContent}}
{{end}}{{end}}{{else}}{{.Content}}{{end}}</{{.Element.TagName}}>