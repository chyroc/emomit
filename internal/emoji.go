package internal

import (
	"fmt"
)

func SelectEmoji() (EmojiConf, error) {
	idx, err := Select("Select Type", emojiSelectStrings)
	if err != nil {
		return EmojiConf{}, err
	}
	return Emojis[idx], err
}

type EmojiConf struct {
	Emoji      string
	Text       string
	Desc       string
	CommitType CommitType
}
type CommitType string

const (
	CommitTypeFeat     CommitType = "feat"     // new feature for the user, not a new feature for build script
	CommitTypeRevert   CommitType = "revert"   // revert some feat or test or config
	CommitTypeFix      CommitType = "fix"      // bug fix for the user, not a fix to a build script
	CommitTypeHotfix   CommitType = "hotfix"   // buf hotfix, may skip code review or test
	CommitTypeDocs     CommitType = "docs"     // changes to the documentation
	CommitTypeStyle    CommitType = "style"    // formatting, missing semi colons, etc; no production code change
	CommitTypeRefactor CommitType = "refactor" // refactoring production code, eg. renaming a variable
	CommitTypeTest     CommitType = "test"     // adding missing tests, refactoring tests; no production code change
	CommitTypeChore    CommitType = "chore"    // updating grunt tasks etc; no production code change
	CommitTypeRelease  CommitType = "release"  // release new tag
	CommitTypeDeploy   CommitType = "deploy"   // deploy code to test or production env
)

// Emojis from: https://gitmoji.dev/
// CommitType from: https://sparkbox.com/foundry/semantic_commit_messages
// https://www.conventionalcommits.org/en/v1.0.0/
var Emojis = []EmojiConf{
	{"🎨", ":art:", "Improve structure / format of the code.", CommitTypeStyle},
	{"⚡️", ":zap:", "Improve performance.", CommitTypeFeat},
	{"🔥", ":fire:", "Remove code or files.", CommitTypeRelease},
	{"🐛", ":bug:", "Fix a bug.", CommitTypeFix},
	{"🚑️", ":ambulance:", "Critical hotfix.", CommitTypeHotfix},
	{"✨", ":sparkles:", "Introduce new features.", CommitTypeFeat},
	{"📝", ":memo:", "Add or update documentation.", CommitTypeDocs},
	{"🚀", ":rocket:", "Deploy stuff.", CommitTypeDeploy},
	{"🎉", ":tada:", "Begin a project.", CommitTypeChore},
	{"✅", ":white_check_mark:", "Add, update, or pass tests.", CommitTypeTest},
	{"🔒️", ":lock:", "Fix security issues.", CommitTypeFix},
	{"🔖", ":bookmark:", "Release / Version tags.", CommitTypeRelease},
	{"🚨", ":rotating_light:", "Fix compiler / linter warnings.", CommitTypeStyle},
	{"🚧", ":construction:", "Work in progress.", CommitTypeFeat},
	{"💚", ":green_heart:", "Fix CI Build.", CommitTypeChore},
	{"⬇️", ":arrow_down:", "Downgrade dependencies.", CommitTypeChore},
	{"⬆️", ":arrow_up:", "Upgrade dependencies.", CommitTypeChore},
	{"📌", ":pushpin:", "Pin dependencies to specific versions.", CommitTypeChore},
	{"👷", ":construction_worker:", "Add or update CI build system.", CommitTypeChore},
	{"📈", ":chart_with_upwards_trend:", "Add or update analytics or track code.", CommitTypeFeat},
	{"♻️", ":recycle:", "Refactor code.", CommitTypeRefactor},
	{"➕", ":heavy_plus_sign:", "Add a dependency.", CommitTypeChore},
	{"➖", ":heavy_minus_sign:", "Remove a dependency.", CommitTypeChore},
	{"🔧", ":wrench:", "Add or update configuration files.", CommitTypeChore},
	{"🔨", ":hammer:", "Add or update development scripts.", CommitTypeChore},
	{"🌐", ":globe_with_meridians:", "Internationalization and localization.", CommitTypeFeat},
	{"✏️", ":pencil2:", "Fix typos.", CommitTypeFix},
	{"💩", ":poop:", "Write bad code that needs to be improved.", CommitTypeFeat},
	{"⏪️", ":rewind:", "Revert changes.", CommitTypeRevert},
	{"📦️", ":package:", "Add or update compiled files or packages.", CommitTypeChore},
	{"👽️", ":alien:", "Update code due to external API changes.", CommitTypeFeat},
	{"🚚", ":truck:", "Move or rename resources (e.g.: files, paths, routes).", CommitTypeRevert},
	{"📄", ":page_facing_up:", "Add or update license.", CommitTypeChore},
	{"💥", ":boom:", "Introduce breaking changes.", CommitTypeChore},
	{"🍱", ":bento:", "Add or update assets.", CommitTypeChore},
	{"💡", ":bulb:", "Add or update comments in source code.", CommitTypeChore},
	{"💬", ":speech_balloon:", "Add or update text and literals.", CommitTypeChore},
	{"🔊", ":loud_sound:", "Add or update logs.", CommitTypeRefactor},
	{"🔇", ":mute:", "Remove logs.", CommitTypeRefactor},
	{"👥", ":busts_in_silhouette:", "Add or update contributor(s).", CommitTypeChore},
	{"🤡", ":clown_face:", "Mock things.", CommitTypeTest},
	{"🙈", ":see_no_evil:", "Add or update a .gitignore file.", CommitTypeChore},
	{"📸", ":camera_flash:", "Add or update snapshots.", CommitTypeChore},
	{"🗑️", ":wastebasket:", "Deprecate code that needs to be cleaned up.", CommitTypeRefactor},
	{"🩹", ":adhesive_bandage:", "Simple fix for a non-critical issue.", CommitTypeFix},
	{"⚰️", ":coffin:", "Remove dead code.", CommitTypeRefactor},
	{"🧪", ":test_tube:", "Add a failing test.", CommitTypeTest},
}

var emojiSelectStrings = []string{}

func init() {
	for _, v := range Emojis {
		emojiSelectStrings = append(emojiSelectStrings, fmt.Sprintf("%s: %s", v.Emoji, v.Desc))
	}
}
