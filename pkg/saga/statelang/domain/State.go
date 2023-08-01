package domain

type State interface {

	/**
	 * name
	 *
	 * @return the state name
	 */
	getName() string

	/**
	 * comment
	 *
	 * @return the state comment
	 */
	getComment() string

	/**
	 * type
	 *
	 * @return the state type
	 */
	getType() string

	/**
	 * next state name
	 *
	 * @return the next state name
	 */
	getNext() string

	/**
	 * extension properties
	 *
	 * @return the state extensions
	 */
	getExtensions() map[string]interface{}

	/**
	 * state machine instance
	 *
	 * @return the state machine
	 */
	getStateMachine() StateMachine
}
