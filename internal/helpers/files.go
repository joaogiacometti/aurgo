package helpers

import (
	"os"
	"os/exec"
)

func runCmd(name string, args []string, dir string) error {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	if dir != "" {
		cmd.Dir = dir
	}

	return cmd.Run()
}

func IsDir(path string) (bool, error) {
	info, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return info.IsDir(), nil
}

func CloneRepo(url, dst string) error {
	return runCmd("git", []string{"clone", "--depth=1", url, dst}, "")
}

func BuildPackage(dir string) error {
	return runCmd("makepkg", []string{"-si", "--noconfirm"}, dir)
}

func RemoveWithPacman(pkgName string) error {
	return runCmd("sudo", []string{"pacman", "-Rns", pkgName}, "")
}
