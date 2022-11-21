package main

import (
	"context"
	log "github.com/sirupsen/logrus"
	"github.com/stssc/transaction"
)

func main() {
	step1 := Step1{
		PStep: transaction.PStep{
			Name:          "step1",
			PrepareParam:  Step1PrepareParam{Param1: 1, Param2: "param"},
			CommitParam:   Step1CommitPram{Param1: "param"},
			RollbackParam: Step1RollbackParam{},
		},
	}
	step2 := Step2{
		PStep: transaction.PStep{
			Name:          "step2",
			PrepareParam:  Step2PrepareParam{Param1: 1, Param2: "param"},
			CommitParam:   Step2CommitPram{Param1: "param"},
			RollbackParam: Step2RollbackParam{},
		},
	}
	t := &transaction.Transaction{
		Steps: []transaction.IStep{step1, step2},
	}
	result := t.Do(context.Background())
	switch result.Stage {
	case transaction.StagePrepare:
		log.Error("transaction prepare failed. prepare error:", result.Errs)
	case transaction.StageCommit:
		log.Info("transaction commit succeeded.")
	case transaction.StageRollback:
		log.Error("transaction commit failed. commit error:", result.Errs)
	}
}
