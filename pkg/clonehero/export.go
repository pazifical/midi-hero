package clonehero

import "os"

func WriteToFile(chart Chart, filePath string) error {
	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(chart.String())
	if err != nil {
		return err
	}

	return nil
}
