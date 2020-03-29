Go uses a map to map a function to string. That special map is called `funcMap`

```var fm = tempalte.funcMap {
  "uc": strings.toUpper,
  "ft": strings.toLower,
}```

Usage in html:

```{{uc .Name}}``` <br>

funcMap is of type `map[string]interface{}`

- funcMap should be assigned before the `parseGlob` or `parseFiles`. If we assign it after, templates would not have function definitions. It would just be some static constants without definitions
- Using `New()` to get pointer to a template. Assign funcMap to that followed the parsing.
