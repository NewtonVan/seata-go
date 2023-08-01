package domain

type StateMachineInstance interface {

	/**
	 * Gets get id.
	 *
	 * @return the get id
	 */
	getId() string

	/**
	 * Sets set id.
	 *
	 * @param id the id
	 */
	setId(id string)

	/**
	 * Gets get machine id.
	 *
	 * @return the get machine id
	 */
	getMachineId() string

	/**
	 * Sets set machine id.
	 *
	 * @param machineId the machine id
	 */
	setMachineId(machineId string)

	/**
	 * Gets get tenant id.
	 *
	 * @return the tenant id
	 */
	getTenantId() string

	/**
	 * Sets set tenant id.
	 *
	 * @param tenantId the tenant id
	 */
	setTenantId(tenantId string)

	/**
	 * Gets get parent id.
	 *
	 * @return the get parent id
	 */
	getParentId() string

	/**
	 * Sets set parent id.
	 *
	 * @param parentId the parent id
	 */
	setParentId(parentId string)

	/**
	 * Gets get gmt started.
	 *
	 * @return the get gmt started
	 */
	getGmtStarted() Date

	/**
	 * Sets set gmt started.
	 *
	 * @param gmtStarted the gmt started
	 */
	setGmtStarted(gmtStarted Date)

	/**
	 * Gets get gmt end.
	 *
	 * @return the get gmt end
	 */
	getGmtEnd() Date

	/**
	 * Sets set gmt end.
	 *
	 * @param gmtEnd the gmt end
	 */
	setGmtEnd(gmtEnd Date)

	/**
	 * Put state instance.
	 *
	 * @param stateId       the state id
	 * @param stateInstance the state instance
	 */
	putStateInstance(stateId string, stateInstance StateInstance)

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
	 * Gets get compensation status.
	 *
	 * @return the get compensation status
	 */
	getCompensationStatus() ExecutionStatus

	/**
	 * Sets set compensation status.
	 *
	 * @param compensationStatus the compensation status
	 */
	setCompensationStatus(compensationStatus ExecutionStatus)

	/**
	 * Is running bool.
	 *
	 * @return the bool
	 */
	isRunning() bool

	/**
	 * Sets set running.
	 *
	 * @param running the running
	 */
	setRunning(running bool)

	/**
	 * Gets get gmt updated.
	 *
	 * @return the get gmt updated
	 */
	getGmtUpdated() Date

	/**
	 * Sets set gmt updated.
	 *
	 * @param gmtUpdated the gmt updated
	 */
	setGmtUpdated(gmtUpdated Date)

	/**
	 * Gets get business key.
	 *
	 * @return the get business key
	 */
	getBusinessKey() string

	/**
	 * Sets set business key.
	 *
	 * @param businessKey the business key
	 */
	setBusinessKey(businessKey string)

	/**
	 * Gets get exception.
	 *
	 * @return the get exception
	 */
	geterror() error

	/**
	 * Sets set exception.
	 *
	 * @param exception the exception
	 */
	seterror(exception error)

	/**
	 * Gets get start params.
	 *
	 * @return the get start params
	 */
	getStartParams() map[string]interface{}

	/**
	 * Sets set start params.
	 *
	 * @param startParams the start params
	 */
	setStartParams(startParams map[string]interface{})

	/**
	 * Gets get end params.
	 *
	 * @return the get end params
	 */
	getEndParams() map[string]interface{}

	/**
	 * Sets set end params.
	 *
	 * @param endParams the end params
	 */
	setEndParams(endParams map[string]interface{})

	/**
	 * Gets get context.
	 *
	 * @return the state machine context
	 */
	getContext() map[string]interface{}

	/**
	 * Sets set context.
	 *
	 * @param context the key and value context
	 */
	setContext(context map[string]interface{})

	/**
	 * Gets get state machine.
	 *
	 * @return the get state machine
	 */
	getStateMachine() StateMachine

	/**
	 * Sets set state machine.
	 *
	 * @param stateMachine the state machine
	 */
	setStateMachine(stateMachine StateMachine)

	/**
	 * Gets get state list.
	 *
	 * @return the get state list
	 */
	getStateList() []StateInstance

	/**
	 * Sets set state list.
	 *
	 * @param stateList the state list
	 */
	setStateList(stateList []StateInstance)

	/**
	 * Gets get state map.
	 *
	 * @return the get state map
	 */
	getStateMap() map[string]StateInstance

	/**
	 * Sets set state map.
	 *
	 * @param stateMap the state map
	 */
	setStateMap(stateMap map[string]StateInstance)

	/**
	 * Gets get serialized start params.
	 *
	 * @return the get serialized start params
	 */
	getSerializedStartParams() interface{}

	/**
	 * Sets set serialized start params.
	 *
	 * @param serializedStartParams the serialized start params
	 */
	setSerializedStartParams(serializedStartParams interface{})

	/**
	 * Gets get serialized end params.
	 *
	 * @return the get serialized end params
	 */
	getSerializedEndParams() interface{}

	/**
	 * Sets set serialized end params.
	 *
	 * @param serializedEndParams the serialized end params
	 */
	setSerializedEndParams(serializedEndParams interface{})

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
}
