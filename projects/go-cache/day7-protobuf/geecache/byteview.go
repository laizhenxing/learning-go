package geecache

type ByteView struct {
	b []byte	// 存储真实的缓存值,只读
}

// 实现 Value 接口
func (v ByteView) Len() int {
	return len(v.b)
}

// 返回一个拷贝，防止缓存值被修改
func (v ByteView) ByteSlice() []byte {
	return cloneBytes(v.b)
}

func (v ByteView) String() string {
	return string(v.b)
}

func cloneBytes(b []byte) []byte {
	c := make([]byte, len(b))
	copy(c, b)
	return c
}
