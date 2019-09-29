package sys

import (
	"io/ioutil"
	"os"
	"os/exec"
	"os/user"
	"strconv"
	"strings"

	"github.com/kcarretto/paragon/script"
	"github.com/pkg/errors"
)

func setBit(n uint32, pos uint) uint32 {
	n |= (1 << pos)
	return n
}

// Remove uses os.Rename to remove a file/folder. WARNING: basically works like rm -rf.
//
// @param file: A string for the path of the file.
//
// @return (nil, nil) iff success; (nil, err) o/w
func Remove(parser script.ArgParser) (script.Retval, error) {
	file, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	err = os.RemoveAll(file)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// ReadFile uses ioutil.ReadFile to read an entire file's contents.
//
// @param file: A string for the path of the file.
//
// @return (fileContents, nil) iff success; (nil, err) o/w
func ReadFile(parser script.ArgParser) (script.Retval, error) {
	file, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	result, err := ioutil.ReadFile(file)
	return string(result), err
}

// WriteFile uses ioutil.WriteFile to write an entire file's contents, perms are set to 0644.
//
// @param file: A string for the path of the file.
// @param content: A string for the content of the file to be written to.
//
// @return (nil, nil) iff success; (nil, err) o/w
func WriteFile(parser script.ArgParser) (script.Retval, error) {
	file, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	content, err := parser.GetString(1)
	if err != nil {
		return nil, err
	}
	// if you fucks want different perms use chmod
	err = ioutil.WriteFile(file, []byte(content), 0644)
	return nil, err
}

// Move uses os.Rename to move a file from source to destination.
//
// @param srcFile: A string for the path of the source file.
// @param dstFile: A string for the path of the destination file.
//
// @return (nil, nil) iff success; (nil, err) o/w
func Move(parser script.ArgParser) (script.Retval, error) {
	srcFile, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	dstFile, err := parser.GetString(1)
	if err != nil {
		return nil, err
	}
	err = os.Rename(srcFile, dstFile)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// Copy uses ioutil.ReadFile and ioutil.WriteFile to copy a file from source to destination.
//
// @param srcFile: A string for the path of the source file.
// @param dstFile: A string for the path of the destination file.
//
// @return (nil, nil) iff success; (nil, err) o/w
func Copy(parser script.ArgParser) (script.Retval, error) {
	srcFile, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	dstFile, err := parser.GetString(1)
	if err != nil {
		return nil, err
	}

	result, err := ioutil.ReadFile(srcFile)
	if err != nil {
		return nil, err
	}

	err = ioutil.WriteFile(dstFile, result, 0777)

	return nil, err
}

// Exec uses os.Command to execute the passed string
//
// @param cmd: A string for execution.
//
// @return (stdoutStderr, nil) iff success; (nil, err) o/w
func Exec(parser script.ArgParser) (script.Retval, error) {
	cmd, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	argv := strings.Fields(cmd)
	if len(argv) == 0 {
		return nil, errors.New("exec expected args but got none")
	}
	bin := exec.Command(argv[0], argv[1:]...)
	result, err := bin.CombinedOutput()
	return string(result), err
}

// Chown uses os.Chown to change the user/group ownership of a file/dir.
//
// @param file: A string for execution.
// @param owner: A string representing a user and/or group, separated by a ":" character.
//
// @return (nil, nil) iff success; (nil, err) o/w
func Chown(parser script.ArgParser) (script.Retval, error) {
	file, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	owner, err := parser.GetString(1)
	if err != nil {
		return nil, err
	}

	// -1 here means don't change, so it is a good default
	var uid int64 = -1
	var gid int64 = -1
	t := strings.Split(owner, ":")

	if t[0] != "" {
		username := t[0]
		userData, err := user.Lookup(username)
		if err != nil {
			return nil, err
		}
		uid, err = strconv.ParseInt(userData.Uid, 10, 32)
		if err != nil {
			return nil, err
		}
	}
	if len(t) > 1 && t[1] != "" {
		group := t[1]
		groupData, err := user.LookupGroup(group)
		if err != nil {
			return nil, err
		}
		gid, err = strconv.ParseInt(groupData.Gid, 10, 32)
		if err != nil {
			return nil, err
		}
	}
	return nil, os.Chown(file, int(uid), int(gid))
}

// Chmod uses os.Chmod to change a file's permissions. All optional params are assumed to be false unless passed.
//
// @param file: A string for the path of the file.
// @param ?setUser: A bool for the set user bit.
// @param ?setGroup: A bool for the set group bit.
// @param ?setSticky: A bool for the sticky bit.
// @param ?ownerRead: A bool for the owner read permission.
// @param ?ownerWrite: A bool for the owner write permission. In Windows this is the only bit that matters (set file
// to read only iff false; true o/w).
// @param ?ownerExec: A bool for the owner execute permission.
// @param ?groupRead: A bool for the group read permission.
// @param ?groupWrite: A bool for the group write permission.
// @param ?groupExec: A bool for the group execute permission.
// @param ?worldRead: A bool for the world read permission.
// @param ?worldWrite: A bool for the world write permission.
// @param ?worldExec: A bool for the world execute permission.
//
// @return (nil, nil) iff success; (nil, err) o/w
func Chmod(parser script.ArgParser) (script.Retval, error) {
	file, err := parser.GetString(0)
	if err != nil {
		return nil, err
	}
	var perms uint32
	setUser, _ := parser.GetBoolByName("setUser")
	setGroup, _ := parser.GetBoolByName("setGroup")
	setSticky, _ := parser.GetBoolByName("setSticky")
	ownerRead, _ := parser.GetBoolByName("ownerRead")
	ownerWrite, _ := parser.GetBoolByName("ownerWrite")
	ownerExec, _ := parser.GetBoolByName("ownerExec")
	groupRead, _ := parser.GetBoolByName("groupRead")
	groupWrite, _ := parser.GetBoolByName("groupWrite")
	groupExec, _ := parser.GetBoolByName("groupExec")
	worldRead, _ := parser.GetBoolByName("worldRead")
	worldWrite, _ := parser.GetBoolByName("worldWrite")
	worldExec, _ := parser.GetBoolByName("worldExec")

	// world perms
	if worldExec {
		setBit(perms, 0)
	}
	if worldWrite {
		setBit(perms, 1)
	}
	if worldRead {
		setBit(perms, 2)
	}

	// group perms
	if groupExec {
		setBit(perms, 3)
	}
	if groupWrite {
		setBit(perms, 4)
	}
	if groupRead {
		setBit(perms, 5)
	}

	// owner perms
	if ownerExec {
		setBit(perms, 6)
	}
	if ownerWrite {
		setBit(perms, 7)
	}
	if ownerRead {
		setBit(perms, 8)
	}

	// other perms
	if setSticky {
		setBit(perms, 9)
	}
	if setGroup {
		setBit(perms, 10)
	}
	if setUser {
		setBit(perms, 11)
	}

	err = os.Chmod(file, os.FileMode(perms))
	return nil, err
}
