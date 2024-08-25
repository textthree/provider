package types

import "github.com/textthree/provider/core"

type BeforStartCallback func(c core.Container)

// clog 配置
type Clog struct {
	Level string
}
