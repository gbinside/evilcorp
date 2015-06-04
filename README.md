# Evilcorp

An incremental kata solution to [this](https://github.com/Gianfrancoalongi/incremental_katas/tree/master/Evil_Corp), written in [GO](http://golang.org)

## Running

Type something like this.

	run.bat "what a nice friend"

to get 

    what a XXXX XXXXXX

## Build

Run 
    
    build.bat

## Pofile

First run the compiled program

    main.exe "what a nice friend"

After thar run

    go tool pprof main.exe main.prof

At the `(pprof)` prompt type 

    top 10 