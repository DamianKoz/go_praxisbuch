package goals

type SubGoal struct {
	id   int
	name string
}

type GoalStatus struct {
	status string
}

type Goal struct {
	id            int
	name          string
	description   string
	subgoals      []SubGoal
	currentStatus GoalStatus
}
