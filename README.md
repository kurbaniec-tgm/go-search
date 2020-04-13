<h1 align="center">
  <br>
  <img src="images/Logo/Logo.png" alt="go search" width="250"></a>
  <br>
  Go Search!
  <br>
</h1>

<h4 align="center">A simple command line based utility to search files and directories on your system.</h4>

## 🛠️ Build

```bash
go build -o ./bin/search[.exe|.sh|...] search/src
```

## 🚴 Run

```
search <File|Directory|Pattern> [Options]
```

The  search requests can be a file or directory name. You can also use glob pattern like `*` or `[Ss]ong.*` to look for your files! 

For more information about the glob pattern visit its [wiki page](https://en.wikipedia.org/wiki/Glob_(programming)).

### 📎 Options

```
-b, --base          The root folder in which the search process begins.
                    On default, it starts at the current folder but you 
                    can specify manually to scan the whole system with
                    `--base "/"`.
-s, --stop          Stop the search process on the first match.
-m, --max           The search process stops automatically when it finds
                    more than 10000 matches. You can extend this limit
                    with this flag.
```



## Sources

* [Find files with Go | 05.03.2020](https://socketloop.com/tutorials/golang-find-files-by-name-cross-platform-example)
* [filepath package | 06.03.2020](https://golang.org/pkg/path/filepath/#pkg-overview)
* [go-flags package | 06.03.2020](https://github.com/jessevdk/go-flags)
* [go-flags api | 06.03.2020](https://godoc.org/github.com/jessevdk/go-flags)
* [Understanding sync.waitGroup | 06.03.2020](https://stackoverflow.com/a/25234899)
* [glob.go package | 06.03.2020](https://github.com/gobwas/glob)
* [The Go Gopher Logo - Renee French | 13.04.2020](https://commons.wikimedia.org/wiki/File:Gogophercolor.png)
* [Folder icon - OpenClipart - Pixabay](https://pixabay.com/images/id-146153/)

## License

MIT

---

> GitHub [kurbaniec](https://github.com/kurbaniec-tgm) &nbsp;&middot;&nbsp;
> Mail [at.kacper.urbaniec@gmail.com](mailto:at.kacper.urbaniec@gmail.com)