package puppetweb

import (
	"errors"
	"fmt"
	wechatyPuppet "github.com/wechaty/go-wechaty/wechaty-puppet"
	file_box "github.com/wechaty/go-wechaty/wechaty-puppet/file-box"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"log"
	"net/http"
	"net/http/cookiejar"
	"time"
)

type Option struct {
	AppId     string
	UserAgent string
	Lang      string
	DeviceID  string
}

// PuppetWeb struct
type PuppetWeb struct {
	*wechatyPuppet.Puppet
	client      *http.Client
	option      *Option
	redirectUri string
	uuid        string
	ticket      xmlTicket
}

// NewPuppetWeb new PuppetWeb struct
func NewPuppetWeb(o *Option) (*PuppetWeb, error) {
	if o == nil {
		o = &Option{}
	}
	if len(o.AppId) < 1 {
		o.AppId = "wx782c26e4c19acffb"
	}
	if len(o.UserAgent) < 1 {
		o.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36"
	}
	if len(o.Lang) < 1 {
		o.Lang = "zh_CN"
	}
	puppetAbstract, err := wechatyPuppet.NewPuppet(wechatyPuppet.Option{})
	if err != nil {
		return nil, err
	}
	puppetWeb := &PuppetWeb{
		Puppet: puppetAbstract,
		option: o,
		client: &http.Client{
			Timeout: time.Second * 100,
			Transport: &http.Transport{
				Proxy: http.ProxyFromEnvironment,
			},
		},
	}
	if len(puppetWeb.option.DeviceID) < 1 {
		puppetWeb.option.DeviceID = puppetWeb.makeDeviceID()
	}
	puppetWeb.client.Jar, err = cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	puppetAbstract.SetPuppetImplementation(puppetWeb)
	return puppetWeb, nil
}

// MessageImage ...
func (p *PuppetWeb) MessageImage(messageID string, imageType schemas.ImageType) (*file_box.FileBox, error) {

	return nil, nil
}

// Start ...
func (p *PuppetWeb) Start() (err error) {
	log.Println("PuppetWeb Start()")
	defer func() {
		if err != nil {
			err = fmt.Errorf("PuppetWeb Start() rejection: %w", err)
		}
	}()

	return nil
}

// Stop ...
func (p *PuppetWeb) Stop() {
	var err error
	defer func() {
		if err != nil {
			log.Printf("PuppetWeb Stop err: %s\n", err)
		}
	}()

}

// Logout ...
func (p *PuppetWeb) Logout() error {
	log.Println("PuppetWeb Logout()")
	//if !p.logonoff() {
	//	return errors.New("logout before login? ")
	//}

	return nil
}

// Ding ...
func (p *PuppetWeb) Ding(data string) {
	log.Printf("PuppetWeb Ding(%s)\n", data)

	//if err != nil {
	//	log.Printf("PuppetWeb Ding() err: %s\n", err)
	//}
}

// SetContactAlias ...
func (p *PuppetWeb) SetContactAlias(contactID string, alias string) error {
	log.Printf("PuppetWeb, SetContactAlias(%s, %s)\n", contactID, alias)

	return nil
}

// ContactAlias ...
func (p *PuppetWeb) ContactAlias(contactID string) (string, error) {
	log.Printf("PuppetWeb, 'ContactAlias(%s)\n", contactID)

	return "", nil
}

// ContactList ...
func (p *PuppetWeb) ContactList() ([]string, error) {
	log.Println("PuppetWeb ContactList()")
	return nil, nil
}

// ContactQRCode ...
func (p *PuppetWeb) ContactQRCode(contactID string) (string, error) {
	log.Printf("PuppetWeb ContactQRCode(%s)\n", contactID)
	if contactID != p.SelfID() {
		return "", errors.New("can not set avatar for others")
	}
	return "", nil
}

// SetContactAvatar ...
func (p *PuppetWeb) SetContactAvatar(contactID string, fileBox *file_box.FileBox) error {
	log.Printf("PuppetWeb SetContactAvatar(%s)\n", contactID)
	return nil
}

// ContactAvatar ...
func (p *PuppetWeb) ContactAvatar(contactID string) (*file_box.FileBox, error) {
	log.Printf("PuppetWeb ContactAvatar(%s)\n", contactID)
	return nil, nil
}

// ContactRawPayload ...
func (p *PuppetWeb) ContactRawPayload(contactID string) (*schemas.ContactPayload, error) {
	log.Printf("PuppetWeb ContactRawPayload(%s)\n", contactID)
	return nil, nil
}

// SetContactSelfName ...
func (p *PuppetWeb) SetContactSelfName(name string) error {
	log.Printf("PuppetWeb SetContactSelfName(%s)\n", name)
	return nil
}

// ContactSelfQRCode ...
func (p *PuppetWeb) ContactSelfQRCode() (string, error) {
	log.Println("PuppetWeb ContactSelfQRCode()")
	return "", nil
}

// SetContactSelfSignature ...
func (p *PuppetWeb) SetContactSelfSignature(signature string) error {
	log.Printf("PuppetWeb SetContactSelfSignature(%s)\n", signature)
	return nil
}

// MessageMiniProgram ...
func (p *PuppetWeb) MessageMiniProgram(messageID string) (*schemas.MiniProgramPayload, error) {
	log.Printf("PuppetWeb MessageMiniProgram(%s)\n", messageID)
	return nil, nil
}

// MessageContact ...
func (p *PuppetWeb) MessageContact(messageID string) (string, error) {
	log.Printf("PuppetWeb MessageContact(%s)\n", messageID)
	return "", nil
}

// MessageSendMiniProgram ...
func (p *PuppetWeb) MessageSendMiniProgram(conversationID string, miniProgramPayload *schemas.MiniProgramPayload) (string, error) {
	return "", nil
}

// MessageRecall ...
func (p *PuppetWeb) MessageRecall(messageID string) (bool, error) {
	log.Printf("PuppetWeb MessageRecall(%s)\n", messageID)
	return false, nil
}

// MessageFile ...
func (p *PuppetWeb) MessageFile(id string) (*file_box.FileBox, error) {
	log.Printf("PuppetWeb MessageFile(%s)\n", id)
	return nil, nil
}

// MessageRawPayload ...
func (p *PuppetWeb) MessageRawPayload(id string) (*schemas.MessagePayload, error) {
	log.Printf("PuppetWeb MessagePayload(%s)\n", id)
	return nil, nil
}

// MessageSendText ...
func (p *PuppetWeb) MessageSendText(conversationID string, text string, mentionIDList ...string) (string, error) {
	log.Printf("PuppetWeb messageSendText(%s, %s)\n", conversationID, text)
	return "", nil
}

// MessageSendFile ...
func (p *PuppetWeb) MessageSendFile(conversationID string, fileBox *file_box.FileBox) (string, error) {
	log.Printf("PuppetWeb MessageSendFile(%s)\n", conversationID)
	return "", nil
}

// MessageSendContact ...
func (p *PuppetWeb) MessageSendContact(conversationID string, contactID string) (string, error) {
	log.Printf("PuppetWeb MessageSendContact(%s, %s)\n", conversationID, contactID)
	return "", nil
}

// MessageSendURL ...
func (p *PuppetWeb) MessageSendURL(conversationID string, urlLinkPayload *schemas.UrlLinkPayload) (string, error) {
	log.Printf("PuppetWeb MessageSendURL(%s, %s)\n", conversationID, urlLinkPayload)
	return "", nil
}

// MessageURL ...
func (p *PuppetWeb) MessageURL(messageID string) (*schemas.UrlLinkPayload, error) {
	log.Printf("PuppetWeb MessageURL(%s)\n", messageID)
	return nil, nil
}

// RoomRawPayload ...
func (p *PuppetWeb) RoomRawPayload(id string) (*schemas.RoomPayload, error) {
	log.Printf("PuppetWeb RoomRawPayload(%s)\n", id)
	return nil, nil
}

// RoomList ...
func (p *PuppetWeb) RoomList() ([]string, error) {
	log.Printf("PuppetWeb RoomList()\n")
	return nil, nil
}

// RoomDel ...
func (p *PuppetWeb) RoomDel(roomID, contactID string) error {
	log.Printf("PuppetWeb roomDel(%s, %s)\n", roomID, contactID)
	return nil
}

// RoomAvatar ...
func (p *PuppetWeb) RoomAvatar(roomID string) (*file_box.FileBox, error) {
	log.Printf("PuppetWeb RoomAvatar(%s)\n", roomID)
	return nil, nil
}

// RoomAdd ...
func (p *PuppetWeb) RoomAdd(roomID, contactID string) error {
	log.Printf("PuppetWeb RoomAdd(%s, %s)\n", roomID, contactID)
	return nil
}

// SetRoomTopic ...
func (p *PuppetWeb) SetRoomTopic(roomID string, topic string) error {
	log.Printf("PuppetWeb setRoomTopic(%s, %s)\n", roomID, topic)
	return nil
}

// RoomTopic ...
func (p *PuppetWeb) RoomTopic(roomID string) (string, error) {
	log.Printf("PuppetWeb RoomTopic(%s)\n", roomID)
	return "", nil
}

// RoomCreate ...
func (p *PuppetWeb) RoomCreate(contactIDList []string, topic string) (string, error) {
	log.Printf("PuppetWeb roomCreate(%s, %s)\n", contactIDList, topic)
	return "", nil
}

// RoomQuit ...
func (p *PuppetWeb) RoomQuit(roomID string) error {
	log.Printf("PuppetWeb RoomQuit(%s)\n", roomID)
	return nil
}

// RoomQRCode ...
func (p *PuppetWeb) RoomQRCode(roomID string) (string, error) {
	log.Printf("PuppetWeb RoomQRCode(%s)\n", roomID)
	return "", nil
}

// RoomMemberList ...
func (p *PuppetWeb) RoomMemberList(roomID string) ([]string, error) {
	log.Printf("PuppetWeb RoomMemberList(%s)\n", roomID)
	return nil, nil
}

// RoomMemberRawPayload ...
func (p *PuppetWeb) RoomMemberRawPayload(roomID string, contactID string) (*schemas.RoomMemberPayload, error) {
	log.Printf("PuppetWeb RoomMemberRawPayload(%s, %s)\n", roomID, contactID)
	return nil, nil
}

// SetRoomAnnounce ...
func (p *PuppetWeb) SetRoomAnnounce(roomID, text string) error {
	log.Printf("PuppetWeb SetRoomAnnounce(%s, %s)\n", roomID, text)
	return nil
}

// RoomAnnounce ...
func (p *PuppetWeb) RoomAnnounce(roomID string) (string, error) {
	log.Printf("PuppetWeb RoomAnnounce(%s)\n", roomID)
	return "", nil
}

// RoomInvitationAccept ...
func (p *PuppetWeb) RoomInvitationAccept(roomInvitationID string) error {
	log.Printf("PuppetWeb RoomInvitationAccept(%s)\n", roomInvitationID)
	return nil
}

// RoomInvitationRawPayload ...
func (p *PuppetWeb) RoomInvitationRawPayload(id string) (*schemas.RoomInvitationPayload, error) {
	log.Printf("PuppetWeb RoomInvitationRawPayload(%s)\n", id)
	return nil, nil
}

// FriendshipSearchPhone ...
func (p *PuppetWeb) FriendshipSearchPhone(phone string) (string, error) {
	log.Printf("PuppetWeb FriendshipSearchPhone(%s)\n", phone)
	return "", nil
}

// FriendshipSearchWeixin ...
func (p *PuppetWeb) FriendshipSearchWeixin(weixin string) (string, error) {
	log.Printf("PuppetWeb FriendshipSearchWeixin(%s)\n", weixin)
	return "", nil
}

// FriendshipRawPayload ...
func (p *PuppetWeb) FriendshipRawPayload(id string) (*schemas.FriendshipPayload, error) {
	log.Printf("PuppetWeb FriendshipRawPayload(%s)\n", id)
	return nil, nil
}

// FriendshipAdd ...
func (p *PuppetWeb) FriendshipAdd(contactID, hello string) (err error) {
	log.Printf("PuppetWeb FriendshipAdd(%s, %s)\n", contactID, hello)
	return err
}

// FriendshipAccept ...
func (p *PuppetWeb) FriendshipAccept(friendshipID string) (err error) {
	log.Printf("PuppetWeb FriendshipAccept(%s)\n", friendshipID)
	return err
}

// TagContactAdd ...
func (p *PuppetWeb) TagContactAdd(id, contactID string) (err error) {
	log.Printf("PuppetWeb TagContactAdd(%s, %s)\n", id, contactID)
	return err
}

// TagContactRemove ...
func (p *PuppetWeb) TagContactRemove(id, contactID string) (err error) {
	log.Printf("PuppetWeb TagContactRemove(%s, %s)\n", id, contactID)
	return err
}

// TagContactDelete ...
func (p *PuppetWeb) TagContactDelete(id string) (err error) {
	log.Printf("PuppetWeb TagContactDelete(%s)\n", id)
	return err
}

// TagContactList ...
func (p *PuppetWeb) TagContactList(contactID string) ([]string, error) {
	log.Printf("PuppetWeb TagContactList(%s)\n", contactID)
	return nil, nil
}

// MessageRawMiniProgramPayload ...
func (p *PuppetWeb) MessageRawMiniProgramPayload(messageID string) (*schemas.MiniProgramPayload, error) {
	panic("implement me")
}
