## Organizar por reglas
***

Ir **[aquí](byRule/configExample.json)** para ver un ejemplo de json de configuración.

```sh
$ toc byRule -f ~/Descargas -c ./configExample.json
```

**Antes**
```sh
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
```sh
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


## Organizar por fecha de modificación
***

**Antes**
```sh
$ ls -l
total 0
-rw-rw-r-- 1 diego diego 0 ene 18  2028 file1.txt
-rw-rw-r-- 1 diego diego 0 ene 18  2038 file2.txt
```

**Después**
```sh
$ toc byMod -f .

Creando directorio 2038/01/18
Creando directorio 2028/01/18
Moviendo file2.txt -> 2038/01/18/file2.txt 
Moviendo file1.txt -> 2028/01/18/file1.txt 

$ tree
.
├── 2028
│   └── 01
│       └── 18
│           └── file1.txt
└── 2038
    └── 01
        └── 18
            └── file2.txt

6 directories, 2 files
diego@vant-N14xWU:/tmp/testtoc$
```