package logging

import (
	"encoding/json"
	"fmt"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func NewLogger() *zap.SugaredLogger {

	configJSON := []byte(
		fmt.Sprintf(
			`{
				"level": "%s",
				"encoding": "json",
				"outputPaths": ["stdout"],
				"errorOutputPaths": ["stdout","stderr"],
				"encoderConfig": {
					"messageKey": "message",
					"levelKey": "level",
					"levelEncoder": "lowercase",
					"timeKey": "@timestamp",
					"timeEncoder": "rfc3339nano"
				}
			}`,
			"debug"),
	)

	var cfg zap.Config
	if err := json.Unmarshal(configJSON, &cfg); err != nil {
		panic(err)
	}

	core := zapcore.NewCore(zapcore.NewJSONEncoder(cfg.EncoderConfig), zapcore.AddSync(os.Stdout), cfg.Level)

	return zap.New(core).Sugar()
}
