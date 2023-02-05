// Code generated by hertz generator. DO NOT EDIT.

package Base

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	base "github.com/linzijie1998/bytedance_camp_douyin/biz/handler/douyin/base"
)

/*
 This file will register all the routes of the services in the master idl.
 And it will update automatically when you use the "update" command for the idl.
 So don't modify the contents of the file, or your code will be deleted when it is updated.
*/

// Register register routes based on the IDL 'api.${HTTP Method}' annotation.
func Register(r *server.Hertz) {

	root := r.Group("/", rootMw()...)
	{
		_douyin := root.Group("/douyin", _douyinMw()...)
		{
			_feed := _douyin.Group("/feed", _feedMw()...)
			_feed.GET("/", append(_feed0Mw(), base.Feed)...)
		}
		{
			_publish := _douyin.Group("/publish", _publishMw()...)
			{
				_action := _publish.Group("/action", _actionMw()...)
				_action.POST("/", append(_publish_ctionMw(), base.PublishAction)...)
			}
			{
				_list := _publish.Group("/list", _listMw()...)
				_list.GET("/", append(_publishlistMw(), base.PublishList)...)
			}
		}
		{
			_user := _douyin.Group("/user", _userMw()...)
			_user.GET("/", append(_userinfoMw(), base.UserInfo)...)
			{
				_login := _user.Group("/login", _loginMw()...)
				_login.POST("/", append(_userloginMw(), base.UserLogin)...)
			}
			{
				_register := _user.Group("/register", _registerMw()...)
				_register.POST("/", append(_userregisterMw(), base.UserRegister)...)
			}
		}
	}
}
