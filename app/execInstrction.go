package app

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"robotcleaner/domain"
)

func ExecInstraction(inssInput string, cleaner *domain.RobotCleaner) error {
	input, err := ioutil.ReadFile(inssInput)
	if err != nil {
		fmt.Printf("read file err, %v", err)
		return err
	}
	inss := InsTraction{}
	if err := json.Unmarshal(input, &inss); err != nil {
		fmt.Printf("unmarshal err, %v", err)
	}

	inssModel := convetInssToModel(inss)
	inssModel.Do(cleaner)
	return nil
}

type InsTraction struct {
	Inss []Ins
}

type Ins struct {
	In string
}

func convetInssToModel(inss InsTraction) domain.InstractionInterface {
	return &domain.Sequence{}
}
