2 points for templates

- parse
- execute

`parseGlob` => to specify a pattern of files
`parseFiles` => to parse a bunch of files

Important points
- `template.Must` will take care of error handling while parsing the glob
- `parseGlob` & `parseFiles` returns pointer to template & error
- `os.Stdout` would write the output onto console
- 2nd argument of `template.Execute` is data to be passed to the template
