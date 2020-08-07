package gazelle

import (
	"flag"
	"github.com/bazelbuild/bazel-gazelle/config"
	"github.com/bazelbuild/bazel-gazelle/rule"
)

type TsGenConfig struct {
	TsLibraryLoad string
}

func GetTsGenConfig(c *config.Config) *TsGenConfig {
	ts := c.Exts["ts"]
	if ts == nil {
		return nil
	}
	return ts.(*TsGenConfig)
}

func (s *tslang) CheckFlags(fs *flag.FlagSet, c *config.Config) error {
	return nil
}

func (*tslang) RegisterFlags(fs *flag.FlagSet, cmd string, c *config.Config) {
	ts := &TsGenConfig{}
	c.Exts["ts"] = ts

	fs.StringVar(&ts.TsLibraryLoad, "load", "@npm//@bazel/typescript:index.bzl", "Load site to use for ts_library")
}

func (*tslang) KnownDirectives() []string {
	return []string{"ts_library"}
}

func (*tslang) Configure(c *config.Config, rel string, f *rule.File) {}
