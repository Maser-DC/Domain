package domain

import "time"

type Activity struct {
	ActivityID   string `json:"ActivityID"`
	OwnerID      string `json:"OwnerID"`
	Title        string `json:"Title"`
	ActivityType string `json:"ActivityType"`
	Description  string `json:"Description"`
	DateCreated  time.Time
	Maintainers  []Maintainer
}
type ActivityRepository interface {
	Store(activity Activity) error
	Delete(ActivityID string) error
	FindById(activityID string) (Activity, error)
	Update(activity Activity) error
	Search(searchText string, category []string, limit int, offset int) ([]Activity, error)
}
