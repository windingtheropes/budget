package argent

import (
	"log"

	"github.com/windingtheropes/budget/types"
)

func UserTagNameExists(tag_name string, user_id int) bool {
	tags, err := GetUserTags(user_id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if len(tags) == 0 {
		return false
	}
	for i := 0; i < len(tags); i++ {
		tag := tags[i]
		if tag.Name == tag_name {
			return true
		}
	}
	return false
}

func GetTransactionTags(entry_id int) ([]types.Tag, error) {
	var tags []types.Tag
	assignments, err := GetTagAssignments(entry_id)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(assignments); i++ {
		res_tags, err := GetTagById(assignments[i].Tag_Id)
		if err != nil {
			return nil, err
		}
		if len(res_tags) == 1 {
			tags = append(tags, res_tags[0])
		}
	}
	return tags, nil
}

func GetUserTags(user_id int) ([]types.Tag, error) {
	var userTags []types.Tag
	ownership_records, err := GetUserTagOwnerships(user_id)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(ownership_records); i++ {
		record := ownership_records[i]
		res_tags, err := GetTagById(record.Tag_Id)
		if err != nil {
			return nil, err
		}
		if len(res_tags) == 1 {
			userTags = append(userTags, res_tags[0])
		}
	}
	return userTags, nil
}
func NewUserTag(name string, user_id int) (int64, error) {
	tag_id, err := NewTag(name)
	if err != nil {
		return 0, err
	}
	if _, err := NewTagOwnership(int(tag_id), user_id); err != nil {
		return 0, err
	}
	return tag_id, nil
}
func TagExistsOnEntry(tag_id int, entry_id int) bool {
	tags, err := GetTagAssignments(entry_id)
	if err != nil {
		log.Fatal(err)
		return false
	}
	if len(tags) == 0 {
		return false
	}
	for i := 0; i < len(tags); i++ {
		tag := tags[i]
		if tag.Id == tag_id {
			return true
		}
	}
	return false
}

func UserOwnsTag(user_id int, tag_id int) bool {
	ownerships, err := GetUserTagOwnerships(user_id)
	if err != nil {
		return false
	}
	if len(ownerships) == 0 {
		return false
	}
	for i := range ownerships {
		ownership := ownerships[i]
		if ownership.Tag_Id == tag_id {
			return true
		}
	}
	return false
}

func TagExists(tag_id int) bool {
	tags, err := GetTagById(tag_id)
	if err != nil {
		return false
	}
	if len(tags) == 0 {
		return false
	} else {
		return true
	}
}
