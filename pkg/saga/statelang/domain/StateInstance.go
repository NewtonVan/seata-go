package domain

type ExecutionStatus string

const (
	RU ExecutionStatus = "Running"
	SU ExecutionStatus = "Succeed"
	FA ExecutionStatus = "Failed"
	UN ExecutionStatus = "Unknown"
	SK ExecutionStatus = "Skipped"
)

type StateInstance interface {
	// GetId /**
	GetId() string

	// SetId /**
	SetId(id string)

	// GetMachineInstanceId /**
	GetMachineInstanceId() string

	// SetMachineInstanceId /**
	SetMachineInstanceId(machineInstanceId string)

	// GetName /**
	GetName() string

	// SetName /**
	SetName(name string)

	// GetType /**
	GetType() string

	// SetType /**
	SetType(tp string)

	// GetServiceName /**
	GetServiceName() string

	// SetServiceName /**
	SetServiceName(serviceName string)

	// GetServiceMethod /**
	GetServiceMethod() string

	// SetServiceMethod /**
	SetServiceMethod(serviceMethod string)

	/**
	 * get service type
	 *
	 * @return the state instance service type
	 */
	getServiceType() string

	/**
	 * get service type
	 *
	 * @param serviceType the state instance service type
	 */
	setServiceType(serviceType string)

	/**
	 * get businessKey
	 *
	 * @return the state instance businessKey
	 */
	getBusinessKey() string

	/**
	 * set business key
	 *
	 * @param businessKey the state instance businessKey
	 */
	setBusinessKey(businessKey string)

	/**
	 * get start time
	 *
	 * @return the state instance start time
	 */
	getGmtStarted() Date

	/**
	 * set start time
	 *
	 * @param gmtStarted the state instance start time
	 */
	setGmtStarted(gmtStarted Date)

	/**
	 * get update time
	 *
	 * @return the state instance update time
	 */
	getGmtUpdated() Date

	/**
	 * set update time
	 *
	 * @param gmtUpdated the state instance update time
	 */
	setGmtUpdated(gmtUpdated Date)

	/**
	 * get end time
	 *
	 * @return the state instance end time
	 */
	getGmtEnd() Date

	/**
	 * set end time
	 *
	 * @param gmtEnd the state instance end time
	 */
	setGmtEnd(gmtEnd Date)

	/**
	 * Is this state task will update data?
	 *
	 * @return the bool
	 */
	isForUpdate() bool

	/**
	 * setForUpdate
	 *
	 * @param forUpdate is for update
	 */
	setForUpdate(forUpdate bool)

	/**
	 * get exception
	 *
	 * @return exception
	 */
	geterror() error

	/**
	 * set exception
	 *
	 * @param exception exception
	 */
	seterror(exception error)

	/**
	 * get input params
	 *
	 * @return input params
	 */
	getInputParams() interface{}

	/**
	 * set inout params
	 *
	 * @param inputParams inputParams
	 */
	setInputParams(inputParams interface{})

	/**
	 * get output params
	 *
	 * @return output params
	 */
	getOutputParams() interface{}

	/**
	 * Sets set output params.
	 *
	 * @param outputParams the output params
	 */
	setOutputParams(outputParams interface{})

	/**
	 * Gets get status.
	 *
	 * @return the get status
	 */
	getStatus() ExecutionStatus

	/**
	 * Sets set status.
	 *
	 * @param status the status
	 */
	setStatus(status ExecutionStatus)

	/**
	 * Gets get state id compensated for.
	 *
	 * @return the get state id compensated for
	 */
	getStateIdCompensatedFor() string

	/**
	 * Sets set state id compensated for.
	 *
	 * @param stateIdCompensatedFor the state id compensated for
	 */
	setStateIdCompensatedFor(stateIdCompensatedFor string)

	/**
	 * Gets get state id retried for.
	 *
	 * @return the get state id retried for
	 */
	getStateIdRetriedFor() string

	/**
	 * Sets set state id retried for.
	 *
	 * @param stateIdRetriedFor the state id retried for
	 */
	setStateIdRetriedFor(stateIdRetriedFor string)

	/**
	 * Gets get compensation state.
	 *
	 * @return the get compensation state
	 */
	getCompensationState() StateInstance

	/**
	 * Sets set compensation state.
	 *
	 * @param compensationState the compensation state
	 */
	setCompensationState(compensationState StateInstance)

	/**
	 * Gets get state machine instance.
	 *
	 * @return the get state machine instance
	 */
	getStateMachineInstance() StateMachineInstance

	/**
	 * Sets set state machine instance.
	 *
	 * @param stateMachineInstance the state machine instance
	 */
	setStateMachineInstance(stateMachineInstance StateMachineInstance)

	/**
	 * Is ignore status bool.
	 *
	 * @return the bool
	 */
	isIgnoreStatus() bool

	/**
	 * Sets set ignore status.
	 *
	 * @param ignoreStatus the ignore status
	 */
	setIgnoreStatus(ignoreStatus bool)

	/**
	 * Is for compensation bool.
	 *
	 * @return the bool
	 */
	isForCompensation() bool

	/**
	 * Gets get serialized input params.
	 *
	 * @return the get serialized input params
	 */
	getSerializedInputParams() interface{}

	/**
	 * Sets set serialized input params.
	 *
	 * @param serializedInputParams the serialized input params
	 */
	setSerializedInputParams(serializedInputParams interface{})

	/**
	 * Gets get serialized output params.
	 *
	 * @return the get serialized output params
	 */
	getSerializedOutputParams() interface{}

	/**
	 * Sets set serialized output params.
	 *
	 * @param serializedOutputParams the serialized output params
	 */
	setSerializedOutputParams(serializedOutputParams interface{})

	/**
	 * Gets get serialized exception.
	 *
	 * @return the get serialized exception
	 */
	getSerializederror() interface{}

	/**
	 * Sets set serialized exception.
	 *
	 * @param serializederror the serialized exception
	 */
	setSerializederror(serializederror interface{})

	/**
	 * Gets get compensation status.
	 *
	 * @return the get compensation status
	 */
	getCompensationStatus() ExecutionStatus
}
