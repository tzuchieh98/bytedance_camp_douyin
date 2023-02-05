namespace go douyin.base

// hz new -idl idl/base.thrift --snake_tag -mod github.com/linzijie1998/bytedance_camp_douyin

struct UserRegisterReq {
    1: required string Username (api.query="username");
    2: required string Password (api.query="password");
}

struct UserRegisterResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: required i64 UserID;
    4: required string Token;
}

struct UserLoginReq {
    1: required string Username (api.query="username");
    2: required string Password (api.query="password");
}

struct UserLoginResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: required i64 UserID;
    4: required string Token;
}

struct UserInfoReq {
    1: required i64 UserID (api.query="user_id");
    2: required string Token (api.query="token");
}

struct UserInfoResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: required User User;
}

struct User {
    1: required i64 ID;
    2: required string Name;
    3: optional i64 FollowCount;
    4: optional i64 FollowerCount;
    5: required bool IsFollow;
}

service UserService {
    UserRegisterResp UserRegister(1: UserRegisterReq req) (api.post="/douyin/user/register/");
    UserLoginResp UserLogin(1: UserLoginReq req) (api.post="/douyin/user/login/");
    UserInfoResp UserInfo(1: UserInfoReq req) (api.get="/douyin/user/");
}


struct PublishActionReq {
    1: required string Token (api.query="token");
    2: required binary Data (api.query="data");
    3: required string Title (api.query="title");
}

struct PublishActionResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
}

struct PublishListReq {
    1: required i64 UserID (api.query="user_id");
    2: required string Token (api.query="token");
}

struct PublishListResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: required list<Video> VideoList;
}

struct Video {
    1: required i64 ID;
    2: required User Author;
    3: required string PlayURL;
    4: required string CoverURL;
    5: required i64 FavoriteCount;
    6: required i64 CommentCount;
    7: required bool IsFavorite;
    8: required string Title;
}

service PublishService {
    PublishActionResp PublishAction(1: PublishActionReq req) (api.post="/douyin/publish/action/");
    PublishListResp PublishList(1: PublishListReq req) (api.get="/douyin/publish/list/");
}

struct FeedReq {
    1: optional i64 LatestTime (api.query="latest_time");
    2: optional string Token (api.query="token");
}

struct FeedResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: required list<Video> VideoList;
    4: optional i64 NextTime;
}

service FeedService {
    FeedResp Feed(1: FeedReq req) (api.get="/douyin/feed/");
}
