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

## Profile

First run the compiled program

    main.exe "what a nice friend"

After thar run

    go tool pprof main.exe main.prof

At the `(pprof)` prompt type

    top 10

to get

    2290ms of 3800ms total (60.26%)
    Dropped 36 nodes (cum <= 19ms)
    Showing top 10 nodes out of 79 (cum >= 470ms)
          flat  flat%   sum%        cum   cum%
        1150ms 30.26% 30.26%     1150ms 30.26%  scanblock
         290ms  7.63% 37.89%      370ms  9.74%  runtime.mallocgc
         210ms  5.53% 43.42%      210ms  5.53%  runtime.MSpan_Sweep
         150ms  3.95% 47.37%      450ms 11.84%  runtime.growslice
         130ms  3.42% 50.79%      170ms  4.47%  unicode.SimpleFold
          90ms  2.37% 53.16%       90ms  2.37%  runtime.findfunc
          80ms  2.11% 55.26%       80ms  2.11%  runtime.pcvalue
          70ms  1.84% 57.11%       70ms  1.84%  runtime.writebarrierptr
          60ms  1.58% 58.68%      100ms  2.63%  regexp.(*machine).add
          60ms  1.58% 60.26%      470ms 12.37%  regexp/syntax.(*compiler).compile
