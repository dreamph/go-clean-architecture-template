package messages

var (
	//Error
	E10001 = New("E10001", "internal server error")
	E10002 = New("E10002", "internal server error")
	E00001 = New("E00001", "%s not found")
	E00002 = New("E00002", "%s field is invalid format")
	E00003 = New("E00003", "%s must more than 0")
	E00004 = New("E00004", "call external service error : [system:%s, api:%s]")
	E00005 = New("E00005", "%s field must be %s digits")
	E00006 = New("E00006", "%s must between %s and %s")
	E00007 = New("E00007", "%s field not allow to update")
	E00008 = New("E00008", "%s field is required")
	E00009 = New("E00009", "%s already exists")
	E00010 = New("E00010", "%s has been already transferred")
	E00011 = New("E00011", "invalid request data")
	E00012 = New("E00012", "verification failed")
	E00013 = New("E00013", "%s has been already assigned")
	E00014 = New("E00014", "%s invalid")
	E00015 = New("E00015", "duplicated %s")
	E00016 = New("E00016", "%s invalid or expired")
	E00017 = New("E00017", "cannot be deleted because it is being used")
	E00018 = New("E00018", "cannot be process because invalid state")
	E00019 = New("E00019", "%s has already expired")
	E00020 = New("E00020", "%s period has not yet arrived")
	E00021 = New("E00021", "%s has already been used.")
	E00022 = New("E00022", "%s cannot update because not allow.")
	E00023 = New("E00023", "%s cannot remove because not allow.")
	E00024 = New("E00024", "%s cannot create because not allow.")
	E00025 = New("E00025", "%s not found or inactive")
	E00026 = New("E00026", "%s cannot be inactive because it is being used")
	E00027 = New("E00027", "%s is inactive")

	//Message
	M00001 = New("M00001", "successfully")
)
