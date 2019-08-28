package hydradblayer

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkMySQLDBReads(b *testing.B) {
	dblayer, err := ConnectDatabase("mysql", "hydea:hydraisme@/Hydra")
	if err != nil {
		b.Fatal("Could not connect to hydra chat system", err)
	}

	allMembersBM(b, dblayer)
}

func BenchmarkMongoDBReads(b *testing.B) {
	dblayer, err := ConnectDatabase("mongodb", "mongodb://127.0.0.1")
	if err != nil {
		b.Error("Could not connect to hydra chat system", err)
		return
	}

	allMembersBM(b, dblayer)
	findMembersBM(b, dblayer)
}

func allMembersBM(b *testing.B, dblayer DBLayer) {

	for i := 0; i < b.N; i++ {
		_, err := dblayer.AllMembers()
		if err != nil {
			b.Error("Query failed ", err)
			return
		}
	}
}

func findMembersBM(b *testing.B, dblayer DBLayer) {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < b.N; i++ {
		_, err := dblayer.FindMember(rand.Intn(16) + 1)
		if err != nil {
			b.Error("The db Query failed ", err)
			return
		}
	}

}
