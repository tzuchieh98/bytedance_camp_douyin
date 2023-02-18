namespace go douyin.relation

include "base.thrift"

struct RelationActionReq {
    1: required string Token;
    2: required i64 ToUserID;
    3: required i32 ActionType;
}

struct RelationActionResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
}

struct RelationFollowListReq {
    1: required i64 UserID;
    2: required string Token;
}

struct RelationFollowListResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: required list<base.User> UserList;
}

struct RelationFollowerListReq {
    1: required i64 UserID;
    2: required string Token;
}

struct RelationFollowerListResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: required list<base.User> UserList;
}

struct FriendUser {
    1: required i64 ID;
    2: required string Name;
    3: optional i64 FollowCount;
    4: optional i64 FollowerCount;
    5: required bool IsFollow;
    6: required string Avatar;
    7: optional string BackgroundImage;
    8: optional string Signture;
    9: optional i64 TotalFacorited;
    10: optional i64 WorkCount;
    11: optional i64 FacoriteCount;
    12: optional string Message;
    13: required i32 MsgType;
}

struct RelationFriendListReq {
    1: required i64 UserID;
    2: required string Token;
}
struct RelationFriendListResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: required list<FriendUser> UserList;
}

service RelationService {
    RelationActionResp RelationAction(1: RelationActionReq req) (api.post="/douyin/relation/action/");
    RelationFollowListResp RelationFollowList(1: RelationFollowListReq req) (api.get="/douyin/relation/follow/list/");
    RelationFollowerListResp RelationFollowerList(1: RelationFollowerListReq req) (api.get="/douyin/relation/follower/list/");
    RelationFriendListResp RelationFriendList(1: RelationFriendListReq req) (api.get="/douyin/relation/friend/list/");
}

struct Message {
    1: required i64 ID;
    2: required i64 ToUserID;
    3: required i64 FromUserID;
    4: required string Content;
    5: optional string CreateTime;
}

struct MessageChatReq {
    1: required string Token;
    2: required i64 ToUserID;
    3: required i64 PreMsgTime;
}

struct MessageChatResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
    3: required list<Message> MessageList;
}

struct MessageActionReq {
    1: required string Token;
    2: required i64 ToUserID;
    3: required i32 ActionType;
    4: required string Content;
}

struct MessageActionResp {
    1: required i32 StatusCode;
    2: optional string StatusMsg;
}

service MessageService {
    MessageChatResp MessageChat(1: MessageChatReq req) (api.get="/douyin/message/chat/");
    MessageActionResp MessageAction(1: MessageActionReq req) (api.post="/douyin/message/action/")
}