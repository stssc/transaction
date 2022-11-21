package transaction

type IStep interface {
	Prepare(IPrepareParam) error
	Commit(ICommitParam) error
	Rollback(IRollbackParam) error

	GetName() string
	GetPrepareParam() IPrepareParam
	GetCommitParam() ICommitParam
	GetRollbackParam() IRollbackParam
}

type PStep struct {
	Name          string
	PrepareParam  IPrepareParam
	CommitParam   ICommitParam
	RollbackParam IRollbackParam
}

type IPrepareParam interface{}

type ICommitParam interface{}

type IRollbackParam interface{}

func (step PStep) Prepare(param IPrepareParam) error {
	panic("implement me")
}

func (step PStep) Commit(param ICommitParam) error {
	panic("implement me")
}

func (step PStep) Rollback(param IRollbackParam) error {
	panic("implement me")
}

func (step PStep) GetName() string {
	return step.Name
}

func (step PStep) GetPrepareParam() IPrepareParam {
	return step.PrepareParam
}

func (step PStep) GetCommitParam() ICommitParam {
	return step.CommitParam
}

func (step PStep) GetRollbackParam() IRollbackParam {
	return step.RollbackParam
}
