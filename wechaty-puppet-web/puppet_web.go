package puppetweb

import (
	"errors"
	"fmt"
	wechatyPuppet "github.com/wechaty/go-wechaty/wechaty-puppet"
	file_box "github.com/wechaty/go-wechaty/wechaty-puppet/file-box"
	"github.com/wechaty/go-wechaty/wechaty-puppet/schemas"
	"log"
)

// PuppetWeb struct
type PuppetWeb struct {
	*wechatyPuppet.Puppet
}

// NewPuppetWeb new PuppetWeb struct
func NewPuppetWeb(o wechatyPuppet.Option) (*PuppetWeb, error) {

	return nil, nil
}

// MessageImage ...
func (p *PuppetWeb) MessageImage(messageID string, imageType schemas.ImageType) (*file_box.FileBox, error) {

	return nil, nil
}

// Start ...
func (p *PuppetWeb) Start() (err error) {
	log.Println("PuppetHostie Start()")
	defer func() {
		if err != nil {
			err = fmt.Errorf("PuppetHostie Start() rejection: %w", err)
		}
	}()

	return nil
}

// Stop ...
func (p *PuppetWeb) Stop() {
	var err error
	defer func() {
		if err != nil {
			log.Printf("PuppetHostie Stop err: %s\n", err)
		}
	}()

}

// Logout ...
func (p *PuppetWeb) Logout() error {
	log.Println("PuppetHostie Logout()")
	//if !p.logonoff() {
	//	return errors.New("logout before login? ")
	//}

	return nil
}

// Ding ...
func (p *PuppetWeb) Ding(data string) {
	log.Printf("PuppetHostie Ding(%s)\n", data)

	//if err != nil {
	//	log.Printf("PuppetHostie Ding() err: %s\n", err)
	//}
}

// SetContactAlias ...
func (p *PuppetWeb) SetContactAlias(contactID string, alias string) error {
	log.Printf("PuppetHostie, SetContactAlias(%s, %s)\n", contactID, alias)

	return nil
}

// ContactAlias ...
func (p *PuppetWeb) ContactAlias(contactID string) (string, error) {
	log.Printf("PuppetHostie, 'ContactAlias(%s)\n", contactID)

	return "", nil
}

// ContactList ...
func (p *PuppetWeb) ContactList() ([]string, error) {
	log.Println("PuppetHostie ContactList()")
	return nil, nil
}

// ContactQRCode ...
func (p *PuppetWeb) ContactQRCode(contactID string) (string, error) {
	log.Printf("PuppetHostie ContactQRCode(%s)\n", contactID)
	if contactID != p.SelfID() {
		return "", errors.New("can not set avatar for others")
	}
	return "", nil
}

// SetContactAvatar ...
func (p *PuppetWeb) SetContactAvatar(contactID string, fileBox *file_box.FileBox) error {
	log.Printf("PuppetHostie SetContactAvatar(%s)\n", contactID)
	return nil
}

// ContactAvatar ...
func (p *PuppetWeb) ContactAvatar(contactID string) (*file_box.FileBox, error) {
	log.Printf("PuppetHostie ContactAvatar(%s)\n", contactID)
	return nil, nil
}

// ContactRawPayload ...
func (p *PuppetWeb) ContactRawPayload(contactID string) (*schemas.ContactPayload, error) {
	log.Printf("PuppetHostie ContactRawPayload(%s)\n", contactID)
	return nil, nil
}

// SetContactSelfName ...
func (p *PuppetWeb) SetContactSelfName(name string) error {
	log.Printf("PuppetHostie SetContactSelfName(%s)\n", name)
	return nil
}

// ContactSelfQRCode ...
func (p *PuppetWeb) ContactSelfQRCode() (string, error) {
	log.Println("PuppetHostie ContactSelfQRCode()")
	return "", nil
}

// SetContactSelfSignature ...
func (p *PuppetWeb) SetContactSelfSignature(signature string) error {
	log.Printf("PuppetHostie SetContactSelfSignature(%s)\n", signature)
	return nil
}

// MessageMiniProgram ...
func (p *PuppetWeb) MessageMiniProgram(messageID string) (*schemas.MiniProgramPayload, error) {
	log.Printf("PuppetHostie MessageMiniProgram(%s)\n", messageID)
	return nil, nil
}

// MessageContact ...
func (p *PuppetWeb) MessageContact(messageID string) (string, error) {
	log.Printf("PuppetHostie MessageContact(%s)\n", messageID)
	return "", nil
}

// MessageSendMiniProgram ...
func (p *PuppetWeb) MessageSendMiniProgram(conversationID string, miniProgramPayload *schemas.MiniProgramPayload) (string, error) {
	return "", nil
}

// MessageRecall ...
func (p *PuppetWeb) MessageRecall(messageID string) (bool, error) {
	log.Printf("PuppetHostie MessageRecall(%s)\n", messageID)
	return false, nil
}

// MessageFile ...
func (p *PuppetWeb) MessageFile(id string) (*file_box.FileBox, error) {
	log.Printf("PuppetHostie MessageFile(%s)\n", id)
	return nil, nil
}

// MessageRawPayload ...
func (p *PuppetWeb) MessageRawPayload(id string) (*schemas.MessagePayload, error) {
	log.Printf("PuppetHostie MessagePayload(%s)\n", id)
	return nil, nil
}

// MessageSendText ...
func (p *PuppetWeb) MessageSendText(conversationID string, text string, mentionIDList ...string) (string, error) {
	log.Printf("PuppetHostie messageSendText(%s, %s)\n", conversationID, text)
	return "", nil
}

// MessageSendFile ...
func (p *PuppetWeb) MessageSendFile(conversationID string, fileBox *file_box.FileBox) (string, error) {
	log.Printf("PuppetHostie MessageSendFile(%s)\n", conversationID)
	return "", nil
}

// MessageSendContact ...
func (p *PuppetWeb) MessageSendContact(conversationID string, contactID string) (string, error) {
	log.Printf("PuppetHostie MessageSendContact(%s, %s)\n", conversationID, contactID)
	return "", nil
}

// MessageSendURL ...
func (p *PuppetWeb) MessageSendURL(conversationID string, urlLinkPayload *schemas.UrlLinkPayload) (string, error) {
	log.Printf("PuppetHostie MessageSendURL(%s, %s)\n", conversationID, urlLinkPayload)
	return "", nil
}

// MessageURL ...
func (p *PuppetWeb) MessageURL(messageID string) (*schemas.UrlLinkPayload, error) {
	log.Printf("PuppetHostie MessageURL(%s)\n", messageID)
	return nil, nil
}

// RoomRawPayload ...
func (p *PuppetWeb) RoomRawPayload(id string) (*schemas.RoomPayload, error) {
	log.Printf("PuppetHostie RoomRawPayload(%s)\n", id)
	return nil, nil
}

// RoomList ...
func (p *PuppetWeb) RoomList() ([]string, error) {
	log.Printf("PuppetHostie RoomList()\n")
	return nil, nil
}

// RoomDel ...
func (p *PuppetWeb) RoomDel(roomID, contactID string) error {
	log.Printf("PuppetHostie roomDel(%s, %s)\n", roomID, contactID)
	return nil
}

// RoomAvatar ...
func (p *PuppetWeb) RoomAvatar(roomID string) (*file_box.FileBox, error) {
	log.Printf("PuppetHostie RoomAvatar(%s)\n", roomID)
	return nil, nil
}

// RoomAdd ...
func (p *PuppetWeb) RoomAdd(roomID, contactID string) error {
	log.Printf("PuppetHostie RoomAdd(%s, %s)\n", roomID, contactID)
	return nil
}

// SetRoomTopic ...
func (p *PuppetWeb) SetRoomTopic(roomID string, topic string) error {
	log.Printf("PuppetHostie setRoomTopic(%s, %s)\n", roomID, topic)
	return nil
}

// RoomTopic ...
func (p *PuppetWeb) RoomTopic(roomID string) (string, error) {
	log.Printf("PuppetHostie RoomTopic(%s)\n", roomID)
	return "", nil
}

// RoomCreate ...
func (p *PuppetWeb) RoomCreate(contactIDList []string, topic string) (string, error) {
	log.Printf("PuppetHostie roomCreate(%s, %s)\n", contactIDList, topic)
	return "", nil
}

// RoomQuit ...
func (p *PuppetWeb) RoomQuit(roomID string) error {
	log.Printf("PuppetHostie RoomQuit(%s)\n", roomID)
	return nil
}

// RoomQRCode ...
func (p *PuppetWeb) RoomQRCode(roomID string) (string, error) {
	log.Printf("PuppetHostie RoomQRCode(%s)\n", roomID)
	return "", nil
}

// RoomMemberList ...
func (p *PuppetWeb) RoomMemberList(roomID string) ([]string, error) {
	log.Printf("PuppetHostie RoomMemberList(%s)\n", roomID)
	return nil, nil
}

// RoomMemberRawPayload ...
func (p *PuppetWeb) RoomMemberRawPayload(roomID string, contactID string) (*schemas.RoomMemberPayload, error) {
	log.Printf("PuppetHostie RoomMemberRawPayload(%s, %s)\n", roomID, contactID)
	return nil, nil
}

// SetRoomAnnounce ...
func (p *PuppetWeb) SetRoomAnnounce(roomID, text string) error {
	log.Printf("PuppetHostie SetRoomAnnounce(%s, %s)\n", roomID, text)
	return nil
}

// RoomAnnounce ...
func (p *PuppetWeb) RoomAnnounce(roomID string) (string, error) {
	log.Printf("PuppetHostie RoomAnnounce(%s)\n", roomID)
	return "", nil
}

// RoomInvitationAccept ...
func (p *PuppetWeb) RoomInvitationAccept(roomInvitationID string) error {
	log.Printf("PuppetHostie RoomInvitationAccept(%s)\n", roomInvitationID)
	return nil
}

// RoomInvitationRawPayload ...
func (p *PuppetWeb) RoomInvitationRawPayload(id string) (*schemas.RoomInvitationPayload, error) {
	log.Printf("PuppetHostie RoomInvitationRawPayload(%s)\n", id)
	return nil, nil
}

// FriendshipSearchPhone ...
func (p *PuppetWeb) FriendshipSearchPhone(phone string) (string, error) {
	log.Printf("PuppetHostie FriendshipSearchPhone(%s)\n", phone)
	return "", nil
}

// FriendshipSearchWeixin ...
func (p *PuppetWeb) FriendshipSearchWeixin(weixin string) (string, error) {
	log.Printf("PuppetHostie FriendshipSearchWeixin(%s)\n", weixin)
	return "", nil
}

// FriendshipRawPayload ...
func (p *PuppetWeb) FriendshipRawPayload(id string) (*schemas.FriendshipPayload, error) {
	log.Printf("PuppetHostie FriendshipRawPayload(%s)\n", id)
	return nil, nil
}

// FriendshipAdd ...
func (p *PuppetWeb) FriendshipAdd(contactID, hello string) (err error) {
	log.Printf("PuppetHostie FriendshipAdd(%s, %s)\n", contactID, hello)
	return err
}

// FriendshipAccept ...
func (p *PuppetWeb) FriendshipAccept(friendshipID string) (err error) {
	log.Printf("PuppetHostie FriendshipAccept(%s)\n", friendshipID)
	return err
}

// TagContactAdd ...
func (p *PuppetWeb) TagContactAdd(id, contactID string) (err error) {
	log.Printf("PuppetHostie TagContactAdd(%s, %s)\n", id, contactID)
	return err
}

// TagContactRemove ...
func (p *PuppetWeb) TagContactRemove(id, contactID string) (err error) {
	log.Printf("PuppetHostie TagContactRemove(%s, %s)\n", id, contactID)
	return err
}

// TagContactDelete ...
func (p *PuppetWeb) TagContactDelete(id string) (err error) {
	log.Printf("PuppetHostie TagContactDelete(%s)\n", id)
	return err
}

// TagContactList ...
func (p *PuppetWeb) TagContactList(contactID string) ([]string, error) {
	log.Printf("PuppetHostie TagContactList(%s)\n", contactID)
	return nil, nil
}
