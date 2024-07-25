package static

import _ "embed"

//go:embed js.js
var Js []byte
//go:embed main.html
var MainHtml []byte