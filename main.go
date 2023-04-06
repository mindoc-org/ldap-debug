package main

import (
	"fmt"
	"github.com/go-ldap/ldap/v3"
	"log"
	"net"
	"time"
)

var DefaultTimeout = 5 * time.Second

func demo(ldapURL string) {
	lc, err := ldap.DialURL(ldapURL, ldap.DialWithDialer(&net.Dialer{Timeout: DefaultTimeout}))
	if err != nil {
		log.Fatalf("DialURL => %v", err)
		return
	}
	defer lc.Close()

	// filter := fmt.Sprintf("(&(objectClass=User)(%s=%s))", "mail", "admin@example.com")
	filter := "(&(objectClass=User))"  // 【5】 查询的筛选条件
	attributes := []string{"dn", "cn"} // 【6】 查询的包含属性
	// baseDn := "DC=example,DC=com"
	baseDn := "DC=example,DC=com" // 【2】 DC
	searchRequest := ldap.NewSearchRequest(
		baseDn, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		filter,     // The filter to apply
		attributes, // A list attributes to retrieve
		nil,
	)
	// account := "cn=admin,dc=example,dc=com"
	// account := "cn=administrator,dc=example,dc=com"
	account := "administrator@example.com" // 【3】 账号
	// password := "123456"
	password := "123+456=579" // 【4】 密码
	err = lc.Bind(account, password)
	if err != nil {
		log.Fatalf("Bind => %v", err)
		return
	}

	sr, err := lc.Search(searchRequest)
	if err != nil {
		log.Fatalf("Search => %v", err)
	}

	for _, entry := range sr.Entries {
		log.Printf("%s: %v\n", entry.DN, entry.GetAttributeValue("cn"))
	}
}

func main() {
	log.Println("Debug go-ldap:")

	ldapURL := fmt.Sprintf("ldap://%s:%d", "10.96.8.9", 389) // 【1】 地址

	demo(ldapURL)
}
