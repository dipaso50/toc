# TOC 
Organiza carpetas como ~/Download en base a la extensión de ficheros.

Las reglas de organización se definen el un fichero json que debe ser pasado como argumento al ejecutable.

Las reglas se definen con expresiones regulares para abarcar el mayor número de casos y combinaciones de nombre de fichero posibles. 



## Construcción
***
En entornos linux con Go instalado:
```sh
$ git clone  https://github.com/dipaso50/toc.git
$ cd toc
$ make
```
En cualquier sistema operativo con Docker instalado 
```sh
$ git clone  https://github.com/dipaso50/toc.git
$ cd toc
$ docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp golang:1.15.6 make
```
Esto generará ejecutables para linux y windows en ./bin/linux y ./bin/windows respectivamente.

## Sintaxis
***

***Ayuda***
```sh
$ toc help
toc es una herramienta de organización de ficheros.
	
$> toc help para ayuda 

La variable de entorno $TOC_GOROUTINES controla el número de goroutinas que realizan el trabajo (por defecto 20)

Usage:
  toc [command]

Available Commands:
  byMod       Organiza carpetas en base a la fecha de modificación de ficheros
  byRule      Organiza carpetas en base a reglas
  help        Help about any command

Flags:
  -h, --help   help for toc

Use "toc [command] --help" for more information about a command.

```
## Ejemplos
***
Ir **[aquí](example/README.MD)** para ver ejemplos