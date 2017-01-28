package dependencyLoader

import (
	"fmt"
	"github.com/suluvir/server/web/handler/externals"
	"os"
	"github.com/suluvir/server/logging"
	"github.com/uber-go/zap"
)

type ExternalFileExtractor struct {
	Externals []External
}

const externalBaseDir = "layout/js/node_modules"

const searchInDirectory = "dist"
const jsFileSuffix = ".min.js"
const cssFileSuffix = ".css"

func NewExternalFileExtractor(externals []External) *ExternalFileExtractor {
	return &ExternalFileExtractor{
		Externals: externals,
	}
}

// this function looks up the external files. It tries an exact match with the external name, for exceptions define
// the file to use in `specialFiles.go`
// FIXME too much duplicated code...
func (e *ExternalFileExtractor) LookupExternalFiles() []External {
	var result []External
	for _, external := range e.Externals {
		externalBaseDir := fmt.Sprintf("%s/%s/%s", externalBaseDir, external.Name, searchInDirectory)
		external.Directory = externalBaseDir

		if specialJsFile, ok := externals.SPECIAL_EXTERNAL_JS_FILES[external.Name]; ok {
			specialJsFileDir := fmt.Sprintf("%s/%s", externalBaseDir, specialJsFile)
			if _, err := os.Stat(specialJsFileDir); err == nil {
				external.JsFile = specialJsFile
				external.HasJs = true
				logging.GetLogger().Info("found special externals file",
					zap.String("path", specialJsFileDir))
			} else {
				external.HasJs = false
				logging.GetLogger().Warn("specal external file given does not exist",
					zap.String("filename", specialJsFile),
					zap.String("expected path", specialJsFileDir))
			}
		} else {
			externalDir := fmt.Sprintf("%s/%s%s", externalBaseDir, external.Name, jsFileSuffix)
			if _, err := os.Stat(externalDir); err == nil {
				external.JsFile = fmt.Sprintf("%s%s", external.Name, jsFileSuffix)
				external.HasJs = true
				logging.GetLogger().Info("found externals file",
					zap.String("path", externalDir))
			} else {
				external.HasJs = false
				logging.GetLogger().Warn("unable to find externals file",
					zap.String("expected path", externalDir))
			}
		}

		if specialCssFile, ok := externals.SPECIAL_EXTERNAL_CSS_FILES[external.Name]; ok {
			specialCssFileDir := fmt.Sprintf("%s/%s", externalBaseDir, specialCssFile)
			if _, err := os.Stat(specialCssFileDir); err == nil {
				external.CssFile = specialCssFile
				external.HasCss = true
				logging.GetLogger().Info("found special externals file",
					zap.String("path", specialCssFileDir))
			} else {
				external.HasCss = false
				logging.GetLogger().Warn("specal external file given does not exist",
					zap.String("filename", specialCssFile),
					zap.String("expected path", specialCssFileDir))
			}
		} else {
			externalDir := fmt.Sprintf("%s/%s%s", externalBaseDir, external.Name, cssFileSuffix)
			if _, err := os.Stat(externalDir); err == nil {
				external.CssFile = fmt.Sprintf("%s%s", external.Name, cssFileSuffix)
				external.HasCss = true
				logging.GetLogger().Info("found externals file",
					zap.String("path", externalDir))
			} else {
				external.HasCss = false
				logging.GetLogger().Warn("unable to find externals file",
					zap.String("extected path", externalDir))
			}
		}
		result = append(result, external)
	}

	return result
}
