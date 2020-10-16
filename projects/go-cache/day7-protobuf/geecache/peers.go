package geecache

import "day7-protobuf/geecache/pb"

type PeerPicker interface {
	// 根据传入的key选择相应节点 PeerGetter
	PickPeer(key string) (peer PeerGetter, ok bool)
}

// 相当于流程中的HTTP客户端
type PeerGetter interface {
	// 从对应的group查找缓存值
	Get(in *pb.Request, out *pb.Response) error
}