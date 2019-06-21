package seed

import (
	"github.com/jinzhu/gorm"
	m "github.com/nicolasjhampton/go-course-api/internal/database/models"
)

var courses = []m.Course{
	m.Course{
		UserID: 1,
		Title:  "Build a Basic Bookcase",
		Description: `High-end furniture projects are great to dream 
		about. But unless you have a well-equipped shop and some 
		serious woodworking experience to draw on, it can be difficult 
		to turn the dream into a reality.\n\nNot every piece of 
		furniture needs to be a museum showpiece, though.`,
		Time: "12 hours",
		Materials: `* 1/2 x 3/4 inch parting strip\n* 1 x 2 common pine\n
		* 1 x 4 common pine\n* 1 x 10 common pine\n* 1/4 inch thick lauan plywood\n
		* Finishing Nails\n* Sandpaper\n* Wood Glue\n* Wood Filler\n
		* Minwax Oil Based Polyurethane\n`,
		Version: "0",
		Reviews: []m.Review{
			m.Review{
				UserID:  1,
				Rating:  5,
				Review:  "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat.",
				Version: "0",
			},
			m.Review{
				UserID:  2,
				Rating:  3,
				Review:  "Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				Version: "0",
			},
		},
		Steps: []m.Step{
			m.Step{
				Title:       "Yakkety smackity...",
				Description: "Here's some things...",
			},
			m.Step{
				Title:       "Blah Blah Blah",
				Description: "And some other stuff...",
			},
		},
	},
	m.Course{
		UserID:      2,
		Title:       "Learn How to Program",
		Description: "In this course, you'll learn how to write code like a pro!",
		Time:        "6 hours",
		Materials:   "* Notebook computer running Mac OS X or Windows\n* Text editor",
		Version:     "0",
		Reviews: []m.Review{
			m.Review{
				UserID:  1,
				Rating:  5,
				Review:  "Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
				Version: "0",
			},
		},
		Steps: []m.Step{
			m.Step{
				Title:       "Cutting the Parts",
				Description: "For precise crosscuts, first make a simple, self-aligning T-guide for your circular saw. Cut a piece of 1/2-in. plywood to 2 1/2 x 24 in. and glue and screw it to a roughly 12-in.-long piece of 1 x 4 pine that will serve as the crossbar of the T. Center the plywood strip along the 1 x 4 and make sure the pieces are perfectly square to each other.\n\nButt the crossbar of the T-guide against the edge of a piece of scrap lumber, tack the guide in place and make a cut through the 1 x 4 with your saw base guided by the plywood strip. Then, trim the 1 x 4 on the opposite side in the same way. Now, the ends of the 1 x 4 can be aligned with layout lines on the stock for precise cut positioning.\n\nBegin construction by using a tape measure to mark the length of a side panel on 1 x 10 stock, and lay out the cut line with a square. The side panels on our bookcase are 48 in. long.",
			},
		},
	},
}

func courseseeds(tx *gorm.DB) (err error) {
	for _, course := range courses {
		if err = tx.Create(&course).Error; err != nil {
			tx.Rollback()
			return
		}
	}
	return
}
