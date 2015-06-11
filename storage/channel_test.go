package storage

import (
	"testing"

	"github.com/khlieng/name_pending/Godeps/_workspace/src/github.com/stretchr/testify/assert"
)

func TestGetSetUsers(t *testing.T) {
	channelStore := NewChannelStore()
	users := []string{"a,b"}
	channelStore.SetUsers(users, "srv", "#chan")
	assert.Equal(t, channelStore.GetUsers("srv", "#chan"), users)
}

func TestAddRemoveUser(t *testing.T) {
	channelStore := NewChannelStore()
	channelStore.AddUser("user", "srv", "#chan")
	channelStore.AddUser("user2", "srv", "#chan")
	assert.Equal(t, channelStore.GetUsers("srv", "#chan"), []string{"user", "user2"})
	channelStore.RemoveUser("user", "srv", "#chan")
	assert.Equal(t, []string{"user2"}, channelStore.GetUsers("srv", "#chan"))
}

func TestRemoveUserAll(t *testing.T) {
	channelStore := NewChannelStore()
	channelStore.AddUser("user", "srv", "#chan1")
	channelStore.AddUser("user", "srv", "#chan2")
	channelStore.RemoveUserAll("user", "srv")
	assert.Empty(t, channelStore.GetUsers("srv", "#chan1"))
	assert.Empty(t, channelStore.GetUsers("srv", "#chan2"))
}

func TestRenameUser(t *testing.T) {
	channelStore := NewChannelStore()
	channelStore.AddUser("user", "srv", "#chan1")
	channelStore.AddUser("user", "srv", "#chan2")
	channelStore.RenameUser("user", "new", "srv")
	assert.Equal(t, []string{"new"}, channelStore.GetUsers("srv", "#chan1"))
	assert.Equal(t, []string{"new"}, channelStore.GetUsers("srv", "#chan2"))
}

func TestMode(t *testing.T) {
	channelStore := NewChannelStore()
	channelStore.AddUser("+user", "srv", "#chan")
	channelStore.SetMode("srv", "#chan", "user", "o", "v")
	assert.Equal(t, []string{"@user"}, channelStore.GetUsers("srv", "#chan"))
	channelStore.SetMode("srv", "#chan", "user", "v", "")
	assert.Equal(t, []string{"+user"}, channelStore.GetUsers("srv", "#chan"))
	channelStore.SetMode("srv", "#chan", "user", "", "v")
	assert.Equal(t, []string{"user"}, channelStore.GetUsers("srv", "#chan"))
}

func TestTopic(t *testing.T) {
	channelStore := NewChannelStore()
	assert.Equal(t, "", channelStore.GetTopic("srv", "#chan"))
	channelStore.SetTopic("the topic", "srv", "#chan")
	assert.Equal(t, "the topic", channelStore.GetTopic("srv", "#chan"))
}
