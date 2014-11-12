/*
The dummy task is designed for regresion test of meritop framework.
This works with
*/
package meritop

import (
	"log"
	"os"
)

type DummyMaster struct {
	framework       Framework
	epochID, taskID uint64
	logger          *log.Logger
}

// This is useful to bring the task up to speed from scratch or if it recovers.
func (t DummyMaster) Init(taskID uint64, framework Framework, config Config) {
	t.taskID = taskID
	t.framework = framework
	t.logger = log.New(os.Stdout, "DummyMaster:", log.Ldate|log.Ltime|log.Lshortfile)
}

// Task need to finish up for exit, last chance to save work?
func (t DummyMaster) Exit() {}

// These are called by framework implementation so that task implementation can
// reacts to parent or children restart.
func (t DummyMaster) ParentRestart(parentID uint64) {}
func (t DummyMaster) ChildRestart(childID uint64)   {}

func (t DummyMaster) ParentDie(parentID uint64) {}
func (t DummyMaster) ChildDie(childID uint64)   {}

// Ideally, we should also have the following:
func (t DummyMaster) ParentReady(taskID uint64, meta TGMeta) {
	t.logger.Fatal("There should be no parent for master")
	t.framework.Exit()
}

func (t DummyMaster) ChildReady(taskID uint64, meta TGMeta) {
	// Get data from child. When all the data is back, starts the next epoch.
}

// This give the task an opportunity to cleanup and regroup.
func (t DummyMaster) SetEpoch(epochID uint64) {
	t.epochID = epochID
	t.framework.FlagReadyForChildren(nil)
}

func (t DummyMaster) GetFromParent(req TGMeta) TGMeta {
	return nil
}
func (t DummyMaster) GetFromChild(reg TGMeta) TGMeta {
	return nil
}

type DummySlave struct {
	framework       Framework
	epochID, taskID uint64
	logger          *log.Logger
}

// This is useful to bring the task up to speed from scratch or if it recovers.
func (t DummySlave) Init(taskID uint64, framework Framework, config Config) {
	t.taskID = taskID
	t.framework = framework
	t.logger = log.New(os.Stdout, "DummySlave:", log.Ldate|log.Ltime|log.Lshortfile)
}

// Task need to finish up for exit, last chance to save work?
func (t DummySlave) Exit() {}

// These are called by framework implementation so that task implementation can
// reacts to parent or children restart.
func (t DummySlave) ParentRestart(parentID uint64) {}
func (t DummySlave) ChildRestart(childID uint64)   {}

func (t DummySlave) ParentDie(parentID uint64) {}
func (t DummySlave) ChildDie(childID uint64)   {}

// Ideally, we should also have the following:
func (t DummySlave) ParentReady(taskID uint64, meta TGMeta) {
	// Get data from parent, and then make it available for children.
	if t.framework.HasChildren() {
		t.framework.FlagReadyForChildren(nil)
	} else {
		t.framework.FlagReadyForParent(nil)
	}
}

func (t DummySlave) ChildReady(taskID uint64, meta TGMeta) {
	// Get data from child. When all the data is back, we flag the parent.
	t.framework.FlagReadyForParent(nil)
}

// This give the task an opportunity to cleanup and regroup.
func (t DummySlave) SetEpoch(epochID uint64) {
	t.epochID = epochID

	// this is client node, it should wait for parent ready then inform
	// its children
}

// These are payload rpc for application purpose.
func (t DummySlave) GetFromParent(req TGMeta) TGMeta {
	return nil
}
func (t DummySlave) GetFromChild(reg TGMeta) TGMeta {
	return nil
}

type simpleTaskBuilder struct{}

// This method is called once by framework implementation to get the
// right task implementation for the node/task. It requires the taskID
// for current node, and also a global array of tasks.
func (tc simpleTaskBuilder) GetTask(taskID uint64) Task {

	if taskID == 0 {
		return DummyMaster{}
	} else {
		return DummySlave{}
	}
}

// This is used to show how to drive the network.
func drive() {
	var framework Framework
	var taskBuilder simpleTaskBuilder
	framework.SetTaskBuilder(taskBuilder)
	framework.SetTopology(NewTreeTopology(2, 127))
	framework.Start()
}
