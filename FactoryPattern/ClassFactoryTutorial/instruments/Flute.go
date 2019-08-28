package instruments

type Flute struct {
	//appliances.Appliance
	typeName string
	sound    string
}

func (mr *Flute) Start() {
	mr.typeName = " Flute "
}

func (mr *Flute) GetPurpose() string {
	return "I am a " + mr.typeName + " I heat stuff up!!"
}

func (mr *Flute) PlayMusic() string {
	mr.sound = "The Beat"
	return "I play " + mr.typeName + mr.sound
}
