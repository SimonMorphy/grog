package entity

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

// 在vo包中补充状态描述
func (s PostStatus) String() string {
	switch s {
	case Published:
		return "已发布"
	case Draft:
		return "草稿"
	case Deleted:
		return "已删除"
	case Scheduled:
		return "定时发布"
	case Archived:
		return "已归档"
	case Private:
		return "私密"
	case Locked:
		return "锁定"
	default:
		return "未知状态"
	}
}

var StatusTransitions = map[PostStatus]map[PostStatus]bool{
	Published: {
		Draft:    true, // 发布→草稿（撤回）
		Deleted:  true, // 发布→删除
		Archived: true, // 发布→归档
	},
	Draft: {
		Published: true, // 草稿→发布
		Deleted:   true, // 草稿→删除
	},
	Deleted: {
		Published: true, // 删除→恢复发布
		Draft:     true, // 删除→恢复草稿
	},
	Scheduled: {
		Published: true, // 定时→发布（到时自动发布）
		Draft:     true, // 定时→草稿（取消定时）
	},
	Archived: {
		Published: true, // 归档→重新发布
	},
	Private: {
		Published: true, // 私密→公开
		Deleted:   true, // 私密→删除
	},
	Locked: {
		Published: true, // 锁定→发布（解除锁定）
	},
}

func PrtToStatus(i *int) PostStatus {
	if i == nil {
		return Published
	}
	if *i < 0 || *i > 6 {
		return Published
	}
	return PostStatus(*i)
}
