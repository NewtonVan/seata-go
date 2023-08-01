package domain

type ServiceTaskState struct {
}

func (s *ServiceTaskState) getName() string {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceTaskState) getComment() string {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceTaskState) getType() string {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceTaskState) getNext() string {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceTaskState) getExtensions() map[string]interface{} {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceTaskState) getStateMachine() StateMachine {
	//TODO implement me
	panic("implement me")
}

// Below are for service task
func (s *ServiceTaskState) getServiceType() string {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceTaskState) getServiceName() string {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceTaskState) getServiceMethod() string {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceTaskState) getParameterTypes() []string {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceTaskState) isPersist() bool {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceTaskState) isRetryPersistModeUpdate() *bool {
	//TODO implement me
	panic("implement me")
}

func (s *ServiceTaskState) isCompensatePersistModeUpdate() *bool {
	//TODO implement me
	panic("implement me")
}
