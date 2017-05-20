package log4

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log4/log4go"
	"strings"
)

type LogConfig struct {
	Filter []struct {
		Name        string `json:"name"`
		Filename    string `json:"filename"`
		Format      string `json:"format"`
		DefaultLvl  string `json:"default_lvl"`
		Rotate      bool   `json:"rotate"`
		MaxSize     int    `json:"max_size"`
		MaxLines    int    `json:"max_lines"`
		RotateDaily bool   `json:"rotate_daily"`
	} `json:"filter"`
}

type LogSturct struct {
	Logger map[string]log4go.Logger
	tag    []string
}

func LoadConfiguration(filename string) (LogSturct, error) {

	Logg := LogSturct{}
	var auxLogger = make(map[string]log4go.Logger)
	var auxTag []string

	// read json file
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return Logg, errors.New(fmt.Sprintf("LoadConfiguration: Error: Could not read %q: %s\n", filename, err))

	}

	// unmarshal json
	config := new(LogConfig)
	err = json.Unmarshal(body, &config)
	if err != nil {
		return Logg, errors.New(fmt.Sprintf("LoadConfiguration: Error: Could not parse JSON configuration in %q: %s\n", filename, err))

	}

	for _, filter := range config.Filter {

		flw := log4go.NewFileLogWriter(filter.Filename, false)
		flw.SetFormat(filter.Format)
		flw.SetRotate(filter.Rotate)
		flw.SetRotateDaily(filter.RotateDaily)
		flw.SetRotateSize(filter.MaxSize)
		flw.SetRotateLines(filter.MaxLines)
		lvl, ok := FindLevel(filter.DefaultLvl)
		if ok == false {
			return Logg, errors.New(fmt.Sprintf("LoadConfiguration: Error : Wrong logging level in %q for log_name : %s", filename, filter.Name))
		} else {
			log := log4go.NewLogger()
			log.AddFilter(filter.Name, lvl, flw)
			auxLogger[filter.Name] = log
			auxTag = append(auxTag, strings.ToLower(filter.Name))

		}

	}

	Logg.Logger = auxLogger
	Logg.tag = auxTag

	return Logg, nil
}

func FindLevel(level string) (lvl log4go.Level, ok bool) {
	level = strings.ToLower(level)
	switch level {
	case "finest":
		lvl = log4go.FINEST
		ok = true
	case "fine":
		lvl = log4go.FINE
		ok = true
	case "debug":
		lvl = log4go.DEBUG
		ok = true
	case "error":
		lvl = log4go.ERROR
		ok = true
	case "critical":
		lvl = log4go.CRITICAL
		ok = true
	case "trace":
		lvl = log4go.TRACE
		ok = true
	case "warn":
		lvl = log4go.WARNING
		ok = true
	case "info":
		lvl = log4go.INFO
		ok = true
	default:
		lvl = log4go.FINEST
	}
	return lvl, ok
}

func IsExist(slice []string, name string) bool {
	for _, a := range slice {
		if a == name {
			return true
		}
	}
	return false
}

func (logg LogSturct) WriteLog(name, lvl, message string) {

	name = strings.ToLower(name)
	lvl = strings.ToLower(lvl)

	if IsExist(logg.tag, name) {
		switch lvl {
		case "finest":
			logg.Logger[name].Finest(message)

		case "fine":
			logg.Logger[name].Fine(message)

		case "debug":
			logg.Logger[name].Debug(message)

		case "error":
			logg.Logger[name].Error(message)

		case "critical":
			logg.Logger[name].Critical(message)

		case "trace":
			logg.Logger[name].Trace(message)

		case "warn":
			logg.Logger[name].Warn(message)

		case "info":
			logg.Logger[name].Info(message)

		}
	}

}
