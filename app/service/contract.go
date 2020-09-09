package service

import (
	"errors"
	"go-im/app/model"
	"time"

	"github.com/jinzhu/gorm"
)

type ContactService struct{}

//添加好友
func (c *ContactService) AddFriend(userid string, dstid string) error {
	if dstid == userid {
		return errors.New("不能添加自己为好友啊")
	}
	//判断是否已经添加了好友
	friend, err := model.CheckFriendRelationShip(userid, dstid)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	//如果好友已经存在，则不添加
	if (*friend).Id != "" {
		return errors.New("该好友已经添加过了")
	}

	//新建好友关系
	err = model.CreatePersonalFriendRelationShip(userid, dstid)

	return err
}

//搜索社群
func (service *ContactService) SearchComunity(userId string) (coms []*model.Community, err error) {
	conatcts := make([]*model.Contact, 0)
	comIds := make([]string, 0)
	coms = make([]*model.Community, 0)

	contact := model.Contact{}
	contact.Ownerid = userId

	conatcts, err = contact.SearchCommunitys()
	if err != nil {
		return
	}

	for _, v := range conatcts {
		comIds = append(comIds, (*v).Dstobj)
	}

	if len(comIds) == 0 {
		return
	}

	coms, err = model.FindInCommunitys(comIds)

	return
}

//根据名字搜索社群
func (c *ContactService) SearchCommunityByName(cname string) (com model.Community, err error) {
	com.Name = cname
	err = com.GetByName()
	return com, err
}

//搜索好友
func (c *ContactService) SearchFriend(userId string) (users []*model.User, err error) {
	objIds := make([]string, 0)
	contact := model.Contact{}
	contact.Ownerid = userId
	friends := make([]*model.Contact, 0)
	friends, err = contact.SearchFriends()
	if err != nil {
		return
	}
	for _, v := range friends {
		objIds = append(objIds, (*v).Dstobj)
	}

	users = make([]*model.User, 0)
	if len(objIds) == 0 {
		return
	}

	//获取好友列表信息
	users, err = model.FindInUsers(objIds)

	return
}

//根据手机号搜索用户
func (c *ContactService) SearchFriendByName(mobile string) (user model.User, err error) {
	user.Mobile = mobile
	err = user.GetByName()

	return
}

func (service *ContactService) SearchComunityIds(userId string) (comIds []string, err error) {
	// 获取用户全部群ID
	conconts := make([]*model.Contact, 0)
	comIds = make([]string, 0)

	contact := model.Contact{}
	contact.Ownerid = userId
	conconts, err = contact.SearchCommunitys()
	if err != nil {
		return
	}

	for _, v := range conconts {
		comIds = append(comIds, (*v).Dstobj)
	}
	return
}

//添加群
func (c *ContactService) CreateCommunity(comm model.Community) (ret model.Community, err error) {
	if len(comm.Name) == 0 {
		err = errors.New("缺少群名称")
		return ret, err
	}
	if comm.Ownerid == "" {
		err = errors.New("请先登录")
		return ret, err
	}
	com := model.Community{
		Ownerid: comm.Ownerid,
	}
	num, err := com.CountCommunitys()

	if num > 5 {
		err = errors.New("一个用户最多只能创见5个群")
		return com, err
	}

	comm.Createat = time.Now()
	err = comm.UserCreateCommunity()
	return com, err
}

//用户加群
func (c *ContactService) JoinCommunity(userId, comId string) error {
	cot := model.Contact{
		Ownerid: userId,
		Dstobj:  comId,
		Cate:    model.ConcatCateComunity,
	}

	//检查是否已经存在过加群记录
	com, err := model.CheckCommunityRelationShip(userId, comId)
	if err != nil && err != gorm.ErrRecordNotFound {
		return err
	}

	//如果群聊已经存在，则不继续添加
	if (*com).Id != "" {
		return errors.New("该群聊已经添加过了")
	}

	//新建群关系
	err = cot.Create()

	return err
}
