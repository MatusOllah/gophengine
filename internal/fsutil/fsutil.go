package fsutil

import "io/fs"

// NumFiles counts all the files in the filesystem and returns the number of files.
func NumFiles(fsys fs.FS) (int, error) {
	var num int
	err := fs.WalkDir(fsys, ".", func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() {
			return nil
		}

		num++
		return nil
	})
	if err != nil {
		return 0, err
	}

	return num, nil
}
