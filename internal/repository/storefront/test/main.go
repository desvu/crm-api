package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-pg/pg/v9"
	"github.com/go-pg/pg/v9/orm"
)

type storefront struct {
	ID        uint      `pg:"id"`
	Name      string    `pg:"name"`
	IsActive  bool      `pg:"-"`
	CreatedAt time.Time `pg:"created_at,default:now()"`

	// relations
	Version version
	// Versions    []version
	// Activation  *activation
	// Activations []activation

	tableName struct{} `pg:"storefronts,alias:sf"`
}

type version struct {
	StorefrontID uint      `pg:"storefront_id,pk,fk"`
	ID           uint      `pg:"id,use_zero,pk"`
	Blocks       []string  `pg:"blocks"`
	CreatedAt    time.Time `pg:"created_at,default:now()"`

	// relations
	// Activation *activation

	tableName struct{} `pg:"storefront_versions"`
}

type activation struct {
	Timestamp    time.Time `pg:"timestamp,default:now(),pk"`
	VersionID    uint      `pg:"version_id,use_zero,fk"`
	StorefrontID uint      `pg:"storefront_id,fk"`

	tableName struct{} `pg:"storefront_activations"`
}

type dbLogger struct{}

func (d dbLogger) BeforeQuery(c context.Context, q *pg.QueryEvent) (context.Context, error) {
	return c, nil
}

func (d dbLogger) AfterQuery(c context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}
func main() {
	db := pg.Connect(&pg.Options{
		Addr:     "localhost:5567",
		User:     "qilin",
		Password: "password",
		Database: "qilin-store",
	})
	defer db.Close()

	db.AddQueryHook(dbLogger{})

	err := db.DropTable((*storefront)(nil), &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	})
	panicIf(err)

	err = db.CreateTable((*storefront)(nil), nil)
	panicIf(err)

	err = db.DropTable((*activation)(nil), &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	})
	panicIf(err)

	err = db.CreateTable((*activation)(nil), nil)
	panicIf(err)

	err = db.DropTable((*version)(nil), &orm.DropTableOptions{
		IfExists: true,
		Cascade:  true,
	})
	panicIf(err)

	err = db.CreateTable((*version)(nil), nil)
	panicIf(err)

	sf := &storefront{
		Name: "some",
		Version: version{
			ID:     0,
			Blocks: []string{"asd0", "sdf0"},
		},
	}
	db.Model(sf).Relation("Version").Insert()

	sf2 := &storefront{
		Name: "some2",
	}
	db.Model(sf2).Insert()

	// v10 := &version{
	// 	StorefrontID: sf.ID,
	// 	ID:           0,
	// 	Blocks:       []string{"asd0", "sdf0"},
	// }
	// db.Model(v10).Insert()
	v11 := &version{
		StorefrontID: sf.ID,
		ID:           1,
		Blocks:       []string{"asd1", "sdf1"},
	}
	db.Model(v11).Insert()
	v12 := &version{
		StorefrontID: sf.ID,
		ID:           2,
		Blocks:       []string{"asd2", "sdf2"},
	}
	db.Model(v12).Insert()
	v20 := &version{
		StorefrontID: sf2.ID,
		ID:           0,
		Blocks:       []string{"asd20", "sdf20"},
	}
	db.Model(v20).Insert()
	v21 := &version{
		StorefrontID: sf2.ID,
		ID:           1,
		Blocks:       []string{"asd21", "sdf21"},
	}
	db.Model(v21).Insert()

	act := &activation{
		StorefrontID: v20.StorefrontID,
		VersionID:    v20.ID,
	}
	db.Model(act).Insert()
	time.Sleep(time.Millisecond)
	act = &activation{
		StorefrontID: sf.ID,
		VersionID:    sf.Version.ID,
	}
	db.Model(act).Insert()
	time.Sleep(time.Millisecond)
	act = &activation{
		StorefrontID: 11,
		VersionID:    11,
	}
	db.Model(act).Insert()
	time.Sleep(time.Millisecond)
	act = &activation{
		StorefrontID: v11.StorefrontID,
		VersionID:    v11.ID,
	}
	db.Model(act).Insert()
	time.Sleep(time.Millisecond)
	act = &activation{
		StorefrontID: v21.StorefrontID,
		VersionID:    v21.ID,
	}
	db.Model(act).Insert()
	query(db)
	queryAll(db)

}

func query(db *pg.DB) {
	rsf := &storefront{}
	// select storefront_id from storefront_activations order by timestamp desc limit 1
	q := db.Model(rsf).
		Column("sf.*").
		ColumnExpr("((?) = sf.id) as is_active", db.Model((*activation)(nil)).Column("storefront_id").Order("timestamp desc").Limit(1)).
		Relation("Version", func(q *orm.Query) (*orm.Query, error) {
			return q.Order("version desc"), nil
		}).
		// Relation("Versions").
		// Relation("Activation", func(q *orm.Query) (*orm.Query, error) {
		// 	return q.Order("timestamp desc"), nil
		// }).
		// Relation("Versions").
		First()
	fmt.Println(q)

	fmt.Println("done")

	x := rsf
	// for _, x := range *rsf {
	fmt.Println("Name: ", x.Name)
	fmt.Println("Last version: ", x.Version.ID, x.Version.Blocks)
	// fmt.Println("Total versions: ", len(x.Versions))
	// fmt.Println("Last activation: ", x.Activation.Timestamp)
	fmt.Println("IsActive: ", x.IsActive)
	// }
}

func queryAll(db *pg.DB) {
	rsf := []storefront{}
	q := db.Model(&rsf).
		DistinctOn("sf.id").Order("sf.id").
		Column("sf.*").
		ColumnExpr("((?) = sf.id) as is_active", db.Model((*activation)(nil)).Column("storefront_id").Order("timestamp desc").Limit(1)).
		Relation("Version").Order("version desc").
		//, func(q *orm.Query) (*orm.Query, error) {
		// return q.Order("version desc"), nil
		// }).
		// Relation("Version.Activation", func(q *orm.Query) (*orm.Query, error) {
		// 	return q.Order("version__activation.timestamp desc"), nil
		// }).
		// Relation("Activation").Order("timestamp desc").
		// Relation("Version.Activation").Order("timestamp desc").
		// , func(q *orm.Query) (*orm.Query, error) {
		// return q.Order("timestamp desc"), nil
		// }).
		// Relation("Versions").
		// Relation("Versions.Activation", func(q *orm.Query) (*orm.Query, error) {
		// 	return q.Order("timestamp desc"), nil
		// }).
		Select()
	fmt.Println(q)

	fmt.Println("done")

	for _, x := range rsf {
		fmt.Println("Name: ", x.ID, x.Name)
		fmt.Println("Last version: ", x.Version.ID, x.Version.Blocks)
		// fmt.Println("Total versions: ", len(x.Versions))
		// fmt.Println("Last version activation: ", x.Version.Activation.Timestamp, x.Version.Activation.StorefrontID)
		// fmt.Println("Last activation: ", x.Activation.Timestamp)
		fmt.Println("IsActive: ", x.IsActive)
	}
}

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}
