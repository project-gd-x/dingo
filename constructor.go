package main

import (
	"github.com/elliotchance/pie/pie"
	"regexp"
	"strings"
)

type Constructor struct {
	Name      string
	Arguments []Argument
	Error     bool
}

type Argument string

func (a Argument) IsService() bool {
	return regexp.MustCompile(`@{(.*?)}`).Match([]byte(a))
}

func (a Argument) IsEnv() bool {
	return regexp.MustCompile(`\${(.*?)}`).Match([]byte(a))
}

func (a Argument) IsContainerValue() bool {
	return regexp.MustCompile(`#{(.*?)}`).Match([]byte(a))
}

func (a Argument) IsValue() bool {
	if a.IsService() || a.IsEnv() || a.IsContainerValue() {
		return false
	}
	return true
}

func (a Argument) Name() string {
	if sub := regexp.MustCompile(`[@|#|\$]+{(.*?)}`).FindStringSubmatch(string(a)); len(sub) > 0 {
		return sub[1]
	}

	return ""
}

func (c Constructor) Dependencies() (deps []string) {
	for _, a := range c.Arguments {
		for _, v := range regexp.MustCompile(`@{(.*?)}`).FindAllStringSubmatch(string(a), -1) {
			deps = append(deps, v[1])
		}
	}

	return pie.Strings(deps).Unique()
}

func (c Constructor) PackageName() string {
	if !strings.Contains(c.Name, ".") {
		return ""
	}

	parts := strings.Split(strings.TrimLeft(c.Name, "*"), ".")
	return strings.Join(parts[:len(parts)-1], ".")
}

func (c Constructor) UnversionedPackageName() string {
	packageName := strings.Split(c.PackageName(), "/")
	if regexp.MustCompile(`^v\d+$`).MatchString(packageName[len(packageName)-1]) {
		packageName = packageName[:len(packageName)-1]
	}

	return strings.Join(packageName, "/")
}

func (c Constructor) LocalPackageName() string {
	pkgNameParts := strings.Split(c.UnversionedPackageName(), "/")
	lastPart := pkgNameParts[len(pkgNameParts)-1]

	return strings.Replace(lastPart, "-", "_", -1)
}

func (c Constructor) ReplaceLocalPackageName(newName string) *Constructor {
	pkgNameParts := strings.Split(c.UnversionedPackageName(), "/")
	lastPart := pkgNameParts[len(pkgNameParts)-1]

	c.Name = strings.Replace(c.Name, lastPart+".", newName+".", -1)
	return &c
}

func (c Constructor) EntityName() string {
	parts := strings.Split(c.Name, ".")
	return strings.TrimLeft(parts[len(parts)-1], "*")
}

func (c Constructor) LocalEntityName() string {
	name := c.LocalPackageName() + "." + c.EntityName()
	return strings.TrimLeft(name, ".")
}
