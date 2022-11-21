package main

import (
	"github.com/pkg/errors"
	"github.com/stssc/transaction"
)

type Step2 struct {
	transaction.PStep
}

type Step2PrepareParam struct {
	Param1 int
	Param2 string
}

type Step2CommitPram struct {
	Param1 string
}

type Step2RollbackParam struct {
}

func (step2 Step2) Prepare(param transaction.IPrepareParam) error {
	println(
		"step2 prepare succeeded.",
		"param1:",
		param.(Step2PrepareParam).Param1,
		"param2:",
		param.(Step2PrepareParam).Param2,
	)
	return nil
}

func (step2 Step2) Commit(param transaction.ICommitParam) error {
	println("step2 commit failed. param1:", param.(Step2CommitPram).Param1)
	return errors.Errorf("step2 commit failed. param1: %s", param.(Step2CommitPram).Param1)
}

func (step2 Step2) Rollback(param transaction.IRollbackParam) error {
	println("step2 rollback failed")
	return errors.Errorf("step2 rollback failed")
}
