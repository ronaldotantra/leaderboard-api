package dbutil

func GetLimitOffset(pageSize, page int) (bool, int, int) {
	limit := pageSize
	if limit <= 0 || page < 1 {
		return false, 0, 0
	}
	offset := (page - 1) * limit

	return true, limit, offset
}
