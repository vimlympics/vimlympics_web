package events

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/vimlympics/vimlympics_web/model"
)

var (
	Summaries = model.Summaries{}
	Events    = model.Events{}
)

func init() {
	// load events from json files
	Events[model.Checkpoint] = make(map[model.Level]model.Event)
	basepath := "./leveldefs"
	err := filepath.Walk(basepath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".json" {
			dir := filepath.Dir(path)
			eventtype, _ := filepath.Rel(basepath, dir)
			filename := filepath.Base(path)
			level := model.Level(filename[0] - '0')
			fmt.Printf("loading %s %v\n", eventtype, level)
			content, err := os.ReadFile(path)
			if err != nil {
				return nil
			}
			switch eventtype {
			case "checkpoint":
				var checkpoint model.Event
				err := json.Unmarshal(content, &checkpoint)
				if err != nil {
					fmt.Println(err)
					return nil
				}
				checkpoint.Type = eventtype
				checkpoint.Name = fmt.Sprintf("%s %v", eventtype, level)
				Events[model.Checkpoint][level] = checkpoint
			}
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	for eventType, levelMap := range Events {
		Summaries[eventType] = make(map[model.Level]string)
		for level, event := range levelMap {
			Summaries[eventType][level] = event.Name
		}
	}
}

func GetEvent(level model.Level, event model.EventType) (model.Event, error) {
	eventMap, ok := Events[event]
	if !ok {
		return model.Event{}, fmt.Errorf("event type %v not found", event)
	}
	levelEvent, ok := eventMap[level]
	if !ok {
		return model.Event{}, fmt.Errorf("level %v not found for event type %v", level, event)
	}
	return levelEvent, nil
}

func ListEvents() model.Summaries {
	return Summaries
}
