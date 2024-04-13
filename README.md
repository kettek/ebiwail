This is a barebones PoC of using Ebitengine + Wails. It is scant on features and any enhancements are left up to the end-user.

Note that an iframe wasn't used so as to allow for javascript interop. It might be possible to communicate to/from the iframe, but I cannot recall so haven't investigated.

Also note that it might be possible to restructure this to configure `vite` to automatically call the build steps below...

## Build ebitengine game
Within the `game` subdirectory, issue the following commands:

### Unix/Linux:

Issue this each time you want to rebuild your game.

```
env GOOS=js GOARCH=wasm go build -o ../frontend/src/assets/game.wasm
```

Issue this once.
```
cp $(go env GOROOT)/misc/wasm/wasm_exec.js ../frontend/src/assets/
```

### On Windows PowerShell:

Issue this each time you want to rebuild your game.
```
$Env:GOOS = 'js'
$Env:GOARCH = 'wasm'
go build -o ../frontend/src/assets/game.wasm
Remove-Item Env:GOOS
Remove-Item Env:GOARCH
```

Issue this once.
```
$goroot = go env GOROOT
cp $goroot\misc\wasm\wasm_exec.js ../frontend/src/assets/
```

## Run Wails

Issue `wails dev` from the project root directory and enjoy.