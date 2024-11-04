```bash
# cue version
cue version v0.10.1

go version go1.23.2
      -buildmode exe
       -compiler gc
       -trimpath true
     CGO_ENABLED 0
          GOARCH amd64
            GOOS linux
         GOAMD64 v1
cue.lang.version v0.10.1

# cue eval root.cue
root: "root"
vals: {
    ex1: "BAR"
    ex2: "b222"
    ex3: "hello"
    ex4: "world"
    ex5: "a111"
    ex6: "abc123"
}
```