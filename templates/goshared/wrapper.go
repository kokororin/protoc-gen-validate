package goshared

const wrapperTpl = `
	{{ $f := .Field }}{{ $r := .Rules }}

	if wrapper := {{ accessor . }}; wrapper != nil {
		{{ render (unwrap . "wrapper") }}
	} {{ if .MessageRules.GetRequired }} else {
		{{ if ne $r.GetErrorMsg "" }}
		err := {{ err . $r.GetErrorMsg }}
		{{ else }}
		err := {{ err . "value is required and must not be nil." }}
		{{ end }}
		if !all { return err }
		errors = append(errors, err)
	} {{ end }}
`
