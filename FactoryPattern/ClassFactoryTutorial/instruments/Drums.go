package instruments

type Drums struct {
	//appliances.Appliance
	typeName string
	sound    string
}

func (mr *Drums) Start() {
	mr.typeName = " drums "
}

func (mr *Drums) GetPurpose() string {
	return "I am a " + mr.typeName + " I heat stuff up!!"
}

func (mr *Drums) PlayMusic() string {
	mr.sound = "The Beat"
	return "I play " + mr.typeName + " The Beat!!"
}
