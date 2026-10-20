package services

import (
	"testing"

	"github.com/go-ldap/ldap/v3"

	"catalina-support/backend/shared/config"
)

func TestObjectGUIDBinarioSeConvierteAlFormatoCanonico(t *testing.T) {
	raw := []byte{0x33, 0x22, 0x11, 0x00, 0x55, 0x44, 0x77, 0x66, 0x88, 0x99, 0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff}
	if got := guidCanonico(raw); got != "00112233-4455-6677-8899-aabbccddeeff" {
		t.Fatalf("GUID canónico inesperado: %q", got)
	}
	if got := guidCanonico(raw[:15]); got != "" {
		t.Fatalf("un objectGUID incompleto no debe persistirse: %q", got)
	}
}

func TestCuentaLDAPConservaIdentificadorTextual(t *testing.T) {
	cliente := NewLDAP(config.Directory{AttrUsername: "uid", AttrEmail: "mail", AttrName: "givenName", AttrLastName: "sn", AttrID: "uid"})
	entrada := &ldap.Entry{DN: "uid=ana,dc=example", Attributes: []*ldap.EntryAttribute{
		{Name: "uid", Values: []string{"ana"}}, {Name: "mail", Values: []string{"ana@example.com"}},
		{Name: "givenName", Values: []string{"Ana"}}, {Name: "sn", Values: []string{"Pérez"}},
	}}
	cuenta := cliente.cuenta(entrada)
	if cuenta.ExternalID != "ana" || cuenta.Username != "ana" {
		t.Fatalf("identidad textual inesperada: %#v", cuenta)
	}
}

func TestFiltroLDAPCombinaBaseUsuarioYCorreoSinInterpolarTexto(t *testing.T) {
	cliente := NewLDAP(config.Directory{
		UserFilter: "(&(objectClass=user)(!(disabled=true)))", AttrUsername: "sAMAccountName", AttrEmail: "mail",
	})
	filtro, atributo := cliente.filtroDeBusqueda(`jcestevez*)(mail=*)`, false)
	esperado := `(&(&(objectClass=user)(!(disabled=true)))(|(sAMAccountName=jcestevez\2a\29\28mail=\2a\29)(mail=jcestevez\2a\29\28mail=\2a\29)))`
	if atributo != "sAMAccountName" || filtro != esperado {
		t.Fatalf("filtro LDAP inesperado: %q (%s)", filtro, atributo)
	}
	parcial, _ := cliente.filtroDeBusqueda("jce", true)
	if parcial != `(&(&(objectClass=user)(!(disabled=true)))(|(sAMAccountName=*jce*)(mail=*jce*)))` {
		t.Fatalf("filtro parcial inesperado: %q", parcial)
	}
}
