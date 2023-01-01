package commands

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"

	"github.com/jackmordaunt/icns/v2"

	"github.com/leaanthony/winicon"
)

type IconOptions struct {
	Input string `description:"The input image file"`
	Sizes string `description:"The sizes to generate in .ico file (comma separated)"`

	WindowsFilename string `description:"The output filename for the Windows icon"`
	MacFilename     string `description:"The output filename for the Mac icon bundle"`
}

func (i *IconOptions) Default() *IconOptions {
	return &IconOptions{
		Sizes:           "256,128,64,48,32,16",
		MacFilename:     "icons.icns",
		WindowsFilename: "icons.ico",
	}
}

func Icon(options *IconOptions) error {
	if options.Input == "" {
		return fmt.Errorf("input is required")
	}

	if options.WindowsFilename == "" && options.MacFilename == "" {
		return fmt.Errorf("at least one output filename is required")
	}

	// Parse sizes
	sizes, err := parseSizes(options.Sizes)
	if err != nil {
		return err
	}

	iconData, err := os.ReadFile(options.Input)
	if err != nil {
		return err
	}

	if options.WindowsFilename != "" {
		err := generateWindowsIcon(iconData, sizes, options)
		if err != nil {
			return err
		}
	}

	if options.MacFilename != "" {
		err := generateMacIcon(iconData, options)
		if err != nil {
			return err
		}
	}

	return nil
}

func parseSizes(sizes string) ([]int, error) {
	// split the input string by comma and confirm that each one is an integer
	parsedSizes := strings.Split(sizes, ",")
	result := make([]int, len(parsedSizes))
	for i, size := range parsedSizes {
		s, err := strconv.Atoi(size)
		if err != nil {
			return nil, err
		}
		if s == 0 {
			continue
		}
		result[i] = s
	}

	// put all integers in a slice and return
	return result, nil
}

func generateMacIcon(iconData []byte, options *IconOptions) error {

	srcImg, _, err := image.Decode(bytes.NewBuffer(iconData))
	if err != nil {
		return err
	}

	dest, err := os.Create(options.MacFilename)
	if err != nil {
		return err

	}
	defer func() {
		err = dest.Close()
		if err == nil {
			return
		}
	}()
	return icns.Encode(dest, srcImg)
}

func generateWindowsIcon(iconData []byte, sizes []int, options *IconOptions) error {

	var output bytes.Buffer

	err := winicon.GenerateIcon(bytes.NewBuffer(iconData), &output, sizes)
	if err != nil {
		return err
	}

	err = os.WriteFile(options.WindowsFilename, output.Bytes(), 0644)
	if err != nil {
		return err
	}
	return nil
}