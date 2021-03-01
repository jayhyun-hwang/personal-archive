package common

import (
	"fmt"
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

type Version struct {
	Major int
	Minor int
	Patch int
}

func NewVersion(verInfo string) (*Version, error) {
	splited := strings.Split(verInfo, ".")
	if len(splited) != 3 {
		return nil, fmt.Errorf("invalid ver info: %s", verInfo)
	}

	major, err := strconv.Atoi(splited[0])
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse major version (%s)", verInfo)
	}
	minor, err := strconv.Atoi(splited[1])
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse minor version (%s)", verInfo)
	}
	patch, err := strconv.Atoi(splited[2])
	if err != nil {
		return nil, errors.Wrapf(err, "failed to parse patch version (%s)", verInfo)
	}

	return &Version{
		Major: major,
		Minor: minor,
		Patch: patch,
	}, nil
}

func (v *Version) IsLessThan(v1 *Version) bool {
	if v.Major < v1.Major {
		return true
	}
	if v.Minor < v1.Minor {
		return true
	}
	if v.Patch < v1.Minor {
		return true
	}
	return false
}
