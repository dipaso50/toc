# TOC 
Organiza carpetas como ~/Download en base a la extensión de ficheros.

Las reglas de organización se definen el un fichero json que debe ser pasado como argumento al ejecutable.

Las reglas se definen con expresiones regulares para abarcar el mayor número de casos y combinaciones de nombre de fichero posibles. 

Ir **[aquí](/example/organizeFolderConfigExample.json)** para ver un ejemplo de json de configuración.

## Construcción
***
En entornos linux con Go instalado:
```sh
$ git clone  https://github.com/dipaso50/toc.git
$ cd toc
$ make
```

Esto generará ejecutables para linux y windows en ./bin/linux y ./bin/windows respectivamente.

## Sintaxis
***

***Ayuda***
```sh
$ toc help
$ toc help organizeFolder
```

***Ejecutar el comando***
```sh
$ toc organizeFolder -d ~/Descargas -f ./config.json
```

## Ejemplo
***

**Antes**
```
Descargas/
├── sample.doc
├── sample.gif
├── sample.jpeg
├── sample.mp3
├── sample.mp4
├── sample.odt
├── sample.pdf
├── sample.rar
├── sample.txt
├── sample.xls
└── sample.zip
```
**Después**
```
Descargas/
├── AudioVideo
│   ├── sample.mp3
│   └── sample.mp4
├── Comprimidos
│   ├── sample.rar
│   └── sample.zip
├── Documentos
│   ├── sample.doc
│   ├── sample.odt
│   └── sample.pdf
├── Excels
│   └── sample.xls
├── Imagenes
│   ├── sample.gif
│   └── sample.jpeg
└── Txt
    └── sample.txt
```