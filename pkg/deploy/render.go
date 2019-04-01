package deploy

import (
	"bytes"
	"errors"
	"fmt"
	"strings"

	"github.com/flant/logboek"
	"github.com/flant/werf/pkg/config"
	"github.com/flant/werf/pkg/deploy/helm"
	"github.com/flant/werf/pkg/tag_strategy"
)

type RenderOptions struct {
	Values       []string
	SecretValues []string
	Set          []string
	SetString    []string
	Env          string
}

func RunRender(projectDir string, werfConfig *config.WerfConfig, opts RenderOptions) error {
	if debug() {
		fmt.Fprintf(logboek.GetOutStream(), "Render options: %#v\n", opts)
	}

	m, err := GetSafeSecretManager(projectDir, opts.SecretValues)
	if err != nil {
		return fmt.Errorf("cannot get project secret: %s", err)
	}

	releaseName := "RELEASE_NAME"
	imagesRepo := "REPO"
	tag := "GIT_BRANCH"
	tagStrategy := tag_strategy.GitBranch
	namespace := "NAMESPACE"

	images := GetImagesInfoGetters(werfConfig.Images, imagesRepo, tag, true)

	serviceValues, err := GetServiceValues(werfConfig.Meta.Project, imagesRepo, namespace, tag, tagStrategy, images, ServiceValuesOptions{Env: opts.Env})

	werfChart, err := PrepareWerfChart(GetTmpWerfChartPath(werfConfig.Meta.Project), werfConfig.Meta.Project, projectDir, m, opts.SecretValues, serviceValues)
	if err != nil {
		return err
	}
	defer ReleaseTmpWerfChart(werfChart.ChartDir)

	out := &bytes.Buffer{}

	renderOptions := helm.RenderOptions{
		ShowNotes: false,
	}

	if err := helm.Render(
		out,
		werfChart.ChartDir,
		releaseName,
		namespace,
		append(werfChart.Values, opts.Values...),
		append(werfChart.Set, opts.Set...),
		append(werfChart.SetString, opts.SetString...),
		renderOptions,
	); err != nil {
		replaceOld := fmt.Sprintf("%s/", werfChart.Name)
		replaceNew := fmt.Sprintf("%s/", ".helm")
		errMsg := strings.Replace(err.Error(), replaceOld, replaceNew, -1)
		return errors.New(errMsg)
	}

	fmt.Print(out.String())

	return nil
}
