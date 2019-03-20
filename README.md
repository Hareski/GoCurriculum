# GoCurriculum
Simple curriculum vitae with Go Language

## Compilation
To compile please use something like :
```sh
$ gccgo ServerCV.go -o serverCV
```

## Usage
To use this program do :
```sh
$ ./serverC [Port] [Name] {[Link name 1] [link url 1]} {...} ...
```
with [Port] the port use for the server.

An example with my personal curriculum vitae on the port 8088 :
```sh
8088 "Areski Guilhem Himeur" "Curriculum vitae" https://drive.google.com/file/d/1bpdVJ-iiFQA3OMho-y6FHdRFy6U_O-XA Linkedin https://www.linkedin.com/in/AreskiHimeur/ Github https://github.com/Hareski
```
You can directly change the index.html with something better for you.

