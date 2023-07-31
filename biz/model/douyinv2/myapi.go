package douyinv2

type VideoPublishRequestv2 struct {
	ID       int64  `thrift:"ID,1,required" frugal:"1,required,i64" json:"ID"`
	Filepath string `thrift:"data,2,required" frugal:"2,required,binary" json:"data"`
	Title    string `thrift:"title,3,required" frugal:"3,required,string" json:"title"`
	Coverurl string `thrift:"cover_url,4,required" frugal:"3,required" json:"Coverurl"`
}
