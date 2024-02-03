package gomaith_test

import (
	"fmt"
	"testing"

	"github.com/johnstcn/gomaith"
	"github.com/stretchr/testify/assert"
)

func ExampleVocative() {
	fmt.Println(gomaith.Vocative("Seán"))
	// Output: a Sheáin
}

func TestVocative(t *testing.T) {
	for _, tt := range []struct {
		in, want string
	}{
		{"", ""},

		// Non-Irish names
		{"Albert", "a Albert"},
		{"Benjamin", "a Benjamin"},
		{"Charlie", "a Charlie"},
		{"David", "a David"},
		{"Edward", "a Edward"},
		{"Frank", "a Frank"},
		{"Godfrey", "a Godfrey"},
		{"Harry", "a Harry"},
		{"Isabella", "a Isabella"},
		{"John", "a John"},
		{"Kevin", "a Kevin"},
		{"Larry", "a Larry"},
		{"Mary", "a Mary"},
		{"Nancy", "a Nancy"},
		{"Oliver", "a Oliver"},
		{"Quincy", "a Quincy"},
		{"Robert", "a Robert"},
		{"Steve", "a Steve"},
		{"Theresa", "a Theresa"},
		{"Ulysses", "a Ulysses"},
		{"Victor", "a Victor"},
		{"William", "a William"},
		{"Xavier", "a Xavier"},
		{"Yancy", "a Yancy"},
		{"Zachary", "a Zachary"},

		// Non-Irish names that can behave like Irish names.
		{"Adam", "a Adaim"},
		{"Bertha", "a Bhertha"},
		{"Bob", "a Bhoib"},
		{"Conor", "a Chonoir"},
		{"George", "a Gheorge"},
		{"Isaac", "a Isaaic"},
		{"Paul", "a Phauil"},
		{"Tom", "a Thoim"},

		// Irish names
		{"Aodh", "a Aoidh"},
		{"Aodhán", "a Aodháin"},
		{"Aonghus", "a Aonghuis"},
		{"Breandán", "a Bhreandáin"},
		{"Brian", "a Bhriain"},
		{"Cathal", "a Chathail"},
		{"Caoimhe", "a Chaoimhe"},
		{"Cillian", "a Chilliain"},
		{"Cian", "a Chiain"},
		{"Ciarán", "a Chiaráin"},
		{"Conchúr", "a Chonchúir"},
		{"Cormac", "a Chormaic"},
		{"Darragh", "a Dharraigh"},
		{"Diarmuid", "a Dhiarmuid"},
		{"Donnchadh", "a Dhonnchaidh"},
		{"Eoghan", "a Eoghain"},
		{"Fergus", "a Fherguis"},
		{"Fionn", "a Fhioinn"},
		{"Fionnán", "a Fhionnáin"},
		{"Fionnuala", "a Fhionnuala"},
		{"Gearód", "a Ghearóid"},
		{"Gráinne", "a Ghráinne"},
		{"Liam", "a Liaim"},
		{"Máire", "a Mháire"},
		{"Seán", "a Sheáin"},
		{"Séamus", "a Shéamuis"},
		{"Seosamh", "a Sheosaimh"},
		{"Síle", "a Shíle"},
	} {
		tt := tt
		name := fmt.Sprintf("%s --> %s", tt.in, tt.want)
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			got := gomaith.Vocative(tt.in)
			assert.Equal(t, tt.want, got)
		})
	}
}
