package instruments

type Guitar struct {
	//appliances.Appliance
	typeName string
	sound    string
}

func (mr *Guitar) Start() {
	mr.typeName = " Guitar "
}

func (mr *Guitar) GetPurpose() string {
	return "I am a " + mr.typeName + "I rock!!"
}

func (mr *Guitar) PlayMusic() string {
	mr.sound = "The Rytthm"
	return "I play " + mr.typeName + mr.sound
}
