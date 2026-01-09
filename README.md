# gostdwebapp

## Template-less GoLang Demo App!

## Using Postgres + std Go packages only.

- ### Invented template-less rendering—no external engine hassles
- ### Pure HTML streaming = zero parse overhead  
- ### HTable() utility = perfect tables from structs instantly

## Based on

- ### This is based on a sample database named 'northwind'.

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