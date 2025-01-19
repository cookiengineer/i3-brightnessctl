package detectors

import "os"

func GetBacklights() []string {

	result := make([]string, 0)

	prefix := "/sys/class/backlight"

	files, err0 := os.ReadDir(prefix)

	if err0 == nil {

		for _, file := range files {

			name := file.Name()

			stat1, err1 := os.Stat(prefix + "/" + name)
			stat2, err2 := os.Lstat(prefix + "/" + name)

			if err1 == nil && err2 == nil {

				if stat1.IsDir() == true && stat2.IsDir() == false {
					result = append(result, name)
				}

			}

		}

	}

	return result

}
