# compareversion
CLI tool for comparing versions

## usage
Use `go build` to build executable inside compareversion directory.

Use `./compareversions 1.3.4 1.3.5` to run CLI tool.
The tool will return 1 to stdout in case version1 > version2, -1 in case
version1 < version2. 0 - otherwise.

Use `go test` to run added tests.