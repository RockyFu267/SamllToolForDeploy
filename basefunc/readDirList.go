package basefunc

import (
	"SamllToolForDeploy/basecmd"
	"log"
)

func ReadDirList(targetpath string) ([]string, error) {
	resTmp, err := basecmd.CmdAndChangeDirToResAllInOne(targetpath, "ls -l|grep  '^d' | awk '(NR > 1){print $9}'")
	if err != nil {
		log.Println(err)
		return resTmp, err
	}

	return resTmp, nil
}
