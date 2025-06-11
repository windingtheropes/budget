package argent

import (
	"log"

	"github.com/windingtheropes/budget/tables"
	"github.com/windingtheropes/budget/types"
)

func UserTagNameExists(tag_name string, user_id int64) bool {
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

func GetTransactionTags(entry_id int64) ([]types.Tag, error) {
	var tags []types.Tag
	assignments, err := tables.TagAssignment.Get("entry_id = ?", entry_id)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(assignments); i++ {
		res_tags, err := tables.Tag.Get("(id=?)",assignments[i].Tag_Id)
		if err != nil {
			return nil, err
		}
		if len(res_tags) == 1 {
			tags = append(tags, res_tags[0])
		}
	}
	return tags, nil
}

func GetUserTags(user_id int64) ([]types.Tag, error) {
	var userTags []types.Tag
	ownership_records, err := tables.TagOwnership.Get("user_id=?", user_id)
	if err != nil {
		return nil, err
	}
	for i := 0; i < len(ownership_records); i++ {
		record := ownership_records[i]
		res_tags, err := tables.Tag.Get("(id=?)",record.Tag_Id)
		if err != nil {
			return nil, err
		}
		if len(res_tags) == 1 {
			userTags = append(userTags, res_tags[0])
		}
	}
	return userTags, nil
}
func NewUserTag(name string, user_id int64) (int64, error) {
	tag_id, err := tables.Tag.New(types.TagForm{Name: name})
	if err != nil {
		return 0, err
	}
	if _, err := tables.TagOwnership.New(types.TagOwnershipForm{User_Id: user_id, Tag_Id: int64(tag_id)}); err != nil {
		return 0, err
	}
	return tag_id, nil
}
func TagExistsOnEntry(tag_id int64, entry_id int64) bool {
	tags, err := tables.TagAssignment.Get("entry_id = ?", entry_id)
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

func UserOwnsTag(user_id int64, tag_id int64) bool {
	ownerships, err := tables.TagOwnership.Get("user_id=?",user_id)
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

func TagExists(tag_id int64) bool {
	tags, err := tables.Tag.Get("(id=?)",tag_id)
	if err != nil {
		return false
	}
	if len(tags) == 0 {
		return false
	} else {
		return true
	}
}
