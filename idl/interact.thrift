namespace go douyin.interact

include "base.thrift"

struct FavoriteActionReq {
    1: required string Token (api.query="token");
    2: required i64 VideoID (api.query="video_id");
    3: required i32 ActionType (api.query="action_type");
}

struct FavoriteActionResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
}

struct FavoriteListReq {
    1: required i64 UserID (api.query="user_id");
    2: required string Token (api.query="token");
}

struct FavoriteListResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: required list<base.Video> VideoList;
}

service FavoriteService {
    FavoriteActionResp FavoriteAction(1: FavoriteActionReq req) (api.post="/douyin/favorite/action/");
    FavoriteListResp FavoriteList(1: FavoriteListReq req) (api.get="/douyin/favorite/list/");
}

struct Comment {
    1: required i64 ID;
    2: required base.User User;
    3: required string Content;
    4: required string CreateDate;
}

struct CommentActionReq {
    1: required string Token (api.query="token");
    2: required i64 VideoID (api.query="video_id");
    3: required i32 ActionType (api.query="action_type");
    4: optional string CommentText (api.query="comment_text");
    5: optional i64 CommentID (api.query="comment_id");
}

struct CommentActionResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: optional Comment Comment;
}

struct CommentListReq {
    1: required string Token (api.query="token");
    2: required i64 VideoID (api.query="video_id");
}

struct CommentListResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: required list<Comment> CommentList;
}

service CommentService {
    CommentActionResp CommentAction(1: CommentActionReq req) (api.post="/douyin/comment/action/");
    CommentListResp CommentList(1: CommentListReq req) (api.get="/douyin/comment/list/");
}
