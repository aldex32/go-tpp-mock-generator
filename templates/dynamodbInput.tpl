{
"psd2-consent-tpptest": [

  {{$addSeparator := addSeparator}}
  {{range .}}
  {{if call $addSeparator}}, {{end}}
  {
  "PutRequest": {{.}}
  }
  {{end}}
]
}
