package metadata

import (
	"github.com/romberli/das/pkg/message"
	"github.com/romberli/go-util/config"
)

func init() {
	initDebugUserMessage()
	initInfoUserMessage()
	initErrorUserMessage()
}

const (
	// debug
	DebugMetadataGetUserAll     = 100901
	DebugMetadataGetUserByID    = 100902
	DebugMetadataAddUser        = 100903
	DebugMetadataUpdateUser     = 100904
	DebugMetadataGetUserByName  = 100905
	DebugMetadataGetEmployeeID  = 100906
	DebugMetadataGetAccountName = 100907
	DebugMetadataGetEmail       = 100908
	DebugMetadataGetTelephone   = 100909
	DebugMetadataGetMobile      = 100910
	DebugMetadataDeleteUserByID = 100911
	// info
	InfoMetadataGetUserAll     = 200901
	InfoMetadataGetUserByID    = 200902
	InfoMetadataAddUser        = 200903
	InfoMetadataUpdateUser     = 200904
	InfoMetadataGetUserByName  = 200905
	InfoMetadataGetEmployeeID  = 200906
	InfoMetadataGetAccountName = 200907
	InfoMetadataGetEmail       = 200908
	InfoMetadataGetTelephone   = 200909
	InfoMetadataGetMobile      = 200910
	InfoMetadataDeleteUserByID = 200911
	// error
	ErrMetadataGetUserAll     = 400901
	ErrMetadataGetUserByID    = 400902
	ErrMetadataAddUser        = 400903
	ErrMetadataUpdateUser     = 400904
	ErrMetadataGetUserByName  = 400905
	ErrMetadataGetEmployeeID  = 400906
	ErrMetadataGetAccountName = 400907
	ErrMetadataGetEmail       = 400908
	ErrMetadataGetTelephone   = 400909
	ErrMetadataGetMobile      = 400910
	ErrMetadataDeleteUserByID = 400911
)

func initDebugUserMessage() {
	message.Messages[DebugMetadataGetUserAll] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetUserAll, "metadata: get all user message: %s")
	message.Messages[DebugMetadataGetUserByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetUserByID, "metadata: get user by id message: %s")
	message.Messages[DebugMetadataAddUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataAddUser, "metadata: add new user message: %s")
	message.Messages[DebugMetadataUpdateUser] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataUpdateUser, "metadata: update user message: %s")
	message.Messages[DebugMetadataGetUserByName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetUserByName, "metadata: get user by username message: %s")
	message.Messages[DebugMetadataGetEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetEmployeeID, "metadata: get user by employeeid message: %s")
	message.Messages[DebugMetadataGetAccountName] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetAccountName, "metadata: get user by accountname message: %s")
	message.Messages[DebugMetadataGetEmail] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetEmail, "metadata: get user by email message: %s")
	message.Messages[DebugMetadataGetTelephone] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetTelephone, "metadata: get user by telephone message: %s")
	message.Messages[DebugMetadataGetMobile] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataGetMobile, "metadata: get user by mobile message: %s")
	message.Messages[DebugMetadataDeleteUserByID] = config.NewErrMessage(message.DefaultMessageHeader, DebugMetadataDeleteUserByID, "metadata: delete user by ID message: %s")
}

func initInfoUserMessage() {
	message.Messages[InfoMetadataGetUserAll] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetUserAll, "metadata: get user all completed")
	message.Messages[InfoMetadataGetUserByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetUserByID, "metadata: get user by id completed. id: %d")
	message.Messages[InfoMetadataAddUser] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataAddUser, "metadata: add new user completed. user_name: %s")
	message.Messages[InfoMetadataUpdateUser] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataUpdateUser, "metadata: update user completed. id: %d")
	message.Messages[InfoMetadataGetUserByName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetUserByName, "metadata: get user by username completed.Name: %s\n%s")
	message.Messages[InfoMetadataGetEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetEmployeeID, "metadata: get user by employeeid completed.employID: %d\n%s")
	message.Messages[InfoMetadataGetAccountName] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetAccountName, "metadata: get user by accountname completed.accountName: %s\n%s")
	message.Messages[InfoMetadataGetEmail] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetEmail, "metadata: get user by email completed.email: %s\n%s")
	message.Messages[InfoMetadataGetTelephone] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetTelephone, "metadata: get user by telephone completed.telephone: %s\n%s")
	message.Messages[InfoMetadataGetMobile] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataGetMobile, "metadata: get user by mobile completed.mobile: %s\n%s")
	message.Messages[InfoMetadataDeleteUserByID] = config.NewErrMessage(message.DefaultMessageHeader, InfoMetadataDeleteUserByID, "metadata: delete user by ID completed. id: %d")
}

func initErrorUserMessage() {
	message.Messages[ErrMetadataGetUserAll] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetUserAll, "metadata: get all user failed.\n%s")
	message.Messages[ErrMetadataGetUserByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetUserByID, "metadata: get user by id failed. id: %d\n%s")
	message.Messages[ErrMetadataAddUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataAddUser, "metadata: add new user failed. user_name: %s\n%s")
	message.Messages[ErrMetadataUpdateUser] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataUpdateUser, "metadata: update user failed. id: %d\n%s")
	message.Messages[ErrMetadataGetUserByName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetUserByName, "metadata: get user by username failed.Name: %s\n%s")
	message.Messages[ErrMetadataGetEmployeeID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetEmployeeID, "metadata: get user by employeeid failed.employID: %d\n%s")
	message.Messages[ErrMetadataGetAccountName] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetAccountName, "metadata: get user by accountname failed.accountName: %s\n%s")
	message.Messages[ErrMetadataGetEmail] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetEmail, "metadata: get user by email failed.email: %s\n%s")
	message.Messages[ErrMetadataGetTelephone] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetTelephone, "metadata: get user by telephone failed.telephone: %s\n%s")
	message.Messages[ErrMetadataGetMobile] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataGetMobile, "metadata: get user by mobile failed.mobile: %s\n%s")
	message.Messages[ErrMetadataDeleteUserByID] = config.NewErrMessage(message.DefaultMessageHeader, ErrMetadataDeleteUserByID, "metadata: delete user by ID failed. id: %d\n%s")
}
