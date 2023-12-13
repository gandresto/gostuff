# hellogo

## Runnable packages

For a package to capable of being built and run as a standalone package, it
needs to be a main package with a `main()` function

## Run code

Run compiles and runs the named main Go package.

```bash
go run main.go
```

## Build and run

Build compiles the packages named by the import paths,
along with their dependencies, but it does not install the results.

```bash
go build
```

Run the compiled file

```bash
./hellogo
```

Compile and run

```bash
go build && ./hellogo
```

## Install binary locally

Install compiles and installs the packages named by the import paths.

Executables are installed in the directory named by the GOBIN environment
variable, which defaults to $GOPATH/bin or $HOME/go/bin if the GOPATH
environment variable is not set. Executables in $GOROOT
are installed in $GOROOT/bin or $GOTOOLDIR instead of $GOBIN.

```bash
go install
```