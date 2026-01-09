# gostdwebapp

## Template-less GoLang Demo App!

### - This concept of template-less rendering was thought of to avoid the hassles of configuring external templating engine

### - No templates. Pure performance. Direct HTML streaming skips parse/render overhead.

### - Created utility package to eliminate repetition; generate perfect tables from Go structs instantly.

## Compilation Script

```
go build -ldflags "-w -s" .
upx --best --lzma $1
ls -l
./$1

- $1 = the runtime binary
```


## Run as:

```
./run gostdwebapp
```

### Screenshot

[View Screenshot](template_less_go_app_home_page.png)