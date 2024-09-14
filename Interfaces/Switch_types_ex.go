package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (m directMessage) importance() int {
	if m.isUrgent {
		return 50
	}
	return m.priorityLevel
}

func (m groupMessage) importance() int {
	return m.priorityLevel
}

func (m systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch m := n.(type) {
	case directMessage:
		return m.senderUsername, m.importance()
	case groupMessage:
		return m.groupName, m.importance()
	case systemAlert:
		return m.alertCode, m.importance()
	default:
		return "", 0
	}
}