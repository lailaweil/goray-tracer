package main

import (
	"fmt"
	genericError "github.com/laiweil/goray-tracer/errors"
	"github.com/laiweil/goray-tracer/models"
	"os"
)

type Projectile struct {
	position *models.Tuple
	velocity *models.Tuple
}

type Environment struct {
	gravity *models.Tuple
	wind    *models.Tuple
}

func main() {
	ticks := 0

	canvas := models.CreateCanvas(1080, 720)
	red := models.Color(0.35, -0.35, 1.9)

	projectile := &Projectile{
		position: models.Point(0, 1, 0),
		velocity: models.Vector(1, 1.8, 0).Normalize().Multiply(12),
	}

	environment := &Environment{
		gravity: models.Vector(0, -0.1, 0),
		wind:    models.Vector(-0.0, 0, 0),
	}

	canvas.WritePixel(int(projectile.position[models.X]), canvas.Height()-int(projectile.position[models.Y]), *red)

	for projectile.position[models.Y] > 0 {
		ticks++
		projectile = Tick(environment, projectile)

		err := canvas.WritePixel(int(projectile.position[models.X]), canvas.Height()-int(projectile.position[models.Y]), *red)

		if err != nil {
			fmt.Sprintf("Error Message %s, err: %s", err.(*genericError.GenericError).Message(), err.Error())
		}

		fmt.Printf("Projectile position is %v\n", projectile.position)

	}

	var f *os.File
	var err error
	if f, err = os.Create("test.ppm"); err != nil {
		return
	}
	if err = canvas.WritePPM(f); err != nil {
		return
	}

	f.Close()
}

func Tick(env *Environment, prj *Projectile) *Projectile {
	return &Projectile{
		position: prj.position.Add(*prj.velocity),
		velocity: prj.velocity.Add(*env.gravity).Add(*env.wind),
	}
}
