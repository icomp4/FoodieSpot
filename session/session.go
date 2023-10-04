/*
Session store
Used to manage user sessions, makes it simple to pass current user to functions
I.E. If you need a user's ID for a function, can retrieve it via the session
*/

package session

import (
	"fmt"
	"foodSharer/models"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var Store *session.Store

func init() {
	store := session.New()
	Store = store
}

func SetSession(sess *session.Session, user *models.User) error {
	userIDStr := fmt.Sprintf("%d", user.ID)
	sess.Set("UserID", userIDStr)
	sess.Set("Username", user.Username)
	sess.Set("Authorized", true)
	if err := sess.Save(); err != nil {
		return err
	}
	return nil
}
