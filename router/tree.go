package router

import (
	"strings"
)

type nodeType uint8

const (
	static nodeType = iota // default
	root
	param
)

type Node struct {
	path            string                 //当前节点路径
	level           int                    //节点等级
	fullPath        string                 //全路径
	wildChild       bool                   //是否是子节点
	nType           nodeType               //节点类型
	paramName       string                 //如果为参数节点  存储该参数名字
	children        []*Node                // 节点的子节点列表
	indices         []string               //下级子节点索引
	parent          *Node                  //父节点
	handlerPipeline *RouterHandlerPipeline //处理器链 包含处理器和过滤器链
}

func NewRoot() *Node {
	return &Node{
		path:      "/",   //当前节点路径
		fullPath:  "/",   //全路径
		level:     0,     //根节点为0
		wildChild: false, //是否是子节点
		nType:     root,  //节点类型
		paramName: "",    //如果为参数节点  存储该参数名字
		children:  nil,   // 节点的子节点列表
		/*indices:nil,				//下级子节点索引
		parent:nil,					//父节点
		handlerPipeline:nil, 				//处理器*/
	}
}

//考虑如下情况
//1.static	  /user/add
//2.root     /*
//3.param    /user/:uid  /user/:uid/info
func (root *Node) addNode(pattern string, handlerPipeline *RouterHandlerPipeline) {
	if "/*" == pattern {
		panic("路径错误")
	}
	//例子 /user/:id
	subPaths := strings.Split(pattern, "/")[1:]
	root.add(subPaths, handlerPipeline)
}

// subPaths=[users :id xxx]  /user/:id/xxx
func (n *Node) add(subPaths []string, handlerPipeline *RouterHandlerPipeline) {
	//如果是根节点 肯定白搭
	var paths_len = len(subPaths)
	if n.level < paths_len {
		//如果比当前path的长度小 说明是该节点的上级节点
		//判断路径长度 跟当前节点等级
		//如果路径为3
		//根节点 0 /
		//一级节点 1/user
		//二级节点 2 /user/:id
		//三级节点 3 /user/:id/xx

		if paths_len-n.level > 1 {
			var zf *Node
			//新增节点的祖父节点
			if len(n.children) > 0 {
				for _, c := range n.children {
					//还有一种情况就是 该节点已经 param类型节点 那么我们就认为已经存在了
					//也就是在同一级中不允许出现2个以上的param节点
					//例如 /questions/:qaid/answer/:id
					// /questions/:qrid/reply/:id
					//此时我们认为 两者是重复的
					zfpath := subPaths[c.level-1]
					if strings.Contains(zfpath, ":") && c.nType == param {
						if strings.Split(zfpath, ":")[1] != c.paramName {
							panic("不允许出现同一级节点不同路径参数名的情况,已经存在一个路径参数:" + c.paramName + "--不要试图添加" + zfpath)
						} else {
							//不需要
							zf = c
							break
						}
					}
					if c.path == subPaths[n.level] {
						//相同则说明 该 祖父节点已经有了
						zf = c
						//不需要创建
						break
					}
				}
			}
			if zf == nil {
				zf = newChild(n, subPaths, false, handlerPipeline)
			}
			zf.add(subPaths, handlerPipeline)
		} else {
			//当前的父节点
			//判断之前父节点中是否已经创建过了该节点了
			checkPath(n, subPaths[n.level])

			newChild(n, subPaths, true, handlerPipeline)
		}
	} else {
		//do nothing
	}
}
func checkPath(parent *Node, path string) {
	indices := parent.indices
	for i, is := range indices {
		if strings.Contains(is, ":") && strings.Contains(path, ":") {
			panic("该节点已经存在了" + parent.fullPath + indices[i] + "不允许再出现叶子节点为" + parent.fullPath + path)
		}

		if is == path {
			panic("该节点已经存在了" + parent.fullPath + indices[i])
		}

	}
}

//新增节点
func newChild(parent *Node, subPaths []string, directParent bool, handlerPipeline *RouterHandlerPipeline) *Node {
	var h *RouterHandlerPipeline
	//是否是直接父节点
	if directParent {
		h = handlerPipeline

	} else {
		h = parent.handlerPipeline
	}
	path := subPaths[parent.level]
	var nt nodeType
	var pname string
	if strings.HasPrefix(path, ":") {
		nt = param
		pname = strings.Split(path, ":")[1]
	} else {
		nt = static
	}

	//判断新增的节点
	child := &Node{
		path:            path,
		fullPath:        parent.fullPath + path + "/",
		parent:          parent,
		level:           parent.level + 1,
		wildChild:       false,
		nType:           nt,
		paramName:       pname,
		handlerPipeline: h,
	}
	//父节点新增子节点索引
	parent.indices = append(parent.indices, subPaths[parent.level])
	//将该节点添加到父节点中
	parent.children = append(parent.children, child)
	return child
}

func (root *Node) Search(pattern string) (*Node, *RouterHandlerPipeline, PathParam) {
	subPaths := strings.Split(pattern, "/")[1:]

	return root.search0(subPaths)
}
func (root *Node) search0(subPaths []string) (rn *Node, h *RouterHandlerPipeline, pathParam PathParam) {
	pathParam = map[string]string{}
walk:
	for {
		if len(subPaths)-root.level > 1 {
			//祖父节点
			if len(root.children) > 0 {
				//没有子节点 啥也不干了 有则循环
				for _, c := range root.children {
					if c.path == subPaths[root.level] {
						//相同则说明 该 祖父节点已经有了
						//调用祖父节点去查询
						if checkIndices(c.indices, subPaths[c.level]) {
							root = c
							continue walk
						}
					}

				}
				for _, c := range root.children {
					if c.nType == param {
						root = c
						pathParam[c.paramName] = subPaths[c.level-1]
						continue walk
					}
				}
			} else {
				//没有子节点 啥也不干了
			}
		} else {
			//父节点
			for _, c := range root.children {
				if c.nType == param {
					pathParam[c.paramName] = subPaths[root.level]
					rn = c
					h = c.handlerPipeline
				}

				if c.path == subPaths[root.level] {
					return c, c.handlerPipeline, nil
				}
			}

		}
		return rn, h, pathParam
	}
}
func checkIndices(indices []string, path string) bool {
	for _, is := range indices {
		if strings.Contains(is, ":") {
			return true
		}
		if is == path {
			return true
		}
	}
	return false
}

func prefix(s string, prefix string) string {
	if !strings.HasPrefix(s, prefix) {
		return prefix + s
	}
	return s
}

func suffix(s string, suffix string) string {
	if !strings.HasSuffix(s, suffix) {
		return s + suffix
	}
	return s
}
