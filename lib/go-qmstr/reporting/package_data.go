package reporting

import (
	"encoding/json"

	"github.com/QMSTR/qmstr/lib/go-qmstr/module"
	"github.com/QMSTR/qmstr/lib/go-qmstr/service"
)

// PackageData contains metadata about a specific package.
type PackageData struct {
	Name            string
	Version         string // Usually a Git hash, but any string can be used
	LicenseDeclared string
	Targets         []*Target
	Project         string // The project this package is associated with.
}

type Target struct {
	Target     *service.FileNode
	Licenses   []string
	Authors    []string
	Copyrights []string
}

func GetPackageData(pkg *module.PackageNodeProxy, projectName string) (*PackageData, error) {
	packageData := PackageData{
		Name:            pkg.GetName(),
		Version:         pkg.GetVersion(),
		LicenseDeclared: "NO License declared",
		Project:         projectName,
	}

	if packageData.Version == "" {
		packageData.Version = "default"
	}

	packageData.Project = service.RemoveSlash(packageData.Project)
	packageData.Name = service.RemoveSlash(packageData.Name)
	packageData.Version = service.RemoveSlash(packageData.Version)

	targets := []*Target{}
	targetProxies, err := pkg.GetTargets()
	if err != nil {
		return nil, err
	}
	for _, fileNodeProxy := range targetProxies {
		targets = append(targets, &Target{Target: &fileNodeProxy.FileNode})
	}
	packageData.Targets = targets

	return &packageData, nil
}

// TODO: GetAuthors, GetLicenses and GetCopyrights can be a single function!
func (p *PackageData) GetAuthors() []string {
	authors := map[string]struct{}{}
	for _, target := range p.Targets {
		for _, author := range target.Authors {
			authors[author] = struct{}{}
		}
	}

	uniqAuthors := []string{}

	for author := range authors {
		uniqAuthors = append(uniqAuthors, author)
	}
	return uniqAuthors
}

func (p *PackageData) GetLicenses() []string {
	licenses := map[string]struct{}{}
	for _, target := range p.Targets {
		for _, license := range target.Licenses {
			licenses[license] = struct{}{}
		}
	}

	uniqLicenses := []string{}

	for license := range licenses {
		uniqLicenses = append(uniqLicenses, license)
	}
	return uniqLicenses
}

func (p *PackageData) GetCopyrights() []string {
	copyrights := map[string]struct{}{}
	for _, target := range p.Targets {
		for _, copyright := range target.Copyrights {
			copyrights[copyright] = struct{}{}
		}
	}

	uniqCopyrights := []string{}

	for copyright := range copyrights {
		uniqCopyrights = append(uniqCopyrights, copyright)
	}
	return uniqCopyrights
}

func (p *PackageData) MarshalJSON() ([]byte, error) {
	type Alias PackageData
	return json.Marshal(&struct {
		Authors    []string
		Licenses   []string
		Copyrights []string
		*Alias
	}{
		Authors:    p.GetAuthors(),
		Licenses:   p.GetLicenses(),
		Copyrights: p.GetCopyrights(),
		Alias:      (*Alias)(p),
	})
}
