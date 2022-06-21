package xworkflow

type DagParser struct {
	
}

// ParseString 从dag json字符串
func (d DagParser) ParseString(dag string) (DagGraph,error)  {
	return DagGraph{},nil
}

// ParseFile 从文件读取dag定义文件
func (d DagParser) ParseFile(path string) (DagGraph,error) {
	return DagGraph{},nil
}