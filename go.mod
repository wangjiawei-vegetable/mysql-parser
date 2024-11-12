module mysql-parser

go 1.21

require (
	github.com/antlr4-go/antlr/v4 v4.13.0
	github.com/pingcap/tidb/parser v0.0.0-20221101143359-5b0be9af540e
	github.com/pkg/errors v0.9.1
	github.com/stretchr/testify v1.8.2
)

require (
	github.com/cznic/mathutil v0.0.0-20181122101859-297441e03548 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pingcap/errors v0.11.5-0.20210425183316-da1aaba5fb63 // indirect
	github.com/pingcap/log v0.0.0-20210625125904-98ed8e2eb1c7 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/remyoudompheng/bigfft v0.0.0-20200410134404-eec4a21b6bb0 // indirect
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
	go.uber.org/zap v1.18.1 // indirect
	golang.org/x/exp v0.0.0-20230515195305-f3d0a9c9a5cc // indirect
	golang.org/x/text v0.19.0 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace github.com/pingcap/tidb/parser => github.com/bytebase/tidb/parser v0.0.0-20230914094316-ec1081216cfb

replace github.com/antlr4-go/antlr/v4 => github.com/bytebase/antlr/v4 v4.0.0-20231103101006-5fe1a93b199f
