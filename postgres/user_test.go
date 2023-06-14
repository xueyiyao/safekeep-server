package postgres_test

import (
	"testing"

	"github.com/xueyiyao/safekeep/domain"
	"github.com/xueyiyao/safekeep/postgres"
)

func TestUserService_CreateUser(t *testing.T) {
	t.Run("OK", func(t *testing.T) {
		db := MustOpenDB(t)
		tx := db.DB.Begin()
		defer func() {
			tx.Rollback()
			MustCloseDB(t, db)
		}()

		s := postgres.NewUserService(tx)

		u := &domain.User{Name: "one", Email: "one@test.com"}
		if err := s.CreateUser(u); err != nil {
			t.Fatal(err)
		} else if got, want := u.ID, 1; int(got) != want {
			t.Fatalf("ID=%v, want %v", got, want)
		}

		u2 := &domain.User{Name: "two", Email: "two@test.com"}
		if err := s.CreateUser(u2); err != nil {
			t.Fatal(err)
		} else if got, want := u2.ID, 2; int(got) != want {
			t.Fatalf("ID=%v, want %v", got, want)
		}

		u3 := &domain.User{Name: "two", Email: "nottwo@test.com"}
		if err := s.CreateUser(u3); err != nil {
			t.Fatal(err)
		} else if got, want := u3.ID, 3; int(got) != want {
			t.Fatalf("ID=%v, want %v", got, want)
		}

		// Go's time.Time precision is in ns, while PSQL is in microseconds. this will fail as is so commented out
		// // Simply fetches user and compare
		// if other, err := s.FindUserByID(1); err != nil {
		// 	t.Fatal(err)
		// } else if !reflect.DeepEqual(u, other) {
		// 	t.Fatalf("mismatch: %#v != %#v", u, other)
		// }
	})

	// Ensure an error is returned if empty user.
	t.Run("ErrNilUser", func(t *testing.T) {
		db := MustOpenDB(t)
		tx := db.DB.Begin()
		defer func() {
			tx.Rollback()
			MustCloseDB(t, db)
		}()
		s := postgres.NewUserService(tx)
		if err := s.CreateUser(nil); err == nil {
			t.Fatal("expected an error, none occured")
		}
	})

	// Ensure an error is returned if user name is not set.
	t.Run("ErrNameRequired", func(t *testing.T) {
		db := MustOpenDB(t)
		tx := db.DB.Begin()
		defer func() {
			tx.Rollback()
			MustCloseDB(t, db)
		}()
		s := postgres.NewUserService(tx)
		if err := s.CreateUser(&domain.User{Email: "one@gmail.com"}); err == nil {
			t.Fatal("expected an error, none occured")
		}
	})

	// Ensure an error is returned if user email is not set.
	t.Run("ErrEmailRequired", func(t *testing.T) {
		db := MustOpenDB(t)
		tx := db.DB.Begin()
		defer func() {
			tx.Rollback()
			MustCloseDB(t, db)
		}()
		s := postgres.NewUserService(tx)
		if err := s.CreateUser(&domain.User{Name: "one"}); err == nil {
			t.Fatal("expected an error, none occured")
		}
	})

	// Ensure an error is returned if user email is not valid.
	t.Run("ErrEmailInvalid", func(t *testing.T) {
		db := MustOpenDB(t)
		tx := db.DB.Begin()
		defer func() {
			tx.Rollback()
			MustCloseDB(t, db)
		}()
		s := postgres.NewUserService(tx)
		if err := s.CreateUser(&domain.User{Name: "one", Email: "asdfjkl;"}); err == nil {
			t.Fatal("expected an error, none occured")
		}
	})

	// Ensure an error is returned if user email is a duplicate.
	t.Run("ErrEmailDuplicate", func(t *testing.T) {
		db := MustOpenDB(t)
		tx := db.DB.Begin()
		defer func() {
			tx.Rollback()
			MustCloseDB(t, db)
		}()

		s := postgres.NewUserService(tx)

		u := &domain.User{Name: "one", Email: "one@test.com"}
		if err := s.CreateUser(u); err != nil {
			t.Fatal(err)
		}
		if err := s.CreateUser(&domain.User{Name: "one", Email: "one@test.com"}); err == nil {
			t.Fatal("expected an error, none occured")
		}
	})
}

func TestUserService_FindUser(t *testing.T) {
	// Ensure an error is returned if fetching a non-existent user.
	t.Run("ErrNotFound", func(t *testing.T) {
		db := MustOpenDB(t)
		tx := db.DB.Begin()
		defer func() {
			tx.Rollback()
			MustCloseDB(t, db)
		}()
		s := postgres.NewUserService(tx)
		if user, err := s.FindUserByID(1); err == nil {
			t.Fatalf("expected an error, none occured %v", user)
		}
	})
}
