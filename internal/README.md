# `/internal`

Private application and library code. This is the code you don't want others importing in their applications or libraries. Note that this layout pattern is enforced by the Go compiler itself. See the Go 1.4 [`release notes`](https://golang.org/doc/go1.4#internalpackages) for more details. Note that you are not limited to the top level `internal` directory. You can have more than one `internal` directory at any level of your project tree.

You can optionally add a bit of extra structure to your internal packages to separate your shared and non-shared internal code. It's not required (especially for smaller projects), but it's nice to have visual clues showing the intended package use. Your actual application code can go in the `/internal/app` directory (e.g., `/internal/app/myapp`) and the code shared by those apps in the `/internal/pkg` directory (e.g., `/internal/pkg/myprivlib`).

For example, if you have a directory structure like this:
```go
myapp/
    internal/
        util.go
    main.go 
```

myapp/internal/util.go can only be referenced within the myapp module. For example, main.go can import "myapp/internal/util".
But code from other modules cannot import myapp/internal. If you try to import "myapp/internal/util", the compiler will report an error:
```bash
cannot use "myapp/internal/util" as external package name: myapp/internal/util.go: does not export package internal 
```

Examples:

* https://github.com/hashicorp/terraform/tree/main/internal
* https://github.com/influxdata/influxdb/tree/master/internal
* https://github.com/perkeep/perkeep/tree/master/internal
* https://github.com/jaegertracing/jaeger/tree/main/internal
* https://github.com/moby/moby/tree/master/internal
* https://github.com/satellity/satellity/tree/main/internal

## `/internal/pkg`

Examples:

* https://github.com/hashicorp/waypoint/tree/main/internal/pkg