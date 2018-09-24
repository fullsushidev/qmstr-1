package builder

import "regexp"

var (
	BoolArgs = map[string]struct{}{
		"-w":                struct{}{},
		"-W":                struct{}{},
		"-O":                struct{}{},
		"-f":                struct{}{},
		"-C":                struct{}{},
		"-std":              struct{}{},
		"-nostdinc":         struct{}{},
		"-nostdlib":         struct{}{},
		"-print-file-name":  struct{}{},
		"--print-file-name": struct{}{},
		"-M":                struct{}{},
		"-MG":               struct{}{},
		"-MM":               struct{}{},
		"-MD":               struct{}{},
		"-MMD":              struct{}{},
		"-MP":               struct{}{},
		"-m":                struct{}{},
		"-v":                struct{}{},
		"-g":                struct{}{},
		"-pg":               struct{}{},
		"-P":                struct{}{},
		"-pipe":             struct{}{},
		"-pie":              struct{}{},
		"-pedantic":         struct{}{},
		"-print-":           struct{}{},
		"-pthread":          struct{}{},
		"-rdynamic":         struct{}{},
		"-shared":           struct{}{},
		"-static":           struct{}{},
		"-dynamiclib":       struct{}{},
		"-dumpversion":      struct{}{},
		"-dM":               struct{}{},
		"--version":         struct{}{},
		"-undef":            struct{}{},
		"-nostartfiles":     struct{}{},
		"-remap":            struct{}{},
		"-r":                struct{}{},
	}

	StringArgs = map[string]struct{}{
		"-e":                     struct{}{},
		"-D":                     struct{}{},
		"-B":                     struct{}{},
		"-Q":                     struct{}{},
		"-T":                     struct{}{},
		"-U":                     struct{}{},
		"-x":                     struct{}{},
		"-MF":                    struct{}{},
		"-MT":                    struct{}{},
		"-MQ":                    struct{}{},
		"-install_name":          struct{}{},
		"-compatibility_version": struct{}{},
		"-current_version":       struct{}{},
	}

	StringArgsRE = "\\s+\\S+={0,1}\\S*\\s"

	FixPosixArgs = map[string]struct{}{
		"-isystem": struct{}{},
		"-include": struct{}{},
	}

	StaticLibPattern = regexp.MustCompile("-static-lib(\\w+)")

	LinkBoolArgs = map[string]struct{}{
		"-P": struct{}{},
		"-E": struct{}{},
		"-F": struct{}{},
		"-g": struct{}{},
		"-r": struct{}{},
		"-i": struct{}{},
		"-m": struct{}{},
	}

	LinkStringArgs = map[string]struct{}{}
)
