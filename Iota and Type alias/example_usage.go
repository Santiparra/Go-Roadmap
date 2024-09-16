package main

//IOTA

type emailStatus int 

const(
	emailBounced emailStatus = iota
	emailInvalid
	emailDelivered
	emailOpened
)

//TYPE ALIAS note that this is not safe and can be sided

type sendingChannel string

const (
    Email sendingChannel = "email"
    SMS   sendingChannel = "sms"
    Phone sendingChannel = "phone"
)

func sendNotification(ch sendingChannel, message string) {
    // send the message
}
