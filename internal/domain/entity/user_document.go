package entity

import "time"

type UserDocument struct {
	ID         uint
	UserID     uint
	DocumentID uint
	CreatedAt  time.Time
}

type UserDocumentArray []UserDocument

func NewUserDocumentArray(items []UserDocument) *UserDocumentArray {
	a := make(UserDocumentArray, len(items))
	for i, v := range items {
		a[i] = v
	}
	return &a
}

func (a *UserDocumentArray) IDs() []uint {
	ids := make([]uint, len(*a))
	for i, id := range *a {
		ids[i] = id.ID
	}
	return ids
}

func (a *UserDocumentArray) DocumentIDs() []uint {
	ids := make([]uint, len(*a))
	for i, id := range *a {
		ids[i] = id.DocumentID
	}
	return ids
}
