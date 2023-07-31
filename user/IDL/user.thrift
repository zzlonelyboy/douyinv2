namespace go user
struct User {
  1:required i64 id = 1; // 用户id
  2:required string name = "2"; // 用户名称
  3:optional i64 follow_count = 3; // 关注总数
  4:optional i64 follower_count = 4; // 粉丝总数
  5:required bool is_follow = false; // true-已关注，false-未关注
  6:optional string avatar = "6"; //用户头像
  7:optional string background_image = "7"; //用户个人页顶部大图
  8:optional string signature = "hello"; //个人简介
  9:optional i64 total_favorited = 9; //获赞数量
  10:optional i64 work_count = 10; //作品数量
  11:optional i64 favorite_count = 11; //点赞数量
}
struct FriendUser{
    1:required i64 id = 1; // 用户id
    2:required string name = "2"; // 用户名称
    3:optional i64 follow_count = 3; // 关注总数
    4:optional i64 follower_count = 4; // 粉丝总数
    5:required bool is_follow = false; // true-已关注，false-未关注
    6:optional string avatar = "6"; //用户头像
    7:optional string background_image = "7"; //用户个人页顶部大图
    8:optional string signature = "hello"; //个人简介
    9:optional i64 total_favorited = 9; //获赞数量
    10:optional i64 work_count = 10; //作品数量
    11:optional i64 favorite_count = 11; //点赞数量
    12:optional string message="1";
    13:required i64 msgType=1;
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
struct UserCountChangeRequest{
    required i64 user_id
    required bool ifadd //增加为true，否则为false
    required string filed
    required string token="2"
}
struct UserCountChangeResponse{
    required i32 status_code=1
    optional string status_msg="2"
}
struct RelationActionRequest{
    required i64 fromid=-1
    required i64 toid=-1
    required i32 action_type=1
}
struct RelationActionResponse{
    required i32 status_code = 1; // 状态码，0-成功，其他值-失败
    optional string status_msg = "2"; // 返回状态描述
}
struct FollowListRequest{
    required i64 userid=-1,
}
struct FollowListResponse{
     required i32 status_code = 1; // 状态码，0-成功，其他值-失败
     optional string status_msg = "2"; // 返回状态描述
     list<User> user_list=[]
}
struct FollowerListRequest{
    required i64 userid=-1,
}
struct FollowerListResponse{
     required i32 status_code = 1; // 状态码，0-成功，其他值-失败
     optional string status_msg = "2"; // 返回状态描述
     list<User> user_list=[]
}
struct FriendListRequest{
    required i64 ID
}
struct FriendListResponse{
     required i32 status_code = 1; // 状态码，0-成功，其他值-失败
     optional string status_msg = "2"; // 返回状态描述
     list<FriendUser> user_list=[]
}

service ApiService {
    UserRegisterResponse RegisterUser(1: UserRegisterRequest req) (api.post="/douyin/user/register")
    UserLoginResponse LoginUser(1: UserLoginRequest req) (api.post="/douyin/user/login")
    UserInfoResponse UserInfo(1:UserInfoRequest req)(api.get="/douyin/user/")
    UserCountChangeResponse CountAdd(1:UserCountChangeRequest req)
    RelationActionResponse RelationAction(1:RelationActionRequest req)
    FollowListResponse FollowList(1:FollowListRequest req)
    FollowerListResponse FollowerList(1:FollowerListRequest req)
    FriendListResponse  FriendList(1:FriendListRequest req)
}

