package tags

import (
	"github.com/mikkyang/id3-go"
	"github.com/suluvir/server/logging"
	"github.com/uber-go/zap"
)

func ExtractTags(fileName string) {
	file, err := id3.Open(fileName)
	defer file.Close()
	if err != nil {
		logging.GetLogger().Error("error loading mp3 file for id extraction", zap.Error(err))
	}
	logging.GetLogger().Info("extracted information",
		zap.String("artist", file.Artist()),
		zap.String("title", file.Title()),
		zap.String("year", file.Year()),
		zap.String("genre", file.Genre()),
		zap.String("album", file.Album()))
}