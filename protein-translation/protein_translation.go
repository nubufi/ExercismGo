package protein

import "errors"

var codons = map[string]string {
	"AUG":"Methionine",
	"UUU":"Phenylalanine",
	"UUC":"Phenylalanine",
	"UUA":"Leucine",
	"UUG":"Leucine",
	"UCU":"Serine",
	"UCC":"Serine",
	"UCA":"Serine",
	"UCG":"Serine",
	"UAU":"Tyrosine",
	"UAC":"Tyrosine",
	"UGU":"Cysteine",
	"UGC":"Cysteine",
	"UGG":"Tryptophan",
	"UAA":"STOP",
	"UAG":"STOP",
	"UGA":"STOP",
}

var ErrStop = errors.New("STOP")
var ErrInvalidBase = errors.New("invalid base")

func FromRNA(rna string) ([]string, error) {
	var pros []string
	for i:=0;i<len(rna)/3;i++ {
		codon := rna[3*i:3*i+3]
		pro,err := FromCodon(codon)
		if err == ErrInvalidBase {
			return pros,err
		}
		if err == ErrStop {
			return pros,nil
		}
		pros = append(pros,pro)

	}

	return pros, nil
}


func FromCodon(codon string) (string, error) {
	pro,ok := codons[codon]

	if !ok {
		return "",ErrInvalidBase
	}

	if pro == "STOP" {
		return "",ErrStop
	}

	return pro,nil
}
