package facade

import "log"

func GetCommentsByEntity(entityID string) {

	//	print that its received
}

func GetCommentsByUser(userID string) {

}

func AddComment(entityID string, comment string, uid string) {

	//	print that its received

	log.Printf("From the facade : User with UID %s and  added a comment on entity %s: %s\n", uid, entityID, comment)

}
