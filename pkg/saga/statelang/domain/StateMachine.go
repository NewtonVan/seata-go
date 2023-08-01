package domain

type Status string

const (
	AC Status = "Active"
	IN Status = "Inactive"
)

type RecoverStrategy int

const (
	COMPENSATE RecoverStrategy = 0
	FORWARD    RecoverStrategy = 1
)

type StateMachine interface {

	/**
	 * name
	 *
	 * @return the state machine name
	 */
	getName() string

	/**
	 * comment
	 *
	 * @return the state machine comment
	 */
	getComment() string

	/**
	 * start state name
	 *
	 * @return the start state name
	 */
	getStartState() string

	setStartState(startState string)

	/**
	 * version
	 *
	 * @return the state machine version
	 */
	getVersion() string

	/**
	 * set version
	 *
	 * @param version the state machine version
	 */
	setVersion(version string)

	/**
	 * states
	 *
	 * @return the state machine key: the state machine name,value: the state machine
	 */
	getStates() map[string]State

	/**
	 * get state
	 *
	 * @param name the state machine name
	 * @return the state machine
	 */
	getState(name string) State

	/**
	 * get id
	 *
	 * @return the state machine id
	 */
	getId() string

	setId(id string)

	/**
	 * get tenantId
	 *
	 * @return the tenant id
	 */
	getTenantId() string

	/**
	 * set tenantId
	 *
	 * @param tenantId the tenant id
	 */
	setTenantId(tenantId string)

	/**
	 * app name
	 *
	 * @return the app name
	 */
	getAppName() string

	/**
	 * type, there is only one type: SSL(SEATA state language)
	 *
	 * @return the state type
	 */
	getType() string

	/**
	 * statue (Active|Inactive)
	 *
	 * @return the state machine status
	 */
	getStatus() Status

	/**
	 * recover strategy: prefer compensation or forward when error occurred
	 *
	 * @return the recover strategy
	 */
	getRecoverStrategy() RecoverStrategy

	/**
	 * set RecoverStrategy
	 *
	 * @param recoverStrategy the recover strategy
	 */
	setRecoverStrategy(recoverStrategy RecoverStrategy)

	/**
	 * Is it persist execution log to storage?, default true
	 *
	 * @return is persist
	 */
	isPersist() bool

	/**
	 * Is update last retry execution log, default append new
	 *
	 * @return the boolean
	 */
	isRetryPersistModeUpdate() bool

	/**
	 * Is update last compensate execution log, default append new
	 *
	 * @return the boolean
	 */
	isCompensatePersistModeUpdate() bool

	/**
	 * State language text
	 *
	 * @return the state language text
	 */
	getContent() string

	setContent(content string)

	/**
	 * get create time
	 *
	 * @return the create gmt
	 */
	getGmtCreate() Date

	/**
	 * set create time
	 *
	 * @param date the create gmt
	 */
	setGmtCreate(date Date)
}
