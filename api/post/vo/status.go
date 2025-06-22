package vo

type PostStatus int

const (
	Published PostStatus = iota
	Draft
	Deleted
	Scheduled
	Archived
	Private
	Locked
)

func PrtToStatus(i *int) PostStatus {
	if i == nil {
		return Published
	}
	if *i < 0 || *i > 6 {
		return Published
	}
	return PostStatus(*i)
}
