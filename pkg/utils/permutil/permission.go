package permutil

func checkPermission(roles []string, required []string) bool {
	for i := range roles {
		for ii := range required {
			if roles[i] == required[ii] {
				return true
			}
		}
	}
	return false
}

func CheckSupporterOrAbove(roles []string) (bool, []string) {
	required := SupporterOrAbove()
	return checkPermission(roles, required), required
}

func CheckModOrAbove(roles []string) (bool, []string) {
	required := ModOrAbove()
	return checkPermission(roles, required), required
}

func CheckSmodOrAbove(roles []string) (bool, []string) {
	required := SmodOrAbove()
	return checkPermission(roles, required), required
}

func CheckAdmin(roles []string) (bool, []string) {
	required := []string{DiscordRoleAdmin}
	return checkPermission(roles, required), required
}

func SupporterOrAbove() []string {
	return []string{
		DiscordRoleAdmin,
		DiscordRoleSmod,
		DiscordRoleMod,
		DiscordRoleSupporter,
	}
}

func ModOrAbove() []string {
	return []string{
		DiscordRoleAdmin,
		DiscordRoleSmod,
		DiscordRoleMod,
	}
}

func SmodOrAbove() []string {
	return []string{
		DiscordRoleAdmin,
		DiscordRoleSmod,
	}
}
