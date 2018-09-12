package npm

import (
	"github.com/fossas/fossa-cli/errors"
	"github.com/fossas/fossa-cli/files"
	"github.com/fossas/fossa-cli/pkg"
)

type manifest struct {
	Name         string
	Version      string
	Dependencies map[string]string
}

func PackageFromManifest(filename string) (pkg.Package, error) {
	var manifest manifest
	err := files.ReadJSON(&manifest, filename)
	if err != nil {
		return pkg.Package{}, err
	}

	return convertManifestToPkg(manifest), nil
}

func FromNodeModules(dir string) (pkg.Package, error) {
	return pkg.Package{}, errors.ErrNotImplemented
}

type Lockfile struct {
	Dependencies map[string]struct {
		Version  string
		Requires map[string]string
	}
}

func FromLockfile(filename string) (Lockfile, error) {
	return Lockfile{}, errors.ErrNotImplemented
}

func convertManifestToPkg(manifest manifest) pkg.Package {
	id := pkg.ID{
		Type:     pkg.NodeJS,
		Name:     manifest.Name,
		Revision: manifest.Version,
	}

	var imports pkg.Imports
	for depName, version := range manifest.Dependencies {
		imports = append(imports, createImport(depName, version))
	}

	return pkg.Package{
		ID:      id,
		Imports: imports,
	}
}

func createImport(packageName string, version string) pkg.Import {
	id := pkg.ID{
		Type:     pkg.NodeJS,
		Name:     packageName,
		Revision: version,
	}

	return pkg.Import{
		Target:   packageName,
		Resolved: id,
	}
}
