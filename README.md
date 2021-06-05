# VFA-314 FULL METAL BITCHES WEB APPLICATION
> Development with Golang, Web Assembly and SCSS.
> Using Web-Framework: Go-Fiber.

## Requirements
1. WSL2: Ubuntu 20.04 LTS
2. Latest Go-Lang

## How Compile and Run?
1. Prepare
Make Typical Go-Lang Workspace Directory Structure at Your Home Root

```
cd ~
mkdir -p go/bin go/pkg go/src
```

2. Cloning
Clone This Repository in Go-Lang Source Directory

```
git clone https://www.github.com/kim-ahri/fmb
```

3. Advance to Workspace and Make Go Mod File
You Can Install All Defendencies With Using Below Commands

```
go mod init
go mod tidy
```

4. Bring "wasm_exec.js" to Workspace
This File Need to Run WASM. Will Async Fetching in Header of "page.html"

```
cp "$(go env GOROOT)/misc/wasm/wasm_exec.js" ./static
```

5. Open Visual Code and Do Test
Please Check "makefile" Before Run Build Command.

```
code .
make build
```