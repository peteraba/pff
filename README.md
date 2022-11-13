Peter's File Finder
===================

I tend to name my files so that they basically include tags, separate by dashes. Examples
1. The Fight Ali vs Frazier-1pro-muhammadali-joefrazier-1971-round15.mpg
2. daughters-birthday-2022-grandma-athome-cut1.mpg
3. daves-birth-inberlin-2017.wav
4. daves-birthday-2018-porto-portugal.jpg

I created this tool to make it easier for myself to find files named in such a manner.

Usage
-----

Find all birthday files:
```bash
pff birthday                                                                                                                                                                             main|✚1…4
[...]/examples/daughters-birthday-2022-grandma-athome-cut1.mpg.txt
[...]/examples/daves-birthday-2018-porto-portugal.jpg.txt
```

Find all birth files. Note that birthday files will not be found.
```bash
pff birth                                                                                                                                                                             main|✚1…4
[...]/examples/daves-birth-inberlin-2017.wav.txt
```

Find all professional boxing related files from 1971. The reason this works is that numeric prefixes and postfixes are ignored.
```bash
pff pro 1071                                                                                                                                                                              main|✚1…4
[...]/examples/The Fight Ali vs Frazier-1pro-muhammadali-joefrazier-1971-round15.mpg.txt
```

Installation
------------

Install one of the binaries 

Options
-------
To get the options available, simply use `pff --help`
```
Usage of pff:
  -delimiters string
        characters used to find word boundaries (default "|-. ")
  -numsOkay
        if true, words prefixed or postfixed with numbers will be found (default true)
  -root string
        directory used to find files (current working directory by default) (default "/home/pabazsolt/Projects/pff")
```