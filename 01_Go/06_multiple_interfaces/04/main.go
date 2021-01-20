// Go program to illustrate embedding interfaces
package main

import "fmt"

type AuthorDetails interface {
	details()
}

type AuthorArticles interface {
	articles()
	picked()
}
type FinalDetails interface {
	details()
	AuthorArticles
	cdeatils()
}

type author struct {
	a_name    string
	branch    string
	college   string
	year      int
	salary    int
	particles int
	tarticles int
	cid       int
	post      string
	pick      int
}

func (a author) details() {

	fmt.Printf("Author Name: %s", a.a_name)
	fmt.Printf("\nBranch: %s and passing year: %d", a.branch, a.year)
	fmt.Printf("\nCollege Name: %s", a.college)
	fmt.Printf("\nSalary: %d", a.salary)
	fmt.Printf("\nPublished articles: %d", a.particles)
}

func (a author) articles() {

	pendingarticles := a.tarticles - a.particles
	fmt.Printf("\nPending articles: %d", pendingarticles)
}

func (a author) picked() {

	fmt.Printf("\nTotal number of picked articles: %d", a.pick)
}

func (a author) cdeatils() {

	fmt.Printf("\nAuthor Id: %d", a.cid)
	fmt.Printf("\nPost: %s", a.post)
}

func main() {
	values := author{
		a_name:    "Mickey",
		branch:    "Computer science",
		college:   "XYZ",
		year:      2012,
		salary:    50000,
		particles: 209,
		tarticles: 309,
		cid:       3087,
		post:      "Technical content writer",
		pick:      58,
	}
	var f FinalDetails = values
	f.details()
	f.articles()
	f.picked()
	f.cdeatils()
}
