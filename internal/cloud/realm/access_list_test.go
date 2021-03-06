package realm_test

import (
	"testing"

	"github.com/10gen/realm-cli/internal/cloud/realm"
	u "github.com/10gen/realm-cli/internal/utils/test"
	"github.com/10gen/realm-cli/internal/utils/test/assert"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// TODO(REALMC-9207): Unskip tests once backend is fully implemented
func TestRealmIPAccess(t *testing.T) {
	u.SkipUnlessRealmServerRunning(t)

	t.Run("should fail without an auth client", func(t *testing.T) {
		client := realm.NewClient(u.RealmServerURL())

		_, err := client.AllowedIPCreate(primitive.NewObjectID().Hex(), primitive.NewObjectID().Hex(), "0.0.0.0/0", "comment", false)
		assert.Equal(t, realm.ErrInvalidSession{}, err)
	})

	t.Run("with an active session", func(t *testing.T) {
		t.Skip("skipping test")
		client := newAuthClient(t)
		groupID := u.CloudGroupID()

		testApp, teardown := setupTestApp(t, client, groupID, "accesslist-test")
		defer teardown()

		t.Run("should create an allowed IP", func(t *testing.T) {
			address := "0.0.0.0"
			comment := "comment"
			useCurrent := false
			allowedIP, err := client.AllowedIPCreate(groupID, testApp.ID, address, comment, useCurrent)

			assert.Nil(t, err)
			assert.Equal(t, address, allowedIP.Address)
			assert.Equal(t, comment, allowedIP.Comment)
		})

	})
}
