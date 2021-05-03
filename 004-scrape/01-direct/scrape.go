package main

import (
	"encoding/json"
	"fmt"

	"github.com/gocolly/colly"
)

const site = "https://www.reddit.com/user/JetBrains_official/comments/fcb7mr/goland_is_a_constantly_evolving_ide_with_a_deep/?instanceId=t3_z%3DgAAAAABgU3EwXgTr9DuHm8NvHID_n4jwTVQJpC1_zK3-ycPKRx0lAFp7ayXPqlpjmvznO7oF7k7pBkY-ZXo0dz6_e0THsBnmCETMFCoAZnfDyrVxCp5FEPAeWLPRjLq6gRnj-YYEOWlr39pNNzfn-RVn-Wi680A0vtZEZBA_SA2mUyqhKWEJSeNhT8He9tTCXrGngpqSt3kC-BX1b6AhTn14i2b1yUZdZeLph1Qa26ydW-U3o7Zld51D7oGRSzYqrRwP6DEJe3luNl5gLYEepU90YlKX_x2dBPJikVNBEOUdYFWA3YZQ9Q7Pyh6mNzJFAlSes29Vm4xQrUaaFzouPnUB5GOcgwpeBp_FeCwYRtY7WviaC2zOpJtGcvLAun2Qdxdd5dSlh30nsp1nINiDjjRrTn18ZacumDfPy3_usYxMDoitZtknAJtJjZPP4WsPgDYdrfyDR37AMm_l3_ynbN47wTc6EscXGYDs5VWYPdsY-FiVdAszg_RxdHYhENRMxFMyWNcJGU_xLp-IQLzvzZ_G3zJjS5mfukIWnFSKLAZTpLLu29gSxCxI8_S9to84SsljIERS3kuZcC0kPTUETki2_s9gVnUNa7QsR769rxZv8haRb6s7JFRVR1tLzTBtswmSCMFTKy_ZR8HOr4XoSzilUAlBW8hIkTtL0BzphMNcN5B-tNJbc1z-0hkamKwmr4Y7BB1tL-fLNc9MwWtx1NA2PmihhLikYYXkgDm6UQ-60Hgw2Q3WR3-hCUX0Tp8ON_FPlG1RLVlulaXeYbEuTD5LsWBgmd__voD-31c5jmKssTPo6z2igj2w1Xvr6Aisv-poMw7M5-8lwluBws_olwrijpoIDitE5bwLFttkLTBokkKHGuapgFuZH5hV8okR55aLyiKjRt4VtT4zpg0Weaw3PaLtXTltUW_kBqGQaCNC0LToVahuZv0%3D"

type comments struct {
	Name    string
	TimeAgo string
	Message string
}

func main() {
	c := colly.NewCollector()

	messages := []comments{}

	c.OnHTML("._3tw__eCCe7j-epNCKGXUKk", func(e *colly.HTMLElement) {
		messages = append(messages, comments{
			Name:    e.ChildText(".f3THgbzMYccGW8vbqZBUH ._23wugcdiaj44hdfugIAlnX"),
			TimeAgo: e.ChildText("._1sA-1jNHouHDpgCp1fCQ_F"),
			Message: e.ChildText("._1qeIAgB0cPwnLhDF9XSiJM"),
		})
	})

	err := c.Visit(site)
	if err != nil {
		panic(err)
	}

	c.Wait()

	bs, err := json.MarshalIndent(messages, "", "\t")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(bs))
	fmt.Println("Number of tweets:", len(messages))
}
