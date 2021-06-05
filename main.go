package main

import (
	"fmt"
	"github.com/laiweil/goray-tracer/models"
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

	projectile := &Projectile{
		position: models.Point(0, 1, 0),
		velocity: models.Vector(1, 1, 0).Normalize(),
	}

	environment := &Environment{
		gravity: models.Vector(0, -0.1, 0),
		wind:    models.Vector(-0.01, 0, 0),
	}

	for projectile.position.Y > 0 {
		ticks++
		projectile = Tick(environment, projectile)

		fmt.Printf("Projectile position is %v\n", projectile.position)

	}

	fmt.Printf("Projectile hitted the ground after %d Ticks\n", ticks)

}

func Tick(env *Environment, prj *Projectile) *Projectile {
	return &Projectile{
		position: prj.position.Add(*prj.velocity),
		velocity: prj.velocity.Add(*env.gravity).Add(*env.wind),
	}
}
