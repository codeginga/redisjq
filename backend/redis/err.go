package redis

func errEmpty(err error) bool {
	if err.Error() == "redis: nil" {
		return true
	}

	return false
}
