package hw10programoptimization

import (
	"bufio"
	"io"
	"strings"

	"github.com/mailru/easyjson"
)

type User struct {
	ID       int
	Name     string
	Username string
	Email    string
	Phone    string
	Password string
	Address  string
}

type DomainStat map[string]int

func GetDomainStat(r io.Reader, domain string) (DomainStat, error) {
	result := make(DomainStat)
	domain = "." + strings.ToLower(domain)

	scanner := bufio.NewScanner(r)
	user := &User{}

	for scanner.Scan() {
		// if err := json.Unmarshal(scanner.Bytes(), user); err != nil {
		if err := easyjson.Unmarshal(scanner.Bytes(), user); err != nil {
			return nil, err
		}

		if strings.HasSuffix(strings.ToLower(user.Email), domain) {
			email := strings.ToLower(strings.SplitN(user.Email, "@", 2)[1])
			result[email]++
		}
	}

	return result, scanner.Err()
}
