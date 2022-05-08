// -----------------------------------------------------------------------------
// go-backup/consts/action.go

package consts

// Action specifies archiving actions, for example:
// CreateArchive, ExtractArchive, ListArchive, VerifyArchive
type Action int

// CreateArchive is an action to create a new archive
// file from the files in the specified source folder.
const CreateArchive = Action(1)

// ExtractArchive is an action to read an existing archive
// file and extract its contents into the specified folder.
const ExtractArchive = Action(2)

// ListArchive is an action to list all the files stored
// in an archive file, without extracting them.
const ListArchive = Action(3)

// VerifyArchive is an action to checks the integrity of all
// files stored in an archive file (by checking file hashes).
const VerifyArchive = Action(4)

// end
