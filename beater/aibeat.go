package beater

import (
	"fmt"
	"time"

	"github.com/elastic/beats/libbeat/beat"
	"github.com/elastic/beats/libbeat/common"
	"github.com/elastic/beats/libbeat/logp"
	"github.com/elastic/beats/libbeat/publisher"

	"github.com/aryou/aibeat/config"
)

type Aibeat struct {
	beatConfig *config.Config
	done       chan struct{}
	period     time.Duration
	client     publisher.Client
}

// Creates beater
func New() *Aibeat {
	return &Aibeat{
		done: make(chan struct{}),
	}
}

/// *** Beater interface methods ***///

func (bt *Aibeat) Config(b *beat.Beat) error {

	// Load beater beatConfig
	err := b.RawConfig.Unpack(&bt.beatConfig)
	if err != nil {
		return fmt.Errorf("Error reading config file: %v", err)
	}

	aiBeatSection := "aibeat"
	rawAibeatConfig, err := b.RawConfig.Child(aiBeatSection, -1)
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return err
	}
	// fmt.Println(rawAibeatConfig)
	
	err = rawAibeatConfig.Unpack(&bt.beatConfig.Aibeat)
	if err != nil {
		logp.Err("Error reading configuration file: %v", err)
		return err
	}
	aibeatConfig := bt.beatConfig.Aibeat
	fmt.Printf("beat config = %#v \n", aibeatConfig.Test)
	return nil
}

func (bt *Aibeat) Setup(b *beat.Beat) error {

	// Setting default period if not set
	if bt.beatConfig.Aibeat.Period == "" {
		bt.beatConfig.Aibeat.Period = "1s"
	}

	bt.client = b.Publisher.Connect()

	var err error
	bt.period, err = time.ParseDuration(bt.beatConfig.Aibeat.Period)
	if err != nil {
		return err
	}

	return nil
}

func (bt *Aibeat) Run(b *beat.Beat) error {
	logp.Info("aibeat is running! Hit CTRL-C to stop it.")
	ticker := time.NewTicker(bt.period)
	counter := 1
	for {
		select {
		case <-bt.done:
			return nil
		case <-ticker.C:
		}

		event := common.MapStr{
			"@timestamp": common.Time(time.Now()),
			"type":       b.Name,
			"counter":    counter,
		}
		bt.client.PublishEvent(event)
		logp.Info("Event sent")
		counter++
		fmt.Println(counter)
	}
}

func (bt *Aibeat) Cleanup(b *beat.Beat) error {
	return nil
}

func (bt *Aibeat) Stop() {
	close(bt.done)
}
