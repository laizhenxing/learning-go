package main

import (
	"io"
	"time"
)

type Artifact interface {
	Title() string
	Creators() []string
	Created() time.Time
}

type Text interface {
	Pages() int
	Words() int
	PageSize() int
}

type Audio interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}

type Video interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
	Resolution() (x, y int)
}

// 提取Video 和 Audio 的公共功能
// 不需要对原来这两类的接口做出改变
type Streamer interface {
	Stream() (io.ReadCloser, error)
	RunningTime() time.Duration
	Format() string
}
