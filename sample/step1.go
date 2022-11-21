package main

import (
	"github.com/stssc/transaction"
)

type Step1 struct {
	transaction.PStep
}

type Step1PrepareParam struct {
	Param1 int
	Param2 string
}

type Step1CommitPram struct {
	Param1 string
}

type Step1RollbackParam struct {
}

func (step1 Step1) Prepare(param transaction.IPrepareParam) error {
	println(
		"step1 prepare succeeded.",
		"param1:",
		param.(Step1PrepareParam).Param1,
		"param2:",
		param.(Step1PrepareParam).Param2,
	)
	return nil
}

func (step1 Step1) Commit(param transaction.ICommitParam) error {
	println("step1 commit succeeded.", "param1:", param.(Step1CommitPram).Param1)
	return nil
}

func (step1 Step1) Rollback(param transaction.IRollbackParam) error {
	println("step1 rollback succeeded.")
	return nil
}
