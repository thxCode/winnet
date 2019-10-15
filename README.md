# winnet

[![Build status](https://ci.appveyor.com/api/projects/status/5dqfpq07sqsoqus1?svg=true)](https://ci.appveyor.com/project/thxCode/winnet)

A Go implementation replaces [rakelkar/gonetsh](https://github.com/rakelkar/gonetsh) to support none-English locale Windows OS.

## Test

```powershell
# test via rancher/dapper
./make.bat test

# test on local
./scripts/test.ps1
```

## Build

```powershell
# build via rancher/dapper
./make.bat build

# build on local
./scripts/build.ps1
```

## License

Copyright (c) 2014-2019 The winnet Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

[http://www.apache.org/licenses/LICENSE-2.0](http://www.apache.org/licenses/LICENSE-2.0)

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.