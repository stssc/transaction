package transaction

import (
	"context"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Transaction struct {
	Steps []IStep
}

type Result struct {
	Stage TransactionStage // 结束在哪一阶段
	Errs  []error          // 报错信息
}

type TransactionStage int

const (
	StagePrepare = iota + 1
	StageCommit
	StageRollback
)

func (transaction *Transaction) Do(ctx context.Context) Result {
	// prepare
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(len(transaction.Steps))
	prepareErrs := make(chan error, len(transaction.Steps))
	for _, step := range transaction.Steps {
		go func(step IStep) {
			if err := step.Prepare(step.GetPrepareParam()); err != nil {
				prepareErrs <- errors.Wrapf(err, "step %s prepare error", step.GetName())
			}
			waitGroup.Done()
		}(step)
	}
	waitGroup.Wait()
	close(prepareErrs)
	if len(prepareErrs) > 0 {
		return Result{
			Stage: StagePrepare,
			Errs: func(errChan chan error) []error {
				errs := []error{}
				for err := range errChan {
					errs = append(errs, err)
				}
				return errs
			}(prepareErrs),
		}
	}
	// commit
	commitErrs := []error{}
	var rollbackLocation int
	for i, step := range transaction.Steps {
		if err := step.Commit(step.GetCommitParam()); err != nil {
			// todo: retry
			rollbackLocation = i
			commitErrs = append(commitErrs, errors.Wrapf(err, "step %s commit error", step.GetName()))
			goto rollback
		}
	}
	return Result{
		Stage: StageCommit,
		Errs:  nil,
	}
	// rollback
rollback:
	rollbackErrs := []error{}
	for i := rollbackLocation; i >= 0; i-- {
		step := transaction.Steps[i]
		if err := step.Rollback(step.GetRollbackParam()); err != nil {
			// todo: rollback retry
			rollbackErrs = append(rollbackErrs, errors.Wrapf(err, "step %s rollback error", step.GetName()))
			// send a message if rollback failed
			log.WithFields(
				log.Fields{
					"traceId": ctx.Value("traceId"),
					"message": errors.Wrapf(err, "step %s rollback error", step.GetName()),
				},
			).Error()
		}
	}
	return Result{
		Stage: StageRollback,
		Errs:  commitErrs,
	}
}
