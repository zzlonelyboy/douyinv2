struct CommentActionRequest {
  required string token = 1; // 用户鉴权token
  required int64 video_id = 2; // 视频id
  required int32 action_type = 3; // 1-发布评论，2-删除评论
  optional string comment_text = 4; // 用户填写的评论内容，在action_type=1的时候使用
  optional int64 comment_id = 5; // 要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse {
  required int32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = 2; // 返回状态描述
  optional Comment comment = 3; // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct Comment {
  required int64 id = 1; // 视频评论id
  required User user =2; // 评论用户信息
  required string content = 3; // 评论内容
  required string create_date = 4; // 评论发布日期，格式 mm-dd
  required int64 like_count = 5;// 该评论点赞数量
  required int64 tease_count = 6; // 该评论diss数量
}