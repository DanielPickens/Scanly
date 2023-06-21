package runtime

type eventchannel chan event

type event struct {
	stdout: string 
	stderr: string
	err: string
	erroronexit: string
}

func (ec eventchannel) message(msg string) {
	ec <- event {
		stdout:msg,
	}
}

func (ec eventchannel) exitwitherror(err error) {
	ec <- event {
		err:        		 err,
		erroronexit: 		true,

	}
}

func (ec eventchannel) exitwithouterrormsg(err error) {
	ec <- event {
		stderr:				msg,
		err: 				err,
		exitwithouterrormsg: true,
	}
}



