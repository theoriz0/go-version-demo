package version

import (
	"encoding/json"
	"fmt"
	"runtime"
)

var (
	// GitVersion 是语义化的版本号.
	GitVersion = "v0.0.0-master+$Format:%h$"
	// BuildDate 是 ISO8601 格式的构建时间, $(date -u +'%Y-%m-%dT%H:%M:%SZ') 命令的输出.
	BuildDate = "1970-01-01T00:00:00Z"
	// GitCommit 是 Git 的 SHA1 值，$(git rev-parse HEAD) 命令的输出.
	GitCommit = "$Format:%H$"
	// GitTreeState 代表构建时 Git 仓库的状态，可能的值有：clean, dirty.
	GitTreeState = ""
)

type Info struct {
	GitVersion   string `json:"gitVersion"`
	GitCommit    string `json:"gitCommit"`
	GitTreeState string `json:"gitTreeState"`
	BuildDate    string `json:"buildDate"`
	GoVersion    string `json:"goVersion"`
	Compiler     string `json:"compiler"`
	Platform     string `json:"platform"`
}

func Get() Info {
	// 以下变量通常由 -ldflags 进行设置
	return Info{
		GitVersion:   GitVersion,
		GitCommit:    GitCommit,
		GitTreeState: GitTreeState,
		BuildDate:    BuildDate,
		GoVersion:    runtime.Version(),
		Compiler:     runtime.Compiler,
		Platform:     fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH),
	}
}

func (info Info) ToJSON() string {
	s, _ := json.Marshal(info)

	return string(s)
}

func PrintJSON() {
	fmt.Println(Get().ToJSON())
}
