package stage

import "github.com/flant/dapp/pkg/config"

func GenerateArtifactImportBeforeSetupStage(dimgBaseConfig *config.DimgBase) Interface {
	imports := getImports(dimgBaseConfig, &getImportsOptions{Before: Setup})
	if len(imports) != 0 {
		return newArtifactImportBeforeSetupStage(imports)
	}

	return nil
}

func newArtifactImportBeforeSetupStage(imports []*config.ArtifactImport) *ArtifactImportBeforeSetupStage {
	s := &ArtifactImportBeforeSetupStage{}
	s.ArtifactImportStage = newArtifactImportStage(imports)
	return s
}

type ArtifactImportBeforeSetupStage struct {
	*ArtifactImportStage
}

func (s *ArtifactImportBeforeSetupStage) Name() StageName {
	return ArtifactImportBeforeSetup
}
