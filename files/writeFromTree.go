package files

import (
	"os"

	utils "github.com/ayushsatyam146/gitpot/utils"
)

func WriteToGitpotFromTree(tree *Tree, path string) string {

  // if _, err := os.Stat(path); os.IsNotExist(err) {
  //   os.Mkdir(path, 0755)
  // } 

  TreeFileContent := "tree\n"
  // make a git-tree wala file here sabse pehle 
  for _, child := range tree.children {
    if(!child.isDir) {
      objectDirFileName, _ := utils.GetSHA1(child.value)
      // objectDirContent := "blob\n" + utils.EncodeBase64(child.value)
      objectDirContent := "blob\n" + string(child.value)
      TreeFileContent += "blob " + objectDirFileName + " " + child.name + "\n"
      os.MkdirAll(path + "/objects/" + objectDirFileName[0:2], 0755)
      file, err := os.Create(path + "/objects/" + objectDirFileName[0:2] + "/" + objectDirFileName[2:])
      if err != nil {
        panic(err)
      }
      defer file.Close()
      file.WriteString(objectDirContent)
    } else {
      hash := WriteToGitpotFromTree(child, path)
      TreeFileContent += "tree " + hash + " " + child.name + "\n"
    }
  }

  finalHash, _ := utils.GetSHA1([]byte(TreeFileContent))
  os.MkdirAll(path + "/objects/" + finalHash[0:2], 0755)
  file, err := os.Create(path + "/objects/" + finalHash[0:2] + "/" + finalHash[2:])
  if err != nil {
    panic(err)
  }
  defer file.Close()
  file.WriteString(TreeFileContent)

  return finalHash
}