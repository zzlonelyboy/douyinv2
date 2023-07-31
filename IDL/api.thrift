namespace go douyinv2
struct User {
  1:required i64 id = 1; // 用户id
  2:required string name = "2"; // 用户名称
  3:optional i64 follow_count = 3; // 关注总数
  4:optional i64 follower_count = 4; // 粉丝总数
  5:required bool is_follow = false; // true-已关注，false-未关注
  6:optional string avatar = "6"; //用户头像
  7:optional string background_image = "7"; //用户个人页顶部大图
  8:optional string signature = "8"; //个人简介
  9:optional i64 total_favorited = 9; //获赞数量
  10:optional i64 work_count = 10; //作品数量
  11:optional i64 favorite_count = 11; //点赞数量
}
struct FriendUser{
    1:required User user;
    2:optional string message="1";
    3:required i64 msgType=1;
}
struct UserRegisterRequest {
  1:required string username = "1"; // 注册用户名，最长32个字符
  2:required string password = "2"; // 密码，最长32个字符
}
struct UserRegisterResponse{
      1:required i32 status_code = 1; // 状态码，0-成功，其他值-失败
      2:optional string status_msg = "2"; // 返回状态描述
      3:required i64 user_id = 3; // 用户id
      4:required string token = "4"; // 用户鉴权token
}
struct UserLoginRequest {
  required string username = "1"; // 登录用户名
  required string password = "2"; // 登录密码
}

struct UserLoginResponse {
  required i32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = "2"; // 返回状态描述
  required i64 user_id = 3; // 用户id
  required string token = "4"; // 用户鉴权token
}
struct UserInfoRequest {
  required i64 user_id = 1; // 用户id
  required string token = "2"; // 用户鉴权token
}

struct UserInfoResponse {
  required i32 status_code = 1; // 状态码，0-成功，其他值-失败
  optional string status_msg = "2"; // 返回状态描述
  required User user; // 用户信息
}
struct Video {
  1:required i64 id = 1; // 视频唯一标识
  2:required User author ; // 视频作者信息
  3:required string play_url = "null"; // 视频播放地址
  4:required string cover_url = "null"; // 视频封面地址
  5:required i64 favorite_count = 0; // 视频的点赞总数
  6:required i64 comment_count = 0 // 视频的评论总数
  7:required bool is_favorite = false; // true-已点赞，false-未点赞
  8:required string title; // 视频标题
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
  1:required string token  // 用户鉴权token
  2:required binary data; // 视频数据
  3:required string title = "null"; // 视频标题
}

struct VideoPublishResponse {
  1:required i32 status_code = 1; // 状态码，0-成功，其他值-失败
  2:optional string status_msg = "error"; // 返回状态描述
}

struct VideoPublishedListRequest {
  1:required i64 user_id = 1; // 用户id
  2:required string token ; // 用户鉴权token
}

struct VideoPublishedListResponse {
  1:required i32 status_code = 1; // 状态码，0-成功，其他值-失败
  2:optional string status_msg = "error"; // 返回状态描述
  3:list< Video> video_list = []; // 用户发布的视频列表
}
struct FavoriteActionRequest{
    required string token
    required i64 video_id
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

struct Comment {
  1:required i64 id ; // 视频评论id
  2:required User user  // 评论用户信息
  3:required string content ; // 评论内容
  4:required string create_date ; // 评论发布日期，格式 mm-dd
  5:required i64 like_count ;// 该评论点赞数量
  6:required i64 tease_count ; // 该评论diss数量
}
struct CommentActionRequest {
  required string token ; // 用户鉴权token
  required i64 video_id ; // 视频id
  required i32 action_type=1 ; // 1-发布评论，2-删除评论
  optional string comment_text="test" ; // 用户填写的评论内容，在action_type=1的时候使用
  optional i64 comment_id=-1; // 要删除的评论id，在action_type=2的时候使用
}

struct CommentActionResponse {
  required i32 status_code=-1 ; // 状态码，0-成功，其他值-失败
  optional string status_msg="error" ; // 返回状态描述
  optional Comment comment ; // 评论成功返回评论内容，不需要重新拉取整个列表
}

struct CommentListRequest{
    required string token="2"
    required i64 video_id=-1
}
struct CommentListResponse{
    required i32 status_code=-1
    optional string status_msg="error"
    list<Comment> comment_list
}
struct RelationActionRequest{
    required string token="null"
    required i64 to_user_id=-1
    required i32 action_type=1
}
struct RelationActionResponse{
    required i32 status_code = 1; // 状态码，0-成功，其他值-失败
    optional string status_msg = "2"; // 返回状态描述
}
struct FollowListRequest{
    required string token="null"
    required i64 user_id=-1,
}
struct FollowListResponse{
     required i32 status_code = 1; // 状态码，0-成功，其他值-失败
     optional string status_msg = "2"; // 返回状态描述
     list<User> user_list=[]
}
struct FollowerListRequest{
    required string token="null"
    required i64 user_id=-1,
}
struct FollowerListResponse{
     required i32 status_code = 1; // 状态码，0-成功，其他值-失败
     optional string status_msg = "2"; // 返回状态描述
     list<User> user_list=[]
}
struct FriendListRequest{
    required string token="null"
    required i64 user_id=-1
}
struct FriendListResponse{
     required i32 status_code = 1; // 状态码，0-成功，其他值-失败
     optional string status_msg = "2"; // 返回状态描述
     list<FriendUser> user_list=[]
}

service ApiService {
    UserRegisterResponse RegisterUser(1: UserRegisterRequest req) (api.post="/douyin/user/register/")
    UserLoginResponse LoginUser(1: UserLoginRequest req) (api.post="/douyin/user/login/")
    UserInfoResponse UserInfo(1:UserInfoRequest req)(api.get="/douyin/user/")
    VideoPublishResponse VideoPublish(1:VideoPublishRequest req)(api.post="/douyin/publish/action/")
    VideoPublishedListResponse GetVideoList(1:VideoPublishedListRequest req) (api.get="/douyin/publish/list/")
    FeedResponse Feed(1:FeedRequest req)(api.get="/douyin/feed/")
    FavoriteActionResponse LikeAction(1:FavoriteActionRequest req)(api.post="/douyin/favorite/action/")
    FavoriteVideoListResponse LikeVideos(1:FavoriteVideoListRequest req)(api.get="/douyin/favorite/list/")
    CommentActionResponse CommentAction(1:CommentActionRequest req)(api.post="/douyin/comment/action/")
    CommentListResponse CommentList(1:CommentListRequest req)(api.get="/douyin/comment/list/")
    RelationActionResponse RelationAction(1:RelationActionRequest req)(api.post="/douyin/relation/action/")
    FollowListResponse FollowList(1:FollowListRequest req)(api.get="/douyin/relation/follow/list/")
    FollowerListResponse FollowerList(1:FollowerListRequest req)(api.get="/douyin/relation/follower/list/")
    FriendListResponse  FriendList(1:FriendListRequest req)(api.get="/douyin/relation/friend/list/")
}
