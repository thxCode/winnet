$ErrorActionPreference = 'Stop'

Import-Module -WarningAction Ignore -Name "$PSScriptRoot\utils.psm1"

Invoke-Script -File "$PSScriptRoot\version.ps1"

$SRC_PATH = (Resolve-Path "$PSScriptRoot\..").Path
Push-Location $SRC_PATH

Remove-Item -Path "$SRC_PATH\bin\*" -Force -ErrorAction Ignore
$null = New-Item -Type Directory -Path bin -ErrorAction Ignore
$env:GOARCH = $env:ARCH
$env:GOOS = 'windows'
$env:CGO_ENABLED = 0
$LINKFLAGS = ('-s -w -X github.com/thxcode/winnet/cmd/main.Version={0} -X github.com/thxcode/winnet/cmd/main.Commit={1} -extldflags "-static"' -f $env:VERSION,  $env:COMMIT)
go build -ldflags "$LINKFLAGS" -o bin\winnet.exe cmd\main.go
if (-not $?) {
    Log-Fatal "go build failed!"
}

Pop-Location
