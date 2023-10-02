package handlers

import (
	"fmt"
	"time"

	"github.com/emarifer/go-docker-trivia/database"
	"github.com/emarifer/go-docker-trivia/models"
	"github.com/gofiber/fiber/v2"
)

// This handler is no longer in use
/* func Home(c *fiber.Ctx) error {
	return c.SendString("Hello, Enrique from Go Trivia App!!")
} */

var year = time.Now().Year()

func ListFact(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Sp.DB.From("facts").Select("*").Execute(&facts)

	// return c.Status(fiber.StatusOK).JSON(facts) // for API REST
	return c.Render("index", fiber.Map{
		"PageTitle": "Gopher Trivia Time",
		"Title":     "Gopher Trivia Time",
		"Subtitle":  "Facts for funtimes with friends!",
		"Facts":     facts,
		"Year":      year,
	})
}

func NewFactView(c *fiber.Ctx) error {
	return c.Render("new", fiber.Map{
		"PageTitle": "Create New Fact",
		"Title":     "New Fact",
		"Subtitle":  "Add a cool fact!",
		"Year":      year,
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		/* return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		}) */
		return NewFactView(c)
	}

	facts := []models.Fact{}
	err := database.DB.Sp.DB.From("facts").Insert(fact).Execute(&facts)
	if err != nil {
		return NewFactView(c)
	}

	// return c.Status(fiber.StatusOK).JSON(fact)
	// return ConfirmationView(c)
	// return ListFact(c)
	return c.Redirect("/")
	// Es necesario redirigir, porque si no cambia la url
	// si el usuario recarga la página se producirá un reenvío
	// del formulario y se duplicará la entrada
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	err := database.DB.Sp.DB.From("facts").Select("*").Single().Eq("id", id).Execute(&fact)
	if err != nil {
		return NotFound(c)
	}

	return c.Render("show", fiber.Map{
		"PageTitle": fmt.Sprintf("Show Fact #%s", id),
		"Title":     "Single Fact",
		"Fact":      fact,
		"Year":      year,
	})
}

func EditFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	err := database.DB.Sp.DB.From("facts").Select("*").Single().Eq("id", id).Execute(&fact)
	if err != nil {
		return NotFound(c)
	}

	return c.Render("edit", fiber.Map{
		"PageTitle": fmt.Sprintf("Edit Fact #%s", id),
		"Title":     "Edit Fact",
		"Subtitle":  fmt.Sprintf("Edit your interesting fact #%s", id),
		"Fact":      fact,
		"Year":      year,
	})
}

func UpadteFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	if err := c.BodyParser(&fact); err != nil {
		return c.Status(fiber.StatusServiceUnavailable).SendString(err.Error())
	}

	err := database.DB.Sp.DB.From("facts").Update(&fact).Eq("id", id).Execute(&fact)
	if err != nil {
		return EditFact(c)
	}

	return ShowFact(c)
}

func DeleteFact(c *fiber.Ctx) error {
	facts := []models.Fact{}
	id := c.Params("id")

	err := database.DB.Sp.DB.From("facts").Delete().Eq("id", id).Execute(&facts)
	if err != nil {
		return NotFound(c)
	}

	return ListFact(c)
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
}

// This handler is no longer in use
/* func ConfirmationView(c *fiber.Ctx) error {
	return c.Render("confirmation", fiber.Map{
		"PageTitle": "Confirmation Page",
		"Title":     "Fact added successfully!",
		"Subtitle":  "Add more wonderful facts to the list!",
	})
} */

/* DOCUMENTACIÓN supabase-go
https://github.com/nedpals/supabase-go
*/
