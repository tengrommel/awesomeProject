package worker

import (
	"awesomeProject/crontab/common"
	"context"
	"go.etcd.io/etcd/clientv3"
)

// 分布式锁
type JobLock struct {
	kv         clientv3.KV
	lease      clientv3.Lease
	jobName    string
	cancelFunc context.CancelFunc
	leaseId    clientv3.LeaseID
	isLocked   bool
}

func InitJobLock(jobName string, kv clientv3.KV, lease clientv3.Lease) (jobLock *JobLock) {
	jobLock = &JobLock{
		kv:      kv,
		lease:   lease,
		jobName: jobName,
	}
	return
}

func (jobLock *JobLock) TryLock() (err error) {

	var (
		leaseGrantResp *clientv3.LeaseGrantResponse
		cancelCtx      context.Context
		cancelFunc     context.CancelFunc
		keepRespChan   <-chan *clientv3.LeaseKeepAliveResponse
		leaseID        clientv3.LeaseID
		txResp         *clientv3.TxnResponse
		txn            clientv3.Txn
		lockKey        string
	)
	if leaseGrantResp, err = jobLock.lease.Grant(context.TODO(), 5); err != nil {
		return
	}
	cancelCtx, cancelFunc = context.WithCancel(context.TODO())
	leaseID = leaseGrantResp.ID
	if keepRespChan, err = jobLock.lease.KeepAlive(cancelCtx, leaseID); err != nil {
		goto FAIL
	}

	go func() {
		var (
			keepResp *clientv3.LeaseKeepAliveResponse
		)
		for {
			select {
			case keepResp = <-keepRespChan:
				if keepResp == nil {
					goto END
				}
			}
		}
	END:
	}()
	//
	txn = jobLock.kv.Txn(context.TODO())
	lockKey = common.JOB_LOCK_DIR + jobLock.jobName

	txn.If(clientv3.Compare(clientv3.CreateRevision(lockKey), "=", 0)).
		Then(clientv3.OpPut(lockKey, "", clientv3.WithLease(leaseID))).
		Else(clientv3.OpGet(lockKey))
	if txResp, err = txn.Commit(); err != nil {
		goto FAIL
	}
	if !txResp.Succeeded {
		err = common.ERR_LOCK_ALREADY_REQUIRED
		goto FAIL
	}
	jobLock.leaseId = leaseID
	jobLock.cancelFunc = cancelFunc
	jobLock.isLocked = true
	return
FAIL:
	cancelFunc()
	jobLock.lease.Revoke(context.TODO(), leaseID)
	return
}

func (jobLock *JobLock) Unlock() {
	if jobLock.isLocked {
		jobLock.cancelFunc()
		jobLock.lease.Revoke(context.TODO(), jobLock.leaseId)
	}
}
