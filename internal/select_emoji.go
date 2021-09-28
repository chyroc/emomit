package internal

import (
	"strings"

	"github.com/manifoldco/promptui"
)

func Select(title string, items []string) (int, error) {
	prompt := promptui.Select{
		Label:             title,
		Items:             items,
		Size:              15,
		StartInSearchMode: true,
		Searcher: func(input string, index int) bool {
			return InText(strings.ToLower(input), strings.ToLower(items[index]))
		},
	}
	idx, _, err := prompt.Run()
	if err != nil {
		return 0, err
	}

	return idx, nil
}
