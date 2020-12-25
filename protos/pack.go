// Copyright (c) 2019,CAOHONGJU All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package protos

// Pack 表示流媒体包
type Pack interface {
	Size() int // 包长度
}
