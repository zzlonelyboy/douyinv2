namespace go video
struct Video {
  1:required i64 id = 1; // 视频唯一标识
  2:required User author ; // 视频作者信息
  3:required string play_url = "null"; // 视频播放地址
  4:required string cover_url = "null"; // 视频封面地址
  5:required i64 favorite_count = 5; // 视频的点赞总数
  6:required i64 comment_count = 6; // 视频的评论总数
  7:required bool is_favorite = 7; // true-已点赞，false-未点赞
  8:required string title; // 视频标题
}

struct User {
  1:required i64 id = 1; // 用户id
  2:required string name; // 用户名称
  3:optional i64 follow_count = 3; // 关注总数
  4:optional i64 follower_count = 4; // 粉丝总数
  5:required bool is_follow = 5; // true-已关注，false-未关注
  6:optional string avatar = "6"; //用户头像
  7:optional string background_image = "7"; //用户个人页顶部大图
  8:optional string signature = "hello"; //个人简介
  9:optional i64 total_favorited = 9; //获赞数量
  10:optional i64 work_count = 10; //作品数量
  11:optional i64 favorite_count = 11; //点赞数量
}


struct Comment {
  1:required i64 id = 1; // 视频评论id
  2:required User user ; // 评论用户信息
  3:required string content; // 评论内容
  4:required string create_date ; // 评论发布日期，格式 mm-dd
  5:required i64 like_count ;// 该评论点赞数量
  6:required i64 tease_count; // 该评论diss数量
}
struct FeedRequest {
  1:optional i64 latest_time = 1; // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
  2:optional string token = "null" // 可选参数，登录用户设置
}

struct FeedResponse {
  1:required i32 status_code = 1; // 状态码，0-成功，其他值-失败
  2:optional string status_msg = "error"; // 返回状态描述
  3:list<Video> video_list = []; // 视频列表
  4:optional i64 next_time = 4; // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
}
struct VideoPublishRequest {
  required i64 ID  // 用户鉴权token
  required string Filepath; // 视频数据
  required string title = "null"; // 视频标题
  required string Cover="null";
}

struct VideoPublishResponse {
  required i32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = "error"; // 返回状态描述
}
struct VideoPublishedListRequest {
  required i64 user_id = 1; // 用户id
  required string token ; // 用户鉴权token
}

struct VideoPublishedListResponse {
  required i32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = "error"; // 返回状态描述
  list< Video> video_list = []; // 用户发布的视频列表
}
struct FavoriteActionRequest{
    required i64 ID
    required i64 Video_ID
    required i32 action_type
}
struct FavoriteActionResponse{
    required i32 status_code = 1; // 状态码，0-成功，其他值-失败
    optional string status_msg = "error"; // 返回状态描述
}
struct FavoriteVideoListRequest{
    required i64 user_id
    required string token
}
struct FavoriteVideoListResponse{
    required i32 status_code = 1; // 状态码，0-成功，其他值-失败
    optional string status_msg = "error"; // 返回状态描述
    list< Video> video_list = []; // 用户发布的视频列表
}
struct DoesFavoriteRequest{
    required i64 user_id
    required i64 video_id
}
struct DoesFavoriteResponse{
       required bool iffollow
   }
struct CommentActionRequest {
  required i64  userid=-1 ; // 用户鉴权token
  required i64 video_id=-1 ; // 视频id
  required i32 action_type=1 ; // 1-发布评论，2-删除评论
  optional string comment_text="ss" ; // 用户填写的评论内容，在action_type=1的时候使用
  optional i64 comment_id=-1; // 要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse {
  required i32 status_code ; // 状态码，0-成功，其他值-失败
  optional string status_msg="error" ; // 返回状态描述
  optional Comment comment ; // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct CommentListRequest{
    required i64 video_id
}
struct CommentListResponse{
    required i32 status_code
    optional string status_msg="error"
    list<Comment> comment_list
}

service VideoService {
    VideoPublishResponse VideoPublish(1:VideoPublishRequest req)(api.post="/douyin/publish/action/")
    VideoPublishedListResponse GetVideoList(1:VideoPublishedListRequest req) (api.post="/douyin/publish/list/")
    FeedResponse Feed(1:FeedRequest req)(api.get="/douyin/feed/")
    FavoriteActionResponse LikeAction(1:FavoriteActionRequest req)
    FavoriteVideoListResponse LikeVideos(1:FavoriteVideoListRequest req)
    DoesFavoriteResponse IfLike(1:DoesFavoriteRequest req)
    CommentActionResponse CommentACtion(1:CommentActionRequest req)
    CommentListResponse CommentList(1:CommentListRequest req)
}