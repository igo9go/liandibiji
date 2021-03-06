// LianDi - 链滴笔记，连接点滴
// Copyright (c) 2020-present, b3log.org
//
// LianDi is licensed under Mulan PSL v2.
// You can use this software according to the terms and conditions of the Mulan PSL v2.
// You may obtain a copy of Mulan PSL v2 at:
//         http://license.coscl.org.cn/MulanPSL2
// THIS SOFTWARE IS PROVIDED ON AN "AS IS" BASIS, WITHOUT WARRANTIES OF ANY KIND, EITHER EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO NON-INFRINGEMENT, MERCHANTABILITY OR FIT FOR A PARTICULAR PURPOSE.
// See the Mulan PSL v2 for more details.

package main

type dirs struct {
	*BaseCmd
}

func (cmd *dirs) Exec() {
	ret := NewCmdResult(cmd.Name(), cmd.id)

	data := []map[string]interface{}{}
	for _, dir := range Conf.Dirs {
		data = append(data, map[string]interface{}{
			"dir": dir,
		})
	}
	ret.Data = data
	Push(ret.Bytes())
}

func (cmd *dirs) Name() string {
	return "dirs"
}
