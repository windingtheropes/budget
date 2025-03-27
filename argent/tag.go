package argent

import ("log"
"github.com/windingtheropes/budget/types")

func TagExists(tag_name string, user_id int) bool {
	tags, err := GetTag(types.UserID(user_id));
	if err != nil {
		log.Fatal(err)
		return false
	}
	if len(tags) == 0 {
		return false
	}
	for i := 0; i < len(tags); i++ {
		tag := tags[i];
		if tag.Name == tag_name {
			return true
		}
	}
	return false
}

func TagExistsOnEntry(tag_id int, entry_id int) bool {
	tags, err := GetEntryTags(entry_id);
	if err != nil {
		log.Fatal(err)
		return false
	}
	if len(tags) == 0 {
		return false
	}
	for i := 0; i < len(tags); i++ {
		tag := tags[i];
		if tag.Id == tag_id {
			return true
		}
	}
	return false
}